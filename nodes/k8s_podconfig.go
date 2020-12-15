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
	opts = append(opts, attr.AssetImage("assets/k8s/podconfig/secret.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("secret", opts...)
}

func (c *k8sPodConfigContainer) Cm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/k8s/podconfig/cm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cm", opts...)
}

func init() {
  const namespace = "k8s.podconfig"
  Register(namespace, "Secret", K8sPodConfig.Secret)
  Register(namespace, "Cm", K8sPodConfig.Cm)
}
