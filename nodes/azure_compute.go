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
	opts = append(opts, attr.AssetImage("assets/azure/compute/vm-windows.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vm-windows", opts...)
}

func (c *azureComputeContainer) ContainerRegistries(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/container-registries.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("container-registries", opts...)
}

func (c *azureComputeContainer) FunctionApps(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/function-apps.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("function-apps", opts...)
}

func (c *azureComputeContainer) VmClassic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/vm-classic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vm-classic", opts...)
}

func (c *azureComputeContainer) VmImages(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/vm-images.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vm-images", opts...)
}

func (c *azureComputeContainer) AvailabilitySets(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/availability-sets.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("availability-sets", opts...)
}

func (c *azureComputeContainer) CloudsimpleVirtualMachines(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/cloudsimple-virtual-machines.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudsimple-virtual-machines", opts...)
}

func (c *azureComputeContainer) ContainerInstances(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/container-instances.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("container-instances", opts...)
}

func (c *azureComputeContainer) DiskSnapshots(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/disk-snapshots.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("disk-snapshots", opts...)
}

func (c *azureComputeContainer) Disks(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/disks.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("disks", opts...)
}

func (c *azureComputeContainer) SapHanaOnAzure(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/sap-hana-on-azure.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sap-hana-on-azure", opts...)
}

func (c *azureComputeContainer) Vm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/vm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vm", opts...)
}

func (c *azureComputeContainer) BatchAccounts(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/batch-accounts.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("batch-accounts", opts...)
}

func (c *azureComputeContainer) CloudServicesClassic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/cloud-services-classic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-services-classic", opts...)
}

func (c *azureComputeContainer) CloudServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/cloud-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-services", opts...)
}

func (c *azureComputeContainer) KubernetesServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/kubernetes-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kubernetes-services", opts...)
}

func (c *azureComputeContainer) MeshApplications(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/mesh-applications.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mesh-applications", opts...)
}

func (c *azureComputeContainer) ServiceFabricClusters(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/service-fabric-clusters.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("service-fabric-clusters", opts...)
}

func (c *azureComputeContainer) VmLinux(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/vm-linux.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vm-linux", opts...)
}

func (c *azureComputeContainer) CitrixVirtualDesktopsEssentials(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/compute/citrix-virtual-desktops-essentials.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("citrix-virtual-desktops-essentials", opts...)
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
