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
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-mediapackage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-mediapackage", opts...)
}

func (c *awsMediaContainer) ElementalConductor(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-conductor.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-conductor", opts...)
}

func (c *awsMediaContainer) ElementalLive(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-live.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-live", opts...)
}

func (c *awsMediaContainer) ElementalMedialive(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-medialive.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-medialive", opts...)
}

func (c *awsMediaContainer) ElementalMediaconvert(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-mediaconvert.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-mediaconvert", opts...)
}

func (c *awsMediaContainer) ElementalMediastore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-mediastore.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-mediastore", opts...)
}

func (c *awsMediaContainer) ElementalMediatailor(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-mediatailor.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-mediatailor", opts...)
}

func (c *awsMediaContainer) ElementalServer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-server.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-server", opts...)
}

func (c *awsMediaContainer) ElasticTranscoder(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elastic-transcoder.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-transcoder", opts...)
}

func (c *awsMediaContainer) ElementalDelta(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-delta.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-delta", opts...)
}

func (c *awsMediaContainer) ElementalMediaconnect(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/media/elemental-mediaconnect.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elemental-mediaconnect", opts...)
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
