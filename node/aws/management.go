package aws

import "github.com/blushft/go-diagrams/node"

type managementContainer struct {
	path string
	opts []node.Option
}

var Management = &managementContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/management",
}

func (c *managementContainer) Codeguru(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/codeguru.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) CommandLineInterface(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/command-line-interface.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) Opsworks(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/opsworks.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) Cloudtrail(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/cloudtrail.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) Cloudformation(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/cloudformation.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) ControlTower(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/control-tower.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) Organizations(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/organizations.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) ServiceCatalog(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/service-catalog.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) WellArchitectedTool(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/well-architected-tool.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) AutoScaling(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/auto-scaling.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) LicenseManager(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/license-manager.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) SystemsManager(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/systems-manager.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) TrustedAdvisor(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/trusted-advisor.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) Cloudwatch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/cloudwatch.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) ManagedServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/managed-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) ManagementConsole(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/management-console.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) SystemsManagerParameterStore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/systems-manager-parameter-store.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *managementContainer) Config(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/management/config.png")}, c.opts, opts)
	return node.New(nopts...)
}
