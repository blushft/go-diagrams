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
	opts = append(opts, attr.AssetImage("assets/gcp/security/key-management-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("key-management-service", opts...)
}

func (c *gcpSecurityContainer) ResourceManager(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/security/resource-manager.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("resource-manager", opts...)
}

func (c *gcpSecurityContainer) SecurityCommandCenter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/security/security-command-center.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security-command-center", opts...)
}

func (c *gcpSecurityContainer) SecurityScanner(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/security/security-scanner.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security-scanner", opts...)
}

func (c *gcpSecurityContainer) Iam(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/security/iam.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iam", opts...)
}

func (c *gcpSecurityContainer) Iap(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/security/iap.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iap", opts...)
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