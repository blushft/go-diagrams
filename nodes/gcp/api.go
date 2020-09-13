package gcp

import "github.com/blushft/go-diagrams/diagram"

type apiContainer struct {
	path string
	opts []diagram.NodeOption
}

var Api = &apiContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/api",
}

func (c *apiContainer) Endpoints(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/api/endpoints.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
