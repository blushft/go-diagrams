package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type gcpComputeContainer struct {
	path  string
	attrs []attr.Attribute
}

var GCPCompute = &gcpComputeContainer{path: "assets/gcp/compute"}

func (c *gcpComputeContainer) ComputeEngine(opts ...attr.Attribute) *node.Node {
	return node.New("compute-engine", attr.AssetImage("assets/gcp/compute/compute-engine.png"), attr.Shape(attr.None))
}

func (c *gcpComputeContainer) ContainerOptimizedOs(opts ...attr.Attribute) *node.Node {
	return node.New("container-optimized-os", attr.AssetImage("assets/gcp/compute/container-optimized-os.png"), attr.Shape(attr.None))
}

func (c *gcpComputeContainer) Functions(opts ...attr.Attribute) *node.Node {
	return node.New("functions", attr.AssetImage("assets/gcp/compute/functions.png"), attr.Shape(attr.None))
}

func (c *gcpComputeContainer) GkeOnPrem(opts ...attr.Attribute) *node.Node {
	return node.New("gke-on-prem", attr.AssetImage("assets/gcp/compute/gke-on-prem.png"), attr.Shape(attr.None))
}

func (c *gcpComputeContainer) Gpu(opts ...attr.Attribute) *node.Node {
	return node.New("gpu", attr.AssetImage("assets/gcp/compute/gpu.png"), attr.Shape(attr.None))
}

func (c *gcpComputeContainer) KubernetesEngine(opts ...attr.Attribute) *node.Node {
	return node.New("kubernetes-engine", attr.AssetImage("assets/gcp/compute/kubernetes-engine.png"), attr.Shape(attr.None))
}

func (c *gcpComputeContainer) Run(opts ...attr.Attribute) *node.Node {
	return node.New("run", attr.AssetImage("assets/gcp/compute/run.png"), attr.Shape(attr.None))
}

func (c *gcpComputeContainer) AppEngine(opts ...attr.Attribute) *node.Node {
	return node.New("app-engine", attr.AssetImage("assets/gcp/compute/app-engine.png"), attr.Shape(attr.None))
}


func init(){
	const namespace = "gcp.compute"
	Register(namespace, "ComputeEngine", GCPCompute.ComputeEngine)
	Register(namespace, "ContainerOptimizedOs", GCPCompute.ContainerOptimizedOs)
	Register(namespace, "Functions", GCPCompute.Functions)
	Register(namespace, "GkeOnPrem", GCPCompute.GkeOnPrem)
	Register(namespace, "Gpu", GCPCompute.Gpu)
	Register(namespace, "KubernetesEngine", GCPCompute.KubernetesEngine)
	Register(namespace, "Run", GCPCompute.Run)
	Register(namespace, "AppEngine", GCPCompute.AppEngine)
}