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
    rate  = 30.0
    burst = 15
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

* `limiter` -  (Optional) Global API rate limiting configuration block. When not specified, requests are sent without rate limiting. Controls request timing across all API operations to prevent HTTP 429 errors. See [Limiter](#limiter) below for details.

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

The `limiter` block helps avoid HTTP 429 rate limit errors by controlling the timing of requests sent to the F5 Distributed Cloud API.

`rate` - (Required) Request rate in **requests per minute**. Set this to match your backend rate quota. Controls how quickly tokens refill. Must be greater than 0 (`Float64`).

`burst` - (Required) Maximum tokens that can accumulate. Determines how many requests can be sent in quick succession. We recommend calculating this value using the formula below. Minimum: 1 (`Int32`).

**How limiter works:** Requests wait for available tokens before being sent. Tokens refill at the configured rate, with a burst capacity for handling temporary spikes.

**Important consideration:** The provider uses a global rate limiter shared across all API operations, while the backend enforces separate limits per endpoint. During bulk operations (e.g., deleting 50+ resources), Terraform's parallelism (default 10, configurable via `-parallelism` flag) can cause many requests to target the same backend endpoint, potentially exceeding that specific endpoint's limit.

**Recommended Configuration:**

To avoid backend endpoint errors during bulk operations, calculate burst as:

```
burst = max(1, backendBurst / 6)
```

Example: If your backend quota is `rate=120/min, burst=30`, use `burst=5` in the provider.

**Why this formula?** The provider has 6 API operation types sharing one limiter. This calculation provides a safety margin accounting for Terraform's parallelism (default 10, configurable higher) and token refill during operations, keeping requests within backend endpoint limits in typical scenarios.

**Note:** Advanced users may experiment with higher burst values. If you encounter HTTP 429 errors during bulk operations, reduce burst using the formula above.

---

**Configuration Examples:**
- Backend `rate=120, burst=30`: Use `rate = 120.0, burst = 5`
- Backend `rate=60, burst=18`: Use `rate = 60.0, burst = 3`
- Backend `rate=300, burst=60`: Use `rate = 300.0, burst = 10`