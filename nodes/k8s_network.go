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
	opts = append(opts, attr.AssetImage("assets/k8s/network/ep.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ep", opts...)
}

func (c *k8sNetworkContainer) Ing(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/network/ing.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ing", opts...)
}

func (c *k8sNetworkContainer) Netpol(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/network/netpol.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("netpol", opts...)
}

func (c *k8sNetworkContainer) Svc(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/network/svc.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("svc", opts...)
}

func init() {
  const namespace = "k8s.network"
  Register(namespace, "Ep", K8sNetwork.Ep)
  Register(namespace, "Ing", K8sNetwork.Ing)
  Register(namespace, "Netpol", K8sNetwork.Netpol)
  Register(namespace, "Svc", K8sNetwork.Svc)
}
