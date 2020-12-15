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
	opts = append(opts, attr.AssetImage("assets/azure/database/database-for-mysql-servers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database-for-mysql-servers", opts...)
}

func (c *azuredatabaseContainer) DatabaseForPostgresqlServers(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/database-for-postgresql-servers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database-for-postgresql-servers", opts...)
}

func (c *azuredatabaseContainer) ElasticDatabasePools(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/elastic-database-pools.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-database-pools", opts...)
}

func (c *azuredatabaseContainer) SqlDatawarehouse(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/sql-datawarehouse.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sql-datawarehouse", opts...)
}

func (c *azuredatabaseContainer) VirtualClusters(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/virtual-clusters.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("virtual-clusters", opts...)
}

func (c *azuredatabaseContainer) CacheForRedis(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/cache-for-redis.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cache-for-redis", opts...)
}

func (c *azuredatabaseContainer) DataLake(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/data-lake.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-lake", opts...)
}

func (c *azuredatabaseContainer) DatabaseForMariadbServers(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/database-for-mariadb-servers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database-for-mariadb-servers", opts...)
}

func (c *azuredatabaseContainer) ManagedDatabases(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/managed-databases.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("managed-databases", opts...)
}

func (c *azuredatabaseContainer) SqlDatabases(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/sql-databases.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sql-databases", opts...)
}

func (c *azuredatabaseContainer) SqlManagedInstances(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/sql-managed-instances.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sql-managed-instances", opts...)
}

func (c *azuredatabaseContainer) SqlServerStretchDatabases(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/sql-server-stretch-databases.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sql-server-stretch-databases", opts...)
}

func (c *azuredatabaseContainer) SqlServers(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/sql-servers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sql-servers", opts...)
}

func (c *azuredatabaseContainer) BlobStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/blob-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("blob-storage", opts...)
}

func (c *azuredatabaseContainer) ElasticJobAgents(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/elastic-job-agents.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-job-agents", opts...)
}

func (c *azuredatabaseContainer) VirtualDatacenter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/virtual-datacenter.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("virtual-datacenter", opts...)
}

func (c *azuredatabaseContainer) CosmosDb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/database/cosmos-db.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cosmos-db", opts...)
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
