package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type elasticSecurityContainer struct {
	path  string
	attrs []attr.Attribute
}

var ElasticSecurity = &elasticSecurityContainer{path: "assets/elastic/security"}

func (c *elasticSecurityContainer) Endpoint(opts ...attr.Attribute) *node.Node {
	return node.New("endpoint", attr.AssetImage("assets/elastic/security/endpoint.png"), attr.Shape(attr.None))
}

func (c *elasticSecurityContainer) Security(opts ...attr.Attribute) *node.Node {
	return node.New("security", attr.AssetImage("assets/elastic/security/security.png"), attr.Shape(attr.None))
}

func (c *elasticSecurityContainer) Siem(opts ...attr.Attribute) *node.Node {
	return node.New("siem", attr.AssetImage("assets/elastic/security/siem.png"), attr.Shape(attr.None))
}
