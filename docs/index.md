# Volterra Provider

Volterra terraform provider can be used to configure Volterra Objects amd these objects represent desired state of the system. The desired state of the system could be configuring a http/tcp loadbalancer, vk8s cluster, service mesh, enabling api security etc.


## Authenticating to Volterra

Volterra terraform provider supports authenticating to Volterra using API Credential P12 file. Learn more on [how to generate api certificate](https://volterra.io/docs/how-to/user-mgmt/credentials).


## Example Usage

```hcl
provider "volterra" {
  api_p12_file     = "/path/to/api_credential.p12"
  url              = "https://<tenant_name>.console.ves.volterra.io/api"
}
```


## Argument Reference

* `api_p12_file` - API credential p12 file path. Either api p12 file or (`api_cert` and `api_key`) must be given. This can also be sourced from `VOLT_API_P12_FILE` env variable.

* `api_cert` - API certificate file path. `api_cert` and `api_key` combo must be given. This can also be sourced from `VOLT_API_CERT` env variable.

* `api_key` - API certificate private key file path. This can also be sourced from `VOLT_API_KEY` env variable.

* `url` - (Required) Tenant API url file path. This can also be sourced from `VOLT_API_URL` env variable.

* `timeout` - Volterra api call timeout, by default its 10 seconds. This can also be sourced from `VOLT_API_TIMEOUT` env variable.