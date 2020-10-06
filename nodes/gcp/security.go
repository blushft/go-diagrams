package gcp

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type securityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Security = &securityContainer{path: "assets/gcp/security"}

func (c *securityContainer) KeyManagementService(opts ...attr.Attribute) *node.Node {
	return node.New("key-management-service", attr.AssetImage("assets/gcp/security/key-management-service.png"), attr.Shape(attr.None))
}

func (c *securityContainer) ResourceManager(opts ...attr.Attribute) *node.Node {
	return node.New("resource-manager", attr.AssetImage("assets/gcp/security/resource-manager.png"), attr.Shape(attr.None))
}

func (c *securityContainer) SecurityCommandCenter(opts ...attr.Attribute) *node.Node {
	return node.New("security-command-center", attr.AssetImage("assets/gcp/security/security-command-center.png"), attr.Shape(attr.None))
}

func (c *securityContainer) SecurityScanner(opts ...attr.Attribute) *node.Node {
	return node.New("security-scanner", attr.AssetImage("assets/gcp/security/security-scanner.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Iam(opts ...attr.Attribute) *node.Node {
	return node.New("iam", attr.AssetImage("assets/gcp/security/iam.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Iap(opts ...attr.Attribute) *node.Node {
	return node.New("iap", attr.AssetImage("assets/gcp/security/iap.png"), attr.Shape(attr.None))
}
