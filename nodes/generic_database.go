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
	return node.New("sql", attr.AssetImage("assets/generic/database/sql.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "generic.database"
	Register(namespace, "Sql", GenericDatabase.Sql)
}
