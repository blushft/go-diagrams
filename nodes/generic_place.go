package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type placeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Place = &placeContainer{path: "assets/generic/place"}

func (c *placeContainer) Datacenter(opts ...attr.Attribute) *node.Node {
	return node.New("datacenter", attr.AssetImage("assets/generic/place/datacenter.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "generic.place"
	Register(namespace, "Datacenter", Place.Datacenter)
}