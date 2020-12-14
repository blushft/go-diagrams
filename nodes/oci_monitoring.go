package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ociMonitoringContainer struct {
	path  string
	attrs []attr.Attribute
}

var OciMonitoring = &ociMonitoringContainer{path: "assets/oci/monitoring"}

func (c *ociMonitoringContainer) Workflow(opts ...attr.Attribute) *node.Node {
	return node.New("workflow", attr.AssetImage("assets/oci/monitoring/workflow.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) HealthCheck(opts ...attr.Attribute) *node.Node {
	return node.New("health-check", attr.AssetImage("assets/oci/monitoring/health-check.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) SearchWhite(opts ...attr.Attribute) *node.Node {
	return node.New("search-white", attr.AssetImage("assets/oci/monitoring/search-white.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) NotificationsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("notifications-white", attr.AssetImage("assets/oci/monitoring/notifications-white.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) QueueWhite(opts ...attr.Attribute) *node.Node {
	return node.New("queue-white", attr.AssetImage("assets/oci/monitoring/queue-white.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) TelemetryWhite(opts ...attr.Attribute) *node.Node {
	return node.New("telemetry-white", attr.AssetImage("assets/oci/monitoring/telemetry-white.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) WorkflowWhite(opts ...attr.Attribute) *node.Node {
	return node.New("workflow-white", attr.AssetImage("assets/oci/monitoring/workflow-white.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) AlarmWhite(opts ...attr.Attribute) *node.Node {
	return node.New("alarm-white", attr.AssetImage("assets/oci/monitoring/alarm-white.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) EventsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("events-white", attr.AssetImage("assets/oci/monitoring/events-white.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) Notifications(opts ...attr.Attribute) *node.Node {
	return node.New("notifications", attr.AssetImage("assets/oci/monitoring/notifications.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) Queue(opts ...attr.Attribute) *node.Node {
	return node.New("queue", attr.AssetImage("assets/oci/monitoring/queue.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) Search(opts ...attr.Attribute) *node.Node {
	return node.New("search", attr.AssetImage("assets/oci/monitoring/search.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) Email(opts ...attr.Attribute) *node.Node {
	return node.New("email", attr.AssetImage("assets/oci/monitoring/email.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) HealthCheckWhite(opts ...attr.Attribute) *node.Node {
	return node.New("health-check-white", attr.AssetImage("assets/oci/monitoring/health-check-white.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) Events(opts ...attr.Attribute) *node.Node {
	return node.New("events", attr.AssetImage("assets/oci/monitoring/events.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) Telemetry(opts ...attr.Attribute) *node.Node {
	return node.New("telemetry", attr.AssetImage("assets/oci/monitoring/telemetry.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) Alarm(opts ...attr.Attribute) *node.Node {
	return node.New("alarm", attr.AssetImage("assets/oci/monitoring/alarm.png"), attr.Shape(attr.None))
}

func (c *ociMonitoringContainer) EmailWhite(opts ...attr.Attribute) *node.Node {
	return node.New("email-white", attr.AssetImage("assets/oci/monitoring/email-white.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "oci.monitoring"
  Register(namespace, "Workflow", OciMonitoring.Workflow)
  Register(namespace, "HealthCheck", OciMonitoring.HealthCheck)
  Register(namespace, "SearchWhite", OciMonitoring.SearchWhite)
  Register(namespace, "NotificationsWhite", OciMonitoring.NotificationsWhite)
  Register(namespace, "QueueWhite", OciMonitoring.QueueWhite)
  Register(namespace, "TelemetryWhite", OciMonitoring.TelemetryWhite)
  Register(namespace, "WorkflowWhite", OciMonitoring.WorkflowWhite)
  Register(namespace, "AlarmWhite", OciMonitoring.AlarmWhite)
  Register(namespace, "EventsWhite", OciMonitoring.EventsWhite)
  Register(namespace, "Notifications", OciMonitoring.Notifications)
  Register(namespace, "Queue", OciMonitoring.Queue)
  Register(namespace, "Search", OciMonitoring.Search)
  Register(namespace, "Email", OciMonitoring.Email)
  Register(namespace, "HealthCheckWhite", OciMonitoring.HealthCheckWhite)
  Register(namespace, "Events", OciMonitoring.Events)
  Register(namespace, "Telemetry", OciMonitoring.Telemetry)
  Register(namespace, "Alarm", OciMonitoring.Alarm)
  Register(namespace, "EmailWhite", OciMonitoring.EmailWhite)
}
