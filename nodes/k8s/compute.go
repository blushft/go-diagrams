package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/k8s/compute"}

func (c *computeContainer) Pod(opts ...attr.Attribute) *node.Node {
	return node.New("pod", attr.AssetImage("assets/k8s/compute/pod.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Rs(opts ...attr.Attribute) *node.Node {
	return node.New("rs", attr.AssetImage("assets/k8s/compute/rs.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Sts(opts ...attr.Attribute) *node.Node {
	return node.New("sts", attr.AssetImage("assets/k8s/compute/sts.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Cronjob(opts ...attr.Attribute) *node.Node {
	return node.New("cronjob", attr.AssetImage("assets/k8s/compute/cronjob.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Deploy(opts ...attr.Attribute) *node.Node {
	return node.New("deploy", attr.AssetImage("assets/k8s/compute/deploy.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Ds(opts ...attr.Attribute) *node.Node {
	return node.New("ds", attr.AssetImage("assets/k8s/compute/ds.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Job(opts ...attr.Attribute) *node.Node {
	return node.New("job", attr.AssetImage("assets/k8s/compute/job.png"), attr.Shape(attr.None))
}
