package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type integrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Integration = &integrationContainer{path: "assets/aws/integration"}

func (c *integrationContainer) ApplicationIntegration(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/integration/application-integration.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("application-integration", opts...)
}

func (c *integrationContainer) Appsync(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/integration/appsync.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("appsync", opts...)
}

func (c *integrationContainer) ConsoleMobileApplication(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/integration/console-mobile-application.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("console-mobile-application", opts...)
}

func (c *integrationContainer) Eventbridge(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/integration/eventbridge.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("eventbridge", opts...)
}

func (c *integrationContainer) Mq(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/integration/mq.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mq", opts...)
}

func (c *integrationContainer) SimpleNotificationServiceSns(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/integration/simple-notification-service-sns.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("simple-notification-service-sns", opts...)
}

func (c *integrationContainer) SimpleQueueServiceSqs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/integration/simple-queue-service-sqs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("simple-queue-service-sqs", opts...)
}

func (c *integrationContainer) StepFunctions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/integration/step-functions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("step-functions", opts...)
}

func init() {
  const namespace = "aws.integration"
  Register(namespace, "ApplicationIntegration", Integration.ApplicationIntegration)
  Register(namespace, "Appsync", Integration.Appsync)
  Register(namespace, "ConsoleMobileApplication", Integration.ConsoleMobileApplication)
  Register(namespace, "Eventbridge", Integration.Eventbridge)
  Register(namespace, "Mq", Integration.Mq)
  Register(namespace, "SimpleNotificationServiceSns", Integration.SimpleNotificationServiceSns)
  Register(namespace, "SimpleQueueServiceSqs", Integration.SimpleQueueServiceSqs)
  Register(namespace, "StepFunctions", Integration.StepFunctions)
}
