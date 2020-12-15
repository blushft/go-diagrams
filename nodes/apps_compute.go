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
	opts = append(opts, attr.AssetImage("assets/apps/compute/nomad.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nomad", opts...)
}

func (c *appsComputeContainer) Server(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/compute/server.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("server", opts...)
}

func init() {
	const namespace = "apps.compute"
	Register(namespace, "Nomad", AppsCompute.Nomad)
	Register(namespace, "Server", AppsCompute.Server)
}

