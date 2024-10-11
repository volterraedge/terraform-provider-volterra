---

page_title: "Volterra: cloud_link"
description: "The cloud_link allows CRUD of Cloud Link resource on Volterra SaaS"

---

Resource volterra_cloud_link
============================

The Cloud Link allows CRUD of Cloud Link resource on Volterra SaaS

~> **Note:** Please refer to [Cloud Link API docs](https://docs.cloud.f5.com/docs-v2/api/cloud-link) to learn more

Example Usage
-------------

```hcl
resource "volterra_cloud_link" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "aws azure gcp" must be set

  aws {
    aws_cred {
      name      = "test1"
      namespace = "staging"
      tenant    = "acmecorp"
    }

    // One of the arguments from this list "byoc f5xc_managed" must be set

    byoc {
      connections {
        auth_key {
          blindfold_secret_info_internal {
            decryption_provider = "value"

            location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

            store_provider = "value"
          }

          secret_encoding_type = "secret_encoding_type"

          // One of the arguments from this list "blindfold_secret_info clear_secret_info vault_secret_info wingman_secret_info" must be set

          blindfold_secret_info {
            decryption_provider = "value"

            location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

            store_provider = "value"
          }
        }

        bgp_asn = "64512"

        connection_id = "dxcon-a2h48678"

        coordinates {
          latitude = "10.0"

          longitude = "20.0"
        }

        enable_sitelink = true

        // One of the arguments from this list "ipv4 ipv6" must be set

        ipv4 {
          aws_router_peer_address = "10.1.0.0/31"

          router_peer_address = "10.1.0.0/31"
        }
        jumbo_mtu = true
        metadata {
          description = "Virtual Host for acmecorp website"

          disable = true

          name = "acmecorp-web"
        }
        region = "us-east-1"

        // One of the arguments from this list "system_generated_name user_assigned_name" can be set

        system_generated_name = true
        tags = {
          "key1" = "value1"
        }
        virtual_interface_type = "virtual_interface_type"
        vlan = "700"
      }
    }

    // One of the arguments from this list "auto custom_asn" must be set

    auto = true
  }

  // One of the arguments from this list "disabled enabled" must be set

  disabled = true
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

###### One of the arguments from this list "aws, azure, gcp" must be set

`aws` - (Optional) CloudLink for AWS Cloud. See [Cloud Provider Aws ](#cloud-provider-aws) below for details.

`azure` - (Optional) CloudLink for Azure Cloud. See [Cloud Provider Azure ](#cloud-provider-azure) below for details.(Deprecated)

`gcp` - (Optional) CloudLink for Google Cloud Platform. See [Cloud Provider Gcp ](#cloud-provider-gcp) below for details.

###### One of the arguments from this list "disabled, enabled" must be set

`disabled` - (Optional) CloudLink connection to RE Site is disabled (`Bool`).

`enabled` - (Optional) CloudLink connection to RE Site is enabled. See [Enable Connection To Re Choice Enabled ](#enable-connection-to-re-choice-enabled) below for details.

### Auth Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Byoc Connections

List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but will be used for connecting sites and REs..

`auth_key` - (Required) This string has a minimum length of 6 characters and and a maximum length of 80 characters.. See [Connections Auth Key ](#connections-auth-key) below for details.

`bgp_asn` - (Required) router for the new virtual interface to be configured on AWS. (`Int`).

`connection_id` - (Required) Id of the existing AWS Direct Connect Connection (`String`).

`coordinates` - (Optional) Coordinates of the CloudLink Connection based on connection's physical location. See [Connections Coordinates ](#connections-coordinates) below for details.(Deprecated)

`enable_sitelink` - (Optional) Enable direct connectivity between Direct Connect points of presence. The Virtual Interface will be created with SiteLink enabled. (`Bool`).(Deprecated)

###### One of the arguments from this list "ipv4, ipv6" must be set

`ipv4` - (Optional) Configure BGP IPv4 peering for endpoints. See [Ip Type Ipv4 ](#ip-type-ipv4) below for details.

`ipv6` - (Optional) Configure BGP IPv6 peering for endpoints. See [Ip Type Ipv6 ](#ip-type-ipv6) below for details.(Deprecated)

`jumbo_mtu` - (Optional) The Virtual Interface will be created with option to use jumbo frames enabled (`Bool`).(Deprecated)

`metadata` - (Required) Specify attributes for the connection including name and description.. See [Connections Metadata ](#connections-metadata) below for details.

`region` - (Required) Region where the connection is setup (`String`).

###### One of the arguments from this list "system_generated_name, user_assigned_name" can be set

`system_generated_name` - (Optional) F5XC will automatically assign a AWS resource name (`Bool`).(Deprecated)

`user_assigned_name` - (Optional) User is managing the AWS resource name (`String`).(Deprecated)

`tags` - (Optional) Specified tags will be added to Virtual interface along with any F5XC specific tags (`String`).

`virtual_interface_type` - (Optional) Specifies the virtual interface type that needs to be configured on AWS (`String`).

`vlan` - (Required) This tag is required for any traffic traversing the AWS Direct Connect connection (`Int`).

### Byoc Connections

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud to facilitate seamless private connectivity. .

`coordinates` - (Optional) Coordinates of the CloudLink Connection based on connection's physical location. See [Connections Coordinates ](#connections-coordinates) below for details.(Deprecated)

`interconnect_attachment_name` - (Required) Name of already-existing GCP Cloud Interconnect Attachment (`String`).

`metadata` - (Required) Specify attributes for the connection including name and description.. See [Connections Metadata ](#connections-metadata) below for details.

###### One of the arguments from this list "project, same_as_credential" must be set

`project` - (Optional) Specify a GCP Project for the interconnect attachment (`String`).

`same_as_credential` - (Optional) GCP Project for the interconnect is the same as the project specified in the credential (`Bool`).

`region` - (Required) GCP Region in which the GCP Cloud Interconnect attachment is configured (`String`).

### Cloud Link Type Byoc

Assumption is that this given AWS account already has direct connect connection provisioned.

`connections` - (Required) List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but will be used for connecting sites and REs.. See [Byoc Connections ](#byoc-connections) below for details.

### Cloud Link Type Byoc

Assumption is that this given GCP account already has Cloud Interconnect provisioned..

`connections` - (Required)Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud to facilitate seamless private connectivity. . See [Byoc Connections ](#byoc-connections) below for details.

### Cloud Link Type F5xc Managed

F5 XC will manage end to end AWS Direct Connect Connection and making it ready to be consumed by the site..

### Cloud Provider Aws

CloudLink for AWS Cloud.

`aws_cred` - (Required) Reference to AWS cloud account credential object used to deploy CloudLink specific object. See [ref](#ref) below for details.

###### One of the arguments from this list "byoc, f5xc_managed" must be set

`byoc` - (Optional) Assumption is that this given AWS account already has direct connect connection provisioned. See [Cloud Link Type Byoc ](#cloud-link-type-byoc) below for details.

`f5xc_managed` - (Optional) F5 XC will manage end to end AWS Direct Connect Connection and making it ready to be consumed by the site.. See [Cloud Link Type F5xc Managed ](#cloud-link-type-f5xc-managed) below for details.(Deprecated)

###### One of the arguments from this list "auto, custom_asn" must be set

`auto` - (Optional) F5XC will automatically generate an ASN to create a Direct Connect Gateway (`Bool`).(Deprecated)

`custom_asn` - (Optional) F5XC will use custom ASN to create a Direct Connect Gateway (`Int`).

### Cloud Provider Azure

CloudLink for Azure Cloud.

### Cloud Provider Gcp

CloudLink for Google Cloud Platform.

###### One of the arguments from this list "byoc" must be set

`byoc` - (Optional) Assumption is that this given GCP account already has Cloud Interconnect provisioned.. See [Cloud Link Type Byoc ](#cloud-link-type-byoc) below for details.

`gcp_cred` - (Required) Reference to GCP cloud account credential object used to deploy CloudLink specific object. See [ref](#ref) below for details.

### Connections Auth Key

This string has a minimum length of 6 characters and and a maximum length of 80 characters..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Auth Key Blindfold Secret Info Internal ](#auth-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Connections Coordinates

Coordinates of the CloudLink Connection based on connection's physical location.

`latitude` - (Optional) Latitude of the site location (`Float`).

`longitude` - (Optional) longitude of site location (`Float`).

### Connections Metadata

Specify attributes for the connection including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Direct Connect Gateway Asn Choice Auto

F5XC will automatically generate an ASN to create a Direct Connect Gateway.

### Enable Connection To Re Choice Enabled

CloudLink connection to RE Site is enabled.

`cloudlink_network_name` - (Required) Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support. (`String`).

### Ip Type Ipv4

Configure BGP IPv4 peering for endpoints.

`aws_router_peer_address` - (Required) The BGP peer IP configured on the AWS endpoint (`String`).

`router_peer_address` - (Required) The BGP peer IP configured on your (customer) endpoint (`String`).

### Ip Type Ipv6

Configure BGP IPv6 peering for endpoints.

### Project Choice Same As Credential

GCP Project for the interconnect is the same as the project specified in the credential.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Resource Name Choice System Generated Name

F5XC will automatically assign a AWS resource name.

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

Attribute Reference
-------------------

-	`id` - This is the id of the configured cloud_link.
