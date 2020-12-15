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
	opts = append(opts, attr.AssetImage("assets/gcp/compute/compute-engine.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("compute-engine", opts...)
}

func (c *gcpComputeContainer) ContainerOptimizedOs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/compute/container-optimized-os.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("container-optimized-os", opts...)
}

func (c *gcpComputeContainer) Functions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/compute/functions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("functions", opts...)
}

func (c *gcpComputeContainer) GkeOnPrem(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/compute/gke-on-prem.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("gke-on-prem", opts...)
}

func (c *gcpComputeContainer) Gpu(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/compute/gpu.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("gpu", opts...)
}

func (c *gcpComputeContainer) KubernetesEngine(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/compute/kubernetes-engine.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kubernetes-engine", opts...)
}

func (c *gcpComputeContainer) Run(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/compute/run.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("run", opts...)
}

func (c *gcpComputeContainer) AppEngine(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/compute/app-engine.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-engine", opts...)
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