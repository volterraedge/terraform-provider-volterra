---

page_title: "Volterra: fastl4_profile"

description: "The fastl4_profile allows CRUD of Fastl4 Profile resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_fastl4_profile
================================

The Fastl4 Profile allows CRUD of Fastl4 Profile resource on Volterra SaaS

~> **Note:** Please refer to [Fastl4 Profile API docs](https://docs.cloud.f5.com/docs-v2/api/fastl4-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_fastl4_profile" "example" {
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

`app_service` - (Optional) The application service to which the object belongs. (`String`).

`client_timeout` - (Optional) Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take. (`Int`).

`explicit_flow_migration` - (Optional) Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware. (`String`).

`full_path` - (Optional) x-displayName: "Full Path" (`String`).(Deprecated)

`idle_timeout` - (Optional) Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable PVA.Scrub time in msec for it to work properly.The default value is 300 seconds. (`Int`).

`ip_df_mode` - (Optional) Describe the Don't Fragment (DF) bit setting in the outgoing packet's IP Header. (`String`).

`ip_tos_to_client` - (Optional) Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify. (`Int`).

`ip_tos_to_server` - (Optional) Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify. (`Int`).

`ip_ttl_mode` - (Optional) Describe the outgoing packet's IP Header TTL value modes. (`String`).

`ip_ttl_v4` - (Optional) Specifies the outgoing IPV4 Header TTL value for ip-ttl-mode SET (`Int`).

`ip_ttl_v6` - (Optional) Specifies the outgoing IPV6 Header TTL value for ip-ttl-mode SET (`Int`).

`keep_alive_interval` - (Optional) Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds). (`Int`).

`late_binding` - (Optional) Enables or disables late binding to control L7 to FastL4 Flow migration. The default value is disabled. (`String`).

`link_qos_to_client` - (Optional) Specifies a Link QoS (VLAN priority) number for the client side. This option specifies the Quality of Service level that the system assigns to packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify. (`Int`).

`link_qos_to_server` - (Optional) Specifies a Link QoS (VLAN priority) number for the server side. This option specifies the Quality of Service level that the system assigns to packets when sending them to servers. The default value is 65535, which indicates, do not modify. (`Int`).

`loose_close` - (Optional) Specifies that the system closes a loosely-initiated connection when the system receives the first FIN packet from either the client or the server. The default value is disabled. (`String`).

`loose_initialization` - (Optional) Specifies that the system initializes a connection when it receives any TCP packet, rather than requiring a SYN packet for connection initiation. The default value is disabled. (`String`).

`mss_override` - (Optional) Specifies a maximum segment size (MSS) override for server-side connections. Note that this is also the MSS advertised to a client when a client first connects. The default value is 0 (zero), which disables this option. You can specify an integer from 256 to 9162. (`Int`).

`other_pva_clientpkts_threshold` - (Optional) Specifies the number of client packets before ePVA hardware offloading occurs for stateless protocol traffic. The valid value is 0~255. The default value is 2. (`Int`).

`other_pva_offload_direction` - (Optional) For stateless protocol traffic only, specifies which side of the traffic can ePVA perform hardware offload for. Bidirectional implies both side is permitted to offload if threshold exceeds. Client-to-server-only implies only the traffic from client to server is allowed to be offloaded. Even if the traffic from server to client exceeds the threshold, it will not be offloaded. Vice versa, server-to-client-only implies only the traffic from server to client is allowed to be offloaded. (`String`).

`other_pva_serverpkts_threshold` - (Optional) Specifies the number of server packets before ePVA hardware offloading occurs for stateless protocol traffic. The valid value is 0~255. The default value is 1. (`Int`).

`other_pva_whento_offload` - (Optional) Specifies when the ePVA performs hardware offload for stateless protocol traffic. After-packets-per-direction implies the client and server traffic is offloaded independently after exceeding their own thresholds. After-packets-both-direction implies both client and server traffic thresholds need to be exceeded, then can both sides get offloaded. (`String`).

`priority_to_client` - (Optional) Specifies internal packet priority for the client side. This option specifies the internal packet priority that the system assigns to packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify. (`Int`).

`priority_to_server` - (Optional) Specifies internal packet priority for the server side. This option specifies the internal packet priority that the system assigns to packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify. (`Int`).

`pva_acceleration` - (Optional) Specifies the Packet Velocity ASIC acceleration policy. (`String`).

`pva_dynamic_client_packets` - (Optional) Specifies the number of client packets before dynamic ePVA hardware re-offloading occurs. The valid value is 0~10. (`Int`).

`pva_dynamic_server_packets` - (Optional) Specifies the number of server packets before dynamic ePVA hardware re-offloading occurs. The valid value is 0~10. (`Int`).

`pva_flow_aging` - (Optional) Specifies if enabling automatic epva flow aging when flow becomes idle or inactive for a period of time. use with care, (`String`).

`pva_flow_evict` - (Optional) Specifies if epva flow can be evicted upon hash collision with a new flow snoop learn request. use with care. (`String`).

`pva_offload_dynamic` - (Optional) Specifies whether PVA flow dynamic offloading is enabled or not. (`String`).

`pva_offload_dynamic_priority` - (Optional) Specifies if dynamic adjustment of epva offload flow priority is turned on or not. Default value is disabled. (`String`).

`pva_offload_initial_priority` - (Optional) Specifies the initial epva offload priority of a flow. Priority can be low, medium or high. The default value is medium (`String`).

`reassemble_fragments` - (Optional) Specifies whether to reassemble fragments. The default value is disabled. (`String`).

`receive_window_size` - (Optional) Specifies the size of the receive window, in bytes. The minimum and default value is 65535 bytes without scale. (`Int`).

`reset_on_client_fin` - (Optional) Specifies whether to reset connections when a TCP FIN is received from the client. The default value is disabled. (`String`).

`reset_on_timeout` - (Optional) Specifies whether you want to reset connections on timeout. The default value is enabled.The default value is enabled. (`String`).

`rtt_from_client` - (Optional) Enables or disables the TCP timestamp options to measure the round trip time to the client. The default value is disabled. (`String`).

`rtt_from_server` - (Optional) Enables or disables the TCP timestamp options to measure the round trip time to the server. The default value is disabled. (`String`).

`server_sack` - (Optional) Specifies whether to support server sack option in cookie response by default. The default is disabled. (`String`).

`server_timestamp` - (Optional) Specifies whether to support server timestamp option in cookie response by default. The default is disabled. (`String`).

`sub_path` - (Optional) x-displayName: "Sub Path" (`String`).(Deprecated)

`syn_cookie_dsr_flow_reset_by` - (Optional) Specifies how TCP SYN Flood is handled when syn-cookie-whitelist is enabled and the attack is detected in Direct Server Return(DSR) mode. (`String`).

`syn_cookie_enable` - (Optional) Specifies whether or not to use SYN Cookie. The default is enabled. (`String`).

`syn_cookie_mss` - (Optional) Specifies a maximum segment size (MSS) for server-side connections when SYN cookie is enabled. Note that this is also the MSS advertised to a client when a client first connects. The default value is 0 (zero), which disables this option. You can specify an integer from 256 to 9162. (`Int`).

`syn_cookie_whitelist` - (Optional) This option to enable SYN Cookie whitelist. The default is false. (`String`).

`tcp_close_timeout` - (Optional) Specifies an TCP close timeout in seconds. The default value is 5 seconds. (`Int`).

`tcp_generate_isn` - (Optional) Specifies whether you want to generate TCP sequence numbers on all SYNs that conform with RFC1948, and allow timestamp recycling. The default value is disabled. (`String`).

`tcp_handshake_timeout` - (Optional) Specifies a TCP handshake timeout in seconds. The default value is 5 seconds. (`Int`).

`tcp_pva_offload_direction` - (Optional) For tcp protocol traffic only, specifies which side of the traffic can ePVA perform hardware offload for. Bidirectional implies both side is permitted to offload if threshold exceeds. Client-to-server-only implies only the traffic from client to server is allowed to be offloaded. Even if the traffic from server to client exceeds the threshold, it will not be offloaded. Vice versa, server-to-client-only implies only the traffic from server to client is allowed to be offloaded. (`String`).

`tcp_pva_whento_offload` - (Optional) Specifies at what stage the ePVA performs hardware offload for TCP traffic. Embryonic implies at TCP SYN packet; establish implies after TCP 3WAY handshaking. (`String`).

`tcp_strip_sack` - (Optional) Specifies whether you want to block the TCP SackOK option from passing to server on an initiating SYN. The default value is disabled. (`String`).

`tcp_time_wait_timeout` - (Optional) Specifies an TCP time_wait timeout in milliseconds. The default value is 0 milliseconds. In another word, by default there is no time_wait period for fastl4 connflow. (`Int`).

`tcp_timestamp_mode` - (Optional) Specifies how you want to handle the TCP timestamp. The default value is preserve. (`String`).

`tcp_wscale_mode` - (Optional) Specifies how you want to handle the TCP window scale. The default value is preserve, which preserves TCP window scale. (`String`).

`timeout_recovery` - (Optional) Specifies late binding timeout recovery mode. This setting specifies the action to take when the client timeout expires. The default is to disconnect the connection. It can be configured to fallback to the normal FastL4 load-balancing for selecting the back-end servers. (`String`).

`user_spec` - (Optional) User specified properties. (`List of String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured fastl4_profile.
