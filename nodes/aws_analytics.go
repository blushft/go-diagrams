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
	opts = append(opts, attr.AssetImage("assets/aws/analytics/glue.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("glue", opts...)
}

func (c *awsAnalytics) Quicksight(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/quicksight.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("quicksight", opts...)
}

func (c *awsAnalytics) Cloudsearch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/cloudsearch.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudsearch", opts...)
}

func (c *awsAnalytics) GlueDataCatalog(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/glue-data-catalog.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("glue-data-catalog", opts...)
}

func (c *awsAnalytics) KinesisDataAnalytics(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/kinesis-data-analytics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kinesis-data-analytics", opts...)
}

func (c *awsAnalytics) KinesisDataFirehose(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/kinesis-data-firehose.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kinesis-data-firehose", opts...)
}

func (c *awsAnalytics) ManagedStreamingForKafka(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/managed-streaming-for-kafka.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("managed-streaming-for-kafka", opts...)
}

func (c *awsAnalytics) Emr(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/emr.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("emr", opts...)
}

func (c *awsAnalytics) KinesisDataStreams(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/kinesis-data-streams.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kinesis-data-streams", opts...)
}

func (c *awsAnalytics) KinesisVideoStreams(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/kinesis-video-streams.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kinesis-video-streams", opts...)
}

func (c *awsAnalytics) Analytics(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/analytics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("analytics", opts...)
}

func (c *awsAnalytics) Athena(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/athena.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("athena", opts...)
}

func (c *awsAnalytics) DataPipeline(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/data-pipeline.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-pipeline", opts...)
}

func (c *awsAnalytics) ElasticsearchService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/elasticsearch-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elasticsearch-service", opts...)
}

func (c *awsAnalytics) EmrHdfsCluster(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/emr-hdfs-cluster.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("emr-hdfs-cluster", opts...)
}

func (c *awsAnalytics) RedshiftDenseStorageNode(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/redshift-dense-storage-node.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("redshift-dense-storage-node", opts...)
}

func (c *awsAnalytics) Redshift(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/redshift.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("redshift", opts...)
}

func (c *awsAnalytics) RedshiftDenseComputeNode(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/redshift-dense-compute-node.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("redshift-dense-compute-node", opts...)
}

func (c *awsAnalytics) CloudsearchSearchDocuments(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/cloudsearch-search-documents.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudsearch-search-documents", opts...)
}

func (c *awsAnalytics) EmrCluster(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/emr-cluster.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("emr-cluster", opts...)
}

func (c *awsAnalytics) GlueCrawlers(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/glue-crawlers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("glue-crawlers", opts...)
}

func (c *awsAnalytics) Kinesis(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/kinesis.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kinesis", opts...)
}

func (c *awsAnalytics) LakeFormation(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/analytics/lake-formation.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("lake-formation", opts...)
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
