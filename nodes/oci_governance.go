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
	opts = append(opts, attr.AssetImage("assets/oci/governance/tagging-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tagging-white", opts...)
}

func (c *ociGovernanceContainer) CompartmentsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/compartments-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("compartments-white", opts...)
}

func (c *ociGovernanceContainer) Groups(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/groups.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("groups", opts...)
}

func (c *ociGovernanceContainer) LoggingWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/logging-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("logging-white", opts...)
}

func (c *ociGovernanceContainer) Ocid(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/ocid.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ocid", opts...)
}

func (c *ociGovernanceContainer) Audit(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/audit.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("audit", opts...)
}

func (c *ociGovernanceContainer) OcidWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/ocid-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ocid-white", opts...)
}

func (c *ociGovernanceContainer) Compartments(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/compartments.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("compartments", opts...)
}

func (c *ociGovernanceContainer) GroupsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/groups-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("groups-white", opts...)
}

func (c *ociGovernanceContainer) Logging(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/logging.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("logging", opts...)
}

func (c *ociGovernanceContainer) Policies(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/policies.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("policies", opts...)
}

func (c *ociGovernanceContainer) AuditWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/audit-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("audit-white", opts...)
}

func (c *ociGovernanceContainer) PoliciesWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/policies-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("policies-white", opts...)
}

func (c *ociGovernanceContainer) Tagging(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/governance/tagging.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tagging", opts...)
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
