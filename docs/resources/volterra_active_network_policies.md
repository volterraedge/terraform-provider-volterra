---

page_title: "Volterra: active_network_policies"

description: "The active_network_policies allows setting of Active Network Policies for a namespace on Volterra SaaS"
---------------------------------------------------------------------------------------------------------------------

Resource volterra_active_network_policies
=========================================

The Active Network Policies allows you to set a list of Active Network Policies on a certain namespace

~> **Note:** Please refer to [Active Network Policies API docs](https://docs.cloud.f5.com/docs/api/namespace#operation/ves.io.schema.namespace.NamespaceCustomAPI.SetActiveNetworkPolicies) to learn more

Example Usage
-------------

```hcl
resource "volterra_namespace" "system" {
	name = "staging"
}
resource "volterra_network_policy" "allow_all" {
	name        = "allow-all-test"
	namespace   = volterra_namespace.system.name
	endpoint {
		any = true
	}
	rules {
		egress_rules {
			action = "ALLOW"
			metadata {
				name = "allow-egress-all"
			}
			any = true
			all_traffic = true
		}
		ingress_rules {
			action = "ALLOW"
			metadata {
				name = "allow-ingress-all"
			}
			any = true
			all_traffic = true
		}
	}

}

resource "volterra_active_network_policies" "active_staging" {
	depends_on = [volterra_network_policy.allow_all]
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

`policies` - (Required) An ordered list of references to network_policy objects.. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of instance of active network policies.
