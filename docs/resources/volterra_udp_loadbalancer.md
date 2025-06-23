---

page_title: "Volterra: udp_loadbalancer"

description: "The udp_loadbalancer allows CRUD of Udp Loadbalancer resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_udp_loadbalancer
==================================

The Udp Loadbalancer allows CRUD of Udp Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Udp Loadbalancer API docs](https://docs.cloud.f5.com/docs-v2/api/views-udp-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_udp_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "advertise_custom advertise_on_public advertise_on_public_default_vip do_not_advertise" must be set

  advertise_on_public {
    public_ip {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
  }

  // One of the arguments from this list "hash_policy_choice_random hash_policy_choice_round_robin hash_policy_choice_source_ip_stickiness" must be set

  hash_policy_choice_source_ip_stickiness = true

  // One of the arguments from this list "udp" must be set

  udp = true

  // One of the arguments from this list "listen_port" must be set

  listen_port = "53"
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

`advertise_on_public` - (Optional) Advertise this load balancer on public network with a user specified public IP address. See [Advertise Choice Advertise On Public ](#advertise-choice-advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Advertise this load balancer on public network with default VIP (`Bool`).

`do_not_advertise` - (Optional) Do not advertise this load balancer (`Bool`).

`dns_volterra_managed` - (Optional) As a prerequisite, the domain to be delegated to F5 Distributed Cloud using the Delegated Domain feature or a DNS CNAME record must be created in your DNS provider's portal. (`Bool`).

`domains` - (Optional) A list of domains (host/authority header) that will be matched to this load balancer. (`List of String`).

`enable_per_packet_load_balancing` - (Optional) If enabled: Each packet is directed to an upstream server as the load balancing algorithm dictates. (`Bool`).

###### One of the arguments from this list "hash_policy_choice_random, hash_policy_choice_round_robin, hash_policy_choice_source_ip_stickiness" must be set

`hash_policy_choice_random` - (Optional) Connections are sent to all eligible origin servers in random fashion (`Bool`).

`hash_policy_choice_round_robin` - (Optional) Connections are sent to all eligible origin servers in round robin fashion (`Bool`).

`hash_policy_choice_source_ip_stickiness` - (Optional) Connections are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server (`Bool`).

`idle_timeout` - (Optional) The amount of time that a session can exist without upstream or downstream activity, in milliseconds. (`Int`).

###### One of the arguments from this list "udp" must be set

`udp` - (Optional) UDP Load Balancer. (`Bool`).

`origin_pools_weights` - (Optional) Origin pools with weights and priorities used for this load balancer.. See [Origin Pools Weights ](#origin-pools-weights) below for details.

###### One of the arguments from this list "listen_port" must be set

`listen_port` - (Optional) Listen Port for this load balancer (`Int`).

### Origin Pools Weights

Origin pools with weights and priorities used for this load balancer..

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

Advertise this load balancer on public network with a user specified public IP address.

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

### Port Choice Use Default Port

Inherit the Load Balancer's Listen Port..

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### V6 Vip Choice Default V6 Vip

Use the default VIP, system allocated or configured in the virtual network.

### Vip Choice Default Vip

Use the default VIP, system allocated or configured in the virtual network.

Attribute Reference
-------------------

-	`id` - This is the id of the configured udp_loadbalancer.
