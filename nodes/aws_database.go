package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsDatabaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSDatabase = &awsDatabaseContainer{path: "assets/aws/database"}

func (c *awsDatabaseContainer) Aurora(opts ...attr.Attribute) *node.Node {
	return node.New("aurora", attr.AssetImage("assets/aws/database/aurora.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) QuantumLedgerDatabaseQldb(opts ...attr.Attribute) *node.Node {
	return node.New("quantum-ledger-database-qldb", attr.AssetImage("assets/aws/database/quantum-ledger-database-qldb.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) Timestream(opts ...attr.Attribute) *node.Node {
	return node.New("timestream", attr.AssetImage("assets/aws/database/timestream.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) DynamodbGlobalSecondaryIndex(opts ...attr.Attribute) *node.Node {
	return node.New("dynamodb-global-secondary-index", attr.AssetImage("assets/aws/database/dynamodb-global-secondary-index.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) DynamodbTable(opts ...attr.Attribute) *node.Node {
	return node.New("dynamodb-table", attr.AssetImage("assets/aws/database/dynamodb-table.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) RdsOnVmware(opts ...attr.Attribute) *node.Node {
	return node.New("rds-on-vmware", attr.AssetImage("assets/aws/database/rds-on-vmware.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) DatabaseMigrationService(opts ...attr.Attribute) *node.Node {
	return node.New("database-migration-service", attr.AssetImage("assets/aws/database/database-migration-service.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) Database(opts ...attr.Attribute) *node.Node {
	return node.New("database", attr.AssetImage("assets/aws/database/database.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) DynamodbDax(opts ...attr.Attribute) *node.Node {
	return node.New("dynamodb-dax", attr.AssetImage("assets/aws/database/dynamodb-dax.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) Rds(opts ...attr.Attribute) *node.Node {
	return node.New("rds", attr.AssetImage("assets/aws/database/rds.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) Redshift(opts ...attr.Attribute) *node.Node {
	return node.New("redshift", attr.AssetImage("assets/aws/database/redshift.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) DocumentdbMongodbCompatibility(opts ...attr.Attribute) *node.Node {
	return node.New("documentdb-mongodb-compatibility", attr.AssetImage("assets/aws/database/documentdb-mongodb-compatibility.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) Dynamodb(opts ...attr.Attribute) *node.Node {
	return node.New("dynamodb", attr.AssetImage("assets/aws/database/dynamodb.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) Elasticache(opts ...attr.Attribute) *node.Node {
	return node.New("elasticache", attr.AssetImage("assets/aws/database/elasticache.png"), attr.Shape(attr.None))
}

func (c *awsDatabaseContainer) Neptune(opts ...attr.Attribute) *node.Node {
	return node.New("neptune", attr.AssetImage("assets/aws/database/neptune.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.database"
  Register(namespace, "Aurora", AWSDatabase.Aurora)
  Register(namespace, "QuantumLedgerDatabaseQldb", AWSDatabase.QuantumLedgerDatabaseQldb)
  Register(namespace, "Timestream", AWSDatabase.Timestream)
  Register(namespace, "DynamodbGlobalSecondaryIndex", AWSDatabase.DynamodbGlobalSecondaryIndex)
  Register(namespace, "DynamodbTable", AWSDatabase.DynamodbTable)
  Register(namespace, "RdsOnVmware", AWSDatabase.RdsOnVmware)
  Register(namespace, "DatabaseMigrationService", AWSDatabase.DatabaseMigrationService)
  Register(namespace, "Database", AWSDatabase.Database)
  Register(namespace, "DynamodbDax", AWSDatabase.DynamodbDax)
  Register(namespace, "Rds", AWSDatabase.Rds)
  Register(namespace, "Redshift", AWSDatabase.Redshift)
  Register(namespace, "DocumentdbMongodbCompatibility", AWSDatabase.DocumentdbMongodbCompatibility)
  Register(namespace, "Dynamodb", AWSDatabase.Dynamodb)
  Register(namespace, "Elasticache", AWSDatabase.Elasticache)
  Register(namespace, "Neptune", AWSDatabase.Neptune)
}
