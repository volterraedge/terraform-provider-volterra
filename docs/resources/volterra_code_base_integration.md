---

page_title: "Volterra: code_base_integration"

description: "The code_base_integration allows CRUD of Code Base Integration resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------------

Resource volterra_code_base_integration
=======================================

The Code Base Integration allows CRUD of Code Base Integration resource on Volterra SaaS

~> **Note:** Please refer to [Code Base Integration API docs](https://docs.cloud.f5.com/docs-v2/api/code-base-integration) to learn more

Example Usage
-------------

```hcl
resource "volterra_code_base_integration" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
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

`code_base_integration` - (Optional) Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection details. See [Code Base Integration ](#code-base-integration) below for details.

### Code Base Integration

Choose your code base (e.g. GitHub, GitLab, Bitbucket, Azure) and provide credentials and connection details.

###### One of the arguments from this list "azure_repos, bitbucket, bitbucket_server, github, github_enterprise, gitlab, gitlab_enterprise" must be set

`azure_repos` - (Optional) x-displayName: "Azure Repos Integration". See [Type Azure Repos ](#type-azure-repos) below for details.

`bitbucket` - (Optional) x-displayName: "BitBucket Cloud Integration". See [Type Bitbucket ](#type-bitbucket) below for details.

`bitbucket_server` - (Optional) x-displayName: "BitBucket Server Integration". See [Type Bitbucket Server ](#type-bitbucket-server) below for details.

`github` - (Optional) x-displayName: "Github Integration". See [Type Github ](#type-github) below for details.

`github_enterprise` - (Optional) x-displayName: "Github Enterprise Integration". See [Type Github Enterprise ](#type-github-enterprise) below for details.

`gitlab` - (Optional) x-displayName: "GitLab Cloud Integration". See [Type Gitlab ](#type-gitlab) below for details.

`gitlab_enterprise` - (Optional) x-displayName: "GitLab Enterprise Integration". See [Type Gitlab Enterprise ](#type-gitlab-enterprise) below for details.

### Azure Repos Access Token

x-required.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Bitbucket Passwd

x-required.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Bitbucket Server Passwd

x-required.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Github Access Token

x-required.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Github Enterprise Access Token

x-required.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Gitlab Access Token

x-required.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Gitlab Enterprise Access Token

x-required.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Type Azure Repos

x-displayName: "Azure Repos Integration".

`access_token` - (Required) x-required. See [Azure Repos Access Token ](#azure-repos-access-token) below for details.

### Type Bitbucket

x-displayName: "BitBucket Cloud Integration".

`passwd` - (Required) x-required. See [Bitbucket Passwd ](#bitbucket-passwd) below for details.

`username` - (Required) x-required (`String`).

### Type Bitbucket Server

x-displayName: "BitBucket Server Integration".

`passwd` - (Required) x-required. See [Bitbucket Server Passwd ](#bitbucket-server-passwd) below for details.

`url` - (Required) x-required (`String`).

`username` - (Required) x-required (`String`).

`verify_ssl` - (Optional) x-displayName: "Verify SSL" (`Bool`).

### Type Github

x-displayName: "Github Integration".

`access_token` - (Required) x-required. See [Github Access Token ](#github-access-token) below for details.

`username` - (Required) x-required (`String`).

`verify_ssl` - (Optional) x-displayName: "GitHub Verify SSL" (`Bool`).

### Type Github Enterprise

x-displayName: "Github Enterprise Integration".

`access_token` - (Required) x-required. See [Github Enterprise Access Token ](#github-enterprise-access-token) below for details.

`hostname` - (Required) x-required (`String`).

`username` - (Required) x-required (`String`).

### Type Gitlab

x-displayName: "GitLab Cloud Integration".

`access_token` - (Required) x-required. See [Gitlab Access Token ](#gitlab-access-token) below for details.

### Type Gitlab Enterprise

x-displayName: "GitLab Enterprise Integration".

`access_token` - (Required) x-required. See [Gitlab Enterprise Access Token ](#gitlab-enterprise-access-token) below for details.

`url` - (Required) x-required (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured code_base_integration.
