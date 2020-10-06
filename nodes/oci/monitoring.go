package oci

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type monitoringContainer struct {
	path  string
	attrs []attr.Attribute
}

var Monitoring = &monitoringContainer{path: "assets/oci/monitoring"}

func (c *monitoringContainer) Workflow(opts ...attr.Attribute) *node.Node {
	return node.New("workflow", attr.AssetImage("assets/oci/monitoring/workflow.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) HealthCheck(opts ...attr.Attribute) *node.Node {
	return node.New("health-check", attr.AssetImage("assets/oci/monitoring/health-check.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) SearchWhite(opts ...attr.Attribute) *node.Node {
	return node.New("search-white", attr.AssetImage("assets/oci/monitoring/search-white.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) NotificationsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("notifications-white", attr.AssetImage("assets/oci/monitoring/notifications-white.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) QueueWhite(opts ...attr.Attribute) *node.Node {
	return node.New("queue-white", attr.AssetImage("assets/oci/monitoring/queue-white.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) TelemetryWhite(opts ...attr.Attribute) *node.Node {
	return node.New("telemetry-white", attr.AssetImage("assets/oci/monitoring/telemetry-white.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) WorkflowWhite(opts ...attr.Attribute) *node.Node {
	return node.New("workflow-white", attr.AssetImage("assets/oci/monitoring/workflow-white.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) AlarmWhite(opts ...attr.Attribute) *node.Node {
	return node.New("alarm-white", attr.AssetImage("assets/oci/monitoring/alarm-white.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) EventsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("events-white", attr.AssetImage("assets/oci/monitoring/events-white.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Notifications(opts ...attr.Attribute) *node.Node {
	return node.New("notifications", attr.AssetImage("assets/oci/monitoring/notifications.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Queue(opts ...attr.Attribute) *node.Node {
	return node.New("queue", attr.AssetImage("assets/oci/monitoring/queue.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Search(opts ...attr.Attribute) *node.Node {
	return node.New("search", attr.AssetImage("assets/oci/monitoring/search.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Email(opts ...attr.Attribute) *node.Node {
	return node.New("email", attr.AssetImage("assets/oci/monitoring/email.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) HealthCheckWhite(opts ...attr.Attribute) *node.Node {
	return node.New("health-check-white", attr.AssetImage("assets/oci/monitoring/health-check-white.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Events(opts ...attr.Attribute) *node.Node {
	return node.New("events", attr.AssetImage("assets/oci/monitoring/events.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Telemetry(opts ...attr.Attribute) *node.Node {
	return node.New("telemetry", attr.AssetImage("assets/oci/monitoring/telemetry.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Alarm(opts ...attr.Attribute) *node.Node {
	return node.New("alarm", attr.AssetImage("assets/oci/monitoring/alarm.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) EmailWhite(opts ...attr.Attribute) *node.Node {
	return node.New("email-white", attr.AssetImage("assets/oci/monitoring/email-white.png"), attr.Shape(attr.None))
}
