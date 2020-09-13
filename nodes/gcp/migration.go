package gcp

import "github.com/blushft/go-diagrams/diagram"

type migrationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Migration = &migrationContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/migration",
}

func (c *migrationContainer) TransferAppliance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/migration/transfer-appliance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
