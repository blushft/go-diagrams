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
	opts = append(opts, attr.AssetImage("assets/k8s/clusterconfig/hpa.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hpa", opts...)
}

func (c *k8sClusterConfigContainer) Limits(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/clusterconfig/limits.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("limits", opts...)
}

func (c *k8sClusterConfigContainer) Quota(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/clusterconfig/quota.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("quota", opts...)
}

func init() {
  const namespace = "k8s.clusterconfig"
  Register(namespace, "Hpa", K8sClusterConfig.Hpa)
  Register(namespace, "Limits", K8sClusterConfig.Limits)
  Register(namespace, "Quota", K8sClusterConfig.Quota)
}
