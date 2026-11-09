---

page_title: "Volterra: tcp_profile"

description: "The tcp_profile allows CRUD of Tcp Profile resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_tcp_profile
=============================

The Tcp Profile allows CRUD of Tcp Profile resource on Volterra SaaS

~> **Note:** Please refer to [Tcp Profile API docs](https://docs.cloud.f5.com/docs-v2/api/tcp-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_tcp_profile" "example" {
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

`abc` - (Optional) When enabled, increases the congestion window by basing the increase amount on the number of previously unacknowledged bytes that each ACK covers. The default value is enabled. (`String`).

`ack_on_push` - (Optional) When enabled, significantly improves performance to Windows and MacOS peers who are writing out on a very small send buffer. The default value is enabled. (`String`).

`app_service` - (Optional) The application service to which the object belongs. (`String`).

`auto_proxy_buffer_size` - (Optional) When enabled, specifies that the system uses the network measurements to set the optimal proxy buffer size. The default value is disabled. (`String`).

`auto_receive_window_size` - (Optional) When enabled, specifies that the system uses the network measurements to set the optimal receive window size. The default value is disabled. (`String`).

`auto_send_buffer_size` - (Optional) When enabled, specifies that the system uses the network measurements to set the optimal send buffer size. The default value is disabled. (`String`).

`close_wait_timeout` - (Optional) Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds. (`Int`).

`cmetrics_cache` - (Optional) Specifies, when enabled, that the system uses a cache for storing congestion metrics. The default value is enabled. (`String`).

`cmetrics_cache_timeout` - (Optional) Specifies the time, in seconds, for which entries in the congestion metrics cache are valid. The default value is 0, which defers to the sys db variable route.metrics.timeout. (`Int`).

`command` - (Optional) x-displayName: "Command" (`String`).

`congestion_control` - (Optional) Specifies the algorithm to use to share network resources among competing users to reduce congestion. The default value is high-speed. (`String`).

`deferred_accept` - (Optional) Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled. (`String`).

`delay_window_control` - (Optional) Specifies whether the system uses an estimate of queuing delay as a measure of congestion, in addition to the normal loss-based control, to control the amount of data sent. The default is disabled. (`String`).

`delayed_acks` - (Optional) When enabled, the traffic management system allows coalescing of multiple ACK responses. The default value is enabled. (`String`).

`dsack` - (Optional) When enabled, specifies the use of the Selective ACKs (SACK) option to acknowledge duplicate segments. The default value is disabled. (`String`).

`early_retransmit` - (Optional) When enabled, specifies that the system uses early fast retransmits (as specified in RFC 5827) to reduce the recovery time for connections that are receive-buffer or user-data limited. The default value is enabled. (`String`).

`ecn` - (Optional) Specifies, when enabled, that the system uses the TCP flags CWR and ECE to notify its peer of congestion and congestion counter-measures. The default value is enabled. (`String`).

`enhanced_loss_recovery` - (Optional) When enabled, specifies that the system uses enhanced loss recovery to recover from random packet losses more effectively. The default value is enabled. (`String`).

`extended_analytics` - (Optional) When enabled, specifies that the system collects extended analytics data for TCP connections. The default value is enabled. (`Bool`).

`fast_open` - (Optional) When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet. This option has no effect on server-side TCP profiles. The default value is enabled. (`String`).

`fast_open_cookie_expiration` - (Optional) Seconds for which a Fast Open Cookie provided by the BIG-IP is valid for incoming SYN packets. The default value is 21,600 seconds (6 hours). The range is from 0 (meaning use the default) to 1000000. (`Int`).

`fin_wait_2_timeout` - (Optional) Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). (`Int`).

`fin_wait_timeout` - (Optional) Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite. (`Int`).

`full_path` - (Optional) x-displayName: "Full Path" (`String`).

`hardware_syn_cookie` - (Optional) *IMPORTANT* This command has been deprecated (as of 13.0.0). Specifies whether or not to use hardware SYN Cookie when cross system limit. The default is enabled. (`String`).

`idle_timeout` - (Optional) Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds. (`Int`).

`init_cwnd` - (Optional) Specifies the initial congestion window size for connections to this destination. The actual window size is this value multiplied by the MSS (Maximal Segment Size) for the same connection. The default value is 10. The range is from 0 to 64. (`Int`).

`init_rwnd` - (Optional) Specifies the initial receive window size for connections to this destination. The actual window size is this value multiplied by the MSS (Maximal Segment Size) for the same connection. The default value is 10. The range is from 0 to 64. (`Int`).

`ip_df_mode` - (Optional) Describe the Don't Fragment (DF) bit setting in the outgoing packet's IP Header. The default setting is PMTU. (`String`).

`ip_tos_to_client` - (Optional) Specifies the Type of Service level that the traffic management system assigns to TCP packets when sending them to clients. (`String`).

`ip_ttl_mode` - (Optional) Describe the outgoing packet's IP Header TTL value modes. (`String`).

`ip_ttl_v4` - (Optional) Specifies the outgoing IPV4 Header TTL value for ip-ttl-mode SET (`Int`).

`ip_ttl_v6` - (Optional) Specifies the outgoing IPV6 Header TTL value for ip-ttl-mode SET (`Int`).

`keep_alive_interval` - (Optional) Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds. (`Int`).

`limited_transmit` - (Optional) When enabled, specifies that the system uses limited transmit recovery revisions for fast retransmits (as specified in RFC 3042) to reduce the recovery time for connections on a lossy network. The default value is enabled. (`String`).

`link_qos_to_client` - (Optional) Specifies the Quality of Service level that the system assigns to TCP packets when sending them to clients. The default value is 0 (zero). (`String`).

`max_retrans` - (Optional) Specifies the maximum number of retransmissions of data segments that the system allows. The default value is 8. (`Int`).

`max_segment_size` - (Optional) Specifies the largest amount of data that the system can receive in a single TCP segment, not including the TCP and IP headers. If the value is 0 (zero), the system calculates the value from the MTU. The default value is 1460 bytes. (`Int`).

`md5_signature` - (Optional) Specifies, when enabled, that the system uses RFC2385 TCP-MD5 signatures to protect TCP traffic against intermediate tampering. The default value is disabled. (`String`).

`md5_signature_passphrase` - (Optional) Specifies, when enabled, a plaintext passphrase which may be between 1 and 80 characters in length, and is used in a shared-secret scheme to implement the spoof-prevention parts of RFC2385. The default value is none. (`String`).

`minimum_rto` - (Optional) Specifies the minimum TCP retransmission timeout in milliseconds. The default value is 1000 milliseconds. (`Int`).

`mptcp` - (Optional) When enabled all incoming flows are handled by the MPTCP stack, allowing for support of multipath-enabled connections. When passthrough MPTCP connections are not terminated by this virtual. (`String`).

`mptcp_csum` - (Optional) If enabled, checksums are supported by this MPTCP-enabled device. (`String`).

`mptcp_csum_verify` - (Optional) If enabled, incoming checksums are verified, and checksum failure causes connection abort. (`String`).

`mptcp_debug` - (Optional) This option is DEPRECATED v12.0.0 onwards and is maintained here for backward compatibility reasons. When enabled, the debug output and statistics are available. (`String`).

`mptcp_fallback` - (Optional) Specifies the algorithm for fallback. The default value is reset. (`String`).

`mptcp_fastjoin` - (Optional) When enabled, permits FAST join, allowing data to be sent on the MP_JOIN SYN, which can allow a server response to occur in parallel with the join. (`String`).

`mptcp_idle_timeout` - (Optional) Number of seconds without traffic before a MPTCP connection is eligible for deletion. The range is from 1 to 2147483647 and the default is 300. (`Int`).

`mptcp_join_max` - (Optional) Specifies the maximum number of simultaneous join attempts on a given flow. The default value is 5. The range is from 1 to 20. (`Int`).

`mptcp_makeafterbreak` - (Optional) When enabled, permit after break functionality, allowing for long-lived MPTCP sessions. (`String`).

`mptcp_nojoindssack` - (Optional) When enabled, no DSS option is sent on the JOIN ACK. (`String`).

`mptcp_rtomax` - (Optional) Specifies the number of RTOs before declaring subflow dead. The default value is 5. The range is from 1 to 12. (`Int`).

`mptcp_rxmitmin` - (Optional) Specifies the minimum value of the retransmission timer for these MPTCP flows. The default value is 1000 msec. The range is from 200 to 5000 msec. (`Int`).

`mptcp_subflowmax` - (Optional) Specifies the maximum number of subflows for a single flow. The range is 0 to 60 where 0 = unlimited. (`Int`).

`mptcp_timeout` - (Optional) Specifies the timeout value to discard long-lived sessions that do not have an active flow. The default value is 3600 seconds. The range is from 60 to 36000 seconds. (`Int`).

`nagle` - (Optional) Specifies, when enabled, that the system applies Nagle's algorithm to reduce the number of short segments on the network. The default value is disabled. Note that for interactive protocols such as Telnet, rlogin, or SSH, F5 recommends disabling this setting on high-latency networks, to improve application responsiveness. When auto, the use of Nagle's algorithm is decided based on network conditions. (`String`).

`pkt_loss_ignore_burst` - (Optional) Specifies the probability of performing congestion control when multiple packets in a row are lost even if the pkt-loss-ignore-rate was not exceeded. Valid values are 0 to 32. The default is 0, meaning that the system should perform congestion control if any packets are lost. Higher values decrease the chance of performing congestion control. (`Int`).

`pkt_loss_ignore_rate` - (Optional) Specifies the threshold of packets lost per million at which the system should perform congestion control. Valid values for n are 0 to 1,000,000. The default is 0, meaning the system should perform congestion control if any packet loss occurs. If you set the ignore rate to 10 and packet loss for a TCP connection is greater than 10 per million, congestion control occurs. (`Int`).

`proxy_buffer_high` - (Optional) Specifies the highest level at which the receive window is closed. The default value is 131072. (`Int`).

`proxy_buffer_low` - (Optional) Specifies the lowest level at which the receive window is closed. The default value is 98304. (`Int`).

`proxy_mss` - (Optional) Specifies, when enabled, that the system advertises the same TCP maximum segment size to the server as was negotiated with the client. The setting is ignored when MRF routing (e.g., httprouter, siprouter, diameterrouter, mqttrouter, messagerouter) is used. The default value is enabled. (`String`).

`proxy_options` - (Optional) Specifies, when enabled, that the system advertises an option, such as a time-stamp to the server only if it was negotiated with the client. The setting is ignored when MRF routing (e.g., httprouter, siprouter, diameterrouter, mqttrouter, messagerouter) is used. The default value is disabled. (`String`).

`push_flag` - (Optional) When default, specifies that the system sets PUSH flag when sending the last segment in the send buffer. When none, specifies that the system never sets PUSH flag for TCP packets. When one, specifies that the system sets one PUSH flag for the FIN segment. When auto, specifies that the system sets PUSH flag based on the application/network conditions. The default value is default. (`String`).

`rate_pace` - (Optional) When enabled, the system will rate pace TCP data transmission. The default value is enabled. (`String`).

`rate_pace_max_rate` - (Optional) If not 0, the maximum rate in bytes per second that TCP connections will be paced to. (`Int`).

`receive_window_size` - (Optional) Specifies the size of the receive window, in bytes. The default value is 65535 bytes. (`Int`).

`reset_on_timeout` - (Optional) Specifies whether to reset connections on timeout. The default value is enabled. (`String`).

`rexmt_thresh` - (Optional) Specifies the number of duplicate ACKs (retransmit threshold) to start fast recovery. Higher values decrease the likelihood of performing fast recovery in a network with high packet reordering. The default value is 3. The range is from 3 to 255. (`Int`).

`selective_acks` - (Optional) Specifies, when enabled, that the system negotiates RFC2018-compliant Selective Acknowledgments with peers. The default value is enabled. (`String`).

`selective_nack` - (Optional) Specifies whether Selective Negative Acknowledgment is enabled or not. The default value is disabled. (`String`).

`send_buffer_size` - (Optional) Specifies the size of the buffer, in bytes. The default value is 131072 bytes. (`Int`).

`slow_start` - (Optional) Specifies, when enabled, that the system uses larger initial window sizes (as specified in RFC 3390) to help reduce round trip times. The default value is enabled. (`String`).

`syn_cookie_enable` - (Optional) Specifies whether or not to use SYN Cookie. The default is enabled. (`String`).

`syn_cookie_whitelist` - (Optional) This option to enable SYN Cookie WhiteList. The default is disabled. (`String`).

`syn_max_retrans` - (Optional) Specifies the maximum number of retransmissions of SYN segments that the system allows. The default value is 3. (`Int`).

`syn_rto_base` - (Optional) Specifies the initial RTO (Retransmission TimeOut) base multiplier for SYN retransmission, in milliseconds. This value is modified by the exponential backoff table to select the interval for subsequent retransmissions. The default value is 3000. (`Int`).

`tail_loss_probe` - (Optional) When enabled, specifies that the system uses tail loss probe to reduce the number of retransmission timeouts. The default value is enabled. (`String`).

`tcp_options` - (Optional) Specifies the option numbers that will be accessible from iRules (TCP::option) for the flow. The format of each entry should be, "{<option number> <first|last>} {<option number> <first|last>}". The keyword "first" means the system records the option the first time it is occurs (after and including the ACK of the three way handshake), while "last" means the system updates the available value every time it occurs. (`String`).

`time_wait_recycle` - (Optional) Specifies whether the system recycles the connection when a SYN packet is received in a TIME-WAIT state. The default value is enabled. (`String`).

`time_wait_timeout` - (Optional) Specifies the number of milliseconds that a connection is in the TIME-WAIT state before closing. The default value is 2000 milliseconds. The range is from 0 to 600,000 (10 minutes). (`String`).

`timestamps` - (Optional) Specifies, when enabled, that the system uses the timestamps extension for TCP (as specified in RFC 1323) to enhance high-speed network performance. The default value is enabled. (`String`).

`verified_accept` - (Optional) When enabled, a SYN-ACK will be sent only if the server port is open. Not compatible with iRules. The default is disabled. (`String`).

`zero_window_timeout` - (Optional) Specifies the timeout in milliseconds for terminating a connection with an effective zero length TCP transmit window. The timeout starts when the peer advertises a zero length TCP window or when enough data has been sent to fill the previously advertised window. The timer is canceled when a non-zero length window is received. The default is 20000 milliseconds. (`Int`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured tcp_profile.
