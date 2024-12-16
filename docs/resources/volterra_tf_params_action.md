---

page_title: "Volterra: volterra_tf_params_action" 

description: "The volterra_tf_params_action runs action api on view sites"

---

Resource volterra_tf_params_action
==================================

====================================

volterra_tf_params_action perform terraform actions for a given view. Supported actions are apply and plan. By default destroy action is performed during delete.

~> **Note:** Please refer to [Run Terraform Action API docs](https://docs.cloud.f5.com/docs/api/views-terraform-parameters#operation/ves.io.schema.views.terraform_parameters.CustomActionAPI.Run) to learn more

Example Usage
-------------

---

```hcl
resource "volterra_tf_params_action" "example" {
  site_name       = "site1"
  site_kind       = "aws_vpc_site
  action          = "plan"
  wait_for_action = true
}

```

### Argument Reference

---

`site_name` - (Required) The value of Site name has to follow DNS-1035 format (`String`).

`site_kind` - (Required) Kind of view of which terraform parameters are requested. Valid values are `aws_vpc_site`, `aws_tgw_site`, `gcp_vpc_site`, `azure_vnet_site` (`String`).

`action` - (Required) Terraform action to be performed for a given view. Valid values are `plan`, `apply`

APPLY: Apply action used to apply the changes required to reach the desired state of the configuration PLAN: Plan action is a convenient way to check whether the execution plan for a set of changes matches your expectations without making any changes to real resources or to the state

`ignore_on_update` - (Optional) Ignore action to perform during update, by default value is true (`Bool`).

`wait_for_action` - (Optional) Wait for terraform action job to complete, by default value is true (`Bool`).

Attribute Reference
-------------------

---

-	`tf_output` - Terraform output string, when action is `apply`
