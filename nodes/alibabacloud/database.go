package alibabacloud

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type databaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var Database = &databaseContainer{path: "assets/alibabacloud/database"}

func (c *databaseContainer) ApsaradbHbase(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-hbase", attr.AssetImage("assets/alibabacloud/database/apsaradb-hbase.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ApsaradbPolardb(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-polardb", attr.AssetImage("assets/alibabacloud/database/apsaradb-polardb.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DataManagementService(opts ...attr.Attribute) *node.Node {
	return node.New("data-management-service", attr.AssetImage("assets/alibabacloud/database/data-management-service.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ApsaradbCassandra(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-cassandra", attr.AssetImage("assets/alibabacloud/database/apsaradb-cassandra.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ApsaradbMemcache(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-memcache", attr.AssetImage("assets/alibabacloud/database/apsaradb-memcache.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ApsaradbPpas(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-ppas", attr.AssetImage("assets/alibabacloud/database/apsaradb-ppas.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) GraphDatabaseService(opts ...attr.Attribute) *node.Node {
	return node.New("graph-database-service", attr.AssetImage("assets/alibabacloud/database/graph-database-service.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ApsaradbSqlserver(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-sqlserver", attr.AssetImage("assets/alibabacloud/database/apsaradb-sqlserver.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DataTransmissionService(opts ...attr.Attribute) *node.Node {
	return node.New("data-transmission-service", attr.AssetImage("assets/alibabacloud/database/data-transmission-service.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DatabaseBackupService(opts ...attr.Attribute) *node.Node {
	return node.New("database-backup-service", attr.AssetImage("assets/alibabacloud/database/database-backup-service.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) HybriddbForMysql(opts ...attr.Attribute) *node.Node {
	return node.New("hybriddb-for-mysql", attr.AssetImage("assets/alibabacloud/database/hybriddb-for-mysql.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) RelationalDatabaseService(opts ...attr.Attribute) *node.Node {
	return node.New("relational-database-service", attr.AssetImage("assets/alibabacloud/database/relational-database-service.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ApsaradbMongodb(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-mongodb", attr.AssetImage("assets/alibabacloud/database/apsaradb-mongodb.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ApsaradbOceanbase(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-oceanbase", attr.AssetImage("assets/alibabacloud/database/apsaradb-oceanbase.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ApsaradbPostgresql(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-postgresql", attr.AssetImage("assets/alibabacloud/database/apsaradb-postgresql.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ApsaradbRedis(opts ...attr.Attribute) *node.Node {
	return node.New("apsaradb-redis", attr.AssetImage("assets/alibabacloud/database/apsaradb-redis.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DisributeRelationalDatabaseService(opts ...attr.Attribute) *node.Node {
	return node.New("disribute-relational-database-service", attr.AssetImage("assets/alibabacloud/database/disribute-relational-database-service.png"), attr.Shape(attr.None))
}
