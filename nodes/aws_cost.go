package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type costContainer struct {
	path  string
	attrs []attr.Attribute
}

var Cost = &costContainer{path: "assets/aws/cost"}

func (c *costContainer) ReservedInstanceReporting(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/cost/reserved-instance-reporting.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("reserved-instance-reporting", opts...)
}

func (c *costContainer) SavingsPlans(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/cost/savings-plans.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("savings-plans", opts...)
}

func (c *costContainer) Budgets(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/cost/budgets.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("budgets", opts...)
}

func (c *costContainer) CostAndUsageReport(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/cost/cost-and-usage-report.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cost-and-usage-report", opts...)
}

func (c *costContainer) CostExplorer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/cost/cost-explorer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cost-explorer", opts...)
}

func init() {
  const namespace = "aws.cost"
  Register(namespace, "ReservedInstanceReporting", Cost.ReservedInstanceReporting)
  Register(namespace, "SavingsPlans", Cost.SavingsPlans)
  Register(namespace, "Budgets", Cost.Budgets)
  Register(namespace, "CostAndUsageReport", Cost.CostAndUsageReport)
  Register(namespace, "CostExplorer", Cost.CostExplorer)
}
