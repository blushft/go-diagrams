package oci

import "github.com/blushft/go-diagrams/node"

type monitoringContainer struct {
	path string
	opts []node.Option
}

var Monitoring = &monitoringContainer{
	opts: node.OptionSet{node.Provider("oci"), node.Shape("none")},
	path: "assets/oci/monitoring",
}

func (c *monitoringContainer) Events(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/events.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Search(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/search.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Alarm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/alarm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Email(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/email.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) HealthCheck(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/health-check.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) SearchWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/search-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Workflow(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/workflow.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) EmailWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/email-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Notifications(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/notifications.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) QueueWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/queue-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Queue(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/queue.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) Telemetry(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/telemetry.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) WorkflowWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/workflow-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) AlarmWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/alarm-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) EventsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/events-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) HealthCheckWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/health-check-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) NotificationsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/notifications-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *monitoringContainer) TelemetryWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/monitoring/telemetry-white.png")}, c.opts, opts)
	return node.New(nopts...)
}
