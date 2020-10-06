package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type databaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var Database = &databaseContainer{path: "assets/aws/database"}

func (c *databaseContainer) Aurora(opts ...attr.Attribute) *node.Node {
	return node.New("aurora", attr.AssetImage("assets/aws/database/aurora.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) QuantumLedgerDatabaseQldb(opts ...attr.Attribute) *node.Node {
	return node.New("quantum-ledger-database-qldb", attr.AssetImage("assets/aws/database/quantum-ledger-database-qldb.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Timestream(opts ...attr.Attribute) *node.Node {
	return node.New("timestream", attr.AssetImage("assets/aws/database/timestream.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DynamodbGlobalSecondaryIndex(opts ...attr.Attribute) *node.Node {
	return node.New("dynamodb-global-secondary-index", attr.AssetImage("assets/aws/database/dynamodb-global-secondary-index.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DynamodbTable(opts ...attr.Attribute) *node.Node {
	return node.New("dynamodb-table", attr.AssetImage("assets/aws/database/dynamodb-table.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) RdsOnVmware(opts ...attr.Attribute) *node.Node {
	return node.New("rds-on-vmware", attr.AssetImage("assets/aws/database/rds-on-vmware.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DatabaseMigrationService(opts ...attr.Attribute) *node.Node {
	return node.New("database-migration-service", attr.AssetImage("assets/aws/database/database-migration-service.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Database(opts ...attr.Attribute) *node.Node {
	return node.New("database", attr.AssetImage("assets/aws/database/database.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DynamodbDax(opts ...attr.Attribute) *node.Node {
	return node.New("dynamodb-dax", attr.AssetImage("assets/aws/database/dynamodb-dax.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Rds(opts ...attr.Attribute) *node.Node {
	return node.New("rds", attr.AssetImage("assets/aws/database/rds.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Redshift(opts ...attr.Attribute) *node.Node {
	return node.New("redshift", attr.AssetImage("assets/aws/database/redshift.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DocumentdbMongodbCompatibility(opts ...attr.Attribute) *node.Node {
	return node.New("documentdb-mongodb-compatibility", attr.AssetImage("assets/aws/database/documentdb-mongodb-compatibility.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Dynamodb(opts ...attr.Attribute) *node.Node {
	return node.New("dynamodb", attr.AssetImage("assets/aws/database/dynamodb.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Elasticache(opts ...attr.Attribute) *node.Node {
	return node.New("elasticache", attr.AssetImage("assets/aws/database/elasticache.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Neptune(opts ...attr.Attribute) *node.Node {
	return node.New("neptune", attr.AssetImage("assets/aws/database/neptune.png"), attr.Shape(attr.None))
}
