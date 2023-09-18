---

page_title: "Volterra: active_alert_policies"

description: "The active_alert_policies allows setting of Active Alert Policies for a namespace on Volterra SaaS"
-----------------------------------------------------------------------------------------------------------------

Resource volterra_active_alert_policies
=======================================

The active alert policies allows you to set a list of active alert policies on a certain namespace

~> **Note:** Please refer to [Active Alert Policies API docs](https://docs.cloud.f5.com/docs/api/namespace#operation/ves.io.schema.namespace.NamespaceCustomAPI.SetActiveAlertPolicies) to learn more

Example Usage
-------------

```hcl
resource "volterra_namespace" "system" {
	name = "staging"
}
resource "volterra_alert_receiver" "email"{
	name = "email-alert"
	namespace   = volterra_namespace.system.name
	email {
		email = "support@ves.io"
	}
}
resource "volterra_alert_policy" "dev-policy" {
	name        = "dev-policy"
	namespace   = volterra_namespace.system.name
	receivers {
		name = volterra_alert_receiver.email.name
		namespace = volterra_namespace.system.name
	}
	routes {
		send = true
		any = true
	}
}

resource "volterra_active_alert_policies" "active_staging" {
	depends_on = [volterra_alert_policy.dev-policy]
	namespace = volterra_namespace.system.name
	policies {
		name = "dev-policy"
		namespace = volterra_namespace.system.name
		tenant = volterra_namespace.system.tenant_name
	}
}

```

Argument Reference
------------------

`namespace` - (Required) (Forces New) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`). If value updated, it will delete and create a new resource (Force New).

`policies` - (Required) An ordered list of references to alert_policy objects.. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

-	`id` - This is the id of instance of active alert policies.
