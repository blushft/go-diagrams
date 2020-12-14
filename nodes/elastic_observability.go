package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type observabilityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Observability = &observabilityContainer{path: "assets/elastic/observability"}

func (c *observabilityContainer) Metrics(opts ...attr.Attribute) *node.Node {
	return node.New("metrics", attr.AssetImage("assets/elastic/observability/metrics.png"), attr.Shape(attr.None))
}

func (c *observabilityContainer) Observability(opts ...attr.Attribute) *node.Node {
	return node.New("observability", attr.AssetImage("assets/elastic/observability/observability.png"), attr.Shape(attr.None))
}

func (c *observabilityContainer) Uptime(opts ...attr.Attribute) *node.Node {
	return node.New("uptime", attr.AssetImage("assets/elastic/observability/uptime.png"), attr.Shape(attr.None))
}

func (c *observabilityContainer) Apm(opts ...attr.Attribute) *node.Node {
	return node.New("apm", attr.AssetImage("assets/elastic/observability/apm.png"), attr.Shape(attr.None))
}

func (c *observabilityContainer) Logs(opts ...attr.Attribute) *node.Node {
	return node.New("logs", attr.AssetImage("assets/elastic/observability/logs.png"), attr.Shape(attr.None))
}
