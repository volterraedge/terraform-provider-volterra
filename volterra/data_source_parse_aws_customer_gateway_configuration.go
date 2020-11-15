//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type xmlConnConfig struct {
	Tunnels []xmlIpsecTunnel `xml:"ipsec_tunnel"`
}

type xmlIpsecTunnel struct {
	OutsideAddress   string `xml:"vpn_gateway>tunnel_outside_address>ip_address"`
	BGPASN           string `xml:"vpn_gateway>bgp>asn"`
	BGPHoldTime      int    `xml:"vpn_gateway>bgp>hold_time"`
	PreSharedKey     string `xml:"ike>pre_shared_key"`
	CgwInsideAddress string `xml:"customer_gateway>tunnel_inside_address>ip_address"`
	VgwInsideAddress string `xml:"vpn_gateway>tunnel_inside_address>ip_address"`
}

// dataSourceVolterraParseAWSCGWConfiguration resource helps to parse aws cgw configuration
func dataSourceVolterraParseAWSCGWConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: dataVolterraParseAWSCGWConfigurationRead,

		Schema: map[string]*schema.Schema{
			"customer_gateway_configuration": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel1_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel1_cgw_inside_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel1_vgw_inside_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel1_bgp_asn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel1_bgp_holdtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"tunnel2_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel2_cgw_inside_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel2_vgw_inside_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel2_bgp_asn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel2_bgp_holdtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

// dataVolterraParseAWSCGWConfigurationRead parse aws cgw configuration
func dataVolterraParseAWSCGWConfigurationRead(d *schema.ResourceData, meta interface{}) error {

	xmlConfig := d.Get("customer_gateway_configuration").(string)
	var vpnConfig xmlConnConfig
	if err := xml.Unmarshal([]byte(xmlConfig), &vpnConfig); err != nil {
		return fmt.Errorf("Error Unmarshalling XML: %s", err)
	}
	log.Printf("[DEBUG] foo vpnconfig: %+v", vpnConfig)
	log.Printf("[DEBUG] foo tunnels: %+v", vpnConfig.Tunnels)

	d.Set("tunnel1_address", vpnConfig.Tunnels[0].OutsideAddress)
	d.Set("tunnel1_cgw_inside_address", vpnConfig.Tunnels[0].CgwInsideAddress)
	d.Set("tunnel1_vgw_inside_address", vpnConfig.Tunnels[0].VgwInsideAddress)
	d.Set("tunnel1_preshared_key", vpnConfig.Tunnels[0].PreSharedKey)
	d.Set("tunnel1_bgp_asn", vpnConfig.Tunnels[0].BGPASN)
	d.Set("tunnel1_bgp_holdtime", vpnConfig.Tunnels[0].BGPHoldTime)
	d.Set("tunnel2_address", vpnConfig.Tunnels[1].OutsideAddress)
	d.Set("tunnel2_preshared_key", vpnConfig.Tunnels[1].PreSharedKey)
	d.Set("tunnel2_cgw_inside_address", vpnConfig.Tunnels[1].CgwInsideAddress)
	d.Set("tunnel2_vgw_inside_address", vpnConfig.Tunnels[1].VgwInsideAddress)
	d.Set("tunnel2_bgp_asn", vpnConfig.Tunnels[1].BGPASN)
	d.Set("tunnel2_bgp_holdtime", vpnConfig.Tunnels[1].BGPHoldTime)
	d.SetId(uuid.New().String())

	return nil
}
