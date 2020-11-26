---
page_title: "Volterra: volterra_tunnel Data Source"

description: "Data source for volterra_tunnel resource"
---

# Data Source: volterra_tunnel

volterra_tunnel is a way to get tunnel information


## Example Usage

```hcl
data "volterra_tunnel" "example" {
  name = "tun-1"
  namespace = "system"
}
```

## Argument Reference

* `name` - (Required) Name of the tunnel

* `namespace` - (Required) Namespace where the tunnel is configure. Example `system`

## Attribute Reference

* `type` - Type of the tunnel created

* `ipsec_psk_b64` - If tunnel type is IPSEC_PSK, it will return the ipsec pre shared key on base64 format