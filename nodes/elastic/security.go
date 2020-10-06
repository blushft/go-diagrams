package elastic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type securityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Security = &securityContainer{path: "assets/elastic/security"}

func (c *securityContainer) Endpoint(opts ...attr.Attribute) *node.Node {
	return node.New("endpoint", attr.AssetImage("assets/elastic/security/endpoint.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Security(opts ...attr.Attribute) *node.Node {
	return node.New("security", attr.AssetImage("assets/elastic/security/security.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Siem(opts ...attr.Attribute) *node.Node {
	return node.New("siem", attr.AssetImage("assets/elastic/security/siem.png"), attr.Shape(attr.None))
}
