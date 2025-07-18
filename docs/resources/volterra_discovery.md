---

page_title: "Volterra: discovery"

description: "The discovery allows CRUD of Discovery resource on Volterra SaaS"
-------------------------------------------------------------------------------

Resource volterra_discovery
===========================

The Discovery allows CRUD of Discovery resource on Volterra SaaS

~> **Note:** Please refer to [Discovery API docs](https://docs.cloud.f5.com/docs-v2/api/discovery) to learn more

Example Usage
-------------

```hcl
resource "volterra_discovery" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "cluster_id no_cluster_id" must be set

  no_cluster_id = true

  // One of the arguments from this list "discovery_cbip discovery_consul discovery_k8s" must be set

  discovery_k8s {
    access_info {
      // One of the arguments from this list "connection_info kubeconfig_url" must be set

      kubeconfig_url {
        // One of the arguments from this list "blindfold_secret_info clear_secret_info" must be set

        blindfold_secret_info {
          decryption_provider = "value"

          location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

          store_provider = "value"
        }
      }

      // One of the arguments from this list "isolated reachable" must be set

      isolated = true
    }

    publish_info {
      // One of the arguments from this list "disable dns_delegation publish publish_fqdns" must be set

      dns_delegation {
        dns_mode = "dns_mode"

        subdomain = "value"
      }
    }
  }
  where {
    // One of the arguments from this list "site virtual_network virtual_site" must be set

    virtual_site {
      // One of the arguments from this list "disable_internet_vip enable_internet_vip" must be set

      enable_internet_vip = true

      network_type = "network_type"

      ref {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
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

###### One of the arguments from this list "cluster_id, no_cluster_id" must be set

`cluster_id` - (Optional) specified in endpoint object to discover only from this discovery object. (`String`).

`no_cluster_id` - (Optional) of the site will discover from this discovery object. (`Bool`).

###### One of the arguments from this list "discovery_cbip, discovery_consul, discovery_k8s" must be set

`discovery_cbip` - (Optional) Discovery configuration for Classic BIG-IP. See [Discovery Choice Discovery Cbip ](#discovery-choice-discovery-cbip) below for details.

`discovery_consul` - (Optional) Discovery configuration for Hashicorp Consul. See [Discovery Choice Discovery Consul ](#discovery-choice-discovery-consul) below for details.

`discovery_k8s` - (Optional) Discovery configuration for K8s.. See [Discovery Choice Discovery K8s ](#discovery-choice-discovery-k8s) below for details.

`where` - (Required) All the sites where this discovery config is valid.. See [Where ](#where) below for details.

### Where

All the sites where this discovery config is valid..

###### One of the arguments from this list "site, virtual_network, virtual_site" must be set

`site` - (Optional) Direct reference to site object. See [Ref Or Selector Site ](#ref-or-selector-site) below for details.

`virtual_network` - (Optional) Direct reference to virtual network object. See [Ref Or Selector Virtual Network ](#ref-or-selector-virtual-network) below for details.

`virtual_site` - (Optional) Direct reference to virtual site object. See [Ref Or Selector Virtual Site ](#ref-or-selector-virtual-site) below for details.

### Access Info Connection Info

Configuration details to access Hashicorp Consul API service using REST..

`api_server` - (Required) API server must be a fully qualified domain string and port specified as host:port pair (`String`).

`tls_info` - (Optional) TLS settings to enable transport layer security. See [Connection Info Tls Info ](#connection-info-tls-info) below for details.

### Access Info Http Basic Auth Info

Username and password used for HTTP/HTTPS access.

`passwd_url` - (Optional) F5XC Secret. URL for password, needs to be fetched from this path. See [Http Basic Auth Info Passwd Url ](#http-basic-auth-info-passwd-url) below for details.

`user_name` - (Optional) username in consul (`String`).

### Admin Credentials Password

Password used to log into an admin account on the BIG-IP device.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Cbip Clusters Cbip Devices

List of Classic BIG-IP devices..

`admin_credentials` - (Required) x-required. See [Cbip Devices Admin Credentials ](#cbip-devices-admin-credentials) below for details.

`cbip_certificate_authority` - (Required) x-required. See [Cbip Devices Cbip Certificate Authority ](#cbip-devices-cbip-certificate-authority) below for details.

`cbip_mgmt_ip` - (Required) IP Address of the Classic BIG-IP device. Hostname is not supported. (`String`).

###### One of the arguments from this list "default_all, namespace_mapping" can be set

`default_all` - (Optional) If configuring in an App Namespace, discovered services across all BIG-IP partitions will be discovered in the current Namespace (`Bool`).

`namespace_mapping` - (Optional) Select the BIG-IP partitions from which services will be discovered. If configuring in Shared Configuration, services can be discovered in selected App Namespaces. If configuring in App Namespace services will be discovered in the current Namespace.. See [Namespace Mapping Choice Namespace Mapping ](#namespace-mapping-choice-namespace-mapping) below for details.

`virtual_server_filter` - (Optional) Filters to discover only required BIG-IP Virtual Servers. The Virtual Server will be discovered only if it matches all criteria specified below. A blank criteria will be treated as match all.. See [Cbip Devices Virtual Server Filter ](#cbip-devices-virtual-server-filter) below for details.

### Cbip Clusters Metadata

Common attributes for the device configuration including name and description..

`description` - (Optional) Human readable description. (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

### Cbip Devices Admin Credentials

x-required.

`password` - (Required) Password used to log into an admin account on the BIG-IP device. See [Admin Credentials Password ](#admin-credentials-password) below for details.

`username` - (Required) Username used to log into an admin account on the BIG-IP device (`String`).

### Cbip Devices Cbip Certificate Authority

x-required.

###### One of the arguments from this list "skip_server_verification, trusted_ca" must be set

`skip_server_verification` - (Optional) Skip origin server verification (`Bool`).

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Origin Pool for verification of server's certificate. See [ref](#ref) below for details.

### Cbip Devices Virtual Server Filter

Filters to discover only required BIG-IP Virtual Servers. The Virtual Server will be discovered only if it matches all criteria specified below. A blank criteria will be treated as match all..

`description_regex` - (Optional) Regex to match Virtual Server description (`String`).

`discover_disabled_virtual_servers` - (Optional) When checked, disabled virtual servers will be included (`Bool`).

`name_regex` - (Optional) Regex to match Virtual Server name (`String`).

`port_ranges` - (Optional) Maximum number of ports allowed is 1024. (`String`).

### Config Type Connection Info

Provide API server access details (endpoint and TLS parameters).

`api_server` - (Required) API server must be a fully qualified domain string and port specified as host:port pair (`String`).

`tls_info` - (Optional) TLS settings to enable transport layer security. See [Connection Info Tls Info ](#connection-info-tls-info) below for details.

### Config Type Kubeconfig Url

Provide kubeconfig file to connect to K8s cluster.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Connection Info Tls Info

TLS settings to enable transport layer security.

`certificate` - (Optional) Client certificate is PEM-encoded certificate or certificate-chain. (`String`).

`key_url` - (Optional) The data may be optionally secured using BlindFold.. See [Tls Info Key Url ](#tls-info-key-url) below for details.

`server_name` - (Optional) server is used (`String`).

`trusted_ca_url` - (Optional) Certificates in PEM format including the PEM headers. (`String`).

### Discovery Cbip Cbip Clusters

are in an Active-Active or Active-Standby setup or even a standalone BIG-IP device..

`cbip_devices` - (Required) List of Classic BIG-IP devices.. See [Cbip Clusters Cbip Devices ](#cbip-clusters-cbip-devices) below for details.

`metadata` - (Required) Common attributes for the device configuration including name and description.. See [Cbip Clusters Metadata ](#cbip-clusters-metadata) below for details.

### Discovery Choice Discovery Cbip

Discovery configuration for Classic BIG-IP.

`cbip_clusters` - (Required) are in an Active-Active or Active-Standby setup or even a standalone BIG-IP device.. See [Discovery Cbip Cbip Clusters ](#discovery-cbip-cbip-clusters) below for details.

### Discovery Choice Discovery Consul

Discovery configuration for Hashicorp Consul.

`access_info` - (Required) Credentials to access Hashicorp Consul service discovery. See [Discovery Consul Access Info ](#discovery-consul-access-info) below for details.

`publish_info` - (Required) Configuration to publish VIPs. See [Discovery Consul Publish Info ](#discovery-consul-publish-info) below for details.

### Discovery Choice Discovery K8s

Discovery configuration for K8s..

`access_info` - (Required) Credentials can be kubeconfig file or mTLS using PKI certificates. See [Discovery K8s Access Info ](#discovery-k8s-access-info) below for details.

`publish_info` - (Required) Configuration to publish VIPs. See [Discovery K8s Publish Info ](#discovery-k8s-publish-info) below for details.

### Discovery Consul Access Info

Credentials to access Hashicorp Consul service discovery.

`connection_info` - (Optional) Configuration details to access Hashicorp Consul API service using REST.. See [Access Info Connection Info ](#access-info-connection-info) below for details.

`http_basic_auth_info` - (Optional) Username and password used for HTTP/HTTPS access. See [Access Info Http Basic Auth Info ](#access-info-http-basic-auth-info) below for details.

### Discovery Consul Publish Info

Configuration to publish VIPs.

###### One of the arguments from this list "disable, publish" must be set

`disable` - (Optional) Disable VIP Publishing (`Bool`).

`publish` - (Optional) Publish domain to VIP mapping. (`Bool`).

### Discovery K8s Access Info

Credentials can be kubeconfig file or mTLS using PKI certificates.

###### One of the arguments from this list "connection_info, kubeconfig_url" must be set

`connection_info` - (Optional) Provide API server access details (endpoint and TLS parameters). See [Config Type Connection Info ](#config-type-connection-info) below for details.

`kubeconfig_url` - (Optional) Provide kubeconfig file to connect to K8s cluster. See [Config Type Kubeconfig Url ](#config-type-kubeconfig-url) below for details.

###### One of the arguments from this list "isolated, reachable" must be set

`isolated` - (Optional) discovered when Kubernetes cluster is in InCluster mode. (`Bool`).

`reachable` - (Optional) always discovers POD IP Address for configured endpoints. (`Bool`).

### Discovery K8s Publish Info

Configuration to publish VIPs.

###### One of the arguments from this list "disable, dns_delegation, publish, publish_fqdns" must be set

`disable` - (Optional) Disable VIP Publishing and DNS Delegation (`Bool`).

`dns_delegation` - (Optional) Program DNS delegation for a sub-domain in external cluster. See [Publish Choice Dns Delegation ](#publish-choice-dns-delegation) below for details.

`publish` - (Optional) Publish domain to VIP mapping.. See [Publish Choice Publish ](#publish-choice-publish) below for details.

`publish_fqdns` - (Optional) Use this option to publish domain to VIP mapping when all domains are expected to be fully qualified i.e. they include the namespace. (`Bool`).

### Http Basic Auth Info Passwd Url

F5XC Secret. URL for password, needs to be fetched from this path.

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

### Internet Vip Choice Disable Internet Vip

Do not enable advertise on external internet vip..

### Internet Vip Choice Enable Internet Vip

Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site..

### K8s Pod Network Choice Isolated

discovered when Kubernetes cluster is in InCluster mode..

### K8s Pod Network Choice Reachable

always discovers POD IP Address for configured endpoints..

### Namespace Mapping Items

Map BIG-IP partition(s) to App Namespaces.

`namespace` - (Optional) Select a namespace (`String`).

`partition_regex` - (Optional) The regex here will be used to match BIG-IP partition(s). (`String`).

### Namespace Mapping Choice Default All

If configuring in an App Namespace, discovered services across all BIG-IP partitions will be discovered in the current Namespace.

### Namespace Mapping Choice Namespace Mapping

Select the BIG-IP partitions from which services will be discovered. If configuring in Shared Configuration, services can be discovered in selected App Namespaces. If configuring in App Namespace services will be discovered in the current Namespace..

`items` - (Optional) Map BIG-IP partition(s) to App Namespaces. See [Namespace Mapping Items ](#namespace-mapping-items) below for details.

### Publish Choice Disable

Disable VIP Publishing.

### Publish Choice Dns Delegation

Program DNS delegation for a sub-domain in external cluster.

`dns_mode` - (Required) Indicates whether external K8S is running core DNS or kube DNS (`String`).

`subdomain` - (Required) The DNS subdomain for which F5XC will respond to DNS queries. (`String`).

### Publish Choice Publish

Publish domain to VIP mapping..

### Publish Choice Publish

Publish domain to VIP mapping..

`namespace` - (Required) The external K8S administrator needs to ensure that the namespace exists. (`String`).

### Publish Choice Publish Fqdns

Use this option to publish domain to VIP mapping when all domains are expected to be fully qualified i.e. they include the namespace..

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Ref Or Selector Site

Direct reference to site object.

###### One of the arguments from this list "disable_internet_vip, enable_internet_vip" must be set

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (`Bool`).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (`Bool`).

`network_type` - (Optional) The type of network on the referred site (`String`).

`ref` - (Required) A site direct reference. See [ref](#ref) below for details.

### Ref Or Selector Virtual Network

Direct reference to virtual network object.

`ref` - (Required) A virtual network direct reference. See [ref](#ref) below for details.

### Ref Or Selector Virtual Site

Direct reference to virtual site object.

###### One of the arguments from this list "disable_internet_vip, enable_internet_vip" must be set

`disable_internet_vip` - (Optional) Do not enable advertise on external internet vip. (`Bool`).

`enable_internet_vip` - (Optional) Enable advertise on internet vip. Only supported for AWS TGW Site or AWS VPC Site. (`Bool`).

`network_type` - (Optional) The type of network on the referred virtual_site (`String`).

`ref` - (Required) A virtual_site direct reference. See [ref](#ref) below for details.

### Secret Info Oneof Blindfold Secret Info

Blindfold Secret is used for the secrets managed by F5XC Secret Management Service.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Secret Info Oneof Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Server Validation Choice Skip Server Verification

Skip origin server verification.

### Tls Info Key Url

The data may be optionally secured using BlindFold..

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

Attribute Reference
-------------------

-	`id` - This is the id of the configured discovery.
