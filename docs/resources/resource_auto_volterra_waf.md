---

page_title: "Volterra: waf"

description: "The waf allows CRUD of Waf resource on Volterra SaaS"
-------------------------------------------------------------------

Resource volterra_waf
=====================

The Waf allows CRUD of Waf resource on Volterra SaaS

~> **Note:** Please refer to [Waf API docs](https://volterra.io/docs/api/waf) to learn more

Example Usage
-------------

```hcl
resource "volterra_waf" "example" {
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

`app_profile` - (Optional) Details of the apps which are going through this WAF instance.. See [App Profile ](#app-profile) below for details.

`disabled_detection_tags` - (Optional) DisabledDetectionTags will allow user to selectively disable the Security Rules associated with Tags. See [Disabled Detection Tags ](#disabled-detection-tags) below for details.

`mode` - (Optional) Mode of this WAF instance , will be either in blocked or alert mode. (`String`).

### App Profile

Details of the apps which are going through this WAF instance..

`cms` - (Optional) The content management system used to build the target application. (`List of Strings`).

`language` - (Optional) The languages used to build the target application. (`List of Strings`).

`webserver` - (Optional) The Web Server used to build the target application. (`List of Strings`).

### Disabled Detection Tags

DisabledDetectionTags will allow user to selectively disable the Security Rules associated with Tags.

`detection_tag_type` - (Optional) List of detection tags to be disabled (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured waf.
