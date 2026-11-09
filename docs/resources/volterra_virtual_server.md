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

  access_profile {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }

  domains = ["www.foo.com"]

  virtual_addresses {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
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

`access_profile` - (Optional) Specified an access policy that determines the authentication rules and access controls applied to user sessions for this virtual server. See [ref](#ref) below for details.

`connection_limit_options` - (Optional) x-displayName: "Connection Limit Options". See [Connection Limit Options ](#connection-limit-options) below for details.

###### One of the arguments from this list "default_pool, default_pool_none" can be set

`default_pool` - (Optional) x-displayName: "Select Pool". See [ref](#ref) below for details.

`default_pool_none` - (Optional) x-displayName: "None" (`Bool`).

`domains` - (Required) Domains also indicate the list of names for which DNS resolution will be automatically resolved to IP addresses by the system. (`List of String`).

`immediate_action_on_service_down` - (Optional) - Drop: Specifies that the system drops the connections when the virtual server is reported Offline or Unavailable.. See [Immediate Action On Service Down ](#immediate-action-on-service-down) below for details.

`irules` - (Optional) Specifies the iRules you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.. See [ref](#ref) below for details.(Deprecated)

`traffic_policies` - (Optional) Specifies the Traffic Policy you want use for this virtual server.. See [ref](#ref) below for details.

`virtual_addresses` - (Required) Specify the virtual address objects to be associated with this virtual server. See [ref](#ref) below for details.

###### One of the arguments from this list "http3, https" can be set

`http3` - (Optional) Selection provides configuration for HTTP/3 services. See [Virtual Server Type Http3 ](#virtual-server-type-http3) below for details.

`https` - (Optional) Selection provides configuration for HTTP and HTTP/2 services with TLS configuration. See [Virtual Server Type Https ](#virtual-server-type-https) below for details.

`vs_score` - (Optional) Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value to load balance traffic in a proportional manner. The default is 0, meaning that no additional metric is applied for the virtual server. (`Int`).

### Connection Limit Options

x-displayName: "Connection Limit Options".

`connection_limit` - (Optional) Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this to 0 turns off connection limits. The default is 0. (`Int`).

`connection_rate_limit` - (Optional) Specifies the maximum number of connections-per-second allowed for a virtual server. When the number of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where connection requests flood a virtual server. Setting this to 0 turns off connection limits. The default is 0. (`Int`).

`connection_rate_limit_mode` - (Optional) - Per Source and Destination Address (All Rate Limiting Virtual Servers): Applies rate limiting based on the specified source and destination addresses for all virtual servers that have rate limits specified.. See [Connection Limit Options Connection Rate Limit Mode ](#connection-limit-options-connection-rate-limit-mode) below for details.

### Immediate Action On Service Down

-	Drop: Specifies that the system drops the connections when the virtual server is reported Offline or Unavailable..

###### One of the arguments from this list "immediate_action_on_service_down_drop, immediate_action_on_service_down_none, immediate_action_on_service_down_reset" can be set

`immediate_action_on_service_down_drop` - (Optional) x-displayName: "Drop". See [Immediate Action On Service Down Choice Immediate Action On Service Down Drop ](#immediate-action-on-service-down-choice-immediate-action-on-service-down-drop) below for details.

`immediate_action_on_service_down_none` - (Optional) x-displayName: "None". See [Immediate Action On Service Down Choice Immediate Action On Service Down None ](#immediate-action-on-service-down-choice-immediate-action-on-service-down-none) below for details.

`immediate_action_on_service_down_reset` - (Optional) x-displayName: "Reset". See [Immediate Action On Service Down Choice Immediate Action On Service Down Reset ](#immediate-action-on-service-down-choice-immediate-action-on-service-down-reset) below for details.

### Connection Limit Options Connection Rate Limit Mode

-	Per Source and Destination Address (All Rate Limiting Virtual Servers): Applies rate limiting based on the specified source and destination addresses for all virtual servers that have rate limits specified..

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

### Http3 Services

x-displayName: "Services".

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Http Server Profile Choice Http Server Profile Same As Client

x-displayName: "Use Client Profile".

### Https Services

x-displayName: "Services".

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Immediate Action On Service Down Choice Immediate Action On Service Down Drop

x-displayName: "Drop".

### Immediate Action On Service Down Choice Immediate Action On Service Down None

x-displayName: "None".

### Immediate Action On Service Down Choice Immediate Action On Service Down Reset

x-displayName: "Reset".

### Protocol Server Profile Choice Protocol Server Profile Same As Client

x-displayName: "Use Client Profile".

### Quic Profile Choice Quic Client Profile None

x-displayName: "None".

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Server App Type Choice Server App Type Default

x-displayName: "TCP".

###### One of the arguments from this list "http_server_profile, http_server_profile_same_as_client" can be set

`http_server_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`http_server_profile_same_as_client` - (Optional) x-displayName: "Use Client Profile" (`Bool`).

`ssl_server_profiles` - (Optional) Specifies the SSL profile for managing server-side SSL traffic. The list contains entries for each already defined server-side SSL profile.. See [ref](#ref) below for details.

`tcp_server_profile` - (Required) Specifies that the selected profile is a server-side profile. Options are: (Use Client Profile), and entries for each already defined server protocol profile. The default is (Use Client Profile).. See [ref](#ref) below for details.

### Server App Type Choice Server App Type Same As Client

x-displayName: "Same As Client".

###### One of the arguments from this list "http_server_profile, http_server_profile_same_as_client" can be set

`http_server_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`http_server_profile_same_as_client` - (Optional) x-displayName: "Use Client Profile" (`Bool`).

###### One of the arguments from this list "protocol_server_profile, protocol_server_profile_same_as_client" must be set

`protocol_server_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`protocol_server_profile_same_as_client` - (Optional) x-displayName: "Use Client Profile" (`Bool`).

`ssl_server_profiles` - (Optional) Specifies the SSL profile for managing server-side SSL traffic. The list contains entries for each already defined server-side SSL profile.. See [ref](#ref) below for details.

### Virtual Server Type Http3

Selection provides configuration for HTTP/3 services.

`http3_client_profile` - (Optional) Specifies that the selected profile is a client-side profile. The list contains entries for each already defined client protocol profile.. See [ref](#ref) below for details.

`http_client_profile` - (Optional) Specifies that the selected profile is a client-side profile. The list contains entries for each already defined client protocol profile.. See [ref](#ref) below for details.

`protocol_client_profile` - (Required) Specifies that the selected profile is a client-side profile. The list contains entries for each already defined client protocol profile.. See [ref](#ref) below for details.

###### One of the arguments from this list "quic_client_profile, quic_client_profile_none" can be set

`quic_client_profile` - (Optional) x-displayName: "Select Profile". See [ref](#ref) below for details.

`quic_client_profile_none` - (Optional) x-displayName: "None" (`Bool`).

###### One of the arguments from this list "server_app_type_default" can be set

`server_app_type_default` - (Optional) x-displayName: "TCP". See [Server App Type Choice Server App Type Default ](#server-app-type-choice-server-app-type-default) below for details.

`services` - (Optional) x-displayName: "Services". See [Http3 Services ](#http3-services) below for details.

`ssl_client_profiles` - (Optional) Specifies the SSL profile for managing client-side SSL traffic. The list contains entries for each already defined client-side SSL profile.. See [ref](#ref) below for details.

### Virtual Server Type Https

Selection provides configuration for HTTP and HTTP/2 services with TLS configuration.

`http_client_profile` - (Required) Specifies that the selected profile is a client-side HTTP profile. The list contains entries for each already defined client protocol profile.. See [ref](#ref) below for details.

`protocol_client_profile` - (Required) Specifies that the selected profile is a client-side profile. The list contains entries for each already defined client protocol profile.. See [ref](#ref) below for details.

###### One of the arguments from this list "server_app_type_same_as_client" can be set

`server_app_type_same_as_client` - (Optional) x-displayName: "Same As Client". See [Server App Type Choice Server App Type Same As Client ](#server-app-type-choice-server-app-type-same-as-client) below for details.

`services` - (Optional) x-displayName: "Services". See [Https Services ](#https-services) below for details.

`ssl_client_profiles` - (Required) Specifies the SSL profile for managing client-side SSL traffic. The list contains entries for each already defined client-side SSL profile.. See [ref](#ref) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured virtual_server.
