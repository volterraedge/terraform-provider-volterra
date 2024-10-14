---

page_title: "Volterra: volterra_tgw_info"
description: "The volterra_tgw_info runs SetTGWInfo custom api for a aws_tgw_site site"

---

Resource volterra_tgw_info
==========================

Configure TGW Information like transit gateway id, vpc id and subnet ids on aws_tgw_site object

~> **Note:** Please refer to [SetTGWInfo docs](https://docs.cloud.f5.com/docs/api/views-aws-tgw-site#operation/ves.io.schema.views.aws_tgw_site.CustomAPI.SetTGWInfo) to learn more ~> **Note:** This resource works well as part of Volterra's aws_tgw_site site bring up terraform script. Please do not use this on already provisioned aws_tgw_site site.

Example Usage
-------------

---

```hcl
resource "volterra_tgw_info" "example" {
  name            = "aws-tgw-site1"
  namespace       = "system"
  tgw_id = "tgw-0d9bf32a5171e5a49"
  vpc_id = "vpc-02d0d70860b330393"
  subnet_ids {
			outside_subnet_id = "subnet-12345678"
			inside_subnet_id = "subnet-12345678"
			workload_subnet_id = "subnet-12345678"
			az = "us-east-2a"
	}
}

```

### Argument Reference

---

-	`name` - (Required) The value of aws_tgw_site name has to follow DNS-1035 format. (`String`).

-	`namespace` - (Required) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

-	`tgw_id` - (Required) List of vpc ip prefixes

-	`vpc_id` - (Required) List of vpc ip prefixes

-	`subnet_ids` - (Required) List of subnetIds associated with the aws_tgw_site

	-	`az` - (Required) Name of the AZ in which the subnet exists (`String`).
	-	`outside_subnet_id` - (Required) Outside Subnet Id.
	-	`inside_subnet_id` - (Required) Inside Subnet Id.
	-	`workload_subnet_id` - (Required) Workload Subnet Id.
