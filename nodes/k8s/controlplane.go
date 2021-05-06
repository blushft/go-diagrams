package k8s

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type controlplaneContainer struct {
	path  string
	attrs []attr.Attribute
}

var Controlplane = &controlplaneContainer{path: "assets/k8s/controlplane"}

func (c *controlplaneContainer) Sched(opts ...attr.Attribute) *node.Node {
	return node.New("sched", attr.AssetImage("assets/k8s/controlplane/sched.png"), attr.Shape(attr.None))
}

func (c *controlplaneContainer) Api(opts ...attr.Attribute) *node.Node {
	return node.New("api", attr.AssetImage("assets/k8s/controlplane/api.png"), attr.Shape(attr.None))
}

func (c *controlplaneContainer) CCM(opts ...attr.Attribute) *node.Node {
	return node.New("c-c-m", attr.AssetImage("assets/k8s/controlplane/c-c-m.png"), attr.Shape(attr.None))
}

func (c *controlplaneContainer) CM(opts ...attr.Attribute) *node.Node {
	return node.New("c-m", attr.AssetImage("assets/k8s/controlplane/c-m.png"), attr.Shape(attr.None))
}

func (c *controlplaneContainer) KProxy(opts ...attr.Attribute) *node.Node {
	return node.New("k-proxy", attr.AssetImage("assets/k8s/controlplane/k-proxy.png"), attr.Shape(attr.None))
}

func (c *controlplaneContainer) Kubelet(opts ...attr.Attribute) *node.Node {
	return node.New("kubelet", attr.AssetImage("assets/k8s/controlplane/kubelet.png"), attr.Shape(attr.None))
}
