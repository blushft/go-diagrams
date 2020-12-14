package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type enterprisesearchContainer struct {
	path  string
	attrs []attr.Attribute
}

var Enterprisesearch = &enterprisesearchContainer{path: "assets/elastic/enterprisesearch"}

func (c *enterprisesearchContainer) AppSearch(opts ...attr.Attribute) *node.Node {
	return node.New("app-search", attr.AssetImage("assets/elastic/enterprisesearch/app-search.png"), attr.Shape(attr.None))
}

func (c *enterprisesearchContainer) EnterpriseSearch(opts ...attr.Attribute) *node.Node {
	return node.New("enterprise-search", attr.AssetImage("assets/elastic/enterprisesearch/enterprise-search.png"), attr.Shape(attr.None))
}

func (c *enterprisesearchContainer) SiteSearch(opts ...attr.Attribute) *node.Node {
	return node.New("site-search", attr.AssetImage("assets/elastic/enterprisesearch/site-search.png"), attr.Shape(attr.None))
}

func (c *enterprisesearchContainer) WorkplaceSearch(opts ...attr.Attribute) *node.Node {
	return node.New("workplace-search", attr.AssetImage("assets/elastic/enterprisesearch/workplace-search.png"), attr.Shape(attr.None))
}
