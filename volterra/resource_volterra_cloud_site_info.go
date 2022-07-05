//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"gopkg.volterra.us/stdlib/codec"

	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/azure_vnet_site"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/gcp_vpc_site"
)

const (
	setCloudInfoRPCFQN = "ves.io.schema.views.%s.CustomAPI.SetCloudSiteInfo"
	setCloudInfoURI    = "/public/namespaces/system/%s/%s/set_cloud_site_info"
)

// resourceVolterraSetCloudSiteInfo is implementation of SetCloudSiteInfo custom api
// for aws_vpc_site azure_vnet_site and gcp_vpc_site
func resourceVolterraSetCloudSiteInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSetCloudSiteInfoCreate,
		Read:   resourceVolterraSetCloudSiteInfoRead,
		Update: resourceVolterraSetCloudSiteInfoUpdate,
		Delete: resourceVolterraSetCloudSiteInfoDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"site_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"aws_vpc_site",
					"azure_vnet_site",
					"gcp_vpc_site",
				}, false),
			},
			"public_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"private_ips": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// resourceVolterraSetCloudSiteInfoCreate applies cloud site information to the site
func resourceVolterraSetCloudSiteInfoCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, siteType, yamlReq string
	var err error
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("site_type"); ok {
		siteType = v.(string)
	}
	var publicIps, privateIps []string
	if v, ok := d.GetOk("public_ips"); ok {
		publicIps = make([]string, len(v.([]interface{})))
		for j, v := range v.([]interface{}) {
			publicIps[j] = v.(string)
		}
	}
	if v, ok := d.GetOk("private_ips"); ok {
		privateIps = make([]string, len(v.([]interface{})))
		for j, v := range v.([]interface{}) {
			privateIps[j] = v.(string)
		}
	}

	switch siteType {
	case "aws_vpc_site":
		req := &aws_vpc_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AwsVpcInfo: &aws_vpc_site.AWSVPCSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}

	case "azure_vnet_site":
		req := &azure_vnet_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AzureVnetInfo: &azure_vnet_site.AzureVnetSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
	case "gcp_vpc_site":
		req := &gcp_vpc_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AwsVpcInfo: &gcp_vpc_site.GCPVPCSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
	}
	log.Printf("[DEBUG] Setting %s SetCloudSiteInfo struct: %+v", siteType, yamlReq)
	if _, err := client.CustomAPI(context.Background(), http.MethodPost,
		fmt.Sprintf(setCloudInfoURI, siteType, name),
		fmt.Sprintf(setCloudInfoRPCFQN, siteType), yamlReq); err != nil {
		return fmt.Errorf("Error Setting %s SetCloudSiteInfo: %s", siteType, err)
	}
	d.SetId(uuid.New().String())
	return nil
}

func resourceVolterraSetCloudSiteInfoRead(d *schema.ResourceData, meta interface{}) error {
	// Ignore dont do anything when refresh/get is trigerred
	return nil
}

// resourceVolterraSetCloudSiteInfoUpdate applies cloud site information to the site
func resourceVolterraSetCloudSiteInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceVolterraSetCloudSiteInfoCreate(d, meta)
}

func resourceVolterraSetCloudSiteInfoDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var name, siteType, yamlReq string
	var err error
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("site_type"); ok {
		siteType = v.(string)
	}
	var publicIps, privateIps []string
	switch siteType {
	case "aws_vpc_site":
		req := &aws_vpc_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AwsVpcInfo: &aws_vpc_site.AWSVPCSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}

	case "azure_vnet_site":
		req := &azure_vnet_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AzureVnetInfo: &azure_vnet_site.AzureVnetSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
	case "gcp_vpc_site":
		req := &gcp_vpc_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AwsVpcInfo: &gcp_vpc_site.GCPVPCSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
	}
	log.Printf("[DEBUG] Setting %s SetCloudSiteInfo struct: %+v", siteType, yamlReq)
	if _, err := client.CustomAPI(context.Background(), http.MethodPost,
		fmt.Sprintf(setCloudInfoURI, siteType, name),
		fmt.Sprintf(setCloudInfoRPCFQN, siteType), yamlReq); err != nil {
		return fmt.Errorf("Error Setting %s SetCloudSiteInfo: %s", siteType, err)
	}
	d.SetId("")
	return nil
}