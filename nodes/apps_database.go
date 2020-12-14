package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsDatabaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var AppsDatabase = &appsDatabaseContainer{path: "assets/apps/database"}

func (c *appsDatabaseContainer) Influxdb(opts ...attr.Attribute) *node.Node {
	return node.New("influxdb", attr.AssetImage("assets/apps/database/influxdb.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Mariadb(opts ...attr.Attribute) *node.Node {
	return node.New("mariadb", attr.AssetImage("assets/apps/database/mariadb.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Postgresql(opts ...attr.Attribute) *node.Node {
	return node.New("postgresql", attr.AssetImage("assets/apps/database/postgresql.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Mysql(opts ...attr.Attribute) *node.Node {
	return node.New("mysql", attr.AssetImage("assets/apps/database/mysql.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Oracle(opts ...attr.Attribute) *node.Node {
	return node.New("oracle", attr.AssetImage("assets/apps/database/oracle.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Scylla(opts ...attr.Attribute) *node.Node {
	return node.New("scylla", attr.AssetImage("assets/apps/database/scylla.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Couchbase(opts ...attr.Attribute) *node.Node {
	return node.New("couchbase", attr.AssetImage("assets/apps/database/couchbase.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Couchdb(opts ...attr.Attribute) *node.Node {
	return node.New("couchdb", attr.AssetImage("assets/apps/database/couchdb.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Hbase(opts ...attr.Attribute) *node.Node {
	return node.New("hbase", attr.AssetImage("assets/apps/database/hbase.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Janusgraph(opts ...attr.Attribute) *node.Node {
	return node.New("janusgraph", attr.AssetImage("assets/apps/database/janusgraph.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Cassandra(opts ...attr.Attribute) *node.Node {
	return node.New("cassandra", attr.AssetImage("assets/apps/database/cassandra.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Clickhouse(opts ...attr.Attribute) *node.Node {
	return node.New("clickhouse", attr.AssetImage("assets/apps/database/clickhouse.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Cockroachdb(opts ...attr.Attribute) *node.Node {
	return node.New("cockroachdb", attr.AssetImage("assets/apps/database/cockroachdb.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Dgraph(opts ...attr.Attribute) *node.Node {
	return node.New("dgraph", attr.AssetImage("assets/apps/database/dgraph.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Druid(opts ...attr.Attribute) *node.Node {
	return node.New("druid", attr.AssetImage("assets/apps/database/druid.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Mongodb(opts ...attr.Attribute) *node.Node {
	return node.New("mongodb", attr.AssetImage("assets/apps/database/mongodb.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Mssql(opts ...attr.Attribute) *node.Node {
	return node.New("mssql", attr.AssetImage("assets/apps/database/mssql.png"), attr.Shape(attr.None))
}

func (c *appsDatabaseContainer) Neo4J(opts ...attr.Attribute) *node.Node {
	return node.New("neo4j", attr.AssetImage("assets/apps/database/neo4j.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.database"
	Register(namespace, "Influxdb", AppsDatabase.Influxdb)
	Register(namespace, "Mariadb", AppsDatabase.Mariadb)
	Register(namespace, "Postgresql", AppsDatabase.Postgresql)
	Register(namespace, "Mysql", AppsDatabase.Mysql)
	Register(namespace, "Oracle", AppsDatabase.Oracle)
	Register(namespace, "Scylla", AppsDatabase.Scylla)
	Register(namespace, "Couchbase", AppsDatabase.Couchbase)
	Register(namespace, "Couchdb", AppsDatabase.Couchdb)
	Register(namespace, "Hbase", AppsDatabase.Hbase)
	Register(namespace, "Janusgraph", AppsDatabase.Janusgraph)
	Register(namespace, "Cassandra", AppsDatabase.Cassandra)
	Register(namespace, "Clickhouse", AppsDatabase.Clickhouse)
	Register(namespace, "Cockroachdb", AppsDatabase.Cockroachdb)
	Register(namespace, "Dgraph", AppsDatabase.Dgraph)
	Register(namespace, "Druid", AppsDatabase.Druid)
	Register(namespace, "Mongodb", AppsDatabase.Mongodb)
	Register(namespace, "Mssql", AppsDatabase.Mssql)
	Register(namespace, "Neo4J", AppsDatabase.Neo4J)
}
