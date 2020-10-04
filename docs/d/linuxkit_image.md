<!--- autogenerated do not edit --->
# linuxkit_image

## Argument Reference

The following arguments are supported:

* `additional_gid_names` - _list_ (Optional)  A list of additional group names for the process
* `additional_gids` - _list_ (Optional)  A list of additional groups for the process
* `ambient` - _list_ (Optional)  The Linux ambient capabilities (capabilities passed to non root users) that are required
* `annotations` - _list_ (Optional)  The map of key value pairs as OCI metadata
* `binds` - _list_ (Optional)  A Simpler interface to specify bind mounts, accepting a string like /src:/dest:opt1,opt2
* `capabilities` - _list_ (Optional)  The Linux capabilities required, for example CAP_SYS_ADMIN, if there is a single capability 'all' then all capabilities are added
* `cgroups_path` - _string_ (Optional)  The path for cgroups
* `command` - _list_ (Optional)  This will override the command and entrypoint in the image with a new list of commands
* `cwd` - _string_ (Optional)  The working directory, defaults to /
* `env` - _list_ (Optional)  This will override the environment in the image with a new environment list. Specify variables as VAR=value
* `gid` - _number_ (Optional)  The group id of the process
* `gid_mappings` - _[gid_mappings](#gid_mappings)_ (Optional)  (Experemental) gid mappings for user namespaces
* `gid_name` - _string_ (Optional)  The group name of the process
* `hostname` - _string_ (Optional)  The hostname inside the image
* `image` - _string_ (Required)  The Docker image to use for the filesystem
* `ipc` - _string_ (Optional)  The ipc namespace, either to a path, or if new is specified it will use a new namespace
* `masked_paths` - _list_ (Optional)  The paths which should be hidden
* `mounts` - _[mounts](#mounts)_ (Optional)  The full form for specifying a mount, which requires type, source, destination and a list of options
* `name` - _string_ (Required)  A unique name for the program being executed, used as the containerd id
* `net` - _string_ (Optional)  The network namespace, either to a path, or if none or new is specified it will use a new namespace
* `no_new_privileges` - _bool_ (Optional)  If set to true means no additional capabilities can be acquired and suid binaries do not work
* `oom_score_adj` - _number_ (Optional)  This changes the OOM score
* `pid` - _string_ (Optional)  The pid namespace, either to a path, or if host is specified it will use the host namespace
* `readonly` - _bool_ (Optional)  The root filesystem to read only, and changes the other default filesystems to read only
* `readonly_paths` - _list_ (Optional)  The paths which should be read only
* `resources` - _[resources](#resources)_ (Optional)  The cgroup resource limits as per the OCI spec
* `rlimits` - _list_ (Optional)  The list of rlimit values in the form name,soft,hard, eg nofile,100,200. You can use unlimited as a value too
* `rootfs_propagation` - _string_ (Optional)  The rootfs propagation, eg shared, slave or (default) private
* `runtime` - _[runtime](#runtime)_ (Optional)  Actions to take place when the container is being started
* `sysctl` - _map_ (Optional)  The map of sysctl key value pairs that are set inside the container namespace
* `tmpfs` - _list_ (Optional)  A simpler interface to mount a tmpfs, like --tmpfs in Docker, taking /dest:opt1,opt2
* `uid` - _number_ (Optional)  The user id of the process
* `uid_mappings` - _[uid_mappings](#uid_mappings)_ (Optional)  (Experemental) uid mappings for user namespaces
* `uid_name` - _string_ (Optional)  The user name of the process
* `uts` - _string_ (Optional)  The uts namespace, either to a path, or if new is specified it will use a new namespace


## Attributes Reference

No additional attributes are exported by this resource.



## Block Reference

Below is the documentation for the argument/attribute blocks in use by this resource:

### bind_ns
* `cgroups` - _string_ (Optional)  
* `ipc` - _string_ (Optional)  
* `mnt` - _string_ (Optional)  
* `net` - _string_ (Optional)  
* `pid` - _string_ (Optional)  
* `user` - _string_ (Optional)  
* `uts` - _string_ (Optional)  


### block_io
* `leaf_weight` - _number_ (Optional)  The tasks' weight in the given cgroup while competing with the cgroup's child cgroups, CFQ scheduler only
* `throttle_read_bps_device` - _[throttle_read_bps_device](#throttle_read_bps_device)_ (Optional)  IO read rate limit per cgroup per device, bytes per second
* `throttle_read_iops_device` - _[throttle_read_iops_device](#throttle_read_iops_device)_ (Optional)  IO read rate limit per cgroup per device, IO per second
* `throttle_write_bps_device` - _[throttle_write_bps_device](#throttle_write_bps_device)_ (Optional)  IO write rate limit per cgroup per device, bytes per second
* `throttle_write_iops_device` - _[throttle_write_iops_device](#throttle_write_iops_device)_ (Optional)  IO read rate limit per cgroup per device, IO per second
* `weight` - _number_ (Optional)  The per cgroup weight
* `weight_device` - _[weight_device](#weight_device)_ (Optional)  The weight per cgroup per device, can override BlkioWeight


### cpu
* `cpus` - _string_ (Optional)  CPUs to use within the cpuset
* `mems` - _string_ (Optional)  List of memory nodes in the cpuset
* `period` - _number_ (Optional)  CPU period to be used for hardcapping in usecs
* `quota` - _number_ (Optional)  CPU hardcap limit in usecs
* `realtime_period` - _number_ (Optional)  CPU period to be used for realtime scheduling in usecs
* `realtime_runtime` - _number_ (Optional)  How much time realtime scheduling may use in usecs
* `shares` - _number_ (Optional)  CPU shares (relative weight vs. other cgroups with cpu shares)


### devices
* `access` - _string_ (Optional)  Cgroup access permissions format, rwm
* `allow` - _bool_ (Required)  Allow or deny device
* `major` - _number_ (Optional)  The device's major number
* `minor` - _number_ (Optional)  The device's minor number
* `type` - _string_ (Optional)  The device type, block, char, etc


### gid_mappings
* `container_id` - _number_ (Optional)  The starting GID in the container
* `host_id` - _number_ (Optional)  The starting GID on the host to be mapped to 'ContainerID'
* `size` - _number_ (Optional)  The number of IDs to be mapped


### hugepage_limits
* `limit` - _number_ (Optional)  The limit of 'hugepagesize' hugetlb usage
* `page_size` - _string_ (Optional)  The hugepage size


### interfaces
* `add` - _string_ (Optional)  The type of interface to be created in the containers namespace, with the specified name
* `create_in_root` - _bool_ (Optional)  The interface being added should be created in the root namespace first, then moved. This is needed for wireguard interfaces
* `name` - _string_ (Optional)  The name of an interface. An existing interface with this name will be moved into the container's network namespace
* `peer` - _string_ (Optional)  The name of the other end when creating a veth interface. This end will remain in the root namespace, where it can be attached to a bridge. Specifying this implies add: veth


### memory
* `disable_oom_killer` - _bool_ (Optional)  This disables the OOM killer for out of memory conditions
* `kernel` - _number_ (Optional)  The kernel memory limit (in bytes)
* `kernel_tcp` - _number_ (Optional)  The kernel memory limit for tcp (in bytes)
* `limit` - _number_ (Optional)  The memory limit (in bytes)
* `reservation` - _number_ (Optional)  The memory reservation or soft_limit (in bytes)
* `swap` - _number_ (Optional)  The total memory limit (memory + swap)
* `swappiness` - _number_ (Optional)  How aggressive the kernel will swap memory pages


### mounts
* `destination` - _list_ (Optional)  The destination path of the mount
* `options` - _list_ (Optional)  The fstab style mount options
* `source` - _list_ (Optional)  The source path of the mount
* `type` - _list_ (Required)  The mount kind


### mounts
* `destination` - _string_ (Optional)  The destination directory to mount onto
* `options` - _list_ (Optional)  The options to use when mounting the directory
* `source` - _string_ (Optional)  The source for the directory you want to mount
* `type` - _string_ (Optional)  The type of the mount


### network
* `class_id` - _number_ (Optional)  The class identifier for container's network packets
* `priorities` - _[priorities](#priorities)_ (Optional)  The priority of network traffic for container


### pids
* `limit` - _number_ (Optional)  Maximum number of PIDs


### priorities
* `name` - _string_ (Optional)  The name of the network interface
* `priority` - _number_ (Optional)  The priority for the interface


### resources
* `block_io` - _[block_io](#block_io)_ (Optional)  The BlockIO restriction configuration
* `cpu` - _[cpu](#cpu)_ (Optional)  The CPU restriction configuration
* `devices` - _[devices](#devices)_ (Optional)  This configures the device whitelist
* `hugepage_limits` - _[hugepage_limits](#hugepage_limits)_ (Optional)  Hugetlb limit (in bytes)
* `memory` - _[memory](#memory)_ (Optional)  The memory restriction configuration
* `network` - _[network](#network)_ (Optional)  The network restriction configuration
* `pids` - _[pids](#pids)_ (Optional)  The task resource restriction configuration


### runtime
* `bind_ns` - _[bind_ns](#bind_ns)_ (Optional)  Specifies a namespace type and a path where the namespace from the container being created will be bound. This allows a namespace to be set up in an onboot container, and then using net: path for a service container to use that network namespace later
* `cgroups` - _list_ (Optional)  A list of cgroups that will be created before the container is run
* `interfaces` - _[interfaces](#interfaces)_ (Optional)  A list of mount specifications
* `mkdir` - _list_ (Optional)  A list of directories to create at runtime, in the root mount namespace
* `mounts` - _[mounts](#mounts)_ (Optional)  A list of mount specifications 
* `namespace` - _string_ (Optional)  Overrides the LinuxKit default containerd namespace to put the container in; only applicable to services


### throttle_read_bps_device
* `rate` - _number_ (Optional)  The IO rate limit per cgroup per device


### throttle_read_iops_device
* `rate` - _number_ (Optional)  The IO rate limit per cgroup per device


### throttle_write_bps_device
* `rate` - _number_ (Optional)  The IO rate limit per cgroup per device


### throttle_write_iops_device
* `rate` - _number_ (Optional)  The IO rate limit per cgroup per device


### uid_mappings
* `container_id` - _number_ (Optional)  The starting UID in the container
* `host_id` - _number_ (Optional)  The starting UID on the host to be mapped to 'ContainerID'
* `size` - _number_ (Optional)  The number of IDs to be mapped


### weight_device
* `leaf_weight` - _number_ (Optional)  The tasks' weight in the given cgroup while competing with the cgroup's child cgroups, CFQ scheduler only
* `weight` - _number_ (Optional)  The weight is the bandwidth rate for the device

