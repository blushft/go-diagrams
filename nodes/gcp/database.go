package gcp

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type databaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var Database = &databaseContainer{path: "assets/gcp/database"}

func (c *databaseContainer) Bigtable(opts ...attr.Attribute) *node.Node {
	return node.New("bigtable", attr.AssetImage("assets/gcp/database/bigtable.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Datastore(opts ...attr.Attribute) *node.Node {
	return node.New("datastore", attr.AssetImage("assets/gcp/database/datastore.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Firestore(opts ...attr.Attribute) *node.Node {
	return node.New("firestore", attr.AssetImage("assets/gcp/database/firestore.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Memorystore(opts ...attr.Attribute) *node.Node {
	return node.New("memorystore", attr.AssetImage("assets/gcp/database/memorystore.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Spanner(opts ...attr.Attribute) *node.Node {
	return node.New("spanner", attr.AssetImage("assets/gcp/database/spanner.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Sql(opts ...attr.Attribute) *node.Node {
	return node.New("sql", attr.AssetImage("assets/gcp/database/sql.png"), attr.Shape(attr.None))
}
