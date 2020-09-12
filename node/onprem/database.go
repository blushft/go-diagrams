package onprem

import "github.com/blushft/go-diagrams/node"

type databaseContainer struct {
	path string
	opts []node.Option
}

var Database = &databaseContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/database",
}

func (c *databaseContainer) Clickhouse(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/clickhouse.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Cockroachdb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/cockroachdb.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Couchdb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/couchdb.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Dgraph(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/dgraph.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Mssql(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/mssql.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Mysql(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/mysql.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Neo4J(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/neo4j.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Cassandra(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/cassandra.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Influxdb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/influxdb.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Janusgraph(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/janusgraph.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Druid(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/druid.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Mongodb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/mongodb.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Oracle(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/oracle.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Postgresql(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/postgresql.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Scylla(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/scylla.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Couchbase(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/couchbase.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Hbase(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/hbase.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Mariadb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/database/mariadb.png")}, c.opts, opts)
	return node.New(nopts...)
}
