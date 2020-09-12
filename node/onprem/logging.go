package onprem

import "github.com/blushft/go-diagrams/node"

type loggingContainer struct {
	path string
	opts []node.Option
}

var Logging = &loggingContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/logging",
}

func (c *loggingContainer) SyslogNg(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/logging/syslog-ng.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *loggingContainer) Fluentbit(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/logging/fluentbit.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *loggingContainer) Fluentd(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/logging/fluentd.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *loggingContainer) Graylog(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/logging/graylog.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *loggingContainer) Loki(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/logging/loki.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *loggingContainer) Rsyslog(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/logging/rsyslog.png")}, c.opts, opts)
	return node.New(nopts...)
}
