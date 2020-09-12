package alibabacloud

import "github.com/blushft/go-diagrams/node"

type databaseContainer struct {
	path string
	opts []node.Option
}

var Database = &databaseContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/database",
}

func (c *databaseContainer) ApsaradbMemcache(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-memcache.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ApsaradbPostgresql(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-postgresql.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) RelationalDatabaseService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/relational-database-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ApsaradbCassandra(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-cassandra.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DataTransmissionService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/data-transmission-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DatabaseBackupService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/database-backup-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) HybriddbForMysql(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/hybriddb-for-mysql.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ApsaradbHbase(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-hbase.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ApsaradbRedis(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-redis.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ApsaradbSqlserver(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-sqlserver.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) GraphDatabaseService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/graph-database-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ApsaradbOceanbase(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-oceanbase.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ApsaradbPolardb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-polardb.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ApsaradbPpas(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-ppas.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DataManagementService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/data-management-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DisributeRelationalDatabaseService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/disribute-relational-database-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ApsaradbMongodb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/database/apsaradb-mongodb.png")}, c.opts, opts)
	return node.New(nopts...)
}
