---

page_title: "Volterra: volterra_cloud_site_labels"

description: "The volterra_cloud_site_labels helps to update the site labels for cloud sites"
---------------------------------------------------------------------------------------------

Resource volterra_cloud_site_labels
===================================

volterra_cloud_site_labels resource updates the labels for the cloud site also known as view sites

~> **Note:** Please add ignore_lifecycle hooks for your cloud sites to ignore updates on labels and control the update of site label through this resource. Below is a sample snippet.

```
lifecycle {
	ignore_changes = [labels]
}
```

Example Usage
-------------

```hcl
resource "volterra_azure_vnet_site" "az-site" {
	name = "az-site-1"
	namespace = "system"
	azure_region = "eastus"
	resource_group = "az-site-1"
	azure_cred {
	name = "az-site-1"
    namespace = "system"
  }
	vnet {
	  new_vnet {
	  	name = "az-site-1"
	  	primary_ipv4 = "192.168.0.0/22"
	  }
	}
	disk_size = 80
	machine_type = "Standard_D4_v2"

	ingress_egress_gw {
	  azure_certified_hw = "azure-byol-multi-nic-voltmesh"
	  az_nodes {
	  	azure_az = "1"
	  	disk_size = "100"
	  	inside_subnet {
	  		subnet_param {
	  			ipv4 = "192.168.0.0/25"
	  		}
	  	}
	  	outside_subnet {
	  		subnet_param {
			  	ipv4 = "192.168.1.0/25"
		  	}
		  }
	  }
	  no_global_network = true
	  no_inside_static_routes = true
	  no_network_policy = true
	  no_outside_static_routes = true
	  no_forward_proxy = true
	}
	logs_streaming_disabled = true
	ssh_key = "ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAD5sRjpKVBvf5afxhysXd4GyvEFaiDOnPhKQcK8SHNUxkGkjhRV6xMFpBBApNctQ73yaHweV//OhBHurwzUodKOWAEyH+ay0V2BAOpx2aiQHxiMh7b0CGYVxv4lRZ4IPZ1Da9Siz1Sz19RYBjVM7v6Dvo2UlYftUyauKPIDPnd19iN10g=="
	nodes_per_az = 2
	coordinates {
	  latitude = 37.404989
	  longitude = -121.942300
	}
	lifecycle {
	  ignore_changes = [labels]
	}
}
resource "volterra_cloud_site_labels" "labels" {
			name = volterra_azure_vnet_site.az-site.name
			site_type = "azure_vnet_site"
			labels = {
				key1 = "value1"
				key2 = "value2"
			}
			ignore_on_delete = true
		}

```

### Argument Reference

---

`name` - (Required) Name of the cloud site (`String`).

`site_type` - (Required) Cloud site type, valid values are `aws_vpc_site`, `aws_tgw_site`, `azure_vnet_site`, `gcp_vpc_site`.

`labels` - (Required) labels which needs to be updated on the cloud site (`String`).

`ignore_on_delete` - (Optional) ignore removing the labels which are set during the apply(`Bool`). Default value is `true`
