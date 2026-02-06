<a href="https://terraform.io">
    <img src=".github/tf.png" alt="Terraform logo" title="Terraform" align="left" height="50" />
</a>

# Terraform Provider for F5XC (formerly Volterra)

The F5XC Terraform Provider allows managing resources within the F5 XC Distributed Cloud Platform.

We recommend using the latest version of Terraform Core ([the latest version can be found here](https://developer.hashicorp.com/terraform/install)).

* [Terraform Website](https://www.terraform.io)
* [F5XC Provider Documentation](https://registry.terraform.io/providers/volterraedge/volterra/latest/docs)
* [F5XC Provider Usage Examples](https://github.com/volterraedge/terraform-provider-volterra/tree/main/examples)

## Usage Example

```hcl
# 1. Specify the version of the F5XC Provider to use
terraform {
  required_providers {
    volterra = {
      source  = "localhost/volterraedge/volterra"
      version = ">=0.11.47"
    }
  }
}

# 2. Configure the F5XC Provider
provider "volterra" {
  api_p12_file     = "/path/to/api_credential.p12"
  url              = "https://<tenant_name>.console.ves.volterra.io/api"
  limiter {
    rate  = 30.0
    burst = 15
  }
}
```

