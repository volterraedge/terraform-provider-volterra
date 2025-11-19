---

page_title: "Volterra: global_log_receiver"

description: "The global_log_receiver allows CRUD of Global Log Receiver resource on Volterra SaaS"
---------------------------------------------------------------------------------------------------

Resource volterra_global_log_receiver
=====================================

The Global Log Receiver allows CRUD of Global Log Receiver resource on Volterra SaaS

~> **Note:** Please refer to [Global Log Receiver API docs](https://docs.cloud.f5.com/docs-v2/api/global-log-receiver) to learn more

Example Usage
-------------

```hcl
resource "volterra_global_log_receiver" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "ns_all ns_current ns_list ns_system" must be set

  ns_current = true

  // One of the arguments from this list "audit_logs dns_logs request_logs security_events" must be set

  request_logs = true

  // One of the arguments from this list "aws_cloud_watch_receiver azure_event_hubs_receiver azure_receiver datadog_receiver elastic_receiver gcp_bucket_receiver http_receiver kafka_receiver new_relic_receiver qradar_receiver s3_receiver splunk_receiver sumo_logic_receiver" must be set

  s3_receiver {
    aws_cred {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    aws_region = "us-east-1"

    batch {
      // One of the arguments from this list "max_bytes max_bytes_disabled" can be set

      max_bytes_disabled = true

      // One of the arguments from this list "max_events max_events_disabled" can be set

      max_events_disabled = true

      // One of the arguments from this list "timeout_seconds timeout_seconds_default" can be set

      timeout_seconds_default = true
    }

    bucket = "my-log-bucket"

    compression {
      // One of the arguments from this list "compression_default compression_gzip compression_none" must be set

      compression_none = true
    }

    filename_options {
      // One of the arguments from this list "custom_folder log_type_folder no_folder" can be set

      custom_folder = "logs"
    }
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

###### One of the arguments from this list "ns_all, ns_current, ns_list, ns_system" must be set

`ns_all` - (Optional) x-displayName: "Select logs from all namespaces" (`Bool`).

`ns_current` - (Optional) x-displayName: "Select logs from current namespace" (`Bool`).

`ns_list` - (Optional) x-displayName: "Select logs in specific namespaces". See [Filter Choice Ns List ](#filter-choice-ns-list) below for details.

`ns_system` - (Optional) x-displayName: "Select logs from System namespace" (`Bool`).(Deprecated)

###### One of the arguments from this list "audit_logs, dns_logs, request_logs, security_events" must be set

`audit_logs` - (Optional) Send Audit Logs (corresponding to Public Audit and Authentication) (`Bool`).

`dns_logs` - (Optional) Send DNS Requests Logs (corresponding to DNS requests received) (`Bool`).

`request_logs` - (Optional) Send Request Logs (corresponding to Load Balancer access logs) (`Bool`).

`security_events` - (Optional) Send Security Events (corresponding to e.g. WAF blocked events or malicious requests) (`Bool`).

###### One of the arguments from this list "aws_cloud_watch_receiver, azure_event_hubs_receiver, azure_receiver, datadog_receiver, elastic_receiver, gcp_bucket_receiver, http_receiver, kafka_receiver, new_relic_receiver, qradar_receiver, s3_receiver, splunk_receiver, sumo_logic_receiver" must be set

`aws_cloud_watch_receiver` - (Optional) Send logs to AWS Cloudwatch. See [Receiver Aws Cloud Watch Receiver ](#receiver-aws-cloud-watch-receiver) below for details.

`azure_event_hubs_receiver` - (Optional) Send logs to Azure Event Hubs. See [Receiver Azure Event Hubs Receiver ](#receiver-azure-event-hubs-receiver) below for details.

`azure_receiver` - (Optional) Send logs to Azure Blob Storage. See [Receiver Azure Receiver ](#receiver-azure-receiver) below for details.

`datadog_receiver` - (Optional) Send logs to a Datadog service. See [Receiver Datadog Receiver ](#receiver-datadog-receiver) below for details.

`elastic_receiver` - (Optional) Send logs to an Elasticsearch endpoint. See [Receiver Elastic Receiver ](#receiver-elastic-receiver) below for details.(Deprecated)

`gcp_bucket_receiver` - (Optional) Send logs to a GCP Bucket. See [Receiver Gcp Bucket Receiver ](#receiver-gcp-bucket-receiver) below for details.

`http_receiver` - (Optional) Send logs to a generic HTTP(s) server. See [Receiver Http Receiver ](#receiver-http-receiver) below for details.

`kafka_receiver` - (Optional) Send logs to a Kafka cluster. See [Receiver Kafka Receiver ](#receiver-kafka-receiver) below for details.

`new_relic_receiver` - (Optional) Send logs to NewRelic. See [Receiver New Relic Receiver ](#receiver-new-relic-receiver) below for details.

`qradar_receiver` - (Optional) Send logs to IBM QRadar. See [Receiver Qradar Receiver ](#receiver-qradar-receiver) below for details.

`s3_receiver` - (Optional) Send logs to an AWS S3 bucket. See [Receiver S3 Receiver ](#receiver-s3-receiver) below for details.

`splunk_receiver` - (Optional) Send logs to a Splunk HEC Logs service. See [Receiver Splunk Receiver ](#receiver-splunk-receiver) below for details.

`sumo_logic_receiver` - (Optional) Send logs to SumoLogic. See [Receiver Sumo Logic Receiver ](#receiver-sumo-logic-receiver) below for details.

### Api Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Auth Basic Password

HTTP Basic Auth Password.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Auth Choice Auth Basic

Basic Authentication specify that HTTP Basic Authentication should be used when connecting to the Elasticsearch endpoint.

`password` - (Optional) HTTP Basic Auth Password. See [Auth Basic Password ](#auth-basic-password) below for details.

`user_name` - (Optional) HTTP Basic Auth User Name (`String`).

### Auth Choice Auth None

No Authentication for the Elasticsearch endpoint.

### Auth Choice Auth Token

Configure an Access Token for authentication to the HTTP(s) server (such as a Bearer Token).

`token` - (Optional) F5XC Secret. URL for token, needs to be fetched from this path. See [Auth Token Token ](#auth-token-token) below for details.

### Auth Token Token

F5XC Secret. URL for token, needs to be fetched from this path.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Token Blindfold Secret Info Internal ](#token-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Aws Cloud Watch Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### Aws Cloud Watch Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### Azure Event Hubs Receiver Connection String

Azure Event Hubs Connection String..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Connection String Blindfold Secret Info Internal ](#connection-string-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Azure Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### Azure Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### Azure Receiver Connection String

Azure Blob Storate Connection String. Note that this field must contain: `AccountKey`, `AccountName` and should contain `DefaultEndpointsProtocol`.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Connection String Blindfold Secret Info Internal ](#connection-string-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Azure Receiver Filename Options

Filename Options allow customization of filename and folder paths for the blob.

###### One of the arguments from this list "custom_folder, log_type_folder, no_folder" can be set

`custom_folder` - (Optional) The folder name must match `/^[a-z_][a-z0-9\\-\\._]*$/i` (`String`).

`log_type_folder` - (Optional) depending on LogType selection, this will be one of: access, security, audit or dns (`Bool`).

`no_folder` - (Optional) Do not use a folder path (`Bool`).

### Batch Bytes Max Bytes Disabled

Batch Bytes Disabled.

### Batch Events Max Events Disabled

Max Events Disabled.

### Batch Timeout Timeout Seconds Default

Use Default Timeout (300 seconds).

### Ca Choice No Ca

Do not use a CA Certificate.

### Compression Choice Compression Default

Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar.

### Compression Choice Compression Gzip

Gzip Compression.

### Compression Choice Compression None

No Compression.

### Connection String Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Datadog Api Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Datadog Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### Datadog Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### Datadog Receiver Datadog Api Key

Secret API key to access the datadog server.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Datadog Api Key Blindfold Secret Info Internal ](#datadog-api-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Elastic Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### Elastic Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### Endpoint Choice Eu

EU Endpoint.

### Endpoint Choice Us

US Endpoint.

### Filter Choice Ns List

x-displayName: "Select logs in specific namespaces".

`namespaces` - (Required) List of namespaces to stream logs for (`String`).

### Folder Log Type Folder

depending on LogType selection, this will be one of: access, security, audit or dns.

### Folder No Folder

Do not use a folder path.

### Gcp Bucket Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### Gcp Bucket Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### Gcp Bucket Receiver Filename Options

Filename Options allow customization of filename and folder paths for the bucket.

###### One of the arguments from this list "custom_folder, log_type_folder, no_folder" can be set

`custom_folder` - (Optional) The folder name must match `/^[a-z_][a-z0-9\\-\\._]*$/i` (`String`).

`log_type_folder` - (Optional) depending on LogType selection, this will be one of: access, security, audit or dns (`Bool`).

`no_folder` - (Optional) Do not use a folder path (`Bool`).

### Http Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### Http Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### Kafka Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### Kafka Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### Key Url Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Mtls Choice Mtls Disabled

mTLS is disabled.

### Mtls Choice Mtls Enable

Enable mTLS configuration.

`certificate` - (Optional) Client certificate is PEM-encoded certificate or certificate-chain. (`String`).

`key_url` - (Optional) The data may be optionally secured using BlindFold.. See [Mtls Enable Key Url ](#mtls-enable-key-url) below for details.

### Mtls Enable Key Url

The data may be optionally secured using BlindFold..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Key Url Blindfold Secret Info Internal ](#key-url-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### New Relic Receiver Api Key

A New Relic License Key.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Api Key Blindfold Secret Info Internal ](#api-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Password Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Qradar Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### Qradar Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### Receiver Aws Cloud Watch Receiver

Send logs to AWS Cloudwatch.

`aws_cred` - (Required) Reference to AWS Cloud Credentials for access to the Cloudwatch Logs. See [ref](#ref) below for details.

`aws_region` - (Required) AWS Region Name (`String`).

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Aws Cloud Watch Receiver Batch ](#aws-cloud-watch-receiver-batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Aws Cloud Watch Receiver Compression ](#aws-cloud-watch-receiver-compression) below for details.

`group_name` - (Required) The group name of the target Cloudwatch Logs stream (`String`).

`stream_name` - (Required) Note that there can only be one writer to a log stream at a time (`String`).

### Receiver Azure Event Hubs Receiver

Send logs to Azure Event Hubs.

`connection_string` - (Required) Azure Event Hubs Connection String.. See [Azure Event Hubs Receiver Connection String ](#azure-event-hubs-receiver-connection-string) below for details.

`instance` - (Required) Event Hubs Instance name into which logs should be stored (`String`).

`namespace` - (Required) Event Hubs Namespace is namespace with instance into which logs should be stored (`String`).

### Receiver Azure Receiver

Send logs to Azure Blob Storage.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Azure Receiver Batch ](#azure-receiver-batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Azure Receiver Compression ](#azure-receiver-compression) below for details.

`connection_string` - (Required) Azure Blob Storate Connection String. Note that this field must contain: `AccountKey`, `AccountName` and should contain `DefaultEndpointsProtocol`. See [Azure Receiver Connection String ](#azure-receiver-connection-string) below for details.

`container_name` - (Required) Container Name is the name of the container into which logs should be stored (`String`).

`filename_options` - (Optional) Filename Options allow customization of filename and folder paths for the blob. See [Azure Receiver Filename Options ](#azure-receiver-filename-options) below for details.

### Receiver Datadog Receiver

Send logs to a Datadog service.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Datadog Receiver Batch ](#datadog-receiver-batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Datadog Receiver Compression ](#datadog-receiver-compression) below for details.

`datadog_api_key` - (Required) Secret API key to access the datadog server. See [Datadog Receiver Datadog Api Key ](#datadog-receiver-datadog-api-key) below for details.

###### One of the arguments from this list "endpoint, site" must be set

`endpoint` - (Optional) Datadog Endpoint, example: `example.com:9000` (`String`).

`site` - (Optional) Datadog Site, example: `datadoghq.com` (`String`).

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) Do not use TLS for the client connection (`Bool`).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

### Receiver Elastic Receiver

Send logs to an Elasticsearch endpoint.

###### One of the arguments from this list "auth_aws, auth_basic, auth_none" must be set

`auth_aws` - (Optional) Reference to AWS Cloud Credentials for Authentication when connecting to the Elasticsearch Endpoint. See [ref](#ref) below for details.(Deprecated)

`auth_basic` - (Optional) Basic Authentication specify that HTTP Basic Authentication should be used when connecting to the Elasticsearch endpoint. See [Auth Choice Auth Basic ](#auth-choice-auth-basic) below for details.

`auth_none` - (Optional) No Authentication for the Elasticsearch endpoint (`Bool`).

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Elastic Receiver Batch ](#elastic-receiver-batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Elastic Receiver Compression ](#elastic-receiver-compression) below for details.

`endpoint` - (Required) Elasticsearch Endpoint URL, example `http://10.9.8.7:9000` (`String`).

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) Do not use TLS for the client connection (`Bool`).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

### Receiver Gcp Bucket Receiver

Send logs to a GCP Bucket.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Gcp Bucket Receiver Batch ](#gcp-bucket-receiver-batch) below for details.

`bucket` - (Required) GCP Bucket Name (`String`).

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Gcp Bucket Receiver Compression ](#gcp-bucket-receiver-compression) below for details.

`filename_options` - (Optional) Filename Options allow customization of filename and folder paths for the bucket. See [Gcp Bucket Receiver Filename Options ](#gcp-bucket-receiver-filename-options) below for details.

`gcp_cred` - (Required) Reference to GCP Cloud Credentials for access to the GCP bucket. See [ref](#ref) below for details.

### Receiver Http Receiver

Send logs to a generic HTTP(s) server.

###### One of the arguments from this list "auth_basic, auth_none, auth_token" must be set

`auth_basic` - (Optional) Use HTTP Basic Auth for authentication to the HTPP(s) server. See [Auth Choice Auth Basic ](#auth-choice-auth-basic) below for details.

`auth_none` - (Optional) No Authentication (`Bool`).

`auth_token` - (Optional) Configure an Access Token for authentication to the HTTP(s) server (such as a Bearer Token). See [Auth Choice Auth Token ](#auth-choice-auth-token) below for details.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Http Receiver Batch ](#http-receiver-batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Http Receiver Compression ](#http-receiver-compression) below for details.

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) Do not use TLS for the client connection (`Bool`).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

`uri` - (Required) HTTP Uri is the Uri of the HTTP endpoint to send logs to, example: `http://example.com:9000/logs` (`String`).

### Receiver Kafka Receiver

Send logs to a Kafka cluster.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Kafka Receiver Batch ](#kafka-receiver-batch) below for details.

`bootstrap_servers` - (Required) List of host:port pairs of the Kafka brokers (`String`).

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Kafka Receiver Compression ](#kafka-receiver-compression) below for details.

`kafka_topic` - (Required) The Kafka topic name to write events to (`String`).

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) Do not use TLS for the client connection (`Bool`).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

### Receiver New Relic Receiver

Send logs to NewRelic.

`api_key` - (Required) A New Relic License Key. See [New Relic Receiver Api Key ](#new-relic-receiver-api-key) below for details.

###### One of the arguments from this list "eu, us" must be set

`eu` - (Optional) EU Endpoint (`Bool`).

`us` - (Optional) US Endpoint (`Bool`).

### Receiver Qradar Receiver

Send logs to IBM QRadar.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Qradar Receiver Batch ](#qradar-receiver-batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Qradar Receiver Compression ](#qradar-receiver-compression) below for details.

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) Do not use TLS for the client connection (`Bool`).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

`uri` - (Required) Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to, example: `http://example.com:9000` (`String`).

### Receiver S3 Receiver

Send logs to an AWS S3 bucket.

`aws_cred` - (Required) Reference to AWS Cloud Credentials for access to the S3 bucket. See [ref](#ref) below for details.

`aws_region` - (Required) AWS Region Name (`String`).

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [S3 Receiver Batch ](#s3-receiver-batch) below for details.

`bucket` - (Required) S3 Bucket Name (`String`).

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [S3 Receiver Compression ](#s3-receiver-compression) below for details.

`filename_options` - (Optional) Filename Options allow customization of filename and folder paths for the bucket. See [S3 Receiver Filename Options ](#s3-receiver-filename-options) below for details.

### Receiver Splunk Receiver

Send logs to a Splunk HEC Logs service.

`batch` - (Optional) Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint. See [Splunk Receiver Batch ](#splunk-receiver-batch) below for details.

`compression` - (Optional) Compression Options allows selection of how data should be compressed when sent to the endpoint. See [Splunk Receiver Compression ](#splunk-receiver-compression) below for details.

`endpoint` - (Required) Splunk HEC Logs Endpoint, example: `https://http-input-hec.splunkcloud.com` (Note: must not contain `/services/collector`) (`String`).

`splunk_hec_token` - (Required) Splunk HEC Logs secret Token. See [Splunk Receiver Splunk Hec Token ](#splunk-receiver-splunk-hec-token) below for details.

###### One of the arguments from this list "no_tls, use_tls" must be set

`no_tls` - (Optional) Do not use TLS for the client connection (`Bool`).

`use_tls` - (Optional) Use TLS for client connections to the endpoint. See [Tls Choice Use Tls ](#tls-choice-use-tls) below for details.

### Receiver Sumo Logic Receiver

Send logs to SumoLogic.

`url` - (Required) The HTTP Source Address URL for the desired SumoLogic HTTP Collector. See [Sumo Logic Receiver Url ](#sumo-logic-receiver-url) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### S3 Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### S3 Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### S3 Receiver Filename Options

Filename Options allow customization of filename and folder paths for the bucket.

###### One of the arguments from this list "custom_folder, log_type_folder, no_folder" can be set

`custom_folder` - (Optional) The folder name must match `/^[a-z_][a-z0-9\\-\\._]*$/i` (`String`).

`log_type_folder` - (Optional) depending on LogType selection, this will be one of: access, security, audit or dns (`Bool`).

`no_folder` - (Optional) Do not use a folder path (`Bool`).

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Secret Info Oneof Vault Secret Info

Vault Secret is used for the secrets managed by Hashicorp Vault.

`key` - (Optional) If not provided entire secret will be returned. (`String`).

`location` - (Required) Path to secret in Vault. (`String`).

`provider` - (Required) Name of the Secret Management Access object that contains information about the backend Vault. (`String`).

`secret_encoding` - (Optional) This field defines the encoding type of the secret BEFORE the secret is put into Hashicorp Vault. (`String`).

`version` - (Optional) If not provided latest version will be returned. (`Int`).

### Secret Info Oneof Wingman Secret Info

Secret is given as bootstrap secret in F5XC Security Sidecar.

`name` - (Required) Name of the secret. (`String`).

### Splunk Hec Token Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Splunk Receiver Batch

Batch Options allow tuning of the conditions for how batches of logs are sent to the endpoint.

###### One of the arguments from this list "max_bytes, max_bytes_disabled" can be set

`max_bytes` - (Optional) Send batch to endpoint after the batch is equal to or larger than this many bytes (`Int`).

`max_bytes_disabled` - (Optional) Batch Bytes Disabled (`Bool`).

###### One of the arguments from this list "max_events, max_events_disabled" can be set

`max_events` - (Optional) Send batch to endpoint after this many log messages are in the batch (`Int`).

`max_events_disabled` - (Optional) Max Events Disabled (`Bool`).

###### One of the arguments from this list "timeout_seconds, timeout_seconds_default" can be set

`timeout_seconds` - (Optional) Send batch to the endpoint after this many seconds (`Int`).

`timeout_seconds_default` - (Optional) Use Default Timeout (300 seconds) (`Bool`).

### Splunk Receiver Compression

Compression Options allows selection of how data should be compressed when sent to the endpoint.

###### One of the arguments from this list "compression_default, compression_gzip, compression_none" must be set

`compression_default` - (Optional) Default Compression defaults to gzip for all endpoint types, except: HTTP, QRadar (`Bool`).

`compression_gzip` - (Optional) Gzip Compression (`Bool`).

`compression_none` - (Optional) No Compression (`Bool`).

### Splunk Receiver Splunk Hec Token

Splunk HEC Logs secret Token.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Splunk Hec Token Blindfold Secret Info Internal ](#splunk-hec-token-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Sumo Logic Receiver Url

The HTTP Source Address URL for the desired SumoLogic HTTP Collector.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Url Blindfold Secret Info Internal ](#url-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Tls Choice No Tls

Do not use TLS for the client connection.

### Tls Choice Use Tls

Use TLS for client connections to the endpoint.

###### One of the arguments from this list "no_ca, trusted_ca_url" must be set

`no_ca` - (Optional) Do not use a CA Certificate (`Bool`).

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

###### One of the arguments from this list "mtls_disabled, mtls_enable" must be set

`mtls_disabled` - (Optional) mTLS is disabled (`Bool`).

`mtls_enable` - (Optional) Enable mTLS configuration. See [Mtls Choice Mtls Enable ](#mtls-choice-mtls-enable) below for details.

###### One of the arguments from this list "disable_verify_certificate, enable_verify_certificate" can be set

`disable_verify_certificate` - (Optional) x-displayName: "Skip Server Certificate Verification" (`Bool`).

`enable_verify_certificate` - (Optional) x-displayName: "Perform Server Certificate Verification" (`Bool`).

###### One of the arguments from this list "disable_verify_hostname, enable_verify_hostname" can be set

`disable_verify_hostname` - (Optional) x-displayName: "Skip Server Hostname Verification" (`Bool`).

`enable_verify_hostname` - (Optional) x-displayName: "Enable Server Hostname Verification" (`Bool`).

### Token Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Url Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Verify Certificate Disable Verify Certificate

x-displayName: "Skip Server Certificate Verification".

### Verify Certificate Enable Verify Certificate

x-displayName: "Perform Server Certificate Verification".

### Verify Hostname Disable Verify Hostname

x-displayName: "Skip Server Hostname Verification".

### Verify Hostname Enable Verify Hostname

x-displayName: "Enable Server Hostname Verification".

Attribute Reference
-------------------

-	`id` - This is the id of the configured global_log_receiver.
