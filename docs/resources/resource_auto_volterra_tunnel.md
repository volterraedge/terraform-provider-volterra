---

page_title: "Volterra: tunnel"

description: "The tunnel allows CRUD of Tunnel resource on Volterra SaaS"
-------------------------------------------------------------------------

Resource volterra_tunnel
========================

The Tunnel allows CRUD of Tunnel resource on Volterra SaaS

~> **Note:** Please refer to [Tunnel API docs](https://volterra.io/docs/api/tunnel) to learn more

Example Usage
-------------

```hcl
resource "volterra_tunnel" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  local_ip {
    // One of the arguments from this list "ip_address intf" must be set

    intf {
      local_intf {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
    }
  }

  remote_ip {
    // One of the arguments from this list "ip endpoints" must be set

    ip {
      // One of the arguments from this list "ipv4 ipv6" must be set

      ipv4 {
        addr = "192.168.1.1"
      }
    }
  }

  tunnel_type = ["tunnel_type"]
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

`local_ip` - (Required) Selects local IP address configuration for tunnel. See [Local Ip ](#local-ip) below for details.

`params` - (Optional) Configuration for supported tunnel types. See [Params ](#params) below for details.

`remote_ip` - (Required) Selects remote endpoint IP address configuration for tunnel. See [Remote Ip ](#remote-ip) below for details.

`tunnel_type` - (Required) Tunnel type supported is IPSEC with pre-shared key (IPSEC_PSK) (`String`).

### Local Ip

Selects local IP address configuration for tunnel.

`intf` - (Optional) Source IP and network is picked from the interface referred here. See [Intf ](#intf) below for details.

`ip_address` - (Optional) Local Source IP configuration, provides IP address with virtual network to be used for transport. See [Ip Address ](#ip-address) below for details.

### Params

Configuration for supported tunnel types.

`ipsec` - (Optional) Configuration for IPSec encapsulation. See [Ipsec ](#ipsec) below for details.

### Remote Ip

Selects remote endpoint IP address configuration for tunnel.

`endpoints` - (Optional) Key is ver node name and value is IP address. See [Endpoints ](#endpoints) below for details.

`ip` - (Optional) Remote IP to which tunnel will be established. See [Ip ](#ip) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured tunnel.
