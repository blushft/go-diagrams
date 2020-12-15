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
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/datalab.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("datalab", opts...)
}

func (c *gcpAnalyticsContainer) Genomics(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/genomics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("genomics", opts...)
}

func (c *gcpAnalyticsContainer) DataFusion(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/data-fusion.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-fusion", opts...)
}

func (c *gcpAnalyticsContainer) Dataflow(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/dataflow.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dataflow", opts...)
}

func (c *gcpAnalyticsContainer) Dataprep(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/dataprep.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dataprep", opts...)
}

func (c *gcpAnalyticsContainer) Dataproc(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/dataproc.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dataproc", opts...)
}

func (c *gcpAnalyticsContainer) Pubsub(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/pubsub.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("pubsub", opts...)
}

func (c *gcpAnalyticsContainer) Bigquery(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/bigquery.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bigquery", opts...)
}

func (c *gcpAnalyticsContainer) Composer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/composer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("composer", opts...)
}

func (c *gcpAnalyticsContainer) DataCatalog(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/analytics/data-catalog.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-catalog", opts...)
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