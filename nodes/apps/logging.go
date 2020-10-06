package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type loggingContainer struct {
	path  string
	attrs []attr.Attribute
}

var Logging = &loggingContainer{path: "assets/apps/logging"}

func (c *loggingContainer) Loki(opts ...attr.Attribute) *node.Node {
	return node.New("loki", attr.AssetImage("assets/apps/logging/loki.png"), attr.Shape(attr.None))
}

func (c *loggingContainer) Rsyslog(opts ...attr.Attribute) *node.Node {
	return node.New("rsyslog", attr.AssetImage("assets/apps/logging/rsyslog.png"), attr.Shape(attr.None))
}

func (c *loggingContainer) SyslogNg(opts ...attr.Attribute) *node.Node {
	return node.New("syslog-ng", attr.AssetImage("assets/apps/logging/syslog-ng.png"), attr.Shape(attr.None))
}

func (c *loggingContainer) Fluentbit(opts ...attr.Attribute) *node.Node {
	return node.New("fluentbit", attr.AssetImage("assets/apps/logging/fluentbit.png"), attr.Shape(attr.None))
}

func (c *loggingContainer) Fluentd(opts ...attr.Attribute) *node.Node {
	return node.New("fluentd", attr.AssetImage("assets/apps/logging/fluentd.png"), attr.Shape(attr.None))
}

func (c *loggingContainer) Graylog(opts ...attr.Attribute) *node.Node {
	return node.New("graylog", attr.AssetImage("assets/apps/logging/graylog.png"), attr.Shape(attr.None))
}
