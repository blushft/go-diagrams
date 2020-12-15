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
	opts = append(opts, attr.AssetImage("assets/generic/place/datacenter.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("datacenter", opts...)
}

func init() {
	const namespace = "generic.place"
	Register(namespace, "Datacenter", Place.Datacenter)
}