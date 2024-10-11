---

page_title: "Volterra: k8s_pod_security_policy"
description: "The k8s_pod_security_policy allows CRUD of K8s Pod Security Policy resource on Volterra SaaS"

---

Resource volterra_k8s_pod_security_policy
=========================================

The K8s Pod Security Policy allows CRUD of K8s Pod Security Policy resource on Volterra SaaS

~> **Note:** Please refer to [K8s Pod Security Policy API docs](https://docs.cloud.f5.com/docs-v2/api/k8s-pod-security-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_k8s_pod_security_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "psp_spec yaml" must be set

  psp_spec {
    allow_privilege_escalation = true

    // One of the arguments from this list "allowed_capabilities no_allowed_capabilities" must be set

    no_allowed_capabilities = true
    allowed_csi_drivers = ["value"]
    allowed_flex_volumes = ["value"]
    allowed_host_paths {
      path_prefix = "value"

      read_only = true
    }
    allowed_proc_mounts = ["value"]
    allowed_unsafe_sysctls = ["value"]
    default_allow_privilege_escalation = true

    // One of the arguments from this list "default_capabilities no_default_capabilities" must be set

    no_default_capabilities = true

    // One of the arguments from this list "drop_capabilities no_drop_capabilities" must be set

    no_drop_capabilities = true
    forbidden_sysctls = ["value"]

    // One of the arguments from this list "fs_group_strategy_options no_fs_groups" must be set

    no_fs_groups = true

    // One of the arguments from this list "no_run_as_group run_as_group" must be set

    no_run_as_group = true
    host_ipc = true
    host_network = true
    host_pid = true
    host_port_ranges = "80,443,8080-8191,9080"
    privileged = true
    read_only_root_filesystem = true

    // One of the arguments from this list "no_runtime_class runtime_class" must be set

    no_runtime_class = true

    // One of the arguments from this list "no_se_linux_options se_linux_options" must be set

    no_se_linux_options = true

    // One of the arguments from this list "no_supplemental_groups supplemental_groups" must be set

    no_supplemental_groups = true

    // One of the arguments from this list "no_run_as_user run_as_user" must be set

    no_run_as_user = true
    volumes = ["gitRepo"]
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

###### One of the arguments from this list "psp_spec, yaml" must be set

`psp_spec` - (Optional) Form based pod security specification. See [Config Method Choice Psp Spec ](#config-method-choice-psp-spec) below for details.

`yaml` - (Optional) K8s YAML for Pod Security Policy (`String`).

### Allowed Capabilities Choice Allowed Capabilities

Allowed Capabilities to add pod spec in addition to default capabilities.

`capabilities` - (Required) List of capabilities that docker container has. (`String`).

### Allowed Capabilities Choice No Allowed Capabilities

Add capabilities is not allowed in POD..

### Config Method Choice Psp Spec

Form based pod security specification.

`allow_privilege_escalation` - (Optional) Pod can request to privilege escalation (`Bool`).

###### One of the arguments from this list "allowed_capabilities, no_allowed_capabilities" must be set

`allowed_capabilities` - (Optional) Allowed Capabilities to add pod spec in addition to default capabilities. See [Allowed Capabilities Choice Allowed Capabilities ](#allowed-capabilities-choice-allowed-capabilities) below for details.

`no_allowed_capabilities` - (Optional) Add capabilities is not allowed in POD. (`Bool`).

`allowed_csi_drivers` - (Optional) Restrict the available CSI drivers for POD, default all drivers are available. (`String`).

`allowed_flex_volumes` - (Optional) Restrict list of Flex volumes, default all volumes are allowed (`String`).

`allowed_host_paths` - (Optional) Restrict list of host paths, default all host paths are allowed. See [Psp Spec Allowed Host Paths ](#psp-spec-allowed-host-paths) below for details.

`allowed_proc_mounts` - (Optional) allowed list of proc mounts, empty list allows default proc mounts. (`String`).

`allowed_unsafe_sysctls` - (Optional) allowed list of unsafe sysctls, empty list allows none. supports prefix reg-ex (`String`).

`default_allow_privilege_escalation` - (Optional) Pod has permission for privilege escalation by default (`Bool`).

###### One of the arguments from this list "default_capabilities, no_default_capabilities" must be set

`default_capabilities` - (Optional) Default capabilities that will be added to container unless, Pod spec drops it.. See [Default Capabilities Choice Default Capabilities ](#default-capabilities-choice-default-capabilities) below for details.

`no_default_capabilities` - (Optional) K8s Default capabilities will be added to container unless pod spec drops it. (`Bool`).

###### One of the arguments from this list "drop_capabilities, no_drop_capabilities" must be set

`drop_capabilities` - (Optional) Capabilities to drop from K8s default capabilities, should not used with custom default capabilities. See [Drop Capabilities Choice Drop Capabilities ](#drop-capabilities-choice-drop-capabilities) below for details.

`no_drop_capabilities` - (Optional) Capabilities are not dropped from K8s default capabilities (`Bool`).

`forbidden_sysctls` - (Optional) Forbidden list of sysctls, empty list forbids none. supports prefix reg-ex (`String`).

###### One of the arguments from this list "fs_group_strategy_options, no_fs_groups" must be set

`fs_group_strategy_options` - (Optional) FS Groups that are used by security context. See [Fs Group Choice Fs Group Strategy Options ](#fs-group-choice-fs-group-strategy-options) below for details.

`no_fs_groups` - (Optional) Default K8s allowed FS group ids can be used (`Bool`).

###### One of the arguments from this list "no_run_as_group, run_as_group" must be set

`no_run_as_group` - (Optional) Default K8s allowed group ids can be used as run as group in POD spec. (`Bool`).

`run_as_group` - (Optional) Controls Allowable run as group values. See [Group Choice Run As Group ](#group-choice-run-as-group) below for details.

`host_ipc` - (Optional) Host IPC determines if the policy allows the use of host IPC in the pod spec. (`Bool`).

`host_network` - (Optional) Host Network determines if the policy allows the use of host network in the pod spec. (`Bool`).

`host_pid` - (Optional) Host PID determines if the policy allows the use of host PID in the pod spec. (`Bool`).

`host_port_ranges` - (Optional) Host port ranges determines which ports ranges are allowed to be exposed (`String`).

`privileged` - (Optional) Privileged determines if a pod can request to be run as privileged. (`Bool`).

`read_only_root_filesystem` - (Optional) Containers can only run with read only root filesystem. (`Bool`).

###### One of the arguments from this list "no_runtime_class, runtime_class" must be set

`no_runtime_class` - (Optional) Default K8s allowed runtime class options can be used (`Bool`).

`runtime_class` - (Optional) Controls Allowable Runtime Class values, if not present all values are allowed. See [Runtime Class Choice Runtime Class ](#runtime-class-choice-runtime-class) below for details.(Deprecated)

###### One of the arguments from this list "no_se_linux_options, se_linux_options" must be set

`no_se_linux_options` - (Optional) Default K8s allowed SE Linux options can be used (`Bool`).

`se_linux_options` - (Optional) Controls Allowable SE Linux labels, if not present all values are allowed. See [Se Linux Choice Se Linux Options ](#se-linux-choice-se-linux-options) below for details.(Deprecated)

###### One of the arguments from this list "no_supplemental_groups, supplemental_groups" must be set

`no_supplemental_groups` - (Optional) Default K8s allowed supplemental group ids can be used (`Bool`).

`supplemental_groups` - (Optional) Supplemental Groups that are used by security context. See [Supplemental Group Choice Supplemental Groups ](#supplemental-group-choice-supplemental-groups) below for details.

###### One of the arguments from this list "no_run_as_user, run_as_user" must be set

`no_run_as_user` - (Optional) Default K8s allowed user ids can be used as run as user in POD spec. (`Bool`).

`run_as_user` - (Optional) Controls Allowable run as user values. See [User Choice Run As User ](#user-choice-run-as-user) below for details.

`volumes` - (Optional) Allow List of volume plugins. Empty no volumes are allowed (`String`).

### Default Capabilities Choice Default Capabilities

Default capabilities that will be added to container unless, Pod spec drops it..

`capabilities` - (Required) List of capabilities that docker container has. (`String`).

### Default Capabilities Choice No Default Capabilities

K8s Default capabilities will be added to container unless pod spec drops it..

### Drop Capabilities Choice Drop Capabilities

Capabilities to drop from K8s default capabilities, should not used with custom default capabilities.

`capabilities` - (Required) List of capabilities that docker container has. (`String`).

### Drop Capabilities Choice No Drop Capabilities

Capabilities are not dropped from K8s default capabilities.

### Fs Group Choice Fs Group Strategy Options

FS Groups that are used by security context.

`id_ranges` - (Optional) List of range of ID(s). See [Fs Group Strategy Options Id Ranges ](#fs-group-strategy-options-id-ranges) below for details.

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

### Fs Group Choice No Fs Groups

Default K8s allowed FS group ids can be used.

### Fs Group Strategy Options Id Ranges

List of range of ID(s).

`max_id` - (Required) Ending(maximum) ID for for ID range (`Int`).

`min_id` - (Required) Starting(minimum) ID for for ID range (`Int`).

### Group Choice No Run As Group

Default K8s allowed group ids can be used as run as group in POD spec..

### Group Choice Run As Group

Controls Allowable run as group values.

`id_ranges` - (Optional) List of range of ID(s). See [Run As Group Id Ranges ](#run-as-group-id-ranges) below for details.

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

### Psp Spec Allowed Host Paths

Restrict list of host paths, default all host paths are allowed.

`path_prefix` - (Required) Host path prefix is the path prefix that the host volume must match. It does not support *. (`String`).

`read_only` - (Optional) This volume will be allowed to mount read only. (`Bool`).

### Run As Group Id Ranges

List of range of ID(s).

`max_id` - (Required) Ending(maximum) ID for for ID range (`Int`).

`min_id` - (Required) Starting(minimum) ID for for ID range (`Int`).

### Run As User Id Ranges

List of range of ID(s).

`max_id` - (Required) Ending(maximum) ID for for ID range (`Int`).

`min_id` - (Required) Starting(minimum) ID for for ID range (`Int`).

### Runtime Class Choice No Runtime Class

Default K8s allowed runtime class options can be used.

### Runtime Class Choice Runtime Class

Controls Allowable Runtime Class values, if not present all values are allowed.

`allowed_runtime_class_names` - (Required) List of allowed runtime class names (`String`).

`default_runtime_class_name` - (Optional) description (`String`).

### Se Linux Choice No Se Linux Options

Default K8s allowed SE Linux options can be used.

### Se Linux Choice Se Linux Options

Controls Allowable SE Linux labels, if not present all values are allowed.

`Level` - (Optional) Run as SE linux level label (`String`).

`role` - (Optional) Run as SE linux label (`String`).

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

`type` - (Optional) Run as SE linux type label (`String`).

`user` - (Optional) Run as SE linux user label (`String`).

### Supplemental Group Choice No Supplemental Groups

Default K8s allowed supplemental group ids can be used.

### Supplemental Group Choice Supplemental Groups

Supplemental Groups that are used by security context.

`id_ranges` - (Optional) List of range of ID(s). See [Supplemental Groups Id Ranges ](#supplemental-groups-id-ranges) below for details.

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

### Supplemental Groups Id Ranges

List of range of ID(s).

`max_id` - (Required) Ending(maximum) ID for for ID range (`Int`).

`min_id` - (Required) Starting(minimum) ID for for ID range (`Int`).

### User Choice No Run As User

Default K8s allowed user ids can be used as run as user in POD spec..

### User Choice Run As User

Controls Allowable run as user values.

`id_ranges` - (Optional) List of range of ID(s). See [Run As User Id Ranges ](#run-as-user-id-ranges) below for details.

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured k8s_pod_security_policy.
