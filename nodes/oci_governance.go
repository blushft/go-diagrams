package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ociGovernanceContainer struct {
	path  string
	attrs []attr.Attribute
}

var OciGovernance = &ociGovernanceContainer{path: "assets/oci/governance"}

func (c *ociGovernanceContainer) TaggingWhite(opts ...attr.Attribute) *node.Node {
	return node.New("tagging-white", attr.AssetImage("assets/oci/governance/tagging-white.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) CompartmentsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("compartments-white", attr.AssetImage("assets/oci/governance/compartments-white.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) Groups(opts ...attr.Attribute) *node.Node {
	return node.New("groups", attr.AssetImage("assets/oci/governance/groups.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) LoggingWhite(opts ...attr.Attribute) *node.Node {
	return node.New("logging-white", attr.AssetImage("assets/oci/governance/logging-white.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) Ocid(opts ...attr.Attribute) *node.Node {
	return node.New("ocid", attr.AssetImage("assets/oci/governance/ocid.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) Audit(opts ...attr.Attribute) *node.Node {
	return node.New("audit", attr.AssetImage("assets/oci/governance/audit.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) OcidWhite(opts ...attr.Attribute) *node.Node {
	return node.New("ocid-white", attr.AssetImage("assets/oci/governance/ocid-white.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) Compartments(opts ...attr.Attribute) *node.Node {
	return node.New("compartments", attr.AssetImage("assets/oci/governance/compartments.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) GroupsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("groups-white", attr.AssetImage("assets/oci/governance/groups-white.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) Logging(opts ...attr.Attribute) *node.Node {
	return node.New("logging", attr.AssetImage("assets/oci/governance/logging.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) Policies(opts ...attr.Attribute) *node.Node {
	return node.New("policies", attr.AssetImage("assets/oci/governance/policies.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) AuditWhite(opts ...attr.Attribute) *node.Node {
	return node.New("audit-white", attr.AssetImage("assets/oci/governance/audit-white.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) PoliciesWhite(opts ...attr.Attribute) *node.Node {
	return node.New("policies-white", attr.AssetImage("assets/oci/governance/policies-white.png"), attr.Shape(attr.None))
}

func (c *ociGovernanceContainer) Tagging(opts ...attr.Attribute) *node.Node {
	return node.New("tagging", attr.AssetImage("assets/oci/governance/tagging.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "oci.governance"
  Register(namespace, "TaggingWhite", OciGovernance.TaggingWhite)
  Register(namespace, "CompartmentsWhite", OciGovernance.CompartmentsWhite)
  Register(namespace, "Groups", OciGovernance.Groups)
  Register(namespace, "LoggingWhite", OciGovernance.LoggingWhite)
  Register(namespace, "Ocid", OciGovernance.Ocid)
  Register(namespace, "Audit", OciGovernance.Audit)
  Register(namespace, "OcidWhite", OciGovernance.OcidWhite)
  Register(namespace, "Compartments", OciGovernance.Compartments)
  Register(namespace, "GroupsWhite", OciGovernance.GroupsWhite)
  Register(namespace, "Logging", OciGovernance.Logging)
  Register(namespace, "Policies", OciGovernance.Policies)
  Register(namespace, "AuditWhite", OciGovernance.AuditWhite)
  Register(namespace, "PoliciesWhite", OciGovernance.PoliciesWhite)
  Register(namespace, "Tagging", OciGovernance.Tagging)
}
