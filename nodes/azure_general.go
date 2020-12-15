package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azuregeneralContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureGeneral = &azuregeneralContainer{path: "assets/azure/general"}

func (c *azuregeneralContainer) Allresources(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/allresources.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("allresources", opts...)
}

func (c *azuregeneralContainer) Shareddashboard(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/shareddashboard.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("shareddashboard", opts...)
}

func (c *azuregeneralContainer) Support(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/support.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("support", opts...)
}

func (c *azuregeneralContainer) Twousericon(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/twousericon.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("twousericon", opts...)
}

func (c *azuregeneralContainer) Whatsnew(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/whatsnew.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("whatsnew", opts...)
}

func (c *azuregeneralContainer) Helpsupport(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/helpsupport.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("helpsupport", opts...)
}

func (c *azuregeneralContainer) Information(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/information.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("information", opts...)
}

func (c *azuregeneralContainer) Quickstartcenter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/quickstartcenter.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("quickstartcenter", opts...)
}

func (c *azuregeneralContainer) Servicehealth(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/servicehealth.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("servicehealth", opts...)
}

func (c *azuregeneralContainer) Templates(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/templates.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("templates", opts...)
}

func (c *azuregeneralContainer) Userhealthicon(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/userhealthicon.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("userhealthicon", opts...)
}

func (c *azuregeneralContainer) Usericon(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/usericon.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("usericon", opts...)
}

func (c *azuregeneralContainer) Userprivacy(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/userprivacy.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("userprivacy", opts...)
}

func (c *azuregeneralContainer) Developertools(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/developertools.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("developertools", opts...)
}

func (c *azuregeneralContainer) Managementgroups(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/managementgroups.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("managementgroups", opts...)
}

func (c *azuregeneralContainer) Marketplace(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/marketplace.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("marketplace", opts...)
}

func (c *azuregeneralContainer) Recent(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/recent.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("recent", opts...)
}

func (c *azuregeneralContainer) Subscriptions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/subscriptions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("subscriptions", opts...)
}

func (c *azuregeneralContainer) Supportrequests(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/supportrequests.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("supportrequests", opts...)
}

func (c *azuregeneralContainer) Tag(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/tag.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tag", opts...)
}

func (c *azuregeneralContainer) Tags(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/tags.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tags", opts...)
}

func (c *azuregeneralContainer) Azurehome(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/azurehome.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("azurehome", opts...)
}

func (c *azuregeneralContainer) Reservations(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/reservations.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("reservations", opts...)
}

func (c *azuregeneralContainer) Resource(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/resource.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("resource", opts...)
}

func (c *azuregeneralContainer) Resourcegroups(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/resourcegroups.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("resourcegroups", opts...)
}

func (c *azuregeneralContainer) Userresource(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/general/userresource.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("userresource", opts...)
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
}
