package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type webContainer struct {
	path string
	opts []diagram.NodeOption
}

var Web = &webContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/web",
}

func (c *webContainer) Dns(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/web/dns.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *webContainer) Domain(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/web/domain.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
