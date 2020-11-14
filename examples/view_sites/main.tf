provider "volterra" {
  api_p12_file = "${var.api_p12_file}"
  api_cert     = "${var.api_p12_file != "" ? "" : var.api_cert}"
  api_key      = "${var.api_p12_file != "" ? "" : var.api_key}"
  api_ca_cert  = "${var.api_ca_cert}"
  url          = "${var.api_url}"
}

resource "volterra_cloud_credentials" "aws_cred" {
  name = var.name
  namespace = "system"
  aws_secret_key {
	  access_key = var.aws_access_key
	  secret_key {
	  	clear_secret_info {
	  		url = var.aws_secret_key_clear
	  	}
	  }
  }
}
resource "volterra_aws_vpc_site" "aws_test" {
  name = var.name
  namespace = "system"
  aws_region = "us-west-2"
  aws_cred {
	  name = volterra_cloud_credentials.aws_cred.name
	  namespace = "system"
  }
  vpc {
	  new_vpc {
	  	name_tag = "tf-test-vpc"
	  	primary_ipv4 = "192.168.0.0/22"
	  }
  }
  disk_size = 80
  instance_type = "t3.xlarge"

  ingress_egress_gw {
	  aws_certified_hw = "aws-byol-multi-nic-voltmesh"
	  az_nodes {
	  	aws_az_name = "us-west-2a"
	  	disk_size = "100"
	  	inside_subnet {
	  		subnet_param {
	  			ipv4 = "192.168.0.0/24"
	  		}
	  	}
	  	outside_subnet {
	  		subnet_param {
	  			ipv4 = "192.168.1.0/24"
	  		}
	  	}
	  }
	  no_global_network = true
	  no_inside_static_routes = true
	  no_network_policy = true
	  no_outside_static_routes = true
	  no_forward_proxy_policy = true
  }
  ssh_key = "ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAD5sRjpKVBvf5afxhysXd4GyvEFaiDOnPhKQcK8SHNUxkGkjhRV6xMFpBBApNctQ73yaHweV//OhBHurwzUodKOWAEyH+ay0V2BAOpx2aiQHxiMh7b0CGYVxv4lRZ4IPZ1Da9Siz1Sz19RYBjVM7v6Dvo2UlYftUyauKPIDPnd19iN10g=="

}
resource "volterra_tf_params_action" "action_test" {
	site_name = volterra_aws_vpc_site.aws_test.name
	site_kind = "aws_vpc_site"
	action = "plan"
	wait_for_action = true
}