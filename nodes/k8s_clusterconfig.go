package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sClusterConfigContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sClusterConfig = &k8sClusterConfigContainer{path: "assets/k8s/clusterconfig"}

func (c *k8sClusterConfigContainer) Hpa(opts ...attr.Attribute) *node.Node {
	return node.New("hpa", attr.AssetImage("assets/k8s/clusterconfig/hpa.png"), attr.Shape(attr.None))
}

func (c *k8sClusterConfigContainer) Limits(opts ...attr.Attribute) *node.Node {
	return node.New("limits", attr.AssetImage("assets/k8s/clusterconfig/limits.png"), attr.Shape(attr.None))
}

func (c *k8sClusterConfigContainer) Quota(opts ...attr.Attribute) *node.Node {
	return node.New("quota", attr.AssetImage("assets/k8s/clusterconfig/quota.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.clusterconfig"
  Register(namespace, "Hpa", K8sClusterConfig.Hpa)
  Register(namespace, "Limits", K8sClusterConfig.Limits)
  Register(namespace, "Quota", K8sClusterConfig.Quota)
}
