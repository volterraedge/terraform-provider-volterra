---

page_title: "Volterra: uztna_flow"

description: "The uztna_flow allows CRUD of Uztna Flow resource on Volterra SaaS"
---------------------------------------------------------------------------------

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

      idp {
        idp_assert_verification_cert {
          certificate {
            name      = "test1"
            namespace = "staging"
            tenant    = "acmecorp"
          }
        }

        idp_security_property {
          // One of the arguments from this list "no yes" must be set

          no = true
        }

        issuer = "https://issuerentity.com"

        // One of the arguments from this list "post redirect" must be set

        post = true
        sso_url = "https://issuersignon.com"
      }

      service_provider_properties {
        audience_uri {
          // One of the arguments from this list "uniform_resource_locator" can be set

          uniform_resource_locator {
            url = "url"
          }
        }

        force_authentication = true

        relay_state = "relay_state"

        sign_authentication_request {
          // One of the arguments from this list "no_auth_req sign_auth" must be set

          no_auth_req = true
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

### Auth Request Signed No

Auth Request is not signed .

### Auth Request Signed Yes

Auth Request is signed .

###### One of the arguments from this list "rsa_sha1, rsa_sha256, rsa_sha384, rsa_sha512" must be set

`rsa_sha1` - (Optional) RSA-SHA1 (`Bool`).

`rsa_sha256` - (Optional) RSA-SHA256 (`Bool`).

`rsa_sha384` - (Optional) RSA-SHA384 (`Bool`).

`rsa_sha512` - (Optional) RSA-SHA512 (`Bool`).

### Flow Type Choice Geo Match

Geolocation Match.

`match_cri` - (Required) Rules to add Matching criteria. See [Geo Match Match Cri ](#geo-match-match-cri) below for details.

### Flow Type Choice Saml Message

a single sign-on (SSO) experience for your applications..

`application_tag` - (Required) Application Tagging. See [Saml Message Application Tag ](#saml-message-application-tag) below for details.

`idp` - (Required) including security certificates and endpoints used for single sign-on (SSO). . See [Saml Message Idp ](#saml-message-idp) below for details.

`service_provider_properties` - (Required) Service Provider Properties. See [Saml Message Service Provider Properties ](#saml-message-service-provider-properties) below for details.

### Geo Match Match Cri

Rules to add Matching criteria.

`custom_geo_location_selector` - (Optional) Select multiple geolocations. See [Match Cri Custom Geo Location Selector ](#match-cri-custom-geo-location-selector) below for details.

### Idp Idp Assert Verification Cert

Used for identity provider's assertion verification .

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this uztna flow. See [ref](#ref) below for details.

### Idp Idp Security Property

Identity Provider's Security Property.

###### One of the arguments from this list "no, yes" must be set

`no` - (Optional) Auth Request is not signed (`Bool`).

`yes` - (Optional) Auth Request is signed . See [Auth Request Signed Yes ](#auth-request-signed-yes) below for details.

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

### Saml Message Idp

including security certificates and endpoints used for single sign-on (SSO). .

`idp_assert_verification_cert` - (Optional) Used for identity provider's assertion verification . See [Idp Idp Assert Verification Cert ](#idp-idp-assert-verification-cert) below for details.

`idp_security_property` - (Required) Identity Provider's Security Property. See [Idp Idp Security Property ](#idp-idp-security-property) below for details.

`issuer` - (Required) and it is used to identify the entity in SAML messages (`String`).

###### One of the arguments from this list "post, redirect" must be set

`post` - (Optional) SSO Service binding as POST request (`Bool`).

`redirect` - (Optional) SSO Service binding as Redirect request (`Bool`).

`sso_url` - (Required) URL required for SSO authentication (`String`).

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

###### One of the arguments from this list "uniform_resource_locator" can be set

`uniform_resource_locator` - (Optional) x-required. See [Audience Uri Choice Uniform Resource Locator ](#audience-uri-choice-uniform-resource-locator) below for details.

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

### Sign Algo Rsa Sha1

RSA-SHA1.

### Sign Algo Rsa Sha256

RSA-SHA256.

### Sign Algo Rsa Sha384

RSA-SHA384.

### Sign Algo Rsa Sha512

RSA-SHA512.

### Sign Auth Type Choice No Auth Req

No Authentication is default authentication.

### Sign Auth Type Choice Sign Auth

Sign Authentication Request, upload certificate for key and certificate..

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this uztna flow. See [ref](#ref) below for details.

### Sso Service Binding Post

SSO Service binding as POST request.

### Sso Service Binding Redirect

SSO Service binding as Redirect request .

### Want Encrypted Request Encrypt Req

Encrypt assertions.

`certificate` - (Optional) Select/Add one or more TLS Certificate objects to associate with this uztna flow. See [ref](#ref) below for details.

### Want Encrypted Request No Encrypt Req

Do not encrypt assertions.

Attribute Reference
-------------------

-	`id` - This is the id of the configured uztna_flow.
