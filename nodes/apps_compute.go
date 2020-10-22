package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsComputeContainer struct {
	path  string
	attrs []attr.Attribute
}

var AppsCompute = &appsComputeContainer{path: "assets/apps/compute"}

func (c *appsComputeContainer) Nomad(opts ...attr.Attribute) *node.Node {
	return node.New("nomad", attr.AssetImage("assets/apps/compute/nomad.png"), attr.Shape(attr.None))
}

func (c *appsComputeContainer) Server(opts ...attr.Attribute) *node.Node {
	return node.New("server", attr.AssetImage("assets/apps/compute/server.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.compute"
	Register(namespace, "Nomad", AppsCompute.Nomad)
	Register(namespace, "Server", AppsCompute.Server)
}

