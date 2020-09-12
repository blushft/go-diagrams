package aws

import "github.com/blushft/go-diagrams/node"

type mediaContainer struct {
	path string
	opts []node.Option
}

var Media = &mediaContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/media",
}

func (c *mediaContainer) ElementalMediaconnect(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-mediaconnect.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElementalMediaconvert(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-mediaconvert.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElementalMediapackage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-mediapackage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElementalMediatailor(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-mediatailor.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElementalServer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-server.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElasticTranscoder(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elastic-transcoder.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElementalConductor(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-conductor.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElementalDelta(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-delta.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElementalLive(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-live.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElementalMedialive(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-medialive.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *mediaContainer) ElementalMediastore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/media/elemental-mediastore.png")}, c.opts, opts)
	return node.New(nopts...)
}
