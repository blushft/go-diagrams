package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type searchContainer struct {
	path  string
	attrs []attr.Attribute
}

var Search = &searchContainer{path: "assets/apps/search"}

func (c *searchContainer) Solr(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/search/solr.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("solr", opts...)
}

func init() {
	const namespace = "apps.search"
	Register(namespace, "Solr", Search.Solr)
}