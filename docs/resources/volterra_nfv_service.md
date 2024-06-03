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

  // One of the arguments from this list "disable_https_management https_management" must be set

  disable_https_management = true

  // One of the arguments from this list "f5_big_ip_aws_service palo_alto_fw_service" must be set

  f5_big_ip_aws_service {
    admin_password {
      blindfold_secret_info_internal {
        decryption_provider = "value"

        location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

        store_provider = "value"
      }

      secret_encoding_type = "secret_encoding_type"

      // One of the arguments from this list "wingman_secret_info blindfold_secret_info vault_secret_info clear_secret_info" must be set

      blindfold_secret_info {
        decryption_provider = "value"

        location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

        store_provider = "value"
      }
    }

    admin_username = "admin"

    endpoint_service {
      // One of the arguments from this list "advertise_on_slo_ip_external disable_advertise_on_slo_ip advertise_on_slo_ip" must be set

      disable_advertise_on_slo_ip = true

      // One of the arguments from this list "automatic_vip configured_vip" must be set

      automatic_vip = true

      // One of the arguments from this list "no_tcp_ports default_tcp_ports http_port https_port custom_tcp_ports" must be set

      https_port = true

      // One of the arguments from this list "no_udp_ports custom_udp_ports" must be set

      no_udp_ports = true
    }

    // One of the arguments from this list "market_place_image byol_image" must be set

    market_place_image {
      // One of the arguments from this list "best_plus_payg_1gbps AWAFPayG200Mbps AWAFPayG3Gbps BestPlusPayG200Mbps" must be set

      best_plus_payg_1gbps = true
    }
    nodes {
      aws_az_name = "us-west-2a"

      // One of the arguments from this list "reserved_mgmt_subnet mgmt_subnet" must be set

      reserved_mgmt_subnet = true
      node_name = "node1"

      // One of the arguments from this list "automatic_prefix tunnel_prefix" must be set

      automatic_prefix = true
    }

    // One of the arguments from this list "aws_tgw_site_params aws_vpc_site_params" must be set

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

`disable_https_management` - (Optional) HTTPS based management is not enabled (`Bool`).

`https_management` - (Optional) Enable HTTPS based management. See [Http Management Choice Https Management ](#http-management-choice-https-management) below for details.

`f5_big_ip_aws_service` - (Optional) Virtual BIG-IP service to be deployed on AWS. See [Service Provider Choice F5 Big Ip Aws Service ](#service-provider-choice-f5-big-ip-aws-service) below for details.

`palo_alto_fw_service` - (Optional) Palo Alto Networks VM-Series Firewall to be deployed on AWS Cloud. See [Service Provider Choice Palo Alto Fw Service ](#service-provider-choice-palo-alto-fw-service) below for details.

`disable_ssh_access` - (Optional) SSH based access is disabled (`Bool`).

`enabled_ssh_access` - (Optional) Enable SSH access to nodes. See [Ssh Management Choice Enabled Ssh Access ](#ssh-management-choice-enabled-ssh-access) below for details.

### Admin Password Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Advertise Choice Advertise On Internet

Advertise this loadbalancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Advertise Choice Advertise On Internet Default Vip

Enable management access on internet with default VIP.

### Advertise Choice Advertise On Public

Advertise this loadbalancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Advertise Choice Advertise On Public Default Vip

Enable management access on internet with default VIP.

### Advertise Choice Advertise On Sli

Enable on Site local inside network, default VIP will be used.

### Advertise Choice Advertise On Sli Vip

Enable on Site local inside network, default VIP will be used.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Advertise On Sli Vip Tls Certificates ](#advertise-on-sli-vip-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Advertise On Sli Vip Tls Config ](#advertise-on-sli-vip-tls-config) below for details.

### Advertise Choice Advertise On Slo

Enable on Site local outside network, default VIP will be used.

### Advertise Choice Advertise On Slo Internet Vip

Enable On Site Local Outside Internet VIP.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Advertise On Slo Internet Vip Tls Certificates ](#advertise-on-slo-internet-vip-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Advertise On Slo Internet Vip Tls Config ](#advertise-on-slo-internet-vip-tls-config) below for details.

### Advertise Choice Advertise On Slo Sli

Enable on Site local inside and outside network, default VIP will be used.

###### One of the arguments from this list "use_mtls, no_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Advertise On Slo Sli Tls Certificates ](#advertise-on-slo-sli-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Advertise On Slo Sli Tls Config ](#advertise-on-slo-sli-tls-config) below for details.

### Advertise Choice Advertise On Slo Sli

Enable on Site local inside and outside network, default VIP will be used.

### Advertise Choice Advertise On Slo Vip

Enable on Site local outside network, default VIP will be used.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Advertise On Slo Vip Tls Certificates ](#advertise-on-slo-vip-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Advertise On Slo Vip Tls Config ](#advertise-on-slo-vip-tls-config) below for details.

### Advertise Choice Disable Local

Disable on Site local network.

### Advertise Choice Do Not Advertise On Internet

Do not enable access to management from internet.

### Advertise On Sli Vip Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, use_system_defaults, disable_ocsp_stapling" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Advertise On Sli Vip Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "default_security, medium_security, low_security, custom_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Advertise On Slo Internet Vip Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "use_system_defaults, disable_ocsp_stapling, custom_hash_algorithms" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Advertise On Slo Internet Vip Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "default_security, medium_security, low_security, custom_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Advertise On Slo Sli Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "use_system_defaults, disable_ocsp_stapling, custom_hash_algorithms" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Advertise On Slo Sli Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "default_security, medium_security, low_security, custom_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Advertise On Slo Vip Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "use_system_defaults, disable_ocsp_stapling, custom_hash_algorithms" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Advertise On Slo Vip Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "default_security, medium_security, low_security, custom_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Ami Choice AWAFPayG200Mbps

F5 Advanced WAF with LTM, IPI, and Threat Campaigns (PAYG, 200Mbps).

### Ami Choice AWAFPayG3Gbps

F5 Advanced WAF with LTM, IPI, and Threat Campaigns (PAYG, 3Gbps).

### Ami Choice Pan Ami Bundle1

VM-Series Next-Generation Firewall Bundle 1.

### Ami Choice Pan Ami Bundle2

VM-Series Next-Generation Firewall Bundle 2.

### Authorization Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Auto Setup Admin Password

Firewall Admin Password.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Admin Password Blindfold Secret Info Internal ](#admin-password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Byol Image License

Secret License data.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [License Blindfold Secret Info Internal ](#license-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "wingman_secret_info, blindfold_secret_info, vault_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Choice Custom Security

Custom selection of TLS versions and cipher suites.

`cipher_suites` - (Required) The TLS listener will only support the specified cipher list. (`String`).

`max_version` - (Optional) Maximum TLS protocol version. (`String`).

`min_version` - (Optional) Minimum TLS protocol version. (`String`).

### Choice Default Security

TLS v1.2+ with PFS ciphers and strong crypto algorithms..

### Choice Low Security

TLS v1.0+ including non-PFS ciphers and weak crypto algorithms..

### Choice Medium Security

TLS v1.0+ with PFS ciphers and medium strength crypto algorithms..

### Choice Subnet Param

Parameters for creating new subnet.

`ipv4` - (Required) IPv4 subnet prefix for this subnet (`String`).

`ipv6` - (Optional) IPv6 subnet prefix for this subnet (`String`).

### Crl Choice No Crl

Client certificate revocation status is not verified.

### Enabled Ssh Access Node Ssh Ports

Enter TCP port and node name per node.

`node_name` - (Required) Node name will be used to match a particular node with the desired TCP port (`String`).

`ssh_port` - (Required) Enter TCP port per node (`Int`).

### External Vip Choice Advertise On Slo Ip

Advertise this loadbalancer on Site Local Outside network address.

### External Vip Choice Advertise On Slo Ip External

Advertise this loadbalancer on Site Local Outside network address and enable cloud external IP.

### External Vip Choice Disable Advertise On Slo Ip

Do not Advertise this loadbalancer on Site Local Outside network address.

### F5 Big Ip Aws Service Admin Password

Secret admin password for BIG ip.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Admin Password Blindfold Secret Info Internal ](#admin-password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "wingman_secret_info, blindfold_secret_info, vault_secret_info, clear_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### F5 Big Ip Aws Service Endpoint Service

External service type is Endpoint service.

###### One of the arguments from this list "disable_advertise_on_slo_ip, advertise_on_slo_ip, advertise_on_slo_ip_external" must be set

`advertise_on_slo_ip` - (Optional) Advertise this loadbalancer on Site Local Outside network address (`Bool`).

`advertise_on_slo_ip_external` - (Optional) Advertise this loadbalancer on Site Local Outside network address and enable cloud external IP (`Bool`).

`disable_advertise_on_slo_ip` - (Optional) Do not Advertise this loadbalancer on Site Local Outside network address (`Bool`).

###### One of the arguments from this list "automatic_vip, configured_vip" must be set

`automatic_vip` - (Optional) System will automatically select a VIP (`Bool`).

`configured_vip` - (Optional) Enter IP address for the default VIP (`String`).

###### One of the arguments from this list "default_tcp_ports, http_port, https_port, custom_tcp_ports, no_tcp_ports" must be set

`custom_tcp_ports` - (Optional) select custom tcp ports. See [Tcp Port Choice Custom Tcp Ports ](#tcp-port-choice-custom-tcp-ports) below for details.

`default_tcp_ports` - (Optional) Select default TCP ports, 80 and 443 (`Bool`).

`http_port` - (Optional) Select HTTP Port 80 (`Bool`).

`https_port` - (Optional) Select HTTPS Port 443 (`Bool`).

`no_tcp_ports` - (Optional) do not select tcp ports (`Bool`).

###### One of the arguments from this list "no_udp_ports, custom_udp_ports" must be set

`custom_udp_ports` - (Optional) select custom udp ports. See [Udp Port Choice Custom Udp Ports ](#udp-port-choice-custom-udp-ports) below for details.

`no_udp_ports` - (Optional) do not select udp ports (`Bool`).

### F5 Big Ip Aws Service Nodes

Specify how and where the service nodes are spawned.

`aws_az_name` - (Required) The AWS Availability Zone must be consistent with the AWS Region chosen. Please select an AZ in the same Region as your TGW Site (`String`).

###### One of the arguments from this list "reserved_mgmt_subnet, mgmt_subnet" must be set

`mgmt_subnet` - (Optional) Select Existing Subnet or Create New. See [Mgmt Subnet Choice Mgmt Subnet ](#mgmt-subnet-choice-mgmt-subnet) below for details.

`reserved_mgmt_subnet` - (Optional) Autogenerate and reserve a subnet from the Primary CIDR (`Bool`).

`node_name` - (Required) Node Name will be used to assign as hostname to the service (`String`).

###### One of the arguments from this list "automatic_prefix, tunnel_prefix" must be set

`automatic_prefix` - (Optional) System will automatically select tunnel prefix (`Bool`).

`tunnel_prefix` - (Optional) Enter IP prefix for the tunnel, it has to be /30 (`String`).

### Http Management Choice Https Management

Enable HTTPS based management.

###### One of the arguments from this list "advertise_on_slo_internet_vip, advertise_on_sli_vip, advertise_on_slo_vip, advertise_on_slo_sli, disable_local, do_not_advertise_on_internet, advertise_on_internet_default_vip, advertise_on_internet" must be set

`advertise_on_internet` - (Optional) Advertise this loadbalancer on public network. See [Advertise Choice Advertise On Internet ](#advertise-choice-advertise-on-internet) below for details.

`advertise_on_internet_default_vip` - (Optional) Enable management access on internet with default VIP (`Bool`).

`advertise_on_sli_vip` - (Optional) Enable on Site local inside network, default VIP will be used. See [Advertise Choice Advertise On Sli Vip ](#advertise-choice-advertise-on-sli-vip) below for details.

`advertise_on_slo_internet_vip` - (Optional) Enable On Site Local Outside Internet VIP. See [Advertise Choice Advertise On Slo Internet Vip ](#advertise-choice-advertise-on-slo-internet-vip) below for details.

`advertise_on_slo_sli` - (Optional) Enable on Site local inside and outside network, default VIP will be used. See [Advertise Choice Advertise On Slo Sli ](#advertise-choice-advertise-on-slo-sli) below for details.

`advertise_on_slo_vip` - (Optional) Enable on Site local outside network, default VIP will be used. See [Advertise Choice Advertise On Slo Vip ](#advertise-choice-advertise-on-slo-vip) below for details.

`disable_local` - (Optional) Disable on Site local network (`Bool`).(Deprecated)

`do_not_advertise_on_internet` - (Optional) Do not enable access to management from internet (`Bool`).(Deprecated)

`domain_suffix` - (Required) Domain suffix will be used along with node name to form URL to access node management (`String`).

###### One of the arguments from this list "do_not_advertise, advertise_on_public_default_vip, advertise_on_public" can be set

`advertise_on_public` - (Optional) Advertise this loadbalancer on public network. See [Internet Choice Advertise On Public ](#internet-choice-advertise-on-public) below for details.(Deprecated)

`advertise_on_public_default_vip` - (Optional) Enable management access on internet with default VIP (`Bool`).(Deprecated)

`do_not_advertise` - (Optional) Do not enable access to management from internet (`Bool`).(Deprecated)

###### One of the arguments from this list "default_https_port, https_port" must be set

`default_https_port` - (Optional) Select default HTTPS 443 (`Bool`).

`https_port` - (Optional) Enter TCP port number (`Int`).

### Image Choice Byol Image

Select the BIG-IP bring your own license image to be used for this service.

`image` - (Required) Select the BIG-IP pay as you go image to be used for this service (`String`).

`license` - (Optional) Secret License data. See [Byol Image License ](#byol-image-license) below for details.

### Image Choice Market Place Image

Select the BIG-IP pay as you go image to be used for this service.

###### One of the arguments from this list "AWAFPayG200Mbps, AWAFPayG3Gbps, BestPlusPayG200Mbps, best_plus_payg_1gbps" must be set

`AWAFPayG200Mbps` - (Optional) F5 Advanced WAF with LTM, IPI, and Threat Campaigns (PAYG, 200Mbps) (`Bool`).

`AWAFPayG3Gbps` - (Optional) F5 Advanced WAF with LTM, IPI, and Threat Campaigns (PAYG, 3Gbps) (`Bool`).

### Inside Vip Choice Automatic Vip

System will automatically select a VIP.

### Internet Choice Advertise On Public

Advertise this loadbalancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Internet Choice Advertise On Public Default Vip

Enable management access on internet with default VIP.

### Internet Choice Do Not Advertise

Do not enable access to management from internet.

### License Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Manual Ssh Keys Private Key

Authorized Public SSH key which will be programmed on the node.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "clear_secret_info, wingman_secret_info, blindfold_secret_info, vault_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Mgmt Subnet Choice Mgmt Subnet

Select Existing Subnet or Create New.

###### One of the arguments from this list "subnet_param, existing_subnet_id" must be set

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Choice Subnet Param ](#choice-subnet-param) below for details.

### Mgmt Subnet Choice Reserved Mgmt Subnet

Autogenerate and reserve a subnet from the Primary CIDR.

### Mtls Choice No Mtls

x-displayName: "Disable".

### Mtls Choice Use Mtls

x-displayName: "Enable".

###### One of the arguments from this list "no_crl, crl" can be set

`crl` - (Optional) Specify the CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

`no_crl` - (Optional) Client certificate revocation status is not verified (`Bool`).

###### One of the arguments from this list "trusted_ca_url, trusted_ca" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Load Balancer. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Load Balancer (`String`).

###### One of the arguments from this list "xfcc_disabled, xfcc_options" can be set

`xfcc_disabled` - (Optional) No X-Forwarded-Client-Cert header will be added (`Bool`).

`xfcc_options` - (Optional) X-Forwarded-Client-Cert header will be added with the configured fields. See [Xfcc Header Xfcc Options ](#xfcc-header-xfcc-options) below for details.

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Palo Alto Fw Service Service Nodes

Specify how and where the Palo Alto Networks Vm-Series AZ Nodes are spawned.

`nodes` - (Required) x-required. See [Service Nodes Nodes ](#service-nodes-nodes) below for details.

### Panaroma Connection Disable Panaroma

Disable Panorama connection during bootstrap, user can always enable it using firewall management console.

### Panaroma Connection Panorama Server

Enabled adding firewall instances to Panorama for config management during bootstrap.

`authorization_key` - (Required) Authentication key for Panorama. See [Panorama Server Authorization Key ](#panorama-server-authorization-key) below for details.

`device_group_name` - (Optional) Device Group Name (`String`).

`server` - (Required) Panorama Server Address to which the firewall should connect to (`String`).

`template_stack_name` - (Optional) Template Stack Name (`String`).

### Panorama Server Authorization Key

Authentication key for Panorama.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Authorization Key Blindfold Secret Info Internal ](#authorization-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Port Choice Default Https Port

Select default HTTPS 443.

### Private Key Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

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

### Service Nodes Nodes

x-required.

`aws_az_name` - (Required) AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is one of the AZ for sites (`String`).

###### One of the arguments from this list "reserved_mgmt_subnet, mgmt_subnet" must be set

`mgmt_subnet` - (Optional) Select Existing Subnet or Create New. See [Mgmt Subnet Choice Mgmt Subnet ](#mgmt-subnet-choice-mgmt-subnet) below for details.

`reserved_mgmt_subnet` - (Optional) Autogenerate and reserve a subnet from the Primary CIDR (`Bool`).

`node_name` - (Required) Node Name will be used to assign as hostname to the service (`String`).

### Service Provider Choice F5 Big Ip Aws Service

Virtual BIG-IP service to be deployed on AWS.

`admin_password` - (Required) Secret admin password for BIG ip. See [F5 Big Ip Aws Service Admin Password ](#f5-big-ip-aws-service-admin-password) below for details.

`admin_username` - (Required) Admin Username for BIG-IP (`String`).

`endpoint_service` - (Optional) External service type is Endpoint service. See [F5 Big Ip Aws Service Endpoint Service ](#f5-big-ip-aws-service-endpoint-service) below for details.

###### One of the arguments from this list "market_place_image, byol_image" must be set

`byol_image` - (Optional) Select the BIG-IP bring your own license image to be used for this service. See [Image Choice Byol Image ](#image-choice-byol-image) below for details.(Deprecated)

`market_place_image` - (Optional) Select the BIG-IP pay as you go image to be used for this service. See [Image Choice Market Place Image ](#image-choice-market-place-image) below for details.

`nodes` - (Required) Specify how and where the service nodes are spawned. See [F5 Big Ip Aws Service Nodes ](#f5-big-ip-aws-service-nodes) below for details.

###### One of the arguments from this list "aws_tgw_site_params, aws_vpc_site_params" must be set

`aws_tgw_site_params` - (Optional) Select AWS transit gateway site. See [Site Type Choice Aws Tgw Site Params ](#site-type-choice-aws-tgw-site-params) below for details.

`aws_vpc_site_params` - (Optional) Select AWS VPC site. See [Site Type Choice Aws Vpc Site Params ](#site-type-choice-aws-vpc-site-params) below for details.(Deprecated)

`ssh_key` - (Required) Public SSH key for accessing the Big IP nodes. (`String`).

`tags` - (Optional) It helps to manage, identify, organize, search for, and filter resources in AWS console. (`String`).

### Service Provider Choice Palo Alto Fw Service

Palo Alto Networks VM-Series Firewall to be deployed on AWS Cloud.

###### One of the arguments from this list "pan_ami_bundle1, pan_ami_bundle2" must be set

`pan_ami_bundle1` - (Optional) VM-Series Next-Generation Firewall Bundle 1 (`Bool`).

`pan_ami_bundle2` - (Optional) VM-Series Next-Generation Firewall Bundle 2 (`Bool`).

`aws_tgw_site` - (Required) Select AWS transit gateway site. See [ref](#ref) below for details.

`instance_type` - (Required) AWS Instance type (`String`).

###### One of the arguments from this list "disable_panaroma, panorama_server" can be set

`disable_panaroma` - (Optional) Disable Panorama connection during bootstrap, user can always enable it using firewall management console (`Bool`).

`panorama_server` - (Optional) Enabled adding firewall instances to Panorama for config management during bootstrap. See [Panaroma Connection Panorama Server ](#panaroma-connection-panorama-server) below for details.

`service_nodes` - (Required) Specify how and where the Palo Alto Networks Vm-Series AZ Nodes are spawned. See [Palo Alto Fw Service Service Nodes ](#palo-alto-fw-service-service-nodes) below for details.

###### One of the arguments from this list "auto_setup, ssh_key" must be set

`auto_setup` - (Optional) Auto Setup API Access & Users. With this firewall api access and given admin user will be auto setup. See [Setup Options Auto Setup ](#setup-options-auto-setup) below for details.

`ssh_key` - (Optional) its corresponding ssh private key. (`String`).

`tags` - (Optional) It helps to manage, identify, organize, search for, and filter resources in AWS console. (`String`).

`version` - (Optional) PAN-OS version (`String`).

### Setup Options Auto Setup

Auto Setup API Access & Users. With this firewall api access and given admin user will be auto setup.

`admin_password` - (Required) Firewall Admin Password. See [Auto Setup Admin Password ](#auto-setup-admin-password) below for details.

`admin_username` - (Required) Firewall Admin Username (`String`).

###### One of the arguments from this list "autogenerated_ssh_keys, manual_ssh_keys" must be set

`autogenerated_ssh_keys` - (Optional) Autogenerated SSH Keys, users will be able to download the keys after external service is created (`Bool`).(Deprecated)

`manual_ssh_keys` - (Optional) User given public and private SSH keys. See [Ssh Keys Choice Manual Ssh Keys ](#ssh-keys-choice-manual-ssh-keys) below for details.

### Site Type Choice Aws Tgw Site Params

Select AWS transit gateway site.

`aws_tgw_site` - (Required) Reference to AWS transit gateway site. See [ref](#ref) below for details.

### Site Type Choice Aws Vpc Site Params

Select AWS VPC site.

`aws_vpc_site` - (Required) Reference to AWS VPC site. See [ref](#ref) below for details.

### Ssh Keys Choice Autogenerated Ssh Keys

Autogenerated SSH Keys, users will be able to download the keys after external service is created.

### Ssh Keys Choice Manual Ssh Keys

User given public and private SSH keys.

`private_key` - (Required) Authorized Public SSH key which will be programmed on the node. See [Manual Ssh Keys Private Key ](#manual-ssh-keys-private-key) below for details.

`public_key` - (Required) Authorized Public SSH key which will be programmed on the node (`String`).

### Ssh Management Choice Enabled Ssh Access

Enable SSH access to nodes.

###### One of the arguments from this list "advertise_on_sli, advertise_on_slo, advertise_on_slo_sli, advertise_on_public_default_vip, advertise_on_public" can be set

`advertise_on_public` - (Optional) Advertise this loadbalancer on public network. See [Advertise Choice Advertise On Public ](#advertise-choice-advertise-on-public) below for details.(Deprecated)

`advertise_on_public_default_vip` - (Optional) Enable management access on internet with default VIP (`Bool`).(Deprecated)

`advertise_on_sli` - (Optional) Enable on Site local inside network, default VIP will be used (`Bool`).

`advertise_on_slo` - (Optional) Enable on Site local outside network, default VIP will be used (`Bool`).

`advertise_on_slo_sli` - (Optional) Enable on Site local inside and outside network, default VIP will be used (`Bool`).

`domain_suffix` - (Required) Domain suffix will be used along with node name to form the hostname for ssh node management (`String`).

`node_ssh_ports` - (Required) Enter TCP port and node name per node. See [Enabled Ssh Access Node Ssh Ports ](#enabled-ssh-access-node-ssh-ports) below for details.

### Tcp Port Choice Custom Tcp Ports

select custom tcp ports.

`ports` - (Required) List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192 (`String`).

### Tcp Port Choice Default Tcp Ports

Select default TCP ports, 80 and 443.

### Tcp Port Choice Http Port

Select HTTP Port 80.

### Tcp Port Choice Https Port

Select HTTPS Port 443.

### Tcp Port Choice No Tcp Ports

do not select tcp ports.

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, vault_secret_info, clear_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Tunnel Prefix Choice Automatic Prefix

System will automatically select tunnel prefix.

### Udp Port Choice Custom Udp Ports

select custom udp ports.

`ports` - (Required) List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192 (`String`).

### Udp Port Choice No Udp Ports

do not select udp ports.

### Xfcc Header Xfcc Disabled

No X-Forwarded-Client-Cert header will be added.

### Xfcc Header Xfcc Options

X-Forwarded-Client-Cert header will be added with the configured fields.

`xfcc_header_elements` - (Required) X-Forwarded-Client-Cert header elements to be added to requests (`List of Strings`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured nfv_service.
