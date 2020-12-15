package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type applicationContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabaApplication = &applicationContainer{path: "assets/alibabacloud/application"}

func (c *applicationContainer) CodePipeline(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/code-pipeline.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("code-pipeline", opts...)
}

func (c *applicationContainer) NodeJsPerformancePlatform(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/node-js-performance-platform.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("node-js-performance-platform", opts...)
}

func (c *applicationContainer) BeeBot(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/bee-bot.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bee-bot", opts...)
}

func (c *applicationContainer) OpenSearch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/open-search.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("open-search", opts...)
}

func (c *applicationContainer) RdCloud(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/rd-cloud.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rd-cloud", opts...)
}

func (c *applicationContainer) SmartConversationAnalysis(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/smart-conversation-analysis.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("smart-conversation-analysis", opts...)
}

func (c *applicationContainer) DirectMail(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/direct-mail.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("direct-mail", opts...)
}

func (c *applicationContainer) LogService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/log-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("log-service", opts...)
}

func (c *applicationContainer) BlockchainAsAService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/blockchain-as-a-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("blockchain-as-a-service", opts...)
}

func (c *applicationContainer) CloudCallCenter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/cloud-call-center.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-call-center", opts...)
}

func (c *applicationContainer) MessageNotificationService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/message-notification-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("message-notification-service", opts...)
}

func (c *applicationContainer) PerformanceTestingService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/performance-testing-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("performance-testing-service", opts...)
}

func (c *applicationContainer) Yida(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/yida.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("yida", opts...)
}

func (c *applicationContainer) ApiGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/application/api-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-gateway", opts...)
}

func init() {
	const namespace = "alibabacloud.application"
	Register(namespace, "CodePipeline", AlibabaApplication.CodePipeline)
	Register(namespace, "NodeJsPerformancePlatform", AlibabaApplication.NodeJsPerformancePlatform)
	Register(namespace, "BeeBot", AlibabaApplication.BeeBot)
	Register(namespace, "OpenSearch", AlibabaApplication.OpenSearch)
	Register(namespace, "RdCloud", AlibabaApplication.RdCloud)
	Register(namespace, "SmartConversationAnalysis", AlibabaApplication.SmartConversationAnalysis)
	Register(namespace, "DirectMail", AlibabaApplication.DirectMail)
	Register(namespace, "LogService", AlibabaApplication.LogService)
	Register(namespace, "BlockchainAsAService", AlibabaApplication.BlockchainAsAService)
	Register(namespace, "CloudCallCenter", AlibabaApplication.CloudCallCenter)
	Register(namespace, "MessageNotificationService", AlibabaApplication.MessageNotificationService)
	Register(namespace, "PerformanceTestingService", AlibabaApplication.PerformanceTestingService)
	Register(namespace, "Yida", AlibabaApplication.Yida)
	Register(namespace, "ApiGateway", AlibabaApplication.ApiGateway)
}