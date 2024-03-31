package aws

import "github.com/blushft/go-diagrams/diagram"

type costContainer struct {
	path string
	opts []diagram.NodeOption
}

var Cost = &costContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/cost",
}

func (c *costContainer) Budgets(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/cost/budgets.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *costContainer) CostAndUsageReport(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/cost/cost-and-usage-report.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *costContainer) CostExplorer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/cost/cost-explorer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *costContainer) ReservedInstanceReporting(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/cost/reserved-instance-reporting.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *costContainer) SavingsPlans(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/cost/savings-plans.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
