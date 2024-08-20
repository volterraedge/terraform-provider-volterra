











---
page_title: "Volterra: virtual_site"
description: "The virtual_site allows CRUD of Virtual Site  resource on Volterra SaaS"
---
# Resource volterra_virtual_site

The Virtual Site  allows CRUD of Virtual Site  resource on Volterra SaaS

~> **Note:** Please refer to [Virtual Site  API docs](https://docs.cloud.f5.com/docs-v2/api/virtual-site) to learn more

## Example Usage

```hcl
resource "volterra_virtual_site" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  site_selector {
    expressions = ["region in (us-west1, us-west2),tier in (staging)"]
  }

  site_type = ["site_type"]
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

`site_selector` - (Required) Label selector for virtual-site. See [Site Selector ](#site-selector) below for details.



`site_type` - (Required) Site type for sites to be selected. A virtual site can only select sites of given type (CE or RE). Sites of different types cannot be selected by one virtual site. (`String`).



### Site Selector 

 Label selector for virtual-site.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).



## Attribute Reference

* `id` - This is the id of the configured virtual_site.

