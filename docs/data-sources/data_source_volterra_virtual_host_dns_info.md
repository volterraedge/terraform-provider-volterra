---
page_title: "Volterra: virtual_host_dns_info Data Source"

description: "Data source for virtual_host_dns_info resource"
---

# Data Source: volterra_virtual_host_dns_info

volterra_virtual_host_dns_info is a way to get DNS information like dns cname or IP address for a given virtual host


## Example Usage

```hcl
data volterra_virtual_host_dns_info" "example" {
  name = "vh-1"
}
```

## Argument Reference

* `name` - (Required) Name of the virtual host whose dns info must be queried


## Attribute Reference

* `host_name` - Host name to be used for the virtual host

* `tenant_name` - Tenant name configured in volterra

* `dns_info` - DNS information for this virtual host
  * `ip_address` - IP address associated with the virtual host