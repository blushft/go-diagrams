package generic

import "github.com/blushft/go-diagrams/node"

type databaseContainer struct {
	path string
	opts []node.Option
}

var Database = &databaseContainer{
	opts: node.OptionSet{node.Provider("generic"), node.Shape("none")},
	path: "assets/generic/database",
}

func (c *databaseContainer) Sql(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/generic/database/sql.png")}, c.opts, opts)
	return node.New(nopts...)
}
