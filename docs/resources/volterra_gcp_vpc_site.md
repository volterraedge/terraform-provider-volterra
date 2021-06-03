---

page_title: "Volterra: gcp_vpc_site"

description: "The gcp_vpc_site allows CRUD of Gcp Vpc Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_gcp_vpc_site
==============================

The Gcp Vpc Site allows CRUD of Gcp Vpc Site resource on Volterra SaaS

~> **Note:** Please refer to [Gcp Vpc Site API docs](https://volterra.io/docs/api/views-gcp-vpc-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_gcp_vpc_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "cloud_credentials assisted" must be set

  cloud_credentials {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }
  gcp_region    = ["us-west1"]
  instance_type = ["n1-standard-4"]
  // One of the arguments from this list "logs_streaming_disabled log_receiver" must be set
  logs_streaming_disabled = true

  // One of the arguments from this list "voltstack_cluster ingress_gw ingress_egress_gw" must be set

  ingress_gw {
    gcp_certified_hw = "gcp-byol-voltmesh"

    gcp_zone_names = ["us-west1-a, us-west1-b, us-west1-c"]

    local_network {
      // One of the arguments from this list "new_network_autogenerate new_network existing_network" must be set

      new_network_autogenerate {
        autogenerate = true
      }
    }

    local_subnet {
      // One of the arguments from this list "existing_subnet new_subnet" must be set

      new_subnet {
        primary_ipv4 = "10.1.0.0/16"
        subnet_name  = "subnet1-in-network1"
      }
    }

    node_number = "node_number"
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

`address` - (Optional) Site's geographical address that can be used determine its latitude and longitude. (`String`).

`coordinates` - (Optional) Site longitude and latitude co-ordinates. See [Coordinates ](#coordinates) below for details.

`assisted` - (Optional) In assisted deployment get GCP parameters generated in status of this objects and run volterra provided terraform script. (bool).

`cloud_credentials` - (Optional) Reference to GCP credentials for automatic deployment. See [ref](#ref) below for details.

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`gcp_region` - (Required) Name for GCP Region. (`String`).

`instance_type` - (Required) Select Instance size based on performance needed (`String`).

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (bool).

`nodes_per_az` - (Optional) Desired Worker Nodes Per AZ. Max limit is up to 21 (`Int`).

`os` - (Optional) Operating System Details. See [Os ](#os) below for details.

`ingress_egress_gw` - (Optional) Two interface site is useful when site is used as ingress/egress gateway to the VPC network.. See [Ingress Egress Gw ](#ingress-egress-gw) below for details.

`ingress_gw` - (Optional) One interface site is useful when site is only used as ingress gateway to the VPC network.. See [Ingress Gw ](#ingress-gw) below for details.

`voltstack_cluster` - (Optional) Voltstack Cluster using single interface, useful for deploying K8s cluster.. See [Voltstack Cluster ](#voltstack-cluster) below for details.

`ssh_key` - (Optional) Public SSH key for accessing the site. (`String`).

`sw` - (Optional) Volterra Software Details. See [Sw ](#sw) below for details.

### Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) List of Forward Proxy Policies. See [ref](#ref) below for details.

### Active Network Policies

Network Policies active for this site..

`network_policies` - (Required) Ordered List of Network Policies active for this network firewall. See [ref](#ref) below for details.

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by Volterra Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted .

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Coordinates

Site longitude and latitude co-ordinates.

`latitude` - (Optional) Latitude of the site location (`Float`).

`longitude` - (Optional) longitude of site location (`Float`).

### Custom Certificate

Certificates for generating intermediate certificate for TLS interception..

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Custom Static Route

Use Custom static route to configure all advanced options.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Nexthop ](#nexthop) below for details.

`subnets` - (Optional) List of route prefixes. See [Subnets ](#subnets) below for details.

### Default Os Version

Will assign latest available OS version.

### Default Sw Version

Will assign latest available SW version.

### Disable Forward Proxy

Forward Proxy is disabled for this connector.

### Disable Interception

Disable Interception.

### Domain Match

Domain value or regular expression to match.

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Enable For All Domains

Enable interception for all domains.

### Enable Forward Proxy

Forward Proxy is enabled for this connector.

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

`max_connect_attempts` - (Optional) Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1. (`Int`).

`no_interception` - (Optional) No TLS interception is enabled for this network connector (bool).

`tls_intercept` - (Optional) Specify TLS interception configuration for the network connector. See [Tls Intercept ](#tls-intercept) below for details.

`white_listed_ports` - (Optional) Example "tmate" server port (`Int`).

`white_listed_prefixes` - (Optional) Example "tmate" server ip (`String`).

### Enable Interception

Enable Interception.

### Existing Network

Name of existing VPC network..

`name` - (Required) Name for your GCP VPC Network (`String`).

### Existing Subnet

Name of existing VPC subnet..

`subnet_name` - (Required) Name of your subnet in VPC network (`String`).

### Forward Proxy Allow All

Enable Forward Proxy for this site and allow all requests..

### Global Network Connections

Global network connections.

`sli_to_global_dr` - (Optional) Site local inside is connected directly to a given global network. See [Sli To Global Dr ](#sli-to-global-dr) below for details.

`slo_to_global_dr` - (Optional) Site local outside is connected directly to a given global network. See [Slo To Global Dr ](#slo-to-global-dr) below for details.

`disable_forward_proxy` - (Optional) Forward Proxy is disabled for this connector (bool).

`enable_forward_proxy` - (Optional) Forward Proxy is enabled for this connector. See [Enable Forward Proxy ](#enable-forward-proxy) below for details.

### Global Network List

List of global network connections.

`global_network_connections` - (Required) Global network connections. See [Global Network Connections ](#global-network-connections) below for details.

### Ingress Egress Gw

Two interface site is useful when site is used as ingress/egress gateway to the VPC network..

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`gcp_certified_hw` - (Required) Name for GCP certified hardware. (`String`).

`gcp_zone_names` - (Required) List of zones when instances will be created, needs to match with region selected. (`String`).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`inside_network` - (Optional) Network for the inside interface of the node. See [Inside Network ](#inside-network) below for details.

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`inside_subnet` - (Optional) Subnet for the inside interface of the node.. See [Inside Subnet ](#inside-subnet) below for details.

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`node_number` - (Optional) Number of nodes to create, either 1 or 3. (`Int`).

`outside_network` - (Optional) Network for the outside interface of the node. See [Outside Network ](#outside-network) below for details.

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`outside_subnet` - (Optional) Subnet for the outside interface of the node.. See [Outside Subnet ](#outside-subnet) below for details.

### Ingress Gw

One interface site is useful when site is only used as ingress gateway to the VPC network..

`gcp_certified_hw` - (Required) Name for GCP certified hardware. (`String`).

`gcp_zone_names` - (Required) List of zones when instances will be created, needs to match with region selected. (`String`).

`local_network` - (Optional) Network for the local interface of the node. See [Local Network ](#local-network) below for details.

`local_subnet` - (Optional) Subnet for the local interface of the node.. See [Local Subnet ](#local-subnet) below for details.

`node_number` - (Optional) Number of nodes to create, either 1 or 3. (`Int`).

### Inside Network

Network for the inside interface of the node.

`existing_network` - (Optional) Name of existing VPC network.. See [Existing Network ](#existing-network) below for details.

`new_network` - (Optional) Create new VPC network with specified name.. See [New Network ](#new-network) below for details.

`new_network_autogenerate` - (Optional) Create new VPC network with autogenerated name.. See [New Network Autogenerate ](#new-network-autogenerate) below for details.

### Inside Static Routes

Manage static routes for inside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Inside Subnet

Subnet for the inside interface of the node..

`existing_subnet` - (Optional) Name of existing VPC subnet.. See [Existing Subnet ](#existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating a new VPC Subnet. See [New Subnet ](#new-subnet) below for details.

### Interception Rules

List of ordered rules to enable or disable for TLS interception.

`domain_match` - (Required) Domain value or regular expression to match. See [Domain Match ](#domain-match) below for details.

`disable_interception` - (Optional) Disable Interception (bool).

`enable_interception` - (Optional) Enable Interception (bool).

### Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Local Network

Network for the local interface of the node.

`existing_network` - (Optional) Name of existing VPC network.. See [Existing Network ](#existing-network) below for details.

`new_network` - (Optional) Create new VPC network with specified name.. See [New Network ](#new-network) below for details.

`new_network_autogenerate` - (Optional) Create new VPC network with autogenerated name.. See [New Network Autogenerate ](#new-network-autogenerate) below for details.

### Local Subnet

Subnet for the local interface of the node..

`existing_subnet` - (Optional) Name of existing VPC subnet.. See [Existing Subnet ](#existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating a new VPC Subnet. See [New Subnet ](#new-subnet) below for details.

### New Network

Create new VPC network with specified name..

`name` - (Required) Name for your GCP VPC Network (`String`).

### New Network Autogenerate

Create new VPC network with autogenerated name..

`autogenerate` - (Optional) Name for your GCP VPC Network will be autogenerated (`Bool`).

### New Subnet

Parameters for creating a new VPC Subnet.

`primary_ipv4` - (Required) IPv4 prefix for this Subnet. It has to be private address space. (`String`).

`subnet_name` - (Optional) Name of new VPC Subnet, will be autogenerated if empty (`String`).

### Nexthop

Nexthop for the route.

`interface` - (Optional) Nexthop is network interface when type is "Network-Interface". See [ref](#ref) below for details.

`nexthop_address` - (Optional) Nexthop address when type is "Use-Configured". See [Nexthop Address ](#nexthop-address) below for details.

`type` - (Optional) Identifies the type of next-hop (`String`).

### Nexthop Address

Nexthop address when type is "Use-Configured".

`ipv4` - (Optional) IPv4 Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ipv6 ](#ipv6) below for details.

### No Forward Proxy

Disable Forward Proxy for this site.

### No Global Network

No global network to connect.

### No Inside Static Routes

Static Routes disabled for inside network..

### No Interception

No TLS interception is enabled for this network connector.

### No K8s Cluster

Site Local K8s API access is disabled.

### No Network Policy

Network Policy is disabled for this site..

### No Outside Static Routes

Static Routes disabled for outside network..

### Os

Operating System Details.

`default_os_version` - (Optional) Will assign latest available OS version (bool).

`operating_system_version` - (Optional) Operating System Version is optional parameter, which allows to specify target OS version for particular site e.g. 7.2009.10. (`String`).

### Outside Network

Network for the outside interface of the node.

`existing_network` - (Optional) Name of existing VPC network.. See [Existing Network ](#existing-network) below for details.

`new_network` - (Optional) Create new VPC network with specified name.. See [New Network ](#new-network) below for details.

`new_network_autogenerate` - (Optional) Create new VPC network with autogenerated name.. See [New Network Autogenerate ](#new-network-autogenerate) below for details.

### Outside Static Routes

Manage static routes for outside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Outside Subnet

Subnet for the outside interface of the node..

`existing_subnet` - (Optional) Name of existing VPC subnet.. See [Existing Subnet ](#existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating a new VPC Subnet. See [New Subnet ](#new-subnet) below for details.

### Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Interception Rules ](#interception-rules) below for details.

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted . See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Site Local Network

Network for the local interface of the node.

`existing_network` - (Optional) Name of existing VPC network.. See [Existing Network ](#existing-network) below for details.

`new_network` - (Optional) Create new VPC network with specified name.. See [New Network ](#new-network) below for details.

`new_network_autogenerate` - (Optional) Create new VPC network with autogenerated name.. See [New Network Autogenerate ](#new-network-autogenerate) below for details.

### Site Local Subnet

Subnet for the local interface of the node..

`existing_subnet` - (Optional) Name of existing VPC subnet.. See [Existing Subnet ](#existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating a new VPC Subnet. See [New Subnet ](#new-subnet) below for details.

### Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Static Route List

List of Static routes.

`custom_static_route` - (Optional) Use Custom static route to configure all advanced options. See [Custom Static Route ](#custom-static-route) below for details.

`simple_static_route` - (Optional) Use simple static route for prefix pointing to single interface in the network (`String`).

### Subnets

List of route prefixes.

`ipv4` - (Optional) IPv4 Subnet Address. See [Ipv4 ](#ipv4) below for details.

`ipv6` - (Optional) IPv6 Subnet Address. See [Ipv6 ](#ipv6) below for details.

### Sw

Volterra Software Details.

`default_sw_version` - (Optional) Will assign latest available SW version (bool).

`volterra_software_version` - (Optional) Volterra Software Version is optional parameter, which allows to specify target SW version for particular site e.g. crt-20210329-1002. (`String`).

### Tls Intercept

Specify TLS interception configuration for the network connector.

`enable_for_all_domains` - (Optional) Enable interception for all domains (bool).

`policy` - (Optional) Policy to enable/disable specific domains, with implicit enable all domains. See [Policy ](#policy) below for details.

`custom_certificate` - (Optional) Certificates for generating intermediate certificate for TLS interception.. See [Custom Certificate ](#custom-certificate) below for details.

`volterra_certificate` - (Optional) Volterra certificates for generating intermediate certificate for TLS interception. (bool).

`trusted_ca_url` - (Optional) Custom trusted CA certificates for validating upstream server certificate (`String`).

`volterra_trusted_ca` - (Optional) Default volterra trusted CA list for validating upstream server certificate (bool).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Volterra Certificate

Volterra certificates for generating intermediate certificate for TLS interception..

### Volterra Trusted Ca

Default volterra trusted CA list for validating upstream server certificate.

### Voltstack Cluster

Voltstack Cluster using single interface, useful for deploying K8s cluster..

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`gcp_certified_hw` - (Required) Name for GCP certified hardware. (`String`).

`gcp_zone_names` - (Required) List of zones when instances will be created, needs to match with region selected. (`String`).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`k8s_cluster` - (Optional) Site Local K8s API access is enabled, using k8s_cluster object. See [ref](#ref) below for details.

`no_k8s_cluster` - (Optional) Site Local K8s API access is disabled (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`node_number` - (Optional) Number of nodes to create, either 1 or 3. (`Int`).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`site_local_network` - (Optional) Network for the local interface of the node. See [Site Local Network ](#site-local-network) below for details.

`site_local_subnet` - (Optional) Subnet for the local interface of the node.. See [Site Local Subnet ](#site-local-subnet) below for details.

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured gcp_vpc_site.
