package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type AppsContainerContainer struct {
	path  string
	attrs []attr.Attribute
}

var AppsContainer = &AppsContainerContainer{path: "assets/apps/container"}

func (c *AppsContainerContainer) Docker(opts ...attr.Attribute) *node.Node {
	return node.New("docker", attr.AssetImage("assets/apps/container/docker.png"), attr.Shape(attr.None))
}

func (c *AppsContainerContainer) Rkt(opts ...attr.Attribute) *node.Node {
	return node.New("rkt", attr.AssetImage("assets/apps/container/rkt.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.container"
	Register(namespace, "Docker", AppsContainer.Docker)
	Register(namespace, "Rkt", AppsContainer.Rkt)
}

