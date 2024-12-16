---

page_title: "Volterra: volterra_site_state" 

description: "The volterra_site_state changes state of the site"

---

Resource volterra_site_state
============================

=============================

volterra_site_state is used to change the state of the site.

~> **Note:** Please refer to [Set Site State API docs](https://docs.cloud.f5.com/docs/api/site#operation/ves.io.schema.site.CustomStateAPI.SetState) to learn more

Example Usage
-------------

---

```hcl
resource "volterra_site_state" "example" {
  name  = "site1"
  state = "DECOMMISSIONING"
  when  = "delete"
}

```

### Argument Reference

---

`name` - (Required) The value of Site name has to follow DNS-1035 format. (`String`).

`state` - (Required) State of Site defines in which operational state site itself is. Site is online and operational. Site is in provisioning state. For instance during site deployment or switching to different connected Regional Edge. Site is in process of upgrade. It transition to ONLINE or FAILED state. Site is in Standby before goes to ONLINE. This is mainly for Regional Edge sites to do their verification before they go to ONLINE state. Site is in failed state. It failed during provisioning or upgrade phase. Site Status Objects contain more details. Reregistration was requested Reregistration is in progress and maurice is waiting for nodes Site deletion is in progress Site is waiting for registration.

Valid values are `REREGISTRATION`, `DECOMMISSIONING` (`String`).

`when` - (Required) Indicates when the action must be performed. Valid values are `create` or `delete` (`String`).

`wait_time` - (Optional) Number of seconds to wait before it retries (`Int`).

`retry` - (Optional) Number of retries to set the site state, before it fails (`Int`).

Attribute Reference
-------------------

---

-	`id` - This is the id of Volterra site object.
