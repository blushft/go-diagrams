package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsMigrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Migration = &awsMigrationContainer{path: "assets/aws/migration"}

func (c *awsMigrationContainer) TransferForSftp(opts ...attr.Attribute) *node.Node {
	return node.New("transfer-for-sftp", attr.AssetImage("assets/aws/migration/transfer-for-sftp.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) ApplicationDiscoveryService(opts ...attr.Attribute) *node.Node {
	return node.New("application-discovery-service", attr.AssetImage("assets/aws/migration/application-discovery-service.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) DatabaseMigrationService(opts ...attr.Attribute) *node.Node {
	return node.New("database-migration-service", attr.AssetImage("assets/aws/migration/database-migration-service.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) MigrationHub(opts ...attr.Attribute) *node.Node {
	return node.New("migration-hub", attr.AssetImage("assets/aws/migration/migration-hub.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) SnowballEdge(opts ...attr.Attribute) *node.Node {
	return node.New("snowball-edge", attr.AssetImage("assets/aws/migration/snowball-edge.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) Snowball(opts ...attr.Attribute) *node.Node {
	return node.New("snowball", attr.AssetImage("assets/aws/migration/snowball.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) Snowmobile(opts ...attr.Attribute) *node.Node {
	return node.New("snowmobile", attr.AssetImage("assets/aws/migration/snowmobile.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) CloudendureMigration(opts ...attr.Attribute) *node.Node {
	return node.New("cloudendure-migration", attr.AssetImage("assets/aws/migration/cloudendure-migration.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) Datasync(opts ...attr.Attribute) *node.Node {
	return node.New("datasync", attr.AssetImage("assets/aws/migration/datasync.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) MigrationAndTransfer(opts ...attr.Attribute) *node.Node {
	return node.New("migration-and-transfer", attr.AssetImage("assets/aws/migration/migration-and-transfer.png"), attr.Shape(attr.None))
}

func (c *awsMigrationContainer) ServerMigrationService(opts ...attr.Attribute) *node.Node {
	return node.New("server-migration-service", attr.AssetImage("assets/aws/migration/server-migration-service.png"), attr.Shape(attr.None))
}

func init() {
  const namespace= "aws.migration"
  Register(namespace, "TransferForSftp", Migration.TransferForSftp)
  Register(namespace, "ApplicationDiscoveryService", Migration.ApplicationDiscoveryService)
  Register(namespace, "DatabaseMigrationService", Migration.DatabaseMigrationService)
  Register(namespace, "MigrationHub", Migration.MigrationHub)
  Register(namespace, "SnowballEdge", Migration.SnowballEdge)
  Register(namespace, "Snowball", Migration.Snowball)
  Register(namespace, "Snowmobile", Migration.Snowmobile)
  Register(namespace, "CloudendureMigration", Migration.CloudendureMigration)
  Register(namespace, "Datasync", Migration.Datasync)
  Register(namespace, "MigrationAndTransfer", Migration.MigrationAndTransfer)
  Register(namespace, "ServerMigrationService", Migration.ServerMigrationService)
}
