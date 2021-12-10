---

page_title: "Volterra: azure_vnet_site"

description: "The azure_vnet_site allows CRUD of Azure Vnet Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_azure_vnet_site
=================================

The Azure Vnet Site allows CRUD of Azure Vnet Site resource on Volterra SaaS

~> **Note:** Please refer to [Azure Vnet Site API docs](https://volterra.io/docs/api/views-azure-vnet-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_azure_vnet_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "azure_cred assisted" must be set

  azure_cred {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }
  // One of the arguments from this list "logs_streaming_disabled log_receiver" must be set
  logs_streaming_disabled = true
  // One of the arguments from this list "azure_region alternate_region" must be set
  azure_region = "eastus"
  resource_group = ["my-resources"]

  // One of the arguments from this list "ingress_egress_gw voltstack_cluster ingress_gw_ar ingress_egress_gw_ar voltstack_cluster_ar ingress_gw" must be set

  ingress_gw_ar {
    azure_certified_hw = "azure-byol-voltmesh"

    node {
      fault_domain = "1"

      local_subnet {
        // One of the arguments from this list "subnet_param subnet" must be set

        subnet_param {
          ipv4 = "10.1.2.0/24"
          ipv6 = "1234:568:abcd:9100::/64"
        }
      }

      node_number   = "node_number"
      update_domain = "1"
    }
  }
  vnet {
    // One of the arguments from this list "new_vnet existing_vnet" must be set

    new_vnet {
      // One of the arguments from this list "name autogenerate" must be set
      name = "name"

      primary_ipv4 = "10.1.0.0/16"
    }
  }
  // One of the arguments from this list "nodes_per_az total_nodes no_worker_nodes" must be set
  no_worker_nodes = true
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

`assisted` - (Optional) In assisted deployment get Azure parameters generated in status of this objects and run volterra provided terraform script. (bool).

`azure_cred` - (Optional) Reference to Azure credentials for automatic deployment. See [ref](#ref) below for details.

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (bool).

`machine_type` - (Optional) Select Instance size based on performance needed (`String`).

`os` - (Optional) Operating System Details. See [Os ](#os) below for details.

`alternate_region` - (Optional) Name of the azure region which does not support availability zones. (`String`).

`azure_region` - (Optional) Name of the azure region which supports availability zones. (`String`).

`resource_group` - (Required) Azure resource group for resources that will be created (`String`).

`site_to_site_tunnel_ip` - (Optional) Optional, VIP in the site_to_site_network_type configured above used for terminating IPSec/SSL tunnels created with SiteMeshGroup. (`String`).

`ingress_egress_gw` - (Optional) Two interface site is useful when site is used as ingress/egress gateway to the Vnet.. See [Ingress Egress Gw ](#ingress-egress-gw) below for details.

`ingress_egress_gw_ar` - (Optional) Two interface site is useful when site is used as ingress/egress gateway to the Vnet.. See [Ingress Egress Gw Ar ](#ingress-egress-gw-ar) below for details.

`ingress_gw` - (Optional) One interface site is useful when site is only used as ingress gateway to the Vnet.. See [Ingress Gw ](#ingress-gw) below for details.

`ingress_gw_ar` - (Optional) One interface site is useful when site is only used as ingress gateway to the Vnet.. See [Ingress Gw Ar ](#ingress-gw-ar) below for details.

`voltstack_cluster` - (Optional) Voltstack Cluster using single interface, useful for deploying K8s cluster.. See [Voltstack Cluster ](#voltstack-cluster) below for details.

`voltstack_cluster_ar` - (Optional) Voltstack Cluster using single interface, useful for deploying K8s cluster.. See [Voltstack Cluster Ar ](#voltstack-cluster-ar) below for details.

`ssh_key` - (Optional) Public SSH key for accessing the site. (`String`).

`sw` - (Optional) Volterra Software Details. See [Sw ](#sw) below for details.

`vnet` - (Required) Choice of using existing Vnet or create new Vnet. See [Vnet ](#vnet) below for details.

`no_worker_nodes` - (Optional) Worker nodes is set to zero (bool).

`nodes_per_az` - (Optional) Desired Worker Nodes Per AZ. Max limit is up to 21 (`Int`).

`total_nodes` - (Optional) Total number of worker nodes to be deployed across all AZ's used in the Site (`Int`).

### Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) List of Forward Proxy Policies. See [ref](#ref) below for details.

### Active Network Policies

Network Policies active for this site..

`network_policies` - (Required) Ordered List of Network Policies active for this network firewall. See [ref](#ref) below for details.

### Autogenerate

Autogenerate the Vnet Name.

### Az Nodes

Only Single AZ or Three AZ(s) nodes are supported currently..

`azure_az` - (Required) Azure availability zone. (`String`).

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`inside_subnet` - (Optional) Subnets for the inside interface of the node. See [Inside Subnet ](#inside-subnet) below for details.

`outside_subnet` - (Optional) Subnets for the outside interface of the node. See [Outside Subnet ](#outside-subnet) below for details.

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

Clear Secret is used for the secrets that are not encrypted.

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

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. Volterra will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) Volterra will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Custom Hash Algorithms

Use hash algorithms in the custom order. Volterra will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Custom Static Route

Use Custom static route to configure all advanced options.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Nexthop ](#nexthop) below for details.

`subnets` - (Optional) List of route prefixes. See [Subnets ](#subnets) below for details.

### Default Os Version

Will assign latest available OS version.

### Default Storage

Use standard storage class configured as AWS EBS.

### Default Sw Version

Will assign latest available SW version.

### Disable Forward Proxy

Forward Proxy is disabled for this connector.

### Disable Interception

Disable Interception.

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

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

### Existing Vnet

Information about existing Vnet.

`resource_group` - (Required) Resource group of existing Vnet (`String`).

`vnet_name` - (Required) Name of existing Vnet (`String`).

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

Two interface site is useful when site is used as ingress/egress gateway to the Vnet..

`az_nodes` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

### Ingress Egress Gw Ar

Two interface site is useful when site is used as ingress/egress gateway to the Vnet..

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`node` - (Optional) Ingress/Egress Gateway (Two Interface) Node information.. See [Node ](#node) below for details.

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

### Ingress Gw

One interface site is useful when site is only used as ingress gateway to the Vnet..

`az_nodes` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

### Ingress Gw Ar

One interface site is useful when site is only used as ingress gateway to the Vnet..

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`node` - (Optional) Ingress Gateway (One Interface) Node information. See [Node ](#node) below for details.

### Inside Static Routes

Manage static routes for inside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Inside Subnet

Subnets for the inside interface of the node.

`subnet` - (Optional) Information about existing subnet.. See [Subnet ](#subnet) below for details.

`subnet_param` - (Optional) Parameters for creating new subnet.. See [Subnet Param ](#subnet-param) below for details.

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

### New Vnet

Parameters for creating new Vnet.

`autogenerate` - (Optional) Autogenerate the Vnet Name (bool).

`name` - (Optional) Specify the Vnet Name (`String`).

`primary_ipv4` - (Required) IPv4 CIDR block for this Vnet. It has to be private address space. (`String`).

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

### Node

Ingress/Egress Gateway (Two Interface) Node information..

`fault_domain` - (Optional) Namuber of fault domains to be used while creating the availability set (`Int`).

`inside_subnet` - (Optional) Subnets for the inside interface of the node. See [Inside Subnet ](#inside-subnet) below for details.

`node_number` - (Required) Number of main nodes to create, either 1 or 3. (`Int`).

`outside_subnet` - (Optional) Subnets for the outside interface of the node. See [Outside Subnet ](#outside-subnet) below for details.

`update_domain` - (Optional) Namuber of update domains to be used while creating the availability set (`Int`).

### Openebs Enterprise

Storage class Device configuration for OpenEBS Enterprise.

`replication` - (Optional) Replication sets the replication factor of the PV, i.e. the number of data replicas to be maintained for it such as 1 or 3. (`Int`).

`storage_class_size` - (Optional) Three 10GB disk will be created and assigned to nodes. (`Int`).

### Os

Operating System Details.

`default_os_version` - (Optional) Will assign latest available OS version (bool).

`operating_system_version` - (Optional) Operating System Version is optional parameter, which allows to specify target OS version for particular site e.g. 7.2009.10. (`String`).

### Outside Static Routes

Manage static routes for outside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Outside Subnet

Subnets for the outside interface of the node.

`subnet` - (Optional) Information about existing subnet.. See [Subnet ](#subnet) below for details.

`subnet_param` - (Optional) Parameters for creating new subnet.. See [Subnet Param ](#subnet-param) below for details.

### Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Interception Rules ](#interception-rules) below for details.

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by Volterra Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in Volterra Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

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

### Storage Class List

Add additional custom storage classes in kubernetes for site.

`storage_classes` - (Optional) List of custom storage classes. See [Storage Classes ](#storage-classes) below for details.

### Storage Classes

List of custom storage classes.

`default_storage_class` - (Optional) Make this storage class default storage class for the K8s cluster (`Bool`).

`openebs_enterprise` - (Optional) Storage class Device configuration for OpenEBS Enterprise. See [Openebs Enterprise ](#openebs-enterprise) below for details.

`storage_class_name` - (Required) Name of the storage class as it will appear in K8s. (`String`).

### Subnet

Information about existing subnet..

`subnet_resource_grp` - (Optional) Specify name of Resource Group (`String`).

`vnet_resource_group` - (Optional) Use the same Resource Group as the Vnet (bool).

`subnet_name` - (Required) Name of existing subnet. (`String`).

### Subnet Param

Parameters for creating new subnet..

`ipv4` - (Required) IPv4 subnet prefix for this subnet (`String`).

`ipv6` - (Optional) IPv6 subnet prefix for this subnet (`String`).

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

### Use System Defaults

Volterra will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Vnet

Choice of using existing Vnet or create new Vnet.

`existing_vnet` - (Optional) Information about existing Vnet. See [Existing Vnet ](#existing-vnet) below for details.

`new_vnet` - (Optional) Parameters for creating new Vnet. See [New Vnet ](#new-vnet) below for details.

### Vnet Resource Group

Use the same Resource Group as the Vnet.

### Volterra Certificate

Volterra certificates for generating intermediate certificate for TLS interception..

### Volterra Trusted Ca

Default volterra trusted CA list for validating upstream server certificate.

### Voltstack Cluster

Voltstack Cluster using single interface, useful for deploying K8s cluster..

`az_nodes` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`k8s_cluster` - (Optional) Site Local K8s API access is enabled, using k8s_cluster object. See [ref](#ref) below for details.

`no_k8s_cluster` - (Optional) Site Local K8s API access is disabled (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`default_storage` - (Optional) Use standard storage class configured as AWS EBS (bool).

`storage_class_list` - (Optional) Add additional custom storage classes in kubernetes for site. See [Storage Class List ](#storage-class-list) below for details.

### Voltstack Cluster Ar

Voltstack Cluster using single interface, useful for deploying K8s cluster..

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`k8s_cluster` - (Optional) Site Local K8s API access is enabled, using k8s_cluster object. See [ref](#ref) below for details.

`no_k8s_cluster` - (Optional) Site Local K8s API access is disabled (bool).

`active_network_policies` - (Optional) Network Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Network Policy is disabled for this site. (bool).

`node` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Node ](#node) below for details.

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`default_storage` - (Optional) Use standard storage class configured as AWS EBS (bool).

`storage_class_list` - (Optional) Add additional custom storage classes in kubernetes for site. See [Storage Class List ](#storage-class-list) below for details.

### Wingman Secret Info

Secret is given as bootstrap secret in Volterra Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured azure_vnet_site.
