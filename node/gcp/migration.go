package gcp

import "github.com/blushft/go-diagrams/node"

type migrationContainer struct {
	path string
	opts []node.Option
}

var Migration = &migrationContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/migration",
}

func (c *migrationContainer) TransferAppliance(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/migration/transfer-appliance.png")}, c.opts, opts)
	return node.New(nopts...)
}
