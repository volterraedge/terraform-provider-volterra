---

page_title: "Volterra: workload"

description: "The workload allows CRUD of Workload resource on Volterra SaaS"
-----------------------------------------------------------------------------

Resource volterra_workload
==========================

The Workload allows CRUD of Workload resource on Volterra SaaS

~> **Note:** Please refer to [Workload API docs](https://docs.cloud.f5.com/docs-v2/api/views-workload) to learn more

Example Usage
-------------

```hcl
resource "volterra_workload" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "job service simple_service stateful_service" must be set

  simple_service {
    // One of the arguments from this list "do_not_advertise service_port simple_advertise" must be set

    do_not_advertise = true

    configuration {
      parameters {
        // One of the arguments from this list "env_var file" must be set

        file {
          data = "value"

          mount {
            mode = "mode"

            mount_path = "value"

            sub_path = "value"
          }

          name = "value"

          volume_name = "value"
        }
      }
    }

    container {
      args = ["value"]

      command = ["value"]

      // One of the arguments from this list "custom_flavor default_flavor flavor" can be set

      flavor = "flavor"
      image {
        name = "quay.io/etcd:0.13"

        pull_policy = "pull_policy"

        // One of the arguments from this list "container_registry public" must be set

        public = true
      }
      init_container = true
      liveness_check {
        // One of the arguments from this list "exec_health_check http_health_check tcp_health_check" must be set

        http_health_check {
          headers = {
            "key1" = "value1"
          }

          host_header = "one.volterra.com"

          path = "/"

          port {
            // One of the arguments from this list "name num" must be set

            num = "num"
          }
        }

        healthy_threshold = "2"

        initial_delay = "30"

        interval = "10"

        timeout = "3"

        unhealthy_threshold = "5"
      }
      name = "value"
      readiness_check {
        // One of the arguments from this list "exec_health_check http_health_check tcp_health_check" must be set

        http_health_check {
          headers = {
            "key1" = "value1"
          }

          host_header = "one.volterra.com"

          path = "/"

          port {
            // One of the arguments from this list "name num" must be set

            num = "num"
          }
        }

        healthy_threshold = "2"

        initial_delay = "30"

        interval = "10"

        timeout = "3"

        unhealthy_threshold = "5"
      }
    }

    // One of the arguments from this list "disabled enabled" must be set

    disabled = true
    scale_to_zero = true
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

###### One of the arguments from this list "job, service, simple_service, stateful_service" must be set

`job` - (Optional) like report generation, billing, parallel data processing, ETL processing, etc.. See [Workload Choice Job ](#workload-choice-job) below for details.

`service` - (Optional) application servers, traditional SQL databases, etc.. See [Workload Choice Service ](#workload-choice-service) below for details.

`simple_service` - (Optional) on default VIP. See [Workload Choice Simple Service ](#workload-choice-simple-service) below for details.

`stateful_service` - (Optional) applications like cassandra, mongodb, redis, etc.. See [Workload Choice Stateful Service ](#workload-choice-stateful-service) below for details.

### Advertise Choice Advertise Custom

Advertise the workload via load balancer on specific sites.

`advertise_where` - (Required) Where should this load balancer be available. See [Advertise Custom Advertise Where ](#advertise-custom-advertise-where) below for details.

`ports` - (Required) Ports to advertise. See [Advertise Custom Ports ](#advertise-custom-ports) below for details.

### Advertise Choice Advertise In Cluster

Advertise the workload locally in the cluster where it is deployed.

###### One of the arguments from this list "multi_ports, port" must be set

`multi_ports` - (Optional) Multi Ports. See [Port Choice Multi Ports ](#port-choice-multi-ports) below for details.

`port` - (Optional) Single Port. See [Port Choice Port ](#port-choice-port) below for details.

### Advertise Choice Advertise On Public

Advertise the workload via load balancer on Internet with default VIP.

###### One of the arguments from this list "multi_ports, port" must be set

`multi_ports` - (Optional) Multiple Ports. See [Advertise Choice Multi Ports ](#advertise-choice-multi-ports) below for details.

`port` - (Optional) Single Port. See [Advertise Choice Port ](#advertise-choice-port) below for details.

### Advertise Choice Do Not Advertise

Do not advertise the workload.

### Advertise Choice Http Loadbalancer

HTTP/HTTPS Load balancer.

`domains` - (Required) Domains also indicate the list of names for which DNS resolution will be done by VER (`String`).

###### One of the arguments from this list "http, https, https_auto_cert" must be set

`http` - (Optional) HTTP Load balancer.. See [Loadbalancer Type Http ](#loadbalancer-type-http) below for details.

`https` - (Optional) User is responsible for managing DNS to this Load Balancer.. See [Loadbalancer Type Https ](#loadbalancer-type-https) below for details.

`https_auto_cert` - (Optional) DNS records will be managed by Volterra.. See [Loadbalancer Type Https Auto Cert ](#loadbalancer-type-https-auto-cert) below for details.

###### One of the arguments from this list "default_route, specific_routes" must be set

`default_route` - (Optional) Default route matching all APIs. See [Route Choice Default Route ](#route-choice-default-route) below for details.

`specific_routes` - (Optional) Routes matching specific APIs. See [Route Choice Specific Routes ](#route-choice-specific-routes) below for details.

### Advertise Choice Multi Ports

Multiple Ports.

`ports` - (Required) Ports to advertise. See [Multi Ports Ports ](#multi-ports-ports) below for details.

### Advertise Choice Port

Single Port.

###### One of the arguments from this list "http_loadbalancer, tcp_loadbalancer" must be set

`http_loadbalancer` - (Optional) HTTP/HTTPS Load balancer. See [Advertise Choice Http Loadbalancer ](#advertise-choice-http-loadbalancer) below for details.

`tcp_loadbalancer` - (Optional) TCP Load balancer. See [Advertise Choice Tcp Loadbalancer ](#advertise-choice-tcp-loadbalancer) below for details.

`port` - (Required) Port information. See [Port Port ](#port-port) below for details.

### Advertise Choice Simple Advertise

Domains and service port to advertise on Internet via HTTP loadbalancer using port 80.

`domains` - (Required) Domains also indicate the list of names for which DNS resolution will be automatically resolved to IP addresses by the system. (`String`).

`service_port` - (Required) Service port to advertise on Internet via HTTP loadbalancer using port 80 (`Int`).

### Advertise Choice Tcp Loadbalancer

TCP Load balancer.

`domains` - (Optional) Domains also indicate the list of names for which DNS resolution will be done by VER (`String`).

`with_sni` - (Optional) Set to true to enable TCP loadbalancer with SNI (`Bool`).

### Advertise Custom Advertise Where

Where should this load balancer be available.

###### One of the arguments from this list "site, virtual_site, vk8s_service" must be set

`site` - (Optional) Advertise on a customer edge site and a given network.. See [Choice Site ](#choice-site) below for details.

`virtual_site` - (Optional) Advertise on a customer edge virtual site and a given network.. See [Choice Virtual Site ](#choice-virtual-site) below for details.

`vk8s_service` - (Optional) Advertise on vK8s Service Network on RE.. See [Choice Vk8s Service ](#choice-vk8s-service) below for details.

### Advertise Custom Ports

Ports to advertise.

###### One of the arguments from this list "http_loadbalancer, tcp_loadbalancer" must be set

`http_loadbalancer` - (Optional) HTTP/HTTPS Load balancer. See [Advertise Choice Http Loadbalancer ](#advertise-choice-http-loadbalancer) below for details.

`tcp_loadbalancer` - (Optional) TCP Load balancer. See [Advertise Choice Tcp Loadbalancer ](#advertise-choice-tcp-loadbalancer) below for details.

`port` - (Required) Port information. See [Ports Port ](#ports-port) below for details.

### Choice Custom Route Object

A custom route uses a route object created outside of this view..

`route_ref` - (Optional) Reference to a custom route object. See [ref](#ref) below for details.

### Choice Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Choice Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Choice Direct Response Route

A direct response route matches on path and/or HTTP method and responds directly to the matching traffic.

`headers` - (Optional) List of (key, value) headers. See [Direct Response Route Headers ](#direct-response-route-headers) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`incoming_port` - (Optional) The port on which the request is received. See [Direct Response Route Incoming Port ](#direct-response-route-incoming-port) below for details.

`path` - (Required) URI path of route. See [Direct Response Route Path ](#direct-response-route-path) below for details.

`route_direct_response` - (Optional) Send direct response. See [Direct Response Route Route Direct Response ](#direct-response-route-route-direct-response) below for details.

### Choice Env Var

Expose as environment variable.

`name` - (Optional) Name of Environment Variable. (`String`).

`value` - (Optional) Value of Environment Variable. (`String`).

### Choice File

Configuration file.

`data` - (Optional) File data (`String`).

`mount` - (Required) Volume Mount information. See [File Mount ](#file-mount) below for details.

`name` - (Required) Name of the file (`String`).

`volume_name` - (Required) Name of the Volume (`String`).

### Choice Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Choice Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Choice Redirect Route

A redirect route matches on path and/or HTTP method and redirects the matching traffic to a different URL.

`headers` - (Optional) List of (key, value) headers. See [Redirect Route Headers ](#redirect-route-headers) below for details.

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`incoming_port` - (Optional) The port on which the request is received. See [Redirect Route Incoming Port ](#redirect-route-incoming-port) below for details.

`path` - (Required) URI path of route. See [Redirect Route Path ](#redirect-route-path) below for details.

`route_redirect` - (Optional) Send redirect response. See [Redirect Route Route Redirect ](#redirect-route-route-redirect) below for details.

### Choice Simple Route

A simple route matches on path and/or HTTP method and forwards the matching traffic to the workload.

###### One of the arguments from this list "auto_host_rewrite, disable_host_rewrite, host_rewrite" must be set

`auto_host_rewrite` - (Optional) Host header will be swapped with hostname of upstream host chosen by the cluster (`Bool`).

`disable_host_rewrite` - (Optional) Host header is not modified (`Bool`).

`host_rewrite` - (Optional) Host header will be swapped with this value (`String`).

`http_method` - (Optional) The name of the HTTP Method (GET, PUT, POST, etc) (`String`).

`path` - (Required) URI path of route. See [Simple Route Path ](#simple-route-path) below for details.

### Choice Site

Advertise on a customer edge site and a given network..

`ip` - (Optional) Use given IP address as VIP on the site (`String`).

`ipv6` - (Optional) Use given IPv6 address as VIP on the site (`String`).

`network` - (Required) By default VIP chosen as ip address of primary network interface in the network (`String`).

`site` - (Required) Reference to site object. See [ref](#ref) below for details.

### Choice Virtual Site

Advertise on a customer edge virtual site and a given network..

`network` - (Required) IP address of primary network interface in the network (`String`).

`virtual_site` - (Required) Reference to virtual site object. See [ref](#ref) below for details.

### Choice Vk8s Service

Advertise on vK8s Service Network on RE..

###### One of the arguments from this list "site, virtual_site" must be set

`site` - (Optional) Reference to site object. See [ref](#ref) below for details.

`virtual_site` - (Optional) Reference to virtual site object. See [ref](#ref) below for details.

### Class Name Choice Default

Default class.

### Configuration Parameters

Parameters for the workload.

###### One of the arguments from this list "env_var, file" must be set

`env_var` - (Optional) Expose as environment variable. See [Choice Env Var ](#choice-env-var) below for details.

`file` - (Optional) Configuration file. See [Choice File ](#choice-file) below for details.

### Container Image

Container image information.

`name` - (Required) Docker public registry is assumed. If tag is not specified, latest is assumed. (`String`).

`pull_policy` - (Optional) Pull Policy determines the policy to use for pulling the image prior to starting the service (`String`).

###### One of the arguments from this list "container_registry, public" must be set

`container_registry` - (Optional) Quay Enterprise, Docker Hub, GCP Container Registry. See [ref](#ref) below for details.

`public` - (Optional) Pull image from a public registry like Docker (`Bool`).

### Container Liveness Check

Container will be restarted if the health check fails..

###### One of the arguments from this list "exec_health_check, http_health_check, tcp_health_check" must be set

`exec_health_check` - (Optional) Exit status of 0 is treated as live/healthy and non-zero is unhealthy. See [Health Check Choice Exec Health Check ](#health-check-choice-exec-health-check) below for details.

`http_health_check` - (Optional) Specifies details for HTTP GET health check requests. See [Health Check Choice Http Health Check ](#health-check-choice-http-health-check) below for details.

`tcp_health_check` - (Optional) Specifies port to use for TCP connection for health checks. See [Health Check Choice Tcp Health Check ](#health-check-choice-tcp-health-check) below for details.

`healthy_threshold` - (Required) required to mark a container healthy. (`Int`).

`initial_delay` - (Optional) Number of seconds after the container has started before health checks are initiated (`Int`).

`interval` - (Required) Time interval in seconds between two health check requests. (`Int`).

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

`unhealthy_threshold` - (Required) the number of unhealthy health checks required before a container is marked unhealthy. (`Int`).

### Container Readiness Check

Container will be removed from service endpoints if the health check fails..

###### One of the arguments from this list "exec_health_check, http_health_check, tcp_health_check" must be set

`exec_health_check` - (Optional) Exit status of 0 is treated as live/healthy and non-zero is unhealthy. See [Health Check Choice Exec Health Check ](#health-check-choice-exec-health-check) below for details.

`http_health_check` - (Optional) Specifies details for HTTP GET health check requests. See [Health Check Choice Http Health Check ](#health-check-choice-http-health-check) below for details.

`tcp_health_check` - (Optional) Specifies port to use for TCP connection for health checks. See [Health Check Choice Tcp Health Check ](#health-check-choice-tcp-health-check) below for details.

`healthy_threshold` - (Required) required to mark a container healthy. (`Int`).

`initial_delay` - (Optional) Number of seconds after the container has started before health checks are initiated (`Int`).

`interval` - (Required) Time interval in seconds between two health check requests. (`Int`).

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

`unhealthy_threshold` - (Required) the number of unhealthy health checks required before a container is marked unhealthy. (`Int`).

### Containers Image

Container image information.

`name` - (Required) Docker public registry is assumed. If tag is not specified, latest is assumed. (`String`).

`pull_policy` - (Optional) Pull Policy determines the policy to use for pulling the image prior to starting the service (`String`).

###### One of the arguments from this list "container_registry, public" must be set

`container_registry` - (Optional) Quay Enterprise, Docker Hub, GCP Container Registry. See [ref](#ref) below for details.

`public` - (Optional) Pull image from a public registry like Docker (`Bool`).

### Containers Liveness Check

Container will be restarted if the health check fails..

###### One of the arguments from this list "exec_health_check, http_health_check, tcp_health_check" must be set

`exec_health_check` - (Optional) Exit status of 0 is treated as live/healthy and non-zero is unhealthy. See [Health Check Choice Exec Health Check ](#health-check-choice-exec-health-check) below for details.

`http_health_check` - (Optional) Specifies details for HTTP GET health check requests. See [Health Check Choice Http Health Check ](#health-check-choice-http-health-check) below for details.

`tcp_health_check` - (Optional) Specifies port to use for TCP connection for health checks. See [Health Check Choice Tcp Health Check ](#health-check-choice-tcp-health-check) below for details.

`healthy_threshold` - (Required) required to mark a container healthy. (`Int`).

`initial_delay` - (Optional) Number of seconds after the container has started before health checks are initiated (`Int`).

`interval` - (Required) Time interval in seconds between two health check requests. (`Int`).

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

`unhealthy_threshold` - (Required) the number of unhealthy health checks required before a container is marked unhealthy. (`Int`).

### Containers Readiness Check

Container will be removed from service endpoints if the health check fails..

###### One of the arguments from this list "exec_health_check, http_health_check, tcp_health_check" must be set

`exec_health_check` - (Optional) Exit status of 0 is treated as live/healthy and non-zero is unhealthy. See [Health Check Choice Exec Health Check ](#health-check-choice-exec-health-check) below for details.

`http_health_check` - (Optional) Specifies details for HTTP GET health check requests. See [Health Check Choice Http Health Check ](#health-check-choice-http-health-check) below for details.

`tcp_health_check` - (Optional) Specifies port to use for TCP connection for health checks. See [Health Check Choice Tcp Health Check ](#health-check-choice-tcp-health-check) below for details.

`healthy_threshold` - (Required) required to mark a container healthy. (`Int`).

`initial_delay` - (Optional) Number of seconds after the container has started before health checks are initiated (`Int`).

`interval` - (Required) Time interval in seconds between two health check requests. (`Int`).

`timeout` - (Required) health check attempt will be considered a failure. (`Int`).

`unhealthy_threshold` - (Required) the number of unhealthy health checks required before a container is marked unhealthy. (`Int`).

### Crl Choice No Crl

Client certificate revocation status is not verified.

### Default Lb Choice Default Loadbalancer

x-displayName: "Yes".

### Default Lb Choice Non Default Loadbalancer

x-displayName: "No".

### Deploy Choice All Res

Deploy the workload on all Regional Edges.

### Deploy Choice Default Virtual Sites

Deploy the workload on virtual sites referred by the Virtual Kubernetes.

### Deploy Choice Deploy Ce Sites

Deploy the workload on specific Customer sites.

`site` - (Required) Which customer sites should this workload be deployed. See [ref](#ref) below for details.

### Deploy Choice Deploy Ce Virtual Sites

Deploy the workload on specific Customer virtual sites.

`virtual_site` - (Required) Which customer virtual sites should this workload be deployed. See [ref](#ref) below for details.

### Deploy Choice Deploy Re Sites

Deploy the workload on specific Regional Edge sites.

`site` - (Required) Which regional edge sites should this workload be deployed. See [ref](#ref) below for details.

### Deploy Choice Deploy Re Virtual Sites

Deploy the workload on specific Regional Edge virtual sites.

`virtual_site` - (Required) Which regional edge virtual sites should this workload be deployed. See [ref](#ref) below for details.

### Direct Response Route Headers

List of (key, value) headers.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, presence, regex" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Direct Response Route Incoming Port

The port on which the request is received.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Direct Response Route Path

URI path of route.

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Direct Response Route Route Direct Response

Send direct response.

### Empty Dir Mount

Volume mount associated with the empty directory.

`mode` - (Optional) Mode of the volume (`String`).

`mount_path` - (Required) Path within the workload container at which the volume should be mounted. Must not contain ':'. (`String`).

`sub_path` - (Optional) Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's root). (`String`).

### Enabled Persistent Volume

Volume containing the Persistent Storage for the workload.

`mount` - (Required) Volume mount associated with the persistent storage. See [Persistent Volume Mount ](#persistent-volume-mount) below for details.

`storage` - (Required) Persistent storage configuration. See [Persistent Volume Storage ](#persistent-volume-storage) below for details.

### File Mount

Volume Mount information.

`mode` - (Optional) Mode of the volume (`String`).

`mount_path` - (Required) Path within the workload container at which the volume should be mounted. Must not contain ':'. (`String`).

`sub_path` - (Optional) Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's root). (`String`).

### Flavor Choice Default Flavor

Container resources based on default workload flavor of the Virtual Kubernetes.

### Header Transformation Choice Default Header Transformation

Normalize the headers to lower case.

### Header Transformation Choice Legacy Header Transformation

Use old header transformation if configured earlier.

### Header Transformation Choice Preserve Case Header Transformation

Preserves the original case of headers without any modifications..

### Header Transformation Choice Proper Case Header Transformation

For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are”.

### Health Check Choice Exec Health Check

Exit status of 0 is treated as live/healthy and non-zero is unhealthy.

`command` - (Required) a shell, you need to explicitly call out to that shell. (`String`).

### Health Check Choice Http Health Check

Specifies details for HTTP GET health check requests.

`headers` - (Optional) health checked container. This is a list of key-value pairs. (`String`).

`host_header` - (Optional) The value of the host header in the HTTP health check request. (`String`).

`path` - (Required) Path to access on the HTTP server. (`String`).

`port` - (Required) Name or number of the port to access on the container.. See [Http Health Check Port ](#http-health-check-port) below for details.

### Health Check Choice Tcp Health Check

Specifies port to use for TCP connection for health checks.

`port` - (Required) Number or name of the port to access on the container.. See [Tcp Health Check Port ](#tcp-health-check-port) below for details.

### Host Path Mount

Volume mount associated with the host path.

`mode` - (Optional) Mode of the volume (`String`).

`mount_path` - (Required) Path within the workload container at which the volume should be mounted. Must not contain ':'. (`String`).

`sub_path` - (Optional) Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's root). (`String`).

### Host Rewrite Params Auto Host Rewrite

Host header will be swapped with hostname of upstream host chosen by the cluster.

### Host Rewrite Params Disable Host Rewrite

Host header is not modified.

### Http Health Check Port

Name or number of the port to access on the container..

###### One of the arguments from this list "name, num" must be set

`name` - (Optional) Port Name (`String`).

`num` - (Optional) Port number (`Int`).

### Https Coalescing Options

Options for coalescing TLS for multiple HTTPS Load Balancers.

### Https Header Transformation Type

Header transformation options for response headers to the client.

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Https Http Protocol Options

HTTP protocol configuration options for downstream connections..

### Https Auto Cert Coalescing Options

Options for coalescing TLS for multiple HTTPS Load Balancers.

### Https Auto Cert Header Transformation Type

Header transformation options for response headers to the client.

###### One of the arguments from this list "default_header_transformation, legacy_header_transformation, preserve_case_header_transformation, proper_case_header_transformation" must be set

`default_header_transformation` - (Optional) Normalize the headers to lower case (`Bool`).

`legacy_header_transformation` - (Optional) Use old header transformation if configured earlier (`Bool`).

`preserve_case_header_transformation` - (Optional) Preserves the original case of headers without any modifications. (`Bool`).

`proper_case_header_transformation` - (Optional) For example, “content-type” becomes “Content-Type”, and “foo$b#$are” becomes “Foo$B#$Are” (`Bool`).

### Https Auto Cert Http Protocol Options

HTTP protocol configuration options for downstream connections..

### Https Auto Cert Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Ip Family Choice Dual V4 Preferred

Dual-stack IPv4 Preferred.

### Ip Family Choice Dual V6 Preferred

Dual-stack IPv6 Preferred.

### Ip Family Choice V4

IPV4 Family.

### Ip Family Choice V6

IPV6 Family.

### Job Configuration

Configuration parameters for the job.

`parameters` - (Optional) Parameters for the workload. See [Configuration Parameters ](#configuration-parameters) below for details.

### Job Containers

Containers to use for the job.

`args` - (Optional) Arguments to the entrypoint. Overrides the docker image's CMD (`String`).

`command` - (Optional) Command to execute. Overrides the docker image's ENTRYPOINT (`String`).

###### One of the arguments from this list "custom_flavor, default_flavor, flavor" can be set

`custom_flavor` - (Optional) Container resources based on Workload Flavor. See [ref](#ref) below for details.

`default_flavor` - (Optional) Container resources based on default workload flavor of the Virtual Kubernetes (`Bool`).

`flavor` - (Optional) Container resources based on system defined flavor (`String`).

`image` - (Required) Container image information. See [Containers Image ](#containers-image) below for details.

`init_container` - (Optional) Specialized container that runs before application container and runs to completion (`Bool`).

`liveness_check` - (Optional) Container will be restarted if the health check fails.. See [Containers Liveness Check ](#containers-liveness-check) below for details.

`name` - (Required) Name of the container (`String`).

`readiness_check` - (Optional) Container will be removed from service endpoints if the health check fails.. See [Containers Readiness Check ](#containers-readiness-check) below for details.

### Job Deploy Options

Deploy options for the job are used to decide where and how to deploy the job.

###### One of the arguments from this list "all_res, default_virtual_sites, deploy_ce_sites, deploy_ce_virtual_sites, deploy_re_sites, deploy_re_virtual_sites" must be set

`all_res` - (Optional) Deploy the workload on all Regional Edges (`Bool`).

`default_virtual_sites` - (Optional) Deploy the workload on virtual sites referred by the Virtual Kubernetes (`Bool`).

`deploy_ce_sites` - (Optional) Deploy the workload on specific Customer sites. See [Deploy Choice Deploy Ce Sites ](#deploy-choice-deploy-ce-sites) below for details.

`deploy_ce_virtual_sites` - (Optional) Deploy the workload on specific Customer virtual sites. See [Deploy Choice Deploy Ce Virtual Sites ](#deploy-choice-deploy-ce-virtual-sites) below for details.

`deploy_re_sites` - (Optional) Deploy the workload on specific Regional Edge sites. See [Deploy Choice Deploy Re Sites ](#deploy-choice-deploy-re-sites) below for details.

`deploy_re_virtual_sites` - (Optional) Deploy the workload on specific Regional Edge virtual sites. See [Deploy Choice Deploy Re Virtual Sites ](#deploy-choice-deploy-re-virtual-sites) below for details.

### Job Volumes

Volumes for the job.

`name` - (Optional) Name of the volume (`String`).

###### One of the arguments from this list "empty_dir, host_path, persistent_volume" must be set

`empty_dir` - (Optional) Volume containing a temporary directory whose lifetime is the same as a replica of a workload. See [Volume Choice Empty Dir ](#volume-choice-empty-dir) below for details.

`host_path` - (Optional) Volume containing a host mapped path into the workload. See [Volume Choice Host Path ](#volume-choice-host-path) below for details.

`persistent_volume` - (Optional) Volume containing the Persistent Storage for the workload. See [Volume Choice Persistent Volume ](#volume-choice-persistent-volume) below for details.

### Loadbalancer Type Http

HTTP Load balancer..

`dns_volterra_managed` - (Optional) or a DNS CNAME record should be created in your DNS provider's portal. (`Bool`).

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) HTTP port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

### Loadbalancer Type Https

User is responsible for managing DNS to this Load Balancer..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`coalescing_options` - (Optional) Options for coalescing TLS for multiple HTTPS Load Balancers. See [Https Coalescing Options ](#https-coalescing-options) below for details.

`connection_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 2 minutes. (`Int`).

###### One of the arguments from this list "default_loadbalancer, non_default_loadbalancer" can be set

`default_loadbalancer` - (Optional) x-displayName: "Yes" (`Bool`).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (`Bool`).

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Https Header Transformation Type ](#https-header-transformation-type) below for details.(Deprecated)

`http_protocol_options` - (Optional) HTTP protocol configuration options for downstream connections.. See [Https Http Protocol Options ](#https-http-protocol-options) below for details.

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" must be set

`disable_path_normalize` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_path_normalize` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "port, port_ranges" must be set

`port` - (Optional) HTTPS port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

###### One of the arguments from this list "append_server_name, default_header, pass_through, server_name" can be set

`append_server_name` - (Optional) If header value is already present, it is not overwritten and passed as-is. (`String`).

`default_header` - (Optional) Response header name is “server” and value is “volt-adc” (`Bool`).

`pass_through` - (Optional) Pass existing server header as is. If server header is absent, a new header is not appended. (`Bool`).

`server_name` - (Optional) This will overwrite existing values, if any, for the server header. (`String`).

###### One of the arguments from this list "tls_cert_params, tls_parameters" must be set

`tls_cert_params` - (Optional) Select/Add one or more TLS Certificate objects to associate with this Load Balancer. See [Tls Certificates Choice Tls Cert Params ](#tls-certificates-choice-tls-cert-params) below for details.

`tls_parameters` - (Optional) Upload a TLS certificate covering all domain names for this Load Balancer. See [Tls Certificates Choice Tls Parameters ](#tls-certificates-choice-tls-parameters) below for details.

### Loadbalancer Type Https Auto Cert

DNS records will be managed by Volterra..

`add_hsts` - (Optional) Add HTTP Strict-Transport-Security response header (`Bool`).

`coalescing_options` - (Optional) Options for coalescing TLS for multiple HTTPS Load Balancers. See [Https Auto Cert Coalescing Options ](#https-auto-cert-coalescing-options) below for details.

`connection_idle_timeout` - (Optional) This is specified in milliseconds. The default value is 2 minutes. (`Int`).

###### One of the arguments from this list "default_loadbalancer, non_default_loadbalancer" can be set

`default_loadbalancer` - (Optional) For traffic terminating at this load balancer, the certificate associated with the first configured domain will be used for TLS termination. (`Bool`).

`non_default_loadbalancer` - (Optional) x-displayName: "No" (`Bool`).

`header_transformation_type` - (Optional) Header transformation options for response headers to the client. See [Https Auto Cert Header Transformation Type ](#https-auto-cert-header-transformation-type) below for details.(Deprecated)

`http_protocol_options` - (Optional) HTTP protocol configuration options for downstream connections.. See [Https Auto Cert Http Protocol Options ](#https-auto-cert-http-protocol-options) below for details.

`http_redirect` - (Optional) Redirect HTTP traffic to HTTPS (`Bool`).

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

###### One of the arguments from this list "disable_path_normalize, enable_path_normalize" must be set

`disable_path_normalize` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_path_normalize` - (Optional) x-displayName: "Enable" (`Bool`).

###### One of the arguments from this list "port, port_ranges" can be set

`port` - (Optional) HTTPS port to Listen. (`Int`).

`port_ranges` - (Optional) Each port range consists of a single port or two ports separated by "-". (`String`).

###### One of the arguments from this list "append_server_name, default_header, pass_through, server_name" can be set

`append_server_name` - (Optional) If header value is already present, it is not overwritten and passed as-is. (`String`).

`default_header` - (Optional) Response header name is “server” and value is “volt-adc” (`Bool`).

`pass_through` - (Optional) Pass existing server header as is. If server header is absent, a new header is not appended. (`Bool`).

`server_name` - (Optional) This will overwrite existing values, if any, for the server header. (`String`).

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Https Auto Cert Tls Config ](#https-auto-cert-tls-config) below for details.

### Mtls Choice No Mtls

x-displayName: "Disable".

### Mtls Choice Use Mtls

x-displayName: "Enable".

`client_certificate_optional` - (Optional) the connection will be accepted. (`Bool`).

###### One of the arguments from this list "crl, no_crl" can be set

`crl` - (Optional) Specify the CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

`no_crl` - (Optional) Client certificate revocation status is not verified (`Bool`).

###### One of the arguments from this list "trusted_ca, trusted_ca_url" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Load Balancer. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Load Balancer (`String`).

###### One of the arguments from this list "xfcc_disabled, xfcc_options" can be set

`xfcc_disabled` - (Optional) No X-Forwarded-Client-Cert header will be added (`Bool`).

`xfcc_options` - (Optional) X-Forwarded-Client-Cert header will be added with the configured fields. See [Xfcc Header Xfcc Options ](#xfcc-header-xfcc-options) below for details.

### Multi Ports Ports

Ports to advertise.

###### One of the arguments from this list "http_loadbalancer, tcp_loadbalancer" must be set

`http_loadbalancer` - (Optional) HTTP/HTTPS Load balancer. See [Advertise Choice Http Loadbalancer ](#advertise-choice-http-loadbalancer) below for details.

`tcp_loadbalancer` - (Optional) TCP Load balancer. See [Advertise Choice Tcp Loadbalancer ](#advertise-choice-tcp-loadbalancer) below for details.

`port` - (Required) Port information. See [Ports Port ](#ports-port) below for details.

### Multi Ports Ports

Ports to advertise.

`info` - (Required) Information of the Port. See [Ports Info ](#ports-info) below for details.

`name` - (Required) Name of the Port (`String`).

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Path Normalize Choice Disable Path Normalize

x-displayName: "Disable".

### Path Normalize Choice Enable Path Normalize

x-displayName: "Enable".

### Persistence Choice Disabled

Disabled.

### Persistence Choice Enabled

x-displayName: "Enabled".

`name` - (Required) Name of the volume (`String`).

`persistent_volume` - (Required) Volume containing the Persistent Storage for the workload. See [Enabled Persistent Volume ](#enabled-persistent-volume) below for details.

### Persistent Volume Mount

Volume mount associated with the persistent storage.

`mode` - (Optional) Mode of the volume (`String`).

`mount_path` - (Required) Path within the workload container at which the volume should be mounted. Must not contain ':'. (`String`).

`sub_path` - (Optional) Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's root). (`String`).

### Persistent Volume Storage

Persistent storage configuration.

`access_mode` - (Optional) AccessMode is used to configure access mode for the persistent storage (`String`).

###### One of the arguments from this list "class_name, default" must be set

`class_name` - (Optional) Use the specified class name (`String`).

`default` - (Optional) Default class (`Bool`).

`storage_size` - (Required) x-required (`Float`).

### Persistent Volumes Persistent Volume

Volume containing the Persistent Storage for the workload.

`mount` - (Required) Volume mount associated with the persistent storage. See [Persistent Volume Mount ](#persistent-volume-mount) below for details.

`storage` - (Required) Persistent storage configuration. See [Persistent Volume Storage ](#persistent-volume-storage) below for details.

### Port Info

Information of the Port.

`port` - (Required) Port the workload can be reached on (`Int`).

`protocol` - (Optional) Application protocol used by the workload (`String`).

###### One of the arguments from this list "same_as_port, target_port" must be set

`same_as_port` - (Optional) Same as Port (`Bool`).

`target_port` - (Optional) Port the workload is listening on (`Int`).

### Port Port

Port information.

`info` - (Required) Information of the Port. See [Port Info ](#port-info) below for details.

### Port Choice Multi Ports

Multi Ports.

`ports` - (Required) Ports to advertise. See [Multi Ports Ports ](#multi-ports-ports) below for details.

### Port Choice Port

Single Port.

`info` - (Required) Information of the Port. See [Port Info ](#port-info) below for details.

### Port Match No Port Match

Disable matching of ports.

### Ports Info

Information of the Port.

`port` - (Required) Port the workload can be reached on (`Int`).

`protocol` - (Optional) Application protocol used by the workload (`String`).

###### One of the arguments from this list "same_as_port, target_port" must be set

`same_as_port` - (Optional) Same as Port (`Bool`).

`target_port` - (Optional) Port the workload is listening on (`Int`).

### Ports Port

Port information.

`info` - (Required) Information of the Port. See [Port Info ](#port-info) below for details.

`name` - (Required) Name of the Port (`String`).

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Redirect Route Headers

List of (key, value) headers.

`invert_match` - (Optional) Invert the result of the match to detect missing header or non-matching value (`Bool`).

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, presence, regex" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`presence` - (Optional) If true, check for presence of header (`Bool`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Redirect Route Incoming Port

The port on which the request is received.

###### One of the arguments from this list "no_port_match, port, port_ranges" can be set

`no_port_match` - (Optional) Disable matching of ports (`Bool`).

`port` - (Optional) Exact Port to match (`Int`).

`port_ranges` - (Optional) Port range to match (`String`).

### Redirect Route Path

URI path of route.

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Redirect Route Route Redirect

Send redirect response.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Registry Choice Public

Pull image from a public registry like Docker.

### Route Choice Default Route

Default route matching all APIs.

###### One of the arguments from this list "auto_host_rewrite, disable_host_rewrite, host_rewrite" must be set

`auto_host_rewrite` - (Optional) Host header will be swapped with hostname of upstream host chosen by the cluster (`Bool`).

`disable_host_rewrite` - (Optional) Host header is not modified (`Bool`).

`host_rewrite` - (Optional) Host header will be swapped with this value (`String`).

### Route Choice Specific Routes

Routes matching specific APIs.

`routes` - (Optional) Routes for this loadbalancer. See [Specific Routes Routes ](#specific-routes-routes) below for details.

### Scaling Choice Scale To Zero

Set number of replicas of service to spawn per site to zero.

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

### Server Header Choice Default Header

Response header name is “server” and value is “volt-adc”.

### Server Header Choice Pass Through

Pass existing server header as is. If server header is absent, a new header is not appended..

### Service Advertise Options

Advertise options for the service control how the service is advertised.

###### One of the arguments from this list "advertise_custom, advertise_in_cluster, advertise_on_public, do_not_advertise" must be set

`advertise_custom` - (Optional) Advertise the workload via load balancer on specific sites. See [Advertise Choice Advertise Custom ](#advertise-choice-advertise-custom) below for details.

`advertise_in_cluster` - (Optional) Advertise the workload locally in the cluster where it is deployed. See [Advertise Choice Advertise In Cluster ](#advertise-choice-advertise-in-cluster) below for details.

`advertise_on_public` - (Optional) Advertise the workload via load balancer on Internet with default VIP. See [Advertise Choice Advertise On Public ](#advertise-choice-advertise-on-public) below for details.

`do_not_advertise` - (Optional) Do not advertise the workload (`Bool`).

### Service Configuration

Configuration parameters for the service.

`parameters` - (Optional) Parameters for the workload. See [Configuration Parameters ](#configuration-parameters) below for details.

### Service Containers

Containers to use for service.

`args` - (Optional) Arguments to the entrypoint. Overrides the docker image's CMD (`String`).

`command` - (Optional) Command to execute. Overrides the docker image's ENTRYPOINT (`String`).

###### One of the arguments from this list "custom_flavor, default_flavor, flavor" can be set

`custom_flavor` - (Optional) Container resources based on Workload Flavor. See [ref](#ref) below for details.

`default_flavor` - (Optional) Container resources based on default workload flavor of the Virtual Kubernetes (`Bool`).

`flavor` - (Optional) Container resources based on system defined flavor (`String`).

`image` - (Required) Container image information. See [Containers Image ](#containers-image) below for details.

`init_container` - (Optional) Specialized container that runs before application container and runs to completion (`Bool`).

`liveness_check` - (Optional) Container will be restarted if the health check fails.. See [Containers Liveness Check ](#containers-liveness-check) below for details.

`name` - (Required) Name of the container (`String`).

`readiness_check` - (Optional) Container will be removed from service endpoints if the health check fails.. See [Containers Readiness Check ](#containers-readiness-check) below for details.

### Service Deploy Options

Deploy options for the service are used to decide where and how to deploy the service.

###### One of the arguments from this list "all_res, default_virtual_sites, deploy_ce_sites, deploy_ce_virtual_sites, deploy_re_sites, deploy_re_virtual_sites" must be set

`all_res` - (Optional) Deploy the workload on all Regional Edges (`Bool`).

`default_virtual_sites` - (Optional) Deploy the workload on virtual sites referred by the Virtual Kubernetes (`Bool`).

`deploy_ce_sites` - (Optional) Deploy the workload on specific Customer sites. See [Deploy Choice Deploy Ce Sites ](#deploy-choice-deploy-ce-sites) below for details.

`deploy_ce_virtual_sites` - (Optional) Deploy the workload on specific Customer virtual sites. See [Deploy Choice Deploy Ce Virtual Sites ](#deploy-choice-deploy-ce-virtual-sites) below for details.

`deploy_re_sites` - (Optional) Deploy the workload on specific Regional Edge sites. See [Deploy Choice Deploy Re Sites ](#deploy-choice-deploy-re-sites) below for details.

`deploy_re_virtual_sites` - (Optional) Deploy the workload on specific Regional Edge virtual sites. See [Deploy Choice Deploy Re Virtual Sites ](#deploy-choice-deploy-re-virtual-sites) below for details.

### Service Family

IP family for the service is used to derive the Service IP.

###### One of the arguments from this list "dual_v4_preferred, dual_v6_preferred, v4, v6" must be set

`dual_v4_preferred` - (Optional) Dual-stack IPv4 Preferred (`Bool`).

`dual_v6_preferred` - (Optional) Dual-stack IPv6 Preferred (`Bool`).

`v4` - (Optional) IPV4 Family (`Bool`).

`v6` - (Optional) IPV6 Family (`Bool`).

### Service Volumes

Volumes for the service.

`name` - (Optional) Name of the volume (`String`).

###### One of the arguments from this list "empty_dir, host_path, persistent_volume" must be set

`empty_dir` - (Optional) Volume containing a temporary directory whose lifetime is the same as a replica of a workload. See [Volume Choice Empty Dir ](#volume-choice-empty-dir) below for details.

`host_path` - (Optional) Volume containing a host mapped path into the workload. See [Volume Choice Host Path ](#volume-choice-host-path) below for details.

`persistent_volume` - (Optional) Volume containing the Persistent Storage for the workload. See [Volume Choice Persistent Volume ](#volume-choice-persistent-volume) below for details.

### Simple Route Path

URI path of route.

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Simple Service Configuration

Configuration parameters for the service.

`parameters` - (Optional) Parameters for the workload. See [Configuration Parameters ](#configuration-parameters) below for details.

### Simple Service Container

Container information.

`args` - (Optional) Arguments to the entrypoint. Overrides the docker image's CMD (`String`).

`command` - (Optional) Command to execute. Overrides the docker image's ENTRYPOINT (`String`).

###### One of the arguments from this list "custom_flavor, default_flavor, flavor" can be set

`custom_flavor` - (Optional) Container resources based on Workload Flavor. See [ref](#ref) below for details.

`default_flavor` - (Optional) Container resources based on default workload flavor of the Virtual Kubernetes (`Bool`).

`flavor` - (Optional) Container resources based on system defined flavor (`String`).

`image` - (Required) Container image information. See [Container Image ](#container-image) below for details.

`init_container` - (Optional) Specialized container that runs before application container and runs to completion (`Bool`).

`liveness_check` - (Optional) Container will be restarted if the health check fails.. See [Container Liveness Check ](#container-liveness-check) below for details.

`name` - (Required) Name of the container (`String`).

`readiness_check` - (Optional) Container will be removed from service endpoints if the health check fails.. See [Container Readiness Check ](#container-readiness-check) below for details.

### Specific Routes Routes

Routes for this loadbalancer.

###### One of the arguments from this list "custom_route_object, direct_response_route, redirect_route, simple_route" must be set

`custom_route_object` - (Optional) A custom route uses a route object created outside of this view.. See [Choice Custom Route Object ](#choice-custom-route-object) below for details.

`direct_response_route` - (Optional) A direct response route matches on path and/or HTTP method and responds directly to the matching traffic. See [Choice Direct Response Route ](#choice-direct-response-route) below for details.

`redirect_route` - (Optional) A redirect route matches on path and/or HTTP method and redirects the matching traffic to a different URL. See [Choice Redirect Route ](#choice-redirect-route) below for details.

`simple_route` - (Optional) A simple route matches on path and/or HTTP method and forwards the matching traffic to the workload. See [Choice Simple Route ](#choice-simple-route) below for details.

### Stateful Service Advertise Options

Advertise options for the service control how the service is advertised.

###### One of the arguments from this list "advertise_custom, advertise_in_cluster, advertise_on_public, do_not_advertise" must be set

`advertise_custom` - (Optional) Advertise the workload via load balancer on specific sites. See [Advertise Choice Advertise Custom ](#advertise-choice-advertise-custom) below for details.

`advertise_in_cluster` - (Optional) Advertise the workload locally in the cluster where it is deployed. See [Advertise Choice Advertise In Cluster ](#advertise-choice-advertise-in-cluster) below for details.

`advertise_on_public` - (Optional) Advertise the workload via load balancer on Internet with default VIP. See [Advertise Choice Advertise On Public ](#advertise-choice-advertise-on-public) below for details.

`do_not_advertise` - (Optional) Do not advertise the workload (`Bool`).

### Stateful Service Configuration

Configuration parameters for the service.

`parameters` - (Optional) Parameters for the workload. See [Configuration Parameters ](#configuration-parameters) below for details.

### Stateful Service Containers

Containers to use for service.

`args` - (Optional) Arguments to the entrypoint. Overrides the docker image's CMD (`String`).

`command` - (Optional) Command to execute. Overrides the docker image's ENTRYPOINT (`String`).

###### One of the arguments from this list "custom_flavor, default_flavor, flavor" can be set

`custom_flavor` - (Optional) Container resources based on Workload Flavor. See [ref](#ref) below for details.

`default_flavor` - (Optional) Container resources based on default workload flavor of the Virtual Kubernetes (`Bool`).

`flavor` - (Optional) Container resources based on system defined flavor (`String`).

`image` - (Required) Container image information. See [Containers Image ](#containers-image) below for details.

`init_container` - (Optional) Specialized container that runs before application container and runs to completion (`Bool`).

`liveness_check` - (Optional) Container will be restarted if the health check fails.. See [Containers Liveness Check ](#containers-liveness-check) below for details.

`name` - (Required) Name of the container (`String`).

`readiness_check` - (Optional) Container will be removed from service endpoints if the health check fails.. See [Containers Readiness Check ](#containers-readiness-check) below for details.

### Stateful Service Deploy Options

Deploy options for the service are used to decide where and how to deploy the service.

###### One of the arguments from this list "all_res, default_virtual_sites, deploy_ce_sites, deploy_ce_virtual_sites, deploy_re_sites, deploy_re_virtual_sites" must be set

`all_res` - (Optional) Deploy the workload on all Regional Edges (`Bool`).

`default_virtual_sites` - (Optional) Deploy the workload on virtual sites referred by the Virtual Kubernetes (`Bool`).

`deploy_ce_sites` - (Optional) Deploy the workload on specific Customer sites. See [Deploy Choice Deploy Ce Sites ](#deploy-choice-deploy-ce-sites) below for details.

`deploy_ce_virtual_sites` - (Optional) Deploy the workload on specific Customer virtual sites. See [Deploy Choice Deploy Ce Virtual Sites ](#deploy-choice-deploy-ce-virtual-sites) below for details.

`deploy_re_sites` - (Optional) Deploy the workload on specific Regional Edge sites. See [Deploy Choice Deploy Re Sites ](#deploy-choice-deploy-re-sites) below for details.

`deploy_re_virtual_sites` - (Optional) Deploy the workload on specific Regional Edge virtual sites. See [Deploy Choice Deploy Re Virtual Sites ](#deploy-choice-deploy-re-virtual-sites) below for details.

### Stateful Service Family

IP family for the service is used to derive the Service IP.

###### One of the arguments from this list "dual_v4_preferred, dual_v6_preferred, v4, v6" must be set

`dual_v4_preferred` - (Optional) Dual-stack IPv4 Preferred (`Bool`).

`dual_v6_preferred` - (Optional) Dual-stack IPv6 Preferred (`Bool`).

`v4` - (Optional) IPV4 Family (`Bool`).

`v6` - (Optional) IPV6 Family (`Bool`).

### Stateful Service Persistent Volumes

Persistent storage configuration for the service.

`name` - (Required) Name of the volume (`String`).

`persistent_volume` - (Required) Volume containing the Persistent Storage for the workload. See [Persistent Volumes Persistent Volume ](#persistent-volumes-persistent-volume) below for details.

### Stateful Service Volumes

Ephemeral volumes for the service.

`name` - (Optional) Name of the volume (`String`).

###### One of the arguments from this list "empty_dir, host_path" must be set

`empty_dir` - (Optional) Volume containing a temporary directory whose lifetime is the same as a replica of a workload. See [Volume Choice Empty Dir ](#volume-choice-empty-dir) below for details.

`host_path` - (Optional) Volume containing a host mapped path into the workload. See [Volume Choice Host Path ](#volume-choice-host-path) below for details.

### Target Port Choice Same As Port

Same as Port.

### Tcp Health Check Port

Number or name of the port to access on the container..

###### One of the arguments from this list "name, num" must be set

`name` - (Optional) Port Name (`String`).

`num` - (Optional) Port number (`Int`).

### Tls Cert Params Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Tls Certificates Choice Tls Cert Params

Select/Add one or more TLS Certificate objects to associate with this Load Balancer.

`certificates` - (Required) Select one or more certificates with any domain names.. See [ref](#ref) below for details.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Cert Params Tls Config ](#tls-cert-params-tls-config) below for details.

### Tls Certificates Choice Tls Parameters

Upload a TLS certificate covering all domain names for this Load Balancer.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Parameters Tls Certificates ](#tls-parameters-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Parameters Tls Config ](#tls-parameters-tls-config) below for details.

### Tls Parameters Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Tls Parameters Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Volume Choice Empty Dir

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

`mount` - (Required) Volume mount associated with the empty directory. See [Empty Dir Mount ](#empty-dir-mount) below for details.

`size_limit` - (Required) x-displayName: "Size Limit (in GiB)" (`Float`).

### Volume Choice Host Path

Volume containing a host mapped path into the workload.

`mount` - (Required) Volume mount associated with the host path. See [Host Path Mount ](#host-path-mount) below for details.

`path` - (Required) Path of the directory on the host (`String`).

### Volume Choice Persistent Volume

Volume containing the Persistent Storage for the workload.

`mount` - (Required) Volume mount associated with the persistent storage. See [Persistent Volume Mount ](#persistent-volume-mount) below for details.

`storage` - (Required) Persistent storage configuration. See [Persistent Volume Storage ](#persistent-volume-storage) below for details.

### Workload Choice Job

like report generation, billing, parallel data processing, ETL processing, etc..

`configuration` - (Optional) Configuration parameters for the job. See [Job Configuration ](#job-configuration) below for details.

`containers` - (Required) Containers to use for the job. See [Job Containers ](#job-containers) below for details.

`deploy_options` - (Required) Deploy options for the job are used to decide where and how to deploy the job. See [Job Deploy Options ](#job-deploy-options) below for details.

`num_replicas` - (Optional) Number of replicas of the batch job to spawn per site (`Int`).

`volumes` - (Optional) Volumes for the job. See [Job Volumes ](#job-volumes) below for details.

### Workload Choice Service

application servers, traditional SQL databases, etc..

`advertise_options` - (Required) Advertise options for the service control how the service is advertised. See [Service Advertise Options ](#service-advertise-options) below for details.

`configuration` - (Optional) Configuration parameters for the service. See [Service Configuration ](#service-configuration) below for details.

`containers` - (Required) Containers to use for service. See [Service Containers ](#service-containers) below for details.

`deploy_options` - (Required) Deploy options for the service are used to decide where and how to deploy the service. See [Service Deploy Options ](#service-deploy-options) below for details.

`family` - (Optional) IP family for the service is used to derive the Service IP. See [Service Family ](#service-family) below for details.

###### One of the arguments from this list "num_replicas, scale_to_zero" must be set

`num_replicas` - (Optional) Number of replicas of service to spawn per site (`Int`).

`scale_to_zero` - (Optional) Set number of replicas of service to spawn per site to zero (`Bool`).

`volumes` - (Optional) Volumes for the service. See [Service Volumes ](#service-volumes) below for details.

### Workload Choice Simple Service

on default VIP.

###### One of the arguments from this list "do_not_advertise, service_port, simple_advertise" must be set

`do_not_advertise` - (Optional) Do not advertise the service (`Bool`).

`service_port` - (Optional) Service port to advertise on Internet via HTTP loadbalancer using port 80 (`Int`).(Deprecated)

`simple_advertise` - (Optional) Domains and service port to advertise on Internet via HTTP loadbalancer using port 80. See [Advertise Choice Simple Advertise ](#advertise-choice-simple-advertise) below for details.

`configuration` - (Optional) Configuration parameters for the service. See [Simple Service Configuration ](#simple-service-configuration) below for details.

`container` - (Required) Container information. See [Simple Service Container ](#simple-service-container) below for details.

###### One of the arguments from this list "disabled, enabled" must be set

`disabled` - (Optional) Disabled (`Bool`).

`enabled` - (Optional) x-displayName: "Enabled". See [Persistence Choice Enabled ](#persistence-choice-enabled) below for details.

`scale_to_zero` - (Optional) Scale down replicas of the service to zero (`Bool`).

### Workload Choice Stateful Service

applications like cassandra, mongodb, redis, etc..

`advertise_options` - (Required) Advertise options for the service control how the service is advertised. See [Stateful Service Advertise Options ](#stateful-service-advertise-options) below for details.

`configuration` - (Optional) Configuration parameters for the service. See [Stateful Service Configuration ](#stateful-service-configuration) below for details.

`containers` - (Required) Containers to use for service. See [Stateful Service Containers ](#stateful-service-containers) below for details.

`deploy_options` - (Required) Deploy options for the service are used to decide where and how to deploy the service. See [Stateful Service Deploy Options ](#stateful-service-deploy-options) below for details.

`family` - (Optional) IP family for the service is used to derive the Service IP. See [Stateful Service Family ](#stateful-service-family) below for details.

`persistent_volumes` - (Required) Persistent storage configuration for the service. See [Stateful Service Persistent Volumes ](#stateful-service-persistent-volumes) below for details.

###### One of the arguments from this list "num_replicas, scale_to_zero" must be set

`num_replicas` - (Optional) Number of replicas of service to spawn per site (`Int`).

`scale_to_zero` - (Optional) Set number of replicas of service to spawn per site to zero (`Bool`).

`volumes` - (Optional) Ephemeral volumes for the service. See [Stateful Service Volumes ](#stateful-service-volumes) below for details.

### Xfcc Header Xfcc Disabled

No X-Forwarded-Client-Cert header will be added.

### Xfcc Header Xfcc Options

X-Forwarded-Client-Cert header will be added with the configured fields.

`xfcc_header_elements` - (Required) X-Forwarded-Client-Cert header elements to be added to requests (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured workload.
