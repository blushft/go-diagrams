package azure

import "github.com/blushft/go-diagrams/diagram"

type migrationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Migration = &migrationContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/migration",
}

func (c *migrationContainer) DatabaseMigrationServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/migration/database-migration-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) MigrationProjects(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/migration/migration-projects.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) RecoveryServicesVaults(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/migration/recovery-services-vaults.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
