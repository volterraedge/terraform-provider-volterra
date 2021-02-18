---

page_title: "Volterra: active_service_policies"

description: "The active_service_policies allows setting of Active Service Policies for a namespace on Volterra SaaS"
---------------------------------------------------------------------------------------------------------------------

Resource volterra_active_service_policies
=========================================

The active service policies allows you to set a list of active service policies on a certain namespace

~> **Note:** Please refer to [Active Service Policies API docs](https://volterra.io/docs/api/namespace#operation/ves.io.schema.namespace.NamespaceCustomAPI.SetActiveNetworkPolicies) to learn more

Example Usage
-------------

```hcl
resource "volterra_namespace" "system" {
	name = "staging"
}
resource "volterra_service_policy" "allow_all" {
	name        = "allow-all-test"
	namespace   = volterra_namespace.system.name
	algo        = "FIRST_MATCH"
	any_server  = true

	simple_rules {
		action = "ALLOW"
		waf_action {
			none = true
			waf_in_monitoring_mode = false
			waf_skip_processing    = false
		}
	}

	allow_all_requests = true
	deny_all_requests  = false
}

resource "volterra_active_service_policies" "active_staging" {
	depends_on = [volterra_service_policy.allow_all]
	namespace = volterra_namespace.system.name
	policies {
		name = "allow-all-test"
		namespace = volterra_namespace.system.name
		tenant = volterra_namespace.system.tenant_name
	}
}

```

Argument Reference
------------------

`namespace` - (Required) (Forces New) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`). If value updated, it will delete and create a new resource (Force New).

`policies` - (Required) An ordered list of references to service_policy objects.. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of instance of active service policies.
