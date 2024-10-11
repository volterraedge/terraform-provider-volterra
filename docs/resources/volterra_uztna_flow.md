---

page_title: "Volterra: uztna_flow"
description: "The uztna_flow allows CRUD of Uztna Flow resource on Volterra SaaS"

---

Resource volterra_uztna_flow
============================

The Uztna Flow allows CRUD of Uztna Flow resource on Volterra SaaS

~> **Note:** Please refer to [Uztna Flow API docs](https://docs.cloud.f5.com/docs-v2/api/uztna-flow) to learn more

Example Usage
-------------

```hcl
resource "volterra_uztna_flow" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  flow_type {
    // One of the arguments from this list "geo_match saml_message" must be set

    saml_message {
      application_tag {
        saml_app_tag = "saml_app_tag"
      }

      provider_metadata {
        idp_provider_metadata = "idp_provider_metadata"

        service_provider_metadata = "service_provider_metadata"
      }

      service_provider_properties {
        audience_uri {
          // One of the arguments from this list "uniform_resource_locator uniform_resource_name" can be set

          uniform_resource_locator {
            url = "url"
          }
        }

        force_authentication = true

        relay_state = "relay_state"

        sign_authentication_request {
          // One of the arguments from this list "no_auth_req sign_auth" must be set

          sign_auth {
            certificate {
              name      = "test1"
              namespace = "staging"
              tenant    = "acmecorp"
            }
          }
        }

        type = "type"

        want_encrypted_assertion {
          // One of the arguments from this list "encrypt_req no_encrypt_req" must be set

          no_encrypt_req = true
        }

        want_signed_assertion = true
      }
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

`flow_type` - (Required) Flow Type. See [Flow Type ](#flow-type) below for details.

### Flow Type

Flow Type.

###### One of the arguments from this list "geo_match, saml_message" must be set

`geo_match` - (Optional) Geolocation Match. See [Flow Type Choice Geo Match ](#flow-type-choice-geo-match) below for details.

`saml_message` - (Optional) a single sign-on (SSO) experience for your applications.. See [Flow Type Choice Saml Message ](#flow-type-choice-saml-message) below for details.

### Audience Uri Choice Uniform Resource Locator

x-required.

`url` - (Required) x-required (`String`).

### Audience Uri Choice Uniform Resource Name

x-required.

`host_name` - (Required) Host name of Uniform resource (`String`).

`urn` - (Required) Name of Uniform resource (`String`).

### Flow Type Choice Geo Match

Geolocation Match.

`match_cri` - (Required) Rules to add Matching criteria. See [Geo Match Match Cri ](#geo-match-match-cri) below for details.

### Flow Type Choice Saml Message

a single sign-on (SSO) experience for your applications..

`application_tag` - (Required) Application Tagging. See [Saml Message Application Tag ](#saml-message-application-tag) below for details.

`provider_metadata` - (Required) Provide metadata is xml kind of input for Service provider metadata and IDP .. See [Saml Message Provider Metadata ](#saml-message-provider-metadata) below for details.

`service_provider_properties` - (Required) Service Provider Properties. See [Saml Message Service Provider Properties ](#saml-message-service-provider-properties) below for details.

### Geo Match Match Cri

Rules to add Matching criteria.

`custom_geo_location_selector` - (Optional) Select multiple geolocations. See [Match Cri Custom Geo Location Selector ](#match-cri-custom-geo-location-selector) below for details.

### Match Cri Custom Geo Location Selector

Select multiple geolocations.

`expressions` - (Required) expressions contains the kubernetes style label expression for selections. (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Saml Message Application Tag

Application Tagging.

`saml_app_tag` - (Required) x-required (`String`).

### Saml Message Provider Metadata

Provide metadata is xml kind of input for Service provider metadata and IDP ..

`idp_provider_metadata` - (Optional) including security certificates and endpoints used for single sign-on (SSO). (`String`).

`service_provider_metadata` - (Optional) Once downloaded, securely send this file to your identity provider so they can complete the federation. (`String`).

### Saml Message Service Provider Properties

Service Provider Properties.

`audience_uri` - (Required) x-required. See [Service Provider Properties Audience Uri ](#service-provider-properties-audience-uri) below for details.

`force_authentication` - (Optional) x-displayName: "Force Authentication" (`Bool`).

`relay_state` - (Optional) x-displayName: "Relay State" (`String`).

`sign_authentication_request` - (Required) x-required. See [Service Provider Properties Sign Authentication Request ](#service-provider-properties-sign-authentication-request) below for details.

`type` - (Required) x-required (`String`).

`want_encrypted_assertion` - (Required) x-required. See [Service Provider Properties Want Encrypted Assertion ](#service-provider-properties-want-encrypted-assertion) below for details.

`want_signed_assertion` - (Optional) x-displayName: "Want Signed Assertion" (`Bool`).

### Service Provider Properties Audience Uri

x-required.

###### One of the arguments from this list "uniform_resource_locator, uniform_resource_name" can be set

`uniform_resource_locator` - (Required) x-required. See [Audience Uri Choice Uniform Resource Locator ](#audience-uri-choice-uniform-resource-locator) below for details.

`uniform_resource_name` - (Required) x-required. See [Audience Uri Choice Uniform Resource Name ](#audience-uri-choice-uniform-resource-name) below for details.

### Service Provider Properties Sign Authentication Request

x-required.

###### One of the arguments from this list "no_auth_req, sign_auth" must be set

`no_auth_req` - (Optional) No Authentication is default authentication (`Bool`).

`sign_auth` - (Optional) Sign Authentication Request, upload certificate for key and certificate.. See [Sign Auth Type Choice Sign Auth ](#sign-auth-type-choice-sign-auth) below for details.

### Service Provider Properties Want Encrypted Assertion

x-required.

###### One of the arguments from this list "encrypt_req, no_encrypt_req" must be set

`encrypt_req` - (Optional) Encrypt assertions. See [Want Encrypted Request Encrypt Req ](#want-encrypted-request-encrypt-req) below for details.

`no_encrypt_req` - (Optional) Do not encrypt assertions (`Bool`).

### Sign Auth Type Choice No Auth Req

No Authentication is default authentication.

### Sign Auth Type Choice Sign Auth

Sign Authentication Request, upload certificate for key and certificate..

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this uztna flow. See [ref](#ref) below for details.

### Want Encrypted Request Encrypt Req

Encrypt assertions.

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this uztna flow. See [ref](#ref) below for details.

### Want Encrypted Request No Encrypt Req

Do not encrypt assertions.

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_flow.
