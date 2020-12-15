package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mobileContainer struct {
	path  string
	attrs []attr.Attribute
}

var Mobile = &mobileContainer{path: "assets/aws/mobile"}

func (c *mobileContainer) ApiGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/mobile/api-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-gateway", opts...)
}

func (c *mobileContainer) Appsync(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/mobile/appsync.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("appsync", opts...)
}

func (c *mobileContainer) DeviceFarm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/mobile/device-farm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("device-farm", opts...)
}

func (c *mobileContainer) Pinpoint(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/mobile/pinpoint.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("pinpoint", opts...)
}

func (c *mobileContainer) Amplify(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/mobile/amplify.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("amplify", opts...)
}

func (c *mobileContainer) ApiGatewayEndpoint(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/mobile/api-gateway-endpoint.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-gateway-endpoint", opts...)
}

func init() {
  const namespace= "aws.mobile"
  Register(namespace, "ApiGateway", Mobile.ApiGateway)
  Register(namespace, "Appsync", Mobile.Appsync)
  Register(namespace, "DeviceFarm", Mobile.DeviceFarm)
  Register(namespace, "Pinpoint", Mobile.Pinpoint)
  Register(namespace, "Amplify", Mobile.Amplify)
  Register(namespace, "ApiGatewayEndpoint", Mobile.ApiGatewayEndpoint)
}
