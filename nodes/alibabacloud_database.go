package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alibabaCloudDatabase struct {
	path  string
	attrs []attr.Attribute
}

var AlibabaCloudDatabase = &alibabaCloudDatabase{path: "assets/alibabacloud/database"}

func (c *alibabaCloudDatabase) ApsaradbMongodb(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-mongodb", attr.AssetImage("assets/alibabacloud/database/apsaradb-mongodb.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) ApsaradbOceanbase(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-oceanbase", attr.AssetImage("assets/alibabacloud/database/apsaradb-oceanbase.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) ApsaradbPostgresql(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-postgresql", attr.AssetImage("assets/alibabacloud/database/apsaradb-postgresql.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) ApsaradbPpas(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-ppas", attr.AssetImage("assets/alibabacloud/database/apsaradb-ppas.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) ApsaradbMemcache(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-memcache", attr.AssetImage("assets/alibabacloud/database/apsaradb-memcache.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) ApsaradbPolardb(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-polardb", attr.AssetImage("assets/alibabacloud/database/apsaradb-polardb.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) ApsaradbRedis(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-redis", attr.AssetImage("assets/alibabacloud/database/apsaradb-redis.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) DatabaseBackupService(opts ...attr.Attribute) *node.Node {
	return node.New("database-backup-service", attr.AssetImage("assets/alibabacloud/database/database-backup-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) RelationalDatabaseService(opts ...attr.Attribute) *node.Node {
	return node.New("relational-database-service", attr.AssetImage("assets/alibabacloud/database/relational-database-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) ApsaradbCassandra(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-cassandra", attr.AssetImage("assets/alibabacloud/database/apsaradb-cassandra.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) ApsaradbHbase(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-hbase", attr.AssetImage("assets/alibabacloud/database/apsaradb-hbase.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) ApsaradbSqlserver(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-sqlserver", attr.AssetImage("assets/alibabacloud/database/apsaradb-sqlserver.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) DataManagementService(opts ...attr.Attribute) *node.Node {
	return node.New("data-management-service", attr.AssetImage("assets/alibabacloud/database/data-management-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) GraphDatabaseService(opts ...attr.Attribute) *node.Node {
	return node.New("graph-database-service", attr.AssetImage("assets/alibabacloud/database/graph-database-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) DataTransmissionService(opts ...attr.Attribute) *node.Node {
	return node.New("data-transmission-service", attr.AssetImage("assets/alibabacloud/database/data-transmission-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) DisributeRelationalDatabaseService(opts ...attr.Attribute) *node.Node {
	return node.New("disribute-relational-database-service", attr.AssetImage("assets/alibabacloud/database/disribute-relational-database-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudDatabase) HybriddbForMysql(opts ...attr.Attribute) *node.Node {
	return node.New("hybriddb-for-mysql", attr.AssetImage("assets/alibabacloud/database/hybriddb-for-mysql.png"), attr.Shape(attr.None))
}

func init(){
	namespace := "alibabacloud.database"
	Register(namespace, "ApsaradbMongodb", AlibabaCloudDatabase.ApsaradbMongodb)
	Register(namespace, "ApsaradbOceanbase", AlibabaCloudDatabase.ApsaradbOceanbase)
	Register(namespace, "ApsaradbPostgresql", AlibabaCloudDatabase.ApsaradbPostgresql)
	Register(namespace, "ApsaradbPpas", AlibabaCloudDatabase.ApsaradbPpas)
	Register(namespace, "ApsaradbMemcache", AlibabaCloudDatabase.ApsaradbMemcache)
	Register(namespace, "ApsaradbPolardb", AlibabaCloudDatabase.ApsaradbPolardb)
	Register(namespace, "ApsaradbRedis", AlibabaCloudDatabase.ApsaradbRedis)
	Register(namespace, "DatabaseBackupService", AlibabaCloudDatabase.DatabaseBackupService)
	Register(namespace, "RelationalDatabaseService", AlibabaCloudDatabase.RelationalDatabaseService)
	Register(namespace, "ApsaradbCassandra", AlibabaCloudDatabase.ApsaradbCassandra)
	Register(namespace, "ApsaradbHbase", AlibabaCloudDatabase.ApsaradbHbase)
	Register(namespace, "ApsaradbSqlserver", AlibabaCloudDatabase.ApsaradbSqlserver)
	Register(namespace, "DataManagementService", AlibabaCloudDatabase.DataManagementService)
	Register(namespace, "GraphDatabaseService", AlibabaCloudDatabase.GraphDatabaseService)
	Register(namespace, "DataTransmissionService", AlibabaCloudDatabase.DataTransmissionService)
	Register(namespace, "DisributeRelationalDatabaseService", AlibabaCloudDatabase.DisributeRelationalDatabaseService)
	Register(namespace, "HybriddbForMysql", AlibabaCloudDatabase.HybriddbForMysql)
}