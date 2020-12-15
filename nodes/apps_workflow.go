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
	opts = append(opts, attr.AssetImage("assets/apps/workflow/airflow.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("airflow", opts...)
}

func (c *appsWorkflowContainer) Digdag(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/workflow/digdag.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("digdag", opts...)
}

func (c *appsWorkflowContainer) Kubeflow(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/workflow/kubeflow.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kubeflow", opts...)
}

func (c *appsWorkflowContainer) Nifi(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/workflow/nifi.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nifi", opts...)
}

func init() {
	const namespace = "apps.workflow"
	Register(namespace, "Airflow", AppsWorkflow.Airflow)
	Register(namespace, "Digdag", AppsWorkflow.Digdag)
	Register(namespace, "Kubeflow", AppsWorkflow.Kubeflow)
	Register(namespace, "Nifi", AppsWorkflow.Nifi)
}
