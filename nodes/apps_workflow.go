package nodes

import (
	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type workflowContainer struct {
	path  string
	attrs []attr.Attribute
}

var Workflow = &workflowContainer{path: "assets/apps/workflow"}

func (c *workflowContainer) Airflow(opts ...attr.Attribute) *node.Node {
	return node.New("airflow", attr.AssetImage("assets/apps/workflow/airflow.png"), attr.Shape(attr.None))
}

func (c *workflowContainer) Digdag(opts ...attr.Attribute) *node.Node {
	return node.New("digdag", attr.AssetImage("assets/apps/workflow/digdag.png"), attr.Shape(attr.None))
}

func (c *workflowContainer) Kubeflow(opts ...attr.Attribute) *node.Node {
	return node.New("kubeflow", attr.AssetImage("assets/apps/workflow/kubeflow.png"), attr.Shape(attr.None))
}

func (c *workflowContainer) Nifi(opts ...attr.Attribute) *node.Node {
	return node.New("nifi", attr.AssetImage("assets/apps/workflow/nifi.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.workflow"
	Register(namespace, "Airflow", Workflow.Airflow)
	Register(namespace, "Digdag", Workflow.Digdag)
	Register(namespace, "Kubeflow", Workflow.Kubeflow)
	Register(namespace, "Nifi", Workflow.Nifi)
}
