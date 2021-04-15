---

page_title: "Volterra: k8s_pod_security_policy"

description: "The k8s_pod_security_policy allows CRUD of K8s Pod Security Policy resource on Volterra SaaS"
-----------------------------------------------------------------------------------------------------------

Resource volterra_k8s_pod_security_policy
=========================================

The K8s Pod Security Policy allows CRUD of K8s Pod Security Policy resource on Volterra SaaS

~> **Note:** Please refer to [K8s Pod Security Policy API docs](https://volterra.io/docs/api/k8s-pod-security-policy) to learn more

Example Usage
-------------

```hcl
resource "volterra_k8s_pod_security_policy" "example" {
  name      = "acmecorp-web"
  namespace = "staging"

  // One of the arguments from this list "psp_spec yaml" must be set

  psp_spec {
    allow_privilege_escalation = true

    // One of the arguments from this list "no_allowed_capabilities allowed_capabilities" must be set
    no_allowed_capabilities = true

    allowed_csi_drivers = ["allowed_csi_drivers"]

    allowed_flex_volumes = ["allowed_flex_volumes"]

    allowed_host_paths {
      path_prefix = "path_prefix"
      read_only   = true
    }

    allowed_proc_mounts = ["allowed_proc_mounts"]

    allowed_unsafe_sysctls             = ["allowed_unsafe_sysctls"]
    default_allow_privilege_escalation = true

    // One of the arguments from this list "no_default_capabilities default_capabilities" must be set
    no_default_capabilities = true

    // One of the arguments from this list "no_drop_capabilities drop_capabilities" must be set
    no_drop_capabilities = true

    forbidden_sysctls = ["forbidden_sysctls"]

    // One of the arguments from this list "no_fs_groups fs_group_strategy_options" must be set
    no_fs_groups = true

    // One of the arguments from this list "no_run_as_group run_as_group" must be set
    no_run_as_group           = true
    host_ipc                  = true
    host_network              = true
    host_pid                  = true
    host_port_ranges          = "80,443,8080-8191,9080"
    privileged                = true
    read_only_root_filesystem = true

    // One of the arguments from this list "no_runtime_class runtime_class" must be set
    no_runtime_class = true

    // One of the arguments from this list "no_se_linux_options se_linux_options" must be set
    no_se_linux_options = true

    // One of the arguments from this list "no_supplemental_groups supplemental_groups" must be set
    no_supplemental_groups = true

    // One of the arguments from this list "run_as_user no_run_as_user" must be set
    no_run_as_user = true

    volumes = ["volumes"]
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

`psp_spec` - (Optional) Form based pod security specification. See [Psp Spec ](#psp-spec) below for details.

`yaml` - (Optional) K8s YAML for Pod Security Policy (`String`).

### Allowed Capabilities

Allowed Capabilities to add pod spec in addition to default capabilities.

`capabilities` - (Required) List of capabilities that docker container has. (`String`).

### Allowed Host Paths

Restrict list of host paths, default all host paths are allowed.

`path_prefix` - (Required) Host path prefix is the path prefix that the host volume must match. It does not support *. (`String`).

`read_only` - (Optional) This volume will be allowed to mount read only. (`Bool`).

### Default Capabilities

Default capabilities that will be added to container unless, Pod spec drops it..

`capabilities` - (Required) List of capabilities that docker container has. (`String`).

### Drop Capabilities

Capabilities to drop from K8s default capabilities, should not used with custom default capabilities.

`capabilities` - (Required) List of capabilities that docker container has. (`String`).

### Fs Group Strategy Options

FS Groups that are used by security context.

`id_ranges` - (Optional) List of range of ID(s). See [Id Ranges ](#id-ranges) below for details.

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

### Id Ranges

List of range of ID(s).

`max_id` - (Required) Ending(maximum) ID for for ID range (`Int`).

`min_id` - (Required) Starting(minimum) ID for for ID range (`Int`).

### No Allowed Capabilities

Add capabilities is not allowed in POD..

### No Default Capabilities

K8s Default capabilities will be added to container unless pod spec drops it..

### No Drop Capabilities

Capabilities are not dropped from K8s default capabilities.

### No Fs Groups

Default K8s allowed FS group ids can be used.

### No Run As Group

Default K8s allowed group ids can be used as run as group in POD spec..

### No Run As User

Default K8s allowed user ids can be used as run as user in POD spec..

### No Runtime Class

Default K8s allowed runtime class options can be used.

### No Se Linux Options

Default K8s allowed SE Linux options can be used.

### No Supplemental Groups

Default K8s allowed supplemental group ids can be used.

### Psp Spec

Form based pod security specification.

`allow_privilege_escalation` - (Optional) Pod can request to privilege escalation (`Bool`).

`allowed_capabilities` - (Optional) Allowed Capabilities to add pod spec in addition to default capabilities. See [Allowed Capabilities ](#allowed-capabilities) below for details.

`no_allowed_capabilities` - (Optional) Add capabilities is not allowed in POD. (bool).

`allowed_csi_drivers` - (Optional) Restrict the available CSI drivers for POD, default all drivers are available. (`String`).

`allowed_flex_volumes` - (Optional) Restrict list of Flex volumes, default all volumes are allowed (`String`).

`allowed_host_paths` - (Optional) Restrict list of host paths, default all host paths are allowed. See [Allowed Host Paths ](#allowed-host-paths) below for details.

`allowed_proc_mounts` - (Optional) allowed list of proc mounts, empty list allows default proc mounts. (`String`).

`allowed_unsafe_sysctls` - (Optional) allowed list of unsafe sysctls, empty list allows none. supports prefix reg-ex (`String`).

`default_allow_privilege_escalation` - (Optional) Pod has permission for privilege escalation by default (`Bool`).

`default_capabilities` - (Optional) Default capabilities that will be added to container unless, Pod spec drops it.. See [Default Capabilities ](#default-capabilities) below for details.

`no_default_capabilities` - (Optional) K8s Default capabilities will be added to container unless pod spec drops it. (bool).

`drop_capabilities` - (Optional) Capabilities to drop from K8s default capabilities, should not used with custom default capabilities. See [Drop Capabilities ](#drop-capabilities) below for details.

`no_drop_capabilities` - (Optional) Capabilities are not dropped from K8s default capabilities (bool).

`forbidden_sysctls` - (Optional) Forbidden list of sysctls, empty list forbids none. supports prefix reg-ex (`String`).

`fs_group_strategy_options` - (Optional) FS Groups that are used by security context. See [Fs Group Strategy Options ](#fs-group-strategy-options) below for details.

`no_fs_groups` - (Optional) Default K8s allowed FS group ids can be used (bool).

`no_run_as_group` - (Optional) Default K8s allowed group ids can be used as run as group in POD spec. (bool).

`run_as_group` - (Optional) Controls Allowable run as group values. See [Run As Group ](#run-as-group) below for details.

`host_ipc` - (Optional) Host IPC determines if the policy allows the use of host IPC in the pod spec. (`Bool`).

`host_network` - (Optional) Host Network determines if the policy allows the use of host network in the pod spec. (`Bool`).

`host_pid` - (Optional) Host PID determines if the policy allows the use of host PID in the pod spec. (`Bool`).

`host_port_ranges` - (Optional) Host port ranges determines which ports ranges are allowed to be exposed (`String`).

`privileged` - (Optional) Privileged determines if a pod can request to be run as privileged. (`Bool`).

`read_only_root_filesystem` - (Optional) Containers can only run with read only root filesystem. (`Bool`).

`no_runtime_class` - (Optional) Default K8s allowed runtime class options can be used (bool).

`runtime_class` - (Optional) Controls Allowable Runtime Class values, if not present all values are allowed. See [Runtime Class ](#runtime-class) below for details.

`no_se_linux_options` - (Optional) Default K8s allowed SE Linux options can be used (bool).

`se_linux_options` - (Optional) Controls Allowable SE Linux labels, if not present all values are allowed. See [Se Linux Options ](#se-linux-options) below for details.

`no_supplemental_groups` - (Optional) Default K8s allowed supplemental group ids can be used (bool).

`supplemental_groups` - (Optional) Supplemental Groups that are used by security context. See [Supplemental Groups ](#supplemental-groups) below for details.

`no_run_as_user` - (Optional) Default K8s allowed user ids can be used as run as user in POD spec. (bool).

`run_as_user` - (Optional) Controls Allowable run as user values. See [Run As User ](#run-as-user) below for details.

`volumes` - (Optional) Allow List of volume plugins. Empty no volumes are allowed (`String`).

### Run As Group

Controls Allowable run as group values.

`id_ranges` - (Optional) List of range of ID(s). See [Id Ranges ](#id-ranges) below for details.

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

### Run As User

Controls Allowable run as user values.

`id_ranges` - (Optional) List of range of ID(s). See [Id Ranges ](#id-ranges) below for details.

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

### Runtime Class

Controls Allowable Runtime Class values, if not present all values are allowed.

`allowed_runtime_class_names` - (Required) List of allowed runtime class names (`String`).

`default_runtime_class_name` - (Optional) description (`String`).

### Se Linux Options

Controls Allowable SE Linux labels, if not present all values are allowed.

`Level` - (Optional) Run as SE linux level label (`String`).

`role` - (Optional) Run as SE linux label (`String`).

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

`type` - (Optional) Run as SE linux type label (`String`).

`user` - (Optional) Run as SE linux user label (`String`).

### Supplemental Groups

Supplemental Groups that are used by security context.

`id_ranges` - (Optional) List of range of ID(s). See [Id Ranges ](#id-ranges) below for details.

`rule` - (Optional) Rule indicated how the FS group ID range is used (`String`).

Attribute Reference
-------------------

-	`id` - This is the id of the configured k8s_pod_security_policy.
