package oci

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/oci/compute"}

func (c *computeContainer) AutoscaleWhite(opts ...attr.Attribute) *node.Node {
	return node.New("autoscale-white", attr.AssetImage("assets/oci/compute/autoscale-white.png"), attr.Shape(attr.None))
}

func (c *computeContainer) FunctionsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("functions-white", attr.AssetImage("assets/oci/compute/functions-white.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Oke(opts ...attr.Attribute) *node.Node {
	return node.New("oke", attr.AssetImage("assets/oci/compute/oke.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Vm(opts ...attr.Attribute) *node.Node {
	return node.New("vm", attr.AssetImage("assets/oci/compute/vm.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Autoscale(opts ...attr.Attribute) *node.Node {
	return node.New("autoscale", attr.AssetImage("assets/oci/compute/autoscale.png"), attr.Shape(attr.None))
}

func (c *computeContainer) InstancePools(opts ...attr.Attribute) *node.Node {
	return node.New("instance-pools", attr.AssetImage("assets/oci/compute/instance-pools.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Ocir(opts ...attr.Attribute) *node.Node {
	return node.New("ocir", attr.AssetImage("assets/oci/compute/ocir.png"), attr.Shape(attr.None))
}

func (c *computeContainer) OkeWhite(opts ...attr.Attribute) *node.Node {
	return node.New("oke-white", attr.AssetImage("assets/oci/compute/oke-white.png"), attr.Shape(attr.None))
}

func (c *computeContainer) BmWhite(opts ...attr.Attribute) *node.Node {
	return node.New("bm-white", attr.AssetImage("assets/oci/compute/bm-white.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Bm(opts ...attr.Attribute) *node.Node {
	return node.New("bm", attr.AssetImage("assets/oci/compute/bm.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Container(opts ...attr.Attribute) *node.Node {
	return node.New("container", attr.AssetImage("assets/oci/compute/container.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ContainerWhite(opts ...attr.Attribute) *node.Node {
	return node.New("container-white", attr.AssetImage("assets/oci/compute/container-white.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Functions(opts ...attr.Attribute) *node.Node {
	return node.New("functions", attr.AssetImage("assets/oci/compute/functions.png"), attr.Shape(attr.None))
}

func (c *computeContainer) InstancePoolsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("instance-pools-white", attr.AssetImage("assets/oci/compute/instance-pools-white.png"), attr.Shape(attr.None))
}

func (c *computeContainer) OcirWhite(opts ...attr.Attribute) *node.Node {
	return node.New("ocir-white", attr.AssetImage("assets/oci/compute/ocir-white.png"), attr.Shape(attr.None))
}

func (c *computeContainer) VmWhite(opts ...attr.Attribute) *node.Node {
	return node.New("vm-white", attr.AssetImage("assets/oci/compute/vm-white.png"), attr.Shape(attr.None))
}
