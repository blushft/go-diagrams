package aws

import "github.com/blushft/go-diagrams/diagram"

type migrationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Migration = &migrationContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/migration",
}

func (c *migrationContainer) TransferForSftp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/transfer-for-sftp.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) Datasync(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/datasync.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) MigrationAndTransfer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/migration-and-transfer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) Snowball(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/snowball.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) MigrationHub(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/migration-hub.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) ServerMigrationService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/server-migration-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) SnowballEdge(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/snowball-edge.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) Snowmobile(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/snowmobile.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) ApplicationDiscoveryService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/application-discovery-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) CloudendureMigration(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/cloudendure-migration.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *migrationContainer) DatabaseMigrationService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/migration/database-migration-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
