package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type k8sComputeContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8sCompute = &k8sComputeContainer{path: "assets/k8s/compute"}

func (c *k8sComputeContainer) Pod(opts ...attr.Attribute) *node.Node {
	return node.New("pod", attr.AssetImage("assets/k8s/compute/pod.png"), attr.Shape(attr.None))
}

func (c *k8sComputeContainer) Rs(opts ...attr.Attribute) *node.Node {
	return node.New("rs", attr.AssetImage("assets/k8s/compute/rs.png"), attr.Shape(attr.None))
}

func (c *k8sComputeContainer) Sts(opts ...attr.Attribute) *node.Node {
	return node.New("sts", attr.AssetImage("assets/k8s/compute/sts.png"), attr.Shape(attr.None))
}

func (c *k8sComputeContainer) Cronjob(opts ...attr.Attribute) *node.Node {
	return node.New("cronjob", attr.AssetImage("assets/k8s/compute/cronjob.png"), attr.Shape(attr.None))
}

func (c *k8sComputeContainer) Deploy(opts ...attr.Attribute) *node.Node {
	return node.New("deploy", attr.AssetImage("assets/k8s/compute/deploy.png"), attr.Shape(attr.None))
}

func (c *k8sComputeContainer) Ds(opts ...attr.Attribute) *node.Node {
	return node.New("ds", attr.AssetImage("assets/k8s/compute/ds.png"), attr.Shape(attr.None))
}

func (c *k8sComputeContainer) Job(opts ...attr.Attribute) *node.Node {
	return node.New("job", attr.AssetImage("assets/k8s/compute/job.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "k8s.compute"
  Register(namespace, "Pod", K8sCompute.Pod)
  Register(namespace, "Rs", K8sCompute.Rs)
  Register(namespace, "Sts", K8sCompute.Sts)
  Register(namespace, "Cronjob", K8sCompute.Cronjob)
  Register(namespace, "Deploy", K8sCompute.Deploy)
  Register(namespace, "Ds", K8sCompute.Ds)
  Register(namespace, "Job", K8sCompute.Job)
}