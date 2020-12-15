package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type optimizationContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackOptimization =&optimizationContainer{path: "assets/openstack/optimization"}

func (c *optimizationContainer) Rally(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/optimization/rally.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rally", opts...)
}

func (c *optimizationContainer) Vitrage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/optimization/vitrage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vitrage", opts...)
}

func (c *optimizationContainer) Watcher(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/optimization/watcher.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("watcher", opts...)
}

func (c *optimizationContainer) Congress(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/openstack/optimization/congress.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("congress", opts...)
}

func init() {
  const namespace = "openstack.optimization"
  Register(namespace, "Rally", OpenstackOptimization.Rally)
  Register(namespace, "Vitrage", OpenstackOptimization.Vitrage)
  Register(namespace, "Watcher", OpenstackOptimization.Watcher)
  Register(namespace, "Congress", OpenstackOptimization.Congress)
}
