package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type databaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var Database = &databaseContainer{path: "assets/azure/database"}

func (c *databaseContainer) SqlServers(opts ...attr.Attribute) *node.Node {
	return node.New("sql-servers", attr.AssetImage("assets/azure/database/sql-servers.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) CacheForRedis(opts ...attr.Attribute) *node.Node {
	return node.New("cache-for-redis", attr.AssetImage("assets/azure/database/cache-for-redis.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DatabaseForPostgresqlServers(opts ...attr.Attribute) *node.Node {
	return node.New("database-for-postgresql-servers", attr.AssetImage("assets/azure/database/database-for-postgresql-servers.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ElasticDatabasePools(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-database-pools", attr.AssetImage("assets/azure/database/elastic-database-pools.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) VirtualDatacenter(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-datacenter", attr.AssetImage("assets/azure/database/virtual-datacenter.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) BlobStorage(opts ...attr.Attribute) *node.Node {
	return node.New("blob-storage", attr.AssetImage("assets/azure/database/blob-storage.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) SqlDatabases(opts ...attr.Attribute) *node.Node {
	return node.New("sql-databases", attr.AssetImage("assets/azure/database/sql-databases.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) VirtualClusters(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-clusters", attr.AssetImage("assets/azure/database/virtual-clusters.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ManagedDatabases(opts ...attr.Attribute) *node.Node {
	return node.New("managed-databases", attr.AssetImage("assets/azure/database/managed-databases.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) SqlDatawarehouse(opts ...attr.Attribute) *node.Node {
	return node.New("sql-datawarehouse", attr.AssetImage("assets/azure/database/sql-datawarehouse.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) SqlManagedInstances(opts ...attr.Attribute) *node.Node {
	return node.New("sql-managed-instances", attr.AssetImage("assets/azure/database/sql-managed-instances.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) CosmosDb(opts ...attr.Attribute) *node.Node {
	return node.New("cosmos-db", attr.AssetImage("assets/azure/database/cosmos-db.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DataLake(opts ...attr.Attribute) *node.Node {
	return node.New("data-lake", attr.AssetImage("assets/azure/database/data-lake.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DatabaseForMysqlServers(opts ...attr.Attribute) *node.Node {
	return node.New("database-for-mysql-servers", attr.AssetImage("assets/azure/database/database-for-mysql-servers.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DatabaseForMariadbServers(opts ...attr.Attribute) *node.Node {
	return node.New("database-for-mariadb-servers", attr.AssetImage("assets/azure/database/database-for-mariadb-servers.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ElasticJobAgents(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-job-agents", attr.AssetImage("assets/azure/database/elastic-job-agents.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) SqlServerStretchDatabases(opts ...attr.Attribute) *node.Node {
	return node.New("sql-server-stretch-databases", attr.AssetImage("assets/azure/database/sql-server-stretch-databases.png"), attr.Shape(attr.None))
}
