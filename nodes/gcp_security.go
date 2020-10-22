package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type gcpSecurityContainer struct {
	path  string
	attrs []attr.Attribute
}

var GcpSecurity = &gcpSecurityContainer{path: "assets/gcp/security"}

func (c *gcpSecurityContainer) KeyManagementService(opts ...attr.Attribute) *node.Node {
	return node.New("key-management-service", attr.AssetImage("assets/gcp/security/key-management-service.png"), attr.Shape(attr.None))
}

func (c *gcpSecurityContainer) ResourceManager(opts ...attr.Attribute) *node.Node {
	return node.New("resource-manager", attr.AssetImage("assets/gcp/security/resource-manager.png"), attr.Shape(attr.None))
}

func (c *gcpSecurityContainer) SecurityCommandCenter(opts ...attr.Attribute) *node.Node {
	return node.New("security-command-center", attr.AssetImage("assets/gcp/security/security-command-center.png"), attr.Shape(attr.None))
}

func (c *gcpSecurityContainer) SecurityScanner(opts ...attr.Attribute) *node.Node {
	return node.New("security-scanner", attr.AssetImage("assets/gcp/security/security-scanner.png"), attr.Shape(attr.None))
}

func (c *gcpSecurityContainer) Iam(opts ...attr.Attribute) *node.Node {
	return node.New("iam", attr.AssetImage("assets/gcp/security/iam.png"), attr.Shape(attr.None))
}

func (c *gcpSecurityContainer) Iap(opts ...attr.Attribute) *node.Node {
	return node.New("iap", attr.AssetImage("assets/gcp/security/iap.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "gcp.security"
	Register(namespace, "KeyManagementService", GcpSecurity.KeyManagementService)
	Register(namespace, "ResourceManager", GcpSecurity.ResourceManager)
	Register(namespace, "SecurityCommandCenter", GcpSecurity.SecurityCommandCenter)
	Register(namespace, "SecurityScanner", GcpSecurity.SecurityScanner)
	Register(namespace, "Iam", GcpSecurity.Iam)
	Register(namespace, "Iap", GcpSecurity.Iap)
}