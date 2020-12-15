package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type qualityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Quality = &qualityContainer{path: "assets/firebase/quality"}

func (c *qualityContainer) Crashlytics(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/quality/crashlytics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("crashlytics", opts...)
}

func (c *qualityContainer) PerformanceMonitoring(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/quality/performance-monitoring.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("performance-monitoring", opts...)
}

func (c *qualityContainer) TestLab(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/quality/test-lab.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("test-lab", opts...)
}

func (c *qualityContainer) AppDistribution(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/quality/app-distribution.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-distribution", opts...)
}

func (c *qualityContainer) CrashReporting(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/quality/crash-reporting.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("crash-reporting", opts...)
}
