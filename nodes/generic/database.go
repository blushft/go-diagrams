package generic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/node"
)

type databaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var Database = &databaseContainer{path: "assets/generic/database"}

func (c *databaseContainer) Sql(opts ...attr.Attribute) *node.Node {
	return node.New("sql", attr.AssetImage("assets/generic/database/sql.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "generic.database"
	diagram.Register(namespace, "Sql", Database.Sql)
}
