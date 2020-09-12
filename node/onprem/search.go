package onprem

import "github.com/blushft/go-diagrams/node"

type searchContainer struct {
	path string
	opts []node.Option
}

var Search = &searchContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/search",
}

func (c *searchContainer) Solr(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/search/solr.png")}, c.opts, opts)
	return node.New(nopts...)
}
