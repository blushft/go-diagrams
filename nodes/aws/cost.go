package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type costContainer struct {
	path  string
	attrs []attr.Attribute
}

var Cost = &costContainer{path: "assets/aws/cost"}

func (c *costContainer) CostAndUsageReport(opts ...attr.Attribute) *node.Node {
	return node.New("cost-and-usage-report", attr.AssetImage("assets/aws/cost/cost-and-usage-report.png"), attr.Shape(attr.None))
}

func (c *costContainer) CostExplorer(opts ...attr.Attribute) *node.Node {
	return node.New("cost-explorer", attr.AssetImage("assets/aws/cost/cost-explorer.png"), attr.Shape(attr.None))
}

func (c *costContainer) ReservedInstanceReporting(opts ...attr.Attribute) *node.Node {
	return node.New("reserved-instance-reporting", attr.AssetImage("assets/aws/cost/reserved-instance-reporting.png"), attr.Shape(attr.None))
}

func (c *costContainer) SavingsPlans(opts ...attr.Attribute) *node.Node {
	return node.New("savings-plans", attr.AssetImage("assets/aws/cost/savings-plans.png"), attr.Shape(attr.None))
}

func (c *costContainer) Budgets(opts ...attr.Attribute) *node.Node {
	return node.New("budgets", attr.AssetImage("assets/aws/cost/budgets.png"), attr.Shape(attr.None))
}
