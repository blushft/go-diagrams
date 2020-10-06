package gcp

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/gcp/compute"}

func (c *computeContainer) ComputeEngine(opts ...attr.Attribute) *node.Node {
	return node.New("compute-engine", attr.AssetImage("assets/gcp/compute/compute-engine.png"), attr.Shape(attr.None))
}

func (c *computeContainer) ContainerOptimizedOs(opts ...attr.Attribute) *node.Node {
	return node.New("container-optimized-os", attr.AssetImage("assets/gcp/compute/container-optimized-os.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Functions(opts ...attr.Attribute) *node.Node {
	return node.New("functions", attr.AssetImage("assets/gcp/compute/functions.png"), attr.Shape(attr.None))
}

func (c *computeContainer) GkeOnPrem(opts ...attr.Attribute) *node.Node {
	return node.New("gke-on-prem", attr.AssetImage("assets/gcp/compute/gke-on-prem.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Gpu(opts ...attr.Attribute) *node.Node {
	return node.New("gpu", attr.AssetImage("assets/gcp/compute/gpu.png"), attr.Shape(attr.None))
}

func (c *computeContainer) KubernetesEngine(opts ...attr.Attribute) *node.Node {
	return node.New("kubernetes-engine", attr.AssetImage("assets/gcp/compute/kubernetes-engine.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Run(opts ...attr.Attribute) *node.Node {
	return node.New("run", attr.AssetImage("assets/gcp/compute/run.png"), attr.Shape(attr.None))
}

func (c *computeContainer) AppEngine(opts ...attr.Attribute) *node.Node {
	return node.New("app-engine", attr.AssetImage("assets/gcp/compute/app-engine.png"), attr.Shape(attr.None))
}
