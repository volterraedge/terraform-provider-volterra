//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	google_protobuf1 "github.com/gogo/protobuf/types"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ves_io_schema_api_credential "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/api_credential"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/svcfw"
)

const (
	apiCredRPCFQN = "ves.io.schema.api_credential.CustomAPI"
	uri           = "/public/namespaces/system/api_credentials"
)

type apiCredentialParams struct {
	name        string
	apiCredType string
	vk8sNS      string
	vk8sName    string
	password    string
	expirydays  int
}

// resourceVolterraAPICredential is implementation of Volterra's API Credential Resource
func resourceVolterraAPICredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAPICredentialCreate,
		Read:   resourceVolterraAPICredentialRead,
		Update: resourceVolterraAPICredentialUpdate,
		Delete: resourceVolterraAPICredentialDelete,

		Schema: map[string]*schema.Schema{

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_credential_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_k8s_namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_k8s_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"api_credential_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expiry_days": {
				Type:     schema.TypeInt,
				Default:  10,
				Optional: true,
			},
		},
	}
}

// resourceVolterraAPICredentialCreate approves registration resource
func resourceVolterraAPICredentialCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	apiCredParams := &apiCredentialParams{}
	if v, ok := d.GetOk("name"); ok {
		apiCredParams.name = v.(string)
	}
	if v, ok := d.GetOk("api_credential_type"); ok {
		apiCredParams.apiCredType = v.(string)
	}
	if v, ok := d.GetOk("virtual_k8s_namespace"); ok {
		apiCredParams.vk8sNS = v.(string)
	}
	if v, ok := d.GetOk("virtual_k8s_name"); ok {
		apiCredParams.vk8sName = v.(string)
	}
	if v, ok := d.GetOk("api_credential_password"); ok {
		apiCredParams.password = v.(string)
	}
	if v, ok := d.GetOk("expiry_days"); ok {
		apiCredParams.expirydays = v.(int)
	}

	apiCredValue, ok := ves_io_schema_api_credential.APICredentialType_value[apiCredParams.apiCredType]
	if !ok {
		return fmt.Errorf("Invalid api_credential_type, valid ones are: API_CERTIFICATE, KUBE_CONFIG, API_TOKEN")
	}

	apiCredReq := &ves_io_schema_api_credential.CreateRequest{
		Name:            apiCredParams.name,
		Namespace:       svcfw.SystemNSVal,
		Expiration_Days: apiCredParams.expirydays,
	}

	apiCredSpec := &ves_io_schema_api_credential.CustomCreateSpecType{
		Type: ves_io_schema_api_credential.APICredentialType(apiCredValue),
	}

	if apiCredParams.vk8sName != "" {
		apiCredSpec.VirtualK8SName = apiCredParams.vk8sName
	}
	if apiCredParams.vk8sNS != "" {
		apiCredSpec.VirtualK8SNamespace = apiCredParams.vk8sNS
	}
	if apiCredParams.password != "" {
		apiCredSpec.Password = apiCredParams.password
	}
	apiCredReq.Spec = apiCredSpec

	yamlReq, err := codec.ToYAML(apiCredReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	rspProto, err := client.CustomAPI(context.Background(), http.MethodPost, uri, fmt.Sprintf("%s.%s", apiCredRPCFQN, "Create"), yamlReq)
	if err != nil {
		return fmt.Errorf("Error creating API Credential: %s", err)
	}

	rspAPICred := rspProto.(*ves_io_schema_api_credential.CreateResponse)
	d.SetId(apiCredParams.name)
	d.Set("data", rspAPICred.Data)
	return resourceVolterraAPICredentialRead(d, meta)
}

func resourceVolterraAPICredentialRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	apiCredReq := &ves_io_schema_api_credential.GetRequest{
		Name:      d.Id(),
		Namespace: svcfw.SystemNSVal,
	}

	yamlReq, err := codec.ToYAML(apiCredReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	_, err = client.CustomAPI(context.Background(), http.MethodGet, uri, fmt.Sprintf("%s.%s", apiCredRPCFQN, "Get"), yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Registration %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra API Credential resource %q: %s", d.Id(), err)
	}
	return nil
}

func resourceVolterraAPICredentialUpdate(d *schema.ResourceData, meta interface{}) error {
	// cannot update api credential object
	return resourceVolterraAPICredentialRead(d, meta)
}

func resourceVolterraAPICredentialDelete(d *schema.ResourceData, meta interface{}) error {
	// cannot delete api credential object

	log.Printf("[DEBUG] Deleting Volterra API Credential obj with name %+v", d.Id())
	d.SetId("")
	return nil
}
