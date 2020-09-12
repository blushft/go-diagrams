package aws

import "github.com/blushft/go-diagrams/node"

type integrationContainer struct {
	path string
	opts []node.Option
}

var Integration = &integrationContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/integration",
}

func (c *integrationContainer) Mq(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/integration/mq.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) SimpleNotificationServiceSns(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/integration/simple-notification-service-sns.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) SimpleQueueServiceSqs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/integration/simple-queue-service-sqs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) StepFunctions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/integration/step-functions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) ApplicationIntegration(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/integration/application-integration.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) Appsync(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/integration/appsync.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) ConsoleMobileApplication(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/integration/console-mobile-application.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *integrationContainer) Eventbridge(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/integration/eventbridge.png")}, c.opts, opts)
	return node.New(nopts...)
}
