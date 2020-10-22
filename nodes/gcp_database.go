package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type gcpDatabaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var GcpDatabase = &gcpDatabaseContainer{path: "assets/gcp/database"}

func (c *gcpDatabaseContainer) Bigtable(opts ...attr.Attribute) *node.Node {
	return node.New("bigtable", attr.AssetImage("assets/gcp/database/bigtable.png"), attr.Shape(attr.None))
}

func (c *gcpDatabaseContainer) Datastore(opts ...attr.Attribute) *node.Node {
	return node.New("datastore", attr.AssetImage("assets/gcp/database/datastore.png"), attr.Shape(attr.None))
}

func (c *gcpDatabaseContainer) Firestore(opts ...attr.Attribute) *node.Node {
	return node.New("firestore", attr.AssetImage("assets/gcp/database/firestore.png"), attr.Shape(attr.None))
}

func (c *gcpDatabaseContainer) Memorystore(opts ...attr.Attribute) *node.Node {
	return node.New("memorystore", attr.AssetImage("assets/gcp/database/memorystore.png"), attr.Shape(attr.None))
}

func (c *gcpDatabaseContainer) Spanner(opts ...attr.Attribute) *node.Node {
	return node.New("spanner", attr.AssetImage("assets/gcp/database/spanner.png"), attr.Shape(attr.None))
}

func (c *gcpDatabaseContainer) Sql(opts ...attr.Attribute) *node.Node {
	return node.New("sql", attr.AssetImage("assets/gcp/database/sql.png"), attr.Shape(attr.None))
}

func init(){
	const namespace = "gcp.database"
	Register(namespace, "Bigtable", GcpDatabase.Bigtable)
	Register(namespace, "Datastore", GcpDatabase.Datastore)
	Register(namespace, "Firestore", GcpDatabase.Firestore)
	Register(namespace, "Memorystore", GcpDatabase.Memorystore)
	Register(namespace, "Spanner", GcpDatabase.Spanner)
	Register(namespace, "Sql", GcpDatabase.Sql)
}