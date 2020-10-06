package gcp

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type apiContainer struct {
	path  string
	attrs []attr.Attribute
}

var Api = &apiContainer{path: "assets/gcp/api"}

func (c *apiContainer) Endpoints(opts ...attr.Attribute) *node.Node {
	return node.New("endpoints", attr.AssetImage("assets/gcp/api/endpoints.png"), attr.Shape(attr.None))
}
