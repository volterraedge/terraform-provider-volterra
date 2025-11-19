---

page_title: "Volterra: protected_application"

description: "The protected_application allows CRUD of Protected Application resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------------

Resource volterra_protected_application
=======================================

The Protected Application allows CRUD of Protected Application resource on Volterra SaaS

~> **Note:** Please refer to [Protected Application API docs](https://docs.cloud.f5.com/docs-v2/api/protected-application) to learn more

Example Usage
-------------

```hcl
resource "volterra_protected_application" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "adobe_commerce_connector big_ip_iapp cloudflare cloudfront custom_connector f5_big_ip not_applicable_connector salesforce_commerce_connector" must be set

  salesforce_commerce_connector = true
  region                        = ["US"]
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

###### One of the arguments from this list "adobe_commerce_connector, big_ip_iapp, cloudflare, cloudfront, custom_connector, f5_big_ip, not_applicable_connector, salesforce_commerce_connector" must be set

`adobe_commerce_connector` - (Optional) Get credentials needed to configure on Adobe(Magento) Commerce. (`Bool`).

`big_ip_iapp` - (Optional) iApp templates simplify the integration and configuration process. (`Bool`).

`cloudflare` - (Optional) Configure Cloudflare connector properties. See [Connector Type Choice Cloudflare ](#connector-type-choice-cloudflare) below for details.

`cloudfront` - (Optional) Configure Amazon CloudFront connector properties. See [Connector Type Choice Cloudfront ](#connector-type-choice-cloudfront) below for details.

`custom_connector` - (Optional) Get credentials needed to configure on other connectors. (`Bool`).

`f5_big_ip` - (Optional) Get credentials needed to configure on BIG-IP. (`Bool`).

`not_applicable_connector` - (Optional) Connector type is not applicable (`Bool`).(Deprecated)

`salesforce_commerce_connector` - (Optional) Get credentials needed to configure on Salesforce Commerce. (`Bool`).

`region` - (Required) Select the region where the application’s origin server resides. (`String`).

### Add Header Choice Add Header

Add a header containing automation information.

### Add Header Choice No Header

Do not add a header containing automation information.

### Aws Configuration Type Choice Aws Configuration Id Selector

Associate your AWS CloudFront distribution(s) with protected endpoint(s) by specifing IDs.

`ids` - (Required) Add AWS CloudFront distribution ID, e.g. ABCDEFGHI0JKLM (`String`).

### Aws Configuration Type Choice Aws Configuration Tag Selector

Associate your AWS CloudFront distribution(s) with protected endpoint(s) by using tag selector.

`tags` - (Required) key is a AWS tag name, and the value is regular expression to match (`String`).

### Aws Configuration Type Choice Disable Aws Configuration

Check this box if you would prefer to manage distribution manaully.

### Client Source Choice Http Header

Request header name and value pairs.

`headers` - (Required) List of HTTP header name and value pairs. See [Http Header Headers ](#http-header-headers) below for details.

### Client Type Choice Mobile Client

Mobile traffic channel..

###### One of the arguments from this list "block, continue" must be set

`block` - (Optional) Respond with a custom block message. See [Mitigation Block ](#mitigation-block) below for details.

`continue` - (Optional) Request is flagged and forwarded to the origin. See [Mitigation Continue ](#mitigation-continue) below for details.

### Client Type Choice Web Client

Web traffic.

###### One of the arguments from this list "block, continue, redirect" must be set

`block` - (Optional) Respond with a custom block message. See [Mitigation Block ](#mitigation-block) below for details.

`continue` - (Optional) Request is flagged and forwarded to the origin. See [Mitigation Continue ](#mitigation-continue) below for details.

`redirect` - (Optional) Respond with a redirect. See [Mitigation Redirect ](#mitigation-redirect) below for details.

### Client Type Choice Web Mobile Client

Web and mobile traffic channel..

###### One of the arguments from this list "block_mobile, continue_mobile" must be set

`block_mobile` - (Optional) Respond with a custom block message. See [Mobile Mitigation Block Mobile ](#mobile-mitigation-block-mobile) below for details.

`continue_mobile` - (Optional) Request is flagged and forwarded to the origin. See [Mobile Mitigation Continue Mobile ](#mobile-mitigation-continue-mobile) below for details.

###### One of the arguments from this list "block_web, continue_web, redirect_web" must be set

`block_web` - (Optional) Respond with a custom block message. See [Web Mitigation Block Web ](#web-mitigation-block-web) below for details.

`continue_web` - (Optional) Request is flagged and forwarded to the origin. See [Web Mitigation Continue Web ](#web-mitigation-continue-web) below for details.

`redirect_web` - (Optional) Respond with a redirect. See [Web Mitigation Redirect Web ](#web-mitigation-redirect-web) below for details.

### Cloudflare Protected Endpoints

List of protected endpoints (max 128 items).

###### One of the arguments from this list "mobile_client, web_client, web_mobile_client" must be set

`mobile_client` - (Optional) Mobile traffic channel.. See [Client Type Choice Mobile Client ](#client-type-choice-mobile-client) below for details.

`web_client` - (Optional) Web traffic. See [Client Type Choice Web Client ](#client-type-choice-web-client) below for details.

`web_mobile_client` - (Optional) Web and mobile traffic channel.. See [Client Type Choice Web Mobile Client ](#client-type-choice-web-mobile-client) below for details.

###### One of the arguments from this list "any_domain, domain" can be set

`any_domain` - (Optional) Any Domain (`Bool`).

`domain` - (Optional) Select Domain matcher. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

`http_methods` - (Required) List of HTTP methods. (`List of Strings`).

`metadata` - (Required) Name and description are for administrative convenience and do no impact the way the endpoint works. See [Protected Endpoints Metadata ](#protected-endpoints-metadata) below for details.

`path` - (Required) Accepts wildcards * to match multiple characters or ? to match a single character. See [Protected Endpoints Path ](#protected-endpoints-path) below for details.

`query` - (Optional) Enter a regular expression to match your query parameters of interest (`String`).

### Cloudflare Trusted Clients

Define your allowlists to skip Bot Defense inference processing.

###### One of the arguments from this list "http_header, ip_prefix" must be set

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IP prefix string (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Trusted Clients Metadata ](#trusted-clients-metadata) below for details.

### Cloudfront Protected Endpoints

List of protected endpoints (max 128 items).

###### One of the arguments from this list "mobile_client, web_client, web_mobile_client" must be set

`mobile_client` - (Optional) Mobile traffic channel.. See [Client Type Choice Mobile Client ](#client-type-choice-mobile-client) below for details.

`web_client` - (Optional) Web traffic. See [Client Type Choice Web Client ](#client-type-choice-web-client) below for details.

`web_mobile_client` - (Optional) Web and mobile traffic channel.. See [Client Type Choice Web Mobile Client ](#client-type-choice-web-mobile-client) below for details.

###### One of the arguments from this list "any_domain, domain" can be set

`any_domain` - (Optional) Any Domain (`Bool`).

`domain` - (Optional) Select Domain matcher. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

###### One of the arguments from this list "flow_label, undefined_flow_label" can be set

`flow_label` - (Optional) x-displayName: "Specify endpoint label category". See [Flow Label Choice Flow Label ](#flow-label-choice-flow-label) below for details.

`undefined_flow_label` - (Optional) x-displayName: "Undefined" (`Bool`).

`http_methods` - (Required) List of HTTP methods. (`List of Strings`).

`metadata` - (Required) Name and description are for administrative convenience and do no impact the way the endpoint works. See [Protected Endpoints Metadata ](#protected-endpoints-metadata) below for details.

`path` - (Required) Accepts wildcards * to match multiple characters or ? to match a single character (`String`).

`query` - (Optional) Enter a regular expression to match your query parameters of interest (`String`).

### Cloudfront Trusted Clients

Define your allowlists to skip Bot Defense inference processing.

###### One of the arguments from this list "http_header, ip_prefix" must be set

`http_header` - (Optional) Request header name and value pairs. See [Client Source Choice Http Header ](#client-source-choice-http-header) below for details.

`ip_prefix` - (Optional) IP prefix string (`String`).

`metadata` - (Required) Common attributes for the rule including name and description.. See [Trusted Clients Metadata ](#trusted-clients-metadata) below for details.

### Connector Type Choice Cloudflare

Configure Cloudflare connector properties.

`continue_mitigation_action_hdr` - (Optional) A case-insensitive HTTP header name for Continue Mitigation Action when add header selected. (`String`).

###### One of the arguments from this list "disable_js_insert, js_insert_all_pages, js_insert_all_pages_except, js_insertion_rules, manual_js_insert" must be set

`disable_js_insert` - (Optional) No web endpoints or to insert JavaScript manually (`Bool`).

`js_insert_all_pages` - (Optional) Insert JavaScript in all pages. See [Java Script Choice Js Insert All Pages ](#java-script-choice-js-insert-all-pages) below for details.(Deprecated)

`js_insert_all_pages_except` - (Optional) Insert JavaScript in all pages with the exceptions. See [Java Script Choice Js Insert All Pages Except ](#java-script-choice-js-insert-all-pages-except) below for details.(Deprecated)

`js_insertion_rules` - (Optional) Specify pages to include or exclude JavaScript. See [Java Script Choice Js Insertion Rules ](#java-script-choice-js-insertion-rules) below for details.

`manual_js_insert` - (Optional) Insert JavaScript manually. See [Java Script Choice Manual Js Insert ](#java-script-choice-manual-js-insert) below for details.

`loglevel` - (Optional) (e.g. Debug includes Error, Warning, and Informational) (`String`).

###### One of the arguments from this list "disable_mobile_sdk, mobile_sdk_config" must be set

`disable_mobile_sdk` - (Optional) (`Bool`).

`mobile_sdk_config` - (Optional). See [Mobile Sdk Choice Mobile Sdk Config ](#mobile-sdk-choice-mobile-sdk-config) below for details.

`protected_endpoints` - (Required) List of protected endpoints (max 128 items). See [Cloudflare Protected Endpoints ](#cloudflare-protected-endpoints) below for details.

`timeout` - (Optional) The timeout for the inference check, in milliseconds (`Int`).

`trusted_clients` - (Optional) Define your allowlists to skip Bot Defense inference processing. See [Cloudflare Trusted Clients ](#cloudflare-trusted-clients) below for details.

### Connector Type Choice Cloudfront

Configure Amazon CloudFront connector properties.

`aws_account_no` - (Optional) We need your AWS account number to provide you access to the F5 solution from your AWS account (`String`).(Deprecated)

###### One of the arguments from this list "aws_configuration_id_selector, aws_configuration_tag_selector, disable_aws_configuration" must be set

`aws_configuration_id_selector` - (Optional) Associate your AWS CloudFront distribution(s) with protected endpoint(s) by specifing IDs. See [Aws Configuration Type Choice Aws Configuration Id Selector ](#aws-configuration-type-choice-aws-configuration-id-selector) below for details.

`aws_configuration_tag_selector` - (Optional) Associate your AWS CloudFront distribution(s) with protected endpoint(s) by using tag selector. See [Aws Configuration Type Choice Aws Configuration Tag Selector ](#aws-configuration-type-choice-aws-configuration-tag-selector) below for details.

`disable_aws_configuration` - (Optional) Check this box if you would prefer to manage distribution manaully (`Bool`).

`continue_mitigation_action_hdr` - (Optional) A case-insensitive HTTP header name for Continue Mitigation Action when add header selected. (`String`).

`data_sample` - (Optional) for analysis (limit 1,048,576 == 1 MiByte) (`Int`).

###### One of the arguments from this list "disable_js_insert, js_insert_all_pages, js_insert_all_pages_except, js_insertion_rules, manual_js_insert" must be set

`disable_js_insert` - (Optional) No web endpoints (`Bool`).

`js_insert_all_pages` - (Optional) Insert JavaScript in all pages. See [Java Script Choice Js Insert All Pages ](#java-script-choice-js-insert-all-pages) below for details.(Deprecated)

`js_insert_all_pages_except` - (Optional) Insert JavaScript in all pages with the exceptions. See [Java Script Choice Js Insert All Pages Except ](#java-script-choice-js-insert-all-pages-except) below for details.(Deprecated)

`js_insertion_rules` - (Optional) Specify pages to include or exclude JavaScript. See [Java Script Choice Js Insertion Rules ](#java-script-choice-js-insertion-rules) below for details.

`manual_js_insert` - (Optional) Insert JavaScript manually. See [Java Script Choice Manual Js Insert ](#java-script-choice-manual-js-insert) below for details.

`js_download_path` - (Optional) If not specified, default to ‘/common.js’. (`String`).(Deprecated)

`loglevel` - (Optional) (e.g. Debug includes Error, Warning, and Informational) (`String`).

###### One of the arguments from this list "disable_mobile_sdk, mobile_sdk_config" must be set

`disable_mobile_sdk` - (Optional) (`Bool`).

`mobile_sdk_config` - (Optional). See [Mobile Sdk Choice Mobile Sdk Config ](#mobile-sdk-choice-mobile-sdk-config) below for details.

`protected_endpoints` - (Required) List of protected endpoints (max 128 items). See [Cloudfront Protected Endpoints ](#cloudfront-protected-endpoints) below for details.

`timeout` - (Optional) The timeout for the inference check, in milliseconds (`Int`).

`trusted_clients` - (Optional) Define your allowlists to skip Bot Defense inference processing. See [Cloudfront Trusted Clients ](#cloudfront-trusted-clients) below for details.

### Domain Matcher Choice Any Domain

Any Domain..

### Domain Matcher Choice Domain

Domain matcher..

###### One of the arguments from this list "exact_value, regex_value, suffix_value" must be set

`exact_value` - (Optional) Exact domain name. (`String`).

`regex_value` - (Optional) Regular Expression value for the domain name (`String`).

`suffix_value` - (Optional) Suffix of domain name e.g "xyz.com" will match "*.xyz.com" and "xyz.com" (`String`).

### Exclude List Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Exclude List Path

URI path matcher..

###### One of the arguments from this list "path, prefix, regex" must be set

`path` - (Optional) Exact path value to match (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

`regex` - (Optional) Regular expression of path match (e.g. the value .* will match on all paths) (`String`).

### Flow Label Choice Account Management

x-displayName: "Account Management".

###### One of the arguments from this list "create, password_reset" must be set

`create` - (Optional) x-displayName: "Account Creation" (`Bool`).

`password_reset` - (Optional) x-displayName: "Password Reset" (`Bool`).

### Flow Label Choice Authentication

x-displayName: "Authentication".

###### One of the arguments from this list "login, login_mfa, login_partner, logout, token_refresh" must be set

`login` - (Optional) x-displayName: "Login". See [Label Choice Login ](#label-choice-login) below for details.

`login_mfa` - (Optional) x-displayName: "Login MFA" (`Bool`).

`login_partner` - (Optional) x-displayName: "Login for a Channel Partner" (`Bool`).

`logout` - (Optional) x-displayName: "Logout" (`Bool`).

`token_refresh` - (Optional) x-displayName: "Token Refresh" (`Bool`).

### Flow Label Choice Financial Services

x-displayName: "Financial Services".

###### One of the arguments from this list "apply, money_transfer" must be set

`apply` - (Optional) x-displayName: "Apply for a Financial Service Account (e.g., credit card, banking, retirement account)" (`Bool`).

`money_transfer` - (Optional) x-displayName: "Money Transfer" (`Bool`).

### Flow Label Choice Flight

x-displayName: "Flight".

###### One of the arguments from this list "checkin" must be set

`checkin` - (Optional) x-displayName: "Check into Flight" (`Bool`).

### Flow Label Choice Flow Label

x-displayName: "Specify endpoint label category".

###### One of the arguments from this list "account_management, authentication, financial_services, flight, profile_management, search, shopping_gift_cards" must be set

`account_management` - (Optional) x-displayName: "Account Management". See [Flow Label Choice Account Management ](#flow-label-choice-account-management) below for details.

`authentication` - (Optional) x-displayName: "Authentication". See [Flow Label Choice Authentication ](#flow-label-choice-authentication) below for details.

`financial_services` - (Optional) x-displayName: "Financial Services". See [Flow Label Choice Financial Services ](#flow-label-choice-financial-services) below for details.

`flight` - (Optional) x-displayName: "Flight". See [Flow Label Choice Flight ](#flow-label-choice-flight) below for details.

`profile_management` - (Optional) x-displayName: "Profile Management". See [Flow Label Choice Profile Management ](#flow-label-choice-profile-management) below for details.

`search` - (Optional) x-displayName: "Search". See [Flow Label Choice Search ](#flow-label-choice-search) below for details.

`shopping_gift_cards` - (Optional) x-displayName: "Shopping & Gift Cards". See [Flow Label Choice Shopping Gift Cards ](#flow-label-choice-shopping-gift-cards) below for details.

### Flow Label Choice Profile Management

x-displayName: "Profile Management".

###### One of the arguments from this list "create, update, view" must be set

`create` - (Optional) x-displayName: "Profile Creation" (`Bool`).

`update` - (Optional) x-displayName: "Profile Update" (`Bool`).

`view` - (Optional) x-displayName: "Profile View" (`Bool`).

### Flow Label Choice Search

x-displayName: "Search".

###### One of the arguments from this list "flight_search, product_search, reservation_search, room_search" can be set

`flight_search` - (Optional) x-displayName: "Flight Search" (`Bool`).

`product_search` - (Optional) x-displayName: "Product Search" (`Bool`).

`reservation_search` - (Optional) x-displayName: "Reservation Search (e.g., sporting events, concerts)" (`Bool`).

`room_search` - (Optional) x-displayName: "Room Search" (`Bool`).

### Flow Label Choice Shopping Gift Cards

x-displayName: "Shopping & Gift Cards".

###### One of the arguments from this list "gift_card_make_purchase_with_gift_card, gift_card_validation, shop_add_to_cart, shop_checkout, shop_choose_seat, shop_enter_drawing_submission, shop_make_payment, shop_order, shop_price_inquiry, shop_promo_code_validation, shop_purchase_gift_card, shop_update_quantity" can be set

`gift_card_make_purchase_with_gift_card` - (Optional) x-displayName: "Purchase with Gift Card" (`Bool`).

`gift_card_validation` - (Optional) x-displayName: "Gift Card Validation" (`Bool`).

`shop_add_to_cart` - (Optional) x-displayName: "Add to Cart" (`Bool`).

`shop_checkout` - (Optional) x-displayName: "Checkout" (`Bool`).

`shop_choose_seat` - (Optional) x-displayName: "Select Seat(s)" (`Bool`).

`shop_enter_drawing_submission` - (Optional) x-displayName: "Enter Drawing Submission" (`Bool`).

`shop_make_payment` - (Optional) x-displayName: "Payment / Billing" (`Bool`).

`shop_order` - (Optional) x-displayName: "Order Submit" (`Bool`).

`shop_price_inquiry` - (Optional) x-displayName: "Price Inquiry" (`Bool`).

`shop_promo_code_validation` - (Optional) x-displayName: "Promo Code Validation" (`Bool`).

`shop_purchase_gift_card` - (Optional) x-displayName: "Purchase a Gift Card" (`Bool`).

`shop_update_quantity` - (Optional) x-displayName: "Update Quantity" (`Bool`).

### Flow Label Choice Undefined Flow Label

x-displayName: "Undefined".

### Http Header Headers

List of HTTP header name and value pairs.

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, regex" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Java Script Choice Disable Js Insert

No web endpoints or to insert JavaScript manually.

### Java Script Choice Js Insert All Pages

Insert JavaScript in all pages.

`javascript_location` - (Optional) Defines where to insert Bot Defense JavaScript in HTML page. (`String`).

`js_download_path` - (Optional) If not specified, default to ‘/common.js’. (`String`).

### Java Script Choice Js Insert All Pages Except

Insert JavaScript in all pages with the exceptions.

`exclude_list` - (Optional) Optional JavaScript insertions exclude list of domain and path matchers.. See [Js Insert All Pages Except Exclude List ](#js-insert-all-pages-except-exclude-list) below for details.

`javascript_location` - (Optional) Defines where to insert Bot Defense JavaScript in HTML page. (`String`).

`js_download_path` - (Optional) If not specified, default to ‘/common.js’. (`String`).

### Java Script Choice Js Insertion Rules

Specify pages to include or exclude JavaScript.

`exclude_list` - (Optional) Optional JavaScript insertions exclude list of domain and path matchers.. See [Js Insertion Rules Exclude List ](#js-insertion-rules-exclude-list) below for details.

`javascript_location` - (Optional) Defines where to insert Bot Defense JavaScript in HTML page. (`String`).

`js_download_path` - (Optional) If not specified, default to ‘/common.js’. (`String`).

`rules` - (Required) Required list of pages to insert Bot Defense client JavaScript.. See [Js Insertion Rules Rules ](#js-insertion-rules-rules) below for details.

### Java Script Choice Manual Js Insert

Insert JavaScript manually.

`javascript_mode` - (Required) The larger chunk can be loaded asynchronously or synchronously. It can also be cacheable or non-cacheable on the browser. (`String`).(Deprecated)

`js_download_path` - (Optional) If not specified, default to ‘/common.js’. (`String`).

### Js Insert All Pages Except Exclude List

Optional JavaScript insertions exclude list of domain and path matchers..

###### One of the arguments from this list "any_domain, domain" must be set

`any_domain` - (Optional) Any Domain. (`Bool`).

`domain` - (Optional) Domain matcher.. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Exclude List Metadata ](#exclude-list-metadata) below for details.

`path` - (Required) URI path matcher.. See [Exclude List Path ](#exclude-list-path) below for details.

### Js Insertion Rules Exclude List

Optional JavaScript insertions exclude list of domain and path matchers..

###### One of the arguments from this list "any_domain, domain" must be set

`any_domain` - (Optional) Any Domain. (`Bool`).

`domain` - (Optional) Domain matcher.. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Exclude List Metadata ](#exclude-list-metadata) below for details.

`path` - (Required) URI path matcher.. See [Exclude List Path ](#exclude-list-path) below for details.

### Js Insertion Rules Rules

Required list of pages to insert Bot Defense client JavaScript..

###### One of the arguments from this list "any_domain, domain" must be set

`any_domain` - (Optional) Any Domain. (`Bool`).

`domain` - (Optional) Domain matcher.. See [Domain Matcher Choice Domain ](#domain-matcher-choice-domain) below for details.

`metadata` - (Required) Common attributes for the rule including name and description.. See [Rules Metadata ](#rules-metadata) below for details.

###### One of the arguments from this list "exact_path, glob, prefix" must be set

`exact_path` - (Optional) Exact path value to match (`String`).

`glob` - (Optional) Accepts wildcards * to match multiple characters or ? to match a single character (`String`).

`prefix` - (Optional) Path prefix to match (e.g. the value / will match on all paths) (`String`).

### Label Choice Apply

x-displayName: "Apply for a Financial Service Account (e.g., credit card, banking, retirement account)".

### Label Choice Checkin

x-displayName: "Check into Flight".

### Label Choice Create

x-displayName: "Account Creation".

### Label Choice Flight Search

x-displayName: "Flight Search".

### Label Choice Gift Card Make Purchase With Gift Card

x-displayName: "Purchase with Gift Card".

### Label Choice Gift Card Validation

x-displayName: "Gift Card Validation".

### Label Choice Login

x-displayName: "Login".

### Label Choice Login Mfa

x-displayName: "Login MFA".

### Label Choice Login Partner

x-displayName: "Login for a Channel Partner".

### Label Choice Logout

x-displayName: "Logout".

### Label Choice Money Transfer

x-displayName: "Money Transfer".

### Label Choice Password Reset

x-displayName: "Password Reset".

### Label Choice Product Search

x-displayName: "Product Search".

### Label Choice Reservation Search

x-displayName: "Reservation Search (e.g., sporting events, concerts)".

### Label Choice Room Search

x-displayName: "Room Search".

### Label Choice Shop Add To Cart

x-displayName: "Add to Cart".

### Label Choice Shop Checkout

x-displayName: "Checkout".

### Label Choice Shop Choose Seat

x-displayName: "Select Seat(s)".

### Label Choice Shop Enter Drawing Submission

x-displayName: "Enter Drawing Submission".

### Label Choice Shop Make Payment

x-displayName: "Payment / Billing".

### Label Choice Shop Order

x-displayName: "Order Submit".

### Label Choice Shop Price Inquiry

x-displayName: "Price Inquiry".

### Label Choice Shop Promo Code Validation

x-displayName: "Promo Code Validation".

### Label Choice Shop Purchase Gift Card

x-displayName: "Purchase a Gift Card".

### Label Choice Shop Update Quantity

x-displayName: "Update Quantity".

### Label Choice Token Refresh

x-displayName: "Token Refresh".

### Label Choice Update

x-displayName: "Profile Update".

### Label Choice View

x-displayName: "Profile View".

### Mitigation Block

Respond with a custom block message.

`body` - (Optional) Custom body message (`String`).

`content_type` - (Optional) Content type to use in a block response (`String`).

`status` - (Required) HTTP Status code to respond with (`String`).

### Mitigation Block

Respond with a custom block message.

`body` - (Optional) Custom body message (`String`).

`content_type` - (Optional) Content type to use in a block response (`String`).

`status` - (Required) HTTP Status code to respond with (`String`).

### Mitigation Continue

Request is flagged and forwarded to the origin.

###### One of the arguments from this list "add_header, no_header" can be set

`add_header` - (Optional) Add a header containing automation information (`Bool`).

`no_header` - (Optional) Do not add a header containing automation information (`Bool`).

### Mitigation Redirect

Respond with a redirect.

`location` - (Required) URI location for redirect reponse (`String`).

`status` - (Required) HTTP Status code to respond with (`String`).

### Mobile Identifier Headers

A list of headers that can be used to identify mobile traffic..

`name` - (Required) Name of the header (`String`).

###### One of the arguments from this list "exact, regex" can be set

`exact` - (Optional) Header value to match exactly (`String`).

`regex` - (Optional) Regex match of the header value in re2 format (`String`).

### Mobile Mitigation Block Mobile

Respond with a custom block message.

`body` - (Optional) Custom body message (`String`).

`content_type` - (Optional) Content type to use in a block response (`String`).

`status` - (Required) HTTP Status code to respond with (`String`).

### Mobile Mitigation Continue Mobile

Request is flagged and forwarded to the origin.

###### One of the arguments from this list "add_header, no_header" can be set

`add_header` - (Optional) Add a header containing automation information (`Bool`).

`no_header` - (Optional) Do not add a header containing automation information (`Bool`).

### Mobile Sdk Choice Disable Mobile Sdk

.

### Mobile Sdk Choice Mobile Sdk Config

.

`mobile_identifier` - (Optional) To be used if web and native mobile app endpoints share same FQDN and URL path.. See [Mobile Sdk Config Mobile Identifier ](#mobile-sdk-config-mobile-identifier) below for details.

### Mobile Sdk Config Mobile Identifier

To be used if web and native mobile app endpoints share same FQDN and URL path..

`headers` - (Optional) A list of headers that can be used to identify mobile traffic.. See [Mobile Identifier Headers ](#mobile-identifier-headers) below for details.

### Protected Endpoints Metadata

Name and description are for administrative convenience and do no impact the way the endpoint works.

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Protected Endpoints Path

Accepts wildcards * to match multiple characters or ? to match a single character.

`caseinsensitive` - (Optional) Should path be searched case insensitive; (`Bool`).

`path` - (Required) Uri Path (`String`).

### Rules Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Transaction Result Failure Conditions

Failure Conditions.

`name` - (Optional) A case-insensitive HTTP header name. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`status` - (Required) HTTP Status code (`String`).

### Transaction Result Success Conditions

Success Conditions.

`name` - (Optional) A case-insensitive HTTP header name. (`String`).

`regex_values` - (Optional) A list of regular expressions to match the input against. (`String`).

`status` - (Required) HTTP Status code (`String`).

### Transaction Result Choice Disable Transaction Result

Disable collection of transaction result..

### Transaction Result Choice Transaction Result

Collect transaction result..

`failure_conditions` - (Optional) Failure Conditions. See [Transaction Result Failure Conditions ](#transaction-result-failure-conditions) below for details.

`success_conditions` - (Optional) Success Conditions. See [Transaction Result Success Conditions ](#transaction-result-success-conditions) below for details.

### Trusted Clients Metadata

Common attributes for the rule including name and description..

`description` - (Optional) Human readable description. (`String`).

`disable` - (Optional) A value of true will administratively disable the object that corresponds to the containing message. (`Bool`).(Deprecated)

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Web Mitigation Block Web

Respond with a custom block message.

`body` - (Optional) Custom body message (`String`).

`content_type` - (Optional) Content type to use in a block response (`String`).

`status` - (Required) HTTP Status code to respond with (`String`).

### Web Mitigation Continue Web

Request is flagged and forwarded to the origin.

###### One of the arguments from this list "add_header, no_header" can be set

`add_header` - (Optional) Add a header containing automation information (`Bool`).

`no_header` - (Optional) Do not add a header containing automation information (`Bool`).

### Web Mitigation Redirect Web

Respond with a redirect.

`location` - (Required) URI location for redirect reponse (`String`).

`status` - (Required) HTTP Status code to respond with (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured protected_application.
