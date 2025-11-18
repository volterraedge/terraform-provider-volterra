---

page_title: "Volterra: dns_proxy"

description: "The dns_proxy allows CRUD of Dns Proxy resource on Volterra SaaS"
-------------------------------------------------------------------------------

Resource volterra_dns_proxy
===========================

The Dns Proxy allows CRUD of Dns Proxy resource on Volterra SaaS

~> **Note:** Please refer to [Dns Proxy API docs](https://docs.cloud.f5.com/docs-v2/api/dns-proxy) to learn more

Example Usage
-------------

```hcl
resource "volterra_dns_proxy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  cache_profile {
    // One of the arguments from this list "cache_size disable_cache_profile" must be set

    cache_size = "1000"
  }

  ddos_profile {
    // One of the arguments from this list "disable_ddos_mitigation enable_ddos_mitigation" can be set

    enable_ddos_mitigation = true
  }

  origin_servers {
    health_checks {
      health_check {
        // One of the arguments from this list "dns_health_check icmp_health_check tcp_health_check udp_health_check" can be set

        dns_health_check {
          expected_rcode = "no-error"

          expected_record_type = "REQUESTED_QUERY_TYPE"

          expected_response = "10.0.0.1"

          query_name = "www.example.com"

          query_type = "A"

          reverse = true
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
      // One of the arguments from this list "k8s_service public_ip public_name" must be set

      public_name {
        dns_name = "value"

        refresh_interval = "20"
      }

      // One of the arguments from this list "no_preference site_preferences" can be set

      no_preference = true
    }
  }

  transport_type = ["transport_type"]
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

`cache_profile` - (Required) which caches DNS replies from the origin DNS servers.. See [Cache Profile ](#cache-profile) below for details.

`ddos_profile` - (Required) to protect the origin DNS servers from external DDoS attacks.. See [Ddos Profile ](#ddos-profile) below for details.

`irules` - (Optional) Options for attaching iRules to dns proxy. See [ref](#ref) below for details.

`lb_algorithm` - (Optional) The algorithm used for load balance request between origin servers. See [Lb Algorithm ](#lb-algorithm) below for details.

`origin_servers` - (Required) Origin Servers for DNS Proxy . See [Origin Servers ](#origin-servers) below for details.

`protocol_inspection` - (Optional) Options for enabling and configuring protocol inspection configuration. See [ref](#ref) below for details.

`proxy_advertisement` - (Optional) Proxy Advertisement choice. See [Proxy Advertisement ](#proxy-advertisement) below for details.

`transport_type` - (Required) DNS Proxy supports TCP and UDP transport (`String`).

### Cache Profile

which caches DNS replies from the origin DNS servers..

###### One of the arguments from this list "cache_size, disable_cache_profile" must be set

`cache_size` - (Optional) cache size (`Int`).

`disable_cache_profile` - (Optional) x-displayName: "Disable" (`Bool`).

### Ddos Profile

to protect the origin DNS servers from external DDoS attacks..

###### One of the arguments from this list "disable_ddos_mitigation, enable_ddos_mitigation" can be set

`disable_ddos_mitigation` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_ddos_mitigation` - (Optional) x-displayName: "Enable" (`Bool`).

### Lb Algorithm

The algorithm used for load balance request between origin servers.

###### One of the arguments from this list "round_robin" must be set

`round_robin` - (Optional) Requests are sent to all the eligible origin servers in round-robin fashion (`Bool`).

### Origin Servers

Origin Servers for DNS Proxy .

`health_checks` - (Required) x-required. See [Origin Servers Health Checks ](#origin-servers-health-checks) below for details.

`origin_servers` - (Required) List of origin servers for Proxy. See [Origin Servers Origin Servers ](#origin-servers-origin-servers) below for details.

### Proxy Advertisement

Proxy Advertisement choice.

###### One of the arguments from this list "advertise_custom, advertise_on_public, advertise_on_public_default_vip, do_not_advertise" must be set

`advertise_custom` - (Optional) Advertise this load balancer on specific sites. See [Advertise Choice Advertise Custom ](#advertise-choice-advertise-custom) below for details.

`advertise_on_public` - (Optional) Advertise this proxy on public network. See [Advertise Choice Advertise On Public ](#advertise-choice-advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Advertise this load balancer on public network with default VIP (`Bool`).

`do_not_advertise` - (Optional) Do not advertise this proxy (`Bool`).

### Advertise Choice Advertise Custom

Advertise this load balancer on specific sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Advertise Custom Advertise Where ](#advertise-custom-advertise-where) below for details.

### Advertise Choice Advertise On Public

Advertise this proxy on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Advertise Choice Advertise On Public Default Vip

Advertise this load balancer on public network with default VIP.

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

### Cache Profile Choice Disable Cache Profile

x-displayName: "Disable".

### Choice Advertise On Public

Advertise this load balancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Choice Dns Health Check

DNS Health Check.

`expected_rcode` - (Required) Specifies an expected Rcode in the answer section of DNS Response, option [no-error, any](`String`).

`expected_record_type` - (Required) options: [REQUESTED_QUERY_TYPE, RECORD_TYPE_ANY] when REQUESTED_QUERY_TYPE is set, health monitor expects record type same as requested query type (`String`).

`expected_response` - (Required) Specifies an IPv4 or IPv6 address in the answer section of DNS Response (`String`).

`query_name` - (Required) The query name that the monitor sends a DNS query for. (`String`).

`query_type` - (Required) The DNS query type that the monitor sends. Supported types are: [a, aaaa] default: a (`String`).

`reverse` - (Optional) string match marks the monitored object down instead of up. (`Bool`).

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

### Choice Public Ip

Specify origin server with public IP.

###### One of the arguments from this list "ip, ipv6" must be set

`ip` - (Optional) Public IPV4 address (`String`).

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

### Choice Udp Health Check

UDP Health Check.

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

### Ddos Mitigation Choice Disable Ddos Mitigation

x-displayName: "Disable".

### Ddos Mitigation Choice Enable Ddos Mitigation

x-displayName: "Enable".

### Health Checks Health Check

List of Health Checks.

###### One of the arguments from this list "dns_health_check, icmp_health_check, tcp_health_check, udp_health_check" can be set

`dns_health_check` - (Optional) DNS Health Check. See [Choice Dns Health Check ](#choice-dns-health-check) below for details.

`icmp_health_check` - (Optional) ICMP Health Check (`Bool`).

`tcp_health_check` - (Optional) TCP Health Check. See [Choice Tcp Health Check ](#choice-tcp-health-check) below for details.

`udp_health_check` - (Optional) UDP Health Check. See [Choice Udp Health Check ](#choice-udp-health-check) below for details.(Deprecated)

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

### Network Choice Inside Network

Inside network on the site.

### Network Choice Outside Network

Outside network on the site.

### Network Choice Vk8s Networks

origin server are on vK8s network on the site.

### Origin Servers Health Checks

x-required.

`health_check` - (Required) List of Health Checks. See [Health Checks Health Check ](#health-checks-health-check) below for details.

`healthy_threshold` - (Required) required to mark a host healthy. (`Int`).

`interval` - (Required) Time interval in seconds between two healthcheck requests. (`Int`).

`jitter` - (Optional) increase the wait time between two healthcheck requests. (`Int`).(Deprecated)

`jitter_percent` - (Optional) Add a random amount of time as a percent value to the interval between successive health check requests. (`Int`).(Deprecated)

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

`unhealthy_threshold` - (Required) this threshold is ignored and the host is considered unhealthy immediately. (`Int`).

### Origin Servers Origin Servers

List of origin servers for Proxy.

###### One of the arguments from this list "k8s_service, public_ip, public_name" must be set

`k8s_service` - (Optional) Specify origin server with K8s service name and site information. See [Choice K8s Service ](#choice-k8s-service) below for details.

`public_ip` - (Optional) Specify origin server with public IP. See [Choice Public Ip ](#choice-public-ip) below for details.

`public_name` - (Optional) Specify origin server with public DNS name. See [Choice Public Name ](#choice-public-name) below for details.

###### One of the arguments from this list "no_preference, site_preferences" can be set

`no_preference` - (Optional) Use all the discovered origins (`Bool`).

`site_preferences` - (Optional) discovered origins. See [Proximity Choice Site Preferences ](#proximity-choice-site-preferences) below for details.

### Port Choice Use Default Port

Inherit the Load Balancer's Listen Port..

### Proximity Choice No Preference

Use all the discovered origins.

### Proximity Choice Site Preferences

discovered origins.

`refs` - (Optional) Reference to one or more sites. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Service Info Service Selector

discovery has to happen. This implicit label is added to service_selector.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Snat Pool Choice No Snat Pool

No configured SNAT Pool to reach Origin Server.

### Snat Pool Choice Snat Pool

Configure SNAT Pool to reach Origin Server.

`ipv6_prefixes` - (Optional) List of IPv6 prefix strings. (`String`).

`prefixes` - (Optional) List of IPv4 prefixes that represent an endpoint (`String`).

### V6 Vip Choice Default V6 Vip

Use the default VIP, system allocated or configured in the virtual network.

### Vip Choice Default Vip

Use the default VIP, system allocated or configured in the virtual network.

Attribute Reference
-------------------

-	`id` - This is the id of the configured dns_proxy.
