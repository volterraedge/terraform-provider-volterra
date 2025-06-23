---

page_title: "Volterra: volterra_tgw_vpc_ip_prefixes"

description: "The volterra_tgw_vpc_ip_prefixes runs set_vpc_ip_prefixes api on aws_tgw_site"

---

Resource volterra_tgw_vpc_ip_prefixes
=====================================

====================================

Configure VPC IP prefix set for aws_tgw_site object

~> **Note:** Please refer to [Set vpc IP prefix docs](https://docs.cloud.f5.com/docs/api/views-aws-tgw-site#operation/ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPCIpPrefixes) to learn more ~> **Note:** This resource works well as part of Volterra's aws_tgw_site site bring up terraform script. Please do not use this on already provisioned aws_tgw_site site.

Example Usage
-------------

---

```hcl
resource "volterra_tgw_vpc_ip_prefixes" "example" {
  name            = "aws-tgw-site1"
  namespace       = "system"
  vpc_ip_prefixes  {
    name = "vpc-12345678"
    value = ["10.0.0.0/16"]
  }
}

```

### Argument Reference

---

-	`name` - (Required) The value of aws_tgw_site name has to follow DNS-1035 format. (`String`).

-	`namespace` - (Required) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

-	`vpc_ip_prefixes` - (Required) List of vpc ip prefixes

	-	`name` - (Required) Name of vpc (`String`\)
	-	`values` - (Required) List of vpc cidrs (`String`\)
