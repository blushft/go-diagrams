package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsAnalytics struct {
	path  string
	attrs []attr.Attribute
}

var AWSAnalyitics = &awsAnalytics{path: "assets/aws/analytics"}

func (c *awsAnalytics) Glue(opts ...attr.Attribute) *node.Node {
	return node.New("glue", attr.AssetImage("assets/aws/analytics/glue.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) Quicksight(opts ...attr.Attribute) *node.Node {
	return node.New("quicksight", attr.AssetImage("assets/aws/analytics/quicksight.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) Cloudsearch(opts ...attr.Attribute) *node.Node {
	return node.New("cloudsearch", attr.AssetImage("assets/aws/analytics/cloudsearch.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) GlueDataCatalog(opts ...attr.Attribute) *node.Node {
	return node.New("glue-data-catalog", attr.AssetImage("assets/aws/analytics/glue-data-catalog.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) KinesisDataAnalytics(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis-data-analytics", attr.AssetImage("assets/aws/analytics/kinesis-data-analytics.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) KinesisDataFirehose(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis-data-firehose", attr.AssetImage("assets/aws/analytics/kinesis-data-firehose.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) ManagedStreamingForKafka(opts ...attr.Attribute) *node.Node {
	return node.New("managed-streaming-for-kafka", attr.AssetImage("assets/aws/analytics/managed-streaming-for-kafka.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) Emr(opts ...attr.Attribute) *node.Node {
	return node.New("emr", attr.AssetImage("assets/aws/analytics/emr.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) KinesisDataStreams(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis-data-streams", attr.AssetImage("assets/aws/analytics/kinesis-data-streams.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) KinesisVideoStreams(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis-video-streams", attr.AssetImage("assets/aws/analytics/kinesis-video-streams.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) Analytics(opts ...attr.Attribute) *node.Node {
	return node.New("analytics", attr.AssetImage("assets/aws/analytics/analytics.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) Athena(opts ...attr.Attribute) *node.Node {
	return node.New("athena", attr.AssetImage("assets/aws/analytics/athena.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) DataPipeline(opts ...attr.Attribute) *node.Node {
	return node.New("data-pipeline", attr.AssetImage("assets/aws/analytics/data-pipeline.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) ElasticsearchService(opts ...attr.Attribute) *node.Node {
	return node.New("elasticsearch-service", attr.AssetImage("assets/aws/analytics/elasticsearch-service.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) EmrHdfsCluster(opts ...attr.Attribute) *node.Node {
	return node.New("emr-hdfs-cluster", attr.AssetImage("assets/aws/analytics/emr-hdfs-cluster.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) RedshiftDenseStorageNode(opts ...attr.Attribute) *node.Node {
	return node.New("redshift-dense-storage-node", attr.AssetImage("assets/aws/analytics/redshift-dense-storage-node.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) Redshift(opts ...attr.Attribute) *node.Node {
	return node.New("redshift", attr.AssetImage("assets/aws/analytics/redshift.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) RedshiftDenseComputeNode(opts ...attr.Attribute) *node.Node {
	return node.New("redshift-dense-compute-node", attr.AssetImage("assets/aws/analytics/redshift-dense-compute-node.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) CloudsearchSearchDocuments(opts ...attr.Attribute) *node.Node {
	return node.New("cloudsearch-search-documents", attr.AssetImage("assets/aws/analytics/cloudsearch-search-documents.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) EmrCluster(opts ...attr.Attribute) *node.Node {
	return node.New("emr-cluster", attr.AssetImage("assets/aws/analytics/emr-cluster.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) GlueCrawlers(opts ...attr.Attribute) *node.Node {
	return node.New("glue-crawlers", attr.AssetImage("assets/aws/analytics/glue-crawlers.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) Kinesis(opts ...attr.Attribute) *node.Node {
	return node.New("kinesis", attr.AssetImage("assets/aws/analytics/kinesis.png"), attr.Shape(attr.None))
}

func (c *awsAnalytics) LakeFormation(opts ...attr.Attribute) *node.Node {
	return node.New("lake-formation", attr.AssetImage("assets/aws/analytics/lake-formation.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.analytics"
  Register(namespace, "Glue", AWSAnalyitics.Glue)
  Register(namespace, "Quicksight", AWSAnalyitics.Quicksight)
  Register(namespace, "Cloudsearch", AWSAnalyitics.Cloudsearch)
  Register(namespace, "GlueDataCatalog", AWSAnalyitics.GlueDataCatalog)
  Register(namespace, "KinesisDataAnalytics", AWSAnalyitics.KinesisDataAnalytics)
  Register(namespace, "KinesisDataFirehose", AWSAnalyitics.KinesisDataFirehose)
  Register(namespace, "ManagedStreamingForKafka", AWSAnalyitics.ManagedStreamingForKafka)
  Register(namespace, "Emr", AWSAnalyitics.Emr)
  Register(namespace, "KinesisDataStreams", AWSAnalyitics.KinesisDataStreams)
  Register(namespace, "KinesisVideoStreams", AWSAnalyitics.KinesisVideoStreams)
  Register(namespace, "Analytics", AWSAnalyitics.Analytics)
  Register(namespace, "Athena", AWSAnalyitics.Athena)
  Register(namespace, "DataPipeline", AWSAnalyitics.DataPipeline)
  Register(namespace, "ElasticsearchService", AWSAnalyitics.ElasticsearchService)
  Register(namespace, "EmrHdfsCluster", AWSAnalyitics.EmrHdfsCluster)
  Register(namespace, "RedshiftDenseStorageNode", AWSAnalyitics.RedshiftDenseStorageNode)
  Register(namespace, "Redshift", AWSAnalyitics.Redshift)
  Register(namespace, "RedshiftDenseComputeNode", AWSAnalyitics.RedshiftDenseComputeNode)
  Register(namespace, "CloudsearchSearchDocuments", AWSAnalyitics.CloudsearchSearchDocuments)
  Register(namespace, "EmrCluster", AWSAnalyitics.EmrCluster)
  Register(namespace, "GlueCrawlers", AWSAnalyitics.GlueCrawlers)
  Register(namespace, "Kinesis", AWSAnalyitics.Kinesis)
  Register(namespace, "LakeFormation", AWSAnalyitics.LakeFormation)
}
