package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/svcfw"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_api_credential "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/api_credential"
)

const (
	serviceCredRPCFQN      = "ves.io.schema.api_credential.CustomAPI"
	serviceCredURI         = "/public/namespaces/system/service_credentials"
	deleteServiceCredURI   = "/public/namespaces/system/revoke/service_credentials"
	getServiceCredURI      = "/public/namespaces/system/service_credentials/%s"
	replaceServiceCredURI  = "/public/namespaces/system/service_credentials/%s"
	activateServiceCredURI = "/public/namespaces/system/activate/service_credentials"
	renewServiceCredURI    = "/public/namespaces/system/renew/service_credentials"
)

type serviceCredentialParams struct {
	name            string
	namespace       string
	namespaceRoles  []*ves_io_schema.NamespaceRoleType
	serviceCredType string
	password        string
	expirationdays  uint32
	userGroupNames  []string
}

func resourceVolterraServiceCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraServiceCredentialCreate,
		Read:   resourceVolterraServiceCredentialRead,
		Update: resourceVolterraServiceCredentialUpdate,
		Delete: resourceVolterraServiceCredentialDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"namespace_roles": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						"role": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"service_credential_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SERVICE_API_CERTIFICATE",
			},

			"service_api_certificate": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"service_api_token": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"site_kubeconfig": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"site": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"vk8s_kubeconfig": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vk8s_cluster_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vk8s_namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"expiration_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},

			"user_group_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"created_at": {
				Type:     schema.TypeString,
				Required: true,
			},

			"expiration_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"automatic_approval_api_token": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceVolterraServiceCredentialCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	serviceCredParams := &serviceCredentialParams{}
	var created_at string

	if v, ok := d.GetOk("name"); ok {
		serviceCredParams.name = v.(string)
	}

	if v, ok := d.GetOk("created_at"); ok {
		created_at = v.(string)
	}

	if v, ok := d.GetOk("namespace_roles"); ok {
		sl := v.([]interface{})
		namespaceRoles := make([]*ves_io_schema.NamespaceRoleType, 0)
		for _, val := range sl {
			temp := &ves_io_schema.NamespaceRoleType{}
			namespaceRolesMapStrToI := val.(map[string]interface{})
			if val != nil {
				if va, ok := namespaceRolesMapStrToI["namespace"]; ok {
					temp.Namespace = va.(string)
				}
				if va, ok := namespaceRolesMapStrToI["role"]; ok {
					temp.Role = va.(string)
				}
				namespaceRoles = append(namespaceRoles, temp)
			}
		}
		serviceCredParams.namespaceRoles = namespaceRoles

	}

	if v, ok := d.GetOk("service_credential_type"); ok {
		serviceCredParams.serviceCredType = v.(string)
	}

	if v, ok := d.GetOk("expiration_days"); ok {
		serviceCredParams.expirationdays = uint32(v.(int))
	}

	if v, ok := d.GetOk("user_group_names"); ok {
		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		serviceCredParams.userGroupNames = ls
	}

	serviceCredValue, ok := ves_io_schema_api_credential.APICredentialType_value[serviceCredParams.serviceCredType]

	if !ok {
		return fmt.Errorf("Invalid service_credential_type, valid ones are: SERVICE_API_TOKEN, SERVICE_API_CERTIFICATE, SERVICE_KUBE_CONFIG, SITE_GLOBAL_KUBE_CONFIG")
	}

	serviceCredRequest := &ves_io_schema_api_credential.CreateServiceCredentialsRequest{
		Name:           serviceCredParams.name,
		Namespace:      svcfw.SystemNSVal,
		Type:           ves_io_schema_api_credential.APICredentialType(serviceCredValue),
		ExpirationDays: serviceCredParams.expirationdays,
		NamespaceRoles: serviceCredParams.namespaceRoles,
		UserGroupNames: serviceCredParams.userGroupNames,
	}

	serviceCredentialChoiceTypeFound := false

	if v, ok := d.GetOk("service_api_certificate"); ok && !serviceCredentialChoiceTypeFound {
		serviceCredentialChoiceTypeFound = true
		sl := v.([]interface{})
		apiCert := &ves_io_schema_api_credential.CreateServiceCredentialsRequest_ApiCertificate{}
		serviceCredRequest.ServiceCredentialChoice = apiCert
		apiCert.ApiCertificate = &ves_io_schema_api_credential.ApiCertificateType{}
		for _, val := range sl {
			apiCertMapStrToI := val.(map[string]interface{})
			if va, ok := apiCertMapStrToI["password"]; ok {
				apiCert.ApiCertificate.Password = va.(string)
			}
		}
	}

	if v, ok := d.GetOk("service_api_token"); ok && !serviceCredentialChoiceTypeFound {
		serviceCredentialChoiceTypeFound = true
		if v.(bool) {
			apiToken := &ves_io_schema_api_credential.CreateServiceCredentialsRequest_ApiToken{}
			apiToken.ApiToken = &ves_io_schema.Empty{}
			serviceCredRequest.ServiceCredentialChoice = apiToken
		}
	}

	if v, ok := d.GetOk("site_kubeconfig"); ok && !serviceCredentialChoiceTypeFound {
		serviceCredentialChoiceTypeFound = true
		sl := v.([]interface{})
		siteKubeconfig := &ves_io_schema_api_credential.CreateServiceCredentialsRequest_SiteKubeconfig{}
		serviceCredRequest.ServiceCredentialChoice = siteKubeconfig
		siteKubeconfig.SiteKubeconfig = &ves_io_schema_api_credential.SiteKubeconfigType{}
		for _, val := range sl {
			siteKubeconfigMapStrToI := val.(map[string]interface{})
			if va, ok := siteKubeconfigMapStrToI["site"]; ok {
				siteKubeconfig.SiteKubeconfig.Site = va.(string)
			}
		}
	}

	if v, ok := d.GetOk("vk8s_kubeconfig"); ok && !serviceCredentialChoiceTypeFound {
		serviceCredentialChoiceTypeFound = true
		sl := v.([]interface{})
		vk8sKubeconfig := &ves_io_schema_api_credential.CreateServiceCredentialsRequest_Vk8SKubeconfig{}
		serviceCredRequest.ServiceCredentialChoice = vk8sKubeconfig
		vk8sKubeconfig.Vk8SKubeconfig = &ves_io_schema_api_credential.Vk8SKubeconfigType{}
		for _, val := range sl {
			vk8sKubeconfigMapStrToI := val.(map[string]interface{})
			if va, ok := vk8sKubeconfigMapStrToI["vk8s_cluster_name"]; ok {
				vk8sKubeconfig.Vk8SKubeconfig.Vk8SClusterName = va.(string)
			}
			if va, ok := vk8sKubeconfigMapStrToI["vk8s_namespace"]; ok {
				vk8sKubeconfig.Vk8SKubeconfig.Vk8SNamespace = va.(string)
			}
		}
	}

	yamlReq, err := codec.ToYAML(serviceCredRequest)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	rspProto, err := client.CustomAPI(context.Background(), http.MethodPost, serviceCredURI, fmt.Sprintf("%s.%s", serviceCredRPCFQN, "CreateServiceCredentials"), yamlReq)
	if err != nil {
		return fmt.Errorf("Error creating Service Credential: %s", err)
	}

	rspServiceCred := rspProto.(*ves_io_schema_api_credential.CreateResponse)

	d.SetId(rspServiceCred.Name)
	d.Set("data", rspServiceCred.Data)

	expTime, err := convertTimestampToString(rspServiceCred.ExpirationTimestamp)
	if err != nil {
		return fmt.Errorf("Error converting expiration_timestamp : %s", err)
	}
	d.Set("expiration_timestamp", expTime)
	d.Set("created_at", created_at)

	return resourceVolterraServiceCredentialRead(d, meta)
}

func resourceVolterraServiceCredentialRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*APIClient)

	serviceCredReq := &ves_io_schema_api_credential.GetRequest{
		Name:      d.Id(),
		Namespace: svcfw.SystemNSVal,
	}

	yamlReq, err := codec.ToYAML(serviceCredReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	_, err = client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(getServiceCredURI, d.Id()), fmt.Sprintf("%s.%s", apiCredRPCFQN, "GetServiceCredentials"), yamlReq)

	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Service Credential %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Service Credential resource %q: %s", d.Id(), err)
	}

	return nil

}

func resourceVolterraServiceCredentialUpdate(d *schema.ResourceData, meta interface{}) error {
	var created_at string
	var expirydays uint32
	var serviceCredType string
	var automatic_approval_api_token bool

	client := meta.(*APIClient)

	updateReq := &ves_io_schema_api_credential.ReplaceServiceCredentialsRequest{
		Name:      d.Id(),
		Namespace: svcfw.SystemNSVal,
	}

	if v, ok := d.GetOk("user_group_names"); ok {
		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		updateReq.UserGroupNames = ls
	}

	if v, ok := d.GetOk("namespace_roles"); ok {
		sl := v.([]interface{})
		temp := &ves_io_schema.NamespaceAccessType{}
		namespaceRoles := make(map[string]*ves_io_schema.RoleListType)
		temp.NamespaceRoleMap = namespaceRoles
		for _, val := range sl {

			namespaceRolesMapStrToI := val.(map[string]interface{})
			if val != nil {
				var namespace string
				roles := ves_io_schema.RoleListType{}
				var names []string
				if va, ok := namespaceRolesMapStrToI["namespace"]; ok {
					namespace = va.(string)
				}
				if va, ok := namespaceRolesMapStrToI["role"]; ok {
					names = append(names, va.(string))
				}
				roles.Names = names
				namespaceRoles[namespace] = &roles

			}
		}
		updateReq.NamespaceAccess = temp
	}

	if v, ok := d.GetOk("created_at"); ok {
		created_at = v.(string)
	}
	if v, ok := d.GetOk("automatic_approval_api_token"); ok {
		automatic_approval_api_token = v.(bool)

	}

	if v, ok := d.GetOk("expiration_days"); ok {
		expirydays = uint32(v.(int))

	}

	if v, ok := d.GetOk("service_credential_type"); ok {
		serviceCredType = v.(string)
	}

	yamlReq, err := codec.ToYAML(updateReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	_, err = client.CustomAPI(context.Background(), http.MethodPut, fmt.Sprintf(replaceServiceCredURI, updateReq.Name), fmt.Sprintf("%s.%s", apiCredRPCFQN, "ReplaceServiceCredentials"), yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Service Credential %s no longer exists", d.Id())
			return nil
		}
	}

	serviceCredReq1 := &ves_io_schema_api_credential.GetRequest{
		Name:      d.Id(),
		Namespace: svcfw.SystemNSVal,
	}

	yamlReq1, err := codec.ToYAML(serviceCredReq1)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	rsp, err := client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(getServiceCredURI, d.Id()), fmt.Sprintf("%s.%s", apiCredRPCFQN, "GetServiceCredentials"), yamlReq1)

	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Service Credential %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Service Credential resource %q: %s", d.Id(), err)
	}

	expTime, err := convertTimestampToString(rsp.(*ves_io_schema_api_credential.GetServiceCredentialsResponse).GetExpiryTimestamp())
	if err != nil {
		return fmt.Errorf("Error converting expiration_timestamp: %s", err)
	}

	expireDate, err := time.Parse(time.RFC3339, expTime)
	if err != nil {
		return fmt.Errorf("invalid expiration date format: %s", expTime)
	}

	d.Set("expiration_timestamp", expTime)
	d.Set("created_at", created_at)

	if (d.HasChange("expiration_days")) && serviceCredType == "SERVICE_API_TOKEN" && automatic_approval_api_token && rsp.(*ves_io_schema_api_credential.GetServiceCredentialsResponse).GetActive() {
		renReq := &ves_io_schema_api_credential.RenewRequest{
			Name:           d.Id(),
			Namespace:      svcfw.SystemNSVal,
			ExpirationDays: expirydays,
		}

		yamlRenReq, err := codec.ToYAML(renReq)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}

		_, err = client.CustomAPI(context.Background(), http.MethodPost, renewServiceCredURI, fmt.Sprintf("%s.%s", apiCredRPCFQN, "RenewServiceCredentials"), yamlRenReq)
		if err != nil {
			return fmt.Errorf("Error Renewing Service Credential: %s", err)
		}

	}

	if time.Now().After(expireDate) && serviceCredType == "SERVICE_API_TOKEN" && automatic_approval_api_token {

		rsp, err := client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(getServiceCredURI, d.Id()), fmt.Sprintf("%s.%s", apiCredRPCFQN, "GetServiceCredentials"), yamlReq)

		if err != nil {
			if strings.Contains(err.Error(), "status code 404") {
				log.Printf("[INFO] Service Credential %s no longer exists", d.Id())
				d.SetId("")
				return nil
			}
			return fmt.Errorf("Error finding Volterra Service Credential resource %q: %s", d.Id(), err)
		}

		activeStatus := rsp.(*ves_io_schema_api_credential.GetServiceCredentialsResponse).GetActive()

		if rsp != nil && !activeStatus {
			renReq := &ves_io_schema_api_credential.RenewRequest{
				Name:           d.Id(),
				Namespace:      svcfw.SystemNSVal,
				ExpirationDays: expirydays,
			}

			yamlRenReq, err := codec.ToYAML(renReq)
			if err != nil {
				return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
			}
			_, err = client.CustomAPI(context.Background(), http.MethodPost, activateServiceCredURI, fmt.Sprintf("%s.%s", apiCredRPCFQN, "ActivateServiceCredentials"), yamlReq1)
			if err != nil {
				return fmt.Errorf("Error activating Service Credential: %s", err)
			}
			_, err = client.CustomAPI(context.Background(), http.MethodPost, renewServiceCredURI, fmt.Sprintf("%s.%s", apiCredRPCFQN, "RenewServiceCredentials"), yamlRenReq)
			if err != nil {
				return fmt.Errorf("Error Renewing Service Credential: %s", err)
			}

		}

	}

	rsp, err = client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(getServiceCredURI, d.Id()), fmt.Sprintf("%s.%s", apiCredRPCFQN, "GetServiceCredentials"), yamlReq1)

	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Service Credential %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Service Credential resource %q: %s", d.Id(), err)
	}

	expTime, err = convertTimestampToString(rsp.(*ves_io_schema_api_credential.GetServiceCredentialsResponse).GetExpiryTimestamp())
	if err != nil {
		return fmt.Errorf("Error converting expiration_timestamp: %s", err)
	}

	d.Set("expiration_timestamp", expTime)

	return resourceVolterraServiceCredentialRead(d, meta)
}

func resourceVolterraServiceCredentialDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*APIClient)
	serviceCredReq := &ves_io_schema_api_credential.GetRequest{
		Name:      d.Id(),
		Namespace: svcfw.SystemNSVal,
	}
	log.Printf("[DEBUG] Deleting/Revoking Volterra Service credential obj %+v ", d.Id())
	yamlReq, err := codec.ToYAML(serviceCredReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	_, err = client.CustomAPI(context.Background(), http.MethodPost, deleteServiceCredURI, fmt.Sprintf("%s.%s", apiCredRPCFQN, "RevokeServiceCredentials"), yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Service Credential %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
	}
	return err
}
