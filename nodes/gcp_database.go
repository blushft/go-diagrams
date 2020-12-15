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
	opts = append(opts, attr.AssetImage("assets/gcp/database/bigtable.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bigtable", opts...)
}

func (c *gcpDatabaseContainer) Datastore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/database/datastore.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("datastore", opts...)
}

func (c *gcpDatabaseContainer) Firestore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/database/firestore.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firestore", opts...)
}

func (c *gcpDatabaseContainer) Memorystore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/database/memorystore.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("memorystore", opts...)
}

func (c *gcpDatabaseContainer) Spanner(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/database/spanner.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("spanner", opts...)
}

func (c *gcpDatabaseContainer) Sql(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/database/sql.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sql", opts...)
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