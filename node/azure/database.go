package azure

import "github.com/blushft/go-diagrams/node"

type databaseContainer struct {
	path string
	opts []node.Option
}

var Database = &databaseContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/database",
}

func (c *databaseContainer) DatabaseForPostgresqlServers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/database-for-postgresql-servers.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ElasticJobAgents(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/elastic-job-agents.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) SqlDatawarehouse(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/sql-datawarehouse.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) BlobStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/blob-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) CosmosDb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/cosmos-db.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DatabaseForMysqlServers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/database-for-mysql-servers.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ManagedDatabases(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/managed-databases.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) SqlDatabases(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/sql-databases.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) SqlServerStretchDatabases(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/sql-server-stretch-databases.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) CacheForRedis(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/cache-for-redis.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DataLake(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/data-lake.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) VirtualClusters(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/virtual-clusters.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) VirtualDatacenter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/virtual-datacenter.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) SqlManagedInstances(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/sql-managed-instances.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) SqlServers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/sql-servers.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DatabaseForMariadbServers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/database-for-mariadb-servers.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ElasticDatabasePools(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/database/elastic-database-pools.png")}, c.opts, opts)
	return node.New(nopts...)
}
