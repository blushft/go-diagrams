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
	opts = append(opts, attr.AssetImage("assets/apps/database/influxdb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("influxdb", opts...)
}

func (c *appsDatabaseContainer) Mariadb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/mariadb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mariadb", opts...)
}

func (c *appsDatabaseContainer) Postgresql(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/postgresql.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("postgresql", opts...)
}

func (c *appsDatabaseContainer) Mysql(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/mysql.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mysql", opts...)
}

func (c *appsDatabaseContainer) Oracle(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/oracle.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("oracle", opts...)
}

func (c *appsDatabaseContainer) Scylla(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/scylla.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("scylla", opts...)
}

func (c *appsDatabaseContainer) Couchbase(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/couchbase.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("couchbase", opts...)
}

func (c *appsDatabaseContainer) Couchdb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/couchdb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("couchdb", opts...)
}

func (c *appsDatabaseContainer) Hbase(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/hbase.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hbase", opts...)
}

func (c *appsDatabaseContainer) Janusgraph(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/janusgraph.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("janusgraph", opts...)
}

func (c *appsDatabaseContainer) Cassandra(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/cassandra.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cassandra", opts...)
}

func (c *appsDatabaseContainer) Clickhouse(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/clickhouse.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("clickhouse", opts...)
}

func (c *appsDatabaseContainer) Cockroachdb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/cockroachdb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cockroachdb", opts...)
}

func (c *appsDatabaseContainer) Dgraph(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/dgraph.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dgraph", opts...)
}

func (c *appsDatabaseContainer) Druid(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/druid.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("druid", opts...)
}

func (c *appsDatabaseContainer) Mongodb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/mongodb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mongodb", opts...)
}

func (c *appsDatabaseContainer) Mssql(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/mssql.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mssql", opts...)
}

func (c *appsDatabaseContainer) Neo4J(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/database/neo4j.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("neo4j", opts...)
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
