package aws

import "github.com/blushft/go-diagrams/node"

type migrationContainer struct {
	path string
	opts []node.Option
}

var Migration = &migrationContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/migration",
}

func (c *migrationContainer) TransferForSftp(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/transfer-for-sftp.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) CloudendureMigration(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/cloudendure-migration.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) DatabaseMigrationService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/database-migration-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) Datasync(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/datasync.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) MigrationAndTransfer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/migration-and-transfer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) ServerMigrationService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/server-migration-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) Snowball(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/snowball.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) Snowmobile(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/snowmobile.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) ApplicationDiscoveryService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/application-discovery-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) MigrationHub(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/migration-hub.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *migrationContainer) SnowballEdge(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/migration/snowball-edge.png")}, c.opts, opts)
	return node.New(nopts...)
}
