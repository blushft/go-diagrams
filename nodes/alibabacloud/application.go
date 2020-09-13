package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type applicationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Application = &applicationContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/application",
}

func (c *applicationContainer) Yida(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/yida.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) ApiGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/api-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) PerformanceTestingService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/performance-testing-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) CodePipeline(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/code-pipeline.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) DirectMail(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/direct-mail.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) MessageNotificationService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/message-notification-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) RdCloud(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/rd-cloud.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) BlockchainAsAService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/blockchain-as-a-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) CloudCallCenter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/cloud-call-center.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) NodeJsPerformancePlatform(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/node-js-performance-platform.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) OpenSearch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/open-search.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) SmartConversationAnalysis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/smart-conversation-analysis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) BeeBot(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/bee-bot.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *applicationContainer) LogService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/application/log-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
