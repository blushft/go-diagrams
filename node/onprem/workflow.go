package onprem

import "github.com/blushft/go-diagrams/node"

type workflowContainer struct {
	path string
	opts []node.Option
}

var Workflow = &workflowContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/workflow",
}

func (c *workflowContainer) Digdag(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/workflow/digdag.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *workflowContainer) Kubeflow(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/workflow/kubeflow.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *workflowContainer) Nifi(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/workflow/nifi.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *workflowContainer) Airflow(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/workflow/airflow.png")}, c.opts, opts)
	return node.New(nopts...)
}
