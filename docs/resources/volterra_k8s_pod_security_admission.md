











---
page_title: "Volterra: k8s_pod_security_admission"
description: "The k8s_pod_security_admission allows CRUD of K8s Pod Security Admission  resource on Volterra SaaS"
---
# Resource volterra_k8s_pod_security_admission

The K8s Pod Security Admission  allows CRUD of K8s Pod Security Admission  resource on Volterra SaaS

~> **Note:** Please refer to [K8s Pod Security Admission  API docs](https://docs.cloud.f5.com/docs-v2/api/k8s-pod-security-admission) to learn more

## Example Usage

```hcl
resource "volterra_k8s_pod_security_admission" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  pod_security_admission_specs {
    // One of the arguments from this list "enforce audit warn" must be set

    audit = true

    // One of the arguments from this list "privileged baseline restricted" must be set

    privileged = true
  }
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

`pod_security_admission_specs` - (Required) x-required. See [Pod Security Admission Specs ](#pod-security-admission-specs) below for details.




		




		




		






		




		




		





### Pod Security Admission Specs 

 x-required.



###### One of the arguments from this list "warn, enforce, audit" must be set

`audit` - (Optional) Policy violations will trigger the addition of an audit annotation to the event recorded in the audit log, but are otherwise allowed. (`Bool`).


`enforce` - (Optional) Policy violations will cause the pod to be rejected. (`Bool`).


`warn` - (Optional) Policy violations will trigger a user-facing warning, but are otherwise allowed. (`Bool`).




###### One of the arguments from this list "privileged, baseline, restricted" must be set

`baseline` - (Optional) Minimally restrictive policy which prevents known privilege escalations (`Bool`).


`privileged` - (Optional) Unrestricted policy, providing the widest possible level of permissions (`Bool`).


`restricted` - (Optional) Heavily restricted policy, following current Pod hardening best practices. (`Bool`).




### Admission Mode Choice Audit 

 Policy violations will trigger the addition of an audit annotation to the event recorded in the audit log, but are otherwise allowed..



### Admission Mode Choice Enforce 

 Policy violations will cause the pod to be rejected..



### Admission Mode Choice Warn 

 Policy violations will trigger a user-facing warning, but are otherwise allowed..



### Policy Type Choice Baseline 

 Minimally restrictive policy which prevents known privilege escalations.



### Policy Type Choice Privileged 

 Unrestricted policy, providing the widest possible level of permissions.



### Policy Type Choice Restricted 

 Heavily restricted policy, following current Pod hardening best practices..



## Attribute Reference

* `id` - This is the id of the configured k8s_pod_security_admission.

