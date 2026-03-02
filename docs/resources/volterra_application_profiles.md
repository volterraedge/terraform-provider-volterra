---

page_title: "Volterra: application_profiles"

description: "The application_profiles allows CRUD of Application Profiles resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------

Resource volterra_application_profiles
======================================

The Application Profiles allows CRUD of Application Profiles resource on Volterra SaaS

~> **Note:** Please refer to [Application Profiles API docs](https://docs.cloud.f5.com/docs-v2/api/application-profiles) to learn more

Example Usage
-------------

```hcl
resource "volterra_application_profiles" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  ddos_profile {
    // One of the arguments from this list "disable_ddos_mitigation enable_ddos_mitigation" can be set

    disable_ddos_mitigation = true
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

`advanced_tcp_profile` - (Optional) On enabling this feature, a default profile is applied to the Proxy. See [Advanced Tcp Profile ](#advanced-tcp-profile) below for details.

`ddos_profile` - (Required) On enabling this feature, a default DDoS profile is applied to the Proxy. See [Ddos Profile ](#ddos-profile) below for details.

`irules` - (Optional) Options for attaching iRules to BIG-IP Proxy. See [ref](#ref) below for details.

`traffic_policies` - (Optional) x-displayName: "Traffic Policy". See [ref](#ref) below for details.

`virtual_server` - (Optional) Specifies configuration related to virtual server. See [Virtual Server ](#virtual-server) below for details.

### Advanced Tcp Profile

On enabling this feature, a default profile is applied to the Proxy.

###### One of the arguments from this list "disable_tcp_advanced_profile, enable_tcp_advanced_profile" can be set

`disable_tcp_advanced_profile` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_tcp_advanced_profile` - (Optional) x-displayName: "Enable" (`Bool`).

### Ddos Profile

On enabling this feature, a default DDoS profile is applied to the Proxy.

###### One of the arguments from this list "disable_ddos_mitigation, enable_ddos_mitigation" can be set

`disable_ddos_mitigation` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_ddos_mitigation` - (Optional) x-displayName: "Enable" (`Bool`).

### Virtual Server

Specifies configuration related to virtual server.

`address_translation` - (Optional) Specifies, when checked (enabled), that the system translates the address of the virtual server. When cleared (disabled), specifies that the system uses the address without translation. This option is useful when the system is load balancing devices that have the same IP address. The default is enabled.. See [Virtual Server Address Translation ](#virtual-server-address-translation) below for details.

`auto_last_hop` - (Optional) When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a default route configured and the client is located on a remote network. This setting is also useful when the system is load balancing transparent devices that do not modify the source IP address of the packet. Without the last hop option enabled, the system could return connections to a different transparent node, resulting in asymmetric routing. You can configure this setting globally and on an object level. You set the global Auto Last Hop value on the System :: Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the Default setting. When you configure Auto Last Hop with a value other than Default at the object level, its setting takes precedence over the global setting. This enables you to configure auto last hop on a per-virtual server basis. The default is Default, meaning that the system uses the global auto-lasthop setting to send back the request.. See [Virtual Server Auto Last Hop ](#virtual-server-auto-last-hop) below for details.

`clone_pool_client` - (Optional) Replicates client-side traffic (that is, prior to address translation) to a member of the specified pool.. See [ref](#ref) below for details.

`clone_pool_server` - (Optional) Replicates server-side traffic (that is, prior to address translation) to a member of the specified pool.. See [ref](#ref) below for details.

`connection_limit` - (Optional) Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this to 0 turns off connection limits. The default is 0. (`Int`).

`connection_rate_limit` - (Optional) Specifies the maximum number of connections-per-second allowed for a virtual server. When the number of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where connection requests flood a virtual server. Setting this to 0 turns off connection limits. The default is 0. (`Int`).

`connection_rate_limit_mode` - (Optional) Controls the scope of the rate limit defined for the virtual server, specifically, per virtual server object, per source address, per destination address, or a combination of source and destination addresses. The default is Per Virtual Server, which does not use the source or destination address.. See [Virtual Server Connection Rate Limit Mode ](#virtual-server-connection-rate-limit-mode) below for details.

`default_persistence_profile` - (Optional) x-displayName: "Default Persistence Profile". See [ref](#ref) below for details.

`default_pool` - (Optional) Specifies the pool name that you want the virtual server to use as the default pool. A load balancing virtual server sends traffic to this pool automatically, unless an iRule directs the server to send the traffic to another pool instead.. See [ref](#ref) below for details.

`fallback_persistence_profile` - (Optional) x-displayName: "Fallback Persistence Profile". See [ref](#ref) below for details.

`fix_profile` - (Optional) x-displayName: "FIX Profile". See [ref](#ref) below for details.

`immediate_action_on_service_down` - (Optional) Drop: Specifies that the system drops the connections when the virtual server is reported Offline or Unavailable.. See [Virtual Server Immediate Action On Service Down ](#virtual-server-immediate-action-on-service-down) below for details.

`last_hop_pool` - (Optional) Directs reply traffic to the last hop router using the specified pool.. See [ref](#ref) below for details.

`nat64` - (Optional) When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a default route configured and the client is located on a remote network. This setting is also useful when the system is load balancing transparent devices that do not modify the source IP address of the packet. Without the last hop option enabled, the system could return connections to a different transparent node, resulting in asymmetric routing. You can configure this setting globally and on an object level. You set the global Auto Last Hop value on the System :: Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the Default setting. When you configure Auto Last Hop with a value other than Default at the object level, its setting takes precedence over the global setting. This enables you to configure auto last hop on a per-virtual server basis. The default is Default, meaning that the system uses the global auto-lasthop setting to send back the request.. See [Virtual Server Nat64 ](#virtual-server-nat64) below for details.

`port_translation` - (Optional) Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any service. The default is enabled.. See [Virtual Server Port Translation ](#virtual-server-port-translation) below for details.

`request_logging_profile` - (Optional) x-displayName: "Request Logging Profile". See [ref](#ref) below for details.

`source_port` - (Optional) Change: Specifies that the system changes the source port. This setting is useful for obfuscating internal network addresses.. See [Virtual Server Source Port ](#virtual-server-source-port) below for details.

`statistics_profile` - (Optional) x-displayName: "Statistics Profile". See [ref](#ref) below for details.

`virtual_server_state` - (Optional) State. See [Virtual Server Virtual Server State ](#virtual-server-virtual-server-state) below for details.

###### One of the arguments from this list "http, https, tcp, udp" can be set

`http` - (Optional) x-displayName: "HTTP". See [Virtual Server Type Http ](#virtual-server-type-http) below for details.

`https` - (Optional) x-displayName: "HTTPS". See [Virtual Server Type Https ](#virtual-server-type-https) below for details.

`tcp` - (Optional) x-displayName: "TCP". See [Virtual Server Type Tcp ](#virtual-server-type-tcp) below for details.

`udp` - (Optional) x-displayName: "UDP". See [Virtual Server Type Udp ](#virtual-server-type-udp) below for details.

`vs_score` - (Optional) Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value to load balance traffic in a proportional manner. The default is 0, meaning that no additional metric is applied for the virtual server. (`Int`).

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

### Ddos Mitigation Choice Disable Ddos Mitigation

x-displayName: "Disable".

### Ddos Mitigation Choice Enable Ddos Mitigation

x-displayName: "Enable".

### Immediate Action On Service Down Choice Immediate Action On Service Down Drop

x-displayName: "Drop".

### Immediate Action On Service Down Choice Immediate Action On Service Down None

x-displayName: "None".

### Immediate Action On Service Down Choice Immediate Action On Service Down Reset

x-displayName: "Reset".

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

### Tcp Advanced Profile Choice Disable Tcp Advanced Profile

x-displayName: "Disable".

### Tcp Advanced Profile Choice Enable Tcp Advanced Profile

x-displayName: "Enable".

### Virtual Server Address Translation

Specifies, when checked (enabled), that the system translates the address of the virtual server. When cleared (disabled), specifies that the system uses the address without translation. This option is useful when the system is load balancing devices that have the same IP address. The default is enabled..

###### One of the arguments from this list "address_translation_disable, address_translation_enable" can be set

`address_translation_disable` - (Optional) x-displayName: "Disable". See [Address Translation Choice Address Translation Disable ](#address-translation-choice-address-translation-disable) below for details.

`address_translation_enable` - (Optional) x-displayName: "Enable". See [Address Translation Choice Address Translation Enable ](#address-translation-choice-address-translation-enable) below for details.

### Virtual Server Auto Last Hop

When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a default route configured and the client is located on a remote network. This setting is also useful when the system is load balancing transparent devices that do not modify the source IP address of the packet. Without the last hop option enabled, the system could return connections to a different transparent node, resulting in asymmetric routing. You can configure this setting globally and on an object level. You set the global Auto Last Hop value on the System :: Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the Default setting. When you configure Auto Last Hop with a value other than Default at the object level, its setting takes precedence over the global setting. This enables you to configure auto last hop on a per-virtual server basis. The default is Default, meaning that the system uses the global auto-lasthop setting to send back the request..

###### One of the arguments from this list "auto_last_hop_default, auto_last_hop_disable, auto_last_hop_enable" can be set

`auto_last_hop_default` - (Optional) x-displayName: "Default". See [Auto Last Hop Choice Auto Last Hop Default ](#auto-last-hop-choice-auto-last-hop-default) below for details.

`auto_last_hop_disable` - (Optional) x-displayName: "Disable". See [Auto Last Hop Choice Auto Last Hop Disable ](#auto-last-hop-choice-auto-last-hop-disable) below for details.

`auto_last_hop_enable` - (Optional) x-displayName: "Enable". See [Auto Last Hop Choice Auto Last Hop Enable ](#auto-last-hop-choice-auto-last-hop-enable) below for details.

### Virtual Server Connection Rate Limit Mode

Controls the scope of the rate limit defined for the virtual server, specifically, per virtual server object, per source address, per destination address, or a combination of source and destination addresses. The default is Per Virtual Server, which does not use the source or destination address..

###### One of the arguments from this list "per_destination_address, per_source_address, per_source_destination_address, per_virtual_server, per_virtual_server_destination_address, per_virtual_server_source_address, per_virtual_server_source_destination_address" can be set

`per_destination_address` - (Optional) x-displayName: "Per Destination Address (All Rate Limiting Virtual Servers)". See [Connection Rate Limit Mode Choice Per Destination Address ](#connection-rate-limit-mode-choice-per-destination-address) below for details.

`per_source_address` - (Optional) x-displayName: "Per Source Address (All Rate Limiting Virtual Servers)". See [Connection Rate Limit Mode Choice Per Source Address ](#connection-rate-limit-mode-choice-per-source-address) below for details.

`per_source_destination_address` - (Optional) x-displayName: "Per Source And Destination Address (All Rate Limiting Virtual Servers)". See [Connection Rate Limit Mode Choice Per Source Destination Address ](#connection-rate-limit-mode-choice-per-source-destination-address) below for details.

`per_virtual_server` - (Optional) x-displayName: "Per Virtual Server". See [Connection Rate Limit Mode Choice Per Virtual Server ](#connection-rate-limit-mode-choice-per-virtual-server) below for details.

`per_virtual_server_destination_address` - (Optional) x-displayName: "Per Virtual Server and Destination Address". See [Connection Rate Limit Mode Choice Per Virtual Server Destination Address ](#connection-rate-limit-mode-choice-per-virtual-server-destination-address) below for details.

`per_virtual_server_source_address` - (Optional) x-displayName: "Per Virtual Server and Source Address". See [Connection Rate Limit Mode Choice Per Virtual Server Source Address ](#connection-rate-limit-mode-choice-per-virtual-server-source-address) below for details.

`per_virtual_server_source_destination_address` - (Optional) x-displayName: "Per Virtual Server, Destination and Source Address". See [Connection Rate Limit Mode Choice Per Virtual Server Source Destination Address ](#connection-rate-limit-mode-choice-per-virtual-server-source-destination-address) below for details.

### Virtual Server Immediate Action On Service Down

Drop: Specifies that the system drops the connections when the virtual server is reported Offline or Unavailable..

###### One of the arguments from this list "immediate_action_on_service_down_drop, immediate_action_on_service_down_none, immediate_action_on_service_down_reset" can be set

`immediate_action_on_service_down_drop` - (Optional) x-displayName: "Drop". See [Immediate Action On Service Down Choice Immediate Action On Service Down Drop ](#immediate-action-on-service-down-choice-immediate-action-on-service-down-drop) below for details.

`immediate_action_on_service_down_none` - (Optional) x-displayName: "None". See [Immediate Action On Service Down Choice Immediate Action On Service Down None ](#immediate-action-on-service-down-choice-immediate-action-on-service-down-none) below for details.

`immediate_action_on_service_down_reset` - (Optional) x-displayName: "Reset". See [Immediate Action On Service Down Choice Immediate Action On Service Down Reset ](#immediate-action-on-service-down-choice-immediate-action-on-service-down-reset) below for details.

### Virtual Server Nat64

When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a default route configured and the client is located on a remote network. This setting is also useful when the system is load balancing transparent devices that do not modify the source IP address of the packet. Without the last hop option enabled, the system could return connections to a different transparent node, resulting in asymmetric routing. You can configure this setting globally and on an object level. You set the global Auto Last Hop value on the System :: Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the Default setting. When you configure Auto Last Hop with a value other than Default at the object level, its setting takes precedence over the global setting. This enables you to configure auto last hop on a per-virtual server basis. The default is Default, meaning that the system uses the global auto-lasthop setting to send back the request..

###### One of the arguments from this list "nat64_disable, nat64_enable" can be set

`nat64_disable` - (Optional) x-displayName: "Disable". See [Nat64 Choice Nat64 Disable ](#nat64-choice-nat64-disable) below for details.

`nat64_enable` - (Optional) x-displayName: "Enable". See [Nat64 Choice Nat64 Enable ](#nat64-choice-nat64-enable) below for details.

### Virtual Server Port Translation

Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any service. The default is enabled..

###### One of the arguments from this list "port_translation_disable, port_translation_enable" can be set

`port_translation_disable` - (Optional) x-displayName: "Disable". See [Port Translation Choice Port Translation Disable ](#port-translation-choice-port-translation-disable) below for details.

`port_translation_enable` - (Optional) x-displayName: "Enable". See [Port Translation Choice Port Translation Enable ](#port-translation-choice-port-translation-enable) below for details.

### Virtual Server Source Port

Change: Specifies that the system changes the source port. This setting is useful for obfuscating internal network addresses..

###### One of the arguments from this list "source_port_change, source_port_preserve, source_port_preserve_strict" can be set

`source_port_change` - (Optional) x-displayName: "Change". See [Source Port Choice Source Port Change ](#source-port-choice-source-port-change) below for details.

`source_port_preserve` - (Optional) x-displayName: "Preserve". See [Source Port Choice Source Port Preserve ](#source-port-choice-source-port-preserve) below for details.

`source_port_preserve_strict` - (Optional) x-displayName: "Preserve Strict". See [Source Port Choice Source Port Preserve Strict ](#source-port-choice-source-port-preserve-strict) below for details.

### Virtual Server Virtual Server State

State.

###### One of the arguments from this list "state_disabled, state_enabled" can be set

`state_disabled` - (Optional) x-displayName: "Disabled". See [State Choice State Disabled ](#state-choice-state-disabled) below for details.

`state_enabled` - (Optional) x-displayName: "Enabled". See [State Choice State Enabled ](#state-choice-state-enabled) below for details.

### Virtual Server Type Http

x-displayName: "HTTP".

`http_client_profile` - (Optional) x-displayName: "HTTP Profile (Client)". See [ref](#ref) below for details.

`http_server_profile` - (Optional) x-displayName: "HTTP Profile (Server)". See [ref](#ref) below for details.

`stream_profile` - (Optional) x-displayName: "Stream Profile". See [ref](#ref) below for details.

`tcp_client_profile` - (Optional) x-displayName: "Protocol Profile (Client)". See [ref](#ref) below for details.

`tcp_server_profile` - (Optional) x-displayName: "Protocol Profile (Server)". See [ref](#ref) below for details.

`websocket_client_profile` - (Optional) x-displayName: "WebSocket Profile Client". See [ref](#ref) below for details.

`websocket_server_profile` - (Optional) x-displayName: "WebSocket Profile Server". See [ref](#ref) below for details.

### Virtual Server Type Https

x-displayName: "HTTPS".

`http_client_profile` - (Optional) x-displayName: "HTTP Profile (Client)". See [ref](#ref) below for details.

`http_server_profile` - (Optional) x-displayName: "HTTP Profile (Server)". See [ref](#ref) below for details.

`stream_profile` - (Optional) x-displayName: "Stream Profile". See [ref](#ref) below for details.

`tcp_client_profile` - (Optional) x-displayName: "Protocol Profile (Client)". See [ref](#ref) below for details.

`tcp_server_profile` - (Optional) x-displayName: "Protocol Profile (Server)". See [ref](#ref) below for details.

`websocket_client_profile` - (Optional) x-displayName: "WebSocket Profile Client". See [ref](#ref) below for details.

`websocket_server_profile` - (Optional) x-displayName: "WebSocket Profile Server". See [ref](#ref) below for details.

### Virtual Server Type Tcp

x-displayName: "TCP".

`tcp_client_profile` - (Optional) x-displayName: "Protocol Profile (Client)". See [ref](#ref) below for details.

`tcp_server_profile` - (Optional) x-displayName: "Protocol Profile (Server)". See [ref](#ref) below for details.

### Virtual Server Type Udp

x-displayName: "UDP".

`tcp_client_profile` - (Optional) x-displayName: "Protocol Profile (Client)". See [ref](#ref) below for details.

`tcp_server_profile` - (Optional) x-displayName: "Protocol Profile (Server)". See [ref](#ref) below for details.

Attribute Reference
-------------------

*   `id` - This is the id of the configured application_profiles.
