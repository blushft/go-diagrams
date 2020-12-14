package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azureidentityContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureIdentity = &saasIdentityContainer{path: "assets/azure/identity"}

func (c *azureidentityContainer) ManagedIdentities(opts ...attr.Attribute) *node.Node {
	return node.New("managed-identities", attr.AssetImage("assets/azure/identity/managed-identities.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) AccessReview(opts ...attr.Attribute) *node.Node {
	return node.New("access-review", attr.AssetImage("assets/azure/identity/access-review.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) ActiveDirectory(opts ...attr.Attribute) *node.Node {
	return node.New("active-directory", attr.AssetImage("assets/azure/identity/active-directory.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) AppRegistrations(opts ...attr.Attribute) *node.Node {
	return node.New("app-registrations", attr.AssetImage("assets/azure/identity/app-registrations.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) AdIdentityProtection(opts ...attr.Attribute) *node.Node {
	return node.New("ad-identity-protection", attr.AssetImage("assets/azure/identity/ad-identity-protection.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) AdPrivilegedIdentityManagement(opts ...attr.Attribute) *node.Node {
	return node.New("ad-privileged-identity-management", attr.AssetImage("assets/azure/identity/ad-privileged-identity-management.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) ConditionalAccess(opts ...attr.Attribute) *node.Node {
	return node.New("conditional-access", attr.AssetImage("assets/azure/identity/conditional-access.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) EnterpriseApplications(opts ...attr.Attribute) *node.Node {
	return node.New("enterprise-applications", attr.AssetImage("assets/azure/identity/enterprise-applications.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) IdentityGovernance(opts ...attr.Attribute) *node.Node {
	return node.New("identity-governance", attr.AssetImage("assets/azure/identity/identity-governance.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) ActiveDirectoryConnectHealth(opts ...attr.Attribute) *node.Node {
	return node.New("active-directory-connect-health", attr.AssetImage("assets/azure/identity/active-directory-connect-health.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) AdB2C(opts ...attr.Attribute) *node.Node {
	return node.New("ad-b2c", attr.AssetImage("assets/azure/identity/ad-b2c.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) AdDomainServices(opts ...attr.Attribute) *node.Node {
	return node.New("ad-domain-services", attr.AssetImage("assets/azure/identity/ad-domain-services.png"), attr.Shape(attr.None))
}

func (c *azureidentityContainer) InformationProtection(opts ...attr.Attribute) *node.Node {
	return node.New("information-protection", attr.AssetImage("assets/azure/identity/information-protection.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.identity"
  Register(namespace, "ManagedIdentities", azureIdentity.ManagedIdentities)
  Register(namespace, "AccessReview", azureIdentity.AccessReview)
  Register(namespace, "ActiveDirectory", azureIdentity.ActiveDirectory)
  Register(namespace, "AppRegistrations", azureIdentity.AppRegistrations)
  Register(namespace, "AdIdentityProtection", azureIdentity.AdIdentityProtection)
  Register(namespace, "AdPrivilegedIdentityManagement", azureIdentity.AdPrivilegedIdentityManagement)
  Register(namespace, "ConditionalAccess", azureIdentity.ConditionalAccess)
  Register(namespace, "EnterpriseApplications", azureIdentity.EnterpriseApplications)
  Register(namespace, "IdentityGovernance", azureIdentity.IdentityGovernance)
  Register(namespace, "ActiveDirectoryConnectHealth", azureIdentity.ActiveDirectoryConnectHealth)
  Register(namespace, "AdB2C", azureIdentity.AdB2C)
  Register(namespace, "AdDomainServices", azureIdentity.AdDomainServices)
  Register(namespace, "InformationProtection", azureIdentity.InformationProtection)
  Register(namespace, "init", azureIdentity.init)
}
