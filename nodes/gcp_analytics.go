package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type gcpAnalyticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var GcpAnalytics = &gcpAnalyticsContainer{path: "assets/gcp/analytics"}

func (c *gcpAnalyticsContainer) Datalab(opts ...attr.Attribute) *node.Node {
	return node.New("datalab", attr.AssetImage("assets/gcp/analytics/datalab.png"), attr.Shape(attr.None))
}

func (c *gcpAnalyticsContainer) Genomics(opts ...attr.Attribute) *node.Node {
	return node.New("genomics", attr.AssetImage("assets/gcp/analytics/genomics.png"), attr.Shape(attr.None))
}

func (c *gcpAnalyticsContainer) DataFusion(opts ...attr.Attribute) *node.Node {
	return node.New("data-fusion", attr.AssetImage("assets/gcp/analytics/data-fusion.png"), attr.Shape(attr.None))
}

func (c *gcpAnalyticsContainer) Dataflow(opts ...attr.Attribute) *node.Node {
	return node.New("dataflow", attr.AssetImage("assets/gcp/analytics/dataflow.png"), attr.Shape(attr.None))
}

func (c *gcpAnalyticsContainer) Dataprep(opts ...attr.Attribute) *node.Node {
	return node.New("dataprep", attr.AssetImage("assets/gcp/analytics/dataprep.png"), attr.Shape(attr.None))
}

func (c *gcpAnalyticsContainer) Dataproc(opts ...attr.Attribute) *node.Node {
	return node.New("dataproc", attr.AssetImage("assets/gcp/analytics/dataproc.png"), attr.Shape(attr.None))
}

func (c *gcpAnalyticsContainer) Pubsub(opts ...attr.Attribute) *node.Node {
	return node.New("pubsub", attr.AssetImage("assets/gcp/analytics/pubsub.png"), attr.Shape(attr.None))
}

func (c *gcpAnalyticsContainer) Bigquery(opts ...attr.Attribute) *node.Node {
	return node.New("bigquery", attr.AssetImage("assets/gcp/analytics/bigquery.png"), attr.Shape(attr.None))
}

func (c *gcpAnalyticsContainer) Composer(opts ...attr.Attribute) *node.Node {
	return node.New("composer", attr.AssetImage("assets/gcp/analytics/composer.png"), attr.Shape(attr.None))
}

func (c *gcpAnalyticsContainer) DataCatalog(opts ...attr.Attribute) *node.Node {
	return node.New("data-catalog", attr.AssetImage("assets/gcp/analytics/data-catalog.png"), attr.Shape(attr.None))
}

func init(){
	const namespace = "gcp.analytics"
	Register(namespace, "Datalab", GcpAnalytics.Datalab)
	Register(namespace, "Genomics", GcpAnalytics.Genomics)
	Register(namespace, "DataFusion", GcpAnalytics.DataFusion)
	Register(namespace, "Dataflow", GcpAnalytics.Dataflow)
	Register(namespace, "Dataprep", GcpAnalytics.Dataprep)
	Register(namespace, "Dataproc", GcpAnalytics.Dataproc)
	Register(namespace, "Pubsub", GcpAnalytics.Pubsub)
	Register(namespace, "Bigquery", GcpAnalytics.Bigquery)
	Register(namespace, "Composer", GcpAnalytics.Composer)
	Register(namespace, "DataCatalog", GcpAnalytics.DataCatalog)
}