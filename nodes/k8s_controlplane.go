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
	opts = append(opts, attr.AssetImage("assets/k8s/controlplane/k-proxy.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("k-proxy", opts...)
}

func (c *k8sContorlPlaneContainer) Kubelet(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/controlplane/kubelet.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kubelet", opts...)
}

func (c *k8sContorlPlaneContainer) Sched(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/controlplane/sched.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sched", opts...)
}

func (c *k8sContorlPlaneContainer) Api(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/controlplane/api.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api", opts...)
}

func (c *k8sContorlPlaneContainer) CCM(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/controlplane/c-c-m.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("c-c-m", opts...)
}

func (c *k8sContorlPlaneContainer) CM(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/controlplane/c-m.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("c-m", opts...)
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
