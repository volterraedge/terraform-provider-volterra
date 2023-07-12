---

page_title: "Volterra: route"

description: "The route allows CRUD of Route resource on Volterra SaaS"
-----------------------------------------------------------------------

Resource volterra_route
=======================

The Route allows CRUD of Route resource on Volterra SaaS

~> **Note:** Please refer to [Route API docs](https://volterra.io/docs/api/route) to learn more

Example Usage
-------------

```hcl
resource "volterra_route" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  routes {
    // One of the arguments from this list "inherited_bot_defense_javascript_injection bot_defense_javascript_injection" must be set

    bot_defense_javascript_injection {
      javascript_location = "javascript_location"

      javascript_tags {
        javascript_url = "https://www.example.com/login/common.js?single"

        tag_attributes {
          javascript_tag = "ID"
          tag_value      = "_imp_apg_dip_"
        }
      }
    }

    bot_defense_javascript_injection_inline_mode {
      element_selector = "value"
      insert_content   = "value"
      position         = "position"
    }

    disable_custom_script = true
    disable_location_add  = true

    match {
      headers {
        invert_match = true
        name         = "Content-Type"

        // One of the arguments from this list "exact regex presence" must be set
        exact = "application/json"
      }

      http_method = "http_method"

      incoming_port {
        // One of the arguments from this list "port port_ranges no_port_match" must be set
        port = "6443"
      }

      path {
        // One of the arguments from this list "path regex prefix" must be set
        prefix = "/register/"
      }

      query_params {
        key = "assignee_username"

        // One of the arguments from this list "exact regex" must be set
        exact = "exact"
      }
    }

    request_headers_to_add {
      append = true
      name   = "value"

      // One of the arguments from this list "value secret_value" must be set
      value = "value"
    }

    request_headers_to_remove = ["host"]

    response_headers_to_add {
      append = true
      name   = "value"

      // One of the arguments from this list "value secret_value" must be set
      value = "value"
    }

    response_headers_to_remove = ["host"]

    // One of the arguments from this list "route_destination route_redirect route_direct_response" must be set

    route_destination {
      buffer_policy {
        disabled          = true
        max_request_bytes = "2048"
        max_request_time  = "30"
      }

      // One of the arguments from this list "do_not_retract_cluster retract_cluster" must be set
      retract_cluster = true

      cors_policy {
        allow_credentials = true
        allow_headers     = "value"
        allow_methods     = "GET"

        allow_origin = ["value"]

        allow_origin_regex = ["value"]
        disabled           = true
        expose_headers     = "value"
        max_age            = "value"
        maximum_age        = "-1"
      }

      destinations {
        cluster {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }

        endpoint_subsets = {
          "key1" = "value1"
        }

        priority = "1"
        weight   = "10"
      }

      endpoint_subsets = {
        "key1" = "value1"
      }

      hash_policy {
        // One of the arguments from this list "header_name cookie source_ip" must be set
        header_name = "host"

        terminal = true
      }

      // One of the arguments from this list "host_rewrite auto_host_rewrite" must be set
      host_rewrite = "one.volterra.com"

      mirror_policy {
        cluster {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }

        percent {
          denominator = "denominator"
          numerator   = "5"
        }
      }

      prefix_rewrite = "/"
      priority       = "priority"

      retry_policy {
        back_off {
          base_interval = "5"
          max_interval  = "60"
        }

        num_retries     = "3"
        per_try_timeout = "1000"

        retriable_status_codes = ["403"]

        retry_condition = ["5xx"]
        retry_on        = "5xx"
      }

      spdy_config {
        use_spdy = true
      }

      timeout = "2000"

      web_socket_config {
        idle_timeout         = "2000"
        max_connect_attempts = "5"
        use_websocket        = true
      }
    }
    service_policy {
      // One of the arguments from this list "disable context_extensions" must be set
      disable = true
    }
    skip_lb_override = true
    waf_type {
      // One of the arguments from this list "app_firewall disable_waf inherit_waf" must be set

      app_firewall {
        app_firewall {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }
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

`routes` - (Required) List of routes to match for incoming request. See [Routes ](#routes) below for details.

### Add Httponly

Add httponly attribute.

### Add Secure

Add secure attribute.

### App Firewall

A direct reference to an Application Firewall configuration object.

`app_firewall` - (Required) References to an Application Firewall configuration object. See [ref](#ref) below for details.

### Back Off

10 times the base interval.

`base_interval` - (Optional) Specifies the base interval between retries in milliseconds (`Int`).

`max_interval` - (Optional) to the base_interval if set. The default is 10 times the base_interval. (`Int`).

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Bot Defense Javascript Injection

Configuration for Bot Defense Javascript Injection.

`javascript_location` - (Optional) Select the location where you would like to insert the Javascript tag(s). (`String`).

`javascript_tags` - (Required) Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first.. See [Javascript Tags ](#javascript-tags) below for details.

### Bot Defense Javascript Injection Inline Mode

Specifies whether bot defense js injection inline mode will be enabled.

`element_selector` - (Required) Element selector to insert into. (`String`).

`insert_content` - (Optional) HTML content to insert. (`String`).

`position` - (Optional) Position of HTML content to be inserted within HTML tag. (`String`).

### Buffer Policy

Route level buffer configuration overrides any configuration at VirtualHost level..

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) manager will stop buffering and return a RequestEntityTooLarge (413) response. (`Int`).

`max_request_time` - (Optional) request before returning a RequestTimeout (408) response (`Int`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Context Extensions

sending additional information to the external authorization server..

`context_extensions` - (Optional) provide extra context for the external authorization server on specific virtual hosts or routes. (`String`).

### Cookie

Hash based on cookie.

`add_httponly` - (Optional) Add httponly attribute (bool).

`ignore_httponly` - (Optional) Ignore httponly attribute (bool).

`name` - (Required) produced (`String`).

`path` - (Optional) will be set for the cookie (`String`).

`ignore_samesite` - (Optional) Ignore Samesite attribute (bool).

`samesite_lax` - (Optional) Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests (bool).

`samesite_none` - (Optional) Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests (bool).

`samesite_strict` - (Optional) Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests (bool).

`add_secure` - (Optional) Add secure attribute (bool).

`ignore_secure` - (Optional) Ignore secure attribute (bool).

`ttl` - (Optional) be a session cookie. TTL value is in milliseconds (`Int`).

### Cors Policy

resources from a server at a different origin.

`allow_credentials` - (Optional) Specifies whether the resource allows credentials (`Bool`).

`allow_headers` - (Optional) Specifies the content for the access-control-allow-headers header (`String`).

`allow_methods` - (Optional) Specifies the content for the access-control-allow-methods header (`String`).

`allow_origin` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`allow_origin_regex` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`expose_headers` - (Optional) Specifies the content for the access-control-expose-headers header (`String`).

`max_age` - (Optional) Specifies the content for the access-control-max-age header (`String`).

`maximum_age` - (Optional) Maximum permitted value is 86400 seconds (24 hours) (`Int`).

### Destinations

sent to the cluster specified in the destination.

`cluster` - (Required) does not exist ServiceUnavailable response will be sent. See [ref](#ref) below for details.

`endpoint_subsets` - (Optional) upstream cluster which match this metadata will be selected for load balancing (`String`).

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) sent to the cluster specified in the destination (`Int`).

### Disable Waf

Any Application Firewall configuration will not be enforced.

### Do Not Retract Cluster

configuration..

### Hash Policy

route the request.

`cookie` - (Optional) Hash based on cookie. See [Cookie ](#cookie) below for details.

`header_name` - (Optional) The name or key of the request header that will be used to obtain the hash key (`String`).

`source_ip` - (Optional) Hash based on source IP address (`Bool`).

`terminal` - (Optional) Specify if its a terminal policy (`Bool`).

### Headers

List of (key, value) headers.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Ignore Httponly

Ignore httponly attribute.

### Ignore Samesite

Ignore Samesite attribute.

### Ignore Secure

Ignore secure attribute.

### Incoming Port

The port on which the request is received.

`no_port_match` - (Optional) Disable matching of ports (bool).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Inherit Waf

Any Application Firewall configuration that was configured on a higher level will be enforced.

### Inherited Bot Defense Javascript Injection

Hence no custom configuration is applied on the route.

### Javascript Tags

Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first..

`javascript_url` - (Required) Please enter the full URL (include domain and path), or relative path. (`String`).

`tag_attributes` - (Optional) Add the tag attributes you want to include in your Javascript tag.. See [Tag Attributes ](#tag-attributes) below for details.

### Match

route match condition.

`headers` - (Optional) List of (key, value) headers. See [Headers ](#headers) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`incoming_port` - (Optional) The port on which the request is received. See [Incoming Port ](#incoming-port) below for details.

`path` - (Optional) URI path of route. See [Path ](#path) below for details.

`query_params` - (Optional) List of (key, value) query parameters. See [Query Params ](#query-params) below for details.

### Mirror Policy

useful for logging. For example, *cluster1* becomes *cluster1-shadow*..

`cluster` - (Required) referred here must be present.. See [ref](#ref) below for details.

`percent` - (Optional) Percentage of requests to be mirrored. See [Percent ](#percent) below for details.

### No Port Match

Disable matching of ports.

### Path

URI path of route.

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Percent

Percentage of requests to be mirrored.

`denominator` - (Required) Samples per denominator. numerator part per 100 or 10000 ro 1000000 (`String`).

`numerator` - (Required) sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000 (`Int`).

### Query Params

List of (key, value) query parameters.

`key` - (Required) In the above example, assignee_username is the key (`String`).

`exact` - (Optional) Exact match value for the query parameter key (`String`).

`regex` - (Optional) Regex match value for the query parameter key (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Remove All Params

x-displayName: "Remove All Parameters".

### Request Headers To Add

enclosing VirtualHost object level.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Secret Value ](#secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Response Headers To Add

enclosing VirtualHost object level.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Secret Value ](#secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Retain All Params

x-displayName: "Retain All Parameters".

### Retract Cluster

for route.

### Retry Policy

Indicates that the route has a retry policy..

`back_off` - (Optional) 10 times the base interval. See [Back Off ](#back-off) below for details.

`num_retries` - (Optional) is used between each retry (`Int`).

`per_try_timeout` - (Optional) Specifies a non-zero timeout per retry attempt. In milliseconds (`Int`).

`retriable_status_codes` - (Optional) HTTP status codes that should trigger a retry in addition to those specified by retry_on. (`Int`).

`retry_condition` - (Optional) matching one defined in retriable_status_codes field (`String`).

`retry_on` - (Optional) matching one defined in retriable_status_codes field (`String`).

### Route Destination

Send request to one of the destination from list of destinations.

`buffer_policy` - (Optional) Route level buffer configuration overrides any configuration at VirtualHost level.. See [Buffer Policy ](#buffer-policy) below for details.

`do_not_retract_cluster` - (Optional) configuration. (bool).

`retract_cluster` - (Optional) for route (bool).

`cors_policy` - (Optional) resources from a server at a different origin. See [Cors Policy ](#cors-policy) below for details.

`destinations` - (Required) sent to the cluster specified in the destination. See [Destinations ](#destinations) below for details.

`endpoint_subsets` - (Optional) upstream cluster which match this metadata will be selected for load balancing (`String`).

`hash_policy` - (Optional) route the request. See [Hash Policy ](#hash-policy) below for details.

`auto_host_rewrite` - (Optional) of the upstream host chosen by the cluster (`Bool`).

`host_rewrite` - (Optional) Indicates that during forwarding, the host header will be swapped with this value (`String`).

`mirror_policy` - (Optional) useful for logging. For example, *cluster1* becomes *cluster1-shadow*.. See [Mirror Policy ](#mirror-policy) below for details.

`prefix_rewrite` - (Optional) while requests to /register/public will be stripped to /public (`String`).

`priority` - (Optional) Also, circuit-breaker configuration at destination cluster is chosen based on the route priority. (`String`).

`retry_policy` - (Optional) Indicates that the route has a retry policy.. See [Retry Policy ](#retry-policy) below for details.

`spdy_config` - (Optional) SPDY configuration for each route. See [Spdy Config ](#spdy-config) below for details.

`timeout` - (Optional) for infinite timeout (`Int`).

`web_socket_config` - (Optional) Websocket configuration for each route. See [Web Socket Config ](#web-socket-config) below for details.

### Route Direct Response

Send direct response.

`response_body` - (Optional) response body to send (`String`).

`response_code` - (Optional) response code to send (`Int`).

### Route Redirect

Send redirect response.

`host_redirect` - (Optional) swap host part of incoming URL in redirect URL (`String`).

`port_redirect` - (Optional) Specify the port value to redirect to a URL with non default port(443) (`Int`).

`proto_redirect` - (Optional) When incoming-proto option is specified, swapping of protocol is not done. (`String`).

`all_params` - (Optional) be removed. Default value is false, which means query portion of the URL will NOT be removed (`Bool`).

`remove_all_params` - (Optional) x-displayName: "Remove All Parameters" (bool).

`replace_params` - (Optional) x-displayName: "Replace All Parameters" (`String`).

`retain_all_params` - (Optional) x-displayName: "Retain All Parameters" (bool).

`strip_query_params` - (Optional) Specifies the list of query params to be removed. Not supported. See [Strip Query Params ](#strip-query-params) below for details.

`path_redirect` - (Optional) swap path part of incoming URL in redirect URL (`String`).

`prefix_rewrite` - (Optional) This option allows redirect URLs be dynamically created based on the request (`String`).

`response_code` - (Optional) The HTTP status code to use in the redirect response. (`Int`).

### Routes

List of routes to match for incoming request.

`bot_defense_javascript_injection` - (Optional) Configuration for Bot Defense Javascript Injection. See [Bot Defense Javascript Injection ](#bot-defense-javascript-injection) below for details.

`inherited_bot_defense_javascript_injection` - (Optional) Hence no custom configuration is applied on the route (bool).

`bot_defense_javascript_injection_inline_mode` - (Optional) Specifies whether bot defense js injection inline mode will be enabled. See [Bot Defense Javascript Injection Inline Mode ](#bot-defense-javascript-injection-inline-mode) below for details.

`disable_custom_script` - (Optional) disable execution of Javascript at route level, if it is configured at virtual-host level (`Bool`).

`disable_location_add` - (Optional) virtual-host level. This configuration is ignored on CE sites. (`Bool`).

`match` - (Optional) route match condition. See [Match ](#match) below for details.

`request_headers_to_add` - (Optional) enclosing VirtualHost object level. See [Request Headers To Add ](#request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) enclosing VirtualHost object level. See [Response Headers To Add ](#response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

`route_destination` - (Optional) Send request to one of the destination from list of destinations. See [Route Destination ](#route-destination) below for details.

`route_direct_response` - (Optional) Send direct response. See [Route Direct Response ](#route-direct-response) below for details.

`route_redirect` - (Optional) Send redirect response. See [Route Redirect ](#route-redirect) below for details.

`service_policy` - (Optional) service policy configuration at route level which overrides configuration at virtual-host level. See [Service Policy ](#service-policy) below for details.

`skip_lb_override` - (Optional) these routes. (`Bool`).

`waf_type` - (Optional) waf_type specified at route level overrides waf configuration at VirtualHost level. See [Waf Type ](#waf-type) below for details.

### Samesite Lax

Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests.

### Samesite None

Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests.

### Samesite Strict

Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests.

### Secret Value

Secret Value of the HTTP header..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Service Policy

service policy configuration at route level which overrides configuration at virtual-host level.

`context_extensions` - (Optional) sending additional information to the external authorization server.. See [Context Extensions ](#context-extensions) below for details.

`disable` - (Optional) disable service policy at route level, if it is configured at virtual-host level (`Bool`).

### Spdy Config

SPDY configuration for each route.

`use_spdy` - (Optional) a SPDY connection (`Bool`).

### Strip Query Params

Specifies the list of query params to be removed. Not supported.

`query_params` - (Optional) Query params keys to strip while manipulating the HTTP request (`String`).

### Tag Attributes

Add the tag attributes you want to include in your Javascript tag..

`javascript_tag` - (Optional) Select from one of the predefined tag attibutes. (`String`).

`tag_value` - (Optional) Add the tag attribute value. (`String`).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Waf Type

waf_type specified at route level overrides waf configuration at VirtualHost level.

`app_firewall` - (Optional) A direct reference to an Application Firewall configuration object. See [App Firewall ](#app-firewall) below for details.

`disable_waf` - (Optional) Any Application Firewall configuration will not be enforced (bool).

`inherit_waf` - (Optional) Any Application Firewall configuration that was configured on a higher level will be enforced (bool).

### Web Socket Config

Websocket configuration for each route.

`idle_timeout` - (Optional) Idle Timeout for Websocket in milli seconds. After timeout, connection will be closed (`Int`).

`max_connect_attempts` - (Optional) giving up. Default is 1 (`Int`).

`use_websocket` - (Optional) a WebSocket connection (`Bool`).

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured route.
