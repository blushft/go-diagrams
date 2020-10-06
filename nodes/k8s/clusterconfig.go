package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type clusterconfigContainer struct {
	path  string
	attrs []attr.Attribute
}

var Clusterconfig = &clusterconfigContainer{path: "assets/k8s/clusterconfig"}

func (c *clusterconfigContainer) Hpa(opts ...attr.Attribute) *node.Node {
	return node.New("hpa", attr.AssetImage("assets/k8s/clusterconfig/hpa.png"), attr.Shape(attr.None))
}

func (c *clusterconfigContainer) Limits(opts ...attr.Attribute) *node.Node {
	return node.New("limits", attr.AssetImage("assets/k8s/clusterconfig/limits.png"), attr.Shape(attr.None))
}

func (c *clusterconfigContainer) Quota(opts ...attr.Attribute) *node.Node {
	return node.New("quota", attr.AssetImage("assets/k8s/clusterconfig/quota.png"), attr.Shape(attr.None))
}
