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
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/workflow.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("workflow", opts...)
}

func (c *ociMonitoringContainer) HealthCheck(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/health-check.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("health-check", opts...)
}

func (c *ociMonitoringContainer) SearchWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/search-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("search-white", opts...)
}

func (c *ociMonitoringContainer) NotificationsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/notifications-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("notifications-white", opts...)
}

func (c *ociMonitoringContainer) QueueWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/queue-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("queue-white", opts...)
}

func (c *ociMonitoringContainer) TelemetryWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/telemetry-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("telemetry-white", opts...)
}

func (c *ociMonitoringContainer) WorkflowWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/workflow-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("workflow-white", opts...)
}

func (c *ociMonitoringContainer) AlarmWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/alarm-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("alarm-white", opts...)
}

func (c *ociMonitoringContainer) EventsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/events-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("events-white", opts...)
}

func (c *ociMonitoringContainer) Notifications(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/notifications.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("notifications", opts...)
}

func (c *ociMonitoringContainer) Queue(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/queue.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("queue", opts...)
}

func (c *ociMonitoringContainer) Search(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/search.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("search", opts...)
}

func (c *ociMonitoringContainer) Email(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/email.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("email", opts...)
}

func (c *ociMonitoringContainer) HealthCheckWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/health-check-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("health-check-white", opts...)
}

func (c *ociMonitoringContainer) Events(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/events.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("events", opts...)
}

func (c *ociMonitoringContainer) Telemetry(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/telemetry.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("telemetry", opts...)
}

func (c *ociMonitoringContainer) Alarm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/alarm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("alarm", opts...)
}

func (c *ociMonitoringContainer) EmailWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/monitoring/email-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("email-white", opts...)
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
