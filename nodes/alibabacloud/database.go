package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type databaseContainer struct {
	path string
	opts []diagram.NodeOption
}

var Database = &databaseContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/database",
}

func (c *databaseContainer) DatabaseBackupService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/database-backup-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) RelationalDatabaseService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/relational-database-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbMongodb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-mongodb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbRedis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-redis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DataManagementService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/data-management-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DataTransmissionService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/data-transmission-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) GraphDatabaseService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/graph-database-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbCassandra(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-cassandra.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbHbase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-hbase.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbPpas(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-ppas.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbOceanbase(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-oceanbase.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbPolardb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-polardb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) HybriddbForMysql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/hybriddb-for-mysql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DisributeRelationalDatabaseService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/disribute-relational-database-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbMemcache(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-memcache.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbPostgresql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-postgresql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ApsaradbSqlserver(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/database/apsaradb-sqlserver.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
