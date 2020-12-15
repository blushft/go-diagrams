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
	opts = append(opts, attr.AssetImage("assets/gcp/migration/transfer-appliance.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("transfer-appliance", opts...)
}

func init(){
	const namespace = "gcp.migration"
	Register(namespace, "TransferAppliance", GcpMigration.TransferAppliance)
}