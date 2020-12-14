package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ociDevopsContainer struct {
	path  string
	attrs []attr.Attribute
}

var OciDevops = &ociDevopsContainer{path: "assets/oci/devops"}

func (c *ociDevopsContainer) ApiService(opts ...attr.Attribute) *node.Node {
	return node.New("api-service", attr.AssetImage("assets/oci/devops/api-service.png"), attr.Shape(attr.None))
}

func (c *ociDevopsContainer) ResourceMgmtWhite(opts ...attr.Attribute) *node.Node {
	return node.New("resource-mgmt-white", attr.AssetImage("assets/oci/devops/resource-mgmt-white.png"), attr.Shape(attr.None))
}

func (c *ociDevopsContainer) ResourceMgmt(opts ...attr.Attribute) *node.Node {
	return node.New("resource-mgmt", attr.AssetImage("assets/oci/devops/resource-mgmt.png"), attr.Shape(attr.None))
}

func (c *ociDevopsContainer) ApiGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("api-gateway-white", attr.AssetImage("assets/oci/devops/api-gateway-white.png"), attr.Shape(attr.None))
}

func (c *ociDevopsContainer) ApiGateway(opts ...attr.Attribute) *node.Node {
	return node.New("api-gateway", attr.AssetImage("assets/oci/devops/api-gateway.png"), attr.Shape(attr.None))
}

func (c *ociDevopsContainer) ApiServiceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("api-service-white", attr.AssetImage("assets/oci/devops/api-service-white.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "oci.devops"
  Register(namespace, "ApiService", OciDevops.ApiService)
  Register(namespace, "ResourceMgmtWhite", OciDevops.ResourceMgmtWhite)
  Register(namespace, "ResourceMgmt", OciDevops.ResourceMgmt)
  Register(namespace, "ApiGatewayWhite", OciDevops.ApiGatewayWhite)
  Register(namespace, "ApiGateway", OciDevops.ApiGateway)
  Register(namespace, "ApiServiceWhite", OciDevops.ApiServiceWhite)
}
