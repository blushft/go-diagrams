package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type migrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Migration = &migrationContainer{path: "assets/azure/migration"}

func (c *migrationContainer) DatabaseMigrationServices(opts ...attr.Attribute) *node.Node {
	return node.New("database-migration-services", attr.AssetImage("assets/azure/migration/database-migration-services.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) MigrationProjects(opts ...attr.Attribute) *node.Node {
	return node.New("migration-projects", attr.AssetImage("assets/azure/migration/migration-projects.png"), attr.Shape(attr.None))
}

func (c *migrationContainer) RecoveryServicesVaults(opts ...attr.Attribute) *node.Node {
	return node.New("recovery-services-vaults", attr.AssetImage("assets/azure/migration/recovery-services-vaults.png"), attr.Shape(attr.None))
}
