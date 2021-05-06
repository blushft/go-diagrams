package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/azure/compute"}

func (c *computeContainer) VmWindows(opts ...attr.Attribute) *node.Node {
	return node.New("vm-windows", attr.AssetImage("assets/azure/compute/vm-windows.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Vm(opts ...attr.Attribute) *node.Node {
	return node.New("vm", attr.AssetImage("assets/azure/compute/vm.png"), attr.Shape(attr.None))
}

func (c *computeContainer) CloudsimpleVirtualMachines(opts ...attr.Attribute) *node.Node {
	return node.New("cloudsimple-virtual-machines", attr.AssetImage("assets/azure/compute/cloudsimple-virtual-machines.png"), attr.Shape(attr.None))
}

func (c *computeContainer) DiskSnapshots(opts ...attr.Attribute) *node.Node {
	return node.New("disk-snapshots", attr.AssetImage("assets/azure/compute/disk-snapshots.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Disks(opts ...attr.Attribute) *node.Node {
	return node.New("disks", attr.AssetImage("assets/azure/compute/disks.png"), attr.Shape(attr.None))
}

func (c *computeContainer) MeshApplications(opts ...attr.Attribute) *node.Node {
	return node.New("mesh-applications", attr.AssetImage("assets/azure/compute/mesh-applications.png"), attr.Shape(attr.None))
}

func (c *computeContainer) VmLinux(opts ...attr.Attribute) *node.Node {
	return node.New("vm-linux", attr.AssetImage("assets/azure/compute/vm-linux.png"), attr.Shape(attr.None))
}

func (c *computeContainer) CitrixVirtualDesktopsEssentials(opts ...attr.Attribute) *node.Node {
	return node.New("citrix-virtual-desktops-essentials", attr.AssetImage("assets/azure/compute/citrix-virtual-desktops-essentials.png"), attr.Shape(attr.None))
}

func (c *computeContainer) CloudServices(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-services", attr.AssetImage("assets/azure/compute/cloud-services.png"), attr.Shape(attr.None))
}

func (c *computeContainer) FunctionApps(opts ...attr.Attribute) *node.Node {
	return node.New("function-apps", attr.AssetImage("assets/azure/compute/function-apps.png"), attr.Shape(attr.None))
}

func (c *computeContainer) VmClassic(opts ...attr.Attribute) *node.Node {
	return node.New("vm-classic", attr.AssetImage("assets/azure/compute/vm-classic.png"), attr.Shape(attr.None))
}

func (c *computeContainer) CloudServicesClassic(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-services-classic", attr.AssetImage("assets/azure/compute/cloud-services-classic.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ContainerInstances(opts ...attr.Attribute) *node.Node {
	return node.New("container-instances", attr.AssetImage("assets/azure/compute/container-instances.png"), attr.Shape(attr.None))
}

func (c *computeContainer) VmImages(opts ...attr.Attribute) *node.Node {
	return node.New("vm-images", attr.AssetImage("assets/azure/compute/vm-images.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ServiceFabricClusters(opts ...attr.Attribute) *node.Node {
	return node.New("service-fabric-clusters", attr.AssetImage("assets/azure/compute/service-fabric-clusters.png"), attr.Shape(attr.None))
}

func (c *computeContainer) AvailabilitySets(opts ...attr.Attribute) *node.Node {
	return node.New("availability-sets", attr.AssetImage("assets/azure/compute/availability-sets.png"), attr.Shape(attr.None))
}

func (c *computeContainer) BatchAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("batch-accounts", attr.AssetImage("assets/azure/compute/batch-accounts.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ContainerRegistries(opts ...attr.Attribute) *node.Node {
	return node.New("container-registries", attr.AssetImage("assets/azure/compute/container-registries.png"), attr.Shape(attr.None))
}

func (c *computeContainer) KubernetesServices(opts ...attr.Attribute) *node.Node {
	return node.New("kubernetes-services", attr.AssetImage("assets/azure/compute/kubernetes-services.png"), attr.Shape(attr.None))
}

func (c *computeContainer) SapHanaOnAzure(opts ...attr.Attribute) *node.Node {
	return node.New("sap-hana-on-azure", attr.AssetImage("assets/azure/compute/sap-hana-on-azure.png"), attr.Shape(attr.None))
}
