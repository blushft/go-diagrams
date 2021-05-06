package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type integrationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Integration = &integrationContainer{path: "assets/aws/integration"}

func (c *integrationContainer) ConsoleMobileApplication(opts ...attr.Attribute) *node.Node {
	return node.New("console-mobile-application", attr.AssetImage("assets/aws/integration/console-mobile-application.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) Eventbridge(opts ...attr.Attribute) *node.Node {
	return node.New("eventbridge", attr.AssetImage("assets/aws/integration/eventbridge.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) Mq(opts ...attr.Attribute) *node.Node {
	return node.New("mq", attr.AssetImage("assets/aws/integration/mq.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) SimpleNotificationServiceSns(opts ...attr.Attribute) *node.Node {
	return node.New("simple-notification-service-sns", attr.AssetImage("assets/aws/integration/simple-notification-service-sns.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) SimpleQueueServiceSqs(opts ...attr.Attribute) *node.Node {
	return node.New("simple-queue-service-sqs", attr.AssetImage("assets/aws/integration/simple-queue-service-sqs.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) StepFunctions(opts ...attr.Attribute) *node.Node {
	return node.New("step-functions", attr.AssetImage("assets/aws/integration/step-functions.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) ApplicationIntegration(opts ...attr.Attribute) *node.Node {
	return node.New("application-integration", attr.AssetImage("assets/aws/integration/application-integration.png"), attr.Shape(attr.None))
}

func (c *integrationContainer) Appsync(opts ...attr.Attribute) *node.Node {
	return node.New("appsync", attr.AssetImage("assets/aws/integration/appsync.png"), attr.Shape(attr.None))
}
