package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ociComputeContainer struct {
	path  string
	attrs []attr.Attribute
}

var OciCompute = &ociComputeContainer{path: "assets/oci/compute"}

func (c *ociComputeContainer) AutoscaleWhite(opts ...attr.Attribute) *node.Node {
	return node.New("autoscale-white", attr.AssetImage("assets/oci/compute/autoscale-white.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) FunctionsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("functions-white", attr.AssetImage("assets/oci/compute/functions-white.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) Oke(opts ...attr.Attribute) *node.Node {
	return node.New("oke", attr.AssetImage("assets/oci/compute/oke.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) Vm(opts ...attr.Attribute) *node.Node {
	return node.New("vm", attr.AssetImage("assets/oci/compute/vm.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) Autoscale(opts ...attr.Attribute) *node.Node {
	return node.New("autoscale", attr.AssetImage("assets/oci/compute/autoscale.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) InstancePools(opts ...attr.Attribute) *node.Node {
	return node.New("instance-pools", attr.AssetImage("assets/oci/compute/instance-pools.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) Ocir(opts ...attr.Attribute) *node.Node {
	return node.New("ocir", attr.AssetImage("assets/oci/compute/ocir.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) OkeWhite(opts ...attr.Attribute) *node.Node {
	return node.New("oke-white", attr.AssetImage("assets/oci/compute/oke-white.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) BmWhite(opts ...attr.Attribute) *node.Node {
	return node.New("bm-white", attr.AssetImage("assets/oci/compute/bm-white.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) Bm(opts ...attr.Attribute) *node.Node {
	return node.New("bm", attr.AssetImage("assets/oci/compute/bm.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) Container(opts ...attr.Attribute) *node.Node {
	return node.New("container", attr.AssetImage("assets/oci/compute/container.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) ContainerWhite(opts ...attr.Attribute) *node.Node {
	return node.New("container-white", attr.AssetImage("assets/oci/compute/container-white.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) Functions(opts ...attr.Attribute) *node.Node {
	return node.New("functions", attr.AssetImage("assets/oci/compute/functions.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) InstancePoolsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("instance-pools-white", attr.AssetImage("assets/oci/compute/instance-pools-white.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) OcirWhite(opts ...attr.Attribute) *node.Node {
	return node.New("ocir-white", attr.AssetImage("assets/oci/compute/ocir-white.png"), attr.Shape(attr.None))
}

func (c *ociComputeContainer) VmWhite(opts ...attr.Attribute) *node.Node {
	return node.New("vm-white", attr.AssetImage("assets/oci/compute/vm-white.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "oci.compute"
  Register(namespace, "AutoscaleWhite", OciCompute.AutoscaleWhite)
  Register(namespace, "FunctionsWhite", OciCompute.FunctionsWhite)
  Register(namespace, "Oke", OciCompute.Oke)
  Register(namespace, "Vm", OciCompute.Vm)
  Register(namespace, "Autoscale", OciCompute.Autoscale)
  Register(namespace, "InstancePools", OciCompute.InstancePools)
  Register(namespace, "Ocir", OciCompute.Ocir)
  Register(namespace, "OkeWhite", OciCompute.OkeWhite)
  Register(namespace, "BmWhite", OciCompute.BmWhite)
  Register(namespace, "Bm", OciCompute.Bm)
  Register(namespace, "Container", OciCompute.Container)
  Register(namespace, "ContainerWhite", OciCompute.ContainerWhite)
  Register(namespace, "Functions", OciCompute.Functions)
  Register(namespace, "InstancePoolsWhite", OciCompute.InstancePoolsWhite)
  Register(namespace, "OcirWhite", OciCompute.OcirWhite)
  Register(namespace, "VmWhite", OciCompute.VmWhite)
}
