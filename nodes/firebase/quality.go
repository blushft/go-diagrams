package firebase

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
	return node.New("crashlytics", attr.AssetImage("assets/firebase/quality/crashlytics.png"), attr.Shape(attr.None))
}

func (c *qualityContainer) PerformanceMonitoring(opts ...attr.Attribute) *node.Node {
	return node.New("performance-monitoring", attr.AssetImage("assets/firebase/quality/performance-monitoring.png"), attr.Shape(attr.None))
}

func (c *qualityContainer) TestLab(opts ...attr.Attribute) *node.Node {
	return node.New("test-lab", attr.AssetImage("assets/firebase/quality/test-lab.png"), attr.Shape(attr.None))
}

func (c *qualityContainer) AppDistribution(opts ...attr.Attribute) *node.Node {
	return node.New("app-distribution", attr.AssetImage("assets/firebase/quality/app-distribution.png"), attr.Shape(attr.None))
}

func (c *qualityContainer) CrashReporting(opts ...attr.Attribute) *node.Node {
	return node.New("crash-reporting", attr.AssetImage("assets/firebase/quality/crash-reporting.png"), attr.Shape(attr.None))
}
