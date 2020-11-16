---
page_title: "Volterra: volterra_parse_aws_cgw_configuration Data Source"

description: "Data source for volterra_parse_aws_cgw_configuration resource"
---

# Data Source: volterra_parse_aws_cgw_configuration

Parses the aws gw configuration provided and gets all the info in a unsorted manner.


## Example Usage

```hcl
data "volterra_parse_aws_cgw_configuration" "example" {
  customer_gateway_configuration = "<aws_vpn_customer_gw_config in xml format>"
}
```

## Argument Reference

* `customer_gateway_configuration` - (Required) Customer gateway configuration received (in XML format) after creating vpn connection object on aws

## Attribute Reference

`tunnel1_address` - Public IP address of the first VPN tunnel.
`tunnel1_cgw_inside_address` - Customer Gateway Side inside address of the first VPN tunnel.
`tunnel1_vgw_inside_address` - AWS VGW side inside address of the first VPN tunnel.
`tunnel1_preshared_key` - Preshared key of the first VPN tunnel.
`tunnel1_bgp_asn` - BGP asn number of the first VPN tunnel.
`tunnel1_bgp_holdtime` - The bgp holdtime of the first VPN tunnel.
`tunnel2_address` - Public IP address of the second VPN tunnel.
`tunnel2_cgw_inside_address` - Customer Gateway Side inside address of the second VPN tunnel.
`tunnel2_vgw_inside_address` - AWS VGW side inside address of the second VPN tunnel.
`tunnel2_preshared_key` - Preshared key of the second VPN tunnel.
`tunnel2_bgp_asn` - BGP asn number of the second VPN tunnel.
`tunnel2_bgp_holdtime` - BGP holdtime of the second VPN tunnel.