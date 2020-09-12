package azure

import "github.com/blushft/go-diagrams/node"

type computeContainer struct {
	path string
	opts []node.Option
}

var Compute = &computeContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/compute",
}

func (c *computeContainer) ContainerRegistries(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/container-registries.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) FunctionApps(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/function-apps.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ServiceFabricClusters(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/service-fabric-clusters.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) VmImages(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/vm-images.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) VmLinux(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/vm-linux.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) AvailabilitySets(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/availability-sets.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) BatchAccounts(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/batch-accounts.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) CloudsimpleVirtualMachines(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/cloudsimple-virtual-machines.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) MeshApplications(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/mesh-applications.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) VmWindows(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/vm-windows.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Vm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/vm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) CloudServicesClassic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/cloud-services-classic.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) CloudServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/cloud-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) KubernetesServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/kubernetes-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) ContainerInstances(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/container-instances.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) Disks(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/disks.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) SapHanaOnAzure(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/sap-hana-on-azure.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) VmClassic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/vm-classic.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) CitrixVirtualDesktopsEssentials(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/citrix-virtual-desktops-essentials.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *computeContainer) DiskSnapshots(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/compute/disk-snapshots.png")}, c.opts, opts)
	return node.New(nopts...)
}
