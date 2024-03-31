package oci

import "github.com/blushft/go-diagrams/diagram"

type monitoringContainer struct {
	path string
	opts []diagram.NodeOption
}

var Monitoring = &monitoringContainer{
	opts: diagram.OptionSet{diagram.Provider("oci"), diagram.NodeShape("none")},
	path: "assets/oci/monitoring",
}

func (c *monitoringContainer) Search(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/search.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) TelemetryWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/telemetry-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Telemetry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/telemetry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Workflow(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/workflow.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Email(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/email.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) HealthCheckWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/health-check-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) QueueWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/queue-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Notifications(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/notifications.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) WorkflowWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/workflow-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) HealthCheck(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/health-check.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) SearchWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/search-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Alarm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/alarm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) EmailWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/email-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Events(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/events.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) Queue(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/queue.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) AlarmWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/alarm-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) EventsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/events-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *monitoringContainer) NotificationsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/monitoring/notifications-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
