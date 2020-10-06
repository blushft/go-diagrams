package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type identityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Identity = &identityContainer{path: "assets/azure/identity"}

func (c *identityContainer) ManagedIdentities(opts ...attr.Attribute) *node.Node {
	return node.New("managed-identities", attr.AssetImage("assets/azure/identity/managed-identities.png"), attr.Shape(attr.None))
}

func (c *identityContainer) AccessReview(opts ...attr.Attribute) *node.Node {
	return node.New("access-review", attr.AssetImage("assets/azure/identity/access-review.png"), attr.Shape(attr.None))
}

func (c *identityContainer) ActiveDirectory(opts ...attr.Attribute) *node.Node {
	return node.New("active-directory", attr.AssetImage("assets/azure/identity/active-directory.png"), attr.Shape(attr.None))
}

func (c *identityContainer) AppRegistrations(opts ...attr.Attribute) *node.Node {
	return node.New("app-registrations", attr.AssetImage("assets/azure/identity/app-registrations.png"), attr.Shape(attr.None))
}

func (c *identityContainer) AdIdentityProtection(opts ...attr.Attribute) *node.Node {
	return node.New("ad-identity-protection", attr.AssetImage("assets/azure/identity/ad-identity-protection.png"), attr.Shape(attr.None))
}

func (c *identityContainer) AdPrivilegedIdentityManagement(opts ...attr.Attribute) *node.Node {
	return node.New("ad-privileged-identity-management", attr.AssetImage("assets/azure/identity/ad-privileged-identity-management.png"), attr.Shape(attr.None))
}

func (c *identityContainer) ConditionalAccess(opts ...attr.Attribute) *node.Node {
	return node.New("conditional-access", attr.AssetImage("assets/azure/identity/conditional-access.png"), attr.Shape(attr.None))
}

func (c *identityContainer) EnterpriseApplications(opts ...attr.Attribute) *node.Node {
	return node.New("enterprise-applications", attr.AssetImage("assets/azure/identity/enterprise-applications.png"), attr.Shape(attr.None))
}

func (c *identityContainer) IdentityGovernance(opts ...attr.Attribute) *node.Node {
	return node.New("identity-governance", attr.AssetImage("assets/azure/identity/identity-governance.png"), attr.Shape(attr.None))
}

func (c *identityContainer) ActiveDirectoryConnectHealth(opts ...attr.Attribute) *node.Node {
	return node.New("active-directory-connect-health", attr.AssetImage("assets/azure/identity/active-directory-connect-health.png"), attr.Shape(attr.None))
}

func (c *identityContainer) AdB2C(opts ...attr.Attribute) *node.Node {
	return node.New("ad-b2c", attr.AssetImage("assets/azure/identity/ad-b2c.png"), attr.Shape(attr.None))
}

func (c *identityContainer) AdDomainServices(opts ...attr.Attribute) *node.Node {
	return node.New("ad-domain-services", attr.AssetImage("assets/azure/identity/ad-domain-services.png"), attr.Shape(attr.None))
}

func (c *identityContainer) InformationProtection(opts ...attr.Attribute) *node.Node {
	return node.New("information-protection", attr.AssetImage("assets/azure/identity/information-protection.png"), attr.Shape(attr.None))
}
