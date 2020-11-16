---

page_title: "Volterra: service_policy_set"

description: "The service_policy_set allows CRUD of Service Policy Set resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_service_policy_set
====================================

The Service Policy Set allows CRUD of Service Policy Set resource on Volterra SaaS

~> **Note:** Please refer to [Service Policy Set API docs](https://volterra.io/docs/api/service-policy-set) to learn more

Example Usage
-------------

```hcl
resource "volterra_service_policy_set" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

`default_action` - (Optional) policy sets for the time being. (`String`).

`deny_info` - (Optional) Detailed information including HTTP response code and error message to be sent when the policy set default action is used and that action is DENY.. See [Deny Info ](#deny-info) below for details.

`policies` - (Optional) An ordered list of references to service_policy objects.. See [ref](#ref) below for details.

`scope` - (Optional) for system created policy sets that are specific to a virtual host. (`String`).

`type` - (Optional) rate limiting policy sets for the time being. (`String`).

### Deny Info

Detailed information including HTTP response code and error message to be sent when the policy set default action is used and that action is DENY..

`error_message` - (Optional) in the body but does not constitute the entire body. (`String`).

`response_code` - (Optional) The HTTP status code to use in the response. The default code is Forbidden (403). (`Int`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of the configured service_policy_set.
