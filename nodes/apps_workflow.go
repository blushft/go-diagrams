package nodes

import (
	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsWorkflowContainer struct {
	path  string
	attrs []attr.Attribute
}

var AppsWorkflow = &appsWorkflowContainer{path: "assets/apps/workflow"}

func (c *appsWorkflowContainer) Airflow(opts ...attr.Attribute) *node.Node {
	return node.New("airflow", attr.AssetImage("assets/apps/workflow/airflow.png"), attr.Shape(attr.None))
}

func (c *appsWorkflowContainer) Digdag(opts ...attr.Attribute) *node.Node {
	return node.New("digdag", attr.AssetImage("assets/apps/workflow/digdag.png"), attr.Shape(attr.None))
}

func (c *appsWorkflowContainer) Kubeflow(opts ...attr.Attribute) *node.Node {
	return node.New("kubeflow", attr.AssetImage("assets/apps/workflow/kubeflow.png"), attr.Shape(attr.None))
}

func (c *appsWorkflowContainer) Nifi(opts ...attr.Attribute) *node.Node {
	return node.New("nifi", attr.AssetImage("assets/apps/workflow/nifi.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.workflow"
	Register(namespace, "Airflow", AppsWorkflow.Airflow)
	Register(namespace, "Digdag", AppsWorkflow.Digdag)
	Register(namespace, "Kubeflow", AppsWorkflow.Kubeflow)
	Register(namespace, "Nifi", AppsWorkflow.Nifi)
}
