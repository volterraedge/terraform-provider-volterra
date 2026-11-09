---

page_title: "Volterra: udp_profile"

description: "The udp_profile allows CRUD of Udp Profile resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_udp_profile
=============================

The Udp Profile allows CRUD of Udp Profile resource on Volterra SaaS

~> **Note:** Please refer to [Udp Profile API docs](https://docs.cloud.f5.com/docs-v2/api/udp-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_udp_profile" "example" {
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

`allow_no_payload` - (Optional) Provides the ability to allow the passage of datagrams that contain header information, but no essential data. The default value is disabled. (`String`).

`app_service` - (Optional) The application service to which the object belongs. (`String`).

`buffer_max_bytes` - (Optional) Specifies ingress buffer byte limit. (`Int`).

`buffer_max_packets` - (Optional) Specifies ingress buffer packet limit. (`Int`).

`command` - (Optional) x-displayName: "Command" (`String`).

`datagram_load_balancing` - (Optional) Provides the ability to load balance UDP datagram by datagram. The default value is disabled. (`String`).

`idle_timeout` - (Optional) Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 60 seconds. (`Int`).

`ip_df_mode` - (Optional) Describe the Don't Fragment (DF) bit setting in the outgoing UDP packet. (`String`).

`ip_tos_to_client` - (Optional) Specifies the Type of Service level that the traffic management system assigns to UDP packets when sending them to clients. (`Int`).

`ip_ttl_mode` - (Optional) Describe the outgoing UDP packet TTL mode. (`String`).

`ip_ttl_v4` - (Optional) x-displayName: "Ip Ttl V4" (`Int`).

`ip_ttl_v6` - (Optional) x-displayName: "Ip Ttl V6" (`Int`).

`link_qos_to_client` - (Optional) Specifies the Quality of Service level that the system assigns to UDP packets when sending them to clients. (`Int`).

`no_checksum` - (Optional) Enables or disables checksum processing. Note that if the datagram is IPv6, the system always performs checksum processing. The default value is disabled. (`String`).

`proxy_mss` - (Optional) Make the system use the same max segment size on both ends. (`String`).

`send_buffer_size` - (Optional) Specifies egress send buffer byte limit. The default value is 655350. The range is from 536 to 16777215. (`Int`).

`user_spec` - (Optional) User specified properties. (`List of String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured udp_profile.
