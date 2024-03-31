package azure

import "github.com/blushft/go-diagrams/diagram"

type computeContainer struct {
	path string
	opts []diagram.NodeOption
}

var Compute = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/compute",
}

func (c *computeContainer) CitrixVirtualDesktopsEssentials(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/citrix-virtual-desktops-essentials.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Disks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/disks.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) FunctionApps(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/function-apps.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ContainerInstances(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/container-instances.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) VmLinux(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/vm-linux.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) Vm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/vm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) CloudServicesClassic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/cloud-services-classic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) CloudServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/cloud-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) CloudsimpleVirtualMachines(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/cloudsimple-virtual-machines.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) SapHanaOnAzure(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/sap-hana-on-azure.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ServiceFabricClusters(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/service-fabric-clusters.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) VmClassic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/vm-classic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) VmWindows(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/vm-windows.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) BatchAccounts(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/batch-accounts.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) ContainerRegistries(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/container-registries.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) MeshApplications(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/mesh-applications.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) VmImages(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/vm-images.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) AvailabilitySets(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/availability-sets.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) DiskSnapshots(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/disk-snapshots.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *computeContainer) KubernetesServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/compute/kubernetes-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
