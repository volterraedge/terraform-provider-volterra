---
page_title: "Volterra: volterra_http_loadbalancer_state Data Source"

description: "Data source for volterra_http_loadbalancer_state resource"
---

# Data Source: volterra_http_loadbalancer_state

volterra_http_loadbalancer_state is used to get the state information related to a http loadbalancer. This includes cname, ip_address and auto_cert info related state


## Example Usage

```hcl
data "volterra_http_loadbalancer_state" "example" {
  name = "aalloc-1"
  namespace = "app-ns"
}
```

## Argument Reference

`name` - (Required) Name of the http loadbalancer

`namespace` - (Required) Namespace where the http loadbalancer is configured. Example `app-ns`

## Attribute Reference

`cname` - Is the cname assigned to the http loadbalancer by Volterra.

`ip_address` - IP address associated with the http loadbalancer.

`state` - State of the loadbalancer.

`auto_cert_state` - State of the auto certificate generation.

`auto_cert_info` - Auto certificate related information. See [AutoCertInfo](#AutoCertInfo) below for details.

### AutoCertInfo

Information related to auto certificate

`auto_cert_state` - State of auto certificate generation.

`auto_cert_expiry` - Auto certificate expiry timestamp.

`auto_cert_subject` - Subject of the auto certificate.

`auto_cert_issuer` - Issuer of the auto certificate.

`dns_records` - DNS Records that are to be added by user in their DNS domain. Currently, this will be populated when auto certificates are desired but DNS delegation is not enabled. See [DNSRecord](#DNSRecord) below for details.

### DNSRecord

DNSRecord info

- `name` - Name of the DNS record
- `type` - DNS record type
- `value` - DNS record value