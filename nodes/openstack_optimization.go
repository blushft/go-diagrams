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
	return node.New("rally", attr.AssetImage("assets/openstack/optimization/rally.png"), attr.Shape(attr.None))
}

func (c *optimizationContainer) Vitrage(opts ...attr.Attribute) *node.Node {
	return node.New("vitrage", attr.AssetImage("assets/openstack/optimization/vitrage.png"), attr.Shape(attr.None))
}

func (c *optimizationContainer) Watcher(opts ...attr.Attribute) *node.Node {
	return node.New("watcher", attr.AssetImage("assets/openstack/optimization/watcher.png"), attr.Shape(attr.None))
}

func (c *optimizationContainer) Congress(opts ...attr.Attribute) *node.Node {
	return node.New("congress", attr.AssetImage("assets/openstack/optimization/congress.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "openstack.optimization"
  Register(namespace, "Rally", OpenstackOptimization.Rally)
  Register(namespace, "Vitrage", OpenstackOptimization.Vitrage)
  Register(namespace, "Watcher", OpenstackOptimization.Watcher)
  Register(namespace, "Congress", OpenstackOptimization.Congress)
}
