package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type databaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var Database = &databaseContainer{path: "assets/apps/database"}

func (c *databaseContainer) Scylla(opts ...attr.Attribute) *node.Node {
	return node.New("scylla", attr.AssetImage("assets/apps/database/scylla.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Influxdb(opts ...attr.Attribute) *node.Node {
	return node.New("influxdb", attr.AssetImage("assets/apps/database/influxdb.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Janusgraph(opts ...attr.Attribute) *node.Node {
	return node.New("janusgraph", attr.AssetImage("assets/apps/database/janusgraph.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Oracle(opts ...attr.Attribute) *node.Node {
	return node.New("oracle", attr.AssetImage("assets/apps/database/oracle.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Cassandra(opts ...attr.Attribute) *node.Node {
	return node.New("cassandra", attr.AssetImage("assets/apps/database/cassandra.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Mssql(opts ...attr.Attribute) *node.Node {
	return node.New("mssql", attr.AssetImage("assets/apps/database/mssql.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Mysql(opts ...attr.Attribute) *node.Node {
	return node.New("mysql", attr.AssetImage("assets/apps/database/mysql.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Postgresql(opts ...attr.Attribute) *node.Node {
	return node.New("postgresql", attr.AssetImage("assets/apps/database/postgresql.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Druid(opts ...attr.Attribute) *node.Node {
	return node.New("druid", attr.AssetImage("assets/apps/database/druid.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Hbase(opts ...attr.Attribute) *node.Node {
	return node.New("hbase", attr.AssetImage("assets/apps/database/hbase.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Neo4J(opts ...attr.Attribute) *node.Node {
	return node.New("neo4j", attr.AssetImage("assets/apps/database/neo4j.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Couchdb(opts ...attr.Attribute) *node.Node {
	return node.New("couchdb", attr.AssetImage("assets/apps/database/couchdb.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Dgraph(opts ...attr.Attribute) *node.Node {
	return node.New("dgraph", attr.AssetImage("assets/apps/database/dgraph.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Mariadb(opts ...attr.Attribute) *node.Node {
	return node.New("mariadb", attr.AssetImage("assets/apps/database/mariadb.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Mongodb(opts ...attr.Attribute) *node.Node {
	return node.New("mongodb", attr.AssetImage("assets/apps/database/mongodb.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Clickhouse(opts ...attr.Attribute) *node.Node {
	return node.New("clickhouse", attr.AssetImage("assets/apps/database/clickhouse.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Cockroachdb(opts ...attr.Attribute) *node.Node {
	return node.New("cockroachdb", attr.AssetImage("assets/apps/database/cockroachdb.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Couchbase(opts ...attr.Attribute) *node.Node {
	return node.New("couchbase", attr.AssetImage("assets/apps/database/couchbase.png"), attr.Shape(attr.None))
}
