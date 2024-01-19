---

page_title: "Volterra: azure_vnet_site"

description: "The azure_vnet_site allows CRUD of Azure Vnet Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------------

Resource volterra_azure_vnet_site
=================================

The Azure Vnet Site allows CRUD of Azure Vnet Site resource on Volterra SaaS

~> **Note:** Please refer to [Azure Vnet Site API docs](https://docs.cloud.f5.com/docs/api/views-azure-vnet-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_azure_vnet_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "default_blocked_services block_all_services blocked_services" must be set
  default_blocked_services = true

  // One of the arguments from this list "azure_cred" must be set

  azure_cred {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }
  // One of the arguments from this list "log_receiver logs_streaming_disabled" must be set
  logs_streaming_disabled = true
  // One of the arguments from this list "alternate_region azure_region" must be set
  azure_region = "eastus"
  resource_group = ["my-resources"]

  // One of the arguments from this list "ingress_gw ingress_egress_gw voltstack_cluster ingress_gw_ar ingress_egress_gw_ar voltstack_cluster_ar" must be set

  ingress_gw_ar {
    accelerated_networking {
      // One of the arguments from this list "disable enable" must be set
      enable = true
    }

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

      node_number   = "1"
      update_domain = "1"
    }

    performance_enhancement_mode {
      // One of the arguments from this list "perf_mode_l7_enhanced perf_mode_l3_enhanced" must be set
      perf_mode_l7_enhanced = true
    }
  }
  ssh_key = ["ssh-rsa AAAAB..."]
  vnet {
    // One of the arguments from this list "new_vnet existing_vnet" must be set

    existing_vnet {
      resource_group = "MyResourceGroup"
      vnet_name      = "MyVnet"
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

`block_all_services` - (Optional) Block DNS, SSH & WebUI services on Site (bool).

`blocked_services` - (Optional) Use custom blocked services configuration. See [Blocked Services ](#blocked-services) below for details.

`default_blocked_services` - (Optional) Allow access to DNS, SSH services on Site (bool).

`coordinates` - (Optional) Site longitude and latitude co-ordinates. See [Coordinates ](#coordinates) below for details.

`azure_cred` - (Optional) Reference to Azure credentials for automatic deployment. See [ref](#ref) below for details.

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (bool).

`machine_type` - (Optional) > advanced options. (`String`).

`offline_survivability_mode` - (Optional) Enable/Disable offline survivability mode. See [Offline Survivability Mode ](#offline-survivability-mode) below for details.

`os` - (Optional) Operating System Details. See [Os ](#os) below for details.

`alternate_region` - (Optional) Name of the azure region which does not support availability zones. (`String`).

`azure_region` - (Optional) Name of the azure region which supports availability zones. (`String`).

`resource_group` - (Required) Azure resource group for resources that will be created (`String`).

`ingress_egress_gw` - (Optional) Two interface site is useful when site is used as ingress/egress gateway to the VNet.. See [Ingress Egress Gw ](#ingress-egress-gw) below for details.

`ingress_egress_gw_ar` - (Optional) Two interface site is useful when site is used as ingress/egress gateway to the VNet.. See [Ingress Egress Gw Ar ](#ingress-egress-gw-ar) below for details.

`ingress_gw` - (Optional) One interface site is useful when site is only used as ingress gateway to the VNet.. See [Ingress Gw ](#ingress-gw) below for details.

`ingress_gw_ar` - (Optional) One interface site is useful when site is only used as ingress gateway to the VNet.. See [Ingress Gw Ar ](#ingress-gw-ar) below for details.

`voltstack_cluster` - (Optional) App Stack Cluster using single interface, useful for deploying K8s cluster.. See [Voltstack Cluster ](#voltstack-cluster) below for details.

`voltstack_cluster_ar` - (Optional) App Stack Cluster using single interface, useful for deploying K8s cluster.. See [Voltstack Cluster Ar ](#voltstack-cluster-ar) below for details.

`ssh_key` - (Required) Public SSH key for accessing the site. (`String`).

`sw` - (Optional) F5XC Software Details. See [Sw ](#sw) below for details.

`tags` - (Optional) It helps to manage, identify, organize, search for, and filter resources in Azure console. (`String`).

`vnet` - (Required) Choice of using existing VNet or create new VNet. See [Vnet ](#vnet) below for details.

`no_worker_nodes` - (Optional) Worker nodes is set to zero (bool).

`nodes_per_az` - (Optional) Desired Worker Nodes Per AZ. Max limit is up to 21 (`Int`).

`total_nodes` - (Optional) Total number of worker nodes to be deployed across all AZ's used in the Site (`Int`).

### Accelerated Networking

disruption will be seen.

`disable` - (Optional) infrastructure. (bool).

`enable` - (Optional) improving networking performance (bool).

### Active Enhanced Firewall Policies

with an additional option for service insertion..

`enhanced_firewall_policies` - (Required) Ordered List of Enhaned Firewall Policy active for this network firewall. See [ref](#ref) below for details.

### Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) List of Forward Proxy Policies. See [ref](#ref) below for details.

### Active Network Policies

Firewall Policies active for this site..

`network_policies` - (Required) Ordered List of Firewall Policies active for this network firewall. See [ref](#ref) below for details.

### Advertise To Route Server

Advertise Spoke Vnet CIDR Routes To Azure Route Server via BGP.

### Authorized Key

Authorization Key created by the circuit owner.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Auto

The subnet CIDR is autogenerated..

### Auto Asn

(Recommended) Automatically set ASN for F5XC Site.

### Autogenerate

Autogenerate the Vnet Name.

### Az Nodes

Only Single AZ or Three AZ(s) nodes are supported currently..

`azure_az` - (Required) Azure availability zone. (`String`).

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`inside_subnet` - (Optional) Subnets for the inside interface of the node. See [Inside Subnet ](#inside-subnet) below for details.

`outside_subnet` - (Optional) Subnets for the outside interface of the node. See [Outside Subnet ](#outside-subnet) below for details.

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

### Blocked Services

Use custom blocked services configuration.

`blocked_sevice` - (Optional) Use custom blocked services configuration. See [Blocked Sevice ](#blocked-sevice) below for details.

### Blocked Sevice

Use custom blocked services configuration.

`dns` - (Optional) Matches DNS port 53 (bool).

`ssh` - (Optional) Matches ssh port 22 (bool).

`web_user_interface` - (Optional) Matches the web user interface port (bool).

`network_type` - (Optional) Network type in which these ports get blocked (`String`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Connections

Add the ExpressRoute Circuit Connections to this site.

`metadata` - (Required) Connection Metadata like name and description. See [Metadata ](#metadata) below for details.

`circuit_id` - (Optional) ExpressRoute Circuit is in same subscription as the site (`String`).

`other_subscription` - (Optional) ExpressRoute Circuit is in a different subscription than the site. In this case both Circuit ID and Authorization key are needed. See [Other Subscription ](#other-subscription) below for details.

`weight` - (Optional) The weight (or priority) for the routes received from this connection. The default value is 10. (`Int`).

### Coordinates

Site longitude and latitude co-ordinates.

`latitude` - (Optional) Latitude of the site location (`Float`).

`longitude` - (Optional) longitude of site location (`Float`).

### Custom Certificate

Certificates for generating intermediate certificate for TLS interception..

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Custom Static Route

Use Custom static route to configure all advanced options.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Nexthop ](#nexthop) below for details.

`subnets` - (Required) List of route prefixes. See [Subnets ](#subnets) below for details.

### Default Os Version

Will assign latest available OS version.

### Default Storage

Use standard storage class configured as AWS EBS.

### Default Sw Version

Will assign latest available SW version.

### Disable

infrastructure..

### Disable Forward Proxy

Forward Proxy is disabled for this connector.

### Disable Interception

Disable Interception.

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Dns

Matches DNS port 53.

### Do Not Advertise To Route Server

Do Not Advertise Spoke Vnet CIDR Routes To Azure Route Server via BGP.

### Domain Match

Domain value or regular expression to match.

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Enable

improving networking performance.

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

### Enable Offline Survivability Mode

When this feature is enabled on an existing site, the pods/services on this site will be restarted..

### Existing Vnet

Information about existing Vnet.

`resource_group` - (Required) Resource group of existing Vnet (`String`).

`vnet_name` - (Required) Name of existing Vnet (`String`).

### Express Route Disabled

Express Route is disabled on this site.

### Express Route Enabled

Express Route is enabled on this site.

`auto_asn` - (Optional) (Recommended) Automatically set ASN for F5XC Site (bool).

`custom_asn` - (Optional) Set custom ASN for F5XC Site (`Int`).

`connections` - (Required) Add the ExpressRoute Circuit Connections to this site. See [Connections ](#connections) below for details.

`site_registration_over_express_route` - (Optional) Site Registration and Site to RE tunnels go over the Azure Express Route. See [Site Registration Over Express Route ](#site-registration-over-express-route) below for details.

`site_registration_over_internet` - (Optional) Site Registration and Site to RE tunnels go over the internet (bool).

`gateway_subnet` - (Optional) Select the type of subnet to be used for VNet Gateway. See [Gateway Subnet ](#gateway-subnet) below for details.

`route_server_subnet` - (Optional) Select the type of subnet to be used for Azure Route Server. See [Route Server Subnet ](#route-server-subnet) below for details.

`sku_ergw1az` - (Optional) ErGw1Az SKU (Standard + Zone protection) (bool).

`sku_ergw2az` - (Optional) ErGw2Az SKU (High Perf + Zone protection) (bool).

`sku_high_perf` - (Optional) High Perf SKU (bool).

`sku_standard` - (Optional) Standard SKU (bool).

`advertise_to_route_server` - (Optional) Advertise Spoke Vnet CIDR Routes To Azure Route Server via BGP (bool).

`do_not_advertise_to_route_server` - (Optional) Do Not Advertise Spoke Vnet CIDR Routes To Azure Route Server via BGP (bool).

### Forward Proxy Allow All

Enable Forward Proxy for this site and allow all requests..

### Gateway Subnet

Select the type of subnet to be used for VNet Gateway.

`auto` - (Optional) The subnet CIDR is autogenerated. (bool).

`subnet` - (Optional) An existing subnet in specified resource group is used.. See [Subnet ](#subnet) below for details.

`subnet_param` - (Optional) A new subnet with specified CIDR is created.. See [Subnet Param ](#subnet-param) below for details.

### Global Network Connections

Global network connections.

`sli_to_global_dr` - (Optional) Site local inside is connected directly to a given global network. See [Sli To Global Dr ](#sli-to-global-dr) below for details.

`slo_to_global_dr` - (Optional) Site local outside is connected directly to a given global network. See [Slo To Global Dr ](#slo-to-global-dr) below for details.

`disable_forward_proxy` - (Optional) Forward Proxy is disabled for this connector (bool).

`enable_forward_proxy` - (Optional) Forward Proxy is enabled for this connector. See [Enable Forward Proxy ](#enable-forward-proxy) below for details.

### Global Network List

List of global network connections.

`global_network_connections` - (Required) Global network connections. See [Global Network Connections ](#global-network-connections) below for details.

### Hub

This VNet is a hub VNet.

`express_route_disabled` - (Optional) Express Route is disabled on this site (bool).

`express_route_enabled` - (Optional) Express Route is enabled on this site. See [Express Route Enabled ](#express-route-enabled) below for details.

`spoke_vnets` - (Optional) Spoke VNet Peering. See [Spoke Vnets ](#spoke-vnets) below for details.

### Ingress Egress Gw

Two interface site is useful when site is used as ingress/egress gateway to the VNet..

`accelerated_networking` - (Optional) disruption will be seen. See [Accelerated Networking ](#accelerated-networking) below for details.

`az_nodes` - (Required) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`dc_cluster_group_inside_vn` - (Optional) This site is member of dc cluster group connected via inside network. See [ref](#ref) below for details.

`dc_cluster_group_outside_vn` - (Optional) This site is member of dc cluster group connected via outside network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (bool).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`hub` - (Optional) This VNet is a hub VNet. See [Hub ](#hub) below for details.

`not_hub` - (Optional) This VNet is a standalone VNet (bool).

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Active Enhanced Firewall Policies ](#active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Firewall Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Firewall Policy is disabled for this site. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

`sm_connection_public_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

### Ingress Egress Gw Ar

Two interface site is useful when site is used as ingress/egress gateway to the VNet..

`accelerated_networking` - (Optional) disruption will be seen. See [Accelerated Networking ](#accelerated-networking) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`dc_cluster_group_inside_vn` - (Optional) This site is member of dc cluster group connected via inside network. See [ref](#ref) below for details.

`dc_cluster_group_outside_vn` - (Optional) This site is member of dc cluster group connected via outside network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (bool).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`hub` - (Optional) This VNet is a hub VNet. See [Hub ](#hub) below for details.

`not_hub` - (Optional) This VNet is a standalone VNet (bool).

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Active Enhanced Firewall Policies ](#active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Firewall Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Firewall Policy is disabled for this site. (bool).

`node` - (Optional) Ingress/Egress Gateway (Two Interface) Node information.. See [Node ](#node) below for details.

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

`sm_connection_public_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

### Ingress Gw

One interface site is useful when site is only used as ingress gateway to the VNet..

`accelerated_networking` - (Optional) disruption will be seen. See [Accelerated Networking ](#accelerated-networking) below for details.

`az_nodes` - (Required) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

### Ingress Gw Ar

One interface site is useful when site is only used as ingress gateway to the VNet..

`accelerated_networking` - (Optional) disruption will be seen. See [Accelerated Networking ](#accelerated-networking) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`node` - (Optional) Ingress Gateway (One Interface) Node information. See [Node ](#node) below for details.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

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

### Jumbo

L3 performance mode enhancement to use jumbo frame.

### Manual

Manually setup routing on spoke VNet.

### Metadata

Connection Metadata like name and description.

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

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

### No Dc Cluster Group

This site is not a member of dc cluster group.

### No Forward Proxy

Disable Forward Proxy for this site.

### No Global Network

No global network to connect.

### No Inside Static Routes

Static Routes disabled for inside network..

### No Interception

No TLS interception is enabled for this network connector.

### No Jumbo

L3 performance mode enhancement without jumbo frame.

### No K8s Cluster

Site Local K8s API access is disabled.

### No Network Policy

Firewall Policy is disabled for this site..

### No Offline Survivability Mode

When this feature is disabled on an existing site, the pods/services on this site will be restarted..

### No Outside Static Routes

Static Routes disabled for outside network..

### Node

Ingress/Egress Gateway (Two Interface) Node information..

`fault_domain` - (Optional) Namuber of fault domains to be used while creating the availability set (`Int`).

`inside_subnet` - (Optional) Subnets for the inside interface of the node. See [Inside Subnet ](#inside-subnet) below for details.

`node_number` - (Required) Number of main nodes to create, either 1 or 3. (`Int`).

`outside_subnet` - (Optional) Subnets for the outside interface of the node. See [Outside Subnet ](#outside-subnet) below for details.

`update_domain` - (Optional) Namuber of update domains to be used while creating the availability set (`Int`).

### Not Hub

This VNet is a standalone VNet.

### Offline Survivability Mode

Enable/Disable offline survivability mode.

`enable_offline_survivability_mode` - (Optional) When this feature is enabled on an existing site, the pods/services on this site will be restarted. (bool).

`no_offline_survivability_mode` - (Optional) When this feature is disabled on an existing site, the pods/services on this site will be restarted. (bool).

### Os

Operating System Details.

`default_os_version` - (Optional) Will assign latest available OS version (bool).

`operating_system_version` - (Optional) Operating System Version is optional parameter, which allows to specify target OS version for particular site e.g. 7.2009.10. (`String`).

### Other Subscription

ExpressRoute Circuit is in a different subscription than the site. In this case both Circuit ID and Authorization key are needed.

`authorized_key` - (Optional) Authorization Key created by the circuit owner. See [Authorized Key ](#authorized-key) below for details.

`circuit_id` - (Optional) Circuit ID (`String`).

### Outside Static Routes

Manage static routes for outside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Outside Subnet

Subnets for the outside interface of the node.

`subnet` - (Optional) Information about existing subnet.. See [Subnet ](#subnet) below for details.

`subnet_param` - (Optional) Parameters for creating new subnet.. See [Subnet Param ](#subnet-param) below for details.

### Perf Mode L3 Enhanced

When the mode is toggled to l3 enhanced, traffic disruption will be seen.

`jumbo` - (Optional) L3 performance mode enhancement to use jumbo frame (bool).

`no_jumbo` - (Optional) L3 performance mode enhancement without jumbo frame (bool).

### Perf Mode L7 Enhanced

When the mode is toggled to l7 enhanced, traffic disruption will be seen.

### Performance Enhancement Mode

Performance Enhancement Mode to optimize for L3 or L7 networking.

`perf_mode_l3_enhanced` - (Optional) When the mode is toggled to l3 enhanced, traffic disruption will be seen. See [Perf Mode L3 Enhanced ](#perf-mode-l3-enhanced) below for details.

`perf_mode_l7_enhanced` - (Optional) When the mode is toggled to l7 enhanced, traffic disruption will be seen (bool).

### Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Interception Rules ](#interception-rules) below for details.

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Route Server Subnet

Select the type of subnet to be used for Azure Route Server.

`auto` - (Optional) The subnet CIDR is autogenerated. (bool).

`subnet` - (Optional) An existing subnet in specified resource group is used.. See [Subnet ](#subnet) below for details.

`subnet_param` - (Optional) A new subnet with specified CIDR is created.. See [Subnet Param ](#subnet-param) below for details.

### Site Registration Over Express Route

Site Registration and Site to RE tunnels go over the Azure Express Route.

`cloudlink_network_name` - (Required) Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support. (`String`).

### Site Registration Over Internet

Site Registration and Site to RE tunnels go over the internet.

### Sku Ergw1az

ErGw1Az SKU (Standard + Zone protection).

### Sku Ergw2az

ErGw2Az SKU (High Perf + Zone protection).

### Sku High Perf

High Perf SKU.

### Sku Standard

Standard SKU.

### Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Sm Connection Public Ip

creating ipsec between two sites which are part of the site mesh group.

### Sm Connection Pvt Ip

creating ipsec between two sites which are part of the site mesh group.

### Spoke Vnets

Spoke VNet Peering.

`labels` - (Optional) These labels used must be from known key and label defined in shared namespace (`String`).

`auto` - (Optional) setup routing for all existing subnets on spoke VNet (bool).

`manual` - (Optional) Manually setup routing on spoke VNet (bool).

`vnet` - (Optional) Information about existing VNet. See [Vnet ](#vnet) below for details.

### Ssh

Matches ssh port 22.

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

F5XC Software Details.

`default_sw_version` - (Optional) Will assign latest available SW version (bool).

`volterra_software_version` - (Optional) F5XC Software Version is optional parameter, which allows to specify target SW version for particular site e.g. crt-20210329-1002. (`String`).

### Tls Intercept

Specify TLS interception configuration for the network connector.

`enable_for_all_domains` - (Optional) Enable interception for all domains (bool).

`policy` - (Optional) Policy to enable/disable specific domains, with implicit enable all domains. See [Policy ](#policy) below for details.

`custom_certificate` - (Optional) Certificates for generating intermediate certificate for TLS interception.. See [Custom Certificate ](#custom-certificate) below for details.

`volterra_certificate` - (Optional) F5XC certificates for generating intermediate certificate for TLS interception. (bool).

`trusted_ca_url` - (Optional) Custom trusted CA certificates for validating upstream server certificate (`String`).

`volterra_trusted_ca` - (Optional) Default volterra trusted CA list for validating upstream server certificate (bool).

### Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Vnet

Information about existing VNet.

`resource_group` - (Required) Resource group of existing Vnet (`String`).

`vnet_name` - (Required) Name of existing Vnet (`String`).

### Vnet Resource Group

Use the same Resource Group as the Vnet.

### Volterra Certificate

F5XC certificates for generating intermediate certificate for TLS interception..

### Volterra Trusted Ca

Default volterra trusted CA list for validating upstream server certificate.

### Voltstack Cluster

App Stack Cluster using single interface, useful for deploying K8s cluster..

`accelerated_networking` - (Optional) disruption will be seen. See [Accelerated Networking ](#accelerated-networking) below for details.

`az_nodes` - (Required) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`dc_cluster_group` - (Optional) This site is member of dc cluster group via Outside Network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (bool).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`k8s_cluster` - (Optional) Site Local K8s API access is enabled, using k8s_cluster object. See [ref](#ref) below for details.

`no_k8s_cluster` - (Optional) Site Local K8s API access is disabled (bool).

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Active Enhanced Firewall Policies ](#active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Firewall Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Firewall Policy is disabled for this site. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`sm_connection_public_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

`default_storage` - (Optional) Use standard storage class configured as AWS EBS (bool).

`storage_class_list` - (Optional) Add additional custom storage classes in kubernetes for site. See [Storage Class List ](#storage-class-list) below for details.

### Voltstack Cluster Ar

App Stack Cluster using single interface, useful for deploying K8s cluster..

`accelerated_networking` - (Optional) disruption will be seen. See [Accelerated Networking ](#accelerated-networking) below for details.

`azure_certified_hw` - (Required) Name for Azure certified hardware. (`String`).

`dc_cluster_group` - (Optional) This site is member of dc cluster group via outside network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (bool).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`k8s_cluster` - (Optional) Site Local K8s API access is enabled, using k8s_cluster object. See [ref](#ref) below for details.

`no_k8s_cluster` - (Optional) Site Local K8s API access is disabled (bool).

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Active Enhanced Firewall Policies ](#active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Firewall Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Firewall Policy is disabled for this site. (bool).

`node` - (Optional) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Node ](#node) below for details.

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`sm_connection_public_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

`default_storage` - (Optional) Use standard storage class configured as AWS EBS (bool).

`storage_class_list` - (Optional) Add additional custom storage classes in kubernetes for site. See [Storage Class List ](#storage-class-list) below for details.

### Web User Interface

Matches the web user interface port.

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured azure_vnet_site.
