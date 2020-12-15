package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsCdContainer struct {
	path  string
	attrs []attr.Attribute
}

var Cd = &appsCdContainer{path: "assets/apps/cd"}

func (c *appsCdContainer) Spinnaker(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/cd/spinnaker.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("spinnaker", opts...)
}

func (c *appsCdContainer) TektonCli(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/cd/tekton-cli.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tekton-cli", opts...)
}

func (c *appsCdContainer) Tekton(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/cd/tekton.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tekton", opts...)
}

func init() {
	const namespace = "apps.cd"
	Register(namespace, "Spinnaker", Cd.Spinnaker)
	Register(namespace, "TektonCli", Cd.TektonCli)
	Register(namespace, "Tekton", Cd.Tekton)
}