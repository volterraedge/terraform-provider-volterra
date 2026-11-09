---

page_title: "Volterra: traffic_policy"

description: "The traffic_policy allows CRUD of Traffic Policy resource on Volterra SaaS"
-----------------------------------------------------------------------------------------

Resource volterra_traffic_policy
================================

The Traffic Policy allows CRUD of Traffic Policy resource on Volterra SaaS

~> **Note:** Please refer to [Traffic Policy API docs](https://docs.cloud.f5.com/docs-v2/api/traffic-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_traffic_policy" "example" {
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

`rules` - (Optional) x-displayName: "Rules". See [Rules ](#rules) below for details.

`strategy` - (Optional) x-displayName: "Strategy" (`String`).

### Rules

x-displayName: "Rules".

`actions` - (Optional) x-displayName: "Actions". See [Rules Actions ](#rules-actions) below for details.

`conditions` - (Optional) x-displayName: "Conditions". See [Rules Conditions ](#rules-conditions) below for details.

`description` - (Optional) x-displayName: "Description" (`String`).

`name` - (Required) x-required (`String`).

### Action Choice Collect Data

x-displayName: "Collect Data".

`event` - (Required) x-required (`String`).

### Action Choice Disable

x-displayName: "Disable".

`event` - (Required) x-required (`String`).

`feature` - (Required) x-required (`String`).

### Action Choice Enable

x-displayName: "Enable".

`auto_mode` - (Optional) x-displayName: "TCP Auto Nagle" (`String`).

`event` - (Required) x-required (`String`).

`feature` - (Required) x-required (`String`).

### Action Choice Forward Traffic

x-displayName: "Forward Traffic".

`event` - (Required) x-required (`String`).

###### One of the arguments from this list "pool, virtual_server" must be set

`pool` - (Optional) x-displayName: "Pool". See [Target Pool ](#target-pool) below for details.

`virtual_server` - (Optional) x-displayName: "Virtual Server". See [Target Virtual Server ](#target-virtual-server) below for details.

### Action Choice Insert

x-displayName: "Insert".

`domain` - (Optional) Value for the domain attribute of a cookie in Set-Cookie header. Tcl command substitutions are allowed for this field. (`String`).

`event` - (Required) x-required (`String`).

`name` - (Optional) x-displayName: "Name" (`String`).

`path` - (Optional) Value for the path attribute of a cookie in Set-Cookie header. Tcl command substitutions are allowed for this field. (`String`).

`target` - (Required) x-required (`String`).

`value` - (Optional) x-displayName: "Value" (`String`).

### Action Choice Log

x-displayName: "Log".

`event` - (Required) x-required (`String`).

`facility` - (Optional) x-displayName: "Facility" (`String`).

`message` - (Required) String message to place into target. A Tcl expression can also be used. (`String`).

`priority` - (Optional) x-displayName: "Priority" (`String`).

###### One of the arguments from this list "disable_remote_server, enable_remote_server" can be set

`disable_remote_server` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_remote_server` - (Optional) x-displayName: "Enable". See [Remote Server Enable Remote Server ](#remote-server-enable-remote-server) below for details.

### Action Choice Persist Session

x-displayName: "Persist Session".

`event` - (Required) x-required (`String`).

###### One of the arguments from this list "carp, cookie_hash, cookie_insert, cookie_passive, cookie_rewrite, destination_address, hash, source_address, universal" must be set

`carp` - (Optional) x-displayName: "Carp". See [Method Carp ](#method-carp) below for details.

`cookie_hash` - (Optional) x-displayName: "Cookie Hash". See [Method Cookie Hash ](#method-cookie-hash) below for details.

`cookie_insert` - (Optional) x-displayName: "Cookie Insert". See [Method Cookie Insert ](#method-cookie-insert) below for details.

`cookie_passive` - (Optional) x-displayName: "Cookie Passive". See [Method Cookie Passive ](#method-cookie-passive) below for details.

`cookie_rewrite` - (Optional) x-displayName: "Cookie Rewrite". See [Method Cookie Rewrite ](#method-cookie-rewrite) below for details.

`destination_address` - (Optional) x-displayName: "Destination Address". See [Method Destination Address ](#method-destination-address) below for details.

`hash` - (Optional) x-displayName: "Hash". See [Method Hash ](#method-hash) below for details.

`source_address` - (Optional) x-displayName: "Source Address". See [Method Source Address ](#method-source-address) below for details.

`universal` - (Optional) x-displayName: "Universal". See [Method Universal ](#method-universal) below for details.

### Action Choice Redirect

x-displayName: "Redirect".

`code` - (Optional) Optional HTTP response code for redirect. (`Int`).

`event` - (Required) x-required (`String`).

`location` - (Optional) The new URL for which a redirect response will be sent. A Tcl command substitution can be used for this field. (`String`).

### Action Choice Remove

x-displayName: "Remove".

`event` - (Required) x-required (`String`).

`name` - (Optional) x-displayName: "Name" (`String`).

`target` - (Required) x-required (`String`).

### Action Choice Replace

x-displayName: "Replace".

`event` - (Required) x-required (`String`).

###### One of the arguments from this list "http_connect, http_header, http_host, http_referer, http_uri" must be set

`http_connect` - (Optional) x-displayName: "HTTP Connect". See [Target Http Connect ](#target-http-connect) below for details.

`http_header` - (Optional) x-displayName: "HTTP Header". See [Target Http Header ](#target-http-header) below for details.

`http_host` - (Optional) x-displayName: "HTTP Host". See [Target Http Host ](#target-http-host) below for details.

`http_referer` - (Optional) x-displayName: "HTTP Referer". See [Target Http Referer ](#target-http-referer) below for details.

`http_uri` - (Optional) x-displayName: "HTTP URI". See [Target Http Uri ](#target-http-uri) below for details.

### Action Choice Reset Traffic

x-displayName: "Reset Traffic".

`event` - (Required) x-required (`String`).

### Action Choice Retry

x-displayName: "Retry".

`event` - (Required) x-required (`String`).

### Action Choice Set Variable

x-displayName: "Set Variable".

`event` - (Required) x-required (`String`).

`expression` - (Optional) Tcl expression to evaluate. (`String`).

`name` - (Optional) Variable name. (`String`).

### Match Type Client Ssl

x-displayName: "Client SSL".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "cipher, cipher_bits, protocol" must be set

`cipher` - (Optional) x-displayName: "Cipher". See [Selector Cipher ](#selector-cipher) below for details.

`cipher_bits` - (Optional) x-displayName: "Cipher Bits". See [Selector Cipher Bits ](#selector-cipher-bits) below for details.

`protocol` - (Optional) x-displayName: "Protocol". See [Selector Protocol ](#selector-protocol) below for details.

### Match Type Cpu Usage

x-displayName: "CPU Usage".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`operator` - (Required) x-required (`String`).

`selector` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Match Type Geo Ip

x-displayName: "Geo. IP".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`interface` - (Optional) The interface (external or internal) on which to apply the condition. (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`operator` - (Required) x-required (`String`).

`selector` - (Required) x-required (`String`).

`side` - (Optional) The side (remote or local traffic) on which to apply the condition. (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Match Type Http Basic Auth

x-displayName: "HTTP Basic Auth".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`operator` - (Required) x-required (`String`).

`selector` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" must be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Match Type Http Connect

x-displayName: "HTTP Connect".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "host, port" must be set

`host` - (Optional) x-displayName: "Host". See [Selector Host ](#selector-host) below for details.

`port` - (Optional) x-displayName: "Port". See [Selector Port ](#selector-port) below for details.

### Match Type Http Cookie

x-displayName: "HTTP Cookie".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`operator` - (Required) x-required (`String`).

`selector` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Match Type Http Header

x-displayName: "HTTP Header".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`operator` - (Required) x-required (`String`).

`selector` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Match Type Http Host

x-displayName: "HTTP Host".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "full_string, host, port" must be set

`full_string` - (Optional) x-displayName: "Full String". See [Selector Full String ](#selector-full-string) below for details.

`host` - (Optional) x-displayName: "Host". See [Selector Host ](#selector-host) below for details.

`port` - (Optional) x-displayName: "Port". See [Selector Port ](#selector-port) below for details.

### Match Type Http Method

x-displayName: "HTTP Method".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Match Type Http Proxy

x-displayName: "HTTP Proxy".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "address, port, route_domain" must be set

`address` - (Optional) x-displayName: "Address". See [Selector Address ](#selector-address) below for details.

`port` - (Optional) x-displayName: "Port". See [Selector Port ](#selector-port) below for details.

`route_domain` - (Optional) x-displayName: "Route Domain". See [Selector Route Domain ](#selector-route-domain) below for details.

### Match Type Http Referer

x-displayName: "HTTP Referer".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`normalize` - (Optional) x-displayName: "Normalize" (`String`).

###### One of the arguments from this list "url_extension, url_full_string, url_host, url_path, url_path_segment, url_port, url_query_parameter, url_query_string, url_scheme, url_unnamed_query_parameter" can be set

`url_extension` - (Optional) x-displayName: "URL Extension". See [Selector Url Extension ](#selector-url-extension) below for details.

`url_full_string` - (Optional) x-displayName: "URL Full String". See [Selector Url Full String ](#selector-url-full-string) below for details.

`url_host` - (Optional) x-displayName: "URL Host". See [Selector Url Host ](#selector-url-host) below for details.

`url_path` - (Optional) x-displayName: "URL Path". See [Selector Url Path ](#selector-url-path) below for details.

`url_path_segment` - (Optional) x-displayName: "URL Path Segment". See [Selector Url Path Segment ](#selector-url-path-segment) below for details.

`url_port` - (Optional) x-displayName: "URL Port". See [Selector Url Port ](#selector-url-port) below for details.

`url_query_parameter` - (Optional) x-displayName: "URL Query Parameter". See [Selector Url Query Parameter ](#selector-url-query-parameter) below for details.

`url_query_string` - (Optional) x-displayName: "URL Query String". See [Selector Url Query String ](#selector-url-query-string) below for details.

`url_scheme` - (Optional) x-displayName: "URL Scheme". See [Selector Url Scheme ](#selector-url-scheme) below for details.

`url_unnamed_query_parameter` - (Optional) x-displayName: "URL Unnamed Query Parameter". See [Selector Url Unnamed Query Parameter ](#selector-url-unnamed-query-parameter) below for details.

### Match Type Http Set Cookie

x-displayName: "HTTP Set Cookie".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`name` - (Required) x-required (`String`).

`operator` - (Required) x-required (`String`).

`selector` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Match Type Http Status

x-displayName: "HTTP Status".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "code, full_string, text" can be set

`code` - (Optional) x-displayName: "Code". See [Selector Code ](#selector-code) below for details.

`full_string` - (Optional) x-displayName: "Full String". See [Selector Full String ](#selector-full-string) below for details.

`text` - (Optional) x-displayName: "Text". See [Selector Text ](#selector-text) below for details.

### Match Type Http Uri

x-displayName: "HTTP URI".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`normalize` - (Optional) x-displayName: "Normalization Choice" (`String`).

###### One of the arguments from this list "url_extension, url_full_string, url_host, url_path, url_path_segment, url_port, url_query_parameter, url_query_string, url_scheme, url_unnamed_query_parameter" can be set

`url_extension` - (Optional) x-displayName: "URL Extension". See [Selector Url Extension ](#selector-url-extension) below for details.

`url_full_string` - (Optional) x-displayName: "URL Full String". See [Selector Url Full String ](#selector-url-full-string) below for details.

`url_host` - (Optional) x-displayName: "URL Host". See [Selector Url Host ](#selector-url-host) below for details.

`url_path` - (Optional) x-displayName: "URL Path". See [Selector Url Path ](#selector-url-path) below for details.

`url_path_segment` - (Optional) x-displayName: "URL Path Segment". See [Selector Url Path Segment ](#selector-url-path-segment) below for details.

`url_port` - (Optional) x-displayName: "URL Port". See [Selector Url Port ](#selector-url-port) below for details.

`url_query_parameter` - (Optional) x-displayName: "URL Query Parameter". See [Selector Url Query Parameter ](#selector-url-query-parameter) below for details.

`url_query_string` - (Optional) x-displayName: "URL Query String". See [Selector Url Query String ](#selector-url-query-string) below for details.

`url_scheme` - (Optional) x-displayName: "URL Scheme". See [Selector Url Scheme ](#selector-url-scheme) below for details.

`url_unnamed_query_parameter` - (Optional) x-displayName: "URL Unnamed Query Parameter". See [Selector Url Unnamed Query Parameter ](#selector-url-unnamed-query-parameter) below for details.

### Match Type Http User Agent

x-displayName: "HTTP User Agent".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "browser_type, browser_version, device_make, device_model, token" can be set

`browser_type` - (Optional) x-displayName: "Browser Type". See [Selector Browser Type ](#selector-browser-type) below for details.

`browser_version` - (Optional) x-displayName: "Browser Version". See [Selector Browser Version ](#selector-browser-version) below for details.

`device_make` - (Optional) x-displayName: "Device Make". See [Selector Device Make ](#selector-device-make) below for details.

`device_model` - (Optional) x-displayName: "Device Model". See [Selector Device Model ](#selector-device-model) below for details.

`token` - (Optional) x-displayName: "User Agent Token". See [Selector Token ](#selector-token) below for details.

### Match Type Http Version

x-displayName: "HTTP Version".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "version_full_string, version_major, version_minor, version_protocol" can be set

`version_full_string` - (Optional) x-displayName: "Full String". See [Selector Version Full String ](#selector-version-full-string) below for details.

`version_major` - (Optional) x-displayName: "Major". See [Selector Version Major ](#selector-version-major) below for details.

`version_minor` - (Optional) x-displayName: "Minor". See [Selector Version Minor ](#selector-version-minor) below for details.

`version_protocol` - (Optional) x-displayName: "Protocol". See [Selector Version Protocol ](#selector-version-protocol) below for details.

### Match Type Ip

x-displayName: "IP".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "ip_protocol, ip_version" can be set

`ip_protocol` - (Optional) x-displayName: "IP Protocol". See [Selector Ip Protocol ](#selector-ip-protocol) below for details.

`ip_version` - (Optional) x-displayName: "IP Version". See [Selector Ip Version ](#selector-ip-version) below for details.

### Match Type Ip Reputation

x-displayName: "IP Reputation".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`interface` - (Optional) The interface (external or internal) on which to apply the condition. (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`operator` - (Required) x-required (`String`).

`side` - (Optional) The side (remote or local traffic) on which to apply the condition. (`String`).

`values` - (Required) x-required (`List of Strings`).

### Match Type Ssl Certificate

x-displayName: "SSL Certificate".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`operator` - (Required) x-required (`String`).

`selector` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Match Type Ssl Extension

x-displayName: "SSL Extension".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "alpn, npn, server_name" can be set

`alpn` - (Optional) x-displayName: "ALPN". See [Selector Alpn ](#selector-alpn) below for details.

`npn` - (Optional) x-displayName: "NPN". See [Selector Npn ](#selector-npn) below for details.

`server_name` - (Optional) x-displayName: "Server Name". See [Selector Server Name ](#selector-server-name) below for details.

### Match Type Tcp

x-displayName: "TCP".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`interface` - (Optional) The interface (external or internal) on which to apply the condition. (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

###### One of the arguments from this list "tcp_address, tcp_mss, tcp_port, tcp_route_domain, tcp_rtt, tcp_vlan, tcp_vlan_id" must be set

`tcp_address` - (Optional) x-displayName: "Address". See [Selector Tcp Address ](#selector-tcp-address) below for details.

`tcp_mss` - (Optional) x-displayName: "MSS". See [Selector Tcp Mss ](#selector-tcp-mss) below for details.

`tcp_port` - (Optional) x-displayName: "Port". See [Selector Tcp Port ](#selector-tcp-port) below for details.

`tcp_route_domain` - (Optional) x-displayName: "Route Domain". See [Selector Tcp Route Domain ](#selector-tcp-route-domain) below for details.

`tcp_rtt` - (Optional) x-displayName: "RTT". See [Selector Tcp Rtt ](#selector-tcp-rtt) below for details.

`tcp_vlan` - (Optional) x-displayName: "VLAN". See [Selector Tcp Vlan ](#selector-tcp-vlan) below for details.

`tcp_vlan_id` - (Optional) x-displayName: "VLAN ID". See [Selector Tcp Vlan Id ](#selector-tcp-vlan-id) below for details.

`side` - (Optional) The side (remote or local traffic) on which to apply the condition. (`String`).

### Match Type Web Socket

x-displayName: "WebSocket".

`case_sensitivity` - (Optional) Use case sensitive string comparison. (`String`).

`event` - (Required) x-required (`String`).

`missing` - (Optional) Skip this condition if it is missing from the request. (`String`).

`operator` - (Required) x-required (`String`).

`selector` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Method Carp

x-displayName: "Carp".

`key` - (Optional) x-displayName: "Key" (`String`).

`timeout` - (Optional) Timeout value in seconds. (`Int`).

### Method Cookie Hash

x-displayName: "Cookie Hash".

`cookie_name` - (Optional) x-displayName: "Cookie Name" (`String`).

`length` - (Optional) Specifies the length of data within the packet in bytes that the system uses to calculate the hash value. (`Int`).

`offset` - (Optional) Specifies the start offset within the packet from which the system begins the hash. (`Int`).

`timeout` - (Optional) Timeout value in seconds. (`Int`).

### Method Cookie Insert

x-displayName: "Cookie Insert".

`cookie_name` - (Optional) x-displayName: "Cookie Name" (`String`).

`timeout` - (Optional) Timeout value in seconds. (`Int`).

### Method Cookie Passive

x-displayName: "Cookie Passive".

`cookie_name` - (Optional) x-displayName: "Cookie Name" (`String`).

### Method Cookie Rewrite

x-displayName: "Cookie Rewrite".

`cookie_name` - (Optional) x-displayName: "Cookie Name" (`String`).

`timeout` - (Optional) Timeout value in seconds. (`Int`).

### Method Destination Address

x-displayName: "Destination Address".

`netmask` - (Optional) x-example: "192.168.13.23/16" (`String`).

`timeout` - (Optional) Timeout value in seconds. (`Int`).

### Method Hash

x-displayName: "Hash".

`key` - (Optional) x-displayName: "Key" (`String`).

`timeout` - (Optional) Timeout value in seconds. (`Int`).

### Method Source Address

x-displayName: "Source Address".

`netmask` - (Optional) x-example: "192.168.13.23/16" (`String`).

`timeout` - (Optional) Timeout value in seconds. (`Int`).

### Method Universal

x-displayName: "Universal".

`key` - (Optional) x-displayName: "Key" (`String`).

`timeout` - (Optional) Timeout value in seconds. (`Int`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Remote Server Disable Remote Server

x-displayName: "Disable".

### Remote Server Enable Remote Server

x-displayName: "Enable".

`host` - (Optional) x-displayName: "Host" (`String`).

`port` - (Optional) x-displayName: "Port" (`Int`).

### Rules Actions

x-displayName: "Actions".

###### One of the arguments from this list "collect_data, disable, enable, forward_traffic, insert, log, persist_session, redirect, remove, replace, reset_traffic, retry, set_variable" can be set

`collect_data` - (Optional) x-displayName: "Collect Data". See [Action Choice Collect Data ](#action-choice-collect-data) below for details.

`disable` - (Optional) x-displayName: "Disable". See [Action Choice Disable ](#action-choice-disable) below for details.

`enable` - (Optional) x-displayName: "Enable". See [Action Choice Enable ](#action-choice-enable) below for details.

`forward_traffic` - (Optional) x-displayName: "Forward Traffic". See [Action Choice Forward Traffic ](#action-choice-forward-traffic) below for details.

`insert` - (Optional) x-displayName: "Insert". See [Action Choice Insert ](#action-choice-insert) below for details.

`log` - (Optional) x-displayName: "Log". See [Action Choice Log ](#action-choice-log) below for details.

`persist_session` - (Optional) x-displayName: "Persist Session". See [Action Choice Persist Session ](#action-choice-persist-session) below for details.

`redirect` - (Optional) x-displayName: "Redirect". See [Action Choice Redirect ](#action-choice-redirect) below for details.

`remove` - (Optional) x-displayName: "Remove". See [Action Choice Remove ](#action-choice-remove) below for details.

`replace` - (Optional) x-displayName: "Replace". See [Action Choice Replace ](#action-choice-replace) below for details.

`reset_traffic` - (Optional) x-displayName: "Reset Traffic". See [Action Choice Reset Traffic ](#action-choice-reset-traffic) below for details.

`retry` - (Optional) x-displayName: "Retry". See [Action Choice Retry ](#action-choice-retry) below for details.

`set_variable` - (Optional) x-displayName: "Set Variable". See [Action Choice Set Variable ](#action-choice-set-variable) below for details.

### Rules Conditions

x-displayName: "Conditions".

###### One of the arguments from this list "client_ssl, cpu_usage, geo_ip, http_basic_auth, http_connect, http_cookie, http_header, http_host, http_method, http_proxy, http_referer, http_set_cookie, http_status, http_uri, http_user_agent, http_version, ip, ip_reputation, ssl_certificate, ssl_extension, tcp, web_socket" can be set

`client_ssl` - (Optional) x-displayName: "Client SSL". See [Match Type Client Ssl ](#match-type-client-ssl) below for details.

`cpu_usage` - (Optional) x-displayName: "CPU Usage". See [Match Type Cpu Usage ](#match-type-cpu-usage) below for details.

`geo_ip` - (Optional) x-displayName: "Geo. IP". See [Match Type Geo Ip ](#match-type-geo-ip) below for details.(Deprecated)

`http_basic_auth` - (Optional) x-displayName: "HTTP Basic Auth". See [Match Type Http Basic Auth ](#match-type-http-basic-auth) below for details.

`http_connect` - (Optional) x-displayName: "HTTP Connect". See [Match Type Http Connect ](#match-type-http-connect) below for details.(Deprecated)

`http_cookie` - (Optional) x-displayName: "HTTP Cookie". See [Match Type Http Cookie ](#match-type-http-cookie) below for details.

`http_header` - (Optional) x-displayName: "HTTP Header". See [Match Type Http Header ](#match-type-http-header) below for details.

`http_host` - (Optional) x-displayName: "HTTP Host". See [Match Type Http Host ](#match-type-http-host) below for details.

`http_method` - (Optional) x-displayName: "HTTP Method". See [Match Type Http Method ](#match-type-http-method) below for details.

`http_proxy` - (Optional) x-displayName: "HTTP Proxy". See [Match Type Http Proxy ](#match-type-http-proxy) below for details.(Deprecated)

`http_referer` - (Optional) x-displayName: "HTTP Referer". See [Match Type Http Referer ](#match-type-http-referer) below for details.

`http_set_cookie` - (Optional) x-displayName: "HTTP Set Cookie". See [Match Type Http Set Cookie ](#match-type-http-set-cookie) below for details.

`http_status` - (Optional) x-displayName: "HTTP Status". See [Match Type Http Status ](#match-type-http-status) below for details.

`http_uri` - (Optional) x-displayName: "HTTP URI". See [Match Type Http Uri ](#match-type-http-uri) below for details.

`http_user_agent` - (Optional) x-displayName: "HTTP User Agent". See [Match Type Http User Agent ](#match-type-http-user-agent) below for details.

`http_version` - (Optional) x-displayName: "HTTP Version". See [Match Type Http Version ](#match-type-http-version) below for details.

`ip` - (Optional) x-displayName: "IP". See [Match Type Ip ](#match-type-ip) below for details.

`ip_reputation` - (Optional) x-displayName: "IP Reputation". See [Match Type Ip Reputation ](#match-type-ip-reputation) below for details.(Deprecated)

`ssl_certificate` - (Optional) x-displayName: "SSL Certificate". See [Match Type Ssl Certificate ](#match-type-ssl-certificate) below for details.

`ssl_extension` - (Optional) x-displayName: "SSL Extension". See [Match Type Ssl Extension ](#match-type-ssl-extension) below for details.

`tcp` - (Optional) x-displayName: "TCP". See [Match Type Tcp ](#match-type-tcp) below for details.

`web_socket` - (Optional) x-displayName: "WebSocket". See [Match Type Web Socket ](#match-type-web-socket) below for details.(Deprecated)

### Selector Address

x-displayName: "Address".

`operator` - (Optional) x-displayName: "Operator" (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "Address Values". See [Values String Values ](#values-string-values) below for details.

### Selector Alpn

x-displayName: "ALPN".

`index` - (Optional) x-displayName: "Index" (`Int`).

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Browser Type

x-displayName: "Browser Type".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Browser Version

x-displayName: "Browser Version".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Cipher

x-displayName: "Cipher".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Cipher Bits

x-displayName: "Cipher Bits".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Code

x-displayName: "Code".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Device Make

x-displayName: "Device Make".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Device Model

x-displayName: "Device Model".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Full String

x-displayName: "Full String".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" must be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Full String

x-displayName: "Full String".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Host

x-displayName: "Host".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Host

x-displayName: "Host".

`operator` - (Optional) x-displayName: "Operator" (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Ip Protocol

x-displayName: "IP Protocol".

`ip_protocols` - (Required) x-required (`String`).

`operator` - (Required) x-required (`String`).

### Selector Ip Version

x-displayName: "IP Version".

`ip_version` - (Required) x-required (`List of Strings`).

`operator` - (Required) x-required (`String`).

### Selector Npn

x-displayName: "NPN".

`index` - (Optional) x-displayName: "Index" (`Int`).

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Port

x-displayName: "Port".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Port

x-displayName: "Port".

`operator` - (Optional) x-displayName: "Operator" (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Port

x-displayName: "Port".

`operator` - (Optional) x-displayName: "Operator" (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Protocol

x-displayName: "Protocol".

`operator` - (Optional) x-displayName: "Operator" (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Route Domain

x-displayName: "Route Domain".

`operator` - (Optional) x-displayName: "Operator" (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Server Name

x-displayName: "Server Name".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Tcp Address

x-displayName: "Address".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "Address Values". See [Values String Values ](#values-string-values) below for details.

### Selector Tcp Mss

x-displayName: "MSS".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Tcp Port

x-displayName: "Port".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Tcp Route Domain

x-displayName: "Route Domain".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Tcp Rtt

x-displayName: "RTT".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Tcp Vlan

x-displayName: "VLAN".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Tcp Vlan Id

x-displayName: "VLAN ID".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "String Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Text

x-displayName: "Text".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Token

x-displayName: "User Agent Token".

`operator` - (Required) x-required (`String`).

`ua_token` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Url Extension

x-displayName: "URL Extension".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Url Full String

x-displayName: "URL Full String".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Url Host

x-displayName: "URL Host".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Url Path

x-displayName: "URL Path".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Url Path Segment

x-displayName: "URL Path Segment".

`index` - (Optional) The numeric order of the selector, starting at 1. Negative values indicate counting right to left. (`Int`).

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Url Port

x-displayName: "URL Port".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Url Query Parameter

x-displayName: "URL Query Parameter".

`name` - (Optional) Name of the particular selector whose value is to be used. (`String`).

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Url Query String

x-displayName: "URL Query String".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Url Scheme

x-displayName: "URL Scheme".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Url Unnamed Query Parameter

x-displayName: "URL Unnamed Query Parameter".

`index` - (Optional) The numeric order of the selector, starting at 1. Negative values indicate counting right to left. (`Int`).

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Version Full String

x-displayName: "Full String".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Selector Version Major

x-displayName: "Major".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Version Minor

x-displayName: "Minor".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, int_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`int_values` - (Optional) x-displayName: "Integer Values". See [Values Int Values ](#values-int-values) below for details.

### Selector Version Protocol

x-displayName: "Protocol".

`operator` - (Required) x-required (`String`).

###### One of the arguments from this list "datagroup, string_values" can be set

`datagroup` - (Optional) x-displayName: "Data Group". See [Values Datagroup ](#values-datagroup) below for details.

`string_values` - (Optional) x-displayName: "String Values". See [Values String Values ](#values-string-values) below for details.

### Target Http Connect

x-displayName: "HTTP Connect".

`destination` - (Optional) x-displayName: "Destination" (`String`).

`port` - (Optional) x-displayName: "Port" (`Int`).

### Target Http Header

x-displayName: "HTTP Header".

`name` - (Optional) x-displayName: "Name" (`String`).

`value` - (Optional) x-displayName: "Value" (`String`).

### Target Http Host

x-displayName: "HTTP Host".

`value` - (Optional) x-displayName: "Value" (`String`).

### Target Http Referer

x-displayName: "HTTP Referer".

`value` - (Optional) x-displayName: "Value" (`String`).

### Target Http Uri

x-displayName: "HTTP URI".

`uri_component` - (Optional) x-displayName: "URI Matching Component" (`String`).

`value` - (Optional) x-displayName: "Value" (`String`).

### Target Pool

x-displayName: "Pool".

`fallback_pool` - (Optional) Forward connection to the specified pool when the default pool does not have active members.. See [ref](#ref) below for details.

`pool` - (Optional) x-displayName: "Pool". See [ref](#ref) below for details.

### Target Virtual Server

x-displayName: "Virtual Server".

`virtual_server` - (Optional) x-displayName: "Virtual Server". See [ref](#ref) below for details.

### Values Datagroup

x-displayName: "Data Group".

`datagroup` - (Optional) A data group direct reference. See [ref](#ref) below for details.

### Values Int Values

x-displayName: "Integer Values".

`values` - (Optional) x-displayName: "Values" (`Int`).

### Values String Values

x-displayName: "String Values".

`values` - (Optional) x-displayName: "Values" (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured traffic_policy.
