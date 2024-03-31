package azure

import "github.com/blushft/go-diagrams/diagram"

type generalContainer struct {
	path string
	opts []diagram.NodeOption
}

var General = &generalContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/general",
}

func (c *generalContainer) Developertools(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/developertools.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Managementgroups(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/managementgroups.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Subscriptions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/subscriptions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Userhealthicon(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/userhealthicon.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Usericon(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/usericon.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Userresource(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/userresource.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Whatsnew(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/whatsnew.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Helpsupport(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/helpsupport.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Quickstartcenter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/quickstartcenter.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Reservations(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/reservations.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Support(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/support.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Supportrequests(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/supportrequests.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Tags(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/tags.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Templates(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/templates.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Allresources(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/allresources.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Azurehome(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/azurehome.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Recent(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/recent.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Resourcegroups(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/resourcegroups.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Servicehealth(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/servicehealth.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Twousericon(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/twousericon.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Userprivacy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/userprivacy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Information(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/information.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Marketplace(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/marketplace.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Resource(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/resource.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Shareddashboard(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/shareddashboard.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *generalContainer) Tag(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/general/tag.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
