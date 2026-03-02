---

page_title: "Volterra: vs_pool"

description: "The vs_pool allows CRUD of Vs Pool resource on Volterra SaaS"
---------------------------------------------------------------------------

Resource volterra_vs_pool
=========================

The Vs Pool allows CRUD of Vs Pool resource on Volterra SaaS

~> **Note:** Please refer to [Vs Pool API docs](https://docs.cloud.f5.com/docs-v2/api/views-vs-pool) to learn more

Example Usage
-------------

```hcl
resource "volterra_vs_pool" "example" {
  name               = "acmecorp-web"
  namespace          = "staging"
  endpoint_selection = ["endpoint_selection"]
  lb_method          = ["lb_method"]

  pool_members {
    description = "value"

    member_selection {
      connection_limit = "connection_limit"

      connection_rate_limit = "connection_rate_limit"

      priority_group = "priority_group"

      ratio = "ratio"
    }

    // One of the arguments from this list "private public" must be set

    public {
      // One of the arguments from this list "all_re regional_re" can be set

      all_re {
        // One of the arguments from this list "custom_ip dns" can be set

        dns {
          dns_name = "value"

          refresh_interval = "20"
        }

        // One of the arguments from this list "f5_global_network" can be set

        f5_global_network = true
      }
    }
    name = "name"
    port = "9080"
    state {
      // One of the arguments from this list "member_state_disabled member_state_enabled member_state_offline" must be set

      member_state_enabled {}
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

`action_on_service_down` - (Optional) Reselect: Specifies that the system manages established client connections by moving them to an alternative pool member when monitors mark the original pool member down.. See [Action On Service Down ](#action-on-service-down) below for details.

`endpoint_selection` - (Required) Policy for selection of endpoints from local site or remote site or both (`String`).

`health_monitoring` - (Optional) Specifies an association between a health or performance monitor and an entire pool, rather than with individual pool members.. See [Health Monitoring ](#health-monitoring) below for details.

`lb_method` - (Required) Specifies the load balancing method. The default is Round Robin (`String`).

`pool_members` - (Required) List of members in this pool. See [Pool Members ](#pool-members) below for details.

`priority_group` - (Optional) Specifies whether the system load balances traffic according to the priority number assigned to the pool member. The default is Disabled. Once you enable this setting, you can specify pool member priority when you create a new pool or on a pool member's properties screen. The system treats same-priority pool members as a group. To enable priority group activation, select Less than from the list, and in the Available Member(s) field, type a number from 0 to 65535 that represents the minimum number of members that must be available in one priority group before the system directs traffic to members in a lower priority group. When a sufficient number of members become available in the higher priority group, the system again directs traffic to the higher priority group.. See [Priority Group ](#priority-group) below for details.

`request_queue_config` - (Optional) Enables TCP request queueing.. See [Request Queue Config ](#request-queue-config) below for details.

`slow_ramp_time` - (Optional) Setting this to a nonzero value can cause unexpected Priority Group behavior, such as load balancing to a low-priority member even with enough high-priority servers. (`Int`).

### Action On Service Down

Reselect: Specifies that the system manages established client connections by moving them to an alternative pool member when monitors mark the original pool member down..

###### One of the arguments from this list "drop, none, reject, reselect" can be set

`drop` - (Optional) x-displayName: "Drop". See [Action On Down Choice Drop ](#action-on-down-choice-drop) below for details.

`none` - (Optional) x-displayName: "None". See [Action On Down Choice None ](#action-on-down-choice-none) below for details.

`reject` - (Optional) x-displayName: "Reject". See [Action On Down Choice Reject ](#action-on-down-choice-reject) below for details.

`reselect` - (Optional) x-displayName: "Reselect". See [Action On Down Choice Reselect ](#action-on-down-choice-reselect) below for details.

### Health Monitoring

Specifies an association between a health or performance monitor and an entire pool, rather than with individual pool members..

###### One of the arguments from this list "health_monitor_availability_all, health_monitor_availability_at_least" can be set

`health_monitor_availability_all` - (Optional) x-displayName: "All". See [Health Monitor Availability Choice Health Monitor Availability All ](#health-monitor-availability-choice-health-monitor-availability-all) below for details.

`health_monitor_availability_at_least` - (Optional) x-displayName: "At least". See [Health Monitor Availability Choice Health Monitor Availability At Least ](#health-monitor-availability-choice-health-monitor-availability-at-least) below for details.

`healthcheck` - (Optional) Reference to healthcheck configuration objects. See [ref](#ref) below for details.

### Pool Members

List of members in this pool.

`description` - (Optional) Description for this Member (`String`).

`member_selection` - (Optional) x-displayName: "Pool Member Selection". See [Pool Members Member Selection ](#pool-members-member-selection) below for details.

###### One of the arguments from this list "private, public" must be set

`private` - (Optional) x-displayName: "Private". See [Member Type Choice Private ](#member-type-choice-private) below for details.

`public` - (Optional) x-displayName: "Public". See [Member Type Choice Public ](#member-type-choice-public) below for details.

`name` - (Required) Name of this Member (`String`).

`port` - (Optional) Endpoint service is available on this port (`Int`).

`state` - (Required) Displays the current status of the pool member.. See [Pool Members State ](#pool-members-state) below for details.

### Priority Group

Specifies whether the system load balances traffic according to the priority number assigned to the pool member. The default is Disabled. Once you enable this setting, you can specify pool member priority when you create a new pool or on a pool member's properties screen. The system treats same-priority pool members as a group. To enable priority group activation, select Less than from the list, and in the Available Member(s) field, type a number from 0 to 65535 that represents the minimum number of members that must be available in one priority group before the system directs traffic to members in a lower priority group. When a sufficient number of members become available in the higher priority group, the system again directs traffic to the higher priority group..

###### One of the arguments from this list "disabled_priority_group, enabled_priority_group" can be set

`disabled_priority_group` - (Optional) x-displayName: "Disabled". See [Priority Group Activation Choice Disabled Priority Group ](#priority-group-activation-choice-disabled-priority-group) below for details.

`enabled_priority_group` - (Optional) x-displayName: "Enabled". See [Priority Group Activation Choice Enabled Priority Group ](#priority-group-activation-choice-enabled-priority-group) below for details.

### Request Queue Config

Enables TCP request queueing..

###### One of the arguments from this list "request_queuing_disabled, request_queuing_enabled" can be set

`request_queuing_disabled` - (Optional) x-displayName: "Disabled". See [Request Queuing Choice Request Queuing Disabled ](#request-queuing-choice-request-queuing-disabled) below for details.

`request_queuing_enabled` - (Optional) x-displayName: "Enabled". See [Request Queuing Choice Request Queuing Enabled ](#request-queuing-choice-request-queuing-enabled) below for details.

### Action On Down Choice Drop

x-displayName: "Drop".

### Action On Down Choice None

x-displayName: "None".

### Action On Down Choice Reject

x-displayName: "Reject".

### Action On Down Choice Reselect

x-displayName: "Reselect".

`reselect_count` - (Optional) Specifies the number of times the system tries to contact a new pool member after a passive failure. A passive failure consists of a server-connect failure or a failure to receive a data response within a user-specified interval. The default is 0 (zero), which indicates no reselects. (`Int`).

### Discovered Service Choice Cbip Service

Specify origin server with Classic BIG-IP service name.

`cbip_service_name` - (Required) Name of the discovered Classic BIG-IP virtual server to be used as origin. (`String`).

### Discovered Service Choice Consul Service

Specify origin server with Hashi Corp Consul service name and site information.

`consul_service_name` - (Required) The format is servicename:cluster-id. (`String`).

### Discovered Service Choice K8s Service

Specify origin server with K8s service name and site information.

`k8s_service_name` - (Optional) Both namespace and cluster-id are optional. (`String`).

### Health Monitor Availability Choice Health Monitor Availability All

x-displayName: "All".

### Health Monitor Availability Choice Health Monitor Availability At Least

x-displayName: "At least".

`minimum_available` - (Optional) Minimum Availability (`Int`).

### Member State Choice Member State Disabled

x-displayName: "Disabled".

### Member State Choice Member State Enabled

x-displayName: "Enabled".

### Member State Choice Member State Offline

x-displayName: "Forced Offline".

### Member Type Choice Custom Ip

IP Address.

###### One of the arguments from this list "ip, ip_v6" can be set

`ip` - (Optional) V4 IP (`String`).

`ip_v6` - (Optional) Optional common virtual V6 IP across all nodes to be used as automatic VIP. (`String`).

### Member Type Choice Discovered

Discovered.

###### One of the arguments from this list "cbip_service, consul_service, k8s_service" can be set

`cbip_service` - (Optional) Specify origin server with Classic BIG-IP service name. See [Discovered Service Choice Cbip Service ](#discovered-service-choice-cbip-service) below for details.

`consul_service` - (Optional) Specify origin server with Hashi Corp Consul service name and site information. See [Discovered Service Choice Consul Service ](#discovered-service-choice-consul-service) below for details.

`k8s_service` - (Optional) Specify origin server with K8s service name and site information. See [Discovered Service Choice K8s Service ](#discovered-service-choice-k8s-service) below for details.

### Member Type Choice Dns

DNS.

`dns_name` - (Required) DNS Name (`String`).

`refresh_interval` - (Optional) Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 (`Int`).

### Member Type Choice Private

x-displayName: "Private".

###### One of the arguments from this list "site, virtual_ce_sites, virtual_network" can be set

`site` - (Optional) x-displayName: "Sites". See [Where Type Choice Site ](#where-type-choice-site) below for details.

`virtual_ce_sites` - (Optional) x-displayName: "Virtual Sites". See [Where Type Choice Virtual Ce Sites ](#where-type-choice-virtual-ce-sites) below for details.

`virtual_network` - (Optional) x-displayName: "Segment". See [Where Type Choice Virtual Network ](#where-type-choice-virtual-network) below for details.

### Member Type Choice Public

x-displayName: "Public".

###### One of the arguments from this list "all_re, regional_re" can be set

`all_re` - (Optional) x-displayName: "Global". See [Where Selection Choice All Re ](#where-selection-choice-all-re) below for details.

`regional_re` - (Optional) x-displayName: "Regional". See [Where Selection Choice Regional Re ](#where-selection-choice-regional-re) below for details.

### Network Choice Site Local Inside Network

x-displayName: "Site Local Inside (Local VRF)".

### Network Choice Site Local Network

x-displayName: "Site Local Outside (Local VRF)".

### Network Selection Choice F5 Global Network

x-displayName: "F5 Global Network".

### Network Selection Choice Network Option

Select Network.

###### One of the arguments from this list "segment_network, site_local_inside_network, site_local_network" can be set

`segment_network` - (Optional) x-displayName: "Segment (Global VRF)". See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) x-displayName: "Site Local Inside (Local VRF)" (`Bool`).

`site_local_network` - (Optional) x-displayName: "Site Local Outside (Local VRF)" (`Bool`).

### Pool Members Member Selection

x-displayName: "Pool Member Selection".

`connection_limit` - (Optional) Specifies a maximum established connection limit for a pool member or node. When the current connections count reaches this number, the system does not send additional connections to that pool member or node. The default is 0, meaning that there is no limit to the number of connections. When used with the weighted least connections load balancing methods, the system uses connection limits to determine the proportional load of each pool member or node. This must be a value other than 0 when specified for the weighted least connections load balancing methods. (`Int`).

`connection_rate_limit` - (Optional) Connection Rate Limit (`Int`).

`priority_group` - (Optional) Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority. To specify a priority, you must activate priority group usage when you create a new pool or when adding or removing pool members. When activated, the system load balances traffic according to the priority group number assigned to the pool member. The higher the number, the higher the priority, so a member with a priority of 3 has higher priority than a member with a priority of 1. (`Int`).

`ratio` - (Optional) Displays the ratio value for each pool member. This value is used by ratio-related load balancing methods to load balance connections. (`Int`).

### Pool Members State

Displays the current status of the pool member..

###### One of the arguments from this list "member_state_disabled, member_state_enabled, member_state_offline" must be set

`member_state_disabled` - (Optional) x-displayName: "Disabled". See [Member State Choice Member State Disabled ](#member-state-choice-member-state-disabled) below for details.

`member_state_enabled` - (Optional) x-displayName: "Enabled". See [Member State Choice Member State Enabled ](#member-state-choice-member-state-enabled) below for details.

`member_state_offline` - (Optional) x-displayName: "Forced Offline". See [Member State Choice Member State Offline ](#member-state-choice-member-state-offline) below for details.

### Priority Group Activation Choice Disabled Priority Group

x-displayName: "Disabled".

### Priority Group Activation Choice Enabled Priority Group

x-displayName: "Enabled".

`count` - (Optional) Specify pool member priority (`Int`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Request Queuing Choice Request Queuing Disabled

x-displayName: "Disabled".

### Request Queuing Choice Request Queuing Enabled

x-displayName: "Enabled".

`queue_depth` - (Optional) Specifies the maximum number of connection requests allowed in the queue. The default value of 0 equates to unlimited connection requests, constrained only by available memory. (`Int`).

`queue_timeout` - (Optional) Specifies the maximum number of milliseconds that a connection request can be queued until capacity becomes available, whereupon the connection request is removed from the queue and reset. The default value of 0 equates to an unlimited time in the queue. (`Int`).

### Where Selection Choice All Re

x-displayName: "Global".

###### One of the arguments from this list "custom_ip, dns" can be set

`custom_ip` - (Optional) IP Address. See [Member Type Choice Custom Ip ](#member-type-choice-custom-ip) below for details.

`dns` - (Optional) DNS. See [Member Type Choice Dns ](#member-type-choice-dns) below for details.

###### One of the arguments from this list "f5_global_network" can be set

`f5_global_network` - (Optional) x-displayName: "F5 Global Network" (`Bool`).

### Where Selection Choice Regional Re

x-displayName: "Regional".

###### One of the arguments from this list "custom_ip, dns" can be set

`custom_ip` - (Optional) IP Address. See [Member Type Choice Custom Ip ](#member-type-choice-custom-ip) below for details.

`dns` - (Optional) DNS. See [Member Type Choice Dns ](#member-type-choice-dns) below for details.

###### One of the arguments from this list "f5_global_network" can be set

`f5_global_network` - (Optional) x-displayName: "F5 Global Network" (`Bool`).

`regional_edge_site` - (Optional) Reference to site object. See [ref](#ref) below for details.

### Where Type Choice Site

x-displayName: "Sites".

###### One of the arguments from this list "custom_ip, discovered, dns" can be set

`custom_ip` - (Optional) IP Address. See [Member Type Choice Custom Ip ](#member-type-choice-custom-ip) below for details.

`discovered` - (Optional) Discovered. See [Member Type Choice Discovered ](#member-type-choice-discovered) below for details.

`dns` - (Optional) DNS. See [Member Type Choice Dns ](#member-type-choice-dns) below for details.

###### One of the arguments from this list "f5_global_network, network_option" can be set

`f5_global_network` - (Optional) x-displayName: "F5 Global Network" (`Bool`).

`network_option` - (Optional) Select Network. See [Network Selection Choice Network Option ](#network-selection-choice-network-option) below for details.

`site_name` - (Optional) x-example: "smsv2-site-name". See [ref](#ref) below for details.

### Where Type Choice Virtual Ce Sites

x-displayName: "Virtual Sites".

###### One of the arguments from this list "custom_ip, discovered, dns" can be set

`custom_ip` - (Optional) IP Address. See [Member Type Choice Custom Ip ](#member-type-choice-custom-ip) below for details.

`discovered` - (Optional) Discovered. See [Member Type Choice Discovered ](#member-type-choice-discovered) below for details.

`dns` - (Optional) DNS. See [Member Type Choice Dns ](#member-type-choice-dns) below for details.

###### One of the arguments from this list "segment_network, site_local_inside_network, site_local_network" can be set

`segment_network` - (Optional) x-displayName: "Segment (Global VRF)". See [ref](#ref) below for details.

`site_local_inside_network` - (Optional) x-displayName: "Site Local Inside (Local VRF)" (`Bool`).

`site_local_network` - (Optional) x-displayName: "Site Local Outside (Local VRF)" (`Bool`).

`virtual_site_name` - (Optional) x-example: "shared/ce-site". See [ref](#ref) below for details.

### Where Type Choice Virtual Network

x-displayName: "Segment".

###### One of the arguments from this list "custom_ip, dns" can be set

`custom_ip` - (Optional) IP Address. See [Member Type Choice Custom Ip ](#member-type-choice-custom-ip) below for details.

`dns` - (Optional) DNS. See [Member Type Choice Dns ](#member-type-choice-dns) below for details.

`segment_network` - (Optional) x-displayName: "Segment". See [ref](#ref) below for details.

Attribute Reference
-------------------

*   `id` - This is the id of the configured vs_pool.
