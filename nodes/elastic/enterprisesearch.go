package elastic

import "github.com/blushft/go-diagrams/diagram"

type enterprisesearchContainer struct {
	path string
	opts []diagram.NodeOption
}

var Enterprisesearch = &enterprisesearchContainer{
	opts: diagram.OptionSet{diagram.Provider("elastic"), diagram.NodeShape("none")},
	path: "assets/elastic/enterprisesearch",
}

func (c *enterprisesearchContainer) AppSearch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/enterprisesearch/app-search.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *enterprisesearchContainer) EnterpriseSearch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/enterprisesearch/enterprise-search.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *enterprisesearchContainer) SiteSearch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/enterprisesearch/site-search.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *enterprisesearchContainer) WorkplaceSearch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/enterprisesearch/workplace-search.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
