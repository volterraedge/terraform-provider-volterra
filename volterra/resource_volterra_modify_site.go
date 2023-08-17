//
// Copyright (c) 2023 F5 Inc. All rights reserved.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"
	"gopkg.volterra.us/stdlib/server"

	"github.com/gogo/protobuf/proto"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
)

// resourceVolterraModifySite is implementation of Volterra's Site resources
func resourceVolterraModifySite() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraModifySiteCreate,
		Read:   resourceVolterraModifySiteRead,
		Update: resourceVolterraModifySiteUpdate,
		Delete: resourceVolterraModifySiteDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"bgp_peer_address": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"bgp_router_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"coordinates": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"latitude": {
							Type:     schema.TypeFloat,
							Optional: true,
						},

						"longitude": {
							Type:     schema.TypeFloat,
							Optional: true,
						},
					},
				},
			},

			"desired_pool_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"inside_nameserver": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"inside_vip": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"operating_system_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volterra_software_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"outside_nameserver": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"outside_vip": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"site_to_site_network_type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"site_to_site_tunnel_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"tunnel_dead_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"tunnel_type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"vip_vrrp_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"volterra_software_overide": {
				Type:     schema.TypeString,
				Optional: true,
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
		},
	}
}

// resourceVolterraModifySiteCreate creates Site resource
func resourceVolterraModifySiteCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceVolterraModifySiteUpdate(d, meta)
}

func resourceVolterraModifySiteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := getVolterraSite(client, name, namespace, server.GetSpecForm)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Site %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Site %q: %s", d.Id(), err)
	}

	return setSiteFields(client, d, resp)
}

func getVolterraSite(client *APIClient, name, namespace string, rspFmt server.ObjectFormat) (*ves_io_schema_site.GetResponse, error) {
	msgFQN := strings.Replace(ves_io_schema_site.ObjectType, "Object", "GetResponse", 1)
	protoMsgType := proto.MessageType(msgFQN).Elem()
	protoMsg := reflect.New(protoMsgType).Interface().(proto.Message)
	callRsp := &server.CallResponse{
		ProtoMsg: protoMsg,
	}
	opts := []vesapi.CallOpt{
		vesapi.WithResponseFormat(rspFmt),
		vesapi.WithOutCallResponse(callRsp),
	}

	_, err := client.GetObject(context.Background(), ves_io_schema_site.ObjectType, namespace, name, opts...)
	if err != nil {
		return nil, err
	}
	resp := callRsp.ProtoMsg.(*ves_io_schema_site.GetResponse)
	return resp, nil
}

func setSiteFields(client *APIClient, d *schema.ResourceData, resp *ves_io_schema_site.GetResponse) error {
	metadata := resp.GetMetadata()

	log.Printf("[DEBUG] REsponse : %+v", resp)

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	d.SetId(resp.SystemMetadata.Uid)

	return nil
}

// resourceVolterraModifySiteUpdate updates Site resource
func resourceVolterraModifySiteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)
	var retry, waitTime int32
	if v, ok := d.GetOk("retry"); ok {
		retry = int32(v.(int))
	}
	if v, ok := d.GetOk("wait_time"); ok {
		waitTime = int32(v.(int))
	}

	for i := 0; i < int(retry); i++ {
		resp, err := getVolterraSite(client, name, namespace, server.ReplaceRequestForm)
		if err != nil {
			if strings.Contains(err.Error(), "status code 404") {
				log.Println("[INFO] Site doesn't exist, retrying")
				time.Sleep(time.Duration(rand.Int31n(waitTime)) * time.Second)
				continue
			}
			return err
		}
		updateReq := getSiteReplaceForm(d, resp)
		log.Printf("[DEBUG] Updating Volterra Site obj with struct: %+v", updateReq)

		err = client.ReplaceObject(context.Background(), ves_io_schema_site.ObjectType, updateReq)
		if err != nil {
			return fmt.Errorf("error updating Site: %s", err)
		}
	}

	return resourceVolterraModifySiteRead(d, meta)
}

func getSiteReplaceForm(d *schema.ResourceData,
	getResp *ves_io_schema_site.GetResponse) *ves_io_schema_site.ReplaceRequest {

	replaceReq := getResp.ReplaceForm
	updateMeta := replaceReq.Metadata
	updateSpec := replaceReq.Spec

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {
		ms := map[string]string{}
		for k, v := range v.(map[string]interface{}) {
			ms[k] = v.(string)
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {
		for k, v := range v.(map[string]interface{}) {
			updateMeta.Labels[k] = v.(string)
		}
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("address"); ok && !isIntfNil(v) {

		updateSpec.Address =
			v.(string)
	}

	if v, ok := d.GetOk("bgp_peer_address"); ok && !isIntfNil(v) {

		updateSpec.BgpPeerAddress =
			v.(string)
	}

	if v, ok := d.GetOk("bgp_router_id"); ok && !isIntfNil(v) {

		updateSpec.BgpRouterId =
			v.(string)
	}

	if v, ok := d.GetOk("coordinates"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()

		coordinates := &ves_io_schema_site.Coordinates{}
		updateSpec.Coordinates = coordinates
		for _, set := range sl {

			coordinatesMapStrToI := set.(map[string]interface{})

			if w, ok := coordinatesMapStrToI["latitude"]; ok && !isIntfNil(w) {
				coordinates.Latitude = w.(float32)
			}

			if w, ok := coordinatesMapStrToI["longitude"]; ok && !isIntfNil(w) {
				coordinates.Longitude = w.(float32)
			}

		}

	}

	if v, ok := d.GetOk("desired_pool_count"); ok && !isIntfNil(v) {

		updateSpec.DesiredPoolCount = int32(v.(int))
	}

	if v, ok := d.GetOk("inside_nameserver"); ok && !isIntfNil(v) {

		updateSpec.InsideNameserver =
			v.(string)
	}

	if v, ok := d.GetOk("inside_vip"); ok && !isIntfNil(v) {

		updateSpec.InsideVip =
			v.(string)
	}

	if v, ok := d.GetOk("operating_system_version"); ok && !isIntfNil(v) {

		updateSpec.OperatingSystemVersion =
			v.(string)
	}

	if v, ok := d.GetOk("outside_nameserver"); ok && !isIntfNil(v) {

		updateSpec.OutsideNameserver =
			v.(string)
	}

	if v, ok := d.GetOk("outside_vip"); ok && !isIntfNil(v) {

		updateSpec.OutsideVip =
			v.(string)
	}

	if v, ok := d.GetOk("site_to_site_network_type"); ok && !isIntfNil(v) {

		updateSpec.SiteToSiteNetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

	}

	if v, ok := d.GetOk("site_to_site_tunnel_ip"); ok && !isIntfNil(v) {

		updateSpec.SiteToSiteTunnelIp =
			v.(string)
	}

	if v, ok := d.GetOk("tunnel_dead_timeout"); ok && !isIntfNil(v) {

		updateSpec.TunnelDeadTimeout =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("tunnel_type"); ok && !isIntfNil(v) {

		updateSpec.TunnelType = ves_io_schema.SiteToSiteTunnelType(ves_io_schema.SiteToSiteTunnelType_value[v.(string)])

	}

	if v, ok := d.GetOk("vip_vrrp_mode"); ok && !isIntfNil(v) {

		updateSpec.VipVrrpMode = ves_io_schema.VipVrrpType(ves_io_schema.VipVrrpType_value[v.(string)])

	}

	if v, ok := d.GetOk("volterra_software_overide"); ok && !isIntfNil(v) {

		updateSpec.VolterraSoftwareOveride = ves_io_schema_site.SiteSoftwareOverrideType(ves_io_schema_site.SiteSoftwareOverrideType_value[v.(string)])

	}

	if v, ok := d.GetOk("volterra_software_version"); ok && !isIntfNil(v) {

		updateSpec.VolterraSoftwareVersion =
			v.(string)
	}

	return replaceReq
}

func resourceVolterraModifySiteDelete(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	log.Printf("[DEBUG] Skipped Deleting Volterra Site obj with name %+v in namespace %+v", name, namespace)
	return nil
}
