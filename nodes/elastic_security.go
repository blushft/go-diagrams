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
	opts = append(opts, attr.AssetImage("assets/elastic/security/endpoint.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("endpoint", opts...)
}

func (c *elasticSecurityContainer) Security(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/security/security.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security", opts...)
}

func (c *elasticSecurityContainer) Siem(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/security/siem.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("siem", opts...)
}
