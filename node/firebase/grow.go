package firebase

import "github.com/blushft/go-diagrams/node"

type growContainer struct {
	path string
	opts []node.Option
}

var Grow = &growContainer{
	opts: node.OptionSet{node.Provider("firebase"), node.Shape("none")},
	path: "assets/firebase/grow",
}

func (c *growContainer) Invites(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/grow/invites.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *growContainer) Messaging(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/grow/messaging.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *growContainer) Predictions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/grow/predictions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *growContainer) RemoteConfig(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/grow/remote-config.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *growContainer) AbTesting(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/grow/ab-testing.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *growContainer) AppIndexing(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/grow/app-indexing.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *growContainer) DynamicLinks(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/grow/dynamic-links.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *growContainer) InAppMessaging(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/firebase/grow/in-app-messaging.png")}, c.opts, opts)
	return node.New(nopts...)
}
