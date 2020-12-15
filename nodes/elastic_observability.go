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
	opts = append(opts, attr.AssetImage("assets/elastic/observability/metrics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("metrics", opts...)
}

func (c *observabilityContainer) Observability(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/observability/observability.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("observability", opts...)
}

func (c *observabilityContainer) Uptime(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/observability/uptime.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("uptime", opts...)
}

func (c *observabilityContainer) Apm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/observability/apm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apm", opts...)
}

func (c *observabilityContainer) Logs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/observability/logs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("logs", opts...)
}
