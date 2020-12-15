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
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-mongodb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-mongodb", opts...)
}

func (c *alibabaCloudDatabase) ApsaradbOceanbase(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-oceanbase.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-oceanbase", opts...)
}

func (c *alibabaCloudDatabase) ApsaradbPostgresql(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-postgresql.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-postgresql", opts...)
}

func (c *alibabaCloudDatabase) ApsaradbPpas(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-ppas.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-ppas", opts...)
}

func (c *alibabaCloudDatabase) ApsaradbMemcache(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-memcache.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-memcache", opts...)
}

func (c *alibabaCloudDatabase) ApsaradbPolardb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-polardb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-polardb", opts...)
}

func (c *alibabaCloudDatabase) ApsaradbRedis(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-redis.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-redis", opts...)
}

func (c *alibabaCloudDatabase) DatabaseBackupService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/database-backup-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database-backup-service", opts...)
}

func (c *alibabaCloudDatabase) RelationalDatabaseService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/relational-database-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("relational-database-service", opts...)
}

func (c *alibabaCloudDatabase) ApsaradbCassandra(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-cassandra.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-cassandra", opts...)
}

func (c *alibabaCloudDatabase) ApsaradbHbase(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-hbase.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-hbase", opts...)
}

func (c *alibabaCloudDatabase) ApsaradbSqlserver(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/apsaradb-sqlserver.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apsaradb-sqlserver", opts...)
}

func (c *alibabaCloudDatabase) DataManagementService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/data-management-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-management-service", opts...)
}

func (c *alibabaCloudDatabase) GraphDatabaseService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/graph-database-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("graph-database-service", opts...)
}

func (c *alibabaCloudDatabase) DataTransmissionService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/data-transmission-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-transmission-service", opts...)
}

func (c *alibabaCloudDatabase) DisributeRelationalDatabaseService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/disribute-relational-database-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("disribute-relational-database-service", opts...)
}

func (c *alibabaCloudDatabase) HybriddbForMysql(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/database/hybriddb-for-mysql.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hybriddb-for-mysql", opts...)
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