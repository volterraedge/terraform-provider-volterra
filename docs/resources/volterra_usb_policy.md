---

page_title: "Volterra: usb_policy"

description: "The usb_policy allows CRUD of Usb Policy resource on Volterra SaaS"
---------------------------------------------------------------------------------

Resource volterra_usb_policy
============================

The Usb Policy allows CRUD of Usb Policy resource on Volterra SaaS

~> **Note:** Please refer to [Usb Policy API docs](https://docs.cloud.f5.com/docs-v2/api/usb-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_usb_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  allowed_devices {
    b_device_class = "hub"

    b_device_protocol = "0002"

    b_device_sub_class = "hub"

    i_serial = "0000:00:14.0"

    id_product = "0x0002"

    id_vendor = "0x1d6b"
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

`allowed_devices` - (Required) List of allowed USB devices. See [Allowed Devices ](#allowed-devices) below for details.

### Allowed Devices

List of allowed USB devices.

`b_device_class` - (Optional) The class of this device (`String`).

`b_device_protocol` - (Optional) The protocol (within the sub-class) of this device (`String`).

`b_device_sub_class` - (Optional) The sub-class (within the class) of this device (`String`).

`i_serial` - (Optional) Index of Serial Number String Descriptor (`String`).

`id_product` - (Optional) Product ID (Assigned by Manufacturer) in hex (`String`).

`id_vendor` - (Optional) Vendor ID (Assigned by USB Org) in hex (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured usb_policy.
