package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type mediaContainer struct {
	path  string
	attrs []attr.Attribute
}

var Media = &mediaContainer{path: "assets/aws/media"}

func (c *mediaContainer) ElasticTranscoder(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-transcoder", attr.AssetImage("assets/aws/media/elastic-transcoder.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalDelta(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-delta", attr.AssetImage("assets/aws/media/elemental-delta.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalLive(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-live", attr.AssetImage("assets/aws/media/elemental-live.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalMediaconnect(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediaconnect", attr.AssetImage("assets/aws/media/elemental-mediaconnect.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalMedialive(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-medialive", attr.AssetImage("assets/aws/media/elemental-medialive.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalMediastore(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediastore", attr.AssetImage("assets/aws/media/elemental-mediastore.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalConductor(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-conductor", attr.AssetImage("assets/aws/media/elemental-conductor.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalMediaconvert(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediaconvert", attr.AssetImage("assets/aws/media/elemental-mediaconvert.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalMediapackage(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediapackage", attr.AssetImage("assets/aws/media/elemental-mediapackage.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalMediatailor(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediatailor", attr.AssetImage("assets/aws/media/elemental-mediatailor.png"), attr.Shape(attr.None))
}

func (c *mediaContainer) ElementalServer(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-server", attr.AssetImage("assets/aws/media/elemental-server.png"), attr.Shape(attr.None))
}
