package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuregeneralContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureGeneral = &generalContainer{path: "assets/azure/general"}

func (c *azuregeneralContainer) Allresources(opts ...attr.Attribute) *node.Node {
	return node.New("allresources", attr.AssetImage("assets/azure/general/allresources.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Shareddashboard(opts ...attr.Attribute) *node.Node {
	return node.New("shareddashboard", attr.AssetImage("assets/azure/general/shareddashboard.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Support(opts ...attr.Attribute) *node.Node {
	return node.New("support", attr.AssetImage("assets/azure/general/support.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Twousericon(opts ...attr.Attribute) *node.Node {
	return node.New("twousericon", attr.AssetImage("assets/azure/general/twousericon.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Whatsnew(opts ...attr.Attribute) *node.Node {
	return node.New("whatsnew", attr.AssetImage("assets/azure/general/whatsnew.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Helpsupport(opts ...attr.Attribute) *node.Node {
	return node.New("helpsupport", attr.AssetImage("assets/azure/general/helpsupport.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Information(opts ...attr.Attribute) *node.Node {
	return node.New("information", attr.AssetImage("assets/azure/general/information.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Quickstartcenter(opts ...attr.Attribute) *node.Node {
	return node.New("quickstartcenter", attr.AssetImage("assets/azure/general/quickstartcenter.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Servicehealth(opts ...attr.Attribute) *node.Node {
	return node.New("servicehealth", attr.AssetImage("assets/azure/general/servicehealth.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Templates(opts ...attr.Attribute) *node.Node {
	return node.New("templates", attr.AssetImage("assets/azure/general/templates.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Userhealthicon(opts ...attr.Attribute) *node.Node {
	return node.New("userhealthicon", attr.AssetImage("assets/azure/general/userhealthicon.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Usericon(opts ...attr.Attribute) *node.Node {
	return node.New("usericon", attr.AssetImage("assets/azure/general/usericon.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Userprivacy(opts ...attr.Attribute) *node.Node {
	return node.New("userprivacy", attr.AssetImage("assets/azure/general/userprivacy.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Developertools(opts ...attr.Attribute) *node.Node {
	return node.New("developertools", attr.AssetImage("assets/azure/general/developertools.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Managementgroups(opts ...attr.Attribute) *node.Node {
	return node.New("managementgroups", attr.AssetImage("assets/azure/general/managementgroups.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Marketplace(opts ...attr.Attribute) *node.Node {
	return node.New("marketplace", attr.AssetImage("assets/azure/general/marketplace.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Recent(opts ...attr.Attribute) *node.Node {
	return node.New("recent", attr.AssetImage("assets/azure/general/recent.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Subscriptions(opts ...attr.Attribute) *node.Node {
	return node.New("subscriptions", attr.AssetImage("assets/azure/general/subscriptions.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Supportrequests(opts ...attr.Attribute) *node.Node {
	return node.New("supportrequests", attr.AssetImage("assets/azure/general/supportrequests.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Tag(opts ...attr.Attribute) *node.Node {
	return node.New("tag", attr.AssetImage("assets/azure/general/tag.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Tags(opts ...attr.Attribute) *node.Node {
	return node.New("tags", attr.AssetImage("assets/azure/general/tags.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Azurehome(opts ...attr.Attribute) *node.Node {
	return node.New("azurehome", attr.AssetImage("assets/azure/general/azurehome.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Reservations(opts ...attr.Attribute) *node.Node {
	return node.New("reservations", attr.AssetImage("assets/azure/general/reservations.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Resource(opts ...attr.Attribute) *node.Node {
	return node.New("resource", attr.AssetImage("assets/azure/general/resource.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Resourcegroups(opts ...attr.Attribute) *node.Node {
	return node.New("resourcegroups", attr.AssetImage("assets/azure/general/resourcegroups.png"), attr.Shape(attr.None))
}

func (c *azuregeneralContainer) Userresource(opts ...attr.Attribute) *node.Node {
	return node.New("userresource", attr.AssetImage("assets/azure/general/userresource.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.general"
  Register(namespace, "Allresources", azureGeneral.Allresources)
  Register(namespace, "Shareddashboard", azureGeneral.Shareddashboard)
  Register(namespace, "Support", azureGeneral.Support)
  Register(namespace, "Twousericon", azureGeneral.Twousericon)
  Register(namespace, "Whatsnew", azureGeneral.Whatsnew)
  Register(namespace, "Helpsupport", azureGeneral.Helpsupport)
  Register(namespace, "Information", azureGeneral.Information)
  Register(namespace, "Quickstartcenter", azureGeneral.Quickstartcenter)
  Register(namespace, "Servicehealth", azureGeneral.Servicehealth)
  Register(namespace, "Templates", azureGeneral.Templates)
  Register(namespace, "Userhealthicon", azureGeneral.Userhealthicon)
  Register(namespace, "Usericon", azureGeneral.Usericon)
  Register(namespace, "Userprivacy", azureGeneral.Userprivacy)
  Register(namespace, "Developertools", azureGeneral.Developertools)
  Register(namespace, "Managementgroups", azureGeneral.Managementgroups)
  Register(namespace, "Marketplace", azureGeneral.Marketplace)
  Register(namespace, "Recent", azureGeneral.Recent)
  Register(namespace, "Subscriptions", azureGeneral.Subscriptions)
  Register(namespace, "Supportrequests", azureGeneral.Supportrequests)
  Register(namespace, "Tag", azureGeneral.Tag)
  Register(namespace, "Tags", azureGeneral.Tags)
  Register(namespace, "Azurehome", azureGeneral.Azurehome)
  Register(namespace, "Reservations", azureGeneral.Reservations)
  Register(namespace, "Resource", azureGeneral.Resource)
  Register(namespace, "Resourcegroups", azureGeneral.Resourcegroups)
  Register(namespace, "Userresource", azureGeneral.Userresource)
  Register(namespace, "init", azureGeneral.init)
}
