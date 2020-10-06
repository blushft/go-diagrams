package oci

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type governanceContainer struct {
	path  string
	attrs []attr.Attribute
}

var Governance = &governanceContainer{path: "assets/oci/governance"}

func (c *governanceContainer) TaggingWhite(opts ...attr.Attribute) *node.Node {
	return node.New("tagging-white", attr.AssetImage("assets/oci/governance/tagging-white.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) CompartmentsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("compartments-white", attr.AssetImage("assets/oci/governance/compartments-white.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) Groups(opts ...attr.Attribute) *node.Node {
	return node.New("groups", attr.AssetImage("assets/oci/governance/groups.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) LoggingWhite(opts ...attr.Attribute) *node.Node {
	return node.New("logging-white", attr.AssetImage("assets/oci/governance/logging-white.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) Ocid(opts ...attr.Attribute) *node.Node {
	return node.New("ocid", attr.AssetImage("assets/oci/governance/ocid.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) Audit(opts ...attr.Attribute) *node.Node {
	return node.New("audit", attr.AssetImage("assets/oci/governance/audit.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) OcidWhite(opts ...attr.Attribute) *node.Node {
	return node.New("ocid-white", attr.AssetImage("assets/oci/governance/ocid-white.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) Compartments(opts ...attr.Attribute) *node.Node {
	return node.New("compartments", attr.AssetImage("assets/oci/governance/compartments.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) GroupsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("groups-white", attr.AssetImage("assets/oci/governance/groups-white.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) Logging(opts ...attr.Attribute) *node.Node {
	return node.New("logging", attr.AssetImage("assets/oci/governance/logging.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) Policies(opts ...attr.Attribute) *node.Node {
	return node.New("policies", attr.AssetImage("assets/oci/governance/policies.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) AuditWhite(opts ...attr.Attribute) *node.Node {
	return node.New("audit-white", attr.AssetImage("assets/oci/governance/audit-white.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) PoliciesWhite(opts ...attr.Attribute) *node.Node {
	return node.New("policies-white", attr.AssetImage("assets/oci/governance/policies-white.png"), attr.Shape(attr.None))
}

func (c *governanceContainer) Tagging(opts ...attr.Attribute) *node.Node {
	return node.New("tagging", attr.AssetImage("assets/oci/governance/tagging.png"), attr.Shape(attr.None))
}
