package oci

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type devopsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Devops = &devopsContainer{path: "assets/oci/devops"}

func (c *devopsContainer) ApiGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("api-gateway-white", attr.AssetImage("assets/oci/devops/api-gateway-white.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) ApiGateway(opts ...attr.Attribute) *node.Node {
	return node.New("api-gateway", attr.AssetImage("assets/oci/devops/api-gateway.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) ApiServiceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("api-service-white", attr.AssetImage("assets/oci/devops/api-service-white.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) ApiService(opts ...attr.Attribute) *node.Node {
	return node.New("api-service", attr.AssetImage("assets/oci/devops/api-service.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) ResourceMgmtWhite(opts ...attr.Attribute) *node.Node {
	return node.New("resource-mgmt-white", attr.AssetImage("assets/oci/devops/resource-mgmt-white.png"), attr.Shape(attr.None))
}

func (c *devopsContainer) ResourceMgmt(opts ...attr.Attribute) *node.Node {
	return node.New("resource-mgmt", attr.AssetImage("assets/oci/devops/resource-mgmt.png"), attr.Shape(attr.None))
}
