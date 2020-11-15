output "new_domain_name_for_web_app" {
  description = "Use the domain name provided to access you helloworld app."
  value       = "${volterra_namespace.volterra_ns.tenant_name}.${var.webapp}.helloworld.app"
}

output "dns_cname" {
  description = "Use the domain name provided to access you helloworld app."
  value       = "${data.volterra_virtual_host_dns_info.dns_info.host_name}"
}

output "vk8s_kubeconfig" {
  value = "${join("", volterra_api_credential.api_cred.*.vk8s_kubeconfig)}"
}
