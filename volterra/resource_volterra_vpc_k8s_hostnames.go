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

	ves_io_schema_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
)

const (
	setVpcK8SHostnamesRPCFQN = "ves.io.schema.views.aws_vpc_site.CustomAPI.SetVPCK8SHostnames"
	setVpcK8SHostnamesURI    = "/public/namespaces/%s/aws_vpc_site/%s/storage/set_vpc_k8s_hostnames"
)

// resourceVolterraSetVpcK8SHostnames is implementation of Volterra's Namespace resources
func resourceVolterraSetVpcK8SHostnames() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSetVpcK8SHostnamesCreate,
		Read:   resourceVolterraSetVpcK8SHostnamesRead,
		Update: resourceVolterraSetVpcK8SHostnamesUpdate,
		Delete: resourceVolterraSetVpcK8SHostnamesDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_names": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// resourceVolterraSetVpcK8SHostnamesCreate creates Namespace resource
func resourceVolterraSetVpcK8SHostnamesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, namespace string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}

	req := &ves_io_schema_vpc_site.SetVPCK8SHostnamesRequest{
		Name:      name,
		Namespace: namespace,
	}
	if v, ok := d.GetOk("node_names"); ok {
		var nodes []string
		for _, v := range v.([]interface{}) {
			nodes = append(nodes, v.(string))
		}
		req.NodeNames = nodes
	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Setting Volterra VPC K8S Hostnames struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setVpcK8SHostnamesURI, namespace, name), setVpcK8SHostnamesRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting Volterra VPC K8S Hostnames: %s", err)
	}

	d.SetId(uuid.New().String())
	return nil
}

func resourceVolterraSetVpcK8SHostnamesRead(d *schema.ResourceData, meta interface{}) error {
	// Ignore dont do anything when refresh/get is trigerred
	return nil
}

// resourceVolterraSetVpcK8SHostnamesUpdate updates Namespace resource
func resourceVolterraSetVpcK8SHostnamesUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, namespace string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}

	req := &ves_io_schema_vpc_site.SetVPCK8SHostnamesRequest{
		Name:      name,
		Namespace: namespace,
	}
	if v, ok := d.GetOk("node_names"); ok {
		req.NodeNames = v.([]string)
	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Setting Volterra VPC K8S Hostnames struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setVpcK8SHostnamesURI, namespace, name), setVpcK8SHostnamesRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting Volterra VPC K8S Hostnames: %s", err)
	}
	return nil
}

func resourceVolterraSetVpcK8SHostnamesDelete(d *schema.ResourceData, meta interface{}) error {
	// Remove the VPC IP prefixes
	client := meta.(*APIClient)

	var name, namespace string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}

	req := &ves_io_schema_vpc_site.SetVPCK8SHostnamesRequest{
		Name:      name,
		Namespace: namespace,
		NodeNames: []string{},
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		log.Printf("Ignored: Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Setting Volterra VPC K8S Hostnames as empty: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setVpcK8SHostnamesURI, namespace, name), setVpcK8SHostnamesRPCFQN, yamlReq)
	if err != nil {
		log.Printf("Ignored: Error Setting Volterra VPC  K8S Hostnames: %s", err)
	}
	d.SetId("")
	return nil
}
