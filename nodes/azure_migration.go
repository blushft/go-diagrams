package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuremigrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureMigration = &azuremigrationContainer{path: "assets/azure/migration"}

func (c *azuremigrationContainer) DatabaseMigrationServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/migration/database-migration-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database-migration-services", opts...)
}

func (c *azuremigrationContainer) MigrationProjects(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/migration/migration-projects.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("migration-projects", opts...)
}

func (c *azuremigrationContainer) RecoveryServicesVaults(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/migration/recovery-services-vaults.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("recovery-services-vaults", opts...)
}

func init() {
  const namespace = "azure.migration"
  Register(namespace, "DatabaseMigrationServices", azureMigration.DatabaseMigrationServices)
  Register(namespace, "MigrationProjects", azureMigration.MigrationProjects)
  Register(namespace, "RecoveryServicesVaults", azureMigration.RecoveryServicesVaults)
}
