package gcp

import "github.com/blushft/go-diagrams/node"

type databaseContainer struct {
	path string
	opts []node.Option
}

var Database = &databaseContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/database",
}

func (c *databaseContainer) Bigtable(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/database/bigtable.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Datastore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/database/datastore.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Firestore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/database/firestore.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Memorystore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/database/memorystore.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Spanner(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/database/spanner.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Sql(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/database/sql.png")}, c.opts, opts)
	return node.New(nopts...)
}
