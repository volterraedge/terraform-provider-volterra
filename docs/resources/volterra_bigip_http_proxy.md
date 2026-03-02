---

page_title: "Volterra: bigip_http_proxy"

description: "The bigip_http_proxy allows CRUD of Bigip Http Proxy resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_bigip_http_proxy
==================================

The Bigip Http Proxy allows CRUD of Bigip Http Proxy resource on Volterra SaaS

~> **Note:** Please refer to [Bigip Http Proxy API docs](https://docs.cloud.f5.com/docs-v2/api/views-bigip-http-proxy) to learn more

Example Usage
-------------

```hcl
resource "volterra_bigip_http_proxy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  ddos_profile {
    // One of the arguments from this list "disable_ddos_mitigation enable_ddos_mitigation" can be set

    enable_ddos_mitigation = true
  }

  origin_pools {
    pools {
      name = "name"

      origin_servers {
        health_checks {
          health_check {
            // One of the arguments from this list "icmp_health_check tcp_health_check" can be set

            tcp_health_check {
              expected_response = ".*"

              send_payload = "send_payload"
            }
          }

          healthy_threshold = "2"

          interval = "10"

          jitter = "1"

          jitter_percent = "25"

          timeout = "1"

          unhealthy_threshold = "5"
        }

        origin_servers {
          // One of the arguments from this list "k8s_service private_ip public_ip public_name" must be set

          public_ip {
            // One of the arguments from this list "ip ipv6" must be set

            ip = "8.8.8.8"
          }
        }

        // One of the arguments from this list "automatic_port lb_port port" must be set

        port = "9080"
      }

      priority = "1"

      weight = "1"
    }
  }

  proxy_config {
    domains = ["www.foo.com"]

    // One of the arguments from this list "http https https_auto_cert" must be set

    http {
      dns_volterra_managed = true

      // One of the arguments from this list "port port_ranges" must be set

      port = "80"
    }
  }
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

`advanced_profile` - (Optional) Advanced Profile. See [Advanced Profile ](#advanced-profile) below for details.

`ddos_profile` - (Required) to protect the origin servers from external DDoS attacks.. See [Ddos Profile ](#ddos-profile) below for details.

`irules` - (Optional) Options for attaching iRules to BIG-IP HTTP Proxy. See [Irules ](#irules) below for details.

`lb_algorithm` - (Optional) The algorithm used for load balance request between origin servers. See [Lb Algorithm ](#lb-algorithm) below for details.

`origin_pools` - (Required) Origin Pools for BIG-IP HTTP Proxy. See [Origin Pools ](#origin-pools) below for details.

`proxy_advertisement` - (Optional) Proxy Advertisement choice. See [Proxy Advertisement ](#proxy-advertisement) below for details.

`proxy_config` - (Required) The type of load balancer, can be "http" or "https". See [Proxy Config ](#proxy-config) below for details.

### Advanced Profile

Advanced Profile.

###### One of the arguments from this list "disable, enable_default_profile" must be set

`disable` - (Optional) Disable Advanced options (`Bool`).

`enable_default_profile` - (Optional) Enable Default Profile (`Bool`).

### Ddos Profile

to protect the origin servers from external DDoS attacks..

###### One of the arguments from this list "disable_ddos_mitigation, enable_ddos_mitigation" can be set

`disable_ddos_mitigation` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_ddos_mitigation` - (Optional) x-displayName: "Enable" (`Bool`).

### Irules

Options for attaching iRules to BIG-IP HTTP Proxy.

`irules` - (Optional) Options for attaching iRules to BIG-IP HTTP Proxy. See [ref](#ref) below for details.

### Lb Algorithm

The algorithm used for load balance request between origin servers.

###### One of the arguments from this list "round_robin" must be set

`round_robin` - (Optional) Requests are sent to all the eligible origin servers in round-robin fashion (`Bool`).

### Origin Pools

Origin Pools for BIG-IP HTTP Proxy.

`pools` - (Optional) List of Origin Pools. See [Origin Pools Pools ](#origin-pools-pools) below for details.

### Proxy Advertisement

Proxy Advertisement choice.

###### One of the arguments from this list "advertise_custom, do_not_advertise" must be set

`advertise_custom` - (Optional) Advertise this load balancer on specific sites. See [Advertise Choice Advertise Custom ](#advertise-choice-advertise-custom) below for details.

`do_not_advertise` - (Optional) Do not advertise this proxy (`Bool`).

### Proxy Config

The type of load balancer, can be "http" or "https".

`domains` - (Required) Domains also indicate the list of names for which DNS resolution will be done by VER (`String`).

###### One of the arguments from this list "http, https, https_auto_cert" must be set

`http` - (Optional) HTTP Load balancer.. See [Loadbalancer Type Http ](#loadbalancer-type-http) below for details.

`https` - (Optional) User is responsible for managing DNS to this Load Balancer.. See [Loadbalancer Type Https ](#loadbalancer-type-https) below for details.

`https_auto_cert` - (Optional) DNS records will be managed by Volterra.. See [Loadbalancer Type Https Auto Cert ](#loadbalancer-type-https-auto-cert) below for details.

### Advertise Choice Advertise Custom

Advertise this load balancer on specific sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Advertise Custom Advertise Where ](#advertise-custom-advertise-where) below for details.

### Advertise Choice Do Not Advertise

Do not advertise this proxy.

### Advertise Custom Advertise Where

Where should this load balancer be available.

###### One of the arguments from this list "advertise_on_public, cloud_edge_segment, segment, site, site_segment, virtual_network, virtual_site, virtual_site_segment, virtual_site_with_vip, vk8s_service" must be set

`advertise_on_public` - (Optional) Advertise this load balancer on public network. See [Choice Advertise On Public ](#choice-advertise-on-public) below for details.

`site` - (Optional) Advertise on a customer site and a given network.. See [Choice Site ](#choice-site) below for details.

`site_segment` - (Optional) Advertise on a segment on a site. See [Choice Site Segment ](#choice-site-segment) below for details.

`virtual_network` - (Optional) Advertise on a virtual network. See [Choice Virtual Network ](#choice-virtual-network) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Choice Virtual Site ](#choice-virtual-site) below for details.

`virtual_site_segment` - (Optional) Advertise on a segment on a virtual site. See [Choice Virtual Site Segment ](#choice-virtual-site-segment) below for details.

`virtual_site_with_vip` - (Optional) Advertise on a customer virtual site and a given network and IP.. See [Choice Virtual Site With Vip ](#choice-virtual-site-with-vip) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Choice Vk8s Service ](#choice-vk8s-service) below for details.

###### One of the arguments from this list "port, port_ranges, use_default_port" must be set

`port` - (Optional) Port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

`use_default_port` - (Optional) Inherit the Load Balancer's Listen Port. (`Bool`).

### Choice Advertise On Public

Advertise this load balancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Choice Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Choice Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Choice Disable

Disable Advanced options.

### Choice Enable Default Profile

Enable Default Profile.

### Choice Icmp Health Check

ICMP Health Check.

### Choice K8s Service

Specify origin server with K8s service name and site information.

###### One of the arguments from this list "inside_network, outside_network, vk8s_networks" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`vk8s_networks` - (Optional) origin server are on vK8s network on the site (`Bool`).

`protocol` - (Optional) Protocol to be used in the discovery. (`String`).

###### One of the arguments from this list "service_name, service_selector" must be set

`service_name` - (Optional) Both namespace and cluster-id are optional. (`String`).

`service_selector` - (Optional) discovery has to happen. This implicit label is added to service_selector. See [Service Info Service Selector ](#service-info-service-selector) below for details.(Deprecated)

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [K8s Service Site Locator ](#k8s-service-site-locator) below for details.

`snat_pool` - (Optional) x-displayName: "SNAT Pool Configuration". See [K8s Service Snat Pool ](#k8s-service-snat-pool) below for details.

### Choice Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Choice Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Choice Private Ip

Specify origin server with Private IP.

###### One of the arguments from this list "inside_network, outside_network, segment" must be set

`inside_network` - (Optional) Inside network on the site (`Bool`).

`outside_network` - (Optional) Outside network on the site (`Bool`).

`segment` - (Optional) Segment where this origin server is located. See [ref](#ref) below for details.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) Private IPV4 address (`String`).

`ipv6` - (Optional) Private IPV6 address (`String`).

`site_locator` - (Required) Site or Virtual site where this origin server is located. See [Private Ip Site Locator ](#private-ip-site-locator) below for details.

`snat_pool` - (Optional) x-displayName: "SNAT Pool Configuration". See [Private Ip Snat Pool ](#private-ip-snat-pool) below for details.

### Choice Public Ip

Specify origin server with public IP.

###### One of the arguments from this list "ip, ipv6" must be set

`ip`- (Optional) Public IPV4 address (`String`).

`ipv6` - (Optional) Public IPV6 address (`String`).

### Choice Public Name

Specify origin server with public DNS name.

`dns_name` - (Required) DNS Name (`String`).

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

### Choice Site

Advertise on a customer site and a given network..

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Choice Site Segment

Advertise on a segment on a site.

`ip` - (Required) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`segment` - (Required) x-required. See [ref](#ref) below for details.

`site` - (Required) x-required. See [ref](#ref) below for details.

### Choice Tcp Health Check

TCP Health Check.

`expected_response` - (Required) Specifies a regular expression pattern which will be matched against response payload (`String`).

`send_payload` - (Required) Text string sent in the request (`String`).

### Choice Virtual Network

Advertise on a virtual network.

###### One of the arguments from this list "default_v6_vip, specific_v6_vip" can be set

`default_v6_vip` - (Optional) Use the default VIP, system allocated or configured in the virtual network (`Bool`).

`specific_v6_vip` - (Optional) Use given IPV6 address as VIP on virtual Network (`String`).

###### One of the arguments from this list "default_vip, specific_vip" can be set

`default_vip` - (Optional) Use the default VIP, system allocated or configured in the virtual network (`Bool`).

`specific_vip` - (Optional) Use given IPV4 address as VIP on virtual Network (`String`).

`virtual_network` - (Required) Select network reference. See [ref](#ref) below for details.

### Choice Virtual Site

Advertise on a customer virtual site and a given network..

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Choice Virtual Site Segment

Advertise on a segment on a virtual site.

`ip` - (Required) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`segment` - (Required) x-required. See [ref](#ref) below for details.

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Choice Virtual Site With Vip

Advertise on a customer virtual site and a given network and IP..

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Choice Vk8s Service

Advertise on vK8s Service Network on RE..

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Coalescing Choice Default Coalescing

or Cipher suite configuration.

### Coalescing Choice Strict Coalescing

and/or Cipher suite configuration.

### Crl Choice No Crl

Client certificate revocation status is not verified.

### Ddos Mitigation Choice Disable Ddos Mitigation

x-displayName: "Disable".

### Ddos Mitigation Choice Enable Ddos Mitigation

x-displayName: "Enable".

### Default Lb Choice Default Loadbalancer

x-displayName: "Yes".

### Default Lb Choice Non Default Loadbalancer

x-displayName: "No".

### Header Transformation Choice Default Header Transformation

Normalize the headers to lower case.

### Header Transformation Choice Legacy Header Transformation

Use old header transformation if configured earlier.

### Header Transformation Choice Preserve Case Header Transformation

Preserves the original case of headers without any modifications..

### Header Transformation Choice Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

### Health Checks Health Check

List of Health Checks.

###### One of the arguments from this list "icmp_health_check, tcp_health_check" can be set

`icmp_health_check` - (Optional) ICMP Health Check (`Bool`).

`tcp_health_check` - (Optional) TCP Health Check. See [Choice Tcp Health Check ](#choice-tcp-health-check) below for details.

### Http Protocol Choice Http Protocol Enable V1 Only

Enable HTTP/1.1 for downstream connections.

`header_transformation` - (Optional) the stateful formatter will take effect, and the stateless formatter will be disregarded.. See [Http Protocol Enable V1 Only Header Transformation ](#http-protocol-enable-v1-only-header-transformation) below for details.

### Http Protocol Choice Http Protocol Enable V1 V2

Enable both HTTP/1.1 and HTTP/2 for downstream connections.

### Http Protocol Choice Http Protocol Enable V2 Only

Enable HTTP/2 for downstream connections.

### Http Protocol Enable V1 Only Header Transformation

the stateful formatter will take effect, and the stateless formatter will be disregarded..

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Https Coalescing Options

Options for coalescing TLS for multiple HTTPS Load Balancers.

###### One of the arguments from this list "default_coalescing, strict_coalescing" must be set

`default_coalescing` - (Optional) or Cipher suite configuration (`Bool`).

`strict_coalescing` - (Optional) and/or Cipher suite configuration (`Bool`).

### Https Header Transformation Type

Header transformation options for response headers to the client.

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Https Http Protocol Options

HTTP protocol configuration options for downstream connections..

###### One of the arguments from this list "http_protocol_enable_v1_only, http_protocol_enable_v1_v2, http_protocol_enable_v2_only" must be set

`http_protocol_enable_v1_only` - (Optional) Enable HTTP/1.1 for downstream connections. See [Http Protocol Choice Http Protocol Enable V1 Only ](#http-protocol-choice-http-protocol-enable-v1-only) below for details.

`http_protocol_enable_v1_v2` - (Optional) Enable both HTTP/1.1 and HTTP/2 for downstream connections (`Bool`).

`http_protocol_enable_v2_only` - (Optional) Enable HTTP/2 for downstream connections (`Bool`).

### Https Auto Cert Coalescing Options

Options for coalescing TLS for multiple HTTPS Load Balancers.

###### One of the arguments from this list "default_coalescing, strict_coalescing" must be set

`default_coalescing` - (Optional) or Cipher suite configuration (`Bool`).

`strict_coalescing` - (Optional) and/or Cipher suite configuration (`Bool`).

### Https Auto Cert Header Transformation Type

Header transformation options for response headers to the client.

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Https Auto Cert Http Protocol Options

HTTP protocol configuration options for downstream connections..

###### One of the arguments from this list "http_protocol_enable_v1_only, http_protocol_enable_v1_v2, http_protocol_enable_v2_only" must be set

`http_protocol_enable_v1_only` - (Optional) Enable HTTP/1.1 for downstream connections. See [Http Protocol Choice Http Protocol Enable V1 Only ](#http-protocol-choice-http-protocol-enable-v1-only) below for details.

`http_protocol_enable_v1_v2` - (Optional) Enable both HTTP/1.1 and HTTP/2 for downstream connections (`Bool`).

`http_protocol_enable_v2_only` - (Optional) Enable HTTP/2 for downstream connections (`Bool`).

### Https Auto Cert Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### K8s Service Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### K8s Service Snat Pool

x-displayName: "SNAT Pool Configuration".

###### One of the arguments from this list "no_snat_pool, snat_pool" can be set

`no_snat_pool` - (Optional) No configured SNAT Pool to reach Origin Server (`Bool`).

`snat_pool` - (Optional) Configure SNAT Pool to reach Origin Server. See [Snat Pool Choice Snat Pool ](#snat-pool-choice-snat-pool) below for details.

### Lb Algorithm Choice Round Robin

Requests are sent to all the eligible origin servers in round-robin fashion.

### Loadbalancer Type Http

HTTP Load balancer..

`dns_volterra_managed` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal. (`Bool`).

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) HTTP port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

### Loadbalancer Type Https

User is responsible for managing DNS to this Load Balancer..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`coalescing_options` - (Optional) Options for coalescing TLS for multiple HTTPS Load Balancers. See [Https Coalescing Options ](#https-coalescing-options) below for details.

`connection_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 2 minutes. (`Int`).

###### One of the arguments from this list "default_loadbalancer, non_default_loadbalancer" can be set

`default_loadbalancer` - (Optional) x-displayName: "Yes" (`Bool`).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (`Bool`).

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Https Header Transformation Type ](#https-header-transformation-type) below for details.(Deprecated)

`http_protocol_options` - (Optional) HTTP protocol configuration options for downstream connections.. See [Https Http Protocol Options ](#https-http-protocol-options) below for details.

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" must be set

`disable_path_normalize` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_path_normalize` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) HTTPS port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

###### One of the arguments from this list "append_server_name, default_header, pass_through, server_name" can be set

`append_server_name` - (Optional) If header value is already present, it is not overwritten and passed as-is. (`String`).

`default_header` - (Optional) Response header name is “server” and value is “volt-adc” (`Bool`).

`pass_through` - (Optional) Pass existing server header as is. If server header is absent, a new header is not appended. (`Bool`).

`server_name` - (Optional) This will overwrite existing values, if any, for the server header. (`String`).

###### One of the arguments from this list "tls_cert_params, tls_parameters" must be set

`tls_cert_params` - (Optional) Select/Add one or more TLS Certificate objects to associate with this Load Balancer. See [Tls Certificates Choice Tls Cert Params ](#tls-certificates-choice-tls-cert-params) below for details.

`tls_parameters` - (Optional) Upload a TLS certificate covering all domain names for this Load Balancer. See [Tls Certificates Choice Tls Parameters ](#tls-certificates-choice-tls-parameters) below for details.

### Loadbalancer Type Https Auto Cert

DNS records will be managed by Volterra..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`coalescing_options` - (Optional) Options for coalescing TLS for multiple HTTPS Load Balancers. See [Https Auto Cert Coalescing Options ](#https-auto-cert-coalescing-options) below for details.

`connection_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 2 minutes. (`Int`).

###### One of the arguments from this list "default_loadbalancer, non_default_loadbalancer" can be set

`default_loadbalancer` - (Optional) For traffic terminating at this load balancer, the certificate associated with the first configured domain will be used for TLS termination. (`Bool`).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (`Bool`).

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Https Auto Cert Header Transformation Type ](#https-auto-cert-header-transformation-type) below for details.(Deprecated)

`http_protocol_options` - (Optional) HTTP protocol configuration options for downstream connections.. See [Https Auto Cert Http Protocol Options ](#https-auto-cert-http-protocol-options) below for details.

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" must be set

`disable_path_normalize` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_path_normalize` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "port, port_ranges" can be set

`port` - (Optional) HTTPS port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

###### One of the arguments from this list "append_server_name, default_header, pass_through, server_name" can be set

`append_server_name` - (Optional) If header value is already present, it is not overwritten and passed as-is. (`String`).

`default_header` - (Optional) Response header name is “server” and value is “volt-adc” (`Bool`).

`pass_through` - (Optional) Pass existing server header as is. If server header is absent, a new header is not appended. (`Bool`).

`server_name` - (Optional) This will overwrite existing values, if any, for the server header. (`String`).

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Https Auto Cert Tls Config ](#https-auto-cert-tls-config) below for details.

### Mtls Choice No Mtls

x-displayName: "Disable".

### Mtls Choice Use Mtls

x-displayName: "Enable".

`client_certificate_optional` - (Optional) the connection will be accepted. (`Bool`).

###### One of the arguments from this list "crl, no_crl" can be set

`crl` - (Optional) Specify the CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

`no_crl` - (Optional) Client certificate revocation status is not verified (`Bool`).

###### One of the arguments from this list "trusted_ca, trusted_ca_url" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Load Balancer. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Load Balancer (`String`).

###### One of the arguments from this list "xfcc_disabled, xfcc_options" can be set

`xfcc_disabled` - (Optional) No X-Forwarded-Client-Cert header will be added (`Bool`).

`xfcc_options` - (Optional) X-Forwarded-Client-Cert header will be added with the configured fields. See [Xfcc Header Xfcc Options ](#xfcc-header-xfcc-options) below for details.

### Network Choice Inside Network

Inside network on the site.

### Network Choice Outside Network

Outside network on the site.

### Network Choice Vk8s Networks

origin server are on vK8s network on the site.

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Origin Pools Pools

List of Origin Pools.

`name` - (Required) Name of the origin pool (`String`).

`origin_servers` - (Required) Create origin pool with multiple backend servers. See [Pools Origin Servers ](#pools-origin-servers) below for details.

`priority` - (Optional) are made active as per the increasing priority. (`Int`).

`weight` - (Optional) the pool (`Int`).

### Origin Servers Health Checks

x-required.

`health_check` - (Required) List of Health Checks. See [Health Checks Health Check ](#health-checks-health-check) below for details.

`healthy_threshold` - (Required) required to mark a host healthy. (`Int`).

`interval` - (Required) Time interval in seconds between two health check requests (`Int`).

`jitter` - (Optional) increase the wait time between two healthcheck requests. (`Int`).(Deprecated)

`jitter_percent` - (Optional) check requests (`Int`).(Deprecated)

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

`unhealthy_threshold` - (Required) this threshold is ignored and the host is considered unhealthy immediately. (`Int`).

### Origin Servers Origin Servers

List of origin servers for Proxy.

###### One of the arguments from this list "k8s_service, private_ip, public_ip, public_name" must be set

`k8s_service` - (Optional) Specify origin server with K8s service name and site information. See [Choice K8s Service ](#choice-k8s-service) below for details.

`private_ip` - (Optional) Specify origin server with Private IP. See [Choice Private Ip ](#choice-private-ip) below for details.

`public_ip` - (Optional) Specify origin server with public IP. See [Choice Public Ip ](#choice-public-ip) below for details.

`public_name` - (Optional) Specify origin server with public DNS name. See [Choice Public Name ](#choice-public-name) below for details.

### Path Normalize Choice Disable Path Normalize

x-displayName: "Disable".

### Path Normalize Choice Enable Path Normalize

x-displayName: "Enable".

### Pools Origin Servers

Create origin pool with multiple backend servers.

`health_checks` - (Required) x-required. See [Origin Servers Health Checks ](#origin-servers-health-checks) below for details.

`origin_servers` - (Required) List of origin servers for Proxy. See [Origin Servers Origin Servers ](#origin-servers-origin-servers) below for details.

###### One of the arguments from this list "automatic_port, lb_port, port" must be set

`automatic_port` - (Optional) if TLS is enabled at Origin Pool and 80 if TLS is disabled (`Bool`).

`lb_port` - (Optional) Endpoint port is selected based on loadbalancer port (`Bool`).

`port` - (Optional) Endpoint service is available on this port (`Int`).

### Port Choice Automatic Port

if TLS is enabled at Origin Pool and 80 if TLS is disabled.

### Port Choice Lb Port

Endpoint port is selected based on loadbalancer port.

### Port Choice Use Default Port

Inherit the Load Balancer's Listen Port..

### Private Ip Site Locator

Site or Virtual site where this origin server is located.

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Private Ip Snat Pool

x-displayName: "SNAT Pool Configuration".

###### One of the arguments from this list "no_snat_pool, snat_pool" can be set

`no_snat_pool` - (Optional) No configured SNAT Pool to reach Origin Server (`Bool`).

`snat_pool` - (Optional) Configure SNAT Pool to reach Origin Server. See [Snat Pool Choice Snat Pool ](#snat-pool-choice-snat-pool) below for details.

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Secret Info Oneof Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Secret Info Oneof Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

### Server Header Choice Default Header

Response header name is “server” and value is “volt-adc”.

### Server Header Choice Pass Through

Pass existing server header as is. If server header is absent, a new header is not appended..

### Service Info Service Selector

discovery has to happen. This implicit label is added to service_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Snat Pool Choice No Snat Pool

No configured SNAT Pool to reach Origin Server.

### Snat Pool Choice Snat Pool

Configure SNAT Pool to reach Origin Server.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### Tls Cert Params Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Tls Certificates Choice Tls Cert Params

Select/Add one or more TLS Certificate objects to associate with this Load Balancer.

`certificates` - (Required) Select one or more certificates with any domain names.. See [ref](#ref) below for details.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Cert Params Tls Config ](#tls-cert-params-tls-config) below for details.

### Tls Certificates Choice Tls Parameters

Upload a TLS certificate covering all domain names for this Load Balancer.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Parameters Tls Certificates ](#tls-parameters-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Parameters Tls Config ](#tls-parameters-tls-config) below for details.

### Tls Parameters Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Tls Parameters Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### V6 Vip Choice Default V6 Vip

Use the default VIP, system allocated or configured in the virtual network.

### Vip Choice Default Vip

Use the default VIP, system allocated or configured in the virtual network.

### Xfcc Header Xfcc Disabled

No X-Forwarded-Client-Cert header will be added.

### Xfcc Header Xfcc Options

X-Forwarded-Client-Cert header will be added with the configured fields.

`xfcc_header_elements` - (Required) X-Forwarded-Client-Cert header elements to be added to requests (`List of Strings`).

Attribute Reference
-------------------

*   `id` - This is the id of the configured bigip_http_proxy.
