---

page_title: "Volterra: k8s_cluster_role"

description: "The k8s_cluster_role allows CRUD of K8s Cluster Role resource on Volterra SaaS"
---------------------------------------------------------------------------------------------

Resource volterra_k8s_cluster_role
==================================

The K8s Cluster Role allows CRUD of K8s Cluster Role resource on Volterra SaaS

~> **Note:** Please refer to [K8s Cluster Role API docs](https://docs.cloud.f5.com/docs-v2/api/k8s-cluster-role) to learn more

Example Usage
-------------

```hcl
resource "volterra_k8s_cluster_role" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "k8s_cluster_role_selector policy_rule_list yaml" must be set

  policy_rule_list {
    policy_rule {
      // One of the arguments from this list "non_resource_url_list resource_list" must be set

      non_resource_url_list {
        urls = ["value"]

        verbs = ["get"]
      }
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

###### One of the arguments from this list "k8s_cluster_role_selector, policy_rule_list, yaml" must be set

`k8s_cluster_role_selector` - (Optional) This role is aggregation of all rules in roles selected by the label expression. See [Rule Choice K8s Cluster Role Selector ](#rule-choice-k8s-cluster-role-selector) below for details.

`policy_rule_list` - (Optional) Policy in terms of rule list.. See [Rule Choice Policy Rule List ](#rule-choice-policy-rule-list) below for details.

`yaml` - (Optional) K8s YAML for ClusterRole (`String`).

### Policy Rule List Policy Rule

List of rules for role permissions.

###### One of the arguments from this list "non_resource_url_list, resource_list" must be set

`non_resource_url_list` - (Optional) permissions for URL(s) that do not represent K8s resource. See [Resource Choice Non Resource Url List ](#resource-choice-non-resource-url-list) below for details.

`resource_list` - (Optional) List of resources in terms of api groups/resource types/resource instances and verbs allowed. See [Resource Choice Resource List ](#resource-choice-resource-list) below for details.

### Resource Choice Non Resource Url List

permissions for URL(s) that do not represent K8s resource.

`urls` - (Required) allowed URL(s) that do not represent any K8s resource. URL can be suffix or regex. (`String`).

`verbs` - (Required) Allowed list of verbs(operations) on resources. Use VerbAll for all operations (`String`).

### Resource Choice Resource List

List of resources in terms of api groups/resource types/resource instances and verbs allowed.

`api_groups` - (Required) Allowed list of API group that contains resources, all resources of a given api group (`String`).

`resource_instances` - (Optional) Allowed list of resource instances within the resource types. (`String`).

`resource_types` - (Required) Allowed list of resource types within the api groups. (`String`).

`verbs` - (Required) Allowed list of verbs(operations) on resources. Use * for all operations (`String`).

### Rule Choice K8s Cluster Role Selector

This role is aggregation of all rules in roles selected by the label expression.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Rule Choice Policy Rule List

Policy in terms of rule list..

`policy_rule` - (Required) List of rules for role permissions. See [Policy Rule List Policy Rule ](#policy-rule-list-policy-rule) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured k8s_cluster_role.
