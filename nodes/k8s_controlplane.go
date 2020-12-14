package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sContorlPlaneContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sContorlPlane = &k8sContorlPlaneContainer{path: "assets/k8s/controlplane"}

func (c *k8sContorlPlaneContainer) KProxy(opts ...attr.Attribute) *node.Node {
	return node.New("k-proxy", attr.AssetImage("assets/k8s/controlplane/k-proxy.png"), attr.Shape(attr.None))
}

func (c *k8sContorlPlaneContainer) Kubelet(opts ...attr.Attribute) *node.Node {
	return node.New("kubelet", attr.AssetImage("assets/k8s/controlplane/kubelet.png"), attr.Shape(attr.None))
}

func (c *k8sContorlPlaneContainer) Sched(opts ...attr.Attribute) *node.Node {
	return node.New("sched", attr.AssetImage("assets/k8s/controlplane/sched.png"), attr.Shape(attr.None))
}

func (c *k8sContorlPlaneContainer) Api(opts ...attr.Attribute) *node.Node {
	return node.New("api", attr.AssetImage("assets/k8s/controlplane/api.png"), attr.Shape(attr.None))
}

func (c *k8sContorlPlaneContainer) CCM(opts ...attr.Attribute) *node.Node {
	return node.New("c-c-m", attr.AssetImage("assets/k8s/controlplane/c-c-m.png"), attr.Shape(attr.None))
}

func (c *k8sContorlPlaneContainer) CM(opts ...attr.Attribute) *node.Node {
	return node.New("c-m", attr.AssetImage("assets/k8s/controlplane/c-m.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.controlplane"
  Register(namespace, "KProxy", K8sContorlPlane.KProxy)
  Register(namespace, "Kubelet", K8sContorlPlane.Kubelet)
  Register(namespace, "Sched", K8sContorlPlane.Sched)
  Register(namespace, "Api", K8sContorlPlane.Api)
  Register(namespace, "CCM", K8sContorlPlane.CCM)
  Register(namespace, "CM", K8sContorlPlane.CM)
}
