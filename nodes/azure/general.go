package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type generalContainer struct {
	path  string
	attrs []attr.Attribute
}

var General = &generalContainer{path: "assets/azure/general"}

func (c *generalContainer) Allresources(opts ...attr.Attribute) *node.Node {
	return node.New("allresources", attr.AssetImage("assets/azure/general/allresources.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Shareddashboard(opts ...attr.Attribute) *node.Node {
	return node.New("shareddashboard", attr.AssetImage("assets/azure/general/shareddashboard.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Support(opts ...attr.Attribute) *node.Node {
	return node.New("support", attr.AssetImage("assets/azure/general/support.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Twousericon(opts ...attr.Attribute) *node.Node {
	return node.New("twousericon", attr.AssetImage("assets/azure/general/twousericon.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Whatsnew(opts ...attr.Attribute) *node.Node {
	return node.New("whatsnew", attr.AssetImage("assets/azure/general/whatsnew.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Helpsupport(opts ...attr.Attribute) *node.Node {
	return node.New("helpsupport", attr.AssetImage("assets/azure/general/helpsupport.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Information(opts ...attr.Attribute) *node.Node {
	return node.New("information", attr.AssetImage("assets/azure/general/information.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Quickstartcenter(opts ...attr.Attribute) *node.Node {
	return node.New("quickstartcenter", attr.AssetImage("assets/azure/general/quickstartcenter.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Servicehealth(opts ...attr.Attribute) *node.Node {
	return node.New("servicehealth", attr.AssetImage("assets/azure/general/servicehealth.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Templates(opts ...attr.Attribute) *node.Node {
	return node.New("templates", attr.AssetImage("assets/azure/general/templates.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Userhealthicon(opts ...attr.Attribute) *node.Node {
	return node.New("userhealthicon", attr.AssetImage("assets/azure/general/userhealthicon.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Usericon(opts ...attr.Attribute) *node.Node {
	return node.New("usericon", attr.AssetImage("assets/azure/general/usericon.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Userprivacy(opts ...attr.Attribute) *node.Node {
	return node.New("userprivacy", attr.AssetImage("assets/azure/general/userprivacy.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Developertools(opts ...attr.Attribute) *node.Node {
	return node.New("developertools", attr.AssetImage("assets/azure/general/developertools.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Managementgroups(opts ...attr.Attribute) *node.Node {
	return node.New("managementgroups", attr.AssetImage("assets/azure/general/managementgroups.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Marketplace(opts ...attr.Attribute) *node.Node {
	return node.New("marketplace", attr.AssetImage("assets/azure/general/marketplace.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Recent(opts ...attr.Attribute) *node.Node {
	return node.New("recent", attr.AssetImage("assets/azure/general/recent.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Subscriptions(opts ...attr.Attribute) *node.Node {
	return node.New("subscriptions", attr.AssetImage("assets/azure/general/subscriptions.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Supportrequests(opts ...attr.Attribute) *node.Node {
	return node.New("supportrequests", attr.AssetImage("assets/azure/general/supportrequests.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Tag(opts ...attr.Attribute) *node.Node {
	return node.New("tag", attr.AssetImage("assets/azure/general/tag.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Tags(opts ...attr.Attribute) *node.Node {
	return node.New("tags", attr.AssetImage("assets/azure/general/tags.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Azurehome(opts ...attr.Attribute) *node.Node {
	return node.New("azurehome", attr.AssetImage("assets/azure/general/azurehome.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Reservations(opts ...attr.Attribute) *node.Node {
	return node.New("reservations", attr.AssetImage("assets/azure/general/reservations.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Resource(opts ...attr.Attribute) *node.Node {
	return node.New("resource", attr.AssetImage("assets/azure/general/resource.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Resourcegroups(opts ...attr.Attribute) *node.Node {
	return node.New("resourcegroups", attr.AssetImage("assets/azure/general/resourcegroups.png"), attr.Shape(attr.None))
}

func (c *generalContainer) Userresource(opts ...attr.Attribute) *node.Node {
	return node.New("userresource", attr.AssetImage("assets/azure/general/userresource.png"), attr.Shape(attr.None))
}
