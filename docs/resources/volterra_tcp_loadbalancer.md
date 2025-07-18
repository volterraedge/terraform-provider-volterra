---

page_title: "Volterra: tcp_loadbalancer"

description: "The tcp_loadbalancer allows CRUD of Tcp Loadbalancer resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_tcp_loadbalancer
==================================

The Tcp Loadbalancer allows CRUD of Tcp Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Tcp Loadbalancer API docs](https://docs.cloud.f5.com/docs-v2/api/views-tcp-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_tcp_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "advertise_custom advertise_on_public advertise_on_public_default_vip do_not_advertise" must be set

  advertise_on_public_default_vip = true

  // One of the arguments from this list "do_not_retract_cluster retract_cluster" must be set

  retract_cluster = true

  // One of the arguments from this list "hash_policy_choice_least_active hash_policy_choice_random hash_policy_choice_round_robin hash_policy_choice_source_ip_stickiness" must be set

  hash_policy_choice_random = true

  // One of the arguments from this list "tcp tls_tcp tls_tcp_auto_cert" must be set

  tcp = true

  // One of the arguments from this list "listen_port port_ranges" must be set

  listen_port = "0"

  // One of the arguments from this list "active_service_policies no_service_policies service_policies_from_namespace" must be set

  service_policies_from_namespace = true

  // One of the arguments from this list "default_lb_with_sni no_sni sni" must be set

  no_sni = true
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

###### One of the arguments from this list "advertise_custom, advertise_on_public, advertise_on_public_default_vip, do_not_advertise" must be set

`advertise_custom` - (Optional) Advertise this VIP on specific sites. See [Advertise Choice Advertise Custom ](#advertise-choice-advertise-custom) below for details.

`advertise_on_public` - (Optional) Advertise this load balancer on public network. See [Advertise Choice Advertise On Public ](#advertise-choice-advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Advertise this load balancer on public network with default VIP (`Bool`).

`do_not_advertise` - (Optional) Do not advertise this load balancer (`Bool`).

###### One of the arguments from this list "do_not_retract_cluster, retract_cluster" must be set

`do_not_retract_cluster` - (Optional) configuration. (`Bool`).

`retract_cluster` - (Optional) for route (`Bool`).

`dns_volterra_managed` - (Optional) This requires the domain to be delegated to F5XC using the Delegated Domain feature. (`Bool`).

`domains` - (Optional) Domains also indicate the list of names for which DNS resolution will be done by VER (`List of String`).

###### One of the arguments from this list "hash_policy_choice_least_active, hash_policy_choice_random, hash_policy_choice_round_robin, hash_policy_choice_source_ip_stickiness" must be set

`hash_policy_choice_least_active` - (Optional) Connections are sent to origin server that has least active connections (`Bool`).

`hash_policy_choice_random` - (Optional) Connections are sent to all eligible origin servers in random fashion (`Bool`).

`hash_policy_choice_round_robin` - (Optional) Connections are sent to all eligible origin servers in round robin fashion (`Bool`).

`hash_policy_choice_source_ip_stickiness` - (Optional) Connections are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server (`Bool`).

`idle_timeout` - (Optional) The amount of time that a stream can exist without upstream or downstream activity, in milliseconds. (`Int`).

###### One of the arguments from this list "tcp, tls_tcp, tls_tcp_auto_cert" must be set

`tcp` - (Optional) TCP Load Balancer. (`Bool`).

`tls_tcp` - (Optional) User is responsible for managing DNS to this load balancer.. See [Loadbalancer Type Tls Tcp ](#loadbalancer-type-tls-tcp) below for details.

`tls_tcp_auto_cert` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal(only for Domains not managed by F5 Distributed Cloud).. See [Loadbalancer Type Tls Tcp Auto Cert ](#loadbalancer-type-tls-tcp-auto-cert) below for details.

`origin_pools_weights` - (Optional) Origin pools and weights used for this load balancer.. See [Origin Pools Weights ](#origin-pools-weights) below for details.

###### One of the arguments from this list "listen_port, port_ranges" must be set

`listen_port` - (Optional) Listen Port for this load balancer (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

###### One of the arguments from this list "active_service_policies, no_service_policies, service_policies_from_namespace" must be set

`active_service_policies` - (Optional) Apply the specified list of service policies and bypass the namespace service policy set. See [Service Policy Choice Active Service Policies ](#service-policy-choice-active-service-policies) below for details.

`no_service_policies` - (Optional) Do not apply any service policies i.e. bypass the namespace service policy set (`Bool`).

`service_policies_from_namespace` - (Optional) Apply the active service policies configured as part of the namespace service policy set (`Bool`).

###### One of the arguments from this list "default_lb_with_sni, no_sni, sni" must be set

`default_lb_with_sni` - (Optional) Also enables usage as Default LB for Non SNI Clients (`Bool`).

`no_sni` - (Optional) Loadbalancer without Server Name Indication support (`Bool`).

`sni` - (Optional) Enables Server Name Indication for Loadbalancer (`Bool`).

### Origin Pools Weights

Origin pools and weights used for this load balancer..

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

###### One of the arguments from this list "cluster, pool" must be set

`cluster` - (Optional) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Optional) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Advertise Choice Advertise Custom

Advertise this VIP on specific sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Advertise Custom Advertise Where ](#advertise-custom-advertise-where) below for details.

### Advertise Choice Advertise On Public

Advertise this load balancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Advertise Custom Advertise Where

Where should this load balancer be available.

###### One of the arguments from this list "advertise_on_public, site, site_segment, virtual_network, virtual_site, virtual_site_segment, virtual_site_with_vip, vk8s_service" must be set

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

### Choice Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Choice Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

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

### Crl Choice No Crl

Client certificate revocation status is not verified.

### Loadbalancer Type Tls Tcp

User is responsible for managing DNS to this load balancer..

###### One of the arguments from this list "tls_cert_params, tls_parameters" must be set

`tls_cert_params` - (Optional) Select/Add one or more TLS Certificate objects to associate with this Load Balancer. See [Tls Certificates Choice Tls Cert Params ](#tls-certificates-choice-tls-cert-params) below for details.

`tls_parameters` - (Optional) Upload a TLS certificate specifically for this Load Balancer (certificate must cover all Load Balancer domain names). See [Tls Certificates Choice Tls Parameters ](#tls-certificates-choice-tls-parameters) below for details.

### Loadbalancer Type Tls Tcp Auto Cert

or a DNS CNAME record should be created in your DNS provider's portal(only for Domains not managed by F5 Distributed Cloud)..

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Tcp Auto Cert Tls Config ](#tls-tcp-auto-cert-tls-config) below for details.

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

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Port Choice Use Default Port

Inherit the Load Balancer's Listen Port..

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

### Service Policy Choice Active Service Policies

Apply the specified list of service policies and bypass the namespace service policy set.

`policies` - (Required) If all policies are evaluated and none match, then the request will be denied by default.. See [ref](#ref) below for details.

### Tls Cert Params Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Tls Certificates Choice Tls Cert Params

Select/Add one or more TLS Certificate objects to associate with this Load Balancer.

`certificates` - (Required) Select one or more certificates with any domain names.. See [ref](#ref) below for details.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Cert Params Tls Config ](#tls-cert-params-tls-config) below for details.

### Tls Certificates Choice Tls Parameters

Upload a TLS certificate specifically for this Load Balancer (certificate must cover all Load Balancer domain names).

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

### Tls Tcp Auto Cert Tls Config

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

-	`id` - This is the id of the configured tcp_loadbalancer.
-	`cname` - This is the hostname of the configured tcp_loadbalancer.
