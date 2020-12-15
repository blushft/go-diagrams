package nodes

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
	opts = append(opts, attr.AssetImage("assets/apps/logging/loki.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("loki", opts...)
}

func (c *loggingContainer) Rsyslog(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/logging/rsyslog.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rsyslog", opts...)
}

func (c *loggingContainer) SyslogNg(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/logging/syslog-ng.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("syslog-ng", opts...)
}

func (c *loggingContainer) Fluentbit(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/logging/fluentbit.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("fluentbit", opts...)
}

func (c *loggingContainer) Fluentd(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/logging/fluentd.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("fluentd", opts...)
}

func (c *loggingContainer) Graylog(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/logging/graylog.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("graylog", opts...)
}
func init() {
	const namespace = "apps.logging"
	Register(namespace, "Loki", Logging.Loki)
	Register(namespace, "Rsyslog", Logging.Rsyslog)
	Register(namespace, "SyslogNg", Logging.SyslogNg)
	Register(namespace, "Fluentbit", Logging.Fluentbit)
	Register(namespace, "Fluentd", Logging.Fluentd)
	Register(namespace, "Graylog", Logging.Graylog)
}