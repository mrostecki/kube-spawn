/*
Copyright 2017 Kinvolk GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

const (
	RuntimeDocker = "docker"
	RuntimeRkt    = "rkt"
	RuntimeCrio   = "crio"
)

// ClusterConfiguration holds the state of a cluster
// with all the information needed to
// run kube-spawn
//
type ClusterConfiguration struct {
	KubeSpawnDir string `toml:"dir" mapstructure:"dir"`

	CNIPluginDir string `toml:"cni-plugin-dir" mapstructure:"cni-plugin-dir"`

	Name              string `toml:"cluster-name" mapstructure:"cluster-name"`
	ContainerRuntime  string `toml:"container-runtime" mapstructure:"container-runtime"`
	KubernetesVersion string `toml:"kubernetes-version" mapstructure:"kubernetes-version"`
	Image             string `toml:"image" mapstructure:"image"`
	Nodes             int    `toml:"nodes" mapstructure:"nodes"`

	// DevCluster indicates if we should run
	// from a local kubernetes build
	DevCluster   bool   `toml:"dev" mapstructure:"dev"`
	HyperkubeTag string `toml:"hyperkube-tag" mapstructure:"hyperkube-tag"`

	RuntimeConfiguration RuntimeConfiguration `toml:"runtime-config,omitempty" mapstructure:"runtime-config"`

	// Files to be copied from host to overlay of all machines
	Copymap []Pathmap `toml:"-" mapstructure:"copymap"`

	// TODO: rename Bindmap
	Bindmount BindmountConfiguration `toml:"bindmount" comment:"syntax: <dst in container> = <src on host>" mapstructure:"bindmount"`

	// Token is the token generated on kubeadm init
	// used to join more workers to the cluster
	Token             string `toml:"token, omitempty" mapstructure:"token"`
	TokenGroupsOption string `toml:"token-groups-option, omitempty" mapstructure:"token-groups-option"`

	Machines []MachineConfiguration `toml:"machines, omitempty" mapstructure:"machines"`
}

// RuntimeConfiguration holds the variables for
// running the cluster with a container runtime
//
type RuntimeConfiguration struct {
	Endpoint              string `toml:"endpoint,omitempty" mapstructure:"endpoint"`
	Timeout               string `toml:"timeout" mapstructure:"timeout"`
	UseLegacyCgroupDriver bool   `toml:"use-legacy-cgroup-driver" mapstructure:"use-legacy-cgroup-driver"`
	CgroupPerQos          bool   `toml:"cgroup-per-qos" mapstructure:"cgroup-per-qos"`
	FailSwapOn            bool   `toml:"fail-swap-on" mapstructure:"fail-swap-on"`

	Rkt  RuntimeConfigurationRkt  `toml:"-" mapstructure:"rkt"`
	Crio RuntimeConfigurationCrio `toml:"-" mapstructure:"crio"`
}

type RuntimeConfigurationRkt struct {
	RktBin      string `toml:"rkt-bin" mapstructure:"rkt-bin"`
	Stage1Image string `toml:"stage1-image" mapstructure:"stage1-image"`
	RktletBin   string `toml:"rktlet-bin" mapstructure:"rktlet-bin"`
}

type RuntimeConfigurationCrio struct {
	CrioBin   string `toml:"crio-bin" mapstructure:"crio-bin"`
	RuncBin   string `toml:"runc-bin" mapstructure:"runc-bin"`
	ConmonBin string `toml:"conmon-bin" mapstructure:"conmon-bin"`
}

// BindmountConfiguration contains the bind args
// for systemd-nspawn
type BindmountConfiguration struct {
	ReadOnly  []Pathmap `toml:"read-only" mapstructure:"read-only"`
	ReadWrite []Pathmap `toml:"read-write" mapstructure:"read-write"`
}

type MachineConfiguration struct {
	Running   bool                   `toml:"running" comment:"autogenerated. do not edit!" mapstructure:"running"`
	Name      string                 `toml:"name" comment:"autogenerated. do not edit!" mapstructure:"name"`
	IP        string                 `toml:"ip" comment:"autogenerated. do not edit!" mapstructure:"ip"`
	Bindmount BindmountConfiguration `toml:"bindmount" comment:"syntax: <dst in container> = <src on host>" mapstructure:"bindmount"`
}

// Pathmap represents mapping a path on one machine
// to a path on another machine
type Pathmap struct {
	Src string `toml:"src" mapstructure:"src"`
	Dst string `toml:"dst" mapstructure:"dst"`
}
