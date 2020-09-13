package aws

import "github.com/blushft/go-diagrams/diagram"

type integrationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Integration = &integrationContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/integration",
}

func (c *integrationContainer) Appsync(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/appsync.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) ConsoleMobileApplication(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/console-mobile-application.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) Eventbridge(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/eventbridge.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) Mq(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/mq.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SimpleNotificationServiceSns(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/simple-notification-service-sns.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) SimpleQueueServiceSqs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/simple-queue-service-sqs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) StepFunctions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/step-functions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *integrationContainer) ApplicationIntegration(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/integration/application-integration.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
