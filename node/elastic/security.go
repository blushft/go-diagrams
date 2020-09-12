package elastic

import "github.com/blushft/go-diagrams/node"

type securityContainer struct {
	path string
	opts []node.Option
}

var Security = &securityContainer{
	opts: node.OptionSet{node.Provider("elastic"), node.Shape("none")},
	path: "assets/elastic/security",
}

func (c *securityContainer) Endpoint(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/security/endpoint.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Security(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/security/security.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Siem(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/security/siem.png")}, c.opts, opts)
	return node.New(nopts...)
}
