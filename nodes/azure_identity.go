package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azureIdentityContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureidentityContainer = &azureIdentityContainer{path: "assets/azure/identity"}

func (c *azureIdentityContainer) ManagedIdentities(opts ...attr.Attribute) *node.Node {
	return node.New("managed-identities", attr.AssetImage("assets/azure/identity/managed-identities.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) AccessReview(opts ...attr.Attribute) *node.Node {
	return node.New("access-review", attr.AssetImage("assets/azure/identity/access-review.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) ActiveDirectory(opts ...attr.Attribute) *node.Node {
	return node.New("active-directory", attr.AssetImage("assets/azure/identity/active-directory.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) AppRegistrations(opts ...attr.Attribute) *node.Node {
	return node.New("app-registrations", attr.AssetImage("assets/azure/identity/app-registrations.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) AdIdentityProtection(opts ...attr.Attribute) *node.Node {
	return node.New("ad-identity-protection", attr.AssetImage("assets/azure/identity/ad-identity-protection.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) AdPrivilegedIdentityManagement(opts ...attr.Attribute) *node.Node {
	return node.New("ad-privileged-identity-management", attr.AssetImage("assets/azure/identity/ad-privileged-identity-management.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) ConditionalAccess(opts ...attr.Attribute) *node.Node {
	return node.New("conditional-access", attr.AssetImage("assets/azure/identity/conditional-access.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) EnterpriseApplications(opts ...attr.Attribute) *node.Node {
	return node.New("enterprise-applications", attr.AssetImage("assets/azure/identity/enterprise-applications.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) IdentityGovernance(opts ...attr.Attribute) *node.Node {
	return node.New("identity-governance", attr.AssetImage("assets/azure/identity/identity-governance.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) ActiveDirectoryConnectHealth(opts ...attr.Attribute) *node.Node {
	return node.New("active-directory-connect-health", attr.AssetImage("assets/azure/identity/active-directory-connect-health.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) AdB2C(opts ...attr.Attribute) *node.Node {
	return node.New("ad-b2c", attr.AssetImage("assets/azure/identity/ad-b2c.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) AdDomainServices(opts ...attr.Attribute) *node.Node {
	return node.New("ad-domain-services", attr.AssetImage("assets/azure/identity/ad-domain-services.png"), attr.Shape(attr.None))
}

func (c *azureIdentityContainer) InformationProtection(opts ...attr.Attribute) *node.Node {
	return node.New("information-protection", attr.AssetImage("assets/azure/identity/information-protection.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.identity"
  Register(namespace, "ManagedIdentities", azureidentityContainer.ManagedIdentities)
  Register(namespace, "AccessReview", azureidentityContainer.AccessReview)
  Register(namespace, "ActiveDirectory", azureidentityContainer.ActiveDirectory)
  Register(namespace, "AppRegistrations", azureidentityContainer.AppRegistrations)
  Register(namespace, "AdIdentityProtection", azureidentityContainer.AdIdentityProtection)
  Register(namespace, "AdPrivilegedIdentityManagement", azureidentityContainer.AdPrivilegedIdentityManagement)
  Register(namespace, "ConditionalAccess", azureidentityContainer.ConditionalAccess)
  Register(namespace, "EnterpriseApplications", azureidentityContainer.EnterpriseApplications)
  Register(namespace, "IdentityGovernance", azureidentityContainer.IdentityGovernance)
  Register(namespace, "ActiveDirectoryConnectHealth", azureidentityContainer.ActiveDirectoryConnectHealth)
  Register(namespace, "AdB2C", azureidentityContainer.AdB2C)
  Register(namespace, "AdDomainServices", azureidentityContainer.AdDomainServices)
  Register(namespace, "InformationProtection", azureidentityContainer.InformationProtection)
}
