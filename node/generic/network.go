package generic

import "github.com/blushft/go-diagrams/node"

type networkContainer struct {
	path string
	opts []node.Option
}

var Network = &networkContainer{
	opts: node.OptionSet{node.Provider("generic"), node.Shape("none")},
	path: "assets/generic/network",
}

func (c *networkContainer) Router(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/network/router.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Switch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/network/switch.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Vpn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/network/vpn.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Firewall(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/network/firewall.png")}, c.opts, opts)
	return node.New(nopts...)
}
