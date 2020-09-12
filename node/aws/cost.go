package aws

import "github.com/blushft/go-diagrams/node"

type costContainer struct {
	path string
	opts []node.Option
}

var Cost = &costContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/cost",
}

func (c *costContainer) Budgets(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/cost/budgets.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *costContainer) CostAndUsageReport(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/cost/cost-and-usage-report.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *costContainer) CostExplorer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/cost/cost-explorer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *costContainer) ReservedInstanceReporting(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/cost/reserved-instance-reporting.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *costContainer) SavingsPlans(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/cost/savings-plans.png")}, c.opts, opts)
	return node.New(nopts...)
}
