package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuremigrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureMigration = &awsMigrationContainer{path: "assets/azure/migration"}

func (c *azuremigrationContainer) DatabaseMigrationServices(opts ...attr.Attribute) *node.Node {
	return node.New("database-migration-services", attr.AssetImage("assets/azure/migration/database-migration-services.png"), attr.Shape(attr.None))
}

func (c *azuremigrationContainer) MigrationProjects(opts ...attr.Attribute) *node.Node {
	return node.New("migration-projects", attr.AssetImage("assets/azure/migration/migration-projects.png"), attr.Shape(attr.None))
}

func (c *azuremigrationContainer) RecoveryServicesVaults(opts ...attr.Attribute) *node.Node {
	return node.New("recovery-services-vaults", attr.AssetImage("assets/azure/migration/recovery-services-vaults.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.migration"
  Register(namespace, "DatabaseMigrationServices", azureMigration.DatabaseMigrationServices)
  Register(namespace, "MigrationProjects", azureMigration.MigrationProjects)
  Register(namespace, "RecoveryServicesVaults", azureMigration.RecoveryServicesVaults)
  Register(namespace, "init", azureMigration.init)
}
