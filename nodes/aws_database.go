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
	opts = append(opts, attr.AssetImage("assets/aws/database/aurora.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("aurora", opts...)
}

func (c *awsDatabaseContainer) QuantumLedgerDatabaseQldb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/quantum-ledger-database-qldb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("quantum-ledger-database-qldb", opts...)
}

func (c *awsDatabaseContainer) Timestream(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/timestream.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("timestream", opts...)
}

func (c *awsDatabaseContainer) DynamodbGlobalSecondaryIndex(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/dynamodb-global-secondary-index.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dynamodb-global-secondary-index", opts...)
}

func (c *awsDatabaseContainer) DynamodbTable(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/dynamodb-table.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dynamodb-table", opts...)
}

func (c *awsDatabaseContainer) RdsOnVmware(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/rds-on-vmware.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rds-on-vmware", opts...)
}

func (c *awsDatabaseContainer) DatabaseMigrationService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/database-migration-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database-migration-service", opts...)
}

func (c *awsDatabaseContainer) Database(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/database.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database", opts...)
}

func (c *awsDatabaseContainer) DynamodbDax(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/dynamodb-dax.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dynamodb-dax", opts...)
}

func (c *awsDatabaseContainer) Rds(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/rds.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rds", opts...)
}

func (c *awsDatabaseContainer) Redshift(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/redshift.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("redshift", opts...)
}

func (c *awsDatabaseContainer) DocumentdbMongodbCompatibility(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/documentdb-mongodb-compatibility.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("documentdb-mongodb-compatibility", opts...)
}

func (c *awsDatabaseContainer) Dynamodb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/dynamodb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dynamodb", opts...)
}

func (c *awsDatabaseContainer) Elasticache(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/elasticache.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elasticache", opts...)
}

func (c *awsDatabaseContainer) Neptune(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/database/neptune.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("neptune", opts...)
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
