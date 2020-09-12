package gcp

import "github.com/blushft/go-diagrams/node"

type apiContainer struct {
	path string
	opts []node.Option
}

var Api = &apiContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/api",
}

func (c *apiContainer) Endpoints(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/api/endpoints.png")}, c.opts, opts)
	return node.New(nopts...)
}
