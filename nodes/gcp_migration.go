package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type migrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var GcpMigration = &migrationContainer{path: "assets/gcp/migration"}

func (c *migrationContainer) TransferAppliance(opts ...attr.Attribute) *node.Node {
	return node.New("transfer-appliance", attr.AssetImage("assets/gcp/migration/transfer-appliance.png"), attr.Shape(attr.None))
}

func init(){
	const namespace = "gcp.migration"
	Register(namespace, "TransferAppliance", GcpMigration.TransferAppliance)
}