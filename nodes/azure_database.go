package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuredatabaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureDatabase = &azuredatabaseContainer{path: "assets/azure/database"}

func (c *azuredatabaseContainer) DatabaseForMysqlServers(opts ...attr.Attribute) *node.Node {
	return node.New("database-for-mysql-servers", attr.AssetImage("assets/azure/database/database-for-mysql-servers.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) DatabaseForPostgresqlServers(opts ...attr.Attribute) *node.Node {
	return node.New("database-for-postgresql-servers", attr.AssetImage("assets/azure/database/database-for-postgresql-servers.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) ElasticDatabasePools(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-database-pools", attr.AssetImage("assets/azure/database/elastic-database-pools.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) SqlDatawarehouse(opts ...attr.Attribute) *node.Node {
	return node.New("sql-datawarehouse", attr.AssetImage("assets/azure/database/sql-datawarehouse.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) VirtualClusters(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-clusters", attr.AssetImage("assets/azure/database/virtual-clusters.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) CacheForRedis(opts ...attr.Attribute) *node.Node {
	return node.New("cache-for-redis", attr.AssetImage("assets/azure/database/cache-for-redis.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) DataLake(opts ...attr.Attribute) *node.Node {
	return node.New("data-lake", attr.AssetImage("assets/azure/database/data-lake.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) DatabaseForMariadbServers(opts ...attr.Attribute) *node.Node {
	return node.New("database-for-mariadb-servers", attr.AssetImage("assets/azure/database/database-for-mariadb-servers.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) ManagedDatabases(opts ...attr.Attribute) *node.Node {
	return node.New("managed-databases", attr.AssetImage("assets/azure/database/managed-databases.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) SqlDatabases(opts ...attr.Attribute) *node.Node {
	return node.New("sql-databases", attr.AssetImage("assets/azure/database/sql-databases.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) SqlManagedInstances(opts ...attr.Attribute) *node.Node {
	return node.New("sql-managed-instances", attr.AssetImage("assets/azure/database/sql-managed-instances.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) SqlServerStretchDatabases(opts ...attr.Attribute) *node.Node {
	return node.New("sql-server-stretch-databases", attr.AssetImage("assets/azure/database/sql-server-stretch-databases.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) SqlServers(opts ...attr.Attribute) *node.Node {
	return node.New("sql-servers", attr.AssetImage("assets/azure/database/sql-servers.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) BlobStorage(opts ...attr.Attribute) *node.Node {
	return node.New("blob-storage", attr.AssetImage("assets/azure/database/blob-storage.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) ElasticJobAgents(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-job-agents", attr.AssetImage("assets/azure/database/elastic-job-agents.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) VirtualDatacenter(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-datacenter", attr.AssetImage("assets/azure/database/virtual-datacenter.png"), attr.Shape(attr.None))
}

func (c *azuredatabaseContainer) CosmosDb(opts ...attr.Attribute) *node.Node {
	return node.New("cosmos-db", attr.AssetImage("assets/azure/database/cosmos-db.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.database"
  Register(namespace, "DatabaseForMysqlServers", azureDatabase.DatabaseForMysqlServers)
  Register(namespace, "DatabaseForPostgresqlServers", azureDatabase.DatabaseForPostgresqlServers)
  Register(namespace, "ElasticDatabasePools", azureDatabase.ElasticDatabasePools)
  Register(namespace, "SqlDatawarehouse", azureDatabase.SqlDatawarehouse)
  Register(namespace, "VirtualClusters", azureDatabase.VirtualClusters)
  Register(namespace, "CacheForRedis", azureDatabase.CacheForRedis)
  Register(namespace, "DataLake", azureDatabase.DataLake)
  Register(namespace, "DatabaseForMariadbServers", azureDatabase.DatabaseForMariadbServers)
  Register(namespace, "ManagedDatabases", azureDatabase.ManagedDatabases)
  Register(namespace, "SqlDatabases", azureDatabase.SqlDatabases)
  Register(namespace, "SqlManagedInstances", azureDatabase.SqlManagedInstances)
  Register(namespace, "SqlServerStretchDatabases", azureDatabase.SqlServerStretchDatabases)
  Register(namespace, "SqlServers", azureDatabase.SqlServers)
  Register(namespace, "BlobStorage", azureDatabase.BlobStorage)
  Register(namespace, "ElasticJobAgents", azureDatabase.ElasticJobAgents)
  Register(namespace, "VirtualDatacenter", azureDatabase.VirtualDatacenter)
  Register(namespace, "CosmosDb", azureDatabase.CosmosDb)
}
