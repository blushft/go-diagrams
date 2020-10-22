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
	return node.New("solr", attr.AssetImage("assets/apps/search/solr.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.search"
	Register(namespace, "Solr", Search.Solr)
}