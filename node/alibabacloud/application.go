package alibabacloud

import "github.com/blushft/go-diagrams/node"

type applicationContainer struct {
	path string
	opts []node.Option
}

var Application = &applicationContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/application",
}

func (c *applicationContainer) ApiGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/api-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) BlockchainAsAService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/blockchain-as-a-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) MessageNotificationService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/message-notification-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) NodeJsPerformancePlatform(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/node-js-performance-platform.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) PerformanceTestingService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/performance-testing-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) CloudCallCenter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/cloud-call-center.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) CodePipeline(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/code-pipeline.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) DirectMail(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/direct-mail.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) LogService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/log-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) OpenSearch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/open-search.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) SmartConversationAnalysis(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/smart-conversation-analysis.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) RdCloud(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/rd-cloud.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) Yida(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/yida.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *applicationContainer) BeeBot(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/application/bee-bot.png")}, c.opts, opts)
	return node.New(nopts...)
}
