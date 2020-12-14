package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sInfraContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sInfra = &k8sInfraContainer{path: "assets/k8s/infra"}

func (c *k8sInfraContainer) Etcd(opts ...attr.Attribute) *node.Node {
	return node.New("etcd", attr.AssetImage("assets/k8s/infra/etcd.png"), attr.Shape(attr.None))
}

func (c *k8sInfraContainer) Master(opts ...attr.Attribute) *node.Node {
	return node.New("master", attr.AssetImage("assets/k8s/infra/master.png"), attr.Shape(attr.None))
}

func (c *k8sInfraContainer) Node(opts ...attr.Attribute) *node.Node {
	return node.New("node", attr.AssetImage("assets/k8s/infra/node.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.infra"
  Register(namespace, "Etcd", K8sInfra.Etcd)
  Register(namespace, "Master", K8sInfra.Master)
  Register(namespace, "Node", K8sInfra.Node)
}
