package azure

import "github.com/blushft/go-diagrams/node"

type generalContainer struct {
	path string
	opts []node.Option
}

var General = &generalContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/general",
}

func (c *generalContainer) Shareddashboard(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/shareddashboard.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Twousericon(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/twousericon.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Azurehome(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/azurehome.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Recent(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/recent.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Resourcegroups(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/resourcegroups.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Quickstartcenter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/quickstartcenter.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Reservations(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/reservations.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Servicehealth(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/servicehealth.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Supportrequests(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/supportrequests.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Tag(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/tag.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Developertools(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/developertools.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Information(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/information.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Managementgroups(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/managementgroups.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Userhealthicon(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/userhealthicon.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Userprivacy(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/userprivacy.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Support(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/support.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Userresource(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/userresource.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Whatsnew(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/whatsnew.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Helpsupport(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/helpsupport.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Marketplace(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/marketplace.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Resource(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/resource.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Templates(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/templates.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Usericon(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/usericon.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Allresources(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/allresources.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Subscriptions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/subscriptions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *generalContainer) Tags(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/general/tags.png")}, c.opts, opts)
	return node.New(nopts...)
}
