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
	opts = append(opts, attr.AssetImage("assets/apps/container/docker.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("docker", opts...)
}

func (c *AppsContainerContainer) Rkt(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/container/rkt.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rkt", opts...)
}

func init() {
	const namespace = "apps.container"
	Register(namespace, "Docker", AppsContainer.Docker)
	Register(namespace, "Rkt", AppsContainer.Rkt)
}

