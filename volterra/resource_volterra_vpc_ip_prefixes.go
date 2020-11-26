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

	ves_io_schema_tgw_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
)

const (
	setVpcIPPrefixesRPCFQN = "ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPCIpPrefixes"
	setVpcIPPrefixesURI    = "/public/namespaces/%s/aws_tgw_site/%s/set_vpc_ip_prefixes"
)

// resourceVolterraSetVpcIPPrefixes is implementation of Volterra's Namespace resources
func resourceVolterraSetVpcIPPrefixes() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSetVpcIPPrefixesCreate,
		Read:   resourceVolterraSetVpcIPPrefixesRead,
		Update: resourceVolterraSetVpcIPPrefixesUpdate,
		Delete: resourceVolterraSetVpcIPPrefixesDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc_ip_prefixes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraSetVpcIPPrefixesCreate creates Namespace resource
func resourceVolterraSetVpcIPPrefixesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, namespace string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}

	req := &ves_io_schema_tgw_site.SetVPCIpPrefixesRequest{
		Name:      name,
		Namespace: namespace,
	}
	if v, ok := d.GetOk("vpc_ip_prefixes"); ok && !isIntfNil(v) {
		sl := v.([]interface{})
		ipPrefixMap := make(map[string]*ves_io_schema_tgw_site.VPCIpPrefixesType)
		for _, set := range sl {
			intfIPPrefixesMapStrToI := set.(map[string]interface{})
			key, ok := intfIPPrefixesMapStrToI["name"]
			if ok && !isIntfNil(key) {
				vpcIPprefixes := &ves_io_schema_tgw_site.VPCIpPrefixesType{}
				ipPrefixMap[key.(string)] = vpcIPprefixes
				val, _ := intfIPPrefixesMapStrToI["value"]
				for _, v := range val.([]interface{}) {
					vpcIPprefixes.Prefixes = append(vpcIPprefixes.Prefixes, v.(string))
				}
			}
		}
		req.VpcIpPrefixes = ipPrefixMap
	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[INFO] Setting Volterra VPC IPPrefixes struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setVpcIPPrefixesURI, namespace, name), setVpcIPPrefixesRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting Volterra VPC IPPrefixes: %s", err)
	}

	d.SetId(uuid.New().String())
	return nil
}

func resourceVolterraSetVpcIPPrefixesRead(d *schema.ResourceData, meta interface{}) error {
	// Ignore dont do anything when refresh/get is trigerred
	return nil
}

// resourceVolterraSetVpcIPPrefixesUpdate updates Namespace resource
func resourceVolterraSetVpcIPPrefixesUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, namespace string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}

	req := &ves_io_schema_tgw_site.SetVPCIpPrefixesRequest{
		Name:      name,
		Namespace: namespace,
	}
	if v, ok := d.GetOk("vpc_ip_prefixes"); ok && !isIntfNil(v) {
		sl := v.([]interface{})
		ipPrefixMap := make(map[string]*ves_io_schema_tgw_site.VPCIpPrefixesType)
		for _, set := range sl {
			intfIPPrefixesMapStrToI := set.(map[string]interface{})
			key, ok := intfIPPrefixesMapStrToI["name"]
			if ok && !isIntfNil(key) {
				vpcIPprefixes := &ves_io_schema_tgw_site.VPCIpPrefixesType{}
				ipPrefixMap[key.(string)] = vpcIPprefixes
				val, _ := intfIPPrefixesMapStrToI["value"]
				for _, v := range val.([]interface{}) {
					vpcIPprefixes.Prefixes = append(vpcIPprefixes.Prefixes, v.(string))
				}
			}
		}
		req.VpcIpPrefixes = ipPrefixMap
	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[INFO] Setting Volterra VPC IPPrefixes struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setVpcIPPrefixesURI, namespace, name), setVpcIPPrefixesRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting Volterra VPC IPPrefixeses: %s", err)
	}
	return nil
}

func resourceVolterraSetVpcIPPrefixesDelete(d *schema.ResourceData, meta interface{}) error {
	// Remove the VPC IP prefixes
	client := meta.(*APIClient)

	var name, namespace string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}

	req := &ves_io_schema_tgw_site.SetVPCIpPrefixesRequest{
		Name:          name,
		Namespace:     namespace,
		VpcIpPrefixes: make(map[string]*ves_io_schema_tgw_site.VPCIpPrefixesType),
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		log.Printf("Ignored: Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[INFO] Setting Volterra VPC IPPrefixes as empty: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setVpcIPPrefixesURI, namespace, name), setVpcIPPrefixesRPCFQN, yamlReq)
	if err != nil {
		log.Printf("Ignored: Error Setting Volterra VPC IPPrefixes: %s", err)
	}
	d.SetId("")
	return nil
}
