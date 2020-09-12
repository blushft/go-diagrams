package aws

import "github.com/blushft/go-diagrams/node"

type databaseContainer struct {
	path string
	opts []node.Option
}

var Database = &databaseContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/database",
}

func (c *databaseContainer) Aurora(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/aurora.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DynamodbGlobalSecondaryIndex(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/dynamodb-global-secondary-index.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Dynamodb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/dynamodb.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Neptune(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/neptune.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Rds(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/rds.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DocumentdbMongodbCompatibility(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/documentdb-mongodb-compatibility.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DynamodbDax(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/dynamodb-dax.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DynamodbTable(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/dynamodb-table.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DatabaseMigrationService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/database-migration-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Database(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/database.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Elasticache(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/elasticache.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) QuantumLedgerDatabaseQldb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/quantum-ledger-database-qldb.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) RdsOnVmware(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/rds-on-vmware.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Redshift(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/redshift.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Timestream(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/database/timestream.png")}, c.opts, opts)
	return node.New(nopts...)
}
