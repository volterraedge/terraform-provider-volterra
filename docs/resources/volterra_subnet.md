---

page_title: "Volterra: subnet"

description: "The subnet allows CRUD of Subnet resource on Volterra SaaS"
-------------------------------------------------------------------------

Resource volterra_subnet
========================

The Subnet allows CRUD of Subnet resource on Volterra SaaS

~> **Note:** Please refer to [Subnet API docs](https://docs.cloud.f5.com/docs-v2/api/subnet) to learn more

Example Usage
-------------

```hcl
resource "volterra_subnet" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  site_subnet_params {
    // One of the arguments from this list "dhcp static_ip" can be set

    dhcp = true

    site {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    subnet_dhcp_server_params {
      dhcp_networks {
        // One of the arguments from this list "network_prefix network_prefix_ipv6" can be set

        network_prefix = "10.1.1.0/24"
      }
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

###### One of the arguments from this list "connect_to_layer2, connect_to_sli, connect_to_slo, isolated_nw" can be set

`connect_to_layer2` - (Optional) x-displayName: "Connect to Layer2 Interface". See [Connection Choice Connect To Layer2 ](#connection-choice-connect-to-layer2) below for details.

`connect_to_sli` - (Optional) x-displayName: "Connect to Site Local Inside Network" (`Bool`).(Deprecated)

`connect_to_slo` - (Optional) x-displayName: "Connect to Site Local Outside Network" (`Bool`).

`isolated_nw` - (Optional) x-displayName: "Not Connected to Other Network" (`Bool`).

`site_subnet_params` - (Required) Configure subnet parameters per site . See [Site Subnet Params ](#site-subnet-params) below for details.

### Site Subnet Params

Configure subnet parameters per site .

###### One of the arguments from this list "dhcp, static_ip" can be set

`dhcp` - (Optional) Subnet uses DHCP for address allocation (`Bool`).

`static_ip` - (Optional) Subnet uses static IP addresses (`Bool`).

`site` - (Required) Site that the subnet parameters are for. See [ref](#ref) below for details.

`subnet_dhcp_server_params` - (Optional) DHCP parameters for a subnet on a site. See [Site Subnet Params Subnet Dhcp Server Params ](#site-subnet-params-subnet-dhcp-server-params) below for details.

### Address Choice Dhcp

Subnet uses DHCP for address allocation.

### Address Choice Static Ip

Subnet uses static IP addresses.

### Connection Choice Connect To Layer2

x-displayName: "Connect to Layer2 Interface".

`layer2_intf_internal_ref` - (Optional) Layer2 interface to connect to. See [ref](#ref) below for details.(Deprecated)

`layer2_intf_ref` - (Required) Layer2 interface to connect to. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site Subnet Params Subnet Dhcp Server Params

DHCP parameters for a subnet on a site.

`dhcp_networks` - (Optional) List of networks from which DHCP server can allocate IP addresses. See [Subnet Dhcp Server Params Dhcp Networks ](#subnet-dhcp-server-params-dhcp-networks) below for details.

### Subnet Dhcp Server Params Dhcp Networks

List of networks from which DHCP server can allocate IP addresses.

###### One of the arguments from this list "network_prefix, network_prefix_ipv6" can be set

`network_prefix` - (Optional) Network prefix for subnet (`String`).

`network_prefix_ipv6` - (Optional) Network prefix for IPV6 subnet (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured subnet.
