package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mobileContainer struct {
	path  string
	attrs []attr.Attribute
}

var Mobile = &mobileContainer{path: "assets/aws/mobile"}

func (c *mobileContainer) Pinpoint(opts ...attr.Attribute) *node.Node {
	return node.New("pinpoint", attr.AssetImage("assets/aws/mobile/pinpoint.png"), attr.Shape(attr.None))
}

func (c *mobileContainer) Amplify(opts ...attr.Attribute) *node.Node {
	return node.New("amplify", attr.AssetImage("assets/aws/mobile/amplify.png"), attr.Shape(attr.None))
}

func (c *mobileContainer) ApiGatewayEndpoint(opts ...attr.Attribute) *node.Node {
	return node.New("api-gateway-endpoint", attr.AssetImage("assets/aws/mobile/api-gateway-endpoint.png"), attr.Shape(attr.None))
}

func (c *mobileContainer) ApiGateway(opts ...attr.Attribute) *node.Node {
	return node.New("api-gateway", attr.AssetImage("assets/aws/mobile/api-gateway.png"), attr.Shape(attr.None))
}

func (c *mobileContainer) Appsync(opts ...attr.Attribute) *node.Node {
	return node.New("appsync", attr.AssetImage("assets/aws/mobile/appsync.png"), attr.Shape(attr.None))
}

func (c *mobileContainer) DeviceFarm(opts ...attr.Attribute) *node.Node {
	return node.New("device-farm", attr.AssetImage("assets/aws/mobile/device-farm.png"), attr.Shape(attr.None))
}
