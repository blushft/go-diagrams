package azure

import "github.com/blushft/go-diagrams/node"

type migrationContainer struct {
	path string
	opts []node.Option
}

var Migration = &migrationContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/migration",
}

func (c *migrationContainer) DatabaseMigrationServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/migration/database-migration-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) MigrationProjects(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/migration/migration-projects.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) RecoveryServicesVaults(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/migration/recovery-services-vaults.png")}, c.opts, opts)
	return node.New(nopts...)
}
