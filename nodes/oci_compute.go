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
	opts = append(opts, attr.AssetImage("assets/oci/compute/autoscale-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("autoscale-white", opts...)
}

func (c *ociComputeContainer) FunctionsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/functions-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("functions-white", opts...)
}

func (c *ociComputeContainer) Oke(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/oke.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("oke", opts...)
}

func (c *ociComputeContainer) Vm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/vm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vm", opts...)
}

func (c *ociComputeContainer) Autoscale(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/autoscale.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("autoscale", opts...)
}

func (c *ociComputeContainer) InstancePools(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/instance-pools.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("instance-pools", opts...)
}

func (c *ociComputeContainer) Ocir(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/ocir.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ocir", opts...)
}

func (c *ociComputeContainer) OkeWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/oke-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("oke-white", opts...)
}

func (c *ociComputeContainer) BmWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/bm-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bm-white", opts...)
}

func (c *ociComputeContainer) Bm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/bm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bm", opts...)
}

func (c *ociComputeContainer) Container(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/container.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("container", opts...)
}

func (c *ociComputeContainer) ContainerWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/container-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("container-white", opts...)
}

func (c *ociComputeContainer) Functions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/functions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("functions", opts...)
}

func (c *ociComputeContainer) InstancePoolsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/instance-pools-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("instance-pools-white", opts...)
}

func (c *ociComputeContainer) OcirWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/ocir-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ocir-white", opts...)
}

func (c *ociComputeContainer) VmWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/compute/vm-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vm-white", opts...)
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
