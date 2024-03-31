package apps

import "github.com/blushft/go-diagrams/diagram"

type loggingContainer struct {
	path string
	opts []diagram.NodeOption
}

var Logging = &loggingContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/logging",
}

func (c *loggingContainer) Rsyslog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/logging/rsyslog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) SyslogNg(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/logging/syslog-ng.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) Fluentbit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/logging/fluentbit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) Fluentd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/logging/fluentd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) Graylog(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/logging/graylog.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *loggingContainer) Loki(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/logging/loki.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
