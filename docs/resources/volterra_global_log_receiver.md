---

page_title: "Volterra: global_log_receiver"

description: "The global_log_receiver allows CRUD of Global Log Receiver resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_global_log_receiver
=====================================

The Global Log Receiver allows CRUD of Global Log Receiver resource on Volterra SaaS

~> **Note:** Please refer to [Global Log Receiver API docs](https://volterra.io/docs/api/global-log-receiver) to learn more

Example Usage
-------------

```hcl
resource "volterra_global_log_receiver" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "ns_current ns_all ns_list ns_system" must be set
  ns_current = true

  // One of the arguments from this list "request_logs security_events audit_logs" must be set
  request_logs = true

  // One of the arguments from this list "elastic_receiver azure_event_hubs_receiver sumo_logic_receiver http_receiver datadog_receiver azure_receiver aws_cloud_watch_receiver kafka_receiver new_relic_receiver qradar_receiver s3_receiver splunk_receiver" must be set

  datadog_receiver {
    batch {
      // One of the arguments from this list "max_bytes_disabled max_bytes" must be set
      max_bytes_disabled = true

      // One of the arguments from this list "max_events_disabled max_events" must be set
      max_events_disabled = true

      // One of the arguments from this list "timeout_seconds_default timeout_seconds" must be set
      timeout_seconds_default = true
    }

    compression {
      // One of the arguments from this list "compression_none compression_gzip" must be set
      compression_none = true
    }

    datadog_api_key {
      blindfold_secret_info_internal {
        decryption_provider = "value"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "value"
      }

      secret_encoding_type = "secret_encoding_type"

      // One of the arguments from this list "clear_secret_info wingman_secret_info blindfold_secret_info vault_secret_info" must be set

      clear_secret_info {
        provider = "box-provider"
        url      = "string:///U2VjcmV0SW5mb3JtYXRpb24="
      }
    }

    // One of the arguments from this list "endpoint site" must be set
    site = "datadoghq.com"

    // One of the arguments from this list "no_tls use_tls" must be set
    no_tls = true
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

`ns_all` - (Optional) x-displayName: "Select logs from all namespaces" (bool).

`ns_current` - (Optional) x-displayName: "Select logs from current namespace" (bool).

`ns_list` - (Optional) x-displayName: "Select logs in specific namespaces". See [Ns List ](#ns-list) below for details.

`ns_system` - (Optional) x-displayName: "Select logs from System namespace" (bool).

`audit_logs` - (Optional) Send Audit Logs (corresponding to Public Audit and Authentication) (bool).

`request_logs` - (Optional) Send Request Logs (corresponding to Load Balancer access logs) (bool).

`security_events` - (Optional) Send Security Events (corresponding to e.g. WAF blocked events or malicious requests) (bool).

`aws_cloud_watch_receiver` - (Optional) Send logs to AWS Cloudwatch. See [Aws Cloud Watch Receiver ](#aws-cloud-watch-receiver) below for details.

`azure_event_hubs_receiver` - (Optional) Send logs to Azure Event Hubs. See [Azure Event Hubs Receiver ](#azure-event-hubs-receiver) below for details.

`azure_receiver` - (Optional) Send logs to Azure Blob Storage. See [Azure Receiver ](#azure-receiver) below for details.

`datadog_receiver` - (Optional) Send logs to a Datadog service. See [Datadog Receiver ](#datadog-receiver) below for details.

`elastic_receiver` - (Optional) Send logs to an Elasticsearch endpoint. See [Elastic Receiver ](#elastic-receiver) below for details.

`http_receiver` - (Optional) Send logs to a generic HTTP(s) server. See [Http Receiver ](#http-receiver) below for details.

`kafka_receiver` - (Optional) Send logs to a Kafka cluster. See [Kafka Receiver ](#kafka-receiver) below for details.

`new_relic_receiver` - (Optional) Send logs to NewRelic. See [New Relic Receiver ](#new-relic-receiver) below for details.

`qradar_receiver` - (Optional) Send logs to IBM QRadar. See [Qradar Receiver ](#qradar-receiver) below for details.

`s3_receiver` - (Optional) Send logs to an AWS S3 bucket. See [S3 Receiver ](#s3-receiver) below for details.

`splunk_receiver` - (Optional) Send logs to a Splunk HEC Logs service. See [Splunk Receiver ](#splunk-receiver) below for details.

`sumo_logic_receiver` - (Optional) Send logs to SumoLogic. See [Sumo Logic Receiver ](#sumo-logic-receiver) below for details.

### Api Key

A New Relic License Key.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Auth Basic

Basic Authentication specify that HTTP Basic Authentication should be used when connecting to the Elasticsearch endpoint.

`password` - (Optional) HTTP Basic Auth Password. See [Password ](#password) below for details.

`user_name` - (Optional) HTTP Basic Auth User Name (`String`).

### Auth None

No Authentication for the Elasticsearch endpoint.

### Auth Token

Configure an Access Token for authentication to the HTTP(s) server (such as a Bearer Token).

`token` - (Optional) F5XC Secret. URL for token, needs to be fetched from this path. See [Token ](#token) below for details.

### Aws Cloud Watch Receiver

Send logs to AWS Cloudwatch.

`aws_cred` - (Required) Reference to AWS Cloud Credentials for access to the Cloudwatch Logs. See [ref](#ref) below for details.

`aws_region` - (Required) AWS Region Name (`String`).

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Batch ](#batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Compression ](#compression) below for details.

`group_name` - (Required) The group name of the target Cloudwatch Logs stream (`String`).

`stream_name` - (Required) Note that there can only be one writer to a log stream at a time (`String`).

### Azure Event Hubs Receiver

Send logs to Azure Event Hubs.

`connection_string` - (Required) Azure Event Hubs Connection String.. See [Connection String ](#connection-string) below for details.

`instance` - (Required) Event Hubs Instance name into which logs should be stored (`String`).

`namespace` - (Required) Event Hubs Namespace is namespace with instance into which logs should be stored (`String`).

### Azure Receiver

Send logs to Azure Blob Storage.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Batch ](#batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Compression ](#compression) below for details.

`connection_string` - (Required) Azure Blob Storate Connection String. Note that this field must contain: `AccountKey`, `AccountName` and should contain `DefaultEndpointsProtocol`. See [Connection String ](#connection-string) below for details.

`container_name` - (Required) Container Name is the name of the container into which logs should be stored (`String`).

### Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (bool).

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (bool).

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (bool).

### Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

`compression_gzip` - (Optional) Gzip Compression (bool).

`compression_none` - (Optional) No Compression (bool).

### Compression Gzip

Gzip Compression.

### Compression None

No Compression.

### Connection String

Azure Event Hubs Connection String..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Datadog Api Key

Secret API key to access the datadog server.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Datadog Receiver

Send logs to a Datadog service.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Batch ](#batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Compression ](#compression) below for details.

`datadog_api_key` - (Required) Secret API key to access the datadog server. See [Datadog Api Key ](#datadog-api-key) below for details.

`endpoint` - (Optional) Datadog Endpoint, example: `example.com:9000` (`String`).

`site` - (Optional) Datadog Site, example: `datadoghq.com` (`String`).

`no_tls` - (Optional) Do not use TLS for the client connection (bool).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Use Tls ](#use-tls) below for details.

### Disable Verify Certificate

x-displayName: "Skip Server Certificate Verification".

### Disable Verify Hostname

x-displayName: "Skip Server Hostname Verification".

### Elastic Receiver

Send logs to an Elasticsearch endpoint.

`auth_aws` - (Optional) Reference to AWS Cloud Credentials for Authentication when connecting to the Elasticsearch Endpoint. See [ref](#ref) below for details.

`auth_basic` - (Optional) Basic Authentication specify that HTTP Basic Authentication should be used when connecting to the Elasticsearch endpoint. See [Auth Basic ](#auth-basic) below for details.

`auth_none` - (Optional) No Authentication for the Elasticsearch endpoint (bool).

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Batch ](#batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Compression ](#compression) below for details.

`endpoint` - (Required) Elasticsearch Endpoint URL, example `http://10.9.8.7:9000` (`String`).

`no_tls` - (Optional) Do not use TLS for the client connection (bool).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Use Tls ](#use-tls) below for details.

### Enable Verify Certificate

x-displayName: "Perform Server Certificate Verification".

### Enable Verify Hostname

x-displayName: "Enable Server Hostname Verification".

### Eu

EU Endpoint.

### Http Receiver

Send logs to a generic HTTP(s) server.

`auth_basic` - (Optional) Use HTTP Basic Auth for authentication to the HTPP(s) server. See [Auth Basic ](#auth-basic) below for details.

`auth_none` - (Optional) No Authentication (bool).

`auth_token` - (Optional) Configure an Access Token for authentication to the HTTP(s) server (such as a Bearer Token). See [Auth Token ](#auth-token) below for details.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Batch ](#batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Compression ](#compression) below for details.

`no_tls` - (Optional) Do not use TLS for the client connection (bool).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Use Tls ](#use-tls) below for details.

`uri` - (Required) HTTP Uri is the Uri of the HTTP endpoint to send logs to, example: `http://example.com:9000/logs` (`String`).

### Kafka Receiver

Send logs to a Kafka cluster.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Batch ](#batch) below for details.

`bootstrap_servers` - (Required) List of host:port pairs of the Kafka brokers (`String`).

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Compression ](#compression) below for details.

`kafka_topic` - (Required) The Kafka topic name to write events to (`String`).

`no_tls` - (Optional) Do not use TLS for the client connection (bool).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Use Tls ](#use-tls) below for details.

### Key Url

The data may be optionally secured using BlindFold..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Max Bytes Disabled

Batch Bytes Disabled.

### Max Events Disabled

Max Events Disabled.

### Mtls Disabled

mTLS is disabled.

### Mtls Enable

Enable mTLS configuration.

`certificate` - (Optional) Client certificate is PEM-encoded certificate or certificate-chain. (`String`).

`key_url` - (Optional) The data may be optionally secured using BlindFold.. See [Key Url ](#key-url) below for details.

### New Relic Receiver

Send logs to NewRelic.

`api_key` - (Required) A New Relic License Key. See [Api Key ](#api-key) below for details.

`eu` - (Optional) EU Endpoint (bool).

`us` - (Optional) US Endpoint (bool).

### No Ca

Do not use a CA Certificate.

### No Tls

Do not use TLS for the client connection.

### Ns List

x-displayName: "Select logs in specific namespaces".

`namespaces` - (Required) List of namespaces to stream logs for (`String`).

### Password

HTTP Basic Auth Password.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Qradar Receiver

Send logs to IBM QRadar.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Batch ](#batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Compression ](#compression) below for details.

`no_tls` - (Optional) Do not use TLS for the client connection (bool).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Use Tls ](#use-tls) below for details.

`uri` - (Required) Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to, example: `http://example.com:9000` (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### S3 Receiver

Send logs to an AWS S3 bucket.

`aws_cred` - (Required) Reference to AWS Cloud Credentials for access to the S3 bucket. See [ref](#ref) below for details.

`aws_region` - (Required) AWS Region Name (`String`).

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Batch ](#batch) below for details.

`bucket` - (Required) S3 Bucket Name (`String`).

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Compression ](#compression) below for details.

### Splunk Hec Token

Splunk HEC Logs secret Token.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Splunk Receiver

Send logs to a Splunk HEC Logs service.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Batch ](#batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Compression ](#compression) below for details.

`endpoint` - (Required) Splunk HEC Logs Endpoint, example: `https://http-input-hec.splunkcloud.com` (`String`).

`splunk_hec_token` - (Required) Splunk HEC Logs secret Token. See [Splunk Hec Token ](#splunk-hec-token) below for details.

`no_tls` - (Optional) Do not use TLS for the client connection (bool).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Use Tls ](#use-tls) below for details.

### Sumo Logic Receiver

Send logs to SumoLogic.

`url` - (Required) The HTTP Source Address URL for the desired SumoLogic HTTP Collector. See [Url ](#url) below for details.

### Timeout Seconds Default

Use Default Timeout (300 seconds).

### Token

F5XC Secret. URL for token, needs to be fetched from this path.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Url

The HTTP Source Address URL for the desired SumoLogic HTTP Collector.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Us

US Endpoint.

### Use Tls

Use TLS for client connections to the endpoint.

`no_ca` - (Optional) Do not use a CA Certificate (bool).

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

`mtls_disabled` - (Optional) mTLS is disabled (bool).

`mtls_enable` - (Optional) Enable mTLS configuration. See [Mtls Enable ](#mtls-enable) below for details.

`disable_verify_certificate` - (Optional) x-displayName: "Skip Server Certificate Verification" (bool).

`enable_verify_certificate` - (Optional) x-displayName: "Perform Server Certificate Verification" (bool).

`disable_verify_hostname` - (Optional) x-displayName: "Skip Server Hostname Verification" (bool).

`enable_verify_hostname` - (Optional) x-displayName: "Enable Server Hostname Verification" (bool).

### Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured global_log_receiver.
