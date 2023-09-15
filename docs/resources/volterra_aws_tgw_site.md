---

page_title: "Volterra: aws_tgw_site"

description: "The aws_tgw_site allows CRUD of Aws Tgw Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_aws_tgw_site
==============================

The Aws Tgw Site allows CRUD of Aws Tgw Site resource on Volterra SaaS

~> **Note:** Please refer to [Aws Tgw Site API docs](https://docs.cloud.f5.com/docs/api/views-aws-tgw-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_aws_tgw_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  aws_parameters {
    aws_certified_hw = "aws-byol-multi-nic-voltmesh"
    aws_region       = "us-east-1"

    az_nodes {
      aws_az_name = "us-west-2a"

      // One of the arguments from this list "reserved_inside_subnet inside_subnet" must be set
      reserved_inside_subnet = true
      disk_size              = "80"

      outside_subnet {
        // One of the arguments from this list "subnet_param existing_subnet_id" must be set

        subnet_param {
          ipv4 = "10.1.2.0/24"
          ipv6 = "1234:568:abcd:9100::/64"
        }
      }

      workload_subnet {
        // One of the arguments from this list "existing_subnet_id subnet_param" must be set
        existing_subnet_id = "subnet-12345678901234567"
      }
    }

    // One of the arguments from this list "aws_cred assisted" must be set

    aws_cred {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
    disk_size     = "80"
    instance_type = "a1.xlarge"
    // One of the arguments from this list "disable_internet_vip enable_internet_vip" must be set
    disable_internet_vip = true
    // One of the arguments from this list "f5xc_security_group custom_security_group" must be set
    f5xc_security_group = true

    // One of the arguments from this list "vpc_id new_vpc" must be set

    new_vpc {
      allocate_ipv6 = true

      // One of the arguments from this list "name_tag autogenerate" must be set
      name_tag = "name_tag"

      primary_ipv4 = "10.1.0.0/16"
    }
    ssh_key = "ssh-rsa AAAAB..."

    // One of the arguments from this list "new_tgw existing_tgw" must be set

    new_tgw {
      // One of the arguments from this list "system_generated user_assigned" must be set
      system_generated = true
    }
    // One of the arguments from this list "nodes_per_az total_nodes no_worker_nodes" must be set
    nodes_per_az = "2"
  }

  // One of the arguments from this list "blocked_services default_blocked_services block_all_services" must be set
  default_blocked_services = true

  // One of the arguments from this list "direct_connect_disabled direct_connect_enabled" must be set
  direct_connect_disabled = true

  // One of the arguments from this list "logs_streaming_disabled log_receiver" must be set
  logs_streaming_disabled = true
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

`aws_parameters` - (Required) Example of the managed AWS resources to name few are VPC, TGW, Route Tables etc. See [Aws Parameters ](#aws-parameters) below for details.

`block_all_services` - (Optional) Block DNS, SSH & WebUI services on Site (bool).

`blocked_services` - (Optional) Use custom blocked services configuration, to list the services which need to be blocked. See [Blocked Services ](#blocked-services) below for details.

`default_blocked_services` - (Optional) Allow access to DNS, SSH services on Site (bool).

`coordinates` - (Optional) Site longitude and latitude co-ordinates. See [Coordinates ](#coordinates) below for details.

`direct_connect_disabled` - (Optional) Direct Connect Connection to Site is disabled (bool).

`direct_connect_enabled` - (Optional) Direct Connect Connection to Site is enabled. See [Direct Connect Enabled ](#direct-connect-enabled) below for details.

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (bool).

`offline_survivability_mode` - (Optional) Enable/Disable offline survivability mode. See [Offline Survivability Mode ](#offline-survivability-mode) below for details.

`os` - (Optional) Operating System Details. See [Os ](#os) below for details.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

`sw` - (Optional) F5XC Software Details. See [Sw ](#sw) below for details.

`tags` - (Optional) It helps to manage, identify, organize, search for, and filter resources in AWS console. (`String`).

`tgw_security` - (Optional) Security Configuration for transit gateway. See [Tgw Security ](#tgw-security) below for details.

`vn_config` - (Optional) Site Network related details will be configured. See [Vn Config ](#vn-config) below for details.

`vpc_attachments` - (Optional) Spoke VPCs to be attached to the AWS TGW Site. See [Vpc Attachments ](#vpc-attachments) below for details.

### Active East West Service Policies

Enable service policy so east-west traffic goes via proxy.

`service_policies` - (Optional) A list of references to service_policy objects.. See [ref](#ref) below for details.

### Active Enhanced Firewall Policies

with an additional option for service insertion..

`enhanced_firewall_policies` - (Required) Ordered List of Enhaned Firewall Policy active for this network firewall. See [ref](#ref) below for details.

### Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) List of Forward Proxy Policies. See [ref](#ref) below for details.

### Active Network Policies

Firewall Policies active for this site..

`network_policies` - (Required) Ordered List of Firewall Policies active for this network firewall. See [ref](#ref) below for details.

### Allowed Vip Port

Allowed VIP Port Configuration.

`custom_ports` - (Optional) Custom list of ports to be allowed. See [Custom Ports ](#custom-ports) below for details.

`disable_allowed_vip_port` - (Optional) HTTP Port (80) & HTTPS Port (443) will be disabled. (bool).

`use_http_https_port` - (Optional) HTTP Port (80) & HTTPS Port (443) will be allowed. (bool).

`use_http_port` - (Optional) Only HTTP Port (80) will be allowed. (bool).

`use_https_port` - (Optional) Only HTTPS Port (443) will be allowed. (bool).

### Allowed Vip Port Sli

Allowed VIP Port Configuration for Inside Network.

`custom_ports` - (Optional) Custom list of ports to be allowed. See [Custom Ports ](#custom-ports) below for details.

`disable_allowed_vip_port` - (Optional) HTTP Port (80) & HTTPS Port (443) will be disabled. (bool).

`use_http_https_port` - (Optional) HTTP Port (80) & HTTPS Port (443) will be allowed. (bool).

`use_http_port` - (Optional) Only HTTP Port (80) will be allowed. (bool).

`use_https_port` - (Optional) Only HTTPS Port (443) will be allowed. (bool).

### Assisted

In assisted deployment get AWS parameters generated in status of this objects and run volterra provided terraform script..

### Auto Asn

Automatically set ASN.

### Autogenerate

Autogenerate the VPC Name.

### Aws Parameters

Example of the managed AWS resources to name few are VPC, TGW, Route Tables etc.

`aws_certified_hw` - (Optional) Name for AWS certified hardware. (`String`).

`aws_region` - (Required) AWS Region of your services vpc, where F5XC site will be deployed. (`String`).

`az_nodes` - (Required) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Az Nodes ](#az-nodes) below for details.

`assisted` - (Optional) In assisted deployment get AWS parameters generated in status of this objects and run volterra provided terraform script. (bool).

`aws_cred` - (Optional) Reference to AWS cloud credential object used to deploy cloud resources. See [ref](#ref) below for details.

`disk_size` - (Optional) Node disk size for all node in the F5XC site. Unit is GiB (`Int`).

`instance_type` - (Required) Instance size based on the performance. (`String`).

`disable_internet_vip` - (Optional) VIPs cannot be advertised to the internet directly on this Site (bool).

`enable_internet_vip` - (Optional) VIPs can be advertised to the internet directly on this Site (bool).

`custom_security_group` - (Optional) With this option, ingress and egress traffic will be controlled via security group ids.. See [Custom Security Group ](#custom-security-group) below for details.

`f5xc_security_group` - (Optional) With this option, ingress and egress traffic will be controlled via f5xc created security group. (bool).

`new_vpc` - (Optional) Details needed to create new VPC. See [New Vpc ](#new-vpc) below for details.

`vpc_id` - (Optional) Existing VPC ID (`String`).

`ssh_key` - (Required) Public SSH key for accessing nodes of the site. (`String`).

`existing_tgw` - (Optional) Information about existing TGW. See [Existing Tgw ](#existing-tgw) below for details.

`new_tgw` - (Optional) Details needed to create new TGW. See [New Tgw ](#new-tgw) below for details.

`no_worker_nodes` - (Optional) Worker nodes is set to zero (bool).

`nodes_per_az` - (Optional) Desired Worker Nodes Per AZ. Max limit is up to 21 (`Int`).

`total_nodes` - (Optional) Total number of worker nodes to be deployed across all AZ's used in the Site (`Int`).

### Az Nodes

Only Single AZ or Three AZ(s) nodes are supported currently..

`aws_az_name` - (Required) AWS availability zone, must be consistent with the selected AWS region. (`String`).

`inside_subnet` - (Optional) Select Existing Subnet or Create New. See [Inside Subnet ](#inside-subnet) below for details.

`reserved_inside_subnet` - (Optional) Autogenerate and reserve a subnet from the Primary CIDR (bool).

`disk_size` - (Optional) Disk size to be used for this instance in GiB. 80 is 80 GiB (`Int`).

`outside_subnet` - (Required) Subnet for the outside interface of the node. See [Outside Subnet ](#outside-subnet) below for details.

`workload_subnet` - (Optional) Subnet in which workloads are launched. See [Workload Subnet ](#workload-subnet) below for details.

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

Use custom blocked services configuration, to list the services which need to be blocked.

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

### Custom Ports

Custom list of ports to be allowed.

`port_ranges` - (Required) Port Ranges (`String`).

### Custom Security Group

With this option, ingress and egress traffic will be controlled via security group ids..

`inside_security_group_id` - (Optional) Security Group ID to be attached to SLI(Site Local Inside) Interface (`String`).

`outside_security_group_id` - (Optional) Security Group ID to be attached to SLO(Site Local Outside) Interface (`String`).

### Custom Static Route

Use Custom static route to configure all advanced options.

`attrs` - (Optional) List of route attributes associated with the static route (`List of Strings`).

`labels` - (Optional) Add Labels for this Static Route, these labels can be used in network policy (`String`).

`nexthop` - (Optional) Nexthop for the route. See [Nexthop ](#nexthop) below for details.

`subnets` - (Required) List of route prefixes. See [Subnets ](#subnets) below for details.

### Default Os Version

Will assign latest available OS version.

### Default Sw Version

Will assign latest available SW version.

### Direct Connect Enabled

Direct Connect Connection to Site is enabled.

`auto_asn` - (Optional) Automatically set ASN (bool).

`custom_asn` - (Optional) Custom Autonomous System Number (`Int`).

`cloud_aggregated_prefix` - (Optional) Aggregated prefix from cloud to be advertised for DC side (`String`).

`dc_connect_aggregated_prefix` - (Optional) Aggregated prefix from direct connect to be advertised for Cloud side (`String`).

`hosted_vifs` - (Optional) and automatically associate provided hosted VIF and also setup BGP Peering.. See [Hosted Vifs ](#hosted-vifs) below for details.

`manual_gw` - (Optional) and a user associate AWS DirectConnect Gateway with it. (bool).

`standard_vifs` - (Optional) and a user associate VIF to the DirectConnect gateway and setup BGP Peering. (bool).

### Disable Allowed Vip Port

HTTP Port (80) & HTTPS Port (443) will be disabled..

### Disable Forward Proxy

Forward Proxy is disabled for this connector.

### Disable Interception

Disable Interception.

### Disable Internet Vip

VIPs cannot be advertised to the internet directly on this Site.

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Dns

Matches DNS port 53.

### Domain Match

Domain value or regular expression to match.

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### East West Service Policy Allow All

Enable service policy with allow all so east-west traffic goes via proxy for monitoring.

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

### Enable Internet Vip

VIPs can be advertised to the internet directly on this Site.

### Enable Offline Survivability Mode

When this feature is enabled on an existing site, the pods/services on this site will be restarted..

### Existing Tgw

Information about existing TGW.

`tgw_asn` - (Optional) TGW ASN. (`Int`).

`tgw_id` - (Optional) Existing TGW ID (`String`).

`volterra_site_asn` - (Optional) F5XC Site ASN. (`Int`).

### F5xc Security Group

With this option, ingress and egress traffic will be controlled via f5xc created security group..

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

### Hosted Vifs

and automatically associate provided hosted VIF and also setup BGP Peering..

`site_registration_over_direct_connect` - (Optional) Site Registration and Site to RE tunnels go over the AWS Direct Connect Connection. See [Site Registration Over Direct Connect ](#site-registration-over-direct-connect) below for details.

`site_registration_over_internet` - (Optional) Site Registration and Site to RE tunnels go over the internet gateway (bool).

`vif_list` - (Optional) List of Hosted VIF Config. See [Vif List ](#vif-list) below for details.

`vifs` - (Optional) VIFs (`String`).

### Inside Static Routes

Manage static routes for inside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Inside Subnet

Select Existing Subnet or Create New.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

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

### Manual Gw

and a user associate AWS DirectConnect Gateway with it..

### New Tgw

Details needed to create new TGW.

`system_generated` - (Optional) F5XC will automatically assign a private ASN for TGW and F5XC Site (bool).

`user_assigned` - (Optional) User is managing the ASN for TGW and F5XC Site.. See [User Assigned ](#user-assigned) below for details.

### New Vpc

Details needed to create new VPC.

`allocate_ipv6` - (Optional) Allocate IPv6 CIDR block from AWS (`Bool`).

`autogenerate` - (Optional) Autogenerate the VPC Name (bool).

`name_tag` - (Optional) Specify the VPC Name (`String`).

`primary_ipv4` - (Required) The Primary IPv4 block cannot be modified. All subnets prefixes in this VPC must be part of this CIDR block. (`String`).

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

### No East West Policy

Disable service policy so that east-west traffic does not go via proxy.

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

### No Network Policy

Firewall Policy is disabled for this site..

### No Offline Survivability Mode

When this feature is disabled on an existing site, the pods/services on this site will be restarted..

### No Outside Static Routes

Static Routes disabled for outside network..

### No Worker Nodes

Worker nodes is set to zero.

### Offline Survivability Mode

Enable/Disable offline survivability mode.

`enable_offline_survivability_mode` - (Optional) When this feature is enabled on an existing site, the pods/services on this site will be restarted. (bool).

`no_offline_survivability_mode` - (Optional) When this feature is disabled on an existing site, the pods/services on this site will be restarted. (bool).

### Os

Operating System Details.

`default_os_version` - (Optional) Will assign latest available OS version (bool).

`operating_system_version` - (Optional) Operating System Version is optional parameter, which allows to specify target OS version for particular site e.g. 7.2009.10. (`String`).

### Outside Static Routes

Manage static routes for outside network..

`static_route_list` - (Required) List of Static routes. See [Static Route List ](#static-route-list) below for details.

### Outside Subnet

Subnet for the outside interface of the node.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

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

### Reserved Inside Subnet

Autogenerate and reserve a subnet from the Primary CIDR.

### Same As Site Region

Use same region as that of the Site.

### Site Registration Over Direct Connect

Site Registration and Site to RE tunnels go over the AWS Direct Connect Connection.

`cloudlink_network_name` - (Required) Cloud Link ADN Network Name for private access connectivity to F5XC ADN. (`String`).

### Site Registration Over Internet

Site Registration and Site to RE tunnels go over the internet gateway.

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

### Ssh

Matches ssh port 22.

### Standard Vifs

and a user associate VIF to the DirectConnect gateway and setup BGP Peering..

### Static Route List

List of Static routes.

`custom_static_route` - (Optional) Use Custom static route to configure all advanced options. See [Custom Static Route ](#custom-static-route) below for details.

`simple_static_route` - (Optional) Use simple static route for prefix pointing to single interface in the network (`String`).

### Subnet Param

Parameters for creating new subnet.

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

### System Generated

F5XC will automatically assign a private ASN for TGW and F5XC Site.

### Tgw Security

Security Configuration for transit gateway.

`active_east_west_service_policies` - (Optional) Enable service policy so east-west traffic goes via proxy. See [Active East West Service Policies ](#active-east-west-service-policies) below for details.

`east_west_service_policy_allow_all` - (Optional) Enable service policy with allow all so east-west traffic goes via proxy for monitoring (bool).

`no_east_west_policy` - (Optional) Disable service policy so that east-west traffic does not go via proxy (bool).

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Active Forward Proxy Policies ](#active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (bool).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (bool).

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Active Enhanced Firewall Policies ](#active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Firewall Policies active for this site.. See [Active Network Policies ](#active-network-policies) below for details.

`no_network_policy` - (Optional) Firewall Policy is disabled for this site. (bool).

### Tls Intercept

Specify TLS interception configuration for the network connector.

`enable_for_all_domains` - (Optional) Enable interception for all domains (bool).

`policy` - (Optional) Policy to enable/disable specific domains, with implicit enable all domains. See [Policy ](#policy) below for details.

`custom_certificate` - (Optional) Certificates for generating intermediate certificate for TLS interception.. See [Custom Certificate ](#custom-certificate) below for details.

`volterra_certificate` - (Optional) F5XC certificates for generating intermediate certificate for TLS interception. (bool).

`trusted_ca_url` - (Optional) Custom trusted CA certificates for validating upstream server certificate (`String`).

`volterra_trusted_ca` - (Optional) Default volterra trusted CA list for validating upstream server certificate (bool).

### Use Http Https Port

HTTP Port (80) & HTTPS Port (443) will be allowed..

### Use Http Port

Only HTTP Port (80) will be allowed..

### Use Https Port

Only HTTPS Port (443) will be allowed..

### Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### User Assigned

User is managing the ASN for TGW and F5XC Site..

`tgw_asn` - (Optional) TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534. (`Int`).

`volterra_site_asn` - (Optional) F5XC Site ASN. (`Int`).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Vif List

List of Hosted VIF Config.

`vif_id` - (Required) AWS Direct Connect VIF ID that needs to be connected to the site (`String`).

`other_region` - (Optional) Other Region (`String`).

`same_as_site_region` - (Optional) Use same region as that of the Site (bool).

### Vn Config

Site Network related details will be configured.

`allowed_vip_port` - (Optional) Allowed VIP Port Configuration. See [Allowed Vip Port ](#allowed-vip-port) below for details.

`allowed_vip_port_sli` - (Optional) Allowed VIP Port Configuration for Inside Network. See [Allowed Vip Port Sli ](#allowed-vip-port-sli) below for details.

`dc_cluster_group_inside_vn` - (Optional) This site is member of dc cluster group connected via inside network. See [ref](#ref) below for details.

`dc_cluster_group_outside_vn` - (Optional) This site is member of dc cluster group connected via outside network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (bool).

`global_network_list` - (Optional) List of global network connections. See [Global Network List ](#global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (bool).

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Routes ](#inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (bool).

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (bool).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Routes ](#outside-static-routes) below for details.

`sm_connection_public_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (bool).

### Volterra Certificate

F5XC certificates for generating intermediate certificate for TLS interception..

### Volterra Trusted Ca

Default volterra trusted CA list for validating upstream server certificate.

### Vpc Attachments

Spoke VPCs to be attached to the AWS TGW Site.

`vpc_list` - (Optional) List of VPC attachments to transit gateway. See [Vpc List ](#vpc-list) below for details.

### Vpc List

List of VPC attachments to transit gateway.

`labels` - (Optional) These labels used must be from known key, label defined in shared namespace and unknown key. (`String`).

`vpc_id` - (Optional) Information about existing VPC (`String`).

### Web User Interface

Matches the web user interface port.

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

### Workload Subnet

Subnet in which workloads are launched.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured aws_tgw_site.
