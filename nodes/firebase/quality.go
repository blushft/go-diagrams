package firebase

import "github.com/blushft/go-diagrams/diagram"

type qualityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Quality = &qualityContainer{
	opts: diagram.OptionSet{diagram.Provider("firebase"), diagram.NodeShape("none")},
	path: "assets/firebase/quality",
}

func (c *qualityContainer) AppDistribution(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/quality/app-distribution.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *qualityContainer) CrashReporting(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/quality/crash-reporting.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *qualityContainer) Crashlytics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/quality/crashlytics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *qualityContainer) PerformanceMonitoring(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/quality/performance-monitoring.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *qualityContainer) TestLab(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/quality/test-lab.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
