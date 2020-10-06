package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type migrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Migration = &migrationContainer{path: "assets/aws/migration"}

func (c *migrationContainer) TransferForSftp(opts ...attr.Attribute) *node.Node {
	return node.New("transfer-for-sftp", attr.AssetImage("assets/aws/migration/transfer-for-sftp.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) ApplicationDiscoveryService(opts ...attr.Attribute) *node.Node {
	return node.New("application-discovery-service", attr.AssetImage("assets/aws/migration/application-discovery-service.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) DatabaseMigrationService(opts ...attr.Attribute) *node.Node {
	return node.New("database-migration-service", attr.AssetImage("assets/aws/migration/database-migration-service.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) MigrationHub(opts ...attr.Attribute) *node.Node {
	return node.New("migration-hub", attr.AssetImage("assets/aws/migration/migration-hub.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) SnowballEdge(opts ...attr.Attribute) *node.Node {
	return node.New("snowball-edge", attr.AssetImage("assets/aws/migration/snowball-edge.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) Snowball(opts ...attr.Attribute) *node.Node {
	return node.New("snowball", attr.AssetImage("assets/aws/migration/snowball.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) Snowmobile(opts ...attr.Attribute) *node.Node {
	return node.New("snowmobile", attr.AssetImage("assets/aws/migration/snowmobile.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) CloudendureMigration(opts ...attr.Attribute) *node.Node {
	return node.New("cloudendure-migration", attr.AssetImage("assets/aws/migration/cloudendure-migration.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) Datasync(opts ...attr.Attribute) *node.Node {
	return node.New("datasync", attr.AssetImage("assets/aws/migration/datasync.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) MigrationAndTransfer(opts ...attr.Attribute) *node.Node {
	return node.New("migration-and-transfer", attr.AssetImage("assets/aws/migration/migration-and-transfer.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) ServerMigrationService(opts ...attr.Attribute) *node.Node {
	return node.New("server-migration-service", attr.AssetImage("assets/aws/migration/server-migration-service.png"), attr.Shape(attr.None))
}
