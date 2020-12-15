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
	opts = append(opts, attr.AssetImage("assets/elastic/enterprisesearch/app-search.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-search", opts...)
}

func (c *enterprisesearchContainer) EnterpriseSearch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/enterprisesearch/enterprise-search.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("enterprise-search", opts...)
}

func (c *enterprisesearchContainer) SiteSearch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/enterprisesearch/site-search.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("site-search", opts...)
}

func (c *enterprisesearchContainer) WorkplaceSearch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/enterprisesearch/workplace-search.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("workplace-search", opts...)
}
