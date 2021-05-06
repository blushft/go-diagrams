package alibabacloud

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type applicationContainer struct {
	path  string
	attrs []attr.Attribute
}

var Application = &applicationContainer{path: "assets/alibabacloud/application"}

func (c *applicationContainer) BeeBot(opts ...attr.Attribute) *node.Node {
	return node.New("bee-bot", attr.AssetImage("assets/alibabacloud/application/bee-bot.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) SmartConversationAnalysis(opts ...attr.Attribute) *node.Node {
	return node.New("smart-conversation-analysis", attr.AssetImage("assets/alibabacloud/application/smart-conversation-analysis.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) CodePipeline(opts ...attr.Attribute) *node.Node {
	return node.New("code-pipeline", attr.AssetImage("assets/alibabacloud/application/code-pipeline.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) DirectMail(opts ...attr.Attribute) *node.Node {
	return node.New("direct-mail", attr.AssetImage("assets/alibabacloud/application/direct-mail.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) OpenSearch(opts ...attr.Attribute) *node.Node {
	return node.New("open-search", attr.AssetImage("assets/alibabacloud/application/open-search.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) PerformanceTestingService(opts ...attr.Attribute) *node.Node {
	return node.New("performance-testing-service", attr.AssetImage("assets/alibabacloud/application/performance-testing-service.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) Yida(opts ...attr.Attribute) *node.Node {
	return node.New("yida", attr.AssetImage("assets/alibabacloud/application/yida.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) ApiGateway(opts ...attr.Attribute) *node.Node {
	return node.New("api-gateway", attr.AssetImage("assets/alibabacloud/application/api-gateway.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) BlockchainAsAService(opts ...attr.Attribute) *node.Node {
	return node.New("blockchain-as-a-service", attr.AssetImage("assets/alibabacloud/application/blockchain-as-a-service.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) CloudCallCenter(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-call-center", attr.AssetImage("assets/alibabacloud/application/cloud-call-center.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) LogService(opts ...attr.Attribute) *node.Node {
	return node.New("log-service", attr.AssetImage("assets/alibabacloud/application/log-service.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) MessageNotificationService(opts ...attr.Attribute) *node.Node {
	return node.New("message-notification-service", attr.AssetImage("assets/alibabacloud/application/message-notification-service.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) NodeJsPerformancePlatform(opts ...attr.Attribute) *node.Node {
	return node.New("node-js-performance-platform", attr.AssetImage("assets/alibabacloud/application/node-js-performance-platform.png"), attr.Shape(attr.None))
}

func (c *applicationContainer) RdCloud(opts ...attr.Attribute) *node.Node {
	return node.New("rd-cloud", attr.AssetImage("assets/alibabacloud/application/rd-cloud.png"), attr.Shape(attr.None))
}
