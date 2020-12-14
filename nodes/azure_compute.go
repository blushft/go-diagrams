package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azureComputeContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureAzureCompute = &azureComputeContainer{path: "assets/azure/compute"}

func (c *azureComputeContainer) VmWindows(opts ...attr.Attribute) *node.Node {
	return node.New("vm-windows", attr.AssetImage("assets/azure/compute/vm-windows.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) ContainerRegistries(opts ...attr.Attribute) *node.Node {
	return node.New("container-registries", attr.AssetImage("assets/azure/compute/container-registries.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) FunctionApps(opts ...attr.Attribute) *node.Node {
	return node.New("function-apps", attr.AssetImage("assets/azure/compute/function-apps.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) VmClassic(opts ...attr.Attribute) *node.Node {
	return node.New("vm-classic", attr.AssetImage("assets/azure/compute/vm-classic.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) VmImages(opts ...attr.Attribute) *node.Node {
	return node.New("vm-images", attr.AssetImage("assets/azure/compute/vm-images.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) AvailabilitySets(opts ...attr.Attribute) *node.Node {
	return node.New("availability-sets", attr.AssetImage("assets/azure/compute/availability-sets.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) CloudsimpleVirtualMachines(opts ...attr.Attribute) *node.Node {
	return node.New("cloudsimple-virtual-machines", attr.AssetImage("assets/azure/compute/cloudsimple-virtual-machines.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) ContainerInstances(opts ...attr.Attribute) *node.Node {
	return node.New("container-instances", attr.AssetImage("assets/azure/compute/container-instances.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) DiskSnapshots(opts ...attr.Attribute) *node.Node {
	return node.New("disk-snapshots", attr.AssetImage("assets/azure/compute/disk-snapshots.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) Disks(opts ...attr.Attribute) *node.Node {
	return node.New("disks", attr.AssetImage("assets/azure/compute/disks.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) SapHanaOnAzure(opts ...attr.Attribute) *node.Node {
	return node.New("sap-hana-on-azure", attr.AssetImage("assets/azure/compute/sap-hana-on-azure.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) Vm(opts ...attr.Attribute) *node.Node {
	return node.New("vm", attr.AssetImage("assets/azure/compute/vm.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) BatchAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("batch-accounts", attr.AssetImage("assets/azure/compute/batch-accounts.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) CloudServicesClassic(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-services-classic", attr.AssetImage("assets/azure/compute/cloud-services-classic.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) CloudServices(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-services", attr.AssetImage("assets/azure/compute/cloud-services.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) KubernetesServices(opts ...attr.Attribute) *node.Node {
	return node.New("kubernetes-services", attr.AssetImage("assets/azure/compute/kubernetes-services.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) MeshApplications(opts ...attr.Attribute) *node.Node {
	return node.New("mesh-applications", attr.AssetImage("assets/azure/compute/mesh-applications.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) ServiceFabricClusters(opts ...attr.Attribute) *node.Node {
	return node.New("service-fabric-clusters", attr.AssetImage("assets/azure/compute/service-fabric-clusters.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) VmLinux(opts ...attr.Attribute) *node.Node {
	return node.New("vm-linux", attr.AssetImage("assets/azure/compute/vm-linux.png"), attr.Shape(attr.None))
}

func (c *azureComputeContainer) CitrixVirtualDesktopsEssentials(opts ...attr.Attribute) *node.Node {
	return node.New("citrix-virtual-desktops-essentials", attr.AssetImage("assets/azure/compute/citrix-virtual-desktops-essentials.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.compute"
  Register(namespace, "VmWindows", azureAzureCompute.VmWindows)
  Register(namespace, "ContainerRegistries", azureAzureCompute.ContainerRegistries)
  Register(namespace, "FunctionApps", azureAzureCompute.FunctionApps)
  Register(namespace, "VmClassic", azureAzureCompute.VmClassic)
  Register(namespace, "VmImages", azureAzureCompute.VmImages)
  Register(namespace, "AvailabilitySets", azureAzureCompute.AvailabilitySets)
  Register(namespace, "CloudsimpleVirtualMachines", azureAzureCompute.CloudsimpleVirtualMachines)
  Register(namespace, "ContainerInstances", azureAzureCompute.ContainerInstances)
  Register(namespace, "DiskSnapshots", azureAzureCompute.DiskSnapshots)
  Register(namespace, "Disks", azureAzureCompute.Disks)
  Register(namespace, "SapHanaOnAzure", azureAzureCompute.SapHanaOnAzure)
  Register(namespace, "Vm", azureAzureCompute.Vm)
  Register(namespace, "BatchAccounts", azureAzureCompute.BatchAccounts)
  Register(namespace, "CloudServicesClassic", azureAzureCompute.CloudServicesClassic)
  Register(namespace, "CloudServices", azureAzureCompute.CloudServices)
  Register(namespace, "KubernetesServices", azureAzureCompute.KubernetesServices)
  Register(namespace, "MeshApplications", azureAzureCompute.MeshApplications)
  Register(namespace, "ServiceFabricClusters", azureAzureCompute.ServiceFabricClusters)
  Register(namespace, "VmLinux", azureAzureCompute.VmLinux)
  Register(namespace, "CitrixVirtualDesktopsEssentials", azureAzureCompute.CitrixVirtualDesktopsEssentials)
}
