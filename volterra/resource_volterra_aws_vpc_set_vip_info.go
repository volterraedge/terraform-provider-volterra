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

	ves_io_schema_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
	ves_io_schema_aws_tgw_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
	ves_io_schema_aws_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
	ves_io_schema_azure_vnet_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/azure_vnet_site"
)

const (
	setVIPInfoRPCFQN = "ves.io.schema.views.%s.CustomAPI.SetVIPInfo"
	setVIPInfoURI    = "/public/namespaces/%s/%s/%s/set_vip_info"
)

// resourceVolterraSiteSetVIPInfo is implementation of SetVipInfo Custom API for AWS VPC Site
func resourceVolterraSiteSetVIPInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSiteSetVIPInfoCreate,
		Read:   resourceVolterraSiteSetVIPInfoRead,
		Update: resourceVolterraSiteSetVIPInfoUpdate,
		Delete: resourceVolterraSiteSetVIPInfoDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"site_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"aws_vpc_site",
					"aws_tgw_site",
					"azure_vnet_site",
				}, false),
			},
			"vip_params_per_az": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside_vip": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"outside_vip": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"outside_vip_cname": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"inside_vip_cname": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"az_name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

type vipInfoRequest interface {
	GetVipParamsPerAz() []*ves_io_schema_site.PublishVIPParamsPerAz
}

// resourceVolterraSiteSetVIPInfoCreate updates the vip info for the site
func resourceVolterraSiteSetVIPInfoCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, namespace, siteType string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}
	if v, ok := d.GetOk("site_type"); ok {
		siteType = v.(string)
	}

	switch siteType {
	case "aws_vpc_site":
		req := &ves_io_schema_aws_vpc_site.SetVIPInfoRequest{
			Name:      name,
			Namespace: namespace,
		}
		if v, ok := d.GetOk("vip_params_per_az"); ok && !isIntfNil(v) {
			req.VipParamsPerAz = getVIPParamsPerAz(v.([]interface{}))
		}
		yamlReq, err := codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] Setting AWS VPC Site Vip Info struct: %+v", req)
		if _, err := client.CustomAPI(context.Background(), http.MethodPost,
			fmt.Sprintf(setVIPInfoURI, namespace, "aws_vpc_site", name),
			fmt.Sprintf(setVIPInfoRPCFQN, "aws_vpc_site"), yamlReq); err != nil {
			return fmt.Errorf("Error Setting AWS VPC Site Vip Info: %s", err)
		}

	case "azure_vnet_site":
		req := &ves_io_schema_azure_vnet_site.SetVIPInfoRequest{
			Name:      name,
			Namespace: namespace,
		}
		if v, ok := d.GetOk("vip_params_per_az"); ok && !isIntfNil(v) {
			req.VipParamsPerAz = getVIPParamsPerAz(v.([]interface{}))
		}
		yamlReq, err := codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] Setting Azure Vnet Site Vip Info struct: %+v", req)
		if _, err := client.CustomAPI(context.Background(), http.MethodPost,
			fmt.Sprintf(setVIPInfoURI, namespace, "azure_vnet_site", name),
			fmt.Sprintf(setVIPInfoRPCFQN, "azure_vnet_site"), yamlReq); err != nil {
			return fmt.Errorf("Error Setting Azure Vnet Site Vip Info: %s", err)
		}
	case "aws_tgw_site":
		req := &ves_io_schema_aws_tgw_site.SetVIPInfoRequest{
			Name:      name,
			Namespace: namespace,
		}
		if v, ok := d.GetOk("vip_params_per_az"); ok && !isIntfNil(v) {
			req.VipParamsPerAz = getVIPParamsPerAz(v.([]interface{}))
		}
		yamlReq, err := codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] Setting AWS TGW Site Vip Info struct: %+v", req)
		if _, err := client.CustomAPI(context.Background(), http.MethodPost,
			fmt.Sprintf(setVIPInfoURI, namespace, "aws_tgw_site", name),
			fmt.Sprintf(setVIPInfoRPCFQN, "aws_tgw_site"), yamlReq); err != nil {
			return fmt.Errorf("Error Setting AWS TGW Site Vip Info: %s", err)
		}
	}
	d.SetId(uuid.New().String())
	return nil
}

func resourceVolterraSiteSetVIPInfoRead(d *schema.ResourceData, meta interface{}) error {
	// Ignore dont do anything when refresh/get is trigerred
	return nil
}

// resourceVolterraSiteSetVIPInfoUpdate updates Namespace resource
func resourceVolterraSiteSetVIPInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceVolterraSiteSetVIPInfoCreate(d, meta)
}

func resourceVolterraSiteSetVIPInfoDelete(d *schema.ResourceData, meta interface{}) error {
	// Remove the VPC IP prefixes
	client := meta.(*APIClient)

	var name, namespace, siteType string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}
	if v, ok := d.GetOk("site_type"); ok {
		siteType = v.(string)
	}

	switch siteType {
	case "aws_vpc_site":
		req := &ves_io_schema_aws_vpc_site.SetVIPInfoRequest{
			Name:           name,
			Namespace:      namespace,
			VipParamsPerAz: make([]*ves_io_schema_site.PublishVIPParamsPerAz, 0),
		}
		yamlReq, err := codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] Setting AWS VPC Site Vip Info as empty: %+v", req)
		if _, err := client.CustomAPI(context.Background(), http.MethodPost,
			fmt.Sprintf(setVIPInfoURI, namespace, "aws_vpc_site", name),
			fmt.Sprintf(setVIPInfoRPCFQN, "aws_vpc_site"), yamlReq); err != nil {
			log.Printf("Ignored: Error Setting Volterra AWS VPC Site Vip Info: %s", err)
		}

	case "azure_vnet_site":
		req := &ves_io_schema_azure_vnet_site.SetVIPInfoRequest{
			Name:           name,
			Namespace:      namespace,
			VipParamsPerAz: make([]*ves_io_schema_site.PublishVIPParamsPerAz, 0),
		}
		yamlReq, err := codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] Setting Azure Vnet Site Vip Info as empty: %+v", req)
		if _, err := client.CustomAPI(context.Background(), http.MethodPost,
			fmt.Sprintf(setVIPInfoURI, namespace, "azure_vnet_site", name),
			fmt.Sprintf(setVIPInfoRPCFQN, "azure_vnet_site"), yamlReq); err != nil {
			log.Printf("Ignored: Error Setting Volterra Azure Vnet Site Vip Info: %s", err)
		}
	case "aws_tgw_site":
		req := &ves_io_schema_aws_tgw_site.SetVIPInfoRequest{
			Name:           name,
			Namespace:      namespace,
			VipParamsPerAz: make([]*ves_io_schema_site.PublishVIPParamsPerAz, 0),
		}
		yamlReq, err := codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] Setting AWS TGW Site Vip Info as empty: %+v", req)
		if _, err := client.CustomAPI(context.Background(), http.MethodPost,
			fmt.Sprintf(setVIPInfoURI, namespace, "aws_tgw_site", name),
			fmt.Sprintf(setVIPInfoRPCFQN, "aws_tgw_site"), yamlReq); err != nil {
			log.Printf("Ignored: Error Setting Volterra AWS TGW Site Vip Info: %s", err)
		}
	}
	d.SetId("")
	return nil
}

func getVIPParamsPerAz(tfVipParmsPerAz []interface{}) []*ves_io_schema_site.PublishVIPParamsPerAz {
	vipParamsPerAZ := make([]*ves_io_schema_site.PublishVIPParamsPerAz, len(tfVipParmsPerAz))
	for i, vipParams := range tfVipParmsPerAz {
		vipParamsMapStrToI := vipParams.(map[string]interface{})
		vipParamsPerAZ[i] = &ves_io_schema_site.PublishVIPParamsPerAz{}
		if v, ok := vipParamsMapStrToI["inside_vip"]; ok && !isIntfNil(v) {
			insideVip := make([]string, len(v.([]interface{})))
			for j, v := range v.([]interface{}) {
				insideVip[j] = v.(string)
			}
			vipParamsPerAZ[i].InsideVip = insideVip
		}

		if v, ok := vipParamsMapStrToI["outside_vip"]; ok && !isIntfNil(v) {
			outsideVip := make([]string, len(v.([]interface{})))
			for j, v := range v.([]interface{}) {
				outsideVip[j] = v.(string)
			}
			vipParamsPerAZ[i].OutsideVip = outsideVip
		}

		if v, ok := vipParamsMapStrToI["outside_vip_cname"]; ok && !isIntfNil(v) {
			vipParamsPerAZ[i].OutsideVipCname = v.(string)
		}

		if v, ok := vipParamsMapStrToI["inside_vip_cname"]; ok && !isIntfNil(v) {
			vipParamsPerAZ[i].InsideVipCname = v.(string)
		}

		if v, ok := vipParamsMapStrToI["az_name"]; ok && !isIntfNil(v) {
			vipParamsPerAZ[i].AzName = v.(string)
		}
	}
	return vipParamsPerAZ
}
