//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	// AllResources will hold all resource definitions of volterra initialized by init() function
	// in respective resource file
	AllResources = map[string]*schema.Resource{}
	// EnvVarP12Password is the name of environment variable that holds password for P12 bundle file
	EnvVarP12Password = "VES_P12_PASSWORD"
)

const (
	activeNetworkPolicies     = "volterra_active_network_policies"
	activeServicePolicies     = "volterra_active_service_policies"
	activeAlertPolicies       = "volterra_active_alert_policies"
	setFastACLForInternetVIPs = "volterra_fast_acl_for_internet_vips"
	approvalResource          = "volterra_registration_approval"
	modifySite                = "volterra_modify_site"
	apiCredential             = "volterra_api_credential"
	publicIP                  = "volterra_public_ip"
	siteState                 = "volterra_site_state"
	setVPCK8SHostnames        = "volterra_vpc_k8s_hostnames"
	setVPCIPPrefixes          = "volterra_tgw_vpc_ip_prefixes"
	setVPNTunnels             = "volterra_tgw_vpn_tunnels"
	setTGWInfo                = "volterra_tgw_info"
	tfParamsAction            = "volterra_tf_params_action"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_p12_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLT_API_P12_FILE", nil),
				Description: "Absolute path to p12 file.",
			},
			"api_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLT_API_CERT", nil),
				Description: "The volt API cert for api operations.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLT_API_KEY", nil),
				Description: "The volt API key for api operations.",
			},
			"api_ca_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLT_API_CA_CERT", nil),
				Description: "The volt API CA Cert for api operations.",
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLT_API_URL", nil),
				Description: "The volt API server url.",
			},
			"test": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLT_API_TEST", false),
				Description: "The volt API test flag .",
			},
			"vesenv": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLT_VESENV", false),
				Description: "The vesenv flag for itnernal use .",
			},
			"tenant": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLT_TENANT", nil),
				Description: "The tenant values, used only when vesenv is set.",
			},
			"timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLT_API_TIMEOUT", "10s"),
				Description: "The Volt API call timeout value, by default its 10s",
			},
		},
		ResourcesMap: getResourceMap(),
		DataSourcesMap: map[string]*schema.Resource{
			"volterra_namespace":                   dataSourceVolterraNamespace(),
			"volterra_virtual_host_dns_info":       dataSourceVolterraVirtualHostDNSInfo(),
			"volterra_address_allocator":           dataSourceVolterraAddressAllocator(),
			"volterra_tunnel":                      dataSourceVolterraTunnel(),
			"volterra_parse_aws_cgw_configuration": dataSourceVolterraParseAWSCGWConfiguration(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {

	config := Config{}
	if v, ok := d.GetOk("url"); ok {
		config.url = v.(string)
	}

	if v, ok := d.GetOk("vesenv"); ok {
		config.vesenv = v.(bool)
		if v, ok := d.GetOk("tenant"); ok {
			config.tenant = v.(string)
		}
		log.Printf(`[DEBUG] VESENV is set for tenant %s`, config.tenant)
	}

	config.timeout = d.Get("timeout").(string)

	if v, ok := d.GetOk("test"); ok {
		config.test = v.(bool)
		return config.Client()
	}

	if v, ok := d.GetOk("api_p12_file"); ok {
		config.apiP12File = fmt.Sprintf("file:///%s", v.(string))
		config.apiP12Password = os.Getenv(EnvVarP12Password)

	} else if v, ok := d.GetOk("api_cert"); ok {
		config.apiCert = fmt.Sprintf("file:///%s", v.(string))
		if v, ok = d.GetOk("api_key"); !ok {
			return nil, fmt.Errorf("api_key must be provided with api_cert as provider config")
		}
		config.apiKey = fmt.Sprintf("file:///%s", v.(string))
	} else {
		return nil, fmt.Errorf("neither api_p12 bundle or api_cert/api_key is provided as provider config")
	}

	if v, ok := d.GetOk("api_ca_cert"); ok {
		config.apiCACert = fmt.Sprintf("file:///%s", v.(string))
	}

	log.Printf(`[DEBUG] Creating volterra client with config
		certFile: %s
		pvtKeyFile: %s
		cacertFile: %s
		p12File: %s
		testURL: %s
		`, config.apiCert, config.apiKey, config.apiCACert,
		config.apiP12File, config.url)

	return config.Client()
}

func getResourceMap() map[string]*schema.Resource {
	// auto generated resource map
	resourceMap := getVolterraResourceMap()
	// add custom respurce map
	resourceMap[activeAlertPolicies] = resourceVolterraActiveAlertPolicies()
	resourceMap[activeNetworkPolicies] = resourceVolterraActiveNetworkPolicies()
	resourceMap[activeServicePolicies] = resourceVolterraActiveServicePolicies()
	resourceMap[apiCredential] = resourceVolterraAPICredential()
	resourceMap[approvalResource] = resourceVolterraRegistrationApproval()
	resourceMap[modifySite] = resourceVolterraModifySite()
	resourceMap[publicIP] = resourceVolterraPublicIp()
	resourceMap[setFastACLForInternetVIPs] = resourceVolterraAFastACLsForInternetVIPs()
	resourceMap[siteState] = resourceVolterraSiteState()
	resourceMap[setVPCK8SHostnames] = resourceVolterraSetVpcK8SHostnames()
	resourceMap[setVPCIPPrefixes] = resourceVolterraSetVpcIPPrefixes()
	resourceMap[setVPNTunnels] = resourceVolterraSetVPNTunnels()
	resourceMap[setTGWInfo] = resourceVolterraSetTGWInfo()
	resourceMap[tfParamsAction] = resourceVolterraTFParamsRunAction()
	return resourceMap
}
