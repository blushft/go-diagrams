package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type analyticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Analytics = &analyticsContainer{path: "assets/aws/analytics"}

func (c *analyticsContainer) Glue(opts ...attr.Attribute) *node.Node {
	return node.New("glue", attr.AssetImage("assets/aws/analytics/glue.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Quicksight(opts ...attr.Attribute) *node.Node {
	return node.New("quicksight", attr.AssetImage("assets/aws/analytics/quicksight.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Cloudsearch(opts ...attr.Attribute) *node.Node {
	return node.New("cloudsearch", attr.AssetImage("assets/aws/analytics/cloudsearch.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) GlueDataCatalog(opts ...attr.Attribute) *node.Node {
	return node.New("glue-data-catalog", attr.AssetImage("assets/aws/analytics/glue-data-catalog.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) KinesisDataAnalytics(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis-data-analytics", attr.AssetImage("assets/aws/analytics/kinesis-data-analytics.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) KinesisDataFirehose(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis-data-firehose", attr.AssetImage("assets/aws/analytics/kinesis-data-firehose.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) ManagedStreamingForKafka(opts ...attr.Attribute) *node.Node {
	return node.New("managed-streaming-for-kafka", attr.AssetImage("assets/aws/analytics/managed-streaming-for-kafka.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Emr(opts ...attr.Attribute) *node.Node {
	return node.New("emr", attr.AssetImage("assets/aws/analytics/emr.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) KinesisDataStreams(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis-data-streams", attr.AssetImage("assets/aws/analytics/kinesis-data-streams.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) KinesisVideoStreams(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis-video-streams", attr.AssetImage("assets/aws/analytics/kinesis-video-streams.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Analytics(opts ...attr.Attribute) *node.Node {
	return node.New("analytics", attr.AssetImage("assets/aws/analytics/analytics.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Athena(opts ...attr.Attribute) *node.Node {
	return node.New("athena", attr.AssetImage("assets/aws/analytics/athena.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) DataPipeline(opts ...attr.Attribute) *node.Node {
	return node.New("data-pipeline", attr.AssetImage("assets/aws/analytics/data-pipeline.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) ElasticsearchService(opts ...attr.Attribute) *node.Node {
	return node.New("elasticsearch-service", attr.AssetImage("assets/aws/analytics/elasticsearch-service.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) EmrHdfsCluster(opts ...attr.Attribute) *node.Node {
	return node.New("emr-hdfs-cluster", attr.AssetImage("assets/aws/analytics/emr-hdfs-cluster.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) RedshiftDenseStorageNode(opts ...attr.Attribute) *node.Node {
	return node.New("redshift-dense-storage-node", attr.AssetImage("assets/aws/analytics/redshift-dense-storage-node.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Redshift(opts ...attr.Attribute) *node.Node {
	return node.New("redshift", attr.AssetImage("assets/aws/analytics/redshift.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) RedshiftDenseComputeNode(opts ...attr.Attribute) *node.Node {
	return node.New("redshift-dense-compute-node", attr.AssetImage("assets/aws/analytics/redshift-dense-compute-node.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) CloudsearchSearchDocuments(opts ...attr.Attribute) *node.Node {
	return node.New("cloudsearch-search-documents", attr.AssetImage("assets/aws/analytics/cloudsearch-search-documents.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) EmrCluster(opts ...attr.Attribute) *node.Node {
	return node.New("emr-cluster", attr.AssetImage("assets/aws/analytics/emr-cluster.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) GlueCrawlers(opts ...attr.Attribute) *node.Node {
	return node.New("glue-crawlers", attr.AssetImage("assets/aws/analytics/glue-crawlers.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Kinesis(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis", attr.AssetImage("assets/aws/analytics/kinesis.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) LakeFormation(opts ...attr.Attribute) *node.Node {
	return node.New("lake-formation", attr.AssetImage("assets/aws/analytics/lake-formation.png"), attr.Shape(attr.None))
}
