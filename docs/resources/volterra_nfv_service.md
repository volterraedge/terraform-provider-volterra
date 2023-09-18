---

page_title: "Volterra: nfv_service"

description: "The nfv_service allows CRUD of Nfv Service resource on Volterra SaaS"
-----------------------------------------------------------------------------------

Resource volterra_nfv_service
=============================

The Nfv Service allows CRUD of Nfv Service resource on Volterra SaaS

~> **Note:** Please refer to [Nfv Service API docs](https://docs.cloud.f5.com/docs/api/nfv-service) to learn more

Example Usage
-------------

```hcl
resource "volterra_nfv_service" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "https_management disable_https_management" must be set

  https_management {
    // One of the arguments from this list "advertise_on_sli_vip advertise_on_slo_vip advertise_on_slo_sli disable_local do_not_advertise_on_internet advertise_on_internet_default_vip advertise_on_internet advertise_on_slo_internet_vip" must be set

    advertise_on_slo_sli {
      // One of the arguments from this list "no_mtls use_mtls" must be set
      no_mtls = true

      tls_certificates {
        certificate_url = "value"
        description     = "Certificate used in production environment"

        // One of the arguments from this list "disable_ocsp_stapling custom_hash_algorithms use_system_defaults" must be set

        disable_ocsp_stapling {}
        private_key {
          blindfold_secret_info_internal {
            decryption_provider = "value"
            location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
            store_provider      = "value"
          }

          secret_encoding_type = "secret_encoding_type"

          // One of the arguments from this list "clear_secret_info wingman_secret_info blindfold_secret_info vault_secret_info" must be set

          blindfold_secret_info {
            decryption_provider = "value"
            location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
            store_provider      = "value"
          }
        }
      }

      tls_config {
        // One of the arguments from this list "custom_security default_security medium_security low_security" must be set
        default_security = true
      }
    }

    domain_suffix = "foo.com"

    // One of the arguments from this list "do_not_advertise advertise_on_public_default_vip advertise_on_public" must be set
    do_not_advertise = true

    // One of the arguments from this list "default_https_port https_port" must be set
    default_https_port = true
  }

  // One of the arguments from this list "f5_big_ip_aws_service palo_alto_fw_service" must be set

  f5_big_ip_aws_service {
    admin_password {
      blindfold_secret_info_internal {
        decryption_provider = "value"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "value"
      }

      secret_encoding_type = "secret_encoding_type"

      // One of the arguments from this list "blindfold_secret_info vault_secret_info clear_secret_info wingman_secret_info" must be set

      blindfold_secret_info {
        decryption_provider = "value"
        location            = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        store_provider      = "value"
      }
    }

    admin_username = "admin"

    endpoint_service {
      // One of the arguments from this list "disable_advertise_on_slo_ip advertise_on_slo_ip advertise_on_slo_ip_external" must be set
      disable_advertise_on_slo_ip = true

      // One of the arguments from this list "automatic_vip configured_vip" must be set
      automatic_vip = true

      // One of the arguments from this list "no_tcp_ports default_tcp_ports http_port https_port custom_tcp_ports" must be set

      custom_tcp_ports {
        ports = ["100-200"]
      }
      // One of the arguments from this list "no_udp_ports custom_udp_ports" must be set
      no_udp_ports = true
    }

    // One of the arguments from this list "byol_image market_place_image" must be set

    market_place_image {
      // One of the arguments from this list "AWAFPayG200Mbps AWAFPayG3Gbps" must be set
      AWAFPayG200Mbps = true
    }
    nodes {
      aws_az_name = "us-west-2a"

      // One of the arguments from this list "reserved_mgmt_subnet mgmt_subnet" must be set
      reserved_mgmt_subnet = true
      node_name            = "node1"

      // One of the arguments from this list "automatic_prefix tunnel_prefix" must be set
      automatic_prefix = true
    }

    // One of the arguments from this list "aws_vpc_site_params aws_tgw_site_params" must be set

    aws_tgw_site_params {
      aws_tgw_site {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
    }
    ssh_key = "ssh-rsa AAAAB..."
    tags = {
      "key1" = "value1"
    }
  }
  // One of the arguments from this list "disable_ssh_access enabled_ssh_access" must be set
  disable_ssh_access = true
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

`disable_https_management` - (Optional) HTTPS based management is not enabled (bool).

`https_management` - (Optional) Enable HTTPS based management. See [Https Management ](#https-management) below for details.

`f5_big_ip_aws_service` - (Optional) Virtual F5 BigIP service to be deployed on AWS. See [F5 Big Ip Aws Service ](#f5-big-ip-aws-service) below for details.

`palo_alto_fw_service` - (Optional) Palo Alto Networks VM-Series Firewall to be deployed on AWS Cloud. See [Palo Alto Fw Service ](#palo-alto-fw-service) below for details.

`disable_ssh_access` - (Optional) SSH based access is disabled (bool).

`enabled_ssh_access` - (Optional) Enable SSH access to nodes. See [Enabled Ssh Access ](#enabled-ssh-access) below for details.

### AWAFPayG200Mbps

F5 Advanced WAF with LTM, IPI, and Threat Campaigns (PAYG, 200Mbps).

### AWAFPayG3Gbps

F5 Advanced WAF with LTM, IPI, and Threat Campaigns (PAYG, 3Gbps).

### Admin Password

Secret admin password for BIG ip.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Advertise On Internet

Advertise this loadbalancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Advertise On Internet Default Vip

Enable management access on internet with default VIP.

### Advertise On Public

Advertise this loadbalancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Advertise On Public Default Vip

Enable management access on internet with default VIP.

### Advertise On Sli Vip

Enable on Site local inside network, default VIP will be used.

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Config ](#tls-config) below for details.

### Advertise On Slo Internet Vip

Enable On Site Local Outside Internet VIP.

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Config ](#tls-config) below for details.

### Advertise On Slo Ip

Advertise this loadbalancer on Site Local Outside network address.

### Advertise On Slo Ip External

Advertise this loadbalancer on Site Local Outside network address and enable cloud external IP.

### Advertise On Slo Sli

Enable on Site local inside and outside network, default VIP will be used.

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Config ](#tls-config) below for details.

### Advertise On Slo Vip

Enable on Site local outside network, default VIP will be used.

`no_mtls` - (Optional) x-displayName: "Disable" (bool).

`use_mtls` - (Optional) x-displayName: "Enable". See [Use Mtls ](#use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Tls Certificates ](#tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Tls Config ](#tls-config) below for details.

### Authorization Key

Authentication key for Panorama.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Auto Setup

Auto Setup API Access & Users. With this firewall api access and given admin user will be auto setup.

`admin_password` - (Required) Firewall Admin Password. See [Admin Password ](#admin-password) below for details.

`admin_username` - (Required) Firewall Admin Username (`String`).

`autogenerated_ssh_keys` - (Optional) Autogenerated SSH Keys, users will be able to download the keys after external service is created (bool).

`manual_ssh_keys` - (Optional) User given public and private SSH keys. See [Manual Ssh Keys ](#manual-ssh-keys) below for details.

### Autogenerated Ssh Keys

Autogenerated SSH Keys, users will be able to download the keys after external service is created.

### Automatic Prefix

System will automatically select tunnel prefix.

### Automatic Vip

System will automatically select a VIP.

### Aws Tgw Site Params

Select AWS transit gateway site.

`aws_tgw_site` - (Required) Reference to AWS transit gateway site. See [ref](#ref) below for details.

### Aws Vpc Site Params

Select AWS VPC site.

`aws_vpc_site` - (Required) Reference to AWS VPC site. See [ref](#ref) below for details.

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

### Byol Image

Select the BigIp bring your own license image to be used for this service.

`image` - (Required) Select the BigIp pay as you go image to be used for this service (`String`).

`license` - (Optional) Secret License data. See [License ](#license) below for details.

### Clear Secret Info

Clear Secret is used for the secrets that are not encrypted.

`provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

`url` - (Required) When asked for this secret, caller will get Secret bytes after Base64 decoding. (`String`).

### Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Custom Tcp Ports

select custom tcp ports.

`ports` - (Required) List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192 (`String`).

### Custom Udp Ports

select custom udp ports.

`ports` - (Required) List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192 (`String`).

### Default Https Port

Select default HTTPS 443.

### Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Default Tcp Ports

Select default TCP ports, 80 and 443.

### Disable Advertise On Slo Ip

Do not Advertise this loadbalancer on Site Local Outside network address.

### Disable Local

Disable on Site local network.

### Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Disable Panaroma

Disable Panorama connection during bootstrap, user can always enable it using firewall management console.

### Do Not Advertise

Do not enable access to management from internet.

### Do Not Advertise On Internet

Do not enable access to management from internet.

### Enabled Ssh Access

Enable SSH access to nodes.

`advertise_on_public` - (Optional) Advertise this loadbalancer on public network. See [Advertise On Public ](#advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Enable management access on internet with default VIP (bool).

`advertise_on_sli_vip` - (Optional) Enable on Site local inside network, default VIP will be used. See [Advertise On Sli Vip ](#advertise-on-sli-vip) below for details.

`advertise_on_slo_internet_vip` - (Optional) Enable On Site Local Outside Internet VIP. See [Advertise On Slo Internet Vip ](#advertise-on-slo-internet-vip) below for details.

`advertise_on_slo_sli` - (Optional) Enable on Site local inside and outside network, default VIP will be used. See [Advertise On Slo Sli ](#advertise-on-slo-sli) below for details.

`advertise_on_slo_vip` - (Optional) Enable on Site local outside network, default VIP will be used. See [Advertise On Slo Vip ](#advertise-on-slo-vip) below for details.

`ssh_ports` - (Required) Enter TCP port per node (`Int`).

### Endpoint Service

External service type is Endpoint service.

`advertise_on_slo_ip` - (Optional) Advertise this loadbalancer on Site Local Outside network address (bool).

`advertise_on_slo_ip_external` - (Optional) Advertise this loadbalancer on Site Local Outside network address and enable cloud external IP (bool).

`disable_advertise_on_slo_ip` - (Optional) Do not Advertise this loadbalancer on Site Local Outside network address (bool).

`automatic_vip` - (Optional) System will automatically select a VIP (bool).

`configured_vip` - (Required) Enter IP address for the default VIP (`String`).

`custom_tcp_ports` - (Optional) select custom tcp ports. See [Custom Tcp Ports ](#custom-tcp-ports) below for details.

`default_tcp_ports` - (Optional) Select default TCP ports, 80 and 443 (bool).

`http_port` - (Optional) Select HTTP Port 80 (bool).

`https_port` - (Optional) Select HTTPS Port 443 (bool).

`no_tcp_ports` - (Optional) do not select tcp ports (bool).

`custom_udp_ports` - (Optional) select custom udp ports. See [Custom Udp Ports ](#custom-udp-ports) below for details.

`no_udp_ports` - (Optional) do not select udp ports (bool).

### F5 Big Ip Aws Service

Virtual F5 BigIP service to be deployed on AWS.

`admin_password` - (Required) Secret admin password for BIG ip. See [Admin Password ](#admin-password) below for details.

`admin_username` - (Required) Admin Username for BIG-IP (`String`).

`endpoint_service` - (Optional) External service type is Endpoint service. See [Endpoint Service ](#endpoint-service) below for details.

`byol_image` - (Required) Select the BigIp bring your own license image to be used for this service. See [Byol Image ](#byol-image) below for details.

`market_place_image` - (Required) Select the BigIp pay as you go image to be used for this service. See [Market Place Image ](#market-place-image) below for details.

`nodes` - (Required) Specify how and where the service nodes are spawned. See [Nodes ](#nodes) below for details.

`aws_tgw_site_params` - (Optional) Select AWS transit gateway site. See [Aws Tgw Site Params ](#aws-tgw-site-params) below for details.

`aws_vpc_site_params` - (Optional) Select AWS VPC site. See [Aws Vpc Site Params ](#aws-vpc-site-params) below for details.

`ssh_key` - (Required) Public SSH key for accessing the Big IP nodes. (`String`).

`tags` - (Optional) It helps to manage, identify, organize, search for, and filter resources in AWS console. (`String`).

### Http Port

Select HTTP Port 80.

### Https Management

Enable HTTPS based management.

`advertise_on_internet` - (Optional) Advertise this loadbalancer on public network. See [Advertise On Internet ](#advertise-on-internet) below for details.

`advertise_on_internet_default_vip` - (Optional) Enable management access on internet with default VIP (bool).

`advertise_on_sli_vip` - (Optional) Enable on Site local inside network, default VIP will be used. See [Advertise On Sli Vip ](#advertise-on-sli-vip) below for details.

`advertise_on_slo_internet_vip` - (Optional) Enable On Site Local Outside Internet VIP. See [Advertise On Slo Internet Vip ](#advertise-on-slo-internet-vip) below for details.

`advertise_on_slo_sli` - (Optional) Enable on Site local inside and outside network, default VIP will be used. See [Advertise On Slo Sli ](#advertise-on-slo-sli) below for details.

`advertise_on_slo_vip` - (Optional) Enable on Site local outside network, default VIP will be used. See [Advertise On Slo Vip ](#advertise-on-slo-vip) below for details.

`disable_local` - (Optional) Disable on Site local network (bool).

`do_not_advertise_on_internet` - (Optional) Do not enable access to management from internet (bool).

`domain_suffix` - (Required) Domain suffix will be used along with node name to form URL to access node management (`String`).

`advertise_on_public` - (Optional) Advertise this loadbalancer on public network. See [Advertise On Public ](#advertise-on-public) below for details.

`advertise_on_public_default_vip` - (Optional) Enable management access on internet with default VIP (bool).

`do_not_advertise` - (Optional) Do not enable access to management from internet (bool).

`default_https_port` - (Optional) Select default HTTPS 443 (bool).

`https_port` - (Optional) Enter TCP port number (`Int`).

### Https Port

Select HTTPS Port 443.

### License

Secret License data.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Manual Ssh Keys

User given public and private SSH keys.

`private_key` - (Required) Authorized Public SSH key which will be programmed on the node. See [Private Key ](#private-key) below for details.

`public_key` - (Required) Authorized Public SSH key which will be programmed on the node (`String`).

### Market Place Image

Select the BigIp pay as you go image to be used for this service.

`AWAFPayG200Mbps` - (Required) F5 Advanced WAF with LTM, IPI, and Threat Campaigns (PAYG, 200Mbps) (bool).

`AWAFPayG3Gbps` - (Required) F5 Advanced WAF with LTM, IPI, and Threat Campaigns (PAYG, 3Gbps) (bool).

### Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Mgmt Subnet

Select Existing Subnet or Create New.

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Subnet Param ](#subnet-param) below for details.

### No Crl

Client certificate revocation status is not verified.

### No Mtls

x-displayName: "Disable".

### No Tcp Ports

do not select tcp ports.

### No Udp Ports

do not select udp ports.

### Nodes

Specify how and where the service nodes are spawned.

`aws_az_name` - (Required) AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is one of the AZ for sites (`String`).

`mgmt_subnet` - (Optional) Select Existing Subnet or Create New. See [Mgmt Subnet ](#mgmt-subnet) below for details.

`reserved_mgmt_subnet` - (Optional) Autogenerate and reserve a subnet from the Primary CIDR (bool).

`node_name` - (Required) Node Name will be used to assign as hostname to the service (`String`).

`automatic_prefix` - (Optional) System will automatically select tunnel prefix (bool).

`tunnel_prefix` - (Optional) Enter IP prefix for the tunnel, it has to be /30 (`String`).

### Palo Alto Fw Service

Palo Alto Networks VM-Series Firewall to be deployed on AWS Cloud.

`pan_ami_bundle1` - (Required) VM-Series Next-Generation Firewall Bundle 1 (bool).

`pan_ami_bundle2` - (Required) VM-Series Next-Generation Firewall Bundle 2 (bool).

`aws_tgw_site` - (Required) Select AWS transit gateway site. See [ref](#ref) below for details.

`instance_type` - (Required) AWS Instance type (`String`).

`disable_panaroma` - (Optional) Disable Panorama connection during bootstrap, user can always enable it using firewall management console (bool).

`panorama_server` - (Optional) Enabled adding firewall instances to Panorama for config management during bootstrap. See [Panorama Server ](#panorama-server) below for details.

`service_nodes` - (Required) Specify how and where the Palo Alto Networks Vm-Series AZ Nodes are spawned. See [Service Nodes ](#service-nodes) below for details.

`auto_setup` - (Optional) Auto Setup API Access & Users. With this firewall api access and given admin user will be auto setup. See [Auto Setup ](#auto-setup) below for details.

`ssh_key` - (Required) its corresponding ssh private key. (`String`).

`tags` - (Optional) It helps to manage, identify, organize, search for, and filter resources in AWS console. (`String`).

`version` - (Optional) PAN-OS version (`String`).

### Pan Ami Bundle1

VM-Series Next-Generation Firewall Bundle 1.

### Pan Ami Bundle2

VM-Series Next-Generation Firewall Bundle 2.

### Panorama Server

Enabled adding firewall instances to Panorama for config management during bootstrap.

`authorization_key` - (Required) Authentication key for Panorama. See [Authorization Key ](#authorization-key) below for details.

`device_group_name` - (Optional) Device Group Name (`String`).

`server` - (Required) Panorama Server Address to which the firewall should connect to (`String`).

`template_stack_name` - (Optional) Template Stack Name (`String`).

### Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Blindfold Secret Info Internal ](#blindfold-secret-info-internal) below for details.

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Blindfold Secret Info ](#blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Clear Secret Info ](#clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Vault Secret Info ](#vault-secret-info) below for details.

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Wingman Secret Info ](#wingman-secret-info) below for details.

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

### Reserved Mgmt Subnet

Autogenerate and reserve a subnet from the Primary CIDR.

### Service Nodes

Specify how and where the Palo Alto Networks Vm-Series AZ Nodes are spawned.

`nodes` - (Required) x-required. See [Nodes ](#nodes) below for details.

### Subnet Param

Parameters for creating new subnet.

`ipv4` - (Required) IPv4 subnet prefix for this subnet (`String`).

`ipv6` - (Optional) IPv6 subnet prefix for this subnet (`String`).

### Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Custom Hash Algorithms ](#custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Disable Ocsp Stapling ](#disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Use System Defaults ](#use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Private Key ](#private-key) below for details.

### Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Custom Security ](#custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (bool).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (bool).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (bool).

### Use Mtls

x-displayName: "Enable".

`crl` - (Optional) Specify the CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

`no_crl` - (Optional) Client certificate revocation status is not verified (bool).

`trusted_ca_url` - (Required) The URL for a trust store (`String`).

### Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

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

-	`id` - This is the id of the configured nfv_service.
