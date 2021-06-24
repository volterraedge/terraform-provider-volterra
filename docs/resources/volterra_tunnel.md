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
    // One of the arguments from this list "intf ip_address" must be set

    ip_address {
      // One of the arguments from this list "ip_address auto" must be set

      ip_address {
        // One of the arguments from this list "ipv4 ipv6" must be set

        ipv4 {
          addr = "192.168.1.1"
        }
      }

      virtual_network_type {
        // One of the arguments from this list "public site_local site_local_inside" must be set
        site_local_inside = true
      }
    }
  }

  remote_ip {
    // One of the arguments from this list "endpoints ip" must be set

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

### Auto

No IP configured, system picks up IP from interface attached to virtual network.

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by Volterra Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted .

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Endpoints

Key is ver node name and value is IP address.

`endpoints` - (Optional) Key is ver node name and value is Remote node attributes. See [Endpoints ](#endpoints) below for details.

### Intf

Source IP and network is picked from the interface referred here.

`local_intf` - (Optional) Local interface to be used for filling in source information of IP and network for transport. See [ref](#ref) below for details.

### Ip

Remote IP to which tunnel will be established.

`ipv4` - (Optional) IPv4 Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ipv6 ](#ipv6) below for details.

### Ip Address

Local Source IP configuration, provides IP address with virtual network to be used for transport.

`auto` - (Optional) No IP configured, system picks up IP from interface attached to virtual network (bool).

`ip_address` - (Required) Local Source IP to be used for tunnel. See [Ip Address ](#ip-address) below for details.

`virtual_network_type` - (Required) Local Virtual network to be used for transporting encapsulated packets. See [Virtual Network Type ](#virtual-network-type) below for details.

### Ipsec

Configuration for IPSec encapsulation.

`ipsec_psk` - (Optional) Pre shared key, valid for tunnel type IPSEC_PSK, SA are computed by IKE dynamically. See [Ipsec Psk ](#ipsec-psk) below for details.

### Ipsec Psk

Pre shared key, valid for tunnel type IPSEC_PSK, SA are computed by IKE dynamically.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Local Ip

Selects local IP address configuration for tunnel.

`intf` - (Optional) Source IP and network is picked from the interface referred here. See [Intf ](#intf) below for details.

`ip_address` - (Optional) Local Source IP configuration, provides IP address with virtual network to be used for transport. See [Ip Address ](#ip-address) below for details.

### Params

Configuration for supported tunnel types.

`ipsec` - (Optional) Configuration for IPSec encapsulation. See [Ipsec ](#ipsec) below for details.

### Public

Indicates use of public network.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Remote Ip

Selects remote endpoint IP address configuration for tunnel.

`endpoints` - (Optional) Key is ver node name and value is IP address. See [Endpoints ](#endpoints) below for details.

`ip` - (Optional) Remote IP to which tunnel will be established. See [Ip ](#ip) below for details.

### Site Local

Indicates use of site local network.

### Site Local Inside

Indicates use of site local inside network.

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Virtual Network Type

Local Virtual network to be used for transporting encapsulated packets.

`public` - (Optional) Indicates use of public network (bool).

`site_local` - (Optional) Indicates use of site local network (bool).

`site_local_inside` - (Optional) Indicates use of site local inside network (bool).

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured tunnel.
