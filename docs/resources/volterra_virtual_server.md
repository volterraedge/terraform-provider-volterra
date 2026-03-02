---

page_title: "Volterra: virtual_server"

description: "The virtual_server allows CRUD of Virtual Server resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_virtual_server
================================

The Virtual Server allows CRUD of Virtual Server resource on Volterra SaaS

~> **Note:** Please refer to [Virtual Server API docs](https://docs.cloud.f5.com/docs-v2/api/views-virtual-server) to learn more

Example Usage
-------------

```hcl
resource "volterra_virtual_server" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "managed not_managed" must be set

  managed {
    domains {
      dns_zone {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }

      prefix = "prefix"
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

`auto_last_hop` - (Optional) When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a default route configured and the client is located on a remote network. This setting is also useful when the system is load balancing transparent devices that do not modify the source IP address of the packet. Without the last hop option enabled, the system could return connections to a different transparent node, resulting in asymmetric routing. You can configure this setting globally and on an object level. You set the global Auto Last Hop value on the System :: Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the Default setting. When you configure Auto Last Hop with a value other than Default at the object level, its setting takes precedence over the global setting. This enables you to configure auto last hop on a per-virtual server basis. The default is Default, meaning that the system uses the global auto-lasthop setting to send back the request.. See [Auto Last Hop ](#auto-last-hop) below for details.

`clone_pool_options` - (Optional) x-displayName: "Clone Pool". See [Clone Pool Options ](#clone-pool-options) below for details.

`connection_limit_options` - (Optional) x-displayName: "Connection Limit Options". See [Connection Limit Options ](#connection-limit-options) below for details.

###### One of the arguments from this list "default_persistence_profile, default_persistence_profile_none" can be set

`default_persistence_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`default_persistence_profile_none` - (Optional) x-displayName: "None" (`Bool`).

###### One of the arguments from this list "default_pool, default_pool_none" can be set

`default_pool` - (Optional) x-displayName: "Select Pool". See [ref](#ref) below for details.

`default_pool_none` - (Optional) x-displayName: "None" (`Bool`).

###### One of the arguments from this list "managed, not_managed" must be set

`managed` - (Optional) Domains that are managed by F5XC platform through setting up of dns_zones. See [Domain Choice Managed ](#domain-choice-managed) below for details.

`not_managed` - (Optional) Not Managed by F5XC. See [Domain Choice Not Managed ](#domain-choice-not-managed) below for details.

###### One of the arguments from this list "fallback_persistence_profile, fallback_persistence_profile_none" can be set

`fallback_persistence_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`fallback_persistence_profile_none` - (Optional) x-displayName: "None" (`Bool`).

`immediate_action_on_service_down` - (Optional) - Drop: Specifies that the system drops the connections when the virtual server is reported Offline or Unavailable.. See [Immediate Action On Service Down ](#immediate-action-on-service-down) below for details.

`irules` - (Optional) Specifies the iRules you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.. See [ref](#ref) below for details.

###### One of the arguments from this list "last_hop_pool, last_hop_pool_none" can be set

`last_hop_pool` - (Optional) x-displayName: "Select Pool". See [ref](#ref) below for details.

`last_hop_pool_none` - (Optional) x-displayName: "None" (`Bool`).

`nat64` - (Optional) When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a default route configured and the client is located on a remote network. This setting is also useful when the system is load balancing transparent devices that do not modify the source IP address of the packet. Without the last hop option enabled, the system could return connections to a different transparent node, resulting in asymmetric routing. You can configure this setting globally and on an object level. You set the global Auto Last Hop value on the System :: Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the Default setting. When you configure Auto Last Hop with a value other than Default at the object level, its setting takes precedence over the global setting. This enables you to configure auto last hop on a per-virtual server basis. The default is Default, meaning that the system uses the global auto-lasthop setting to send back the request.. See [Nat64 ](#nat64) below for details.

###### One of the arguments from this list "request_logging_profile, request_logging_profile_none" can be set

`request_logging_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`request_logging_profile_none` - (Optional) x-displayName: "None" (`Bool`).

`state` - (Optional) State. See [State ](#state) below for details.

`traffic_policies` - (Optional) Specifies the Traffic Policy you want use for this virtual server.. See [ref](#ref) below for details.

`translations` - (Optional) x-displayName: "Translation". See [Translations ](#translations) below for details.

`virtual_addresses` - (Optional) Specify the virtual address objects to be associated with this virtual server. See [ref](#ref) below for details.

###### One of the arguments from this list "http, https, tcp, udp" can be set

`http` - (Optional) x-displayName: "HTTP". See [Virtual Server Type Http ](#virtual-server-type-http) below for details.

`https` - (Optional) x-displayName: "HTTPS". See [Virtual Server Type Https ](#virtual-server-type-https) below for details.

`tcp` - (Optional) x-displayName: "TCP". See [Virtual Server Type Tcp ](#virtual-server-type-tcp) below for details.

`udp` - (Optional) x-displayName: "UDP". See [Virtual Server Type Udp ](#virtual-server-type-udp) below for details.

`vs_score` - (Optional) Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value to load balance traffic in a proportional manner. The default is 0, meaning that no additional metric is applied for the virtual server. (`Int`).

### Auto Last Hop

When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a default route configured and the client is located on a remote network. This setting is also useful when the system is load balancing transparent devices that do not modify the source IP address of the packet. Without the last hop option enabled, the system could return connections to a different transparent node, resulting in asymmetric routing. You can configure this setting globally and on an object level. You set the global Auto Last Hop value on the System :: Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the Default setting. When you configure Auto Last Hop with a value other than Default at the object level, its setting takes precedence over the global setting. This enables you to configure auto last hop on a per-virtual server basis. The default is Default, meaning that the system uses the global auto-lasthop setting to send back the request..

###### One of the arguments from this list "auto_last_hop_default, auto_last_hop_disable, auto_last_hop_enable" can be set

`auto_last_hop_default` - (Optional) x-displayName: "Default". See [Auto Last Hop Choice Auto Last Hop Default ](#auto-last-hop-choice-auto-last-hop-default) below for details.

`auto_last_hop_disable` - (Optional) x-displayName: "Disable". See [Auto Last Hop Choice Auto Last Hop Disable ](#auto-last-hop-choice-auto-last-hop-disable) below for details.

`auto_last_hop_enable` - (Optional) x-displayName: "Enable". See [Auto Last Hop Choice Auto Last Hop Enable ](#auto-last-hop-choice-auto-last-hop-enable) below for details.

### Clone Pool Options

x-displayName: "Clone Pool".

###### One of the arguments from this list "pool_client, pool_client_none" can be set

`pool_client` - (Optional) x-displayName: "Select Pool". See [ref](#ref) below for details.

`pool_client_none` - (Optional) x-displayName: "None" (`Bool`).

###### One of the arguments from this list "pool_server, pool_server_none" can be set

`pool_server` - (Optional) x-displayName: "Select Pool". See [ref](#ref) below for details.

`pool_server_none` - (Optional) x-displayName: "None" (`Bool`).

### Connection Limit Options

x-displayName: "Connection Limit Options".

`connection_limit` - (Optional) Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this to 0 turns off connection limits. The default is 0. (`Int`).

`connection_rate_limit` - (Optional) Specifies the maximum number of connections-per-second allowed for a virtual server. When the number of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where connection requests flood a virtual server. Setting this to 0 turns off connection limits. The default is 0. (`Int`).

`connection_rate_limit_mode` - (Optional) - Per Source and Destination Address (All Rate Limiting Virtual Servers): Applies rate limiting based on the specified source and destination addresses for all virtual servers that have rate limits specified.. See [Connection Limit Options Connection Rate Limit Mode ](#connection-limit-options-connection-rate-limit-mode) below for details.

### Immediate Action On Service Down

-   Drop: Specifies that the system drops the connections when the virtual server is reported Offline or Unavailable..

###### One of the arguments from this list "immediate_action_on_service_down_drop, immediate_action_on_service_down_none, immediate_action_on_service_down_reset" can be set

`immediate_action_on_service_down_drop` - (Optional) x-displayName: "Drop". See [Immediate Action On Service Down Choice Immediate Action On Service Down Drop ](#immediate-action-on-service-down-choice-immediate-action-on-service-down-drop) below for details.

`immediate_action_on_service_down_none` - (Optional) x-displayName: "None". See [Immediate Action On Service Down Choice Immediate Action On Service Down None ](#immediate-action-on-service-down-choice-immediate-action-on-service-down-none) below for details.

`immediate_action_on_service_down_reset` - (Optional) x-displayName: "Reset". See [Immediate Action On Service Down Choice Immediate Action On Service Down Reset ](#immediate-action-on-service-down-choice-immediate-action-on-service-down-reset) below for details.

### Nat64

When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a default route configured and the client is located on a remote network. This setting is also useful when the system is load balancing transparent devices that do not modify the source IP address of the packet. Without the last hop option enabled, the system could return connections to a different transparent node, resulting in asymmetric routing. You can configure this setting globally and on an object level. You set the global Auto Last Hop value on the System :: Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the Default setting. When you configure Auto Last Hop with a value other than Default at the object level, its setting takes precedence over the global setting. This enables you to configure auto last hop on a per-virtual server basis. The default is Default, meaning that the system uses the global auto-lasthop setting to send back the request..

###### One of the arguments from this list "nat64_disable, nat64_enable" can be set

`nat64_disable` - (Optional) x-displayName: "Disable". See [Nat64 Choice Nat64 Disable ](#nat64-choice-nat64-disable) below for details.

`nat64_enable` - (Optional) x-displayName: "Enable". See [Nat64 Choice Nat64 Enable ](#nat64-choice-nat64-enable) below for details.

### State

State.

###### One of the arguments from this list "state_disabled, state_enabled" can be set

`state_disabled` - (Optional) x-displayName: "Disabled". See [State Choice State Disabled ](#state-choice-state-disabled) below for details.

`state_enabled` - (Optional) x-displayName: "Enabled". See [State Choice State Enabled ](#state-choice-state-enabled) below for details.

### Translations

x-displayName: "Translation".

`address_translation` - (Optional) Specifies, when checked (enabled), that the system translates the address of the virtual server. When cleared (disabled), specifies that the system uses the address without translation. This option is useful when the system is load balancing devices that have the same IP address. The default is enabled.. See [Translations Address Translation ](#translations-address-translation) below for details.

`port_translation` - (Optional) Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any service. The default is enabled.. See [Translations Port Translation ](#translations-port-translation) below for details.

`source_port` - (Optional) Change: Specifies that the system changes the source port. This setting is useful for obfuscating internal network addresses.. See [Translations Source Port ](#translations-source-port) below for details.

### Address Translation Choice Address Translation Disable

x-displayName: "Disable".

### Address Translation Choice Address Translation Enable

x-displayName: "Enable".

### Auto Last Hop Choice Auto Last Hop Default

x-displayName: "Default".

### Auto Last Hop Choice Auto Last Hop Disable

x-displayName: "Disable".

### Auto Last Hop Choice Auto Last Hop Enable

x-displayName: "Enable".

### Clone Pool Client Choice Pool Client None

x-displayName: "None".

### Clone Pool Server Choice Pool Server None

x-displayName: "None".

### Connection Limit Options Connection Rate Limit Mode

-   Per Source and Destination Address (All Rate Limiting Virtual Servers): Applies rate limiting based on the specified source and destination addresses for all virtual servers that have rate limits specified..

###### One of the arguments from this list "per_destination_address, per_source_address, per_source_destination_address, per_virtual_server, per_virtual_server_destination_address, per_virtual_server_source_address, per_virtual_server_source_destination_address" can be set

`per_destination_address` - (Optional) x-displayName: "Per Destination Address (All Rate Limiting Virtual Servers)". See [Connection Rate Limit Mode Choice Per Destination Address ](#connection-rate-limit-mode-choice-per-destination-address) below for details.

`per_source_address` - (Optional) x-displayName: "Per Source Address (All Rate Limiting Virtual Servers)". See [Connection Rate Limit Mode Choice Per Source Address ](#connection-rate-limit-mode-choice-per-source-address) below for details.

`per_source_destination_address` - (Optional) x-displayName: "Per Source And Destination Address (All Rate Limiting Virtual Servers)". See [Connection Rate Limit Mode Choice Per Source Destination Address ](#connection-rate-limit-mode-choice-per-source-destination-address) below for details.

`per_virtual_server` - (Optional) x-displayName: "Per Virtual Server". See [Connection Rate Limit Mode Choice Per Virtual Server ](#connection-rate-limit-mode-choice-per-virtual-server) below for details.

`per_virtual_server_destination_address` - (Optional) x-displayName: "Per Virtual Server and Destination Address". See [Connection Rate Limit Mode Choice Per Virtual Server Destination Address ](#connection-rate-limit-mode-choice-per-virtual-server-destination-address) below for details.

`per_virtual_server_source_address` - (Optional) x-displayName: "Per Virtual Server and Source Address". See [Connection Rate Limit Mode Choice Per Virtual Server Source Address ](#connection-rate-limit-mode-choice-per-virtual-server-source-address) below for details.

`per_virtual_server_source_destination_address` - (Optional) x-displayName: "Per Virtual Server, Destination and Source Address". See [Connection Rate Limit Mode Choice Per Virtual Server Source Destination Address ](#connection-rate-limit-mode-choice-per-virtual-server-source-destination-address) below for details.

### Connection Rate Limit Mode Choice Per Destination Address

x-displayName: "Per Destination Address (All Rate Limiting Virtual Servers)".

`destination_mask` - (Optional) x-displayName: "Destination Mask" (`Int`).

### Connection Rate Limit Mode Choice Per Source Address

x-displayName: "Per Source Address (All Rate Limiting Virtual Servers)".

`source_mask` - (Optional) x-displayName: "Source Mask" (`Int`).

### Connection Rate Limit Mode Choice Per Source Destination Address

x-displayName: "Per Source And Destination Address (All Rate Limiting Virtual Servers)".

`destination_mask` - (Optional) x-displayName: "Destination Mask" (`Int`).

`source_mask` - (Optional) x-displayName: "Source Mask" (`Int`).

### Connection Rate Limit Mode Choice Per Virtual Server

x-displayName: "Per Virtual Server".

### Connection Rate Limit Mode Choice Per Virtual Server Destination Address

x-displayName: "Per Virtual Server and Destination Address".

`destination_mask` - (Optional) x-displayName: "Destination Mask" (`Int`).

### Connection Rate Limit Mode Choice Per Virtual Server Source Address

x-displayName: "Per Virtual Server and Source Address".

`source_mask` - (Optional) x-displayName: "Source Mask" (`Int`).

### Connection Rate Limit Mode Choice Per Virtual Server Source Destination Address

x-displayName: "Per Virtual Server, Destination and Source Address".

`destination_mask` - (Optional) x-displayName: "Destination Mask" (`Int`).

`source_mask` - (Optional) x-displayName: "Source Mask" (`Int`).

### Domain Choice Managed

Domains that are managed by F5XC platform through setting up of dns_zones.

`domains` - (Required) A list of Domains (host/authority header) that will be matched to load balancer.. See [Managed Domains ](#managed-domains) below for details.

### Domain Choice Not Managed

Not Managed by F5XC.

`domains` - (Required) Domains also indicate the list of names for which DNS resolution will be automatically resolved to IP addresses by the system. (`String`).

### Fix Profile Choice Fix Profile None

x-displayName: "None".

### Http Http Profiles

Specifies that the selected HTTP profile is a client-side and server-side profile. The list contains entries for each already defined protocol profile..

`client_profile` - (Required) Specifies that the selected profile is a client-side HTTP profile. The list contains entries for each already defined client protocol profile.. See [ref](#ref) below for details.

###### One of the arguments from this list "server_profile, server_profile_same_as_client" can be set

`server_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`server_profile_same_as_client` - (Optional) x-displayName: "Use Client Profile" (`Bool`).

### Http Services

x-displayName: "Services".

###### One of the arguments from this list "port, port_ranges" can be set

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Http Tcp Profiles

Specifies that the selected profile is a client-side and server-side profile. The list contains entries for each already defined protocol profile..

`client_profile` - (Required) Specifies that the selected profile is a client-side profile. The list contains entries for each already defined client protocol profile.. See [ref](#ref) below for details.

###### One of the arguments from this list "server_profile, server_profile_same_as_client" can be set

`server_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`server_profile_same_as_client` - (Optional) x-displayName: "Use Client Profile" (`Bool`).

### Http Websocket Profiles

Specifies that the selected Websocket profile is a client-side and server-side profile. The list contains entries for each already defined protocol profile..

###### One of the arguments from this list "client_profile, client_profile_none" can be set

`client_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`client_profile_none` - (Optional) x-displayName: "None" (`Bool`).

###### One of the arguments from this list "server_profile, server_profile_same_as_client" can be set

`server_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`server_profile_same_as_client` - (Optional) x-displayName: "Use Client Profile" (`Bool`).

### Http Server Profile Choice Server Profile Same As Client

x-displayName: "Use Client Profile".

### Https Services

x-displayName: "Services".

###### One of the arguments from this list "port, port_ranges" can be set

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Immediate Action On Service Down Choice Immediate Action On Service Down Drop

x-displayName: "Drop".

### Immediate Action On Service Down Choice Immediate Action On Service Down None

x-displayName: "None".

### Immediate Action On Service Down Choice Immediate Action On Service Down Reset

x-displayName: "Reset".

### Managed Domains

A list of Domains (host/authority header) that will be matched to load balancer..

`dns_zone` - (Required) DNS Zone managed by F5XC. See [ref](#ref) below for details.

`prefix` - (Required) x-required (`String`).

### Nat64 Choice Nat64 Disable

x-displayName: "Disable".

### Nat64 Choice Nat64 Enable

x-displayName: "Enable".

### Port Translation Choice Port Translation Disable

x-displayName: "Disable".

### Port Translation Choice Port Translation Enable

x-displayName: "Enable".

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Source Port Choice Source Port Change

x-displayName: "Change".

### Source Port Choice Source Port Preserve

x-displayName: "Preserve".

### Source Port Choice Source Port Preserve Strict

x-displayName: "Preserve Strict".

### State Choice State Disabled

x-displayName: "Disabled".

### State Choice State Enabled

x-displayName: "Enabled".

### Stream Profile Choice Stream Profile None

x-displayName: "None".

### Tcp Services

x-displayName: "Services".

###### One of the arguments from this list "port, port_ranges" can be set

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Tcp Server Profile Choice Server Profile Same As Client

x-displayName: "Use Client Profile".

### Translations Address Translation

Specifies, when checked (enabled), that the system translates the address of the virtual server. When cleared (disabled), specifies that the system uses the address without translation. This option is useful when the system is load balancing devices that have the same IP address. The default is enabled..

###### One of the arguments from this list "address_translation_disable, address_translation_enable" can be set

`address_translation_disable` - (Optional) x-displayName: "Disable". See [Address Translation Choice Address Translation Disable ](#address-translation-choice-address-translation-disable) below for details.

`address_translation_enable` - (Optional) x-displayName: "Enable". See [Address Translation Choice Address Translation Enable ](#address-translation-choice-address-translation-enable) below for details.

### Translations Port Translation

Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any service. The default is enabled..

###### One of the arguments from this list "port_translation_disable, port_translation_enable" can be set

`port_translation_disable` - (Optional) x-displayName: "Disable". See [Port Translation Choice Port Translation Disable ](#port-translation-choice-port-translation-disable) below for details.

`port_translation_enable` - (Optional) x-displayName: "Enable". See [Port Translation Choice Port Translation Enable ](#port-translation-choice-port-translation-enable) below for details.

### Translations Source Port

Change: Specifies that the system changes the source port. This setting is useful for obfuscating internal network addresses..

###### One of the arguments from this list "source_port_change, source_port_preserve, source_port_preserve_strict" can be set

`source_port_change` - (Optional) x-displayName: "Change". See [Source Port Choice Source Port Change ](#source-port-choice-source-port-change) below for details.

`source_port_preserve` - (Optional) x-displayName: "Preserve". See [Source Port Choice Source Port Preserve ](#source-port-choice-source-port-preserve) below for details.

`source_port_preserve_strict` - (Optional) x-displayName: "Preserve Strict". See [Source Port Choice Source Port Preserve Strict ](#source-port-choice-source-port-preserve-strict) below for details.

### Udp Services

x-displayName: "Services".

###### One of the arguments from this list "port, port_ranges" can be set

`port`- (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Virtual Server Type Http

x-displayName: "HTTP".

###### One of the arguments from this list "fix_profile, fix_profile_none" can be set

`fix_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`fix_profile_none` - (Optional) x-displayName: "None" (`Bool`).

`http_profiles` - (Optional) Specifies that the selected HTTP profile is a client-side and server-side profile. The list contains entries for each already defined protocol profile.. See [Http Http Profiles ](#http-http-profiles) below for details.

`services` - (Optional) x-displayName: "Services". See [Http Services ](#http-services) below for details.

###### One of the arguments from this list "stream_profile, stream_profile_none" can be set

`stream_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`stream_profile_none` - (Optional) x-displayName: "None" (`Bool`).

`tcp_profiles` - (Optional) Specifies that the selected profile is a client-side and server-side profile. The list contains entries for each already defined protocol profile.. See [Http Tcp Profiles ](#http-tcp-profiles) below for details.

`websocket_profiles` - (Optional) Specifies that the selected Websocket profile is a client-side and server-side profile. The list contains entries for each already defined protocol profile.. See [Http Websocket Profiles ](#http-websocket-profiles) below for details.

### Virtual Server Type Https

x-displayName: "HTTPS".

`services` - (Optional) x-displayName: "Services". See [Https Services ](#https-services) below for details.

### Virtual Server Type Tcp

x-displayName: "TCP".

`services` - (Optional) x-displayName: "Services". See [Tcp Services ](#tcp-services) below for details.

### Virtual Server Type Udp

x-displayName: "UDP".

`services` - (Optional) x-displayName: "Services". See [Udp Services ](#udp-services) below for details.

### Websocket Client Profile Choice Client Profile None

x-displayName: "None".

### Websocket Server Profile Choice Server Profile Same As Client

x-displayName: "Use Client Profile".

Attribute Reference
-------------------

*   `id` - This is the id of the configured virtual_server.
