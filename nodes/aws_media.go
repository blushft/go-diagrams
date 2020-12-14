package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsMediaContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSMedia = &awsMediaContainer{path: "assets/aws/media"}

func (c *awsMediaContainer) ElementalMediapackage(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediapackage", attr.AssetImage("assets/aws/media/elemental-mediapackage.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElementalConductor(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-conductor", attr.AssetImage("assets/aws/media/elemental-conductor.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElementalLive(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-live", attr.AssetImage("assets/aws/media/elemental-live.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElementalMedialive(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-medialive", attr.AssetImage("assets/aws/media/elemental-medialive.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElementalMediaconvert(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediaconvert", attr.AssetImage("assets/aws/media/elemental-mediaconvert.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElementalMediastore(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediastore", attr.AssetImage("assets/aws/media/elemental-mediastore.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElementalMediatailor(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediatailor", attr.AssetImage("assets/aws/media/elemental-mediatailor.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElementalServer(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-server", attr.AssetImage("assets/aws/media/elemental-server.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElasticTranscoder(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-transcoder", attr.AssetImage("assets/aws/media/elastic-transcoder.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElementalDelta(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-delta", attr.AssetImage("assets/aws/media/elemental-delta.png"), attr.Shape(attr.None))
}

func (c *awsMediaContainer) ElementalMediaconnect(opts ...attr.Attribute) *node.Node {
	return node.New("elemental-mediaconnect", attr.AssetImage("assets/aws/media/elemental-mediaconnect.png"), attr.Shape(attr.None))
}

func init() {
  const namespace= "aws.media"
  Register(namespace, "ElementalMediapackage", AWSMedia.ElementalMediapackage)
  Register(namespace, "ElementalConductor", AWSMedia.ElementalConductor)
  Register(namespace, "ElementalLive", AWSMedia.ElementalLive)
  Register(namespace, "ElementalMedialive", AWSMedia.ElementalMedialive)
  Register(namespace, "ElementalMediaconvert", AWSMedia.ElementalMediaconvert)
  Register(namespace, "ElementalMediastore", AWSMedia.ElementalMediastore)
  Register(namespace, "ElementalMediatailor", AWSMedia.ElementalMediatailor)
  Register(namespace, "ElementalServer", AWSMedia.ElementalServer)
  Register(namespace, "ElasticTranscoder", AWSMedia.ElasticTranscoder)
  Register(namespace, "ElementalDelta", AWSMedia.ElementalDelta)
  Register(namespace, "ElementalMediaconnect", AWSMedia.ElementalMediaconnect)
}
