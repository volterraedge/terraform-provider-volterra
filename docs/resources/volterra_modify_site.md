---

page_title: "Volterra: volterra_modify_site"

description: "The volterra_modify_site helps update site fields Volterra SaaS"

---

Resource volterra_modify_site
=============================

================================

volterra_modify_site helps to update fields of site object. This is not applicable for site created through view style APIs.

~> **Note:** Please refer to [Site api docs](https://docs.cloud.f5.com/docs/api/site) to learn more

Example Usage
-------------

---

```hcl
resource "volterra_modify_site" "example" {
  name      = "site1"
  labels    = {
    "ves.io/fleet" = "fleet-car"
  }
}

```

Argument Reference
------------------

---

`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).

`description` - (Optional) Human readable description for the object (`String`).

`disable` - (Optional) A value of true will administratively disable the object (`Bool`).

`labels` - (Optional) by selector expression (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`namespace` - (Required) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

`address` - (Optional) Site's geographical address that can be used determine its latitude and longitude. Example: "123 Street, city, country, postal code" (`String`).

`bgp_peer_address` - (Optional) bgp peer address that can be used as parameter for BGP configuration when BGP is configured to fetch BGP peer address from site Object. This can be used to change peer address per site in fleet (`String`).

`bgp_router_id` - (Optional) bgp router id that can be used as parameter for BGP configuration when BGP is configured to fetch BGP router ID from site object. This can be used to change router id per site in a fleet (`String`).

`coordinates` - (Optional) Coordinates of the site which provides the site physical location `latitude` - (Optional) Latitude of the site location (`String`). `longitude` - (Optional) Longitude of site location (`String`).

`desired_pool_count` - (Optional) Desire pool count represent number of nodes in scaling group for manual scaling. It is valid only for k8s worker nodes not masters. The desired count must be less than or equal to the maximum size of the group. If new value for Desired is greater than Max, then Max must be updated in cloud provider configuration (`Int`).

`inside_nameserver` - (Optional) DNS server IP to be used for name resolution in inside network (`String`).

`inside_vip` - (Optional) Virtual IP to be used as automatic VIP for site local inside network. See documentation for "VIP" in advertise policy to see when Inside VIP is used. When configured, this is used as VIP (depending on advertise policy configuration). When not configured, site local inside interface ip will be used as VIP (`String`).

`operating_system_version` - (Optional) Desired Operating System version for this site (`String`).

`volterra_software_version` - (Optional) Desired Volterra software version for this site, a string matching released set of software components (`String`).

`outside_nameserver` - (Optional) DNS server IP to be used for name resolution in outside network (`String`).

`outside_vip` - (Optional) Virtual IP to be used as automatic VIP for site local outside network. See documentation for "VIP" in advertise policy to see when Outside VIP is used. When configured, this is used as VIP (depending on advertise policy configuration). When not configured, site local interface ip will be used as VIP (`String`).

`site_to_site_network_type` - (Optional) Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL provides connectivity to public (outside) network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created automatically and present on all sites Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE is a private network inside site. It is a secure network and is not connected to public network. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on CE sites. This network is created during provisioning of site User defined per-site virtual network. Scope of this virtual network is limited to the site. This is not yet supported Virtual-network of type VIRTUAL_NETWORK_PUBLIC directly conects to the public internet. Virtual-network of this type is local to every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type is supported on RE sites only It is an internally created by the system. They must not be created by user Virtual Neworks with global scope across different sites in Volterra domain. An example global virtual-network called "AIN Network" is created for every tenant. for volterra fabric

Constraints: It is currently only supported as internally created by the system. Vk8s service network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant Constraints: It is an internally created by the system. Must not be created by user VER internal network for the site. It can only be used for virtual hosts with SMA_PROXY type proxy Constraints: It is an internally created by the system. Must not be created by user Virtual-network of type VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE represents both VIRTUAL_NETWORK_SITE_LOCAL and VIRTUAL_NETWORK_SITE_LOCAL_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of type VIRTUAL_NETWORK_IP_AUTO is selected for an endpoint, VER will try to determine the network based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

Valid values are: `VIRTUAL_NETWORK_SITE_LOCAL` `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE` `VIRTUAL_NETWORK_PER_SITE` `VIRTUAL_NETWORK_PUBLIC` `VIRTUAL_NETWORK_GLOBAL` `VIRTUAL_NETWORK_SITE_SERVICE` `VIRTUAL_NETWORK_VER_INTERNAL` `VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE` `VIRTUAL_NETWORK_IP_AUTO` (`String`).

`site_to_site_tunnel_ip` - (Optional) VIP in the site_to_site_network_type configured above used for terminating IPSec/SSL tunnels created with SiteMeshGroup (`String`).

`tunnel_dead_timeout` - (Optional) Time interval, in millisec, within which any ipsec / ssl connection from the site going down is detected. When not set (== 0), a default value of 10000 msec will be used (`Int`).

`tunnel_type` - (Optional) Tunnel encapsulation to be used between sites. Valid values are `SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL` `SITE_TO_SITE_TUNNEL_IPSEC` `SITE_TO_SITE_TUNNEL_SSL` `SITE_TO_SITE_CLEAR` (`String`).

`vip_vrrp_mode` - (Optional) VRRP advertisement mode for VIP. Valid values are `VIP_VRRP_INVALID` `VIP_VRRP_ENABLE` `VIP_VRRP_DISABLE` (`String`).

`volterra_software_overide` - (Optional) Decide which software version takes effect in case of conflict between site and fleet. Valid values are `"SITE_SOFTWARE_OVERRIDE_SITE` `SITE_SOFTWARE_OVERRIDE_NEWER` `SITE_SOFTWARE_OVERRIDE_FLEET` (`String`).

`retry` - (Optional) Number of retries to get the volterra site config, before it fails (`Int`).

`wait_time` - (Optional) Number of seconds to wait before it retries (`Int`).

Attribute Reference
-------------------

---

-	`id` - This is the id of Volterra site object.
