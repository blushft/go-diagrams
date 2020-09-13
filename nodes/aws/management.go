package aws

import "github.com/blushft/go-diagrams/diagram"

type managementContainer struct {
	path string
	opts []diagram.NodeOption
}

var Management = &managementContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/management",
}

func (c *managementContainer) Config(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/config.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ManagementConsole(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/management-console.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) WellArchitectedTool(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/well-architected-tool.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Cloudformation(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudformation.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) LicenseManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/license-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ManagedServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/managed-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ServiceCatalog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/service-catalog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) TrustedAdvisor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/trusted-advisor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Cloudtrail(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudtrail.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) ControlTower(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/control-tower.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Organizations(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/organizations.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) SystemsManagerParameterStore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/systems-manager-parameter-store.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) AutoScaling(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/auto-scaling.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Cloudwatch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/cloudwatch.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Codeguru(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/codeguru.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) CommandLineInterface(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/command-line-interface.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *managementContainer) Opsworks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/management/opsworks.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
