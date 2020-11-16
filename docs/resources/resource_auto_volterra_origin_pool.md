---

page_title: "Volterra: origin_pool"

description: "The origin_pool allows CRUD of Origin Pool resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_origin_pool
=============================

The Origin Pool allows CRUD of Origin Pool resource on Volterra SaaS

~> **Note:** Please refer to [Origin Pool API docs](https://volterra.io/docs/api/views-origin-pool) to learn more

Example Usage
-------------

```hcl
resource "volterra_origin_pool" "example" {
  name                   = "acmecorp-web"
  namespace              = "staging"
  endpoint_selection     = ["endpoint_selection"]
  loadbalancer_algorithm = ["loadbalancer_algorithm"]

  origin_servers {
    // One of the arguments from this list "public_ip public_name private_ip private_name k8s_service consul_service custom_endpoint_object" must be set

    consul_service {
      // One of the arguments from this list "inside_network outside_network" must be set
      inside_network = true
      service_name   = "service_name"

      site_locator {
        // One of the arguments from this list "virtual_site site" must be set

        site {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }
    }
  }

  port = ["9080"]

  // One of the arguments from this list "no_tls use_tls" must be set
  no_tls = true
}

```

Argument Reference
------------------

### Metadata Argument Reference

`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).

`description` - (Optional) Human readable description for the object (`String`).

`disable` - (Optional) A value of true will administratively disable the object (`Bool`).

`labels` - (Optional) by selector expression (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

### Spec Argument Reference

`endpoint_selection` - (Required) Policy for selection of endpoints from local site or remote site or both (`String`).

`healthcheck` - (Optional) Reference to healthcheck configuration objects. See [ref](#ref) below for details.

`loadbalancer_algorithm` - (Required) loadbalancer_algorithm to determine which host is selected. (`String`).

`origin_servers` - (Required) List of origin servers in this pool. See [Origin Servers ](#origin-servers) below for details.

`port` - (Required) Endpoint service is available on this port (`Int`).

`no_tls` - (Optional) Origin servers do not use TLS (bool).

`use_tls` - (Optional) Origin servers use TLS. See [Use Tls ](#use-tls) below for details.

### Origin Servers

List of origin servers in this pool.

`consul_service` - (Optional) Specify origin server with Hashi Corp Consul service name and site information. See [Consul Service ](#consul-service) below for details.

`custom_endpoint_object` - (Optional) Specify origin server with a reference to endpoint object. See [Custom Endpoint Object ](#custom-endpoint-object) below for details.

`k8s_service` - (Optional) Specify origin server with k8s service name and site information. See [K8s Service ](#k8s-service) below for details.

`private_ip` - (Optional) Specify origin server with private or public IP address and site information. See [Private Ip ](#private-ip) below for details.

`private_name` - (Optional) Specify origin server with private or public DNS name and site information. See [Private Name ](#private-name) below for details.

`public_ip` - (Optional) Specify origin server with public IP. See [Public Ip ](#public-ip) below for details.

`public_name` - (Optional) Specify origin server with public DNS name. See [Public Name ](#public-name) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Tls Config

TLS parameters such as min/max TLS version and ciphers.

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Custom Security ](#custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers with strong crypto algorithms. (bool).

`low_security` - (Optional) Low Security chooses TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (bool).

`medium_security` - (Optional) Medium Security chooses TLS v1.0+ with only PFS ciphers and medium strength crypto algorithms. (bool).

### Use Tls

Origin servers use TLS.

`no_mtls` - (Optional) Do not use MTLS for this pool (bool).

`use_mtls` - (Optional) Use MTLS for this pool using provided certificates. See [Use Mtls ](#use-mtls) below for details.

`skip_server_verification` - (Optional) Skip Server Verification (bool).

`use_server_verification` - (Optional) Verify server identity using provided information. See [Use Server Verification ](#use-server-verification) below for details.

`volterra_trusted_ca` - (Optional) Perform origin server verification using Volterra default trusted CA list (bool).

`disable_sni` - (Optional) Do not use SNI. (bool).

`sni` - (Optional) SNI value to be used. (`String`).

`use_host_header_as_sni` - (Optional) Use the host header as SNI - not supported. (bool).

`tls_config` - (Required) TLS parameters such as min/max TLS version and ciphers. See [Tls Config ](#tls-config) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured origin_pool.
