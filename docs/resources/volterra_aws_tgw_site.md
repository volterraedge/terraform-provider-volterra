---

page_title: "Volterra: aws_tgw_site"

description: "The aws_tgw_site allows CRUD of Aws Tgw Site resource on Volterra SaaS"
-------------------------------------------------------------------------------------

Resource volterra_aws_tgw_site
==============================

The Aws Tgw Site allows CRUD of Aws Tgw Site resource on Volterra SaaS

~> **Note:** Please refer to [Aws Tgw Site API docs](https://docs.cloud.f5.com/docs-v2/api/views-aws-tgw-site) to learn more

Example Usage
-------------

```hcl
resource "volterra_aws_tgw_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  aws_parameters {
    admin_password {
      // One of the arguments from this list "blindfold_secret_info clear_secret_info" must be set

      blindfold_secret_info {
        decryption_provider = "value"

        location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

        store_provider = "value"
      }
    }

    aws_region = "us-east-1"

    az_nodes {
      aws_az_name = "us-west-2a"

      // One of the arguments from this list "inside_subnet reserved_inside_subnet" must be set

      reserved_inside_subnet = true
      outside_subnet {
        // One of the arguments from this list "existing_subnet_id subnet_param" must be set

        subnet_param {
          ipv4 = "10.1.2.0/24"

          ipv6 = "1234:568:abcd:9100::/64"
        }
      }
      workload_subnet {
        // One of the arguments from this list "existing_subnet_id subnet_param" must be set

        subnet_param {
          ipv4 = "10.1.2.0/24"

          ipv6 = "1234:568:abcd:9100::/64"
        }
      }
    }

    // One of the arguments from this list "aws_cred" must be set

    aws_cred {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }
    disk_size = "80"
    instance_type = "a1.xlarge"

    // One of the arguments from this list "disable_internet_vip enable_internet_vip" must be set

    enable_internet_vip = true

    // One of the arguments from this list "custom_security_group f5xc_security_group" must be set

    f5xc_security_group = true

    // One of the arguments from this list "new_vpc vpc_id" must be set

    new_vpc {
      // One of the arguments from this list "autogenerate name_tag" must be set

      name_tag = "name_tag"

      primary_ipv4 = "10.1.0.0/16"
    }
    ssh_key = "ssh-rsa AAAAB..."

    // One of the arguments from this list "existing_tgw new_tgw" must be set

    new_tgw {
      // One of the arguments from this list "system_generated user_assigned" must be set

      system_generated = true
    }

    // One of the arguments from this list "reserved_tgw_cidr tgw_cidr" must be set

    reserved_tgw_cidr = true

    // One of the arguments from this list "no_worker_nodes nodes_per_az total_nodes" must be set

    nodes_per_az = "2"
  }

  // One of the arguments from this list "block_all_services blocked_services default_blocked_services" must be set

  blocked_services {
    blocked_sevice {
      // One of the arguments from this list "dns ssh web_user_interface" can be set

      web_user_interface = true

      network_type = "network_type"
    }
  }

  // One of the arguments from this list "direct_connect_disabled direct_connect_enabled private_connectivity" must be set

  private_connectivity {
    cloud_link {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    // One of the arguments from this list "inside outside" can be set

    outside = true
  }

  // One of the arguments from this list "log_receiver logs_streaming_disabled" must be set

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

###### One of the arguments from this list "block_all_services, blocked_services, default_blocked_services" must be set

`block_all_services` - (Optional) Block DNS, SSH & WebUI services on Site (`Bool`).

`blocked_services` - (Optional) Use custom blocked services configuration, to list the services which need to be blocked. See [Blocked Services Choice Blocked Services ](#blocked-services-choice-blocked-services) below for details.

`default_blocked_services` - (Optional) Allow access to DNS, SSH services on Site (`Bool`).

`coordinates` - (Optional) Site longitude and latitude co-ordinates. See [Coordinates ](#coordinates) below for details.

`custom_dns` - (Optional) custom dns configure to the CE site. See [Custom Dns ](#custom-dns) below for details.

###### One of the arguments from this list "direct_connect_disabled, direct_connect_enabled, private_connectivity" must be set

`direct_connect_disabled` - (Optional) Disable Private Connectivity to Site (`Bool`).

`direct_connect_enabled` - (Optional) Direct Connect Connection to Site is enabled(Legacy). See [Direct Connect Choice Direct Connect Enabled ](#direct-connect-choice-direct-connect-enabled) below for details.

`private_connectivity` - (Optional) Enable Private Connectivity to Site via CloudLink. See [Direct Connect Choice Private Connectivity ](#direct-connect-choice-private-connectivity) below for details.

`kubernetes_upgrade_drain` - (Optional) Enable Kubernetes Drain during OS or SW upgrade. See [Kubernetes Upgrade Drain ](#kubernetes-upgrade-drain) below for details.

###### One of the arguments from this list "log_receiver, logs_streaming_disabled" must be set

`log_receiver` - (Optional) Select log receiver for logs streaming. See [ref](#ref) below for details.

`logs_streaming_disabled` - (Optional) Logs Streaming is disabled (`Bool`).

`offline_survivability_mode` - (Optional) Enable/Disable offline survivability mode. See [Offline Survivability Mode ](#offline-survivability-mode) below for details.

`os` - (Optional) Operating System Details. See [Os ](#os) below for details.

`performance_enhancement_mode` - (Optional) Performance Enhancement Mode to optimize for L3 or L7 networking. See [Performance Enhancement Mode ](#performance-enhancement-mode) below for details.

`sw` - (Optional) F5XC Software Details. See [Sw ](#sw) below for details.

`tags` - (Optional) It helps to manage, identify, organize, search for, and filter resources in AWS console. (`String`).

`tgw_security` - (Optional) Security Configuration for transit gateway. See [Tgw Security ](#tgw-security) below for details.

`vn_config` - (Optional) Site Network related details will be configured. See [Vn Config ](#vn-config) below for details.

`vpc_attachments` - (Optional) Note that this choice would be deprecated in the near release.. See [Vpc Attachments ](#vpc-attachments) below for details.

### Aws Parameters

Example of the managed AWS resources to name few are VPC, TGW, Route Tables etc.

`admin_password` - (Optional) Admin password user for accessing site through serial console .. See [Aws Parameters Admin Password ](#aws-parameters-admin-password) below for details.

`aws_region` - (Required) AWS Region of your services vpc, where F5XC site will be deployed. (`String`).

`az_nodes` - (Required) Only Single AZ or Three AZ(s) nodes are supported currently.. See [Aws Parameters Az Nodes ](#aws-parameters-az-nodes) below for details.

###### One of the arguments from this list "aws_cred" must be set

`aws_cred` - (Optional) Reference to AWS cloud credential object used to deploy cloud resources. See [ref](#ref) below for details.

`disk_size` - (Optional) Node disk size for all node in the F5XC site. Unit is GiB (`Int`).

`instance_type` - (Required) Instance size based on the performance. (`String`).

###### One of the arguments from this list "disable_internet_vip, enable_internet_vip" must be set

`disable_internet_vip` - (Optional) VIPs cannot be advertised to the internet directly on this Site (`Bool`).

`enable_internet_vip` - (Optional) VIPs can be advertised to the internet directly on this Site (`Bool`).

###### One of the arguments from this list "custom_security_group, f5xc_security_group" must be set

`custom_security_group` - (Optional) With this option, ingress and egress traffic will be controlled via security group ids.. See [Security Group Choice Custom Security Group ](#security-group-choice-custom-security-group) below for details.

`f5xc_security_group` - (Optional) With this option, ingress and egress traffic will be controlled via f5xc created security group. (`Bool`).

###### One of the arguments from this list "new_vpc, vpc_id" must be set

`new_vpc` - (Optional) Details needed to create new VPC. See [Service Vpc Choice New Vpc ](#service-vpc-choice-new-vpc) below for details.

`vpc_id` - (Optional) Existing VPC ID (`String`).

`ssh_key` - (Required) Public SSH key for accessing nodes of the site. (`String`).

###### One of the arguments from this list "existing_tgw, new_tgw" must be set

`existing_tgw` - (Optional) Information about existing TGW. See [Tgw Choice Existing Tgw ](#tgw-choice-existing-tgw) below for details.

`new_tgw` - (Optional) Details needed to create new TGW. See [Tgw Choice New Tgw ](#tgw-choice-new-tgw) below for details.

###### One of the arguments from this list "reserved_tgw_cidr, tgw_cidr" must be set

`reserved_tgw_cidr` - (Optional) Autogenerate and reserve a TGW CIDR Block from the Primary CIDR (`Bool`).

`tgw_cidr` - (Optional) Specify TGW CIDR block. See [Tgw Cidr Choice Tgw Cidr ](#tgw-cidr-choice-tgw-cidr) below for details.

###### One of the arguments from this list "no_worker_nodes, nodes_per_az, total_nodes" must be set

`no_worker_nodes` - (Optional) Worker nodes is set to zero (`Bool`).

`nodes_per_az` - (Optional) Desired Worker Nodes Per AZ. Max limit is up to 21 (`Int`).

`total_nodes` - (Optional) Total number of worker nodes to be deployed across all AZ's used in the Site (`Int`).

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

### Performance Enhancement Mode

Performance Enhancement Mode to optimize for L3 or L7 networking.

###### One of the arguments from this list "perf_mode_l3_enhanced, perf_mode_l7_enhanced" must be set

`perf_mode_l3_enhanced` - (Optional) Site optimized for L3 traffic processing. See [Perf Mode Choice Perf Mode L3 Enhanced ](#perf-mode-choice-perf-mode-l3-enhanced) below for details.

`perf_mode_l7_enhanced` - (Optional) Site optimized for L7 traffic processing (`Bool`).

### Sw

F5XC Software Details.

###### One of the arguments from this list "default_sw_version, volterra_software_version" must be set

`default_sw_version` - (Optional) Will assign latest available F5XC Software Version (`Bool`).

`volterra_software_version` - (Optional) Specify a F5XC Software Version to be used e.g. crt-20210329-1002. (`String`).

### Tgw Security

Security Configuration for transit gateway.

###### One of the arguments from this list "active_east_west_service_policies, east_west_service_policy_allow_all, no_east_west_policy" must be set

`active_east_west_service_policies` - (Optional) Enable service policy so east-west traffic goes via proxy. See [East West Service Policy Choice Active East West Service Policies ](#east-west-service-policy-choice-active-east-west-service-policies) below for details.

`east_west_service_policy_allow_all` - (Optional) Enable service policy with allow all so east-west traffic goes via proxy for monitoring (`Bool`).

`no_east_west_policy` - (Optional) Disable service policy so that east-west traffic does not go via proxy (`Bool`).

###### One of the arguments from this list "active_forward_proxy_policies, forward_proxy_allow_all, no_forward_proxy" must be set

`active_forward_proxy_policies` - (Optional) Enable Forward Proxy for this site and manage policies. See [Forward Proxy Choice Active Forward Proxy Policies ](#forward-proxy-choice-active-forward-proxy-policies) below for details.

`forward_proxy_allow_all` - (Optional) Enable Forward Proxy for this site and allow all requests. (`Bool`).

`no_forward_proxy` - (Optional) Disable Forward Proxy for this site (`Bool`).

###### One of the arguments from this list "active_enhanced_firewall_policies, active_network_policies, no_network_policy" must be set

`active_enhanced_firewall_policies` - (Optional) with an additional option for service insertion.. See [Network Policy Choice Active Enhanced Firewall Policies ](#network-policy-choice-active-enhanced-firewall-policies) below for details.

`active_network_policies` - (Optional) Firewall Policies active for this site.. See [Network Policy Choice Active Network Policies ](#network-policy-choice-active-network-policies) below for details.

`no_network_policy` - (Optional) Firewall Policy is disabled for this site. (`Bool`).

### Vn Config

Site Network related details will be configured.

`allowed_vip_port` - (Optional) Allowed VIP Port Configuration. See [Vn Config Allowed Vip Port ](#vn-config-allowed-vip-port) below for details.

`allowed_vip_port_sli` - (Optional) Allowed VIP Port Configuration for Inside Network. See [Vn Config Allowed Vip Port Sli ](#vn-config-allowed-vip-port-sli) below for details.

###### One of the arguments from this list "dc_cluster_group_inside_vn, dc_cluster_group_outside_vn, no_dc_cluster_group" must be set

`dc_cluster_group_inside_vn` - (Optional) This site is member of dc cluster group connected via inside network. See [ref](#ref) below for details.

`dc_cluster_group_outside_vn` - (Optional) This site is member of dc cluster group connected via outside network. See [ref](#ref) below for details.

`no_dc_cluster_group` - (Optional) This site is not a member of dc cluster group (`Bool`).

###### One of the arguments from this list "global_network_list, no_global_network" must be set

`global_network_list` - (Optional) List of global network connections. See [Global Network Choice Global Network List ](#global-network-choice-global-network-list) below for details.

`no_global_network` - (Optional) No global network to connect (`Bool`).

###### One of the arguments from this list "inside_static_routes, no_inside_static_routes" must be set

`inside_static_routes` - (Optional) Manage static routes for inside network.. See [Inside Static Route Choice Inside Static Routes ](#inside-static-route-choice-inside-static-routes) below for details.

`no_inside_static_routes` - (Optional) Static Routes disabled for inside network. (`Bool`).

###### One of the arguments from this list "no_outside_static_routes, outside_static_routes" must be set

`no_outside_static_routes` - (Optional) Static Routes disabled for outside network. (`Bool`).

`outside_static_routes` - (Optional) Manage static routes for outside network.. See [Outside Static Route Choice Outside Static Routes ](#outside-static-route-choice-outside-static-routes) below for details.

###### One of the arguments from this list "sm_connection_public_ip, sm_connection_pvt_ip" must be set

`sm_connection_public_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (`Bool`).

`sm_connection_pvt_ip` - (Optional) creating ipsec between two sites which are part of the site mesh group (`Bool`).

### Vpc Attachments

Note that this choice would be deprecated in the near release..

`vpc_list` - (Optional) List of VPC attachments to transit gateway. See [Vpc Attachments Vpc List ](#vpc-attachments-vpc-list) below for details.

### Asn Choice Auto Asn

Automatically set ASN.

### Asn Choice System Generated

F5XC will automatically assign a private ASN for TGW and F5XC Site.

### Asn Choice User Assigned

User is managing the ASN for TGW and F5XC Site..

`tgw_asn` - (Optional) TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534. (`Int`).

`volterra_site_asn` - (Optional) F5XC Site ASN. (`Int`).

### Aws Parameters Admin Password

Admin password user for accessing site through serial console ..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Aws Parameters Az Nodes

Only Single AZ or Three AZ(s) nodes are supported currently..

`aws_az_name` - (Required) AWS availability zone, must be consistent with the selected AWS region. (`String`).

###### One of the arguments from this list "inside_subnet, reserved_inside_subnet" must be set

`inside_subnet` - (Optional) Select Existing Subnet or Create New. See [Choice Inside Subnet ](#choice-inside-subnet) below for details.

`reserved_inside_subnet` - (Optional) Autogenerate and reserve a subnet from the Primary CIDR (`Bool`).

`outside_subnet` - (Required) Subnet for the outside interface of the node. See [Az Nodes Outside Subnet ](#az-nodes-outside-subnet) below for details.

`workload_subnet` - (Optional) Subnet in which workloads are launched. See [Az Nodes Workload Subnet ](#az-nodes-workload-subnet) below for details.

### Az Nodes Outside Subnet

Subnet for the outside interface of the node.

###### One of the arguments from this list "existing_subnet_id, subnet_param" must be set

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Choice Subnet Param ](#choice-subnet-param) below for details.

### Az Nodes Workload Subnet

Subnet in which workloads are launched.

###### One of the arguments from this list "existing_subnet_id, subnet_param" must be set

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Choice Subnet Param ](#choice-subnet-param) below for details.

### Blocked Services Blocked Sevice

x-displayName: "Disable Node Local Services".

###### One of the arguments from this list "dns, ssh, web_user_interface" can be set

`dns` - (Optional) Matches DNS port 53 (`Bool`).

`ssh` - (Optional) x-displayName: "SSH" (`Bool`).

`web_user_interface` - (Optional) x-displayName: "Web UI" (`Bool`).

`network_type` - (Optional) Site Local VRF on which this service will be disabled (`String`).

### Blocked Services Choice Blocked Services

Use custom blocked services configuration, to list the services which need to be blocked.

`blocked_sevice` - (Optional) x-displayName: "Disable Node Local Services". See [Blocked Services Blocked Sevice ](#blocked-services-blocked-sevice) below for details.

### Blocked Services Value Type Choice Dns

Matches DNS port 53.

### Blocked Services Value Type Choice Ssh

x-displayName: "SSH".

### Blocked Services Value Type Choice Web User Interface

x-displayName: "Web UI".

### Choice Inside Subnet

Select Existing Subnet or Create New.

###### One of the arguments from this list "existing_subnet_id, subnet_param" must be set

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Choice Subnet Param ](#choice-subnet-param) below for details.

### Choice Reserved Inside Subnet

Autogenerate and reserve a subnet from the Primary CIDR.

### Choice Subnet Param

Parameters for creating new subnet.

`ipv4` - (Required) IPv4 subnet prefix for this subnet (`String`).

`ipv6` - (Optional) IPv6 subnet prefix for this subnet (`String`).

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

### Connectivity Options Site Registration Over Direct Connect

Site Registration and Site to RE tunnels go over the AWS Direct Connect Connection.

`cloudlink_network_name` - (Required) Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support. (`String`).

### Connectivity Options Site Registration Over Internet

Site Registration and Site to RE tunnels go over the internet gateway.

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

### Direct Connect Choice Direct Connect Enabled

Direct Connect Connection to Site is enabled(Legacy).

###### One of the arguments from this list "auto_asn, custom_asn" must be set

`auto_asn` - (Optional) Automatically set ASN (`Bool`).

`custom_asn` - (Optional) Custom Autonomous System Number (`Int`).

###### One of the arguments from this list "hosted_vifs, standard_vifs" must be set

`hosted_vifs` - (Optional) and automatically associate provided hosted VIF and also setup BGP Peering.. See [Vif Choice Hosted Vifs ](#vif-choice-hosted-vifs) below for details.

`standard_vifs` - (Optional) and a user associate VIF to the DirectConnect gateway and setup BGP Peering. (`Bool`).

### Direct Connect Choice Private Connectivity

Enable Private Connectivity to Site via CloudLink.

`cloud_link` - (Required) Reference to Cloud Link. See [ref](#ref) below for details.

###### One of the arguments from this list "inside, outside" can be set

`inside` - (Optional) CloudLink will be associated, and routes will be propagated with the Site Local Inside Network of this Site (`Bool`).

`outside` - (Optional) CloudLink will be associated, and routes will be propagated with the Site Local Outside Network of this Site (`Bool`).

### East West Service Policy Choice Active East West Service Policies

Enable service policy so east-west traffic goes via proxy.

`service_policies` - (Optional) A list of references to service_policy objects.. See [ref](#ref) below for details.

### East West Service Policy Choice East West Service Policy Allow All

Enable service policy with allow all so east-west traffic goes via proxy for monitoring.

### East West Service Policy Choice No East West Policy

Disable service policy so that east-west traffic does not go via proxy.

### Forward Proxy Choice Active Forward Proxy Policies

Enable Forward Proxy for this site and manage policies.

`forward_proxy_policies` - (Required) Ordered List of Forward Proxy Policies active. See [ref](#ref) below for details.

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

### Hosted Vifs Vif List

List of Hosted VIF Config.

`vif_id` - (Required) AWS Direct Connect VIF ID that needs to be connected to the site (`String`).

###### One of the arguments from this list "other_region, same_as_site_region" must be set

`other_region` - (Optional) Other Region (`String`).

`same_as_site_region` - (Optional) Use same region as that of the Site (`Bool`).

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

### Internet Vip Choice Disable Internet Vip

VIPs cannot be advertised to the internet directly on this Site.

### Internet Vip Choice Enable Internet Vip

VIPs can be advertised to the internet directly on this Site.

### Kubernetes Upgrade Drain Enable Choice Disable Upgrade Drain

x-displayName: "Disable".

### Kubernetes Upgrade Drain Enable Choice Enable Upgrade Drain

x-displayName: "Enable".

###### One of the arguments from this list "drain_max_unavailable_node_count" must be set

`drain_max_unavailable_node_count` - (Optional) x-example: "1" (`Int`).

`drain_node_timeout` - (Required) (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is recommended to use the default value). (`Int`).

### Name Choice Autogenerate

Autogenerate the VPC Name.

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

### Port Choice Custom Ports

Custom list of ports to be allowed.

`port_ranges` - (Required) Port Ranges (`String`).

### Port Choice Disable Allowed Vip Port

HTTP Port (80) & HTTPS Port (443) will be disabled..

### Port Choice Use Http Https Port

HTTP Port (80) & HTTPS Port (443) will be allowed..

### Port Choice Use Http Port

Only HTTP Port (80) will be allowed..

### Port Choice Use Https Port

Only HTTPS Port (443) will be allowed..

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Security Group Choice Custom Security Group

With this option, ingress and egress traffic will be controlled via security group ids..

`inside_security_group_id` - (Optional) Security Group ID to be attached to SLI(Site Local Inside) Interface (`String`).

`outside_security_group_id` - (Optional) Security Group ID to be attached to SLO(Site Local Outside) Interface (`String`).

### Security Group Choice F5xc Security Group

With this option, ingress and egress traffic will be controlled via f5xc created security group..

### Service Vpc Choice New Vpc

Details needed to create new VPC.

###### One of the arguments from this list "autogenerate, name_tag" must be set

`autogenerate` - (Optional) Autogenerate the VPC Name (`Bool`).

`name_tag` - (Optional) Specify the VPC Name (`String`).

`primary_ipv4` - (Required) The Primary IPv4 block cannot be modified. All subnets prefixes in this VPC must be part of this CIDR block. (`String`).

### Site Mesh Group Choice Sm Connection Public Ip

creating ipsec between two sites which are part of the site mesh group.

### Site Mesh Group Choice Sm Connection Pvt Ip

creating ipsec between two sites which are part of the site mesh group.

### Tgw Choice Existing Tgw

Information about existing TGW.

`tgw_asn` - (Optional) TGW ASN. (`Int`).

`tgw_id` - (Optional) Existing TGW ID (`String`).

`volterra_site_asn` - (Optional) F5XC Site ASN. (`Int`).

### Tgw Choice New Tgw

Details needed to create new TGW.

###### One of the arguments from this list "system_generated, user_assigned" must be set

`system_generated` - (Optional) F5XC will automatically assign a private ASN for TGW and F5XC Site (`Bool`).

`user_assigned` - (Optional) User is managing the ASN for TGW and F5XC Site.. See [Asn Choice User Assigned ](#asn-choice-user-assigned) below for details.

### Tgw Cidr Choice Reserved Tgw Cidr

Autogenerate and reserve a TGW CIDR Block from the Primary CIDR.

### Tgw Cidr Choice Tgw Cidr

Specify TGW CIDR block.

`ipv4` - (Required) IPv4 subnet prefix for this subnet (`String`).

`ipv6` - (Optional) IPv6 subnet prefix for this subnet (`String`).

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

### Vif Choice Hosted Vifs

and automatically associate provided hosted VIF and also setup BGP Peering..

###### One of the arguments from this list "site_registration_over_direct_connect, site_registration_over_internet" can be set

`site_registration_over_direct_connect` - (Optional) Site Registration and Site to RE tunnels go over the AWS Direct Connect Connection. See [Connectivity Options Site Registration Over Direct Connect ](#connectivity-options-site-registration-over-direct-connect) below for details.

`site_registration_over_internet` - (Optional) Site Registration and Site to RE tunnels go over the internet gateway (`Bool`).

`vif_list` - (Optional) List of Hosted VIF Config. See [Hosted Vifs Vif List ](#hosted-vifs-vif-list) below for details.

### Vif Choice Standard Vifs

and a user associate VIF to the DirectConnect gateway and setup BGP Peering..

### Vif Region Choice Same As Site Region

Use same region as that of the Site.

### Vn Config Allowed Vip Port

Allowed VIP Port Configuration.

###### One of the arguments from this list "custom_ports, disable_allowed_vip_port, use_http_https_port, use_http_port, use_https_port" can be set

`custom_ports` - (Optional) Custom list of ports to be allowed. See [Port Choice Custom Ports ](#port-choice-custom-ports) below for details.

`disable_allowed_vip_port` - (Optional) HTTP Port (80) & HTTPS Port (443) will be disabled. (`Bool`).

`use_http_https_port` - (Optional) HTTP Port (80) & HTTPS Port (443) will be allowed. (`Bool`).

`use_http_port` - (Optional) Only HTTP Port (80) will be allowed. (`Bool`).

`use_https_port` - (Optional) Only HTTPS Port (443) will be allowed. (`Bool`).

### Vn Config Allowed Vip Port Sli

Allowed VIP Port Configuration for Inside Network.

###### One of the arguments from this list "custom_ports, disable_allowed_vip_port, use_http_https_port, use_http_port, use_https_port" can be set

`custom_ports` - (Optional) Custom list of ports to be allowed. See [Port Choice Custom Ports ](#port-choice-custom-ports) below for details.

`disable_allowed_vip_port` - (Optional) HTTP Port (80) & HTTPS Port (443) will be disabled. (`Bool`).

`use_http_https_port` - (Optional) HTTP Port (80) & HTTPS Port (443) will be allowed. (`Bool`).

`use_http_port` - (Optional) Only HTTP Port (80) will be allowed. (`Bool`).

`use_https_port` - (Optional) Only HTTPS Port (443) will be allowed. (`Bool`).

### Volterra Sw Version Choice Default Sw Version

Will assign latest available F5XC Software Version.

### Vpc Attachments Vpc List

List of VPC attachments to transit gateway.

`labels` - (Optional) Add labels for the VPC attachment. These labels can then be used in policies such as enhanced firewall. (`String`).

`vpc_id` - (Optional) Information about existing VPC (`String`).

### Worker Nodes No Worker Nodes

Worker nodes is set to zero.

Attribute Reference
-------------------

-	`id` - This is the id of the configured aws_tgw_site.
