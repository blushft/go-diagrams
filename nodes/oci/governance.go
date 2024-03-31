package oci

import "github.com/blushft/go-diagrams/diagram"

type governanceContainer struct {
	path string
	opts []diagram.NodeOption
}

var Governance = &governanceContainer{
	opts: diagram.OptionSet{diagram.Provider("oci"), diagram.NodeShape("none")},
	path: "assets/oci/governance",
}

func (c *governanceContainer) Audit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/audit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) CompartmentsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/compartments-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) LoggingWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/logging-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) OcidWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/ocid-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) Policies(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/policies.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) Compartments(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/compartments.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) Groups(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/groups.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) TaggingWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/tagging-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) Logging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/logging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) Ocid(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/ocid.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) AuditWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/audit-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) GroupsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/groups-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) PoliciesWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/policies-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *governanceContainer) Tagging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/governance/tagging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
