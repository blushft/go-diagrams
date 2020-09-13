package gcp

import "github.com/blushft/go-diagrams/diagram"

type databaseContainer struct {
	path string
	opts []diagram.NodeOption
}

var Database = &databaseContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/database",
}

func (c *databaseContainer) Datastore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/database/datastore.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Firestore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/database/firestore.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Memorystore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/database/memorystore.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Spanner(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/database/spanner.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Sql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/database/sql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Bigtable(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/database/bigtable.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
