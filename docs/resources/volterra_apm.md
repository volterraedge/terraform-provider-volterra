---

page_title: "Volterra: apm"
description: "The apm allows CRUD of Apm resource on Volterra SaaS"

---

Resource volterra_apm
=====================

The Apm allows CRUD of Apm resource on Volterra SaaS

~> **Note:** Please refer to [Apm API docs](https://docs.cloud.f5.com/docs-v2/api/bigip-apm) to learn more

Example Usage
-------------

```hcl
resource "volterra_apm" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "https_management" must be set

  https_management {
    // One of the arguments from this list "advertise_on_internet advertise_on_internet_default_vip advertise_on_sli_vip advertise_on_slo_internet_vip advertise_on_slo_sli advertise_on_slo_vip disable_local do_not_advertise_on_internet" must be set

    advertise_on_slo_internet_vip {
      // One of the arguments from this list "no_mtls use_mtls" must be set

      no_mtls = true

      tls_certificates {
        certificate_url = "value"

        description = "Certificate used in production environment"

        // One of the arguments from this list "custom_hash_algorithms disable_ocsp_stapling use_system_defaults" can be set

        use_system_defaults {}
        private_key {
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
      }

      tls_config {
        // One of the arguments from this list "custom_security default_security low_security medium_security" must be set

        default_security = true
      }
    }

    domain_suffix = "foo.com"

    // One of the arguments from this list "advertise_on_public advertise_on_public_default_vip do_not_advertise" can be set

    advertise_on_public {
      public_ip {
        name      = "test1"
        namespace = "staging"
        tenant    = "acmecorp"
      }
    }

    // One of the arguments from this list "default_https_port https_port" must be set

    default_https_port = true
  }

  // One of the arguments from this list "aws_site_type_choice baremetal_site_type_choice" must be set

  aws_site_type_choice {
    apm_aws_site {
      admin_password {
        blindfold_secret_info_internal {
          decryption_provider = "value"

          location = "string:///U2VjcmV0SW5mb3JtYXRpb24="

          store_provider = "value"
        }

        secret_encoding_type = "secret_encoding_type"

        // One of the arguments from this list "blindfold_secret_info clear_secret_info vault_secret_info wingman_secret_info" must be set

        clear_secret_info {
          provider = "box-provider"

          url = "string:///U2VjcmV0SW5mb3JtYXRpb24="
        }
      }

      admin_username = "admin"

      aws_tgw_site {
        aws_tgw_site {
          name      = "test1"
          namespace = "staging"
          tenant    = "acmecorp"
        }
      }

      endpoint_service {
        // One of the arguments from this list "advertise_on_slo_ip advertise_on_slo_ip_external disable_advertise_on_slo_ip" must be set

        disable_advertise_on_slo_ip = true

        // One of the arguments from this list "automatic_vip configured_vip" must be set

        automatic_vip = true

        // One of the arguments from this list "custom_tcp_ports default_tcp_ports http_port https_port no_tcp_ports" must be set

        custom_tcp_ports {
          ports = ["100-200"]
        }

        // One of the arguments from this list "custom_udp_ports no_udp_ports" must be set

        no_udp_ports = true
      }

      nodes {
        aws_az_name = "us-west-2a"

        // One of the arguments from this list "mgmt_subnet reserved_mgmt_subnet" must be set

        mgmt_subnet {
          // One of the arguments from this list "existing_subnet_id subnet_param" must be set

          existing_subnet_id = "subnet-12345678901234567"
        }
        node_name = "node1"

        // One of the arguments from this list "automatic_prefix tunnel_prefix" must be set

        automatic_prefix = true
      }

      ssh_key = "ssh-rsa AAAAB..."

      tags = {
        "key1" = "value1"
      }
    }

    // One of the arguments from this list "market_place_image" must be set

    market_place_image {
      // One of the arguments from this list "BestPlusPayG200Mbps best_plus_payg_1gbps" must be set

      BestPlusPayG200Mbps = true
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

###### One of the arguments from this list "https_management" must be set

`https_management` - (Optional) Enable HTTPS based management. See [Http Management Choice Https Management ](#http-management-choice-https-management) below for details.

###### One of the arguments from this list "aws_site_type_choice, baremetal_site_type_choice" must be set

`aws_site_type_choice` - (Optional) Virtual F5 BIG-IP APM service to be deployed on AWS Transit Gateway Site. See [Site Type Choice Aws Site Type Choice ](#site-type-choice-aws-site-type-choice) below for details.

`baremetal_site_type_choice` - (Optional) Virtual F5 BIG-IP APM service to be deployed on App Stack Bare Metal Site. See [Site Type Choice Baremetal Site Type Choice ](#site-type-choice-baremetal-site-type-choice) below for details.

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

### Advertise Choice Advertise On Sli Vip

Enable on Site local inside network, default VIP will be used.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Advertise On Sli Vip Tls Certificates ](#advertise-on-sli-vip-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Advertise On Sli Vip Tls Config ](#advertise-on-sli-vip-tls-config) below for details.

### Advertise Choice Advertise On Slo Internet Vip

Enable On Site Local Outside Internet VIP.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Advertise On Slo Internet Vip Tls Certificates ](#advertise-on-slo-internet-vip-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Advertise On Slo Internet Vip Tls Config ](#advertise-on-slo-internet-vip-tls-config) below for details.

### Advertise Choice Advertise On Slo Sli

Enable on Site local inside and outside network, default VIP will be used.

###### One of the arguments from this list "no_mtls, use_mtls" must be set

`no_mtls` - (Optional) x-displayName: "Disable" (`Bool`).

`use_mtls` - (Optional) x-displayName: "Enable". See [Mtls Choice Use Mtls ](#mtls-choice-use-mtls) below for details.

`tls_certificates` - (Required) for example, domain.com and *.domain.com - but use different signature algorithms. See [Advertise On Slo Sli Tls Certificates ](#advertise-on-slo-sli-tls-certificates) below for details.

`tls_config` - (Optional) Configuration of TLS settings such as min/max TLS version and ciphersuites. See [Advertise On Slo Sli Tls Config ](#advertise-on-slo-sli-tls-config) below for details.

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

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Advertise On Sli Vip Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Advertise On Slo Internet Vip Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Advertise On Slo Internet Vip Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Advertise On Slo Sli Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Advertise On Slo Sli Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Advertise On Slo Vip Tls Certificates

for example, domain.com and *.domain.com - but use different signature algorithms.

`certificate_url` - (Required) Certificate or certificate chain in PEM format including the PEM headers. (`String`).

`description` - (Optional) Description for the certificate (`String`).

###### One of the arguments from this list "custom_hash_algorithms, disable_ocsp_stapling, use_system_defaults" can be set

`custom_hash_algorithms` - (Optional) Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set.. See [Ocsp Stapling Choice Custom Hash Algorithms ](#ocsp-stapling-choice-custom-hash-algorithms) below for details.

`disable_ocsp_stapling` - (Optional) This is the default behavior if no choice is selected.. See [Ocsp Stapling Choice Disable Ocsp Stapling ](#ocsp-stapling-choice-disable-ocsp-stapling) below for details.

`use_system_defaults` - (Optional) F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order.. See [Ocsp Stapling Choice Use System Defaults ](#ocsp-stapling-choice-use-system-defaults) below for details.

`private_key` - (Required) TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate.. See [Tls Certificates Private Key ](#tls-certificates-private-key) below for details.

### Advertise On Slo Vip Tls Config

Configuration of TLS settings such as min/max TLS version and ciphersuites.

###### One of the arguments from this list "custom_security, default_security, low_security, medium_security" must be set

`custom_security` - (Optional) Custom selection of TLS versions and cipher suites. See [Choice Custom Security ](#choice-custom-security) below for details.

`default_security` - (Optional) TLS v1.2+ with PFS ciphers and strong crypto algorithms. (`Bool`).

`low_security` - (Optional) TLS v1.0+ including non-PFS ciphers and weak crypto algorithms. (`Bool`).

`medium_security` - (Optional) TLS v1.0+ with PFS ciphers and medium strength crypto algorithms. (`Bool`).

### Ami Choice BestPlusPayG200Mbps

F5 Best Plus with all modules in 200Mbps flavor.

### Ami Choice Best Plus Payg 1gbps

F5 Best Plus with all modules in 1Gbps flavor.

### Apm Aws Site Admin Password

Secret admin password for BIG-IP.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Admin Password Blindfold Secret Info Internal ](#admin-password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### Apm Aws Site Aws Tgw Site

Reference to AWS transit gateway site.

`aws_tgw_site` - (Required) Reference to AWS transit gateway site. See [ref](#ref) below for details.

### Apm Aws Site Endpoint Service

External service type is Endpoint service.

###### One of the arguments from this list "advertise_on_slo_ip, advertise_on_slo_ip_external, disable_advertise_on_slo_ip" must be set

`advertise_on_slo_ip` - (Optional) Advertise this loadbalancer on Site Local Outside network address (`Bool`).

`advertise_on_slo_ip_external` - (Optional) Advertise this loadbalancer on Site Local Outside network address and enable cloud external IP (`Bool`).

`disable_advertise_on_slo_ip` - (Optional) Do not Advertise this loadbalancer on Site Local Outside network address (`Bool`).

###### One of the arguments from this list "automatic_vip, configured_vip" must be set

`automatic_vip` - (Optional) System will automatically select a VIP (`Bool`).

`configured_vip` - (Optional) Enter IP address for the default VIP (`String`).

###### One of the arguments from this list "custom_tcp_ports, default_tcp_ports, http_port, https_port, no_tcp_ports" must be set

`custom_tcp_ports` - (Optional) Select custom TCP Ports. See [Tcp Port Choice Custom Tcp Ports ](#tcp-port-choice-custom-tcp-ports) below for details.

`default_tcp_ports` - (Optional) Select default TCP Ports, 80 and 443 (`Bool`).

`http_port` - (Optional) Select HTTP Port 80 (`Bool`).

`https_port` - (Optional) Select HTTPS Port 443 (`Bool`).

`no_tcp_ports` - (Optional) Do not select TCP Ports (`Bool`).

###### One of the arguments from this list "custom_udp_ports, no_udp_ports" must be set

`custom_udp_ports` - (Optional) select custom udp ports. See [Udp Port Choice Custom Udp Ports ](#udp-port-choice-custom-udp-ports) below for details.

`no_udp_ports` - (Optional) do not select udp ports (`Bool`).

### Apm Aws Site Nodes

Specify how and where the service nodes are spawned.

`aws_az_name` - (Required) The AWS Availability Zone must be consistent with the AWS Region chosen. Please select an AZ in the same Region as your TGW Site (`String`).

###### One of the arguments from this list "mgmt_subnet, reserved_mgmt_subnet" must be set

`mgmt_subnet` - (Optional) Select Existing Subnet or Create New. See [Mgmt Subnet Choice Mgmt Subnet ](#mgmt-subnet-choice-mgmt-subnet) below for details.

`reserved_mgmt_subnet` - (Optional) Autogenerate and reserve a subnet from the Primary CIDR (`Bool`).

`node_name` - (Required) Node Name will be used to assign as hostname to the service (`String`).

###### One of the arguments from this list "automatic_prefix, tunnel_prefix" must be set

`automatic_prefix` - (Optional) System will automatically select tunnel prefix (`Bool`).

`tunnel_prefix` - (Optional) Enter IP prefix for the tunnel, it has to be /30 (`String`).

### Aws Site Type Choice Apm Aws Site

Virtual F5 BIG-IP service to be deployed on AWS.

`admin_password` - (Required) Secret admin password for BIG-IP. See [Apm Aws Site Admin Password ](#apm-aws-site-admin-password) below for details.

`admin_username` - (Required) Admin Username for BIG-IP (`String`).

`aws_tgw_site` - (Required) Reference to AWS transit gateway site. See [Apm Aws Site Aws Tgw Site ](#apm-aws-site-aws-tgw-site) below for details.

`endpoint_service` - (Optional) External service type is Endpoint service. See [Apm Aws Site Endpoint Service ](#apm-aws-site-endpoint-service) below for details.

`nodes` - (Required) Specify how and where the service nodes are spawned. See [Apm Aws Site Nodes ](#apm-aws-site-nodes) below for details.

`ssh_key` - (Required) Public SSH key for accessing the BIG-IP nodes. (`String`).

`tags` - (Optional) It helps to manage, identify, organize, search for, and filter resources in AWS console. (`String`).

### Baremetal Site Type Choice F5 Bare Metal Site

Virtual BIG-IP specification for App Stack Bare Metal Site.

`admin_password` - (Required) Secret admin password for BIG-IP. See [F5 Bare Metal Site Admin Password ](#f5-bare-metal-site-admin-password) below for details.

`admin_username` - (Required) Admin Username for BIG-IP (`String`).

`bare_metal_site` - (Required) Reference to bare metal site on which BIG-IP should be deployed. See [ref](#ref) below for details.

`bigiq_instance` - (Required) Details of BIG-IQ Instance used for activating licenses.. See [F5 Bare Metal Site Bigiq Instance ](#f5-bare-metal-site-bigiq-instance) below for details.

`nodes` - (Required) Specify how and where the service nodes are spawned. See [F5 Bare Metal Site Nodes ](#f5-bare-metal-site-nodes) below for details.

`public_download_url` - (Required) Public URL where BIG-IP VE image (qcow2) is hosted (`String`).

`ssh_key` - (Required) Public SSH key for accessing the BIG-IP nodes. (`String`).

### Bigiq Instance Password

Password of the user used to access BIG-IQ to activate the license.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Password Blindfold Secret Info Internal ](#password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

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

### External Vip Choice Advertise On Slo Ip

Advertise this loadbalancer on Site Local Outside network address.

### External Vip Choice Advertise On Slo Ip External

Advertise this loadbalancer on Site Local Outside network address and enable cloud external IP.

### External Vip Choice Disable Advertise On Slo Ip

Do not Advertise this loadbalancer on Site Local Outside network address.

### F5 Bare Metal Site Admin Password

Secret admin password for BIG-IP.

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Admin Password Blindfold Secret Info Internal ](#admin-password-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

`blindfold_secret_info` - (Optional) Blindfold Secret is used for the secrets managed by F5XC Secret Management Service. See [Secret Info Oneof Blindfold Secret Info ](#secret-info-oneof-blindfold-secret-info) below for details.

`clear_secret_info` - (Optional) Clear Secret is used for the secrets that are not encrypted. See [Secret Info Oneof Clear Secret Info ](#secret-info-oneof-clear-secret-info) below for details.

`vault_secret_info` - (Optional) Vault Secret is used for the secrets managed by Hashicorp Vault. See [Secret Info Oneof Vault Secret Info ](#secret-info-oneof-vault-secret-info) below for details.(Deprecated)

`wingman_secret_info` - (Optional) Secret is given as bootstrap secret in F5XC Security Sidecar. See [Secret Info Oneof Wingman Secret Info ](#secret-info-oneof-wingman-secret-info) below for details.(Deprecated)

### F5 Bare Metal Site Bigiq Instance

Details of BIG-IQ Instance used for activating licenses..

`license_pool_name` - (Required) Name of Utility Pool on BIG-IQ (`String`).

`license_server_ip` - (Required) IP Address from the TCP Load Balancer which is configured to communicate with License Server (`String`).

`password` - (Required) Password of the user used to access BIG-IQ to activate the license. See [Bigiq Instance Password ](#bigiq-instance-password) below for details.

`sku_name` - (Required) License offering name aka SKU name (`String`).

`username` - (Required) User Name used to access BIG-IQ to activate the license (`String`).

### F5 Bare Metal Site Nodes

Specify how and where the service nodes are spawned.

`bm_node_memory_size` - (Required) x-required (`String`).

`bm_virtual_cpu_count` - (Required) x-required (`String`).

`external_interface` - (Optional). See [Nodes External Interface ](#nodes-external-interface) below for details.

`internal_interface` - (Optional). See [Nodes Internal Interface ](#nodes-internal-interface) below for details.

`node_name` - (Required) Node Name will be used to assign as hostname to the service (`String`).

### Http Management Choice Https Management

Enable HTTPS based management.

###### One of the arguments from this list "advertise_on_internet, advertise_on_internet_default_vip, advertise_on_sli_vip, advertise_on_slo_internet_vip, advertise_on_slo_sli, advertise_on_slo_vip, disable_local, do_not_advertise_on_internet" must be set

`advertise_on_internet` - (Optional) Advertise this loadbalancer on public network. See [Advertise Choice Advertise On Internet ](#advertise-choice-advertise-on-internet) below for details.

`advertise_on_internet_default_vip` - (Optional) Enable management access on internet with default VIP (`Bool`).

`advertise_on_sli_vip` - (Optional) Enable on Site local inside network, default VIP will be used. See [Advertise Choice Advertise On Sli Vip ](#advertise-choice-advertise-on-sli-vip) below for details.

`advertise_on_slo_internet_vip` - (Optional) Enable On Site Local Outside Internet VIP. See [Advertise Choice Advertise On Slo Internet Vip ](#advertise-choice-advertise-on-slo-internet-vip) below for details.

`advertise_on_slo_sli` - (Optional) Enable on Site local inside and outside network, default VIP will be used. See [Advertise Choice Advertise On Slo Sli ](#advertise-choice-advertise-on-slo-sli) below for details.

`advertise_on_slo_vip` - (Optional) Enable on Site local outside network, default VIP will be used. See [Advertise Choice Advertise On Slo Vip ](#advertise-choice-advertise-on-slo-vip) below for details.

`disable_local` - (Optional) Disable on Site local network (`Bool`).(Deprecated)

`do_not_advertise_on_internet` - (Optional) Do not enable access to management from internet (`Bool`).(Deprecated)

`domain_suffix` - (Required) Domain suffix will be used along with node name to form URL to access node management (`String`).

###### One of the arguments from this list "advertise_on_public, advertise_on_public_default_vip, do_not_advertise" can be set

`advertise_on_public` - (Optional) Advertise this loadbalancer on public network. See [Internet Choice Advertise On Public ](#internet-choice-advertise-on-public) below for details.(Deprecated)

`advertise_on_public_default_vip` - (Optional) Enable management access on internet with default VIP (`Bool`).(Deprecated)

`do_not_advertise` - (Optional) Do not enable access to management from internet (`Bool`).(Deprecated)

###### One of the arguments from this list "default_https_port, https_port" must be set

`default_https_port` - (Optional) Select default HTTPS 443 (`Bool`).

`https_port` - (Optional) Enter TCP port number (`Int`).

### Inside Vip Choice Automatic Vip

System will automatically select a VIP.

### Internet Choice Advertise On Public

Advertise this loadbalancer on public network.

`public_ip` - (Required) Dedicated Public IP, which is allocated by F5 Distributed Cloud on request, is used as a VIP.. See [ref](#ref) below for details.

### Internet Choice Advertise On Public Default Vip

Enable management access on internet with default VIP.

### Internet Choice Do Not Advertise

Do not enable access to management from internet.

### License Type Market Place Image

Select the BIG-IP pay as you go image to be used for this service.

###### One of the arguments from this list "BestPlusPayG200Mbps, best_plus_payg_1gbps" must be set

`BestPlusPayG200Mbps` - (Optional) F5 Best Plus with all modules in 200Mbps flavor (`Bool`).

`best_plus_payg_1gbps` - (Optional) F5 Best Plus with all modules in 1Gbps flavor (`Bool`).

### Mgmt Subnet Choice Mgmt Subnet

Select Existing Subnet or Create New.

###### One of the arguments from this list "existing_subnet_id, subnet_param" must be set

`existing_subnet_id` - (Optional) Information about existing subnet ID (`String`).

`subnet_param` - (Optional) Parameters for creating new subnet. See [Choice Subnet Param ](#choice-subnet-param) below for details.

### Mgmt Subnet Choice Reserved Mgmt Subnet

Autogenerate and reserve a subnet from the Primary CIDR.

### Mtls Choice No Mtls

x-displayName: "Disable".

### Mtls Choice Use Mtls

x-displayName: "Enable".

`client_certificate_optional` - (Optional) the connection will be accepted. (`Bool`).

###### One of the arguments from this list "crl, no_crl" can be set

`crl` - (Optional) Specify the CRL server information to download the certificate revocation list. See [ref](#ref) below for details.

`no_crl` - (Optional) Client certificate revocation status is not verified (`Bool`).

###### One of the arguments from this list "trusted_ca, trusted_ca_url" must be set

`trusted_ca` - (Optional) Select/Add a Root CA Certificate object to associate with this Load Balancer. See [ref](#ref) below for details.

`trusted_ca_url` - (Optional) Upload a Root CA Certificate specifically for this Load Balancer (`String`).

###### One of the arguments from this list "xfcc_disabled, xfcc_options" can be set

`xfcc_disabled` - (Optional) No X-Forwarded-Client-Cert header will be added (`Bool`).

`xfcc_options` - (Optional) X-Forwarded-Client-Cert header will be added with the configured fields. See [Xfcc Header Xfcc Options ](#xfcc-header-xfcc-options) below for details.

### Nodes External Interface

.

`interface` - (Required) L2 Interface on Site to be connected as interface on BIG-IP. See [ref](#ref) below for details.

`network_gateway` - (Optional) (`String`).

`network_self_ip` - (Required) Self IP CIDR (`String`).

### Nodes Internal Interface

.

`interface` - (Required) L2 Interface on Site to be connected as interface on BIG-IP. See [ref](#ref) below for details.

`network_gateway` - (Optional) (`String`).

`network_self_ip` - (Required) Self IP CIDR (`String`).

### Ocsp Stapling Choice Custom Hash Algorithms

Use hash algorithms in the custom order. F5XC will try to fetch ocsp response from the CA in the given order. Additionally, LoadBalancer will not become active until ocspResponse cannot be fetched if the certificate has MustStaple extension set..

`hash_algorithms` - (Required) Ordered list of hash algorithms to be used. (`List of Strings`).

### Ocsp Stapling Choice Disable Ocsp Stapling

This is the default behavior if no choice is selected..

### Ocsp Stapling Choice Use System Defaults

F5XC will try to fetch OCSPResponse with sha256 and sha1 as HashAlgorithm, in that order..

### Password Blindfold Secret Info Internal

Blindfold Secret Internal is used for the putting re-encrypted blindfold secret.

`decryption_provider` - (Optional) Name of the Secret Management Access object that contains information about the backend Secret Management service. (`String`).

`location` - (Required) Or it could be a path if the store provider is an http/https location (`String`).

`store_provider` - (Optional) This field needs to be provided only if the url scheme is not string:/// (`String`).

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

### Site Type Choice Aws Site Type Choice

Virtual F5 BIG-IP APM service to be deployed on AWS Transit Gateway Site.

`apm_aws_site` - (Required) Virtual F5 BIG-IP service to be deployed on AWS. See [Aws Site Type Choice Apm Aws Site ](#aws-site-type-choice-apm-aws-site) below for details.

###### One of the arguments from this list "market_place_image" must be set

`market_place_image` - (Optional) Select the BIG-IP pay as you go image to be used for this service. See [License Type Market Place Image ](#license-type-market-place-image) below for details.

### Site Type Choice Baremetal Site Type Choice

Virtual F5 BIG-IP APM service to be deployed on App Stack Bare Metal Site.

`f5_bare_metal_site` - (Required) Virtual BIG-IP specification for App Stack Bare Metal Site. See [Baremetal Site Type Choice F5 Bare Metal Site ](#baremetal-site-type-choice-f5-bare-metal-site) below for details.

### Tcp Port Choice Custom Tcp Ports

Select custom TCP Ports.

`ports` - (Required) List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192 (`String`).

### Tcp Port Choice Default Tcp Ports

Select default TCP Ports, 80 and 443.

### Tcp Port Choice Http Port

Select HTTP Port 80.

### Tcp Port Choice Https Port

Select HTTPS Port 443.

### Tcp Port Choice No Tcp Ports

Do not select TCP Ports.

### Tls Certificates Private Key

TLS Private Key data in unencrypted PEM format including the PEM headers. The data may be optionally secured using BlindFold. TLS key has to match the accompanying certificate..

`blindfold_secret_info_internal` - (Optional) Blindfold Secret Internal is used for the putting re-encrypted blindfold secret. See [Private Key Blindfold Secret Info Internal ](#private-key-blindfold-secret-info-internal) below for details.(Deprecated)

`secret_encoding_type` - (Optional) e.g. if a secret is base64 encoded and then put into vault. (`String`).(Deprecated)

###### One of the arguments from this list "blindfold_secret_info, clear_secret_info, vault_secret_info, wingman_secret_info" must be set

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

-	`id` - This is the id of the configured apm.
