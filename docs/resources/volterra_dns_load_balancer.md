











---
page_title: "Volterra: dns_load_balancer"
description: "The dns_load_balancer allows CRUD of Dns Load Balancer  resource on Volterra SaaS"
---
# Resource volterra_dns_load_balancer

The Dns Load Balancer  allows CRUD of Dns Load Balancer  resource on Volterra SaaS

~> **Note:** Please refer to [Dns Load Balancer  API docs](https://docs.cloud.f5.com/docs-v2/api/dns-load-balancer) to learn more

## Example Usage

```hcl
resource "volterra_dns_load_balancer" "example" {
  name        = "acmecorp-web"
  namespace   = "staging"
  record_type = ["record_type"]

  rule_list {
    rules {
      // One of the arguments from this list "pool nxdomain" must be set

      pool {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }

      // One of the arguments from this list "asn_matcher ip_prefix_list ip_prefix_set geo_location_label_selector geo_location_set asn_list" must be set

      ip_prefix_list {
        invert_match = true

        ip_prefixes = ["192.168.20.0/24"]
      }
      score = "50"
    }
  }
}

```

## Argument Reference

### Metadata Argument Reference
`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).


`description` - (Optional) Human readable description for the object (`String`).


`disable` - (Optional) A value of true will administratively disable the object (`Bool`).


`labels` - (Optional) by selector expression (`String`).


`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).


`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).



### Spec Argument Reference

`fallback_pool` - (Optional) Fallback Pool to be used for load balancing if none of the Load Balancing rules match. See [ref](#ref) below for details.

`record_type` - (Required) x-required (`String`).



`response_cache` - (Optional) Response Cache Parameters. See [Response Cache ](#response-cache) below for details.




		




		




		








`rule_list` - (Required) Load Balancing Rules. See [Rule List ](#rule-list) below for details.


		




		







		





		





		






		






		








### Response Cache 

 Response Cache Parameters.



###### One of the arguments from this list "disable, default_response_cache_parameters, response_cache_parameters" must be set

`default_response_cache_parameters` - (Optional) Default Parameters for caching the DNS responses (`Bool`).


`disable` - (Optional) When Response Cache is disabled, responses will be computed for each request (`Bool`).


`response_cache_parameters` - (Optional) Customize the parameters for Response cache. See [Response Cache Parameters Choice Response Cache Parameters ](#response-cache-parameters-choice-response-cache-parameters) below for details.




### Rule List 

 Load Balancing Rules.

`rules` - (Required) Rules to perform load balancing. See [Rule List Rules ](#rule-list-rules) below for details.



### Action Choice Nxdomain 

 Do not perform any load-balancing. Instead return NXDOMAIN.



### Client Choice Asn List 

 The rule evaluates to true if the origin ASN is present in the ASN list..

`as_numbers` - (Required) An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer. (`Int`).



### Client Choice Asn Matcher 

 The rule evaluates to true if the origin ASN is present in one of the BGP ASN Set objects..

`asn_sets` - (Required) A list of references to bgp_asn_set objects.. See [ref](#ref) below for details.



### Client Choice Geo Location Label Selector 

 with the translated geo locations derived from incoming EDNS0 client-subnet in the DNS request..

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).



### Client Choice Ip Prefix List 

 IP Prefix list..

`invert_match` - (Optional) Invert the match result. (`Bool`).

`ip_prefixes` - (Optional) List of IPv4 prefix strings. (`String`).



### Client Choice Ip Prefix Set 

 IP Prefix set..

`invert_matcher` - (Optional) Invert the match result. (`Bool`).

`prefix_sets` - (Required) A list of references to ip_prefix_set objects.. See [ref](#ref) below for details.



### Ref 


Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).



### Response Cache Parameters Choice Default Response Cache Parameters 

 Default Parameters for caching the DNS responses.



### Response Cache Parameters Choice Disable 

 When Response Cache is disabled, responses will be computed for each request.



### Response Cache Parameters Choice Response Cache Parameters 

 Customize the parameters for Response cache.

`cache_cidr_ipv4` - (Optional) Length of CIDR masks used to group IPv4 clients (`Int`).

`cache_cidr_ipv6` - (Optional) Length of CIDR masks used to group IPv6 clients (`Int`).

`cache_ttl` - (Optional) TTL for response cache (`Int`).



### Rule List Rules 

 Rules to perform load balancing.



###### One of the arguments from this list "pool, nxdomain" must be set

`nxdomain` - (Optional) Do not perform any load-balancing. Instead return NXDOMAIN (`Bool`).(Deprecated)


`pool` - (Optional) Use this pool for the Load Balancing.. See [ref](#ref) below for details.




###### One of the arguments from this list "asn_list, asn_matcher, ip_prefix_list, ip_prefix_set, geo_location_label_selector, geo_location_set" must be set

`asn_list` - (Optional) The rule evaluates to true if the origin ASN is present in the ASN list.. See [Client Choice Asn List ](#client-choice-asn-list) below for details.


`asn_matcher` - (Optional) The rule evaluates to true if the origin ASN is present in one of the BGP ASN Set objects.. See [Client Choice Asn Matcher ](#client-choice-asn-matcher) below for details.


`geo_location_label_selector` - (Optional) with the translated geo locations derived from incoming EDNS0 client-subnet in the DNS request.. See [Client Choice Geo Location Label Selector ](#client-choice-geo-location-label-selector) below for details.


`geo_location_set` - (Optional) with the translated geo locations derived from incoming EDNS0 client-subnet in the DNS request.. See [ref](#ref) below for details.


`ip_prefix_list` - (Optional) IP Prefix list.. See [Client Choice Ip Prefix List ](#client-choice-ip-prefix-list) below for details.


`ip_prefix_set` - (Optional) IP Prefix set.. See [Client Choice Ip Prefix Set ](#client-choice-ip-prefix-set) below for details.


`score` - (Optional) When multiple load balancing rules match a query, the one with the highest score is chosen (`Int`).



## Attribute Reference

* `id` - This is the id of the configured dns_load_balancer.

