---
page_title: "linuxkit_image Data Source - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_image is a single image to be included in the system.
---

# Data Source `linuxkit_image`

`linuxkit_image` is a single image to be included in the system.



## Schema

### Required

- **image** (String) The Docker image to use for the filesystem
- **name** (String) A unique name for the program being executed, used as the containerd id

### Optional

- **additional_gid_names** (List of String) A list of additional group names for the process
- **additional_gids** (List of Number) A list of additional groups for the process
- **ambient** (List of String) The Linux ambient capabilities (capabilities passed to non root users) that are required
- **annotations** (List of String) The map of key value pairs as OCI metadata
- **binds** (List of String) A Simpler interface to specify bind mounts, accepting a string like /src:/dest:opt1,opt2
- **capabilities** (List of String) The Linux capabilities required, for example CAP_SYS_ADMIN, if there is a single capability 'all' then all capabilities are added
- **cgroups_path** (String) The path for cgroups
- **command** (List of String) This will override the command and entrypoint in the image with a new list of commands
- **cwd** (String) The working directory, defaults to /
- **env** (List of String) This will override the environment in the image with a new environment list. Specify variables as VAR=value
- **gid** (Number) The group id of the process
- **gid_mappings** (Block List) (Experemental) gid mappings for user namespaces (see [below for nested schema](#nestedblock--gid_mappings))
- **gid_name** (String) The group name of the process
- **hostname** (String) The hostname inside the image
- **id** (String) The ID of this resource.
- **ipc** (String) The ipc namespace, either to a path, or if new is specified it will use a new namespace
- **masked_paths** (List of String) The paths which should be hidden
- **mounts** (Block List) The full form for specifying a mount, which requires type, source, destination and a list of options (see [below for nested schema](#nestedblock--mounts))
- **net** (String) The network namespace, either to a path, or if none or new is specified it will use a new namespace
- **no_new_privileges** (Boolean) If set to true means no additional capabilities can be acquired and suid binaries do not work
- **oom_score_adj** (Number) This changes the OOM score
- **pid** (String) The pid namespace, either to a path, or if host is specified it will use the host namespace
- **readonly** (Boolean) The root filesystem to read only, and changes the other default filesystems to read only
- **readonly_paths** (List of String) The paths which should be read only
- **resources** (Block List, Max: 1) The cgroup resource limits as per the OCI spec (see [below for nested schema](#nestedblock--resources))
- **rlimits** (List of String) The list of rlimit values in the form name,soft,hard, eg nofile,100,200. You can use unlimited as a value too
- **rootfs_propagation** (String) The rootfs propagation, eg shared, slave or (default) private
- **runtime** (Block List) Actions to take place when the container is being started (see [below for nested schema](#nestedblock--runtime))
- **sysctl** (Map of String) The map of sysctl key value pairs that are set inside the container namespace
- **tmpfs** (List of String) A simpler interface to mount a tmpfs, like --tmpfs in Docker, taking /dest:opt1,opt2
- **uid** (Number) The user id of the process
- **uid_mappings** (Block List) (Experemental) uid mappings for user namespaces (see [below for nested schema](#nestedblock--uid_mappings))
- **uid_name** (String) The user name of the process
- **uts** (String) The uts namespace, either to a path, or if new is specified it will use a new namespace

<a id="nestedblock--gid_mappings"></a>
### Nested Schema for `gid_mappings`

Optional:

- **container_id** (Number) The starting GID in the container
- **host_id** (Number) The starting GID on the host to be mapped to 'ContainerID'
- **size** (Number) The number of IDs to be mapped


<a id="nestedblock--mounts"></a>
### Nested Schema for `mounts`

Required:

- **type** (List of String) The mount kind

Optional:

- **destination** (List of String) The destination path of the mount
- **options** (List of String) The fstab style mount options
- **source** (List of String) The source path of the mount


<a id="nestedblock--resources"></a>
### Nested Schema for `resources`

Optional:

- **block_io** (Block List, Max: 1) The BlockIO restriction configuration (see [below for nested schema](#nestedblock--resources--block_io))
- **cpu** (Block List, Max: 1) The CPU restriction configuration (see [below for nested schema](#nestedblock--resources--cpu))
- **devices** (Block List) This configures the device whitelist (see [below for nested schema](#nestedblock--resources--devices))
- **hugepage_limits** (Block List) Hugetlb limit (in bytes) (see [below for nested schema](#nestedblock--resources--hugepage_limits))
- **memory** (Block List, Max: 1) The memory restriction configuration (see [below for nested schema](#nestedblock--resources--memory))
- **network** (Block List, Max: 1) The network restriction configuration (see [below for nested schema](#nestedblock--resources--network))
- **pids** (Block List, Max: 1) The task resource restriction configuration (see [below for nested schema](#nestedblock--resources--pids))

<a id="nestedblock--resources--block_io"></a>
### Nested Schema for `resources.block_io`

Optional:

- **leaf_weight** (Number) The tasks' weight in the given cgroup while competing with the cgroup's child cgroups, CFQ scheduler only
- **throttle_read_bps_device** (Block List) IO read rate limit per cgroup per device, bytes per second (see [below for nested schema](#nestedblock--resources--block_io--throttle_read_bps_device))
- **throttle_read_iops_device** (Block List) IO read rate limit per cgroup per device, IO per second (see [below for nested schema](#nestedblock--resources--block_io--throttle_read_iops_device))
- **throttle_write_bps_device** (Block List) IO write rate limit per cgroup per device, bytes per second (see [below for nested schema](#nestedblock--resources--block_io--throttle_write_bps_device))
- **throttle_write_iops_device** (Block List) IO read rate limit per cgroup per device, IO per second (see [below for nested schema](#nestedblock--resources--block_io--throttle_write_iops_device))
- **weight** (Number) The per cgroup weight
- **weight_device** (Block List) The weight per cgroup per device, can override BlkioWeight (see [below for nested schema](#nestedblock--resources--block_io--weight_device))

<a id="nestedblock--resources--block_io--throttle_read_bps_device"></a>
### Nested Schema for `resources.block_io.throttle_read_bps_device`

Optional:

- **rate** (Number) The IO rate limit per cgroup per device


<a id="nestedblock--resources--block_io--throttle_read_iops_device"></a>
### Nested Schema for `resources.block_io.throttle_read_iops_device`

Optional:

- **rate** (Number) The IO rate limit per cgroup per device


<a id="nestedblock--resources--block_io--throttle_write_bps_device"></a>
### Nested Schema for `resources.block_io.throttle_write_bps_device`

Optional:

- **rate** (Number) The IO rate limit per cgroup per device


<a id="nestedblock--resources--block_io--throttle_write_iops_device"></a>
### Nested Schema for `resources.block_io.throttle_write_iops_device`

Optional:

- **rate** (Number) The IO rate limit per cgroup per device


<a id="nestedblock--resources--block_io--weight_device"></a>
### Nested Schema for `resources.block_io.weight_device`

Optional:

- **leaf_weight** (Number) The tasks' weight in the given cgroup while competing with the cgroup's child cgroups, CFQ scheduler only
- **weight** (Number) The weight is the bandwidth rate for the device



<a id="nestedblock--resources--cpu"></a>
### Nested Schema for `resources.cpu`

Optional:

- **cpus** (String) CPUs to use within the cpuset
- **mems** (String) List of memory nodes in the cpuset
- **period** (Number) CPU period to be used for hardcapping in usecs
- **quota** (Number) CPU hardcap limit in usecs
- **realtime_period** (Number) CPU period to be used for realtime scheduling in usecs
- **realtime_runtime** (Number) How much time realtime scheduling may use in usecs
- **shares** (Number) CPU shares (relative weight vs. other cgroups with cpu shares)


<a id="nestedblock--resources--devices"></a>
### Nested Schema for `resources.devices`

Required:

- **allow** (Boolean) Allow or deny device

Optional:

- **access** (String) Cgroup access permissions format, rwm
- **major** (Number) The device's major number
- **minor** (Number) The device's minor number
- **type** (String) The device type, block, char, etc


<a id="nestedblock--resources--hugepage_limits"></a>
### Nested Schema for `resources.hugepage_limits`

Optional:

- **limit** (Number) The limit of 'hugepagesize' hugetlb usage
- **page_size** (String) The hugepage size


<a id="nestedblock--resources--memory"></a>
### Nested Schema for `resources.memory`

Optional:

- **disable_oom_killer** (Boolean) This disables the OOM killer for out of memory conditions
- **kernel** (Number) The kernel memory limit (in bytes)
- **kernel_tcp** (Number) The kernel memory limit for tcp (in bytes)
- **limit** (Number) The memory limit (in bytes)
- **reservation** (Number) The memory reservation or soft_limit (in bytes)
- **swap** (Number) The total memory limit (memory + swap)
- **swappiness** (Number) How aggressive the kernel will swap memory pages


<a id="nestedblock--resources--network"></a>
### Nested Schema for `resources.network`

Optional:

- **class_id** (Number) The class identifier for container's network packets
- **priorities** (Block List) The priority of network traffic for container (see [below for nested schema](#nestedblock--resources--network--priorities))

<a id="nestedblock--resources--network--priorities"></a>
### Nested Schema for `resources.network.priorities`

Optional:

- **name** (String) The name of the network interface
- **priority** (Number) The priority for the interface



<a id="nestedblock--resources--pids"></a>
### Nested Schema for `resources.pids`

Optional:

- **limit** (Number) Maximum number of PIDs



<a id="nestedblock--runtime"></a>
### Nested Schema for `runtime`

Optional:

- **bind_ns** (Block List) Specifies a namespace type and a path where the namespace from the container being created will be bound. This allows a namespace to be set up in an onboot container, and then using net: path for a service container to use that network namespace later (see [below for nested schema](#nestedblock--runtime--bind_ns))
- **cgroups** (List of String) A list of cgroups that will be created before the container is run
- **interfaces** (Block List) A list of mount specifications (see [below for nested schema](#nestedblock--runtime--interfaces))
- **mkdir** (List of String) A list of directories to create at runtime, in the root mount namespace
- **mounts** (Block List) A list of mount specifications (see [below for nested schema](#nestedblock--runtime--mounts))
- **namespace** (String) Overrides the LinuxKit default containerd namespace to put the container in; only applicable to services

<a id="nestedblock--runtime--bind_ns"></a>
### Nested Schema for `runtime.bind_ns`

Optional:

- **cgroups** (String)
- **ipc** (String)
- **mnt** (String)
- **net** (String)
- **pid** (String)
- **user** (String)
- **uts** (String)


<a id="nestedblock--runtime--interfaces"></a>
### Nested Schema for `runtime.interfaces`

Optional:

- **add** (String) The type of interface to be created in the containers namespace, with the specified name
- **create_in_root** (Boolean) The interface being added should be created in the root namespace first, then moved. This is needed for wireguard interfaces
- **name** (String) The name of an interface. An existing interface with this name will be moved into the container's network namespace
- **peer** (String) The name of the other end when creating a veth interface. This end will remain in the root namespace, where it can be attached to a bridge. Specifying this implies add: veth


<a id="nestedblock--runtime--mounts"></a>
### Nested Schema for `runtime.mounts`

Optional:

- **destination** (String) The destination directory to mount onto
- **options** (List of String) The options to use when mounting the directory
- **source** (String) The source for the directory you want to mount
- **type** (String) The type of the mount



<a id="nestedblock--uid_mappings"></a>
### Nested Schema for `uid_mappings`

Optional:

- **container_id** (Number) The starting UID in the container
- **host_id** (Number) The starting UID on the host to be mapped to 'ContainerID'
- **size** (Number) The number of IDs to be mapped


