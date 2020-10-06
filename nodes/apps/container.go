package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type containerContainer struct {
	path  string
	attrs []attr.Attribute
}

var Container = &containerContainer{path: "assets/apps/container"}

func (c *containerContainer) Docker(opts ...attr.Attribute) *node.Node {
	return node.New("docker", attr.AssetImage("assets/apps/container/docker.png"), attr.Shape(attr.None))
}

func (c *containerContainer) Rkt(opts ...attr.Attribute) *node.Node {
	return node.New("rkt", attr.AssetImage("assets/apps/container/rkt.png"), attr.Shape(attr.None))
}
