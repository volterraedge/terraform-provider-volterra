//
// Copyright (c) 2023 F5 Inc. All rights reserved.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/svcfw"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_api_credential "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/api_credential"
)

const (
	apiCredRPCFQN      = "ves.io.schema.api_credential.CustomAPI"
	apiCredURI         = "/public/namespaces/system/api_credentials"
	getAPICredURI      = "/public/namespaces/%s/api_credentials/%s"
	deleteAPICredURI   = "/public/namespaces/%s/revoke/api_credentials"
	activateAPICredURI = "/public/namespaces/system/activate/api_credentials"
	renewAPICredURI    = "/public/namespaces/system/renew/api_credentials"
)

type apiCredentialParams struct {
	name        string
	apiCredType string
	vk8sNS      string
	vk8sName    string
	password    string
	expirydays  uint32
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
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "This field is deprecated and will be removed in future release.",
			},
			"virtual_k8s_namespace": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "This field is deprecated and will be removed in future release.",
			},
			"virtual_k8s_name": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "This field is deprecated and will be removed in future release.",
			},
			"api_credential_password": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: "This field is deprecated and will be removed in future release.",
			},
			"api_certificate": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
					},
				},
			},
			"api_token": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expiration_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expiry_days": {
				Type:     schema.TypeInt,
				Default:  10,
				Optional: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Required: true,
			},
			"automatic_approval_api_token": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

// resourceVolterraAPICredentialCreate approves registration resource
func resourceVolterraAPICredentialCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var created_at string

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
		apiCredParams.expirydays = uint32(v.(int))
	}

	if v, ok := d.GetOk("created_at"); ok {
		created_at = v.(string)
	}

	apiCredReq := &ves_io_schema_api_credential.CreateRequest{
		Name:           apiCredParams.name,
		Namespace:      svcfw.SystemNSVal,
		ExpirationDays: apiCredParams.expirydays,
	}

	var apiCredSpec *ves_io_schema_api_credential.CustomCreateSpecType
	if apiCredParams.apiCredType != "" {
		apiCredValue, ok := ves_io_schema_api_credential.APICredentialType_value[apiCredParams.apiCredType]
		if !ok {
			return fmt.Errorf("Invalid api_credential_type, valid ones are: API_CERTIFICATE, KUBE_CONFIG, API_TOKEN")
		}

		apiCredSpec = &ves_io_schema_api_credential.CustomCreateSpecType{
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
	}

	credentialChoiceTypeFound := false

	if v, ok := d.GetOk("api_certificate"); ok && !credentialChoiceTypeFound {
		credentialChoiceTypeFound = true
		sl := v.([]interface{})
		apiCert := &ves_io_schema_api_credential.CreateRequest_ApiCertificate{
			ApiCertificate: &ves_io_schema_api_credential.ApiCertificateType{},
		}
		for _, val := range sl {
			apiCertMapStrToI := val.(map[string]interface{})
			if va, ok := apiCertMapStrToI["password"]; ok {
				apiCert.ApiCertificate.Password = va.(string)
			}
		}
		apiCredReq.SelfCredentialChoice = apiCert
	}

	if v, ok := d.GetOk("api_token"); ok && !credentialChoiceTypeFound {
		credentialChoiceTypeFound = true
		if v.(bool) {
			apiCredReq.SelfCredentialChoice = &ves_io_schema_api_credential.CreateRequest_ApiToken{
				ApiToken: &ves_io_schema.Empty{},
			}
		}
	}

	if apiCredReq.Spec == nil && apiCredReq.SelfCredentialChoice == nil {
		return fmt.Errorf("one of api_credential_type, api_certificate, or api_token must be configured")
	}

	yamlReq, err := codec.ToYAML(apiCredReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	rspProto, err := client.CustomAPI(context.Background(), http.MethodPost, apiCredURI, fmt.Sprintf("%s.%s", apiCredRPCFQN, "Create"), yamlReq)
	if err != nil {
		return fmt.Errorf("Error creating API Credential: %s", err)
	}

	rspAPICred := rspProto.(*ves_io_schema_api_credential.CreateResponse)
	d.SetId(rspAPICred.Name)
	d.Set("data", rspAPICred.Data)
	expTime, err := convertTimestampToString(rspAPICred.ExpirationTimestamp)
	if err != nil {
		return fmt.Errorf("Error converting expiration_timestamp : %s", err)
	}
	d.Set("expiration_timestamp", expTime)
	d.Set("created_at", created_at)
	return resourceVolterraAPICredentialRead(d, meta)
}

func convertTimestampToString(timestamp *types.Timestamp) (string, error) {
	if timestamp == nil {
		return "", fmt.Errorf("timestamp is nil")
	}
	t := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
	utcTime := t.UTC()
	return utcTime.Format(time.RFC3339), nil
}

func resourceVolterraAPICredentialRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(credentialCustomAPIClient)

	_, err := getAPICredential(client, d)
	if err != nil {
		if isCredentialNotFoundError(err) {
			log.Printf("[INFO] API Credential %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra API Credential resource %q: %s", d.Id(), err)
	}

	return nil
}

func apiCredentialReadNamespaces(d *schema.ResourceData) []string {
	if credentialType, ok := d.GetOk("api_credential_type"); ok && credentialType.(string) == "KUBE_CONFIG" {
		if namespace, ok := d.GetOk("virtual_k8s_namespace"); ok && namespace.(string) != "" && namespace.(string) != svcfw.SystemNSVal {
			return []string{namespace.(string), svcfw.SystemNSVal}
		}
	}

	return []string{svcfw.SystemNSVal}
}

func getAPICredential(client credentialCustomAPIClient, d *schema.ResourceData) (proto.Message, error) {
	var lastErr error
	for _, namespace := range apiCredentialReadNamespaces(d) {
		apiCredReq := &ves_io_schema_api_credential.GetRequest{
			Name:      d.Id(),
			Namespace: namespace,
		}

		yamlReq, err := codec.ToYAML(apiCredReq)
		if err != nil {
			return nil, fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}

		rsp, err := client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(getAPICredURI, namespace, d.Id()), fmt.Sprintf("%s.%s", apiCredRPCFQN, "Get"), yamlReq)
		if err == nil {
			return rsp, nil
		}
		if !isCredentialNotFoundError(err) {
			return nil, err
		}
		lastErr = err
	}

	return nil, lastErr
}

func resourceVolterraAPICredentialUpdate(d *schema.ResourceData, meta interface{}) error {
	// cannot update api credential object
	var created_at string
	var expirydays uint32
	var apiCredType string
	var apiToken bool
	var automatic_approval_api_token bool
	client := meta.(*APIClient)

	if v, ok := d.GetOk("expiry_days"); ok {
		expirydays = uint32(v.(int))
	}

	if v, ok := d.GetOk("api_credential_type"); ok {
		apiCredType = v.(string)
	}
	if v, ok := d.GetOk("api_token"); ok {
		apiToken = v.(bool)
	}

	if v, ok := d.GetOk("automatic_approval_api_token"); ok {
		automatic_approval_api_token = v.(bool)
	}

	if v, ok := d.GetOk("created_at"); ok {
		created_at = v.(string)
	}

	apiCredReq := &ves_io_schema_api_credential.GetRequest{
		Name:      d.Id(),
		Namespace: svcfw.SystemNSVal,
	}

	yamlReq, err := codec.ToYAML(apiCredReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	rsp, err := getAPICredential(client, d)
	if err != nil {
		if isCredentialNotFoundError(err) {
			log.Printf("[INFO] API Credential %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra API Credential resource %q: %s", d.Id(), err)
	}

	apiCredSpec := rsp.(*ves_io_schema_api_credential.GetResponse).GetObject().GetSpec().GetGcSpec()
	isActive := apiCredSpec.GetActive()
	isAPIToken := apiCredType == "API_TOKEN" || apiToken || apiCredSpec.GetType() == ves_io_schema_api_credential.API_TOKEN

	expirationDate, err := convertTimestampToString(apiCredSpec.GetExpirationTimestamp())

	expireDate, err := time.Parse(time.RFC3339, expirationDate)
	if err != nil {
		return fmt.Errorf("invalid expiration date format: %s", err)
	}

	log.Println(time.Now().After(expireDate), isAPIToken, automatic_approval_api_token)

	// Check if the certificate is expired
	if time.Now().After(expireDate) && isAPIToken && automatic_approval_api_token {
		renReq := &ves_io_schema_api_credential.RenewRequest{
			Name:           d.Id(),
			Namespace:      svcfw.SystemNSVal,
			ExpirationDays: expirydays,
		}

		yamlRenReq, err := codec.ToYAML(renReq)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
		if !isActive {
			_, err = client.CustomAPI(context.Background(), http.MethodPost, activateAPICredURI, fmt.Sprintf("%s.%s", apiCredRPCFQN, "Activate"), yamlReq)
			if err != nil {
				return fmt.Errorf("Error activating API Credential: %s", err)
			}

			_, err = client.CustomAPI(context.Background(), http.MethodPost, renewAPICredURI, fmt.Sprintf("%s.%s", apiCredRPCFQN, "Renew"), yamlRenReq)
			if err != nil {
				return fmt.Errorf("Error Renewing API Credential: %s", err)
			}

		}

	}

	rsp, err = getAPICredential(client, d)
	if err != nil {
		if isCredentialNotFoundError(err) {
			log.Printf("[INFO] API Credential %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra API Credential resource %q: %s", d.Id(), err)
	}

	expTime, err := convertTimestampToString(rsp.(*ves_io_schema_api_credential.GetResponse).GetObject().GetSpec().GcSpec.GetExpirationTimestamp())
	if err != nil {
		return fmt.Errorf("Error converting expiration_timestamp: %s", err)
	}

	d.Set("expiration_timestamp", expTime)
	d.Set("created_at", created_at)

	return resourceVolterraAPICredentialRead(d, meta)
}

func resourceVolterraAPICredentialDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(credentialCustomAPIClient)
	log.Printf("[DEBUG] Deleting/Revoking Volterra API credential obj %+v ", d.Id())
	var lastErr error
	for _, namespace := range apiCredentialReadNamespaces(d) {
		apiCredReq := &ves_io_schema_api_credential.GetRequest{
			Name:      d.Id(),
			Namespace: namespace,
		}
		yamlReq, err := codec.ToYAML(apiCredReq)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
		_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(deleteAPICredURI, namespace), fmt.Sprintf("%s.%s", apiCredRPCFQN, "Revoke"), yamlReq)
		if err == nil {
			return nil
		}
		if !isCredentialNotFoundError(err) {
			return err
		}
		lastErr = err
	}

	if isCredentialNotFoundError(lastErr) {
		log.Printf("[INFO] API Credential %s no longer exists", d.Id())
		d.SetId("")
		return nil
	}
	return lastErr
}
