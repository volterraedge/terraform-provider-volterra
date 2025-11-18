---

page_title: "Volterra: external_connector"

description: "The external_connector allows CRUD of External Connector resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_external_connector
====================================

The External Connector allows CRUD of External Connector resource on Volterra SaaS

~> **Note:** Please refer to [External Connector API docs](https://docs.cloud.f5.com/docs-v2/api/views-external-connector) to learn more

Example Usage
-------------

```hcl
resource "volterra_external_connector" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  ce_site_reference {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  // One of the arguments from this list "direct_connection gre ipsec" must be set

  gre {
    gre_parameters {
      peer_ip_address {
        addr = "192.168.1.1"
      }

      tunnel_eps {
        interface = "interface"

        local_tunnel_ip = "10.10.15.1/24"

        node = "node"

        remote_tunnel_ip = "10.10.15.2/24"
      }

      tunnel_mtu = "tunnel_mtu"

      // One of the arguments from this list "segment site_local_inside_network site_local_network" can be set

      site_local_network = true
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

`ce_site_reference` - (Required) Choose a CE site that will terminate the tunnel. See [ref](#ref) below for details.

###### One of the arguments from this list "direct_connection, gre, ipsec" must be set

`direct_connection` - (Optional) External Connector with direct connection. See [Connection Type Direct Connection ](#connection-type-direct-connection) below for details.(Deprecated)

`gre` - (Optional) External Connector with GRE tunnel. See [Connection Type Gre ](#connection-type-gre) below for details.(Deprecated)

`ipsec` - (Optional) External Connector with IPSec tunnel. See [Connection Type Ipsec ](#connection-type-ipsec) below for details.

### Connection Type Direct Connection

External Connector with direct connection.

### Connection Type Gre

External Connector with GRE tunnel.

`gre_parameters` - (Optional) GRE configuration parameters required for GRE Connection type. See [Gre Gre Parameters ](#gre-gre-parameters) below for details.

### Connection Type Ipsec

External Connector with IPSec tunnel.

`ike_parameters` - (Required) This section involves choosing the IKE parameters including the Phase 1 & Phase 2 Policies as well as other parameters.. See [Ipsec Ike Parameters ](#ipsec-ike-parameters) below for details.

`ipsec_tunnel_parameters` - (Required) In this section, we will configure the tunnel parameters, source, destination, IP addresses, and segment.. See [Ipsec Ipsec Tunnel Parameters ](#ipsec-ipsec-tunnel-parameters) below for details.

### Dpd Choice Dpd Disabled

Disabled Dead Peer Detection .

### Dpd Choice Dpd Keep Alive Timer

Enable Dead Peer Detection (DPD) by setting the keepalive timer..

`timeout` - (Optional) x-inlineHint: "The range is between 1 and 5 seconds." (`Int`).

### Gre Gre Parameters

GRE configuration parameters required for GRE Connection type.

`peer_ip_address` - (Optional) Configure tunnel remote endpoint's IP address. See [Gre Parameters Peer Ip Address ](#gre-parameters-peer-ip-address) below for details.

`tunnel_eps` - (Required) Configure tunnel parameters, source, destination, IP addresses.. See [Gre Parameters Tunnel Eps ](#gre-parameters-tunnel-eps) below for details.

`tunnel_mtu` - (Optional) Configure MTU for the GRE tunnel interface. (`Int`).

###### One of the arguments from this list "segment, site_local_inside_network, site_local_network" can be set

`segment` - (Optional) Interface belongs to user configured inside network segment. See [Tunnel Source Vn Segment ](#tunnel-source-vn-segment) below for details.

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

### Gre Parameters Peer Ip Address

Configure tunnel remote endpoint's IP address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Gre Parameters Tunnel Eps

Configure tunnel parameters, source, destination, IP addresses..

`interface` - (Required) For the chosen node, specify the interface that will be the tunnel source. (`String`).

`local_tunnel_ip` - (Required) For a particular tunnel on a node, specify the local tunnel IP Address i.e. the IP address of the tunnel on the CE node itself and a subnet prefix length (`String`).

`node` - (Required) A CE site is composed of multiple nodes. Choose a node that will be part of this external connection. (`String`).

`remote_tunnel_ip` - (Required) For a particular tunnel on a node, specify the remote tunnel IP Address i.e. the IP address of the tunnel on the remote gateway and a subnet prefix length (`String`).

### Ipsec Ike Parameters

This section involves choosing the IKE parameters including the Phase 1 & Phase 2 Policies as well as other parameters..

###### One of the arguments from this list "dpd_disabled, dpd_keep_alive_timer" can be set

`dpd_disabled` - (Optional) Disabled Dead Peer Detection (`Bool`).

`dpd_keep_alive_timer` - (Optional) Enable Dead Peer Detection (DPD) by setting the keepalive timer.. See [Dpd Choice Dpd Keep Alive Timer ](#dpd-choice-dpd-keep-alive-timer) below for details.

`ike_phase1_profile` - (Required) IKE Phase 1 profile defines mainly the encryption and authentication algorithms to be used for IKE SA. See [ref](#ref) below for details.

`ike_phase2_profile` - (Required) IKE Phase 2 profile defines mainly the encryption and authentication algorithms to be used for ESP SA. See [ref](#ref) below for details.

###### One of the arguments from this list "lc_hostname, lc_ip_address, use_default_local_ike_id" can be set

`lc_hostname` - (Optional) Configure an hostname Local IKE ID (`String`).(Deprecated)

`lc_ip_address` - (Optional) Configure an IP address as Local IKE ID. See [Local Ike Id Lc Ip Address ](#local-ike-id-lc-ip-address) below for details.(Deprecated)

`use_default_local_ike_id` - (Optional) Use default local IKE ID (IP address of tunnel source) (`Bool`).

###### One of the arguments from this list "initiator, responder" must be set

`initiator` - (Optional) Configure as Tunnel Initiator (`Bool`).

`responder` - (Optional) Configure as Tunnel Responder (`Bool`).

###### One of the arguments from this list "rm_hostname, rm_ip_address, use_default_remote_ike_id" can be set

`rm_hostname` - (Optional) Configure an hostname Remote IKE ID (`String`).

`rm_ip_address` - (Optional) Configure an IP address as Remote IKE ID. See [Remote Ike Id Rm Ip Address ](#remote-ike-id-rm-ip-address) below for details.

`use_default_remote_ike_id` - (Optional) Use default remote IKE ID (IP address of tunnel gateway) (`Bool`).

### Ipsec Ipsec Tunnel Parameters

In this section, we will configure the tunnel parameters, source, destination, IP addresses, and segment..

`peer_ip_address` - (Required) This is the reachable address for the remote gateway. For instance if the remote gateway is reachable over the public internet with a public address, then the remote Gateway IP Address would be this public address.. See [Ipsec Tunnel Parameters Peer Ip Address ](#ipsec-tunnel-parameters-peer-ip-address) below for details.

`psk` - (Required) The IKE pre-shared key (PSK) is required to ensure the IKE peers can authenticate one another within IKE phase 1 negotiation. (`String`).

`tunnel_eps` - (Required) Configure tunnel parameters, local and remote IP addresses . See [Ipsec Tunnel Parameters Tunnel Eps ](#ipsec-tunnel-parameters-tunnel-eps) below for details.

`tunnel_mtu` - (Required) The tunnel MTU defines the maximum size of the packet that can be sent through the tunnel without needing to be fragmented (`Int`).

###### One of the arguments from this list "segment, site_local_inside_network, site_local_network" must be set

`segment` - (Optional) Interface belongs to user configured network segment. See [Tunnel Source Vn Segment ](#tunnel-source-vn-segment) below for details.

`site_local_inside_network` - (Optional) Interface belongs to site local network inside (`Bool`).

`site_local_network` - (Optional) Interface belongs to site local network (outside) (`Bool`).

### Ipsec Tunnel Parameters Peer Ip Address

This is the reachable address for the remote gateway. For instance if the remote gateway is reachable over the public internet with a public address, then the remote Gateway IP Address would be this public address..

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ipsec Tunnel Parameters Tunnel Eps

Configure tunnel parameters, local and remote IP addresses .

`interface` - (Required) For the chosen node, specify the interface that will be the tunnel source. (`String`).

`local_tunnel_ip` - (Required) For a particular tunnel on a node, specify the local tunnel IP Address i.e. the IP address of the tunnel on the CE node itself and a subnet prefix length (`String`).

`node` - (Required) A CE site is composed of multiple nodes. Choose a node that will be part of this external connection. (`String`).

`remote_tunnel_ip` - (Required) For a particular tunnel on a node, specify the remote tunnel IP Address i.e. the IP address of the tunnel on the remote gateway and a subnet prefix length (`String`).

### Local Ike Id Lc Ip Address

Configure an IP address as Local IKE ID.

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Local Ike Id Use Default Local Ike Id

Use default local IKE ID (IP address of tunnel source).

### Mode Choice Initiator

Configure as Tunnel Initiator.

### Mode Choice Responder

Configure as Tunnel Responder.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Remote Ike Id Rm Ip Address

Configure an IP address as Remote IKE ID.

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Remote Ike Id Use Default Remote Ike Id

Use default remote IKE ID (IP address of tunnel gateway).

### Tunnel Source Vn Segment

Interface belongs to user configured inside network segment.

`refs` - (Required) Reference to Segment Object. See [ref](#ref) below for details.

`virtual_networks` - (Optional) Internally used to resolve segment ref to networks.. See [ref](#ref) below for details.(Deprecated)

### Tunnel Source Vn Site Local Inside Network

Interface belongs to site local network inside.

### Tunnel Source Vn Site Local Network

Interface belongs to site local network (outside).

### Ver Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ver Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured external_connector.
