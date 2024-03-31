package aws

import "github.com/blushft/go-diagrams/diagram"

type mediaContainer struct {
	path string
	opts []diagram.NodeOption
}

var Media = &mediaContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/media",
}

func (c *mediaContainer) ElementalMediapackage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-mediapackage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElementalMediastore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-mediastore.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElasticTranscoder(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elastic-transcoder.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElementalConductor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-conductor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElementalDelta(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-delta.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElementalLive(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-live.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElementalMediaconnect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-mediaconnect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElementalMediaconvert(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-mediaconvert.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElementalMedialive(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-medialive.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElementalMediatailor(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-mediatailor.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mediaContainer) ElementalServer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/media/elemental-server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
