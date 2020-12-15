package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type genericDatabaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var GenericDatabase = &genericDatabaseContainer{path: "assets/generic/database"}

func (c *genericDatabaseContainer) Sql(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/database/sql.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sql", opts...)
}

func init() {
	const namespace = "generic.database"
	Register(namespace, "Sql", GenericDatabase.Sql)
}
