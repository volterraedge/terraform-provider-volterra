---

page_title: "Volterra: fast_acl_for_internet_vips"

description: "The fast_acl_for_internet_vips activates the passed list of FastACLs for Internet VIPs"
-----------------------------------------------------------------------------------------------------

Resource volterra_fast_acl_for_internet_vips
============================================

The fast_acl for internet vips resource activates the passed list of FastACLs for Internet VIPs

~> **Note:** Please refer to [Fast ACL For Internet VIPs API docs](https://docs.cloud.f5.com/docs/api/namespace#operation/ves.io.schema.namespace.NamespaceCustomAPI.SetFastACLsForInternetVIPs) to learn more

Example Usage
-------------

```hcl
data "volterra_namespace" "system" {
	name = "system"
}
resource "volterra_policer" "this" {
	name = "test-policer"
	namespace = "system"
	policer_type = "POLICER_SINGLE_RATE_TWO_COLOR"
	committed_information_rate = 10000
	burst_size = 5000
	policer_mode = "POLICER_MODE_SHARED"
}
resource "volterra_protocol_policer" "this" {
	name = "test-proto-policer"
	namespace = "system"
	protocol_policer {
		policer {
			name = volterra_policer.this.name
			namespace = "system"
		}
		protocol {
			tcp {
				flags = ["ALL_TCP_FLAGS"]
			}
		}
	}
}
resource "volterra_fast_acl" "this" {
	name        = facl-test"
	namespace   = "system"
	re_acl {
		all_public_vips = true
		fast_acl_rules {
			action {
				simple_action = "DENY"
			}
			metadata {
				name = "deny_certain_ip"
			}
			prefix {
				prefix = ["1.1.1.1/32"]
			}
		}
	}
}

resource "volterra_fast_acl_for_internet_vips" "this" {
	depends_on = [volterra_fast_acl.this]
	fast_acls {
		name = "%[1]s-test"
		tenant = data.volterra_namespace.system.tenant_name
	}
}

```

Argument Reference
------------------

`fast_acls` - (Required) An ordered list of references to fast_acl objects.. See [ref](#ref) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

Attribute Reference
-------------------

-	`id` - This is the id of instance of Fast ACL For Internet VIPs.
