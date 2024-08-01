











---
page_title: "Volterra: ztna"
description: "The ztna allows CRUD of Ztna  resource on Volterra SaaS"
---
# Resource volterra_ztna

The Ztna  allows CRUD of Ztna  resource on Volterra SaaS

~> **Note:** Please refer to [Ztna  API docs](https://docs.cloud.f5.com/docs-v2/api/ztna) to learn more

## Example Usage

```hcl
resource "volterra_ztna" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
}

```

## Argument Reference

### Metadata Argument Reference
`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).


`description` - (Optional) Human readable description for the object (`String`).


`disable` - (Optional) A value of true will administratively disable the object (`Bool`).


`labels` - (Optional) by selector expression (`String`).


`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).


`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).



### Spec Argument Reference

`connectivity` - (Optional) List of Connectivity Policy Type . See [Connectivity ](#connectivity) below for details.





`per_session` - (Optional) List of configuration for per session policy type.. See [Per Session ](#per-session) below for details.


		


		


		








		







		






		









		








		


		




		





		





		











		




		


		




		













		




		




		



		




		


		




		


		




		












		








		


		






		


		


		























		






### Connectivity 

 List of Connectivity Policy Type .

`connectivities` - (Optional) Connectivity policy configuration. Select connectivity policy as PPP/DTLS (`String`).

`connectivity_policy_name` - (Optional) Name of connectivity policy (`String`).



### Per Session 

 List of configuration for per session policy type..

`external_server` - (Optional) List of configuration of external server . See [Per Session External Server ](#per-session-external-server) below for details.

`lease_pool_member` - (Optional) List of Lease pool member for Session Policy. See [Per Session Lease Pool Member ](#per-session-lease-pool-member) below for details.

`per_session_policy_name` - (Optional) This contains the name of session policy (`String`).

`resources` - (Optional) List of the Policy of session type.. See [Per Session Resources ](#per-session-resources) below for details.

`sequence_action` - (Optional) Active Directory Query/Resource Assign  . See [Per Session Sequence Action ](#per-session-sequence-action) below for details.



### Acl Entries 

 Entries Config for Acl Resource.

`action` - (Optional) action is either allow/reject for acl resource type  (`String`).

`destination_end_port` - (Optional) Destination End Port for acl resource type. (`Int`).

`destination_ip` - (Optional) Destination IP address for the acl resource type. . See [Entries Destination Ip ](#entries-destination-ip) below for details.

`destination_start_port` - (Optional) Destination start port for acl resource type. (`Int`).



### Active Directory Auth Active Directory Query 

 Send active directory query configuration on success of allow directory auth..

`deny_name` - (Optional) Send deny as next item if query gets failed. (`String`).

`filter` - (Optional) Filter setting for filtering out the assigned resources  (`String`).

`resource_assign` - (Optional) List of all assigned resources list . See [Active Directory Query Resource Assign ](#active-directory-query-resource-assign) below for details.

`server` - (Optional) Name of the server which gets assigned. (`String`).



### Active Directory Query Resource Assign 

 List of all assigned resources list .

`allow` - (Optional) This sends resource assign configuration. See [Resource Assign Allow ](#resource-assign-allow) below for details.



### Admin Password Blindfold Secret Info Internal 

 Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).



### Allow Rules Resource Assign 

 Rules for assigning the resources .

`acls` - (Optional) Put the list of all acls resource type name. (`String`).

`network_access_rule` - (Optional) List of network access rules configuration (`String`).

`web_top_name` - (Optional) Put the name of webtopfull resource type (`String`).



### Customization Password 

 This field contains the secret password for log on..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)



###### One of the arguments from this list "vault_secret_info, clear_secret_info, wingman_secret_info, blindfold_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.


`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.


`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)


`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)




### Entries Destination Ip 

 Destination IP address for the acl resource type. .




###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.


`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.




### External Server Admin Password 

 Secret password credential used for external server.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Admin Password Blindfold Secret Info Internal ](#admin-password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)



###### One of the arguments from this list "wingman_secret_info, blindfold_secret_info, vault_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.


`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.


`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)


`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)




### Item Type Active Directory Auth 

 Action type is active Directory auth..

`active_directory_query` - (Optional) Send active directory query configuration on success of allow directory auth.. See [Active Directory Auth Active Directory Query ](#active-directory-auth-active-directory-query) below for details.

`deny_name` - (Optional) Send deny as next item in case it fails to find the query. (`String`).

`expression` - (Optional) Expression string to filter the Directory query. (`String`).

`server` - (Optional) Server name on which log on is success. (`String`).



### Item Type Active Directory Query 

 Action type is active Directory query..

`deny_name` - (Optional) Send deny as next item if query gets failed. (`String`).

`filter` - (Optional) Filter setting for filtering out the assigned resources  (`String`).

`resource_assign` - (Optional) List of all assigned resources list . See [Active Directory Query Resource Assign ](#active-directory-query-resource-assign) below for details.

`server` - (Optional) Name of the server which gets assigned. (`String`).



### Item Type Log On 

 Action type is logon .

`active_directory_auth` - (Optional) After logon next item is to verify the group and directory.. See [Log On Active Directory Auth ](#log-on-active-directory-auth) below for details.

`customization` - (Optional) This contains language , username and secret to login. See [Log On Customization ](#log-on-customization) below for details.



### Item Type Resource Assign 

 Action type is resource assign..

`allow` - (Optional) This sends resource assign configuration. See [Resource Assign Allow ](#resource-assign-allow) below for details.



### Lease Pool Member End Address 

 End Address of Lease Pool member.




###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.


`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.




### Lease Pool Member Start Address 

 startAddress of Lease Pool member.




###### One of the arguments from this list "ipv4, ipv6" can be set

`ipv4` - (Optional) IPv4 Address. See [Ver Ipv4 ](#ver-ipv4) below for details.


`ipv6` - (Optional) IPv6 Address. See [Ver Ipv6 ](#ver-ipv6) below for details.




### Log On Active Directory Auth 

 After logon next item is to verify the group and directory..

`active_directory_query` - (Optional) Send active directory query configuration on success of allow directory auth.. See [Active Directory Auth Active Directory Query ](#active-directory-auth-active-directory-query) below for details.

`deny_name` - (Optional) Send deny as next item in case it fails to find the query. (`String`).

`expression` - (Optional) Expression string to filter the Directory query. (`String`).

`server` - (Optional) Server name on which log on is success. (`String`).



### Log On Customization 

 This contains language , username and secret to login.

`password` - (Optional) This field contains the secret password for log on.. See [Customization Password ](#customization-password) below for details.

`user_name` - (Optional)This contains the user name to log on. (`String`).



### Password Blindfold Secret Info Internal 

 Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).



### Per Session External Server 

 List of configuration of external server .

`admin_password` - (Optional) Secret password credential used for external server. See [External Server Admin Password ](#external-server-admin-password) below for details.

`admin_username` - (Optional) Credentials used for external server  (`String`).

`domain` - (Optional) Domain name for the external server   (`String`).

`server_type` - (Optional) serverType of External server (`String`).



### Per Session Lease Pool Member 

 List of Lease pool member for Session Policy.

`end_address` - (Optional) End Address of Lease Pool member. See [Lease Pool Member End Address ](#lease-pool-member-end-address) below for details.

`start_address` - (Optional) startAddress of Lease Pool member. See [Lease Pool Member Start Address ](#lease-pool-member-start-address) below for details.



### Per Session Resources 

 List of the Policy of session type..




###### One of the arguments from this list "web, connectivity, acl" can be set

`acl` - (Optional) Resource type is ACL.. See [Resource Type Acl ](#resource-type-acl) below for details.


`connectivity` - (Optional) Resource type is connectivity . (`Bool`).


`web` - (Optional) Resource type is webtopfull. (`Bool`).




### Per Session Sequence Action 

 Active Directory Query/Resource Assign  .

`connectivity_access_policy_name` - (Optional) Name of connectivity policy  (`String`).

`next_action` - (Optional) Active Directory Query/Resource Assign . See [Sequence Action Next Action ](#sequence-action-next-action) below for details.



### Resource Assign Allow 

 This sends resource assign configuration.

`allow_name` - (Optional) Send allow as next item in case of success (`String`).

`deny_name` - (Optional) Send deny as next item in case of failure (`String`).

`rules_resource_assign` - (Optional) Rules for assigning the resources . See [Allow Rules Resource Assign ](#allow-rules-resource-assign) below for details.

`server` - (Optional) Server name for assigning the resources (`String`).



### Resource Type Acl 

 Resource type is ACL..

`entries` - (Optional) Entries Config for Acl Resource. See [Acl Entries ](#acl-entries) below for details.



### Resource Type Connectivity 

 Resource type is connectivity ..



### Resource Type Web 

 Resource type is webtopfull..



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



### Sequence Action Next Action 

 Active Directory Query/Resource Assign .




###### One of the arguments from this list "resource_assign, log_on, active_directory_auth, active_directory_query" can be set

`active_directory_auth` - (Optional) Action type is active Directory auth.. See [Item Type Active Directory Auth ](#item-type-active-directory-auth) below for details.


`active_directory_query` - (Optional) Action type is active Directory query.. See [Item Type Active Directory Query ](#item-type-active-directory-query) below for details.


`log_on` - (Optional) Action type is logon . See [Item Type Log On ](#item-type-log-on) below for details.


`resource_assign` - (Optional) Action type is resource assign.. See [Item Type Resource Assign ](#item-type-resource-assign) below for details.




### Ver Ipv4 

 IPv4 Address.

`addr` - (Optional) IPv4 Address in string form with dot-decimal notation (`String`).



### Ver Ipv6 

 IPv6 Address.

`addr` - (Optional) e.g. '2001:db8:0:0:0:0:2:1' becomes '2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::' (`String`).



## Attribute Reference

* `id` - This is the id of the configured ztna.

