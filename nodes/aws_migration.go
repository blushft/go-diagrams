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
	opts = append(opts, attr.AssetImage("assets/aws/migration/transfer-for-sftp.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("transfer-for-sftp", opts...)
}

func (c *awsMigrationContainer) ApplicationDiscoveryService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/application-discovery-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("application-discovery-service", opts...)
}

func (c *awsMigrationContainer) DatabaseMigrationService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/database-migration-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database-migration-service", opts...)
}

func (c *awsMigrationContainer) MigrationHub(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/migration-hub.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("migration-hub", opts...)
}

func (c *awsMigrationContainer) SnowballEdge(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/snowball-edge.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("snowball-edge", opts...)
}

func (c *awsMigrationContainer) Snowball(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/snowball.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("snowball", opts...)
}

func (c *awsMigrationContainer) Snowmobile(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/snowmobile.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("snowmobile", opts...)
}

func (c *awsMigrationContainer) CloudendureMigration(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/cloudendure-migration.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudendure-migration", opts...)
}

func (c *awsMigrationContainer) Datasync(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/datasync.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("datasync", opts...)
}

func (c *awsMigrationContainer) MigrationAndTransfer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/migration-and-transfer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("migration-and-transfer", opts...)
}

func (c *awsMigrationContainer) ServerMigrationService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/migration/server-migration-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("server-migration-service", opts...)
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
