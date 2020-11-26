---

page_title: "Volterra: protocol_policer"

description: "The protocol_policer allows CRUD of Protocol Policer resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_protocol_policer
==================================

The Protocol Policer allows CRUD of Protocol Policer resource on Volterra SaaS

~> **Note:** Please refer to [Protocol Policer API docs](https://volterra.io/docs/api/protocol-policer) to learn more

Example Usage
-------------

```hcl
resource "volterra_protocol_policer" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

`protocol_policer` - (Optional) List of L4 protocol match condition and associated traffic rate limits. See [Protocol Policer ](#protocol-policer) below for details.

### Dns

Match all DNS packets.

### Icmp

ICMP message types to be matched in packet.

`type` - (Optional)type (`List of Strings`).

### Protocol

Protocol specifys L4 match criteria in a packet.

`dns` - (Optional) Match all DNS packets. See [Dns ](#dns) below for details.

`icmp` - (Optional) ICMP message types to be matched in packet. See [Icmp ](#icmp) below for details.

`tcp` - (Optional) TCP flags to be matched in packet. See [Tcp ](#tcp) below for details.

`udp` - (Optional) Match all UDP packets. See [Udp ](#udp) below for details.

### Protocol Policer

List of L4 protocol match condition and associated traffic rate limits.

`policer` - (Required) Reference to policer object to apply traffic rate limits. See [ref](#ref) below for details.

`protocol` - (Required) Protocol specifys L4 match criteria in a packet. See [Protocol ](#protocol) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Tcp

TCP flags to be matched in packet.

`flags` - (Optional)flags (`List of Strings`).

### Udp

Match all UDP packets.

Attribute Reference
-------------------

-	`id` - This is the id of the configured protocol_policer.
