---

page_title: "Volterra: tcp_loadbalancer"

description: "The tcp_loadbalancer allows CRUD of Tcp Loadbalancer resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_tcp_loadbalancer
==================================

The Tcp Loadbalancer allows CRUD of Tcp Loadbalancer resource on Volterra SaaS

~> **Note:** Please refer to [Tcp Loadbalancer API docs](https://volterra.io/docs/api/views-tcp-loadbalancer) to learn more

Example Usage
-------------

```hcl
resource "volterra_tcp_loadbalancer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "do_not_advertise advertise_on_public_default_vip advertise_on_public advertise_custom" must be set
  advertise_on_public_default_vip = true

  // One of the arguments from this list "retract_cluster do_not_retract_cluster" must be set
  retract_cluster = true

  // One of the arguments from this list "hash_policy_choice_least_active hash_policy_choice_random hash_policy_choice_source_ip_stickiness hash_policy_choice_round_robin" must be set
  hash_policy_choice_round_robin = true
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

`advertise_custom` - (Optional) Advertise this VIP on specific sites. See [Advertise Custom ](#advertise-custom) below for details.

`advertise_on_public` - (Optional) Advertise this loadbalancer on public network. See [Advertise On Public ](#advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Advertise this loadbalancer on public network with default VIP (bool).

`do_not_advertise` - (Optional) Do not advertise this loadbalancer (bool).

`do_not_retract_cluster` - (Optional) configuration. (bool).

`retract_cluster` - (Optional) for route (bool).

`dns_volterra_managed` - (Optional) This requires the domain to be delegated to Volterra using the Delegated Domain feature. (`Bool`).

`domains` - (Optional) Domains also indicate the list of names for which DNS resolution will be done by VER (`List of String`).

`hash_policy_choice_least_active` - (Optional) Connections are sent to origin server that has least active connections (bool).

`hash_policy_choice_random` - (Optional) Connections are sent to all eligible origin servers in random fashion (bool).

`hash_policy_choice_round_robin` - (Optional) Connections are sent to all eligible origin servers in round robin fashion (bool).

`hash_policy_choice_source_ip_stickiness` - (Optional) Connections are sent to all eligible origin servers using hash of source ip. Consistent hashing algorithm, ring hash, is used to select origin server (bool).

`idle_timeout` - (Optional) The amount of time that a stream can exist without upstream or downstream activity, in milliseconds. (`Int`).

`listen_port` - (Optional) Listen Port for this TCP proxy (`Int`).

`origin_pools_weights` - (Optional) Origin pools and weights used for this loadbalancer.. See [Origin Pools Weights ](#origin-pools-weights) below for details.

`with_sni` - (Optional) Set to true to enable TCP loadbalancer with SNI (`Bool`).

### Advertise Custom

Advertise this VIP on specific sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Advertise Where ](#advertise-where) below for details.

### Advertise On Public

Advertise this loadbalancer on public network.

`public_ip` - (Required) Dedicated public ip are allocated by volterra on request. See [ref](#ref) below for details.

### Advertise Where

Where should this load balancer be available.

`site` - (Optional) Advertise on a customer site and a given network. . See [Site ](#site) below for details.

`virtual_network` - (Optional) Advertise on a virtual network. See [Virtual Network ](#virtual-network) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Virtual Site ](#virtual-site) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Vk8s Service ](#vk8s-service) below for details.

`port` - (Optional) TCP port to Listen. (`Int`).

`use_default_port` - (Optional) For HTTP, default is 80. For HTTPS/SNI, default is 443. (bool).

### Default Vip

Use the default VIP, system allocated or configured in the virtual network.

### Origin Pools Weights

Origin pools and weights used for this loadbalancer..

`endpoint_subsets` - (Optional) upstream origin pool which match this metadata will be selected for load balancing (`String`).

`cluster` - (Required) More flexible, advanced feature control with cluster. See [ref](#ref) below for details.

`pool` - (Required) Simple, commonly used pool parameters with origin pool. See [ref](#ref) below for details.

`weight` - (Optional) Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool (`Int`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site

Advertise on a customer site and a given network. .

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Use Default Port

For HTTP, default is 80. For HTTPS/SNI, default is 443..

### Virtual Network

Advertise on a virtual network.

`default_vip` - (Optional) Use the default VIP, system allocated or configured in the virtual network (bool).

`specific_vip` - (Optional) Use given IP address as VIP on VoltADN private Network (`String`).

`virtual_network` - (Required) Select virtual network reference. See [ref](#ref) below for details.

### Virtual Site

Advertise on a customer virtual site and a given network..

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Vk8s Service

Advertise on vK8s Service Network on RE..

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured tcp_loadbalancer.
