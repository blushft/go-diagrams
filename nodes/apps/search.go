package apps

import "github.com/blushft/go-diagrams/diagram"

type searchContainer struct {
	path string
	opts []diagram.NodeOption
}

var Search = &searchContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/search",
}

func (c *searchContainer) Solr(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/search/solr.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
