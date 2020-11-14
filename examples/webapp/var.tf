
variable "api_cert" {
	default = ""
}
variable "api_key" {
	default = ""
}
variable "api_ca_cert" {
	default = ""
}

variable "api_url" {}

variable "webapp" {}

variable "https_offload" {
	default = "false"
}

variable "site_name" {
	default = "ny8-nyc"
}

variable "name" {}

variable "api_p12_file" {}

variable "diable_api_cred" {
  default = false
}

