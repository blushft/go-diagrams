package apps

import "github.com/blushft/go-diagrams/diagram"

type workflowContainer struct {
	path string
	opts []diagram.NodeOption
}

var Workflow = &workflowContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/workflow",
}

func (c *workflowContainer) Digdag(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/workflow/digdag.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *workflowContainer) Kubeflow(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/workflow/kubeflow.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *workflowContainer) Nifi(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/workflow/nifi.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *workflowContainer) Airflow(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/workflow/airflow.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
