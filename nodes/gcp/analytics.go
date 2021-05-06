package gcp

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type analyticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Analytics = &analyticsContainer{path: "assets/gcp/analytics"}

func (c *analyticsContainer) Composer(opts ...attr.Attribute) *node.Node {
	return node.New("composer", attr.AssetImage("assets/gcp/analytics/composer.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) DataCatalog(opts ...attr.Attribute) *node.Node {
	return node.New("data-catalog", attr.AssetImage("assets/gcp/analytics/data-catalog.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Dataproc(opts ...attr.Attribute) *node.Node {
	return node.New("dataproc", attr.AssetImage("assets/gcp/analytics/dataproc.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Genomics(opts ...attr.Attribute) *node.Node {
	return node.New("genomics", attr.AssetImage("assets/gcp/analytics/genomics.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Bigquery(opts ...attr.Attribute) *node.Node {
	return node.New("bigquery", attr.AssetImage("assets/gcp/analytics/bigquery.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) DataFusion(opts ...attr.Attribute) *node.Node {
	return node.New("data-fusion", attr.AssetImage("assets/gcp/analytics/data-fusion.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Dataflow(opts ...attr.Attribute) *node.Node {
	return node.New("dataflow", attr.AssetImage("assets/gcp/analytics/dataflow.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Datalab(opts ...attr.Attribute) *node.Node {
	return node.New("datalab", attr.AssetImage("assets/gcp/analytics/datalab.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Dataprep(opts ...attr.Attribute) *node.Node {
	return node.New("dataprep", attr.AssetImage("assets/gcp/analytics/dataprep.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Pubsub(opts ...attr.Attribute) *node.Node {
	return node.New("pubsub", attr.AssetImage("assets/gcp/analytics/pubsub.png"), attr.Shape(attr.None))
}
