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
	"gopkg.volterra.us/stdlib/codec"

	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_tgw_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
)

const (
	setTGWInfoRPCFQN = "ves.io.schema.views.aws_tgw_site.CustomAPI.SetTGWInfo"
	setTGWInfoURI    = "/public/namespaces/%s/aws_tgw_site/%s/set_tgw_info"
	systemNS         = "system"
)

// resourceVolterraSetTGWInfo is implementation of Volterra's Namespace resources
func resourceVolterraSetTGWInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSetTGWInfoCreate,
		Read:   resourceVolterraSetTGWInfoRead,
		Update: resourceVolterraSetTGWInfoUpdate,
		Delete: resourceVolterraSetTGWInfoDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tgw_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"az": {
							Type:     schema.TypeString,
							Required: true,
						},
						"outside_subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"inside_subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"workload_subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraSetTGWInfoCreate creates Namespace resource
func resourceVolterraSetTGWInfoCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, tgwID, vpcID string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("tgw_id"); ok {
		tgwID = v.(string)
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		vpcID = v.(string)
	}

	subnetIDs := getSubnetIDs(d)
	req := &ves_io_schema_tgw_site.SetTGWInfoRequest{
		Name:      name,
		Namespace: systemNS,
		TgwInfo: &ves_io_schema_tgw_site.AWSTGWInfoConfigType{
			TgwId:     tgwID,
			VpcId:     vpcID,
			SubnetIds: subnetIDs,
		},
	}
	log.Printf("[INFO] Foo Setting Id %s, %s\n", tgwID, vpcID)
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[INFO] Setting Volterra TGW Info struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setTGWInfoURI, systemNS, name), setTGWInfoRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting Volterra TGW Info: %s", err)
	}

	d.SetId(uuid.New().String())
	return nil
}

func resourceVolterraSetTGWInfoRead(d *schema.ResourceData, meta interface{}) error {
	// Ignore dont do anything when refresh/get is trigerred
	return nil
}

// resourceVolterraSetTGWInfoUpdate updates Namespace resource
func resourceVolterraSetTGWInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, tgwID, vpcID string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("tgw_id"); ok {
		tgwID = v.(string)
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		vpcID = v.(string)
	}
	subnetIDs := getSubnetIDs(d)

	req := &ves_io_schema_tgw_site.SetTGWInfoRequest{
		Name:      name,
		Namespace: systemNS,
		TgwInfo: &ves_io_schema_tgw_site.AWSTGWInfoConfigType{
			TgwId:     tgwID,
			VpcId:     vpcID,
			SubnetIds: subnetIDs,
		},
	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[INFO] Setting Volterra TGW Info struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setTGWInfoURI, systemNS, name), setTGWInfoRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting Volterra TGW Info: %s", err)
	}
	return nil
}

func resourceVolterraSetTGWInfoDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func getSubnetIDs(d *schema.ResourceData) []*ves_io_schema_views.AWSSubnetIdsType {
	subnetIDs := []*ves_io_schema_views.AWSSubnetIdsType{}
	if v, ok := d.GetOk("subnet_ids"); ok && !isIntfNil(v) {
		sl := v.([]interface{})
		for _, set := range sl {
			subnetID := &ves_io_schema_views.AWSSubnetIdsType{}
			subnetIDMapStrToI := set.(map[string]interface{})
			if az, ok := subnetIDMapStrToI["az"]; ok && !isIntfNil(az) {
				subnetID.AzName = az.(string)
			}
			if subnet_id, ok := subnetIDMapStrToI["outside_subnet_id"]; ok && !isIntfNil(subnet_id) {
				subnetID.OutsideSubnetId = subnet_id.(string)
			}
			if subnet_id, ok := subnetIDMapStrToI["inside_subnet_id"]; ok && !isIntfNil(subnet_id) {
				subnetID.InsideSubnetId = subnet_id.(string)
			}
			if subnet_id, ok := subnetIDMapStrToI["workload_subnet_id"]; ok && !isIntfNil(subnet_id) {
				subnetID.WorkloadSubnetId = subnet_id.(string)
			}
			subnetIDs = append(subnetIDs, subnetID)
		}
	}
	return subnetIDs
}
