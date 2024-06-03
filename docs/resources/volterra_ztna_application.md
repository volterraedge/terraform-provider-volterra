---

page_title: "Volterra: ztna_application"

description: "The ztna_application allows CRUD of Ztna Application resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_ztna_application
==================================

The Ztna Application allows CRUD of Ztna Application resource on Volterra SaaS

~> **Note:** Please refer to [Ztna Application API docs](https://docs.cloud.f5.com/docs/api/ztna-application) to learn more

Example Usage
-------------

```hcl
resource "volterra_ztna_application" "example" {
  name           = "acmecorp-web"
  namespace      = "staging"
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

`domain_name` - (Optional) ztna_application's Domain name (`String`).

`msg` - (Optional) Denotes the service IP address and port. See [Msg ](#msg) below for details.

`policies` - (Optional) Denotes the ZTNA policies. See [Policies ](#policies) below for details.

`proxy_advertisement` - (Optional) Proxy Advertisement choice. See [Proxy Advertisement ](#proxy-advertisement) below for details.

`transport_type` - (Required) DNS Proxy supports TCP and UDP transport (`String`).

### Msg

Denotes the service IP address and port.

`ip` - (Optional) Used to set the Service IP. See [Msg Ip ](#msg-ip) below for details.

`port` - (Optional) Matches port (`Int`).

### Policies

Denotes the ZTNA policies.

`connectivity_policy_name` - (Optional) Used to set the Connectivity policy name (`String`).

`message_policy_name` - (Optional) Used to set the Message policy name (`String`).

`session_policy_name` - (Optional) Used to set the session policy name (`String`).

### Proxy Advertisement

Proxy Advertisement choice.

###### One of the arguments from this list "advertise_on_public_default_vip, advertise_on_public, advertise_custom, do_not_advertise" must be set

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

###### One of the arguments from this list "cloud_edge_segment, segment, site, virtual_site, vk8s_service, virtual_network, site_segment, virtual_site_segment" must be set

`site` - (Optional) Advertise on a customer site and a given network.. See [Choice Site ](#choice-site) below for details.

`site_segment` - (Optional) Advertise on a segment on a site. See [Choice Site Segment ](#choice-site-segment) below for details.

`virtual_network` - (Optional) Advertise on a virtual network. See [Choice Virtual Network ](#choice-virtual-network) below for details.

`virtual_site` - (Optional) Advertise on a customer virtual site and a given network.. See [Choice Virtual Site ](#choice-virtual-site) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Choice Vk8s Service ](#choice-vk8s-service) below for details.

###### One of the arguments from this list "use_default_port, port" must be set

`port` - (Optional) TCP port to Listen. (`Int`).

`use_default_port` - (Optional) For HTTP, default is 80. For HTTPS/SNI, default is 443. (`Bool`).

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

### Choice Vk8s Service

Advertise on vK8s Service Network on RE..

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Msg Ip

Used to set the Service IP.

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Port Choice Use Default Port

For HTTP, default is 80. For HTTPS/SNI, default is 443..

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### V6 Vip Choice Default V6 Vip

Use the default VIP, system allocated or configured in the virtual network.

### Ver Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ver Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Vip Choice Default Vip

Use the default VIP, system allocated or configured in the virtual network.

Attribute Reference
-------------------

-	`id` - This is the id of the configured ztna_application.
