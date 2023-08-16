//
// Copyright (c) 2023 F5 Inc. All rights reserved.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_registration "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/registration"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/svcfw"
)

const (
	listRegRPCFQN = "ves.io.schema.registration.CustomAPI.ListRegistrationsByState"
	listRegURI    = "/public/namespaces/system/listregistrationsbystate"

	regApprovalRPCFQN = "ves.io.schema.registration.CustomAPI.RegistrationApprove"
	regApprovalURI    = "/public/namespaces/system/registration/%s/approve"
)

type approvalParams struct {
	clusterName     string
	clusterSize     int32
	latitude        float32
	longitude       float32
	connectedRegion string
	hostname        string
}

// resourceVolterraRegistrationApproval is implementation of Volterra's Node Registration Approval Resource
func resourceVolterraRegistrationApproval() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraRegistrationApprovalCreate,
		Read:   resourceVolterraRegistrationApprovalRead,
		Update: resourceVolterraRegistrationApprovalUpdate,
		Delete: resourceVolterraRegistrationApprovalDelete,

		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"latitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"connected_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"retry": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"wait_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"registration_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
					"SITE_TO_SITE_TUNNEL_IPSEC",
					"SITE_TO_SITE_TUNNEL_SSL",
				}, false),
			},
		},
	}
}

// resourceVolterraRegistrationApprovalCreate approves registration resource
func resourceVolterraRegistrationApprovalCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	aParams := &approvalParams{}
	// validation
	if v, ok := d.GetOk("cluster_name"); ok {
		aParams.clusterName = v.(string)
	}
	if v, ok := d.GetOk("cluster_size"); ok {
		aParams.clusterSize = int32(v.(int))
	}
	if v, ok := d.GetOk("latitude"); ok {
		aParams.latitude = float32(v.(float64))
	}
	if v, ok := d.GetOk("longitude"); ok {
		aParams.longitude = float32(v.(float64))
	}
	if v, ok := d.GetOk("connected_region"); ok {
		aParams.connectedRegion = v.(string)
	}
	if v, ok := d.GetOk("hostname"); ok {
		aParams.hostname = v.(string)
	}

	var retry, waitTime int32
	if v, ok := d.GetOk("retry"); ok {
		retry = int32(v.(int))
	}
	if v, ok := d.GetOk("wait_time"); ok {
		waitTime = int32(v.(int))
	}

	approvalReq := &ves_io_schema_registration.ApprovalReq{
		Namespace: svcfw.SystemNSVal,
		State:     ves_io_schema_registration.APPROVED,
	}
	if v, ok := d.GetOk("tunnel_type"); ok {
		if tunType, ok := ves_io_schema.SiteToSiteTunnelType_value[v.(string)]; ok {
			approvalReq.TunnelType = ves_io_schema.SiteToSiteTunnelType(tunType)
		}
	}

	// Get list of all PENDING registrations
	listRegReq := &ves_io_schema_registration.ListStateReq{
		Namespace: svcfw.SystemNSVal,
		State:     ves_io_schema_registration.PENDING,
	}

	yamlReq, err := codec.ToYAML(listRegReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	for i := 0; i < int(retry); i++ {
		err := findRegistration(client, approvalReq, aParams, yamlReq)
		if err != nil {
			return err
		}
		if approvalReq.Name != "" {
			break
		} else {
			time.Sleep(time.Duration(rand.Int31n(waitTime)) * time.Second)
		}
	}

	if approvalReq.Name == "" {
		return fmt.Errorf("Registration %s does not exist", aParams.clusterName)
	}

	log.Printf("[DEBUG] Approving Volterra Registration object with struct: %+v", approvalReq)

	yamlReq, err = codec.ToYAML(approvalReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	rspProto, err := client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(regApprovalURI, approvalReq.Name), regApprovalRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error approving registration: %s", err)
	}
	rspApp := rspProto.(*ves_io_schema_registration.ObjectChangeResp)
	d.SetId(rspApp.Obj.SystemMetadata.GetUid())
	d.Set("registration_name", rspApp.Obj.Metadata.Name)

	return resourceVolterraRegistrationApprovalRead(d, meta)
}

func resourceVolterraRegistrationApprovalRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	regName := d.Get("registration_name").(string)
	_, err := client.GetObject(context.Background(), ves_io_schema_registration.ObjectType, svcfw.SystemNSVal, regName)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Registration %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Registration %q: %s", d.Id(), err)
	}
	return nil
}

func resourceVolterraRegistrationApprovalUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceVolterraRegistrationApprovalRead(d, meta)
}

func resourceVolterraRegistrationApprovalDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	regName := d.Get("registration_name").(string)
	aParams := &approvalParams{}
	// validation
	if v, ok := d.GetOk("cluster_name"); ok {
		aParams.clusterName = v.(string)
	}
	if v, ok := d.GetOk("cluster_size"); ok {
		aParams.clusterSize = int32(v.(int))
	}
	approvalReq := &ves_io_schema_registration.ApprovalReq{Name: regName, Namespace: svcfw.SystemNSVal, State: ves_io_schema_registration.RETIRED,
		Passport: &ves_io_schema_registration.Passport{
			ClusterName: aParams.clusterName,
			ClusterSize: aParams.clusterSize,
		},
	}

	log.Printf("[DEBUG] Deleting Volterra Registration obj with name %+v in namespace %+v", approvalReq.Name, approvalReq.Namespace)
	yamlReq, err := codec.ToYAML(approvalReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(regApprovalURI, approvalReq.Name), regApprovalRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Registration Approval %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
	}
	return err
}

func findRegistration(client *APIClient, approvalReq *ves_io_schema_registration.ApprovalReq,
	aParams *approvalParams, yamlReq string) error {

	rspProto, err := client.CustomAPI(context.Background(), http.MethodPost, listRegURI, listRegRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Unable to list registrations by state: %s", err)
	}
	rsp := rspProto.(*ves_io_schema_registration.ListResponse)
	for _, it := range rsp.Items {
		if it.Object != nil || it.Object.Spec != nil || it.Object.Spec.GcSpec != nil || it.Object.Spec.GcSpec.Passport != nil || it.Object.Spec.GcSpec.Infra != nil {
			if it.Object.Spec.GcSpec.Passport.ClusterName == aParams.clusterName && it.Object.Spec.GcSpec.Infra.Hostname == aParams.hostname {
				approvalReq.Passport = it.Object.Spec.GcSpec.Passport
				approvalReq.Passport.ClusterSize = aParams.clusterSize

				if aParams.latitude != 0 {
					approvalReq.Passport.Latitude = aParams.latitude
				}
				if aParams.longitude != 0 {
					approvalReq.Passport.Longitude = aParams.longitude
				}
				if aParams.connectedRegion != "" {
					approvalReq.ConnectedRegion = aParams.connectedRegion
				}
				approvalReq.Name = it.Object.Metadata.Name
				break
			}
		}
	}
	return nil
}
