package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sNetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sNetwork = &k8sNetworkContainer{path: "assets/k8s/network"}

func (c *k8sNetworkContainer) Ep(opts ...attr.Attribute) *node.Node {
	return node.New("ep", attr.AssetImage("assets/k8s/network/ep.png"), attr.Shape(attr.None))
}

func (c *k8sNetworkContainer) Ing(opts ...attr.Attribute) *node.Node {
	return node.New("ing", attr.AssetImage("assets/k8s/network/ing.png"), attr.Shape(attr.None))
}

func (c *k8sNetworkContainer) Netpol(opts ...attr.Attribute) *node.Node {
	return node.New("netpol", attr.AssetImage("assets/k8s/network/netpol.png"), attr.Shape(attr.None))
}

func (c *k8sNetworkContainer) Svc(opts ...attr.Attribute) *node.Node {
	return node.New("svc", attr.AssetImage("assets/k8s/network/svc.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.network"
  Register(namespace, "Ep", K8sNetwork.Ep)
  Register(namespace, "Ing", K8sNetwork.Ing)
  Register(namespace, "Netpol", K8sNetwork.Netpol)
  Register(namespace, "Svc", K8sNetwork.Svc)
}
