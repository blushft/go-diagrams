package alibabacloud

import "github.com/blushft/go-diagrams/node"

type webContainer struct {
	path string
	opts []node.Option
}

var Web = &webContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/web",
}

func (c *webContainer) Dns(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/web/dns.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *webContainer) Domain(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/web/domain.png")}, c.opts, opts)
	return node.New(nopts...)
}
