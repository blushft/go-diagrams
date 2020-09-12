package oci

import "github.com/blushft/go-diagrams/node"

type governanceContainer struct {
	path string
	opts []node.Option
}

var Governance = &governanceContainer{
	opts: node.OptionSet{node.Provider("oci"), node.Shape("none")},
	path: "assets/oci/governance",
}

func (c *governanceContainer) Tagging(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/tagging.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) Audit(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/audit.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) PoliciesWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/policies-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) Policies(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/policies.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) TaggingWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/tagging-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) LoggingWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/logging-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) Ocid(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/ocid.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) CompartmentsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/compartments-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) Compartments(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/compartments.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) GroupsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/groups-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) OcidWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/ocid-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) AuditWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/audit-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) Groups(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/groups.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *governanceContainer) Logging(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/governance/logging.png")}, c.opts, opts)
	return node.New(nopts...)
}
