---

page_title: "Volterra: gcp_vpc_site"

description: "The gcp_vpc_site allows CRUD of Gcp Vpc Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_gcp_vpc_site
==============================

The Gcp Vpc Site allows CRUD of Gcp Vpc Site resource on Volterra SaaS

~> **Note:** Please refer to [Gcp Vpc Site API docs](https://docs.cloud.f5.com/docs-v2/api/views-gcp-vpc-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_gcp_vpc_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "block_all_services blocked_services default_blocked_services" must be set

  default_blocked_services = true

  // One of the arguments from this list "cloud_credentials" must be set

  cloud_credentials {
    name      = "test1"
    namespace = "staging"
    tenant    = "acmecorp"
  }
  gcp_region    = ["us-west1"]
  instance_type = ["n1-standard-4"]

  // One of the arguments from this list "log_receiver logs_streaming_disabled" must be set

  logs_streaming_disabled = true

  // One of the arguments from this list "private_connect_disabled private_connectivity" must be set

  private_connect_disabled = true

  // One of the arguments from this list "ingress_egress_gw ingress_gw voltstack_cluster" must be set

  ingress_gw {
    gcp_certified_hw = "gcp-byol-voltmesh"

    gcp_zone_names = ["us-west1-a, us-west1-b, us-west1-c"]

    local_network {
      // One of the arguments from this list "existing_network new_network new_network_autogenerate" must be set

      new_network_autogenerate {
        autogenerate = true
      }
    }

    local_subnet {
      // One of the arguments from this list "existing_subnet new_subnet" must be set

      new_subnet {
        primary_ipv4 = "10.1.0.0/16"

        subnet_name = "subnet1-in-network1"
      }
    }

    node_number = "1"

    performance_enhancement_mode {
      // One of the arguments from this list "perf_mode_l3_enhanced perf_mode_l7_enhanced" must be set

      perf_mode_l3_enhanced {
        // One of the arguments from this list "jumbo no_jumbo" must be set

        jumbo = true
      }
    }
  }
  ssh_key = ["ssh-rsa AAAAB..."]
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

`address` - (Optional) Site's geographical address that can be used to determine its latitude and longitude. (`String`).

`admin_password` - (Optional) Admin password user for accessing site through serial console .. See [Admin Password ](#admin-password) below for details.

###### One of the arguments from this list "block_all_services, blocked_services, default_blocked_services" must be set

`block_all_services` - (Optional) Block DNS, SSH & WebUI services on Site (`Bool`).

`blocked_services` - (Optional) Use custom blocked services configuration. See [Blocked Services Choice Blocked Services ](#blocked-services-choice-blocked-services) below for details.

`default_blocked_services` - (Optional) Allow access to DNS, SSH services on Site (`Bool`).

`coordinates` - (Optional) Site longitude and latitude co-ordinates. See [Coordinates ](#coordinates) below for details.

`custom_dns` - (Optional) custom dns configure to the CE site. See [Custom Dns ](#custom-dns) below for details.

###### One of the arguments from this list "cloud_credentials" must be set

`cloud_credentials` - (Optional) Reference to GCP credentials for automatic deployment. See [ref](#ref) below for details.

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`gcp_labels` - (Optional) It helps to manage, identify, organize, search for, and filter resources in GCP console. (`String`).

`gcp_region` - (Required) Name for GCP Region. (`String`).

`instance_type` - (Required) Select Instance size based on performance needed (`String`).

`kubernetes_upgrade_drain` - (Optional) Enable Kubernetes Drain during OS or SW upgrade. See [Kubernetes Upgrade Drain ](#kubernetes-upgrade-drain) below for details.

###### One of the arguments from this list "log_receiver, logs_streaming_disabled" must be set

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (`Bool`).

`nodes_per_az` - (Optional) Desired Worker Nodes Per AZ. Max limit is up to 21 (`Int`).(Deprecated)

`offline_survivability_mode` - (Optional) Enable/Disable offline survivability mode. See [Offline Survivability Mode ](#offline-survivability-mode) below for details.

`os` - (Optional) Operating System Details. See [Os ](#os) below for details.

###### One of the arguments from this list "private_connect_disabled, private_connectivity" must be set

`private_connect_disabled` - (Optional)Disable Private Connectivity to Site (`Bool`).

`private_connectivity` - (Optional) Enable Private Connectivity to Site. See [Private Connectivity Choice Private Connectivity ](#private-connectivity-choice-private-connectivity) below for details.

###### One of the arguments from this list "ingress_egress_gw, ingress_gw, voltstack_cluster" must be set

`ingress_egress_gw` - (Optional) Two interface site is useful when site is used as ingress/egress gateway to the VPC network.. See [Site Type Ingress Egress Gw ](#site-type-ingress-egress-gw) below for details.

`ingress_gw` - (Optional) One interface site is useful when site is only used as ingress gateway to the VPC network.. See [Site Type Ingress Gw ](#site-type-ingress-gw) below for details.

`voltstack_cluster` - (Optional) App Stack Cluster using single interface, useful for deploying K8s cluster.. See [Site Type Voltstack Cluster ](#site-type-voltstack-cluster) below for details.

`ssh_key` - (Required) Public SSH key for accessing the site. (`String`).

`sw` - (Optional) F5XC Software Details. See [Sw ](#sw) below for details.

### Admin Password

Admin password user for accessing site through serial console ..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Admin Password Blindfold Secret Info Internal ](#admin-password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Coordinates

Site longitude and latitude co-ordinates.

`latitude` - (Optional) Latitude of the site location (`Float`).

`longitude` - (Optional) longitude of site location (`Float`).

### Custom Dns

custom dns configure to the CE site.

`inside_nameserver` - (Optional) Optional DNS server IP to be used for name resolution in inside network (`String`).

`inside_nameserver_v6` - (Optional) Optional DNS server IPv6 to be used for name resolution in inside network (`String`).

`outside_nameserver` - (Optional) Optional DNS server IP to be used for name resolution in outside network (`String`).

`outside_nameserver_v6` - (Optional) Optional DNS server IPv6 to be used for name resolution in outside network (`String`).

### Kubernetes Upgrade Drain

Enable Kubernetes Drain during OS or SW upgrade.

###### One of the arguments from this list "disable_upgrade_drain, enable_upgrade_drain" must be set

`disable_upgrade_drain` - (Optional) x-displayName: "Disable" (`Bool`).

`enable_upgrade_drain` - (Optional) x-displayName: "Enable". See [Kubernetes Upgrade Drain Enable Choice Enable Upgrade Drain ](#kubernetes-upgrade-drain-enable-choice-enable-upgrade-drain) below for details.

### Offline Survivability Mode

Enable/Disable offline survivability mode.

###### One of the arguments from this list "enable_offline_survivability_mode, no_offline_survivability_mode" must be set

`enable_offline_survivability_mode` - (Optional) x-displayName: "Enabled" (`Bool`).

`no_offline_survivability_mode` - (Optional) x-displayName: "Disabled" (`Bool`).

### Os

Operating System Details.

###### One of the arguments from this list "default_os_version, operating_system_version" must be set

`default_os_version` - (Optional) Will assign latest available OS version (`Bool`).

`operating_system_version` - (Optional) Specify a OS version to be used e.g. 9.2024.6. (`String`).

### Sw

F5XC Software Details.

###### One of the arguments from this list "default_sw_version, volterra_software_version" must be set

`default_sw_version` - (Optional) Will assign latest available F5XC Software Version (`Bool`).

`volterra_software_version` - (Optional) Specify a F5XC Software Version to be used e.g. crt-20210329-1002. (`String`).

### Admin Password Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blocked Services Blocked Sevice

x-displayName: "Disable Node Local Services".

###### One of the arguments from this list "dns, ssh, web_user_interface" can be set

`dns` - (Optional) Matches DNS port 53 (`Bool`).

`ssh` - (Optional) x-displayName: "SSH" (`Bool`).

`web_user_interface` - (Optional) x-displayName: "Web UI" (`Bool`).

`network_type` - (Optional) Site Local VRF on which this service will be disabled (`String`).

### Blocked Services Choice Blocked Services

Use custom blocked services configuration.

`blocked_sevice` - (Optional) x-displayName: "Disable Node Local Services". See [Blocked Services Blocked Sevice ](#blocked-services-blocked-sevice) below for details.

### Blocked Services Value Type Choice Dns

Matches DNS port 53.

### Blocked Services Value Type Choice Ssh

x-displayName: "SSH".

### Blocked Services Value Type Choice Web User Interface

x-displayName: "Web UI".

### Choice Existing Network

Name of existing VPC network..

`name` - (Required) Name for your GCP VPC Network (`String`).

###### One of the arguments from this list "f5_orchestrated_routing, manual_routing" can be set

`f5_orchestrated_routing` - (Optional) F5 will orchestrate required routes for SLO Route Table towards Internet and SLI RT towards the CE. (`Bool`).(Deprecated)

`manual_routing` - (Optional) In this mode, F5 will not create nor alter any route tables or routes within the existing VPCs/Vnets providing better integration for existing environments. (`Bool`).(Deprecated)

### Choice Existing Subnet

Name of existing VPC subnet..

`subnet_name` - (Required) Name of your subnet in VPC network (`String`).

### Choice New Network

Create new VPC network with specified name..

`name` - (Required) Name for your GCP VPC Network (`String`).

### Choice New Network Autogenerate

Create new VPC network with autogenerated name..

`autogenerate` - (Optional) Name for your GCP VPC Network will be autogenerated (`Bool`).(Deprecated)

### Choice New Subnet

Parameters for creating a new VPC Subnet.

`primary_ipv4` - (Required) IPv4 prefix for this Subnet. It has to be private address space. (`String`).

`subnet_name` - (Optional) Name of new VPC Subnet, will be autogenerated if empty (`String`).

### Config Mode Choice Custom Static Route

Use Custom static route to configure all advanced options.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Custom Static Route Nexthop ](#custom-static-route-nexthop) below for details.

`subnets` - (Required) List of route prefixes. See [Custom Static Route Subnets ](#custom-static-route-subnets) below for details.

### Connection Choice Sli To Global Dr

Site local inside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Connection Choice Slo To Global Dr

Site local outside is connected directly to a given global network.

`global_vn` - (Required) Select Virtual Network of Global Type. See [ref](#ref) below for details.

### Custom Certificate Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Custom Static Route Nexthop

Nexthop for the route.

`interface` - (Optional) Nexthop is network interface when type is "Network-Interface". See [ref](#ref) below for details.

`nexthop_address` - (Optional) Nexthop address when type is "Use-Configured". See [Nexthop Nexthop Address ](#nexthop-nexthop-address) below for details.

`type` - (Optional) Identifies the type of next-hop (`String`).

### Custom Static Route Subnets

List of route prefixes.

###### One of the arguments from this list "ipv4, ipv6" must be set

`ipv4` - (Optional) IPv4 Subnet Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Subnet Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Dc Cluster Group Choice No Dc Cluster Group

This site is not a member of dc cluster group.

### Enable Disable Choice Disable Interception

Disable Interception.

### Enable Disable Choice Enable Interception

Enable Interception.

### Forward Proxy Choice Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) Ordered List of Forward Proxy Policies active. See [ref](#ref) below for details.

### Forward Proxy Choice Disable Forward Proxy

Forward Proxy is disabled for this connector.

### Forward Proxy Choice Enable Forward Proxy

Forward Proxy is enabled for this connector.

`connection_timeout` - (Optional) This is specified in milliseconds. The default value is 2000 (2 seconds) (`Int`).

`max_connect_attempts` - (Optional) Specifies the allowed number of retries on connect failure to upstream server. Defaults to 1. (`Int`).

###### One of the arguments from this list "no_interception, tls_intercept" can be set

`no_interception` - (Optional) No TLS interception is enabled for this network connector (`Bool`).(Deprecated)

`tls_intercept` - (Optional) Specify TLS interception configuration for the network connector. See [Tls Interception Choice Tls Intercept ](#tls-interception-choice-tls-intercept) below for details.(Deprecated)

`white_listed_ports` - (Optional) Example "tmate" server port (`Int`).

`white_listed_prefixes` - (Optional) Example "tmate" server ip (`String`).

### Forward Proxy Choice Forward Proxy Allow All

Enable Forward Proxy for this site and allow all requests..

### Forward Proxy Choice No Forward Proxy

Disable Forward Proxy for this site.

### Global Network Choice Global Network List

List of global network connections.

`global_network_connections` - (Required) Global network connections. See [Global Network List Global Network Connections ](#global-network-list-global-network-connections) below for details.

### Global Network Choice No Global Network

No global network to connect.

### Global Network List Global Network Connections

Global network connections.

###### One of the arguments from this list "sli_to_global_dr, slo_to_global_dr" must be set

`sli_to_global_dr` - (Optional) Site local inside is connected directly to a given global network. See [Connection Choice Sli To Global Dr ](#connection-choice-sli-to-global-dr) below for details.

`slo_to_global_dr` - (Optional) Site local outside is connected directly to a given global network. See [Connection Choice Slo To Global Dr ](#connection-choice-slo-to-global-dr) below for details.

###### One of the arguments from this list "disable_forward_proxy, enable_forward_proxy" can be set

`disable_forward_proxy` - (Optional) Forward Proxy is disabled for this connector (`Bool`).(Deprecated)

`enable_forward_proxy` - (Optional) Forward Proxy is enabled for this connector. See [Forward Proxy Choice Enable Forward Proxy ](#forward-proxy-choice-enable-forward-proxy) below for details.(Deprecated)

### Ingress Egress Gw Inside Network

Network for the inside interface of the node.

###### One of the arguments from this list "existing_network, new_network, new_network_autogenerate" must be set

`existing_network` - (Optional) Name of existing VPC network.. See [Choice Existing Network ](#choice-existing-network) below for details.

`new_network` - (Optional) Create new VPC network with specified name.. See [Choice New Network ](#choice-new-network) below for details.

`new_network_autogenerate` - (Optional) Create new VPC network with autogenerated name.. See [Choice New Network Autogenerate ](#choice-new-network-autogenerate) below for details.

### Ingress Egress Gw Inside Subnet

Subnet for the inside interface of the node..

###### One of the arguments from this list "existing_subnet, new_subnet" must be set

`existing_subnet` - (Optional) Name of existing VPC subnet.. See [Choice Existing Subnet ](#choice-existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating a new VPC Subnet. See [Choice New Subnet ](#choice-new-subnet) below for details.

### Ingress Egress Gw Outside Network

Network for the outside interface of the node.

###### One of the arguments from this list "existing_network, new_network, new_network_autogenerate" must be set

`existing_network` - (Optional) Name of existing VPC network.. See [Choice Existing Network ](#choice-existing-network) below for details.

`new_network` - (Optional) Create new VPC network with specified name.. See [Choice New Network ](#choice-new-network) below for details.

`new_network_autogenerate` - (Optional) Create new VPC network with autogenerated name.. See [Choice New Network Autogenerate ](#choice-new-network-autogenerate) below for details.

### Ingress Egress Gw Outside Subnet

Subnet for the outside interface of the node..

###### One of the arguments from this list "existing_subnet, new_subnet" must be set

`existing_subnet` - (Optional) Name of existing VPC subnet.. See [Choice Existing Subnet ](#choice-existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating a new VPC Subnet. See [Choice New Subnet ](#choice-new-subnet) below for details.

### Ingress Egress Gw Performance Enhancement Mode

Performance Enhancement Mode to optimize for L3 or L7 networking.

###### One of the arguments from this list "perf_mode_l3_enhanced, perf_mode_l7_enhanced" must be set

`perf_mode_l3_enhanced` - (Optional) Site optimized for L3 traffic processing. See [Perf Mode Choice Perf Mode L3 Enhanced ](#perf-mode-choice-perf-mode-l3-enhanced) below for details.

`perf_mode_l7_enhanced` - (Optional) Site optimized for L7 traffic processing (`Bool`).

### Ingress Gw Local Network

Network for the local interface of the node.

###### One of the arguments from this list "existing_network, new_network, new_network_autogenerate" must be set

`existing_network` - (Optional) Name of existing VPC network.. See [Choice Existing Network ](#choice-existing-network) below for details.

`new_network` - (Optional) Create new VPC network with specified name.. See [Choice New Network ](#choice-new-network) below for details.

`new_network_autogenerate` - (Optional) Create new VPC network with autogenerated name.. See [Choice New Network Autogenerate ](#choice-new-network-autogenerate) below for details.

### Ingress Gw Local Subnet

Subnet for the local interface of the node..

###### One of the arguments from this list "existing_subnet, new_subnet" must be set

`existing_subnet` - (Optional) Name of existing VPC subnet.. See [Choice Existing Subnet ](#choice-existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating a new VPC Subnet. See [Choice New Subnet ](#choice-new-subnet) below for details.

### Ingress Gw Performance Enhancement Mode

Performance Enhancement Mode to optimize for L3 or L7 networking.

###### One of the arguments from this list "perf_mode_l3_enhanced, perf_mode_l7_enhanced" must be set

`perf_mode_l3_enhanced` - (Optional) Site optimized for L3 traffic processing. See [Perf Mode Choice Perf Mode L3 Enhanced ](#perf-mode-choice-perf-mode-l3-enhanced) below for details.

`perf_mode_l7_enhanced` - (Optional) Site optimized for L7 traffic processing (`Bool`).

### Inside Static Route Choice Inside Static Routes

Manage static routes for inside network..

`static_route_list` - (Required) List of Static routes. See [Inside Static Routes Static Route List ](#inside-static-routes-static-route-list) below for details.

### Inside Static Route Choice No Inside Static Routes

Static Routes disabled for inside network..

### Inside Static Routes Static Route List

List of Static routes.

###### One of the arguments from this list "custom_static_route, simple_static_route" must be set

`custom_static_route` - (Optional) Use Custom static route to configure all advanced options. See [Config Mode Choice Custom Static Route ](#config-mode-choice-custom-static-route) below for details.

`simple_static_route` - (Optional) Use simple static route for prefix pointing to single interface in the network (`String`).

### Interception Policy Choice Enable For All Domains

Enable interception for all domains.

### Interception Policy Choice Policy

Policy to enable/disable specific domains, with implicit enable all domains.

`interception_rules` - (Required) List of ordered rules to enable or disable for TLS interception. See [Policy Interception Rules ](#policy-interception-rules) below for details.

### Interception Rules Domain Match

Domain value or regular expression to match.

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### K8s Cluster Choice No K8s Cluster

Site Local K8s API access is disabled.

### Kubernetes Upgrade Drain Enable Choice Disable Upgrade Drain

x-displayName: "Disable".

### Kubernetes Upgrade Drain Enable Choice Enable Upgrade Drain

x-displayName: "Enable".

###### One of the arguments from this list "drain_max_unavailable_node_count, drain_max_unavailable_node_percentage" must be set

`drain_max_unavailable_node_count` - (Optional) x-example: "1" (`Int`).

`drain_max_unavailable_node_percentage` - (Optional) Max number of worker nodes to be upgraded in parallel by percentage. Note: 1% would mean batch size of 1 worker node. (`Int`).(Deprecated)

`drain_node_timeout` - (Required) (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is recommended to use the default value). (`Int`).

###### One of the arguments from this list "disable_vega_upgrade_mode, enable_vega_upgrade_mode" must be set

`disable_vega_upgrade_mode` - (Optional) Disable Vega Upgrade Mode (`Bool`).(Deprecated)

`enable_vega_upgrade_mode` - (Optional) When enabled, vega will inform RE to stop traffic to the specific node. (`Bool`).(Deprecated)

### Network Options Inside

CloudLink will be associated, and routes will be propagated with the Site Local Inside Network of this Site.

### Network Options Outside

CloudLink will be associated, and routes will be propagated with the Site Local Outside Network of this Site.

### Network Policy Choice Active Enhanced Firewall Policies

with an additional option for service insertion..

`enhanced_firewall_policies` - (Required) Ordered List of Enhanced Firewall Policies active. See [ref](#ref) below for details.

### Network Policy Choice Active Network Policies

Firewall Policies active for this site..

`network_policies` - (Required) Ordered List of Firewall Policies active for this network firewall. See [ref](#ref) below for details.

### Network Policy Choice No Network Policy

Firewall Policy is disabled for this site..

### Nexthop Nexthop Address

Nexthop address when type is "Use-Configured".

###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.

`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Offline Survivability Mode Choice Enable Offline Survivability Mode

x-displayName: "Enabled".

### Offline Survivability Mode Choice No Offline Survivability Mode

x-displayName: "Disabled".

### Operating System Version Choice Default Os Version

Will assign latest available OS version.

### Outside Static Route Choice No Outside Static Routes

Static Routes disabled for outside network..

### Outside Static Route Choice Outside Static Routes

Manage static routes for outside network..

`static_route_list` - (Required) List of Static routes. See [Outside Static Routes Static Route List ](#outside-static-routes-static-route-list) below for details.

### Outside Static Routes Static Route List

List of Static routes.

###### One of the arguments from this list "custom_static_route, simple_static_route" must be set

`custom_static_route` - (Optional) Use Custom static route to configure all advanced options. See [Config Mode Choice Custom Static Route ](#config-mode-choice-custom-static-route) below for details.

`simple_static_route` - (Optional) Use simple static route for prefix pointing to single interface in the network (`String`).

### Perf Mode Choice Jumbo

x-displayName: "Enabled".

### Perf Mode Choice No Jumbo

x-displayName: "Disabled".

### Perf Mode Choice Perf Mode L3 Enhanced

Site optimized for L3 traffic processing.

###### One of the arguments from this list "jumbo, no_jumbo" must be set

`jumbo` - (Optional) x-displayName: "Enabled" (`Bool`).

`no_jumbo` - (Optional) x-displayName: "Disabled" (`Bool`).

### Perf Mode Choice Perf Mode L7 Enhanced

Site optimized for L7 traffic processing.

### Policy Interception Rules

List of ordered rules to enable or disable for TLS interception.

`domain_match` - (Required) Domain value or regular expression to match. See [Interception Rules Domain Match ](#interception-rules-domain-match) below for details.

###### One of the arguments from this list "disable_interception, enable_interception" must be set

`disable_interception` - (Optional) Disable Interception (`Bool`).

`enable_interception` - (Optional) Enable Interception (`Bool`).

### Private Connectivity Choice Private Connectivity

Enable Private Connectivity to Site.

`cloud_link` - (Required) Reference to Cloud Link. See [ref](#ref) below for details.

###### One of the arguments from this list "inside, outside" can be set

`inside` - (Optional) CloudLink will be associated, and routes will be propagated with the Site Local Inside Network of this Site (`Bool`).

`outside` - (Optional) CloudLink will be associated, and routes will be propagated with the Site Local Outside Network of this Site (`Bool`).

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Routing Type F5 Orchestrated Routing

F5 will orchestrate required routes for SLO Route Table towards Internet and SLI RT towards the CE..

### Routing Type Manual Routing

In this mode, F5 will not create nor alter any route tables or routes within the existing VPCs/Vnets providing better integration for existing environments. .

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

### Signing Cert Choice Custom Certificate

Certificates for generating intermediate certificate for TLS interception..

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Custom Certificate Private Key ](#custom-certificate-private-key) below for details.

### Signing Cert Choice Volterra Certificate

F5XC certificates for generating intermediate certificate for TLS interception..

### Site Mesh Group Choice Sm Connection Public Ip

creating ipsec between two sites which are part of the site mesh group.

### Site Mesh Group Choice Sm Connection Pvt Ip

creating ipsec between two sites which are part of the site mesh group.

### Site Type Ingress Egress Gw

Two interface site is useful when site is used as ingress/egress gateway to the VPC network..

###### One of the arguments from this list "dc_cluster_group_inside_vn, dc_cluster_group_outside_vn, no_dc_cluster_group" must be set

`dc_cluster_group_inside_vn` - (Optional) This site is member of dc cluster group connected via inside network. See [ref](#ref) below for details.

`dc_cluster_group_outside_vn` - (Optional) This site is member of dc cluster group connected via outside network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (`Bool`).

###### One of the arguments from this list "active_forward_proxy_policies, forward_proxy_allow_all, no_forward_proxy" must be set

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Forward Proxy Choice Active Forward Proxy Policies ](#forward-proxy-choice-active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (`Bool`).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (`Bool`).

`gcp_certified_hw` - (Required) Name for GCP certified hardware. (`String`).

`gcp_zone_names` - (Required) List of zones when instances will be created, needs to match with region selected. (`String`).

###### One of the arguments from this list "global_network_list, no_global_network" must be set

`global_network_list` - (Optional) List of global network connections. See [Global Network Choice Global Network List ](#global-network-choice-global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (`Bool`).

`inside_network` - (Optional) Network for the inside interface of the node. See [Ingress Egress Gw Inside Network ](#ingress-egress-gw-inside-network) below for details.

###### One of the arguments from this list "inside_static_routes, no_inside_static_routes" must be set

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Route Choice Inside Static Routes ](#inside-static-route-choice-inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (`Bool`).

`inside_subnet` - (Optional) Subnet for the inside interface of the node.. See [Ingress Egress Gw Inside Subnet ](#ingress-egress-gw-inside-subnet) below for details.

###### One of the arguments from this list "active_enhanced_firewall_policies, active_network_policies, no_network_policy" must be set

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Network Policy Choice Active Enhanced Firewall Policies ](#network-policy-choice-active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Firewall Policies active for this site.. See [Network Policy Choice Active Network Policies ](#network-policy-choice-active-network-policies) below for details.

`no_network_policy` - (Optional) Firewall Policy is disabled for this site. (`Bool`).

`node_number` - (Optional) Number of main nodes to create, either 1 or 3. (`Int`).

`outside_network` - (Optional) Network for the outside interface of the node. See [Ingress Egress Gw Outside Network ](#ingress-egress-gw-outside-network) below for details.

###### One of the arguments from this list "no_outside_static_routes, outside_static_routes" must be set

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (`Bool`).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Route Choice Outside Static Routes ](#outside-static-route-choice-outside-static-routes) below for details.

`outside_subnet` - (Optional) Subnet for the outside interface of the node.. See [Ingress Egress Gw Outside Subnet ](#ingress-egress-gw-outside-subnet) below for details.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Ingress Egress Gw Performance Enhancement Mode ](#ingress-egress-gw-performance-enhancement-mode) below for details.

###### One of the arguments from this list "sm_connection_public_ip, sm_connection_pvt_ip" must be set

`sm_connection_public_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (`Bool`).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (`Bool`).

### Site Type Ingress Gw

One interface site is useful when site is only used as ingress gateway to the VPC network..

`gcp_certified_hw` - (Required) Name for GCP certified hardware. (`String`).

`gcp_zone_names` - (Required) List of zones when instances will be created, needs to match with region selected. (`String`).

`local_network` - (Optional) Network for the local interface of the node. See [Ingress Gw Local Network ](#ingress-gw-local-network) below for details.

`local_subnet` - (Optional) Subnet for the local interface of the node.. See [Ingress Gw Local Subnet ](#ingress-gw-local-subnet) below for details.

`node_number` - (Optional) Number of main nodes to create, either 1 or 3. (`Int`).

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Ingress Gw Performance Enhancement Mode ](#ingress-gw-performance-enhancement-mode) below for details.

### Site Type Voltstack Cluster

App Stack Cluster using single interface, useful for deploying K8s cluster..

###### One of the arguments from this list "dc_cluster_group, no_dc_cluster_group" must be set

`dc_cluster_group` - (Optional) This site is member of dc cluster group connected via outside network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (`Bool`).

###### One of the arguments from this list "active_forward_proxy_policies, forward_proxy_allow_all, no_forward_proxy" must be set

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Forward Proxy Choice Active Forward Proxy Policies ](#forward-proxy-choice-active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (`Bool`).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (`Bool`).

`gcp_certified_hw` - (Required) Name for GCP certified hardware. (`String`).

`gcp_zone_names` - (Required) List of zones when instances will be created, needs to match with region selected. (`String`).

###### One of the arguments from this list "global_network_list, no_global_network" must be set

`global_network_list` - (Optional) List of global network connections. See [Global Network Choice Global Network List ](#global-network-choice-global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (`Bool`).

###### One of the arguments from this list "k8s_cluster, no_k8s_cluster" must be set

`k8s_cluster` - (Optional) Site Local K8s API access is enabled, using k8s_cluster object. See [ref](#ref) below for details.

`no_k8s_cluster` - (Optional) Site Local K8s API access is disabled (`Bool`).

###### One of the arguments from this list "active_enhanced_firewall_policies, active_network_policies, no_network_policy" must be set

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Network Policy Choice Active Enhanced Firewall Policies ](#network-policy-choice-active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Firewall Policies active for this site.. See [Network Policy Choice Active Network Policies ](#network-policy-choice-active-network-policies) below for details.

`no_network_policy` - (Optional) Firewall Policy is disabled for this site. (`Bool`).

`node_number` - (Optional) Number of main nodes to create, either 1 or 3. (`Int`).

###### One of the arguments from this list "no_outside_static_routes, outside_static_routes" must be set

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (`Bool`).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Route Choice Outside Static Routes ](#outside-static-route-choice-outside-static-routes) below for details.

`site_local_network` - (Optional) Network for the local interface of the node. See [Voltstack Cluster Site Local Network ](#voltstack-cluster-site-local-network) below for details.

`site_local_subnet` - (Optional) Subnet for the local interface of the node.. See [Voltstack Cluster Site Local Subnet ](#voltstack-cluster-site-local-subnet) below for details.

###### One of the arguments from this list "sm_connection_public_ip, sm_connection_pvt_ip" must be set

`sm_connection_public_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (`Bool`).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (`Bool`).

###### One of the arguments from this list "default_storage, storage_class_list" must be set

`default_storage` - (Optional) Use standard storage class configured as AWS EBS (`Bool`).

`storage_class_list` - (Optional) Add additional custom storage classes in kubernetes for site. See [Storage Class Choice Storage Class List ](#storage-class-choice-storage-class-list) below for details.

### Storage Class Choice Default Storage

Use standard storage class configured as AWS EBS.

### Storage Class Choice Storage Class List

Add additional custom storage classes in kubernetes for site.

`storage_classes` - (Optional) List of custom storage classes. See [Storage Class List Storage Classes ](#storage-class-list-storage-classes) below for details.

### Storage Class List Storage Classes

List of custom storage classes.

`default_storage_class` - (Optional) Make this storage class default storage class for the K8s cluster (`Bool`).

`storage_class_name` - (Required) Name of the storage class as it will appear in K8s. (`String`).

### Tls Interception Choice No Interception

No TLS interception is enabled for this network connector.

### Tls Interception Choice Tls Intercept

Specify TLS interception configuration for the network connector.

###### One of the arguments from this list "enable_for_all_domains, policy" must be set

`enable_for_all_domains` - (Optional) Enable interception for all domains (`Bool`).

`policy` - (Optional) Policy to enable/disable specific domains, with implicit enable all domains. See [Interception Policy Choice Policy ](#interception-policy-choice-policy) below for details.

###### One of the arguments from this list "custom_certificate, volterra_certificate" must be set

`custom_certificate` - (Optional) Certificates for generating intermediate certificate for TLS interception.. See [Signing Cert Choice Custom Certificate ](#signing-cert-choice-custom-certificate) below for details.

`volterra_certificate` - (Optional) F5XC certificates for generating intermediate certificate for TLS interception. (`Bool`).

###### One of the arguments from this list "trusted_ca_url, volterra_trusted_ca" must be set

`trusted_ca_url` - (Optional) Custom Root CA Certificate for validating upstream server certificate (`String`).

`volterra_trusted_ca` - (Optional) F5XC Root CA Certificate for validating upstream server certificate (`Bool`).

### Trusted Ca Choice Volterra Trusted Ca

F5XC Root CA Certificate for validating upstream server certificate.

### Vega Upgrade Mode Toggle Choice Disable Vega Upgrade Mode

Disable Vega Upgrade Mode.

### Vega Upgrade Mode Toggle Choice Enable Vega Upgrade Mode

When enabled, vega will inform RE to stop traffic to the specific node..

### Ver Ipv4

IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).

### Ver Ipv4

IPv4 Subnet Address.

`plen` - (Optional) Prefix-length of the IPv4 subnet. Must be <= 32 (`Int`).

`prefix` - (Optional) Prefix part of the IPv4 subnet in string form with dot-decimal notation (`String`).

### Ver Ipv6

IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).

### Ver Ipv6

IPv6 Subnet Address.

`plen` - (Optional) Prefix length of the IPv6 subnet. Must be <= 128 (`Int`).

`prefix` - (Optional) e.g. "2001:db8::2::" (`String`).

### Volterra Sw Version Choice Default Sw Version

Will assign latest available F5XC Software Version.

### Voltstack Cluster Site Local Network

Network for the local interface of the node.

###### One of the arguments from this list "existing_network, new_network, new_network_autogenerate" must be set

`existing_network` - (Optional) Name of existing VPC network.. See [Choice Existing Network ](#choice-existing-network) below for details.

`new_network` - (Optional) Create new VPC network with specified name.. See [Choice New Network ](#choice-new-network) below for details.

`new_network_autogenerate` - (Optional) Create new VPC network with autogenerated name.. See [Choice New Network Autogenerate ](#choice-new-network-autogenerate) below for details.

### Voltstack Cluster Site Local Subnet

Subnet for the local interface of the node..

###### One of the arguments from this list "existing_subnet, new_subnet" must be set

`existing_subnet` - (Optional) Name of existing VPC subnet.. See [Choice Existing Subnet ](#choice-existing-subnet) below for details.

`new_subnet` - (Optional) Parameters for creating a new VPC Subnet. See [Choice New Subnet ](#choice-new-subnet) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured gcp_vpc_site.
