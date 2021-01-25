package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func imageDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "`linuxkit_image` is a single image to be included in the system.",

		Read: imageRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A unique name for the program being executed, used as the containerd id",
				Required:    true,
				ForceNew:    true,
			},
			"image": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The Docker image to use for the filesystem",
				Required:    true,
				ForceNew:    true,
			},
			"capabilities": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The Linux capabilities required, for example CAP_SYS_ADMIN, if there is a single capability 'all' then all capabilities are added",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"ambient": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The Linux ambient capabilities (capabilities passed to non root users) that are required",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"mounts": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The full form for specifying a mount, which requires type, source, destination and a list of options",
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The mount kind",
							Required:    true,
							ForceNew:    true,
						},
						"source": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The source path of the mount",
							Optional:    true,
							ForceNew:    true,
						},
						"destination": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The destination path of the mount",
							Optional:    true,
							ForceNew:    true,
						},
						"options": &schema.Schema{
							Type:        schema.TypeList,
							Description: "The fstab style mount options",
							Optional:    true,
							ForceNew:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"binds": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A Simpler interface to specify bind mounts, accepting a string like /src:/dest:opt1,opt2",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"tmpfs": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A simpler interface to mount a tmpfs, like --tmpfs in Docker, taking /dest:opt1,opt2",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"command": &schema.Schema{
				Type:        schema.TypeList,
				Description: "This will override the command and entrypoint in the image with a new list of commands",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"env": &schema.Schema{
				Type:        schema.TypeList,
				Description: "This will override the environment in the image with a new environment list. Specify variables as VAR=value",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cwd": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The working directory, defaults to /",
				Optional:    true,
				ForceNew:    true,
			},
			"net": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The network namespace, either to a path, or if none or new is specified it will use a new namespace",
				Optional:    true,
				ForceNew:    true,
			},
			"ipc": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ipc namespace, either to a path, or if new is specified it will use a new namespace",
				Optional:    true,
				ForceNew:    true,
			},
			"uts": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The uts namespace, either to a path, or if new is specified it will use a new namespace",
				Optional:    true,
				ForceNew:    true,
			},
			"pid": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The pid namespace, either to a path, or if host is specified it will use the host namespace",
				Optional:    true,
				ForceNew:    true,
			},
			"readonly": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "The root filesystem to read only, and changes the other default filesystems to read only",
				Optional:    true,
				ForceNew:    true,
			},
			"masked_paths": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The paths which should be hidden",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"readonly_paths": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The paths which should be read only",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"uid": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The user id of the process",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"gid": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The group id of the process",
				Optional:    true,
				ForceNew:    true,
			},
			"uid_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The user name of the process",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"gid_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The group name of the process",
				Optional:    true,
				ForceNew:    true,
			},
			"additional_gids": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A list of additional groups for the process",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"additional_gid_names": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A list of additional group names for the process",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"no_new_privileges": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If set to true means no additional capabilities can be acquired and suid binaries do not work",
				Optional:    true,
				ForceNew:    true,
			},
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The hostname inside the image",
				Optional:    true,
				ForceNew:    true,
			},
			"oom_score_adj": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "This changes the OOM score",
				Optional:    true,
				ForceNew:    true,
			},
			"rootfs_propagation": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The rootfs propagation, eg shared, slave or (default) private",
				Optional:    true,
				ForceNew:    true,
			},
			"cgroups_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The path for cgroups",
				Optional:    true,
				ForceNew:    true,
			},
			"resources": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The cgroup resource limits as per the OCI spec",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"devices": &schema.Schema{
							Type:        schema.TypeList,
							Description: "This configures the device whitelist",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "Allow or deny device",
										Required:    true,
										ForceNew:    true,
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The device type, block, char, etc",
										Optional:    true,
										ForceNew:    true,
									},
									"major": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The device's major number",
										Optional:    true,
										ForceNew:    true,
									},
									"minor": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The device's minor number",
										Optional:    true,
										ForceNew:    true,
									},
									"access": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Cgroup access permissions format, rwm",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"memory": &schema.Schema{
							Type:        schema.TypeList,
							Description: "The memory restriction configuration",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The memory limit (in bytes)",
										Optional:    true,
										ForceNew:    true,
									},
									"reservation": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The memory reservation or soft_limit (in bytes)",
										Optional:    true,
										ForceNew:    true,
									},
									"swap": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The total memory limit (memory + swap)",
										Optional:    true,
										ForceNew:    true,
									},
									"kernel": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The kernel memory limit (in bytes)",
										Optional:    true,
										ForceNew:    true,
									},
									"kernel_tcp": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The kernel memory limit for tcp (in bytes)",
										Optional:    true,
										ForceNew:    true,
									},
									"swappiness": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "How aggressive the kernel will swap memory pages",
										Optional:    true,
										ForceNew:    true,
									},
									"disable_oom_killer": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "This disables the OOM killer for out of memory conditions",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"cpu": &schema.Schema{
							Type:        schema.TypeList,
							Description: "The CPU restriction configuration",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"shares": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "CPU shares (relative weight vs. other cgroups with cpu shares)",
										Optional:    true,
										ForceNew:    true,
									},
									"quota": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "CPU hardcap limit in usecs",
										Optional:    true,
										ForceNew:    true,
									},
									"period": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "CPU period to be used for hardcapping in usecs",
										Optional:    true,
										ForceNew:    true,
									},
									"realtime_runtime": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "How much time realtime scheduling may use in usecs",
										Optional:    true,
										ForceNew:    true,
									},
									"realtime_period": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "CPU period to be used for realtime scheduling in usecs",
										Optional:    true,
										ForceNew:    true,
									},
									"cpus": &schema.Schema{
										Type:        schema.TypeString,
										Description: "CPUs to use within the cpuset",
										Optional:    true,
										ForceNew:    true,
									},
									"mems": &schema.Schema{
										Type:        schema.TypeString,
										Description: "List of memory nodes in the cpuset",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"pids": &schema.Schema{
							Type:        schema.TypeList,
							Description: "The task resource restriction configuration",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum number of PIDs",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"block_io": &schema.Schema{
							Type:        schema.TypeList,
							Description: "The BlockIO restriction configuration",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"weight": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The per cgroup weight",
										Optional:    true,
										ForceNew:    true,
									},
									"leaf_weight": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The tasks' weight in the given cgroup while competing with the cgroup's child cgroups, CFQ scheduler only",
										Optional:    true,
										ForceNew:    true,
									},
									"weight_device": &schema.Schema{
										Type:        schema.TypeList,
										Description: "The weight per cgroup per device, can override BlkioWeight",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"weight": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The weight is the bandwidth rate for the device",
													Optional:    true,
													ForceNew:    true,
												},
												"leaf_weight": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The tasks' weight in the given cgroup while competing with the cgroup's child cgroups, CFQ scheduler only",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"throttle_read_bps_device": &schema.Schema{
										Type:        schema.TypeList,
										Description: "IO read rate limit per cgroup per device, bytes per second",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rate": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The IO rate limit per cgroup per device",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"throttle_write_bps_device": &schema.Schema{
										Type:        schema.TypeList,
										Description: "IO write rate limit per cgroup per device, bytes per second",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rate": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The IO rate limit per cgroup per device",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"throttle_read_iops_device": &schema.Schema{
										Type:        schema.TypeList,
										Description: "IO read rate limit per cgroup per device, IO per second",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rate": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The IO rate limit per cgroup per device",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
									"throttle_write_iops_device": &schema.Schema{
										Type:        schema.TypeList,
										Description: "IO read rate limit per cgroup per device, IO per second",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rate": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The IO rate limit per cgroup per device",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
								},
							},
						},
						"hugepage_limits": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Hugetlb limit (in bytes)",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"page_size": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The hugepage size",
										Optional:    true,
										ForceNew:    true,
									},
									"limit": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The limit of 'hugepagesize' hugetlb usage",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"network": &schema.Schema{
							Type:        schema.TypeList,
							Description: "The network restriction configuration",
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_id": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The class identifier for container's network packets",
										Optional:    true,
										ForceNew:    true,
									},
									"priorities": &schema.Schema{
										Type:        schema.TypeList,
										Description: "The priority of network traffic for container",
										Optional:    true,
										ForceNew:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The name of the network interface",
													Optional:    true,
													ForceNew:    true,
												},
												"priority": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The priority for the interface",
													Optional:    true,
													ForceNew:    true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"sysctl": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "The map of sysctl key value pairs that are set inside the container namespace",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"rlimits": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The list of rlimit values in the form name,soft,hard, eg nofile,100,200. You can use unlimited as a value too",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"uid_mappings": &schema.Schema{
				Type:        schema.TypeList,
				Description: "(Experemental) uid mappings for user namespaces",
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "The starting UID on the host to be mapped to 'ContainerID'",
							Optional:    true,
							ForceNew:    true,
						},
						"container_id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "The starting UID in the container",
							Optional:    true,
							ForceNew:    true,
						},
						"size": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "The number of IDs to be mapped",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"gid_mappings": &schema.Schema{
				Type:        schema.TypeList,
				Description: "(Experemental) gid mappings for user namespaces",
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "The starting GID on the host to be mapped to 'ContainerID'",
							Optional:    true,
							ForceNew:    true,
						},
						"container_id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "The starting GID in the container",
							Optional:    true,
							ForceNew:    true,
						},
						"size": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "The number of IDs to be mapped",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"annotations": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The map of key value pairs as OCI metadata",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"runtime": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Actions to take place when the container is being started",
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cgroups": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A list of cgroups that will be created before the container is run",
							Optional:    true,
							ForceNew:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"mounts": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A list of mount specifications ",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The destination directory to mount onto",
										Optional:    true,
										ForceNew:    true,
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The type of the mount",
										Optional:    true,
										ForceNew:    true,
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The source for the directory you want to mount",
										Optional:    true,
										ForceNew:    true,
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Description: "The options to use when mounting the directory",
										Optional:    true,
										ForceNew:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"mkdir": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A list of directories to create at runtime, in the root mount namespace",
							Optional:    true,
							ForceNew:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"interfaces": &schema.Schema{
							Type:        schema.TypeList,
							Description: "A list of mount specifications",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The name of an interface. An existing interface with this name will be moved into the container's network namespace",
										Optional:    true,
										ForceNew:    true,
									},
									"add": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The type of interface to be created in the containers namespace, with the specified name",
										Optional:    true,
										ForceNew:    true,
									},
									"create_in_root": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "The interface being added should be created in the root namespace first, then moved. This is needed for wireguard interfaces",
										Optional:    true,
										ForceNew:    true,
									},
									"peer": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The name of the other end when creating a veth interface. This end will remain in the root namespace, where it can be attached to a bridge. Specifying this implies add: veth",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"bind_ns": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Specifies a namespace type and a path where the namespace from the container being created will be bound. This allows a namespace to be set up in an onboot container, and then using net: path for a service container to use that network namespace later",
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cgroups": &schema.Schema{
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"ipc": &schema.Schema{
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"mnt": &schema.Schema{
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"net": &schema.Schema{
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"pid": &schema.Schema{
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"user": &schema.Schema{
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
									"uts": &schema.Schema{
										Type:        schema.TypeString,
										Description: "",
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
						"namespace": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Overrides the LinuxKit default containerd namespace to put the container in; only applicable to services",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func imageRead(d *schema.ResourceData, meta interface{}) error {
	image := &moby.Image{}

	if v, ok := d.GetOk("name"); ok {
		image.Name = v.(string)
	}

	if v, ok := d.GetOk("image"); ok {
		image.Image = v.(string)
	}

	if v, ok := d.GetOk("capabilities"); ok {
		image.Capabilities = interfaceSliceToStringSliceRef(v.([]interface{}))
	}

	if v, ok := d.GetOk("ambient"); ok {
		image.Ambient = interfaceSliceToStringSliceRef(v.([]interface{}))
	}

	if v, ok := d.GetOk("mounts"); ok {
		image.Mounts = &[]specs.Mount{}

		for _, raw := range v.([]interface{}) {
			v := raw.(map[string]interface{})

			*image.Mounts = append(*image.Mounts, specs.Mount{
				Type:        v["type"].(string),
				Source:      v["source"].(string),
				Destination: v["destination"].(string),
				Options:     interfaceSliceToStringSlice(v["options"].([]interface{})),
			})
		}
	}

	if v, ok := d.GetOk("binds"); ok {
		image.Binds = interfaceSliceToStringSliceRef(v.([]interface{}))
	}

	if v, ok := d.GetOk("tmpfd"); ok {
		image.Tmpfs = interfaceSliceToStringSliceRef(v.([]interface{}))
	}

	if v, ok := d.GetOk("command"); ok {
		image.Command = interfaceSliceToStringSliceRef(v.([]interface{}))
	}

	if v, ok := d.GetOk("env"); ok {
		image.Env = interfaceSliceToStringSliceRef(v.([]interface{}))
	}

	if v, ok := d.GetOk("net"); ok {
		image.Net = v.(string)
	}

	if v, ok := d.GetOk("ipc"); ok {
		image.Ipc = v.(string)
	}

	if v, ok := d.GetOk("uts"); ok {
		image.Uts = v.(string)
	}

	if v, ok := d.GetOk("pid"); ok {
		image.Pid = v.(string)
	}

	if v, ok := d.GetOk("readonly"); ok {
		b := v.(bool)
		image.Readonly = &b
	}

	if v, ok := d.GetOk("readonly_paths"); ok {
		image.ReadonlyPaths = interfaceSliceToStringSliceRef(v.([]interface{}))
	}

	if v, ok := d.GetOk("uid"); ok {
		image.UID = &v
	}

	if v, ok := d.GetOk("gid"); ok {
		image.GID = &v
	}

	if v, ok := d.GetOk("uid_name"); ok {
		image.UID = &v
	}

	if v, ok := d.GetOk("gid_name"); ok {
		image.GID = &v
	}

	if v, ok := d.GetOk("additional_gids"); ok {
		i := v.([]interface{})
		image.AdditionalGids = &i
	}

	if v, ok := d.GetOk("additional_gid_names"); ok {
		i := v.([]interface{})
		if image.AdditionalGids == nil {
			image.AdditionalGids = &[]interface{}{}
		}

		*image.AdditionalGids = append(*image.AdditionalGids, i...)
	}

	if v, ok := d.GetOk("no_new_privileges"); ok {
		b := v.(bool)
		image.NoNewPrivileges = &b
	}

	if v, ok := d.GetOk("hostname"); ok {
		image.Hostname = v.(string)
	}

	if v, ok := d.GetOk("oom_score_adj"); ok {
		i := v.(int)
		image.OOMScoreAdj = &i
	}

	if v, ok := d.GetOk("rootfs_propagation"); ok {
		s := v.(string)
		image.RootfsPropagation = &s
	}

	if v, ok := d.GetOk("cgroups_path"); ok {
		s := v.(string)
		image.CgroupsPath = &s
	}

	if v, ok := d.GetOk("resources"); ok {
		for _, raw := range v.([]interface{}) {
			v := raw.(map[string]interface{})
			image.Resources = &specs.LinuxResources{}

			if v, ok := v["devices"]; ok {
				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})
					device := specs.LinuxDeviceCgroup{}

					if v, ok := v["allow"]; ok {
						device.Allow = v.(bool)
					}

					if v, ok := v["type"]; ok {
						device.Type = v.(string)
					}

					if v, ok := v["major"]; ok {
						i := int64(v.(int))
						device.Major = &i
					}

					if v, ok := v["minor"]; ok {
						i := int64(v.(int))
						device.Minor = &i
					}

					if v, ok := v["access"]; ok {
						device.Access = v.(string)
					}

					image.Resources.Devices = append(image.Resources.Devices, device)
				}
			}

			if v, ok := v["memory"]; ok {
				image.Resources.Memory = &specs.LinuxMemory{}

				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})

					if v, ok := v["limit"]; ok {
						i := int64(v.(int))
						image.Resources.Memory.Limit = &i
					}

					if v, ok := v["reservation"]; ok {
						i := int64(v.(int))
						image.Resources.Memory.Reservation = &i
					}

					if v, ok := v["swap"]; ok {
						i := int64(v.(int))
						image.Resources.Memory.Swap = &i
					}

					if v, ok := v["kernel"]; ok {
						i := int64(v.(int))
						image.Resources.Memory.Kernel = &i
					}

					if v, ok := v["kernel_tcp"]; ok {
						i := int64(v.(int))
						image.Resources.Memory.KernelTCP = &i
					}

					if v, ok := v["swappiness"]; ok {
						i := uint64(v.(int))
						image.Resources.Memory.Swappiness = &i
					}

					if v, ok := v["disable_oom_killer"]; ok {
						b := v.(bool)
						image.Resources.Memory.DisableOOMKiller = &b
					}
				}
			}

			if v, ok := v["cpu"]; ok {
				image.Resources.CPU = &specs.LinuxCPU{}

				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})

					if v, ok := v["shares"]; ok {
						i := uint64(v.(int))
						image.Resources.CPU.Shares = &i
					}

					if v, ok := v["quota"]; ok {
						i := int64(v.(int))
						image.Resources.CPU.Quota = &i
					}

					if v, ok := v["period"]; ok {
						i := uint64(v.(int))
						image.Resources.CPU.Shares = &i
					}

					if v, ok := v["realtime_runtime"]; ok {
						i := int64(v.(int))
						image.Resources.CPU.RealtimeRuntime = &i
					}

					if v, ok := v["realtime_period"]; ok {
						i := uint64(v.(int))
						image.Resources.CPU.RealtimePeriod = &i
					}

					if v, ok := v["cpus"]; ok {
						image.Resources.CPU.Cpus = v.(string)
					}

					if v, ok := v["mems"]; ok {
						image.Resources.CPU.Mems = v.(string)
					}
				}
			}

			if v, ok := v["pids"]; ok {
				image.Resources.Pids = &specs.LinuxPids{}

				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})

					if v, ok := v["limit"]; ok {
						image.Resources.Pids.Limit = int64(v.(int))
					}
				}
			}

			if v, ok := v["block_io"]; ok {
				image.Resources.BlockIO = &specs.LinuxBlockIO{}

				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})

					if v, ok := v["weight"]; ok {
						i := uint16(v.(int))
						image.Resources.BlockIO.Weight = &i
					}

					if v, ok := v["leaf_weight"]; ok {
						i := uint16(v.(int))
						image.Resources.BlockIO.LeafWeight = &i
					}

					if v, ok := v["weight_device"]; ok {
						for _, raw := range v.([]interface{}) {
							v := raw.(map[string]interface{})
							device := specs.LinuxWeightDevice{}

							if v, ok := v["weight"]; ok {
								i := uint16(v.(int))
								device.Weight = &i
							}

							if v, ok := v["leaf_weight"]; ok {
								i := uint16(v.(int))
								device.LeafWeight = &i
							}

							image.Resources.BlockIO.WeightDevice = append(image.Resources.BlockIO.WeightDevice, device)
						}
					}

					if v, ok := v["throttle_read_bps_device"]; ok {
						for _, raw := range v.([]interface{}) {
							v := raw.(map[string]interface{})
							throttle := specs.LinuxThrottleDevice{}

							if v, ok := v["rate"]; ok {
								throttle.Rate = uint64(v.(int))
							}

							image.Resources.BlockIO.ThrottleReadBpsDevice = append(image.Resources.BlockIO.ThrottleReadBpsDevice, throttle)
						}
					}

					if v, ok := v["throttle_write_bps_device"]; ok {
						for _, raw := range v.([]interface{}) {
							v := raw.(map[string]interface{})
							throttle := specs.LinuxThrottleDevice{}

							if v, ok := v["rate"]; ok {
								throttle.Rate = uint64(v.(int))
							}

							image.Resources.BlockIO.ThrottleWriteBpsDevice = append(image.Resources.BlockIO.ThrottleWriteBpsDevice, throttle)
						}
					}

					if v, ok := v["throttle_read_iops_device"]; ok {
						for _, raw := range v.([]interface{}) {
							v := raw.(map[string]interface{})
							throttle := specs.LinuxThrottleDevice{}

							if v, ok := v["rate"]; ok {
								throttle.Rate = uint64(v.(int))
							}

							image.Resources.BlockIO.ThrottleReadIOPSDevice = append(image.Resources.BlockIO.ThrottleReadIOPSDevice, throttle)
						}
					}

					if v, ok := v["throttle_write_iops_device"]; ok {
						for _, raw := range v.([]interface{}) {
							v := raw.(map[string]interface{})
							throttle := specs.LinuxThrottleDevice{}

							if v, ok := v["rate"]; ok {
								throttle.Rate = uint64(v.(int))
							}

							image.Resources.BlockIO.ThrottleWriteIOPSDevice = append(image.Resources.BlockIO.ThrottleWriteIOPSDevice, throttle)
						}
					}
				}
			}

			if v, ok := v["hugepage_limits"]; ok {
				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})
					limit := specs.LinuxHugepageLimit{}

					if v, ok := v["page_size"]; ok {
						limit.Pagesize = v.(string)
					}

					if v, ok := v["limit"]; ok {
						limit.Limit = uint64(v.(int))
					}

					image.Resources.HugepageLimits = append(image.Resources.HugepageLimits, limit)
				}
			}

			if v, ok := v["network"]; ok {
				image.Resources.Network = &specs.LinuxNetwork{}

				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})

					if v, ok := v["class_id"]; ok {
						i := uint32(v.(int))
						image.Resources.Network.ClassID = &i
					}

					if v, ok := v["priorities"]; ok {
						for _, raw := range v.([]interface{}) {
							v := raw.(map[string]interface{})
							priority := specs.LinuxInterfacePriority{}

							if v, ok := v["name"]; ok {
								priority.Name = v.(string)
							}

							if v, ok := v["priority"]; ok {
								priority.Priority = uint32(v.(int))
							}

							image.Resources.Network.Priorities = append(image.Resources.Network.Priorities, priority)
						}
					}
				}
			}
		}
	}

	if v, ok := d.GetOk("sysctl"); ok {
		image.Sysctl = interfaceMapToStringMapRef(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("rlimits"); ok {
		image.Rlimits = interfaceSliceToStringSliceRef(v.([]interface{}))
	}

	if v, ok := d.GetOk("uid_mappings"); ok {
		image.UIDMappings = &[]specs.LinuxIDMapping{}

		for _, raw := range v.([]interface{}) {
			v := raw.(map[string]interface{})
			idMapping := specs.LinuxIDMapping{}

			if v, ok := v["host_id"]; ok {
				idMapping.HostID = uint32(v.(int))
			}

			if v, ok := v["container_id"]; ok {
				idMapping.ContainerID = uint32(v.(int))
			}

			if v, ok := v["size"]; ok {
				idMapping.Size = uint32(v.(int))
			}

			*image.UIDMappings = append(*image.UIDMappings, idMapping)
		}
	}

	if v, ok := d.GetOk("gid_mappings"); ok {
		image.GIDMappings = &[]specs.LinuxIDMapping{}

		for _, raw := range v.([]interface{}) {
			v := raw.(map[string]interface{})
			idMapping := specs.LinuxIDMapping{}

			if v, ok := v["host_id"]; ok {
				idMapping.HostID = uint32(v.(int))
			}

			if v, ok := v["container_id"]; ok {
				idMapping.ContainerID = uint32(v.(int))
			}

			if v, ok := v["size"]; ok {
				idMapping.Size = uint32(v.(int))
			}

			*image.GIDMappings = append(*image.GIDMappings, idMapping)
		}
	}

	if v, ok := d.GetOk("annotations"); ok {
		image.Annotations = interfaceMapToStringMapRef(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("runtime"); ok {
		image.Runtime = &moby.Runtime{}

		for _, raw := range v.([]interface{}) {
			v := raw.(map[string]interface{})

			if v, ok := v["cgroups"]; ok {
				image.Runtime.Cgroups = &[]string{}

				for _, raw := range v.([]interface{}) {
					v := raw.(string)
					*image.Runtime.Cgroups = append(*image.Runtime.Cgroups, v)
				}
			}

			if v, ok := v["mounts"]; ok {
				image.Runtime.Mounts = &[]specs.Mount{}

				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})
					m := specs.Mount{}

					if v, ok := v["destination"]; ok {
						m.Destination = v.(string)
					}

					if v, ok := v["type"]; ok {
						m.Type = v.(string)
					}

					if v, ok := v["source"]; ok {
						m.Source = v.(string)
					}

					if v, ok := v["options"]; ok {
						for _, raw := range v.([]interface{}) {
							v := raw.(string)
							m.Options = append(m.Options, v)
						}
					}

					*image.Runtime.Mounts = append(*image.Runtime.Mounts, m)
				}
			}

			if v, ok := v["mkdir"]; ok {
				image.Runtime.Mkdir = &[]string{}

				for _, raw := range v.([]interface{}) {
					v := raw.(string)
					*image.Runtime.Mkdir = append(*image.Runtime.Mkdir, v)
				}
			}

			if v, ok := v["interfaces"]; ok {
				image.Runtime.Interfaces = &[]moby.Interface{}

				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})
					i := moby.Interface{}

					if v, ok := v["name"]; ok {
						i.Name = v.(string)
					}

					if v, ok := v["add"]; ok {
						i.Add = v.(string)
					}

					if v, ok := v["peer"]; ok {
						i.Peer = v.(string)
					}

					if v, ok := v["create_in_root"]; ok {
						i.CreateInRoot = v.(bool)
					}

					*image.Runtime.Interfaces = append(*image.Runtime.Interfaces, i)
				}
			}

			if v, ok := v["bind_ns"]; ok {
				for _, raw := range v.([]interface{}) {
					v := raw.(map[string]interface{})

					if v, ok := v["cgroups"]; ok {
						image.Runtime.BindNS.Cgroup = stringPtr(v.(string))
					}

					if v, ok := v["ipc"]; ok {
						image.Runtime.BindNS.Ipc = stringPtr(v.(string))
					}

					if v, ok := v["mnt"]; ok {
						image.Runtime.BindNS.Mnt = stringPtr(v.(string))
					}

					if v, ok := v["net"]; ok {
						image.Runtime.BindNS.Net = stringPtr(v.(string))
					}

					if v, ok := v["pid"]; ok {
						image.Runtime.BindNS.Pid = stringPtr(v.(string))
					}

					if v, ok := v["user"]; ok {
						image.Runtime.BindNS.User = stringPtr(v.(string))
					}

					if v, ok := v["uts"]; ok {
						image.Runtime.BindNS.Uts = stringPtr(v.(string))
					}
				}
			}

			if v, ok := v["namespace"]; ok {
				image.Runtime.Namespace = stringPtr(v.(string))
			}
		}
	}

	d.SetId(globalCache.addImage(image))

	return nil
}

func interfaceSliceToStringSlice(i []interface{}) []string {
	s := make([]string, len(i))
	for n, v := range i {
		s[n] = v.(string)
	}
	return s
}

func interfaceSliceToStringSliceRef(i []interface{}) *[]string {
	s := make([]string, len(i))
	for n, v := range i {
		s[n] = v.(string)
	}
	return &s
}

func interfaceMapToStringMapRef(i map[string]interface{}) *map[string]string {
	s := make(map[string]string, len(i))
	for k, v := range i {
		s[k] = v.(string)
	}
	return &s
}
