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
	opts = append(opts, attr.AssetImage("assets/oci/devops/api-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-service", opts...)
}

func (c *ociDevopsContainer) ResourceMgmtWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/devops/resource-mgmt-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("resource-mgmt-white", opts...)
}

func (c *ociDevopsContainer) ResourceMgmt(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/devops/resource-mgmt.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("resource-mgmt", opts...)
}

func (c *ociDevopsContainer) ApiGatewayWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/devops/api-gateway-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-gateway-white", opts...)
}

func (c *ociDevopsContainer) ApiGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/devops/api-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-gateway", opts...)
}

func (c *ociDevopsContainer) ApiServiceWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/devops/api-service-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("api-service-white", opts...)
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
