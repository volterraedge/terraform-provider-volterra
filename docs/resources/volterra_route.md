---

page_title: "Volterra: route"

description: "The route allows CRUD of Route resource on Volterra SaaS"
-----------------------------------------------------------------------

Resource volterra_route
=======================

The Route allows CRUD of Route resource on Volterra SaaS

~> **Note:** Please refer to [Route API docs](https://docs.cloud.f5.com/docs/api/route) to learn more

Example Usage
-------------

```hcl
resource "volterra_route" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  routes {
    // One of the arguments from this list "bot_defense_javascript_injection inherited_bot_defense_javascript_injection" must be set

    bot_defense_javascript_injection {
      javascript_location = "javascript_location"

      javascript_tags {
        javascript_url = "https://www.example.com/login/common.js?single"

        tag_attributes {
          javascript_tag = "ID"

          tag_value = "_imp_apg_dip_"
        }
      }
    }

    bot_defense_javascript_injection_inline_mode {
      element_selector = "value"

      insert_content = "value"

      position = "position"
    }

    disable_custom_script = true

    disable_location_add = true

    match {
      headers {
        invert_match = true

        name = "Content-Type"

        // One of the arguments from this list "presence exact regex" must be set
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

        // One of the arguments from this list "regex exact" must be set
        exact = "exact"
      }
    }

    request_headers_to_add {
      append = true

      name = "value"

      // One of the arguments from this list "value secret_value" must be set
      value = "value"
    }

    request_headers_to_remove = ["host"]

    response_headers_to_add {
      append = true

      name = "value"

      // One of the arguments from this list "value secret_value" must be set
      value = "value"
    }

    response_headers_to_remove = ["host"]

    // One of the arguments from this list "route_destination route_redirect route_direct_response" must be set

    route_destination {
      buffer_policy {
        disabled = true

        max_request_bytes = "2048"

        max_request_time = "30"
      }

      // One of the arguments from this list "retract_cluster do_not_retract_cluster" must be set
      retract_cluster = true

      cors_policy {
        allow_credentials = true

        allow_headers = "value"

        allow_methods = "GET"

        allow_origin = ["value"]

        allow_origin_regex = ["value"]

        disabled = true

        expose_headers = "value"

        max_age = "value"

        maximum_age = "-1"
      }

      csrf_policy {
        // One of the arguments from this list "custom_domain_list disabled all_load_balancer_domains" must be set

        custom_domain_list {
          domains = ["www.foo.com"]
        }
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

        weight = "10"
      }

      endpoint_subsets = {
        "key1" = "value1"
      }

      hash_policy {
        // One of the arguments from this list "cookie source_ip header_name" must be set
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

          numerator = "5"
        }
      }

      priority = "priority"

      retry_policy {
        back_off {
          base_interval = "5"

          max_interval = "60"
        }

        num_retries = "3"

        per_try_timeout = "1000"

        retriable_status_codes = ["403"]

        retry_condition = ["5xx"]

        retry_on = "5xx"
      }

      // One of the arguments from this list "prefix_rewrite regex_rewrite" must be set
      prefix_rewrite = "/"

      spdy_config {
        use_spdy = true
      }

      timeout = "2000"

      web_socket_config {
        idle_timeout = "2000"

        max_connect_attempts = "5"

        use_websocket = true
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

### Routes

List of routes to match for incoming request.

###### One of the arguments from this list "inherited_bot_defense_javascript_injection, bot_defense_javascript_injection" can be set

`bot_defense_javascript_injection` - (Optional) Configuration for Bot Defense Javascript Injection. See [Bot Defense Javascript Injection Choice Bot Defense Javascript Injection ](#bot-defense-javascript-injection-choice-bot-defense-javascript-injection) below for details.

`inherited_bot_defense_javascript_injection` - (Optional) Hence no custom configuration is applied on the route (`Bool`).

`bot_defense_javascript_injection_inline_mode` - (Optional) Specifies whether bot defense js injection inline mode will be enabled. See [Routes Bot Defense Javascript Injection Inline Mode ](#routes-bot-defense-javascript-injection-inline-mode) below for details.(Deprecated)

`disable_custom_script` - (Optional) disable execution of Javascript at route level, if it is configured at virtual-host level (`Bool`).(Deprecated)

`disable_location_add` - (Optional) virtual-host level. This configuration is ignored on CE sites. (`Bool`).

`match` - (Optional) route match condition. See [Routes Match ](#routes-match) below for details.

`request_headers_to_add` - (Optional) enclosing VirtualHost object level. See [Routes Request Headers To Add ](#routes-request-headers-to-add) below for details.

`request_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP request being sent towards upstream. (`String`).

`response_headers_to_add` - (Optional) enclosing VirtualHost object level. See [Routes Response Headers To Add ](#routes-response-headers-to-add) below for details.

`response_headers_to_remove` - (Optional) List of keys of Headers to be removed from the HTTP response being sent towards downstream. (`String`).

###### One of the arguments from this list "route_redirect, route_direct_response, route_destination" must be set

`route_destination` - (Optional) Send request to one of the destination from list of destinations. See [Route Action Route Destination ](#route-action-route-destination) below for details.

`route_direct_response` - (Optional) Send direct response. See [Route Action Route Direct Response ](#route-action-route-direct-response) below for details.

`route_redirect` - (Optional) Send redirect response. See [Route Action Route Redirect ](#route-action-route-redirect) below for details.

`service_policy` - (Optional) service policy configuration at route level which overrides configuration at virtual-host level. See [Routes Service Policy ](#routes-service-policy) below for details.

`skip_lb_override` - (Optional) these routes. (`Bool`).(Deprecated)

`waf_type` - (Optional) waf_type specified at route level overrides waf configuration at VirtualHost level. See [Routes Waf Type ](#routes-waf-type) below for details.

### Allowed Domains All Load Balancer Domains

Add All load balancer domains to source origin (allow) list..

### Allowed Domains Custom Domain List

Add one or more domains to source origin (allow) list..

`domains` - (Required) Wildcard names are supported in the suffix or prefix form. (`String`).

### Allowed Domains Disabled

Allow all source origin domains..

### Bot Defense Javascript Injection Javascript Tags

Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first..

`javascript_url` - (Required) Please enter the full URL (include domain and path), or relative path. (`String`).

`tag_attributes` - (Optional) Add the tag attributes you want to include in your Javascript tag.. See [Javascript Tags Tag Attributes ](#javascript-tags-tag-attributes) below for details.

### Bot Defense Javascript Injection Choice Bot Defense Javascript Injection

Configuration for Bot Defense Javascript Injection.

`javascript_location` - (Optional) Select the location where you would like to insert the Javascript tag(s). (`String`).

`javascript_tags` - (Required) Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot Javascript should be added first.. See [Bot Defense Javascript Injection Javascript Tags ](#bot-defense-javascript-injection-javascript-tags) below for details.

### Bot Defense Javascript Injection Choice Inherited Bot Defense Javascript Injection

Hence no custom configuration is applied on the route.

### Cluster Retract Choice Do Not Retract Cluster

configuration..

### Cluster Retract Choice Retract Cluster

for route.

### Httponly Add Httponly

Add httponly attribute.

### Httponly Ignore Httponly

Ignore httponly attribute.

### Javascript Tags Tag Attributes

Add the tag attributes you want to include in your Javascript tag..

`javascript_tag` - (Optional) Select from one of the predefined tag attibutes. (`String`).

`tag_value` - (Optional) Add the tag attribute value. (`String`).

### Match Headers

List of (key, value) headers.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, regex, presence" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Match Incoming Port

The port on which the request is received.

###### One of the arguments from this list "port, port_ranges, no_port_match" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Match Path

URI path of route.

###### One of the arguments from this list "regex, prefix, path" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Match Query Params

List of (key, value) query parameters.

`key` - (Required) In the above example, assignee_username is the key (`String`).

###### One of the arguments from this list "exact, regex" can be set

`exact` - (Optional) Exact match value for the query parameter key (`String`).

`regex` - (Optional) Regex match value for the query parameter key (`String`).

### Mirror Policy Percent

Percentage of requests to be mirrored.

`denominator` - (Required) Samples per denominator. numerator part per 100 or 10000 ro 1000000 (`String`).

`numerator` - (Required) sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000 (`Int`).

### Policy Specifier Cookie

Hash based on cookie.

###### One of the arguments from this list "ignore_httponly, add_httponly" can be set

`add_httponly` - (Optional) Add httponly attribute (`Bool`).

`ignore_httponly` - (Optional) Ignore httponly attribute (`Bool`).

`name` - (Required) produced (`String`).

`path` - (Optional) will be set for the cookie (`String`).

###### One of the arguments from this list "samesite_lax, samesite_none, ignore_samesite, samesite_strict" can be set

`ignore_samesite` - (Optional) Ignore Samesite attribute (`Bool`).

`samesite_lax` - (Optional) Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests (`Bool`).

`samesite_none` - (Optional) Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests (`Bool`).

`samesite_strict` - (Optional) Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests (`Bool`).

###### One of the arguments from this list "ignore_secure, add_secure" can be set

`add_secure` - (Optional) Add secure attribute (`Bool`).

`ignore_secure` - (Optional) Ignore secure attribute (`Bool`).

`ttl` - (Optional) be a session cookie. TTL value is in milliseconds (`Int`).

### Port Match No Port Match

Disable matching of ports.

### Query Params Remove All Params

x-displayName: "Remove All Parameters".

### Query Params Retain All Params

x-displayName: "Retain All Parameters".

### Query Params Strip Query Params

Specifies the list of query params to be removed. Not supported.

`query_params` - (Optional) Query params keys to strip while manipulating the HTTP request (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Ref Type App Firewall

A direct reference to an Application Firewall configuration object.

`app_firewall` - (Required) References to an Application Firewall configuration object. See [ref](#ref) below for details.

### Ref Type Disable Waf

Any Application Firewall configuration will not be enforced.

### Ref Type Inherit Waf

Any Application Firewall configuration that was configured on a higher level will be enforced.

### Retry Policy Back Off

10 times the base interval.

`base_interval` - (Optional) Specifies the base interval between retries in milliseconds (`Int`).

`max_interval` - (Optional) to the base_interval if set. The default is 10 times the base_interval. (`Int`).

### Route Action Route Destination

Send request to one of the destination from list of destinations.

`buffer_policy` - (Optional) Route level buffer configuration overrides any configuration at VirtualHost level.. See [Route Destination Buffer Policy ](#route-destination-buffer-policy) below for details.

###### One of the arguments from this list "retract_cluster, do_not_retract_cluster" can be set

`do_not_retract_cluster` - (Optional) configuration. (`Bool`).

`retract_cluster` - (Optional) for route (`Bool`).

`cors_policy` - (Optional) resources from a server at a different origin. See [Route Destination Cors Policy ](#route-destination-cors-policy) below for details.

`csrf_policy` - (Optional) Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.).. See [Route Destination Csrf Policy ](#route-destination-csrf-policy) below for details.

`destinations` - (Required) sent to the cluster specified in the destination. See [Route Destination Destinations ](#route-destination-destinations) below for details.

`endpoint_subsets` - (Optional) upstream cluster which match this metadata will be selected for load balancing (`String`).

`hash_policy` - (Optional) route the request. See [Route Destination Hash Policy ](#route-destination-hash-policy) below for details.

###### One of the arguments from this list "host_rewrite, auto_host_rewrite" must be set

`auto_host_rewrite` - (Optional) of the upstream host chosen by the cluster (`Bool`).

`host_rewrite` - (Optional) Indicates that during forwarding, the host header will be swapped with this value (`String`).

`mirror_policy` - (Optional) useful for logging. For example, *cluster1* becomes *cluster1-shadow*.. See [Route Destination Mirror Policy ](#route-destination-mirror-policy) below for details.

`priority` - (Optional) Also, circuit-breaker configuration at destination cluster is chosen based on the route priority. (`String`).

`retry_policy` - (Optional) Indicates that the route has a retry policy.. See [Route Destination Retry Policy ](#route-destination-retry-policy) below for details.

###### One of the arguments from this list "prefix_rewrite, regex_rewrite" can be set

`prefix_rewrite` - (Optional) while requests to /register/public will be stripped to /public (`String`).

`regex_rewrite` - (Optional) would transform "/service/foo/v1/api" into "/v1/api/instance/foo".. See [Route Destination Rewrite Regex Rewrite ](#route-destination-rewrite-regex-rewrite) below for details.(Deprecated)

`spdy_config` - (Optional) SPDY configuration for each route. See [Route Destination Spdy Config ](#route-destination-spdy-config) below for details.

`timeout` - (Optional) for infinite timeout (`Int`).

`web_socket_config` - (Optional) Websocket configuration for each route. See [Route Destination Web Socket Config ](#route-destination-web-socket-config) below for details.

### Route Action Route Direct Response

Send direct response.

`response_body` - (Optional) response body to send (`String`).

`response_code` - (Optional) response code to send (`Int`).

### Route Action Route Redirect

Send redirect response.

`host_redirect` - (Optional) swap host part of incoming URL in redirect URL (`String`).

`port_redirect` - (Optional) Specify the port value to redirect to a URL with non default port(443) (`Int`).(Deprecated)

`proto_redirect` - (Optional) When incoming-proto option is specified, swapping of protocol is not done. (`String`).

###### One of the arguments from this list "retain_all_params, remove_all_params, replace_params, strip_query_params, all_params" can be set

`all_params` - (Optional) be removed. Default value is false, which means query portion of the URL will NOT be removed (`Bool`).(Deprecated)

`remove_all_params` - (Optional) x-displayName: "Remove All Parameters" (`Bool`).

`replace_params` - (Optional) x-displayName: "Replace All Parameters" (`String`).

`retain_all_params` - (Optional) x-displayName: "Retain All Parameters" (`Bool`).

`strip_query_params` - (Optional) Specifies the list of query params to be removed. Not supported. See [Query Params Strip Query Params ](#query-params-strip-query-params) below for details.(Deprecated)

###### One of the arguments from this list "path_redirect, prefix_rewrite" can be set

`path_redirect` - (Optional) swap path part of incoming URL in redirect URL (`String`).

`prefix_rewrite` - (Optional) This option allows redirect URLs be dynamically created based on the request (`String`).

`response_code` - (Optional) The HTTP status code to use in the redirect response. (`Int`).

### Route Destination Buffer Policy

Route level buffer configuration overrides any configuration at VirtualHost level..

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`max_request_bytes` - (Optional) manager will stop buffering and return a RequestEntityTooLarge (413) response. (`Int`).

`max_request_time` - (Optional) request before returning a RequestTimeout (408) response (`Int`).(Deprecated)

### Route Destination Cors Policy

resources from a server at a different origin.

`allow_credentials` - (Optional) Specifies whether the resource allows credentials (`Bool`).

`allow_headers` - (Optional) Specifies the content for the access-control-allow-headers header (`String`).

`allow_methods` - (Optional) Specifies the content for the access-control-allow-methods header (`String`).

`allow_origin` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`allow_origin_regex` - (Optional) An origin is allowed if either allow_origin or allow_origin_regex match (`String`).

`disabled` - (Optional) The value of this field is ignored for virtual-host (`Bool`).

`expose_headers` - (Optional) Specifies the content for the access-control-expose-headers header (`String`).

`max_age` - (Optional) Specifies the content for the access-control-max-age header (`String`).(Deprecated)

`maximum_age` - (Optional) Maximum permitted value is 86400 seconds (24 hours) (`Int`).

### Route Destination Csrf Policy

Because CSRF attacks specifically target state-changing requests, the policy only acts on the HTTP requests that have state-changing method (PUT,POST, etc.)..

###### One of the arguments from this list "all_load_balancer_domains, custom_domain_list, disabled" must be set

`all_load_balancer_domains` - (Optional) Add All load balancer domains to source origin (allow) list. (`Bool`).

`custom_domain_list` - (Optional) Add one or more domains to source origin (allow) list.. See [Allowed Domains Custom Domain List ](#allowed-domains-custom-domain-list) below for details.

`disabled` - (Optional) Allow all source origin domains. (`Bool`).

### Route Destination Destinations

sent to the cluster specified in the destination.

`cluster` - (Required) does not exist ServiceUnavailable response will be sent. See [ref](#ref) below for details.

`endpoint_subsets` - (Optional) upstream cluster which match this metadata will be selected for load balancing (`String`).

`priority` - (Optional) made active as per the increasing priority. (`Int`).

`weight` - (Optional) sent to the cluster specified in the destination (`Int`).

### Route Destination Hash Policy

route the request.

###### One of the arguments from this list "header_name, cookie, source_ip" must be set

`cookie` - (Optional) Hash based on cookie. See [Policy Specifier Cookie ](#policy-specifier-cookie) below for details.

`header_name` - (Optional) The name or key of the request header that will be used to obtain the hash key (`String`).

`source_ip` - (Optional) Hash based on source IP address (`Bool`).

`terminal` - (Optional) Specify if its a terminal policy (`Bool`).

### Route Destination Mirror Policy

useful for logging. For example, *cluster1* becomes *cluster1-shadow*..

`cluster` - (Required) referred here must be present.. See [ref](#ref) below for details.

`percent` - (Optional) Percentage of requests to be mirrored. See [Mirror Policy Percent ](#mirror-policy-percent) below for details.

### Route Destination Retry Policy

Indicates that the route has a retry policy..

`back_off` - (Optional) 10 times the base interval. See [Retry Policy Back Off ](#retry-policy-back-off) below for details.

`num_retries` - (Optional) is used between each retry (`Int`).

`per_try_timeout` - (Optional) Specifies a non-zero timeout per retry attempt. In milliseconds (`Int`).

`retriable_status_codes` - (Optional) HTTP status codes that should trigger a retry in addition to those specified by retry_on. (`Int`).

`retry_condition` - (Required) (disconnect/reset/read timeout.) (`String`).

`retry_on` - (Optional) matching one defined in retriable_status_codes field (`String`).(Deprecated)

### Route Destination Spdy Config

SPDY configuration for each route.

`use_spdy` - (Optional) a SPDY connection (`Bool`).

### Route Destination Web Socket Config

Websocket configuration for each route.

`idle_timeout` - (Optional) Idle Timeout for Websocket in milli seconds. After timeout, connection will be closed (`Int`).(Deprecated)

`max_connect_attempts` - (Optional) giving up. Default is 1 (`Int`).(Deprecated)

`use_websocket` - (Optional) a WebSocket connection (`Bool`).

### Route Destination Rewrite Regex Rewrite

would transform "/service/foo/v1/api" into "/v1/api/instance/foo"..

`pattern` - (Optional) The regular expression used to find portions of a string that should be replaced. (`String`).

`substitution` - (Optional) substitution operation to produce a new string. (`String`).

### Routes Bot Defense Javascript Injection Inline Mode

Specifies whether bot defense js injection inline mode will be enabled.

`element_selector` - (Required) Element selector to insert into. (`String`).

`insert_content` - (Optional) HTML content to insert. (`String`).

`position` - (Optional) Position of HTML content to be inserted within HTML tag. (`String`).

### Routes Match

route match condition.

`headers` - (Optional) List of (key, value) headers. See [Match Headers ](#match-headers) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`incoming_port` - (Optional) The port on which the request is received. See [Match Incoming Port ](#match-incoming-port) below for details.

`path` - (Optional) URI path of route. See [Match Path ](#match-path) below for details.

`query_params` - (Optional) List of (key, value) query parameters. See [Match Query Params ](#match-query-params) below for details.

### Routes Request Headers To Add

enclosing VirtualHost object level.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "value, secret_value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Routes Response Headers To Add

enclosing VirtualHost object level.

`append` - (Optional) Default value is do not append (`Bool`).

`name` - (Required) Name of the HTTP header. (`String`).

###### One of the arguments from this list "value, secret_value" must be set

`secret_value` - (Optional) Secret Value of the HTTP header.. See [Value Choice Secret Value ](#value-choice-secret-value) below for details.

`value` - (Optional) Value of the HTTP header. (`String`).

### Routes Service Policy

service policy configuration at route level which overrides configuration at virtual-host level.

###### One of the arguments from this list "context_extensions, disable" can be set

`context_extensions` - (Optional) sending additional information to the external authorization server.. See [Service Policy Choice Context Extensions ](#service-policy-choice-context-extensions) below for details.(Deprecated)

`disable` - (Optional) disable service policy at route level, if it is configured at virtual-host level (`Bool`).

### Routes Waf Type

waf_type specified at route level overrides waf configuration at VirtualHost level.

###### One of the arguments from this list "app_firewall, disable_waf, inherit_waf" can be set

`app_firewall` - (Optional) A direct reference to an Application Firewall configuration object. See [Ref Type App Firewall ](#ref-type-app-firewall) below for details.

`disable_waf` - (Optional) Any Application Firewall configuration will not be enforced (`Bool`).

`inherit_waf` - (Optional) Any Application Firewall configuration that was configured on a higher level will be enforced (`Bool`).

### Samesite Ignore Samesite

Ignore Samesite attribute.

### Samesite Samesite Lax

Add Samesite attribute with Lax. Means that the cookie is not sent on cross-site requests.

### Samesite Samesite None

Add Samesite attribute with None. Means that the browser sends the cookie with both cross-site and same-site requests.

### Samesite Samesite Strict

Add Samesite attribute with Strict. Means that the browser sends the cookie only for same-site requests.

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Secret Info Oneof Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Secret Info Oneof Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

### Secret Value Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secure Add Secure

Add secure attribute.

### Secure Ignore Secure

Ignore secure attribute.

### Service Policy Choice Context Extensions

sending additional information to the external authorization server..

`context_extensions` - (Optional) provide extra context for the external authorization server on specific virtual hosts or routes. (`String`).

### Value Choice Secret Value

Secret Value of the HTTP header..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Secret Value Blindfold Secret Info Internal ](#secret-value-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

Attribute Reference
-------------------

-	`id` - This is the id of the configured route.
