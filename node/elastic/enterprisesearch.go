package elastic

import "github.com/blushft/go-diagrams/node"

type enterprisesearchContainer struct {
	path string
	opts []node.Option
}

var Enterprisesearch = &enterprisesearchContainer{
	opts: node.OptionSet{node.Provider("elastic"), node.Shape("none")},
	path: "assets/elastic/enterprisesearch",
}

func (c *enterprisesearchContainer) AppSearch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/enterprisesearch/app-search.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *enterprisesearchContainer) EnterpriseSearch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/enterprisesearch/enterprise-search.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *enterprisesearchContainer) SiteSearch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/enterprisesearch/site-search.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *enterprisesearchContainer) WorkplaceSearch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/enterprisesearch/workplace-search.png")}, c.opts, opts)
	return node.New(nopts...)
}
