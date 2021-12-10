---

page_title: "Volterra: volterra_tgw_vpn_tunnels"

description: "The volterra_tgw_vpn_tunnels runs set_vpc_ip_prefixes api on aws_tgw_site"
----------------------------------------------------------------------------------------

Resource volterra_tgw_vpn_tunnels
=================================

====================================

Configure VPN tunnels for aws_tgw_site object

~> **Note:** Please refer to [Set vpn tunnels docs](https://volterra.io/docs/api/views-aws-tgw-site#operation/ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPNTunnels) to learn more ~> **Note:** This resource works well as part of Volterra's aws_tgw_site site bring up terraform script. Please do not use this on already provisioned aws_tgw_site site.

Example Usage
-------------

---

```hcl
resource "volterra_tgw_vpn_tunnels" "example" {
  name            = "aws-tgw-site1"
  namespace       = "system"
  vpn_tunnel_config  {
    node_name = "master-0"
    node_id = "node-1234"
    tunnel_remote_ips = ["1.1.1.1","2.2.2.2"]
    type = "HUB"
  }
}

```

### Argument Reference

---

-	`name` - (Required) The value of aws_tgw_site name has to follow DNS-1035 format. (`String`).

-	`namespace` - (Required) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

-	`vpn_tunnel_config` - (Required) List of vpc ip prefixes

	-	`node_name` - (Required) Name of the node (`String`).
	-	`node_id` - (Required) Name of node-id (`String`).
	-	`tunnel_remote_ips` - (Required) List of remote IP on AWS tgw side (`String`).
	-	`type` - (Optional) Type of VPN tunnel, valid values are `HUB` and `SERVICES`.
