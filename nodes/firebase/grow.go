package firebase

import "github.com/blushft/go-diagrams/diagram"

type growContainer struct {
	path string
	opts []diagram.NodeOption
}

var Grow = &growContainer{
	opts: diagram.OptionSet{diagram.Provider("firebase"), diagram.NodeShape("none")},
	path: "assets/firebase/grow",
}

func (c *growContainer) Invites(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/grow/invites.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *growContainer) Messaging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/grow/messaging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *growContainer) Predictions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/grow/predictions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *growContainer) RemoteConfig(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/grow/remote-config.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *growContainer) AbTesting(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/grow/ab-testing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *growContainer) AppIndexing(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/grow/app-indexing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *growContainer) DynamicLinks(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/grow/dynamic-links.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *growContainer) InAppMessaging(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/firebase/grow/in-app-messaging.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
