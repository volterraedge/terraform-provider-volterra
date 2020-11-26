---

page_title: "Volterra: volterra_registration_approval"

description: "The volterra_registration_approval helps approve registration request"
---

# Resource volterra_registration_approval
=============================

volterra_registration_approval approves pending registration request.

~> **Note:** Please refer to [Registration api docs](https://volterra.io/docs/api/registration#operation/ves.io.schema.registration.CustomAPI.RegistrationApprove) to learn more

## Example Usage
----------------

```hcl
resource "volterra_registration_approval" "example" {
  cluster_name  = "site1"
  cluster_size  = 3
  retry = 5
  wait_time = 60
  latitude = 37.3382
  longitude = 121.8863
}

```

### Argument Reference
----------------------

`cluster_name` - (Required) Name of the site/cluster for which you are approving the request (`String`).

`cluster_size` - (Required) Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time, cluster won't start until ALL nodes are ADMITTED This value can't be changed after installation. It does not interact with auto-scaling as only pool nodes are scaled (`Int`).

`latitude` - (Optional) Geographic latitude location of this site (`Float`).

`longitude` - (Optional) Geographic longitude location of this site (`Float`).

`connected_region` - (Optional) closest REs to which to connect in selected region. If there is more than one RE with the label then select 2 of them. Otherwise default "geo-ip" database is consulted to find closest RE(s) to  public ip with which node registration request was received.

`hostname` - (Optional) Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname (`String`).

`wait_time` - (Optional) Number of seconds to wait before it retries (`Int`).

`retry` - (Optional) Number of retries to get the volterra registration config, before it fails (`Int`).


## Attribute Reference
----------------------

-	`id` - This is the id of Volterra registration object.

- `registration_name` - (Optional) Name of the registration object
