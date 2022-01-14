//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
)

var (
	xmlConfigData = `\u003c?xml version=\"1.0\" encoding=\"UTF-8\"?\u003e\n\u003cvpn_connection id=\"vpn-0c0cf5d09c1101f53\"\u003e\n  \u003ccustomer_gateway_id\u003ecgw-0166f7516dca3e384\u003c/customer_gateway_id\u003e\n  \u003cvpn_gateway_id\u003e\u003c/vpn_gateway_id\u003e\n  \u003cvpn_connection_type\u003eipsec.1\u003c/vpn_connection_type\u003e\n  \u003cipsec_tunnel\u003e\n    \u003ccustomer_gateway\u003e\n      \u003ctunnel_outside_address\u003e\n        \u003cip_address\u003e3.128.248.240\u003c/ip_address\u003e\n      \u003c/tunnel_outside_address\u003e\n      \u003ctunnel_inside_address\u003e\n        \u003cip_address\u003e169.254.192.10\u003c/ip_address\u003e\n        \u003cnetwork_mask\u003e255.255.255.252\u003c/network_mask\u003e\n        \u003cnetwork_cidr\u003e30\u003c/network_cidr\u003e\n      \u003c/tunnel_inside_address\u003e\n      \u003cbgp\u003e\n        \u003casn\u003e64513\u003c/asn\u003e\n        \u003chold_time\u003e30\u003c/hold_time\u003e\n      \u003c/bgp\u003e\n    \u003c/customer_gateway\u003e\n    \u003cvpn_gateway\u003e\n      \u003ctunnel_outside_address\u003e\n        \u003cip_address\u003e1.1.1.1\u003c/ip_address\u003e\n      \u003c/tunnel_outside_address\u003e\n      \u003ctunnel_inside_address\u003e\n        \u003cip_address\u003e169.254.192.9\u003c/ip_address\u003e\n        \u003cnetwork_mask\u003e255.255.255.252\u003c/network_mask\u003e\n        \u003cnetwork_cidr\u003e30\u003c/network_cidr\u003e\n      \u003c/tunnel_inside_address\u003e\n      \u003cbgp\u003e\n        \u003casn\u003e64512\u003c/asn\u003e\n        \u003chold_time\u003e30\u003c/hold_time\u003e\n      \u003c/bgp\u003e\n    \u003c/vpn_gateway\u003e\n    \u003cike\u003e\n      \u003cauthentication_protocol\u003esha1\u003c/authentication_protocol\u003e\n      \u003cencryption_protocol\u003eaes-128-cbc\u003c/encryption_protocol\u003e\n      \u003clifetime\u003e28800\u003c/lifetime\u003e\n      \u003cperfect_forward_secrecy\u003egroup2\u003c/perfect_forward_secrecy\u003e\n      \u003cmode\u003emain\u003c/mode\u003e\n      \u003cpre_shared_key\u003e6a71ad36748141079bed579a51d4f2aa\u003c/pre_shared_key\u003e\n    \u003c/ike\u003e\n    \u003cipsec\u003e\n      \u003cprotocol\u003eesp\u003c/protocol\u003e\n      \u003cauthentication_protocol\u003ehmac-sha1-96\u003c/authentication_protocol\u003e\n      \u003cencryption_protocol\u003eaes-128-cbc\u003c/encryption_protocol\u003e\n      \u003clifetime\u003e3600\u003c/lifetime\u003e\n      \u003cperfect_forward_secrecy\u003egroup2\u003c/perfect_forward_secrecy\u003e\n      \u003cmode\u003etunnel\u003c/mode\u003e\n      \u003cclear_df_bit\u003etrue\u003c/clear_df_bit\u003e\n      \u003cfragmentation_before_encryption\u003etrue\u003c/fragmentation_before_encryption\u003e\n      \u003ctcp_mss_adjustment\u003e1379\u003c/tcp_mss_adjustment\u003e\n      \u003cdead_peer_detection\u003e\n        \u003cinterval\u003e10\u003c/interval\u003e\n        \u003cretries\u003e3\u003c/retries\u003e\n      \u003c/dead_peer_detection\u003e\n    \u003c/ipsec\u003e\n  \u003c/ipsec_tunnel\u003e\n  \u003cipsec_tunnel\u003e\n    \u003ccustomer_gateway\u003e\n      \u003ctunnel_outside_address\u003e\n        \u003cip_address\u003e3.128.248.240\u003c/ip_address\u003e\n      \u003c/tunnel_outside_address\u003e\n      \u003ctunnel_inside_address\u003e\n        \u003cip_address\u003e169.254.192.14\u003c/ip_address\u003e\n        \u003cnetwork_mask\u003e255.255.255.252\u003c/network_mask\u003e\n        \u003cnetwork_cidr\u003e30\u003c/network_cidr\u003e\n      \u003c/tunnel_inside_address\u003e\n      \u003cbgp\u003e\n        \u003casn\u003e64513\u003c/asn\u003e\n        \u003chold_time\u003e30\u003c/hold_time\u003e\n      \u003c/bgp\u003e\n    \u003c/customer_gateway\u003e\n    \u003cvpn_gateway\u003e\n      \u003ctunnel_outside_address\u003e\n        \u003cip_address\u003e2.2.2.2\u003c/ip_address\u003e\n      \u003c/tunnel_outside_address\u003e\n      \u003ctunnel_inside_address\u003e\n        \u003cip_address\u003e169.254.192.13\u003c/ip_address\u003e\n        \u003cnetwork_mask\u003e255.255.255.252\u003c/network_mask\u003e\n        \u003cnetwork_cidr\u003e30\u003c/network_cidr\u003e\n      \u003c/tunnel_inside_address\u003e\n      \u003cbgp\u003e\n        \u003casn\u003e64512\u003c/asn\u003e\n        \u003chold_time\u003e30\u003c/hold_time\u003e\n      \u003c/bgp\u003e\n    \u003c/vpn_gateway\u003e\n    \u003cike\u003e\n      \u003cauthentication_protocol\u003esha1\u003c/authentication_protocol\u003e\n      \u003cencryption_protocol\u003eaes-128-cbc\u003c/encryption_protocol\u003e\n      \u003clifetime\u003e28800\u003c/lifetime\u003e\n      \u003cperfect_forward_secrecy\u003egroup2\u003c/perfect_forward_secrecy\u003e\n      \u003cmode\u003emain\u003c/mode\u003e\n      \u003cpre_shared_key\u003e614413b372ca4e0db65de0fd43cd54b0\u003c/pre_shared_key\u003e\n    \u003c/ike\u003e\n    \u003cipsec\u003e\n      \u003cprotocol\u003eesp\u003c/protocol\u003e\n      \u003cauthentication_protocol\u003ehmac-sha1-96\u003c/authentication_protocol\u003e\n      \u003cencryption_protocol\u003eaes-128-cbc\u003c/encryption_protocol\u003e\n      \u003clifetime\u003e3600\u003c/lifetime\u003e\n      \u003cperfect_forward_secrecy\u003egroup2\u003c/perfect_forward_secrecy\u003e\n      \u003cmode\u003etunnel\u003c/mode\u003e\n      \u003cclear_df_bit\u003etrue\u003c/clear_df_bit\u003e\n      \u003cfragmentation_before_encryption\u003etrue\u003c/fragmentation_before_encryption\u003e\n      \u003ctcp_mss_adjustment\u003e1379\u003c/tcp_mss_adjustment\u003e\n      \u003cdead_peer_detection\u003e\n        \u003cinterval\u003e10\u003c/interval\u003e\n        \u003cretries\u003e3\u003c/retries\u003e\n      \u003c/dead_peer_detection\u003e\n    \u003c/ipsec\u003e\n  \u003c/ipsec_tunnel\u003e\n\u003c/vpn_connection\u003e`
)

func TestAccParseCGWData(t *testing.T) {
	dataSourceName := "data.volterra_parse_aws_cgw_configuration.test"
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{ves_io_schema_ns.ObjectType})
	defer stopFunc()
	os.Setenv("VOLT_API_TEST", "true")
	//os.Setenv("TF_LOG", "TRACE")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "ves-io")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testCGWData(xmlConfigData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "tunnel1_address", "1.1.1.1"),
					resource.TestCheckResourceAttr(dataSourceName, "tunnel2_address", "2.2.2.2"),
				),
			},
		},
	})
}

func testCGWData(xmlConfig string) string {
	return fmt.Sprintf(`
	data "volterra_parse_aws_cgw_configuration" "test" {
		customer_gateway_configuration = "%[1]s"
	}`, xmlConfig)
}
