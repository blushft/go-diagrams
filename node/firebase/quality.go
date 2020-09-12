package firebase

import "github.com/blushft/go-diagrams/node"

type qualityContainer struct {
	path string
	opts []node.Option
}

var Quality = &qualityContainer{
	opts: node.OptionSet{node.Provider("firebase"), node.Shape("none")},
	path: "assets/firebase/quality",
}

func (c *qualityContainer) PerformanceMonitoring(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/quality/performance-monitoring.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *qualityContainer) TestLab(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/quality/test-lab.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *qualityContainer) AppDistribution(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/quality/app-distribution.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *qualityContainer) CrashReporting(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/quality/crash-reporting.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *qualityContainer) Crashlytics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/quality/crashlytics.png")}, c.opts, opts)
	return node.New(nopts...)
}
