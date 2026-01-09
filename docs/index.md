# Volterra Provider

Volterra terraform provider can be used to configure Volterra Objects amd these objects represent desired state of the system. The desired state of the system could be configuring a http/tcp loadbalancer, vk8s cluster, service mesh, enabling api security etc.


## Authenticating to Volterra

Volterra terraform provider supports authenticating to Volterra using API Credential P12 file. Learn more on [how to generate api certificate](https://docs.cloud.f5.com/docs/how-to/user-mgmt/credentials).


## Example Usage

```hcl
provider "volterra" {
  api_p12_file     = "/path/to/api_credential.p12"
  url              = "https://<tenant_name>.console.ves.volterra.io/api"
  limiter {
    rate  = 10.0
    burst = 20
  }
}
```


## Argument Reference

* `api_p12_file` - API credential p12 file path (`String`). See [ref](#api-credential-options) below for details.

* `api_cert` - API certificate file path. `api_cert`, `api_key` duo must be given. This can also be sourced from `VOLT_API_CERT` env variable (`String`).

* `api_key` - API certificate private key file path. This can also be sourced from `VOLT_API_KEY` env variable (`String`).

* `api_ca_cert` - API CA certificate file path. This can also be sourced from `VOLT_API_CA_CERT` env variable (`String`).

* `url` - (Required) Tenant API url file path. This can also be sourced from `VOLT_API_URL` env variable (`String`).

* `timeout` - Volterra api call timeout, by default its 20 seconds. This can also be sourced from `VOLT_API_TIMEOUT` env variable (`String`).

* `limiter` -  (Optional) API rate limiting configuration using a token bucket algorithm. Controls how frequently API requests are allowed to happen by limiting the rate to a specified number of tokens per second with a maximum burst size. See [Limiter](#limiter) below for details.

### API Credential Options
- Provide the path to the API P12 file using the `api_p12_file` variable.
  - **Password:** Set the P12 file password via the `VES_P12_PASSWORD` environment variable.

**OR**

- Use `api_cert`, `api_key` variables for authentication.

**OR**

- If sourcing from environment variables:
  - `api_p12_file`: Use `VOLT_API_P12_FILE`.
  - **Password:** Set the P12 file password.

  **OR**

  - `VES_P12_CONTENT`: Base64-encoded P12 bundle content.
  - **Password:** Set the P12 file password via the `VES_P12_PASSWORD` environment variable.

### Limiter

API rate limiter configuration. Implements a token bucket algorithm that controls request frequency using the WaitN method, which blocks API requests until sufficient tokens are available. The bucket starts full and refills at the specified rate. In any sufficiently large time interval, the limiter restricts requests to the configured rate per second, with allowance for bursts up to the maximum burst size.

`rate` - (Required) The rate at which tokens are refilled in the bucket, measured in tokens per second. This determines the sustained request rate allowed over time. When tokens are not available, requests will wait until they can be obtained. For example, a rate of 10.0 allows 10 requests per second on average. (`Float64`).

`burst` - (Required) Burst returns the maximum burst size. Burst is the maximum number of tokens that can be consumed in a single call to Allow, Reserve, or Wait, so higher Burst values allow more events to happen at once. A zero Burst allows no events, unless limit == Inf. (`Int32`).