package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sPodConfigContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sPodConfig = &k8sPodConfigContainer{path: "assets/k8s/podconfig"}

func (c *k8sPodConfigContainer) Secret(opts ...attr.Attribute) *node.Node {
	return node.New("secret", attr.AssetImage("assets/k8s/podconfig/secret.png"), attr.Shape(attr.None))
}

func (c *k8sPodConfigContainer) Cm(opts ...attr.Attribute) *node.Node {
	return node.New("cm", attr.AssetImage("assets/k8s/podconfig/cm.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.podconfig"
  Register(namespace, "Secret", K8sPodConfig.Secret)
  Register(namespace, "Cm", K8sPodConfig.Cm)
}
