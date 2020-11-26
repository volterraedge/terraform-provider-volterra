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

	ves_io_schema_tgw_site "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
)

const (
	setVPNTunnelsRPCFQN = "ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPNTunnels"
	setVPNTunnelsURI    = "/public/namespaces/%s/aws_tgw_site/%s/set_vpn_tunnels"
)

// resourceVolterraSetVPNTunnels is implementation of Volterra's Namespace resources
func resourceVolterraSetVPNTunnels() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSetVPNTunnelsCreate,
		Read:   resourceVolterraSetVPNTunnelsRead,
		Update: resourceVolterraSetVPNTunnelsUpdate,
		Delete: resourceVolterraSetVPNTunnelsDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vpn_tunnel_config": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"node_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tunnel_remote_ips": {
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

// resourceVolterraSetVPNTunnelsCreate creates Namespace resource
func resourceVolterraSetVPNTunnelsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, namespace string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}

	req := &ves_io_schema_tgw_site.SetVPNTunnelsRequest{
		Name:      name,
		Namespace: namespace,
	}
	if v, ok := d.GetOk("vpn_tunnel_config"); ok && !isIntfNil(v) {
		sl := v.([]interface{})
		vpnTunnelConfigList := []*ves_io_schema_tgw_site.AWSVPNTunnelConfigType{}
		for _, set := range sl {
			intfVPNTunnelsMapStrToI := set.(map[string]interface{})
			key, ok := intfVPNTunnelsMapStrToI["node_name"]
			if ok && !isIntfNil(key) {
				vpnTunnelConfig := &ves_io_schema_tgw_site.AWSVPNTunnelConfigType{
					NodeName: key.(string),
				}
				id, ok := intfVPNTunnelsMapStrToI["node_id"]
				if ok && !isIntfNil(id) {
					vpnTunnelConfig.NodeId = id.(string)
				}
				val, _ := intfVPNTunnelsMapStrToI["tunnel_remote_ips"]
				for _, v := range val.([]interface{}) {
					vpnTunnelConfig.TunnelRemoteIp = append(vpnTunnelConfig.TunnelRemoteIp, v.(string))
				}
				vpnTunnelConfigList = append(vpnTunnelConfigList, vpnTunnelConfig)
			}
		}
		req.Tunnels = vpnTunnelConfigList
	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[INFO] Setting Volterra VPN Tunnels struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setVPNTunnelsURI, namespace, name), setVPNTunnelsRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting Volterra VPN Tunnels: %s", err)
	}

	d.SetId(uuid.New().String())
	return nil
}

func resourceVolterraSetVPNTunnelsRead(d *schema.ResourceData, meta interface{}) error {
	// Ignore dont do anything when refresh/get is trigerred
	return nil
}

// resourceVolterraSetVPNTunnelsUpdate updates Namespace resource
func resourceVolterraSetVPNTunnelsUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, namespace string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}

	req := &ves_io_schema_tgw_site.SetVPNTunnelsRequest{
		Name:      name,
		Namespace: namespace,
	}
	if v, ok := d.GetOk("vpn_tunnel_config"); ok && !isIntfNil(v) {
		sl := v.([]interface{})
		vpnTunnelConfigList := []*ves_io_schema_tgw_site.AWSVPNTunnelConfigType{}
		for _, set := range sl {
			intfVPNTunnelsMapStrToI := set.(map[string]interface{})
			key, ok := intfVPNTunnelsMapStrToI["node_name"]
			if ok && !isIntfNil(key) {
				vpnTunnelConfig := &ves_io_schema_tgw_site.AWSVPNTunnelConfigType{
					NodeName: key.(string),
				}
				id, ok := intfVPNTunnelsMapStrToI["node_id"]
				if ok && !isIntfNil(id) {
					vpnTunnelConfig.NodeId = id.(string)
				}
				val, _ := intfVPNTunnelsMapStrToI["tunnel_remote_ips"]
				for _, v := range val.([]interface{}) {
					vpnTunnelConfig.TunnelRemoteIp = append(vpnTunnelConfig.TunnelRemoteIp, v.(string))
				}
				vpnTunnelConfigList = append(vpnTunnelConfigList, vpnTunnelConfig)
			}
		}
		req.Tunnels = vpnTunnelConfigList
	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[INFO] Setting Volterra VPN Tunnels struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setVPNTunnelsURI, namespace, name), setVPNTunnelsRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting Volterra VPN Tunnels: %s", err)
	}
	return nil
}

func resourceVolterraSetVPNTunnelsDelete(d *schema.ResourceData, meta interface{}) error {
	// Remove the VPN tunnel config
	client := meta.(*APIClient)
	var name, namespace string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}

	req := &ves_io_schema_tgw_site.SetVPNTunnelsRequest{
		Name:      name,
		Namespace: namespace,
		Tunnels:   []*ves_io_schema_tgw_site.AWSVPNTunnelConfigType{},
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		log.Printf("Ignored: Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[INFO] Setting Volterra VPN Tunnels as empty: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(setVPNTunnelsURI, namespace, name), setVPNTunnelsRPCFQN, yamlReq)
	if err != nil {
		log.Printf("Ignored: Error Setting Volterra VPN Tunnels: %s", err)
	}
	d.SetId("")
	return nil
}
