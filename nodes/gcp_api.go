package nodes

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
	opts = append(opts, attr.AssetImage("assets/gcp/api/endpoints.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("endpoints", opts...)
}

func init(){
	const namespace = "gcp.api"
	Register(namespace, "Endpoints", Api.Endpoints)
}