package aws

import "github.com/blushft/go-diagrams/node"

type analyticsContainer struct {
	path string
	opts []node.Option
}

var Analytics = &analyticsContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/analytics",
}

func (c *analyticsContainer) KinesisDataFirehose(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/kinesis-data-firehose.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) KinesisVideoStreams(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/kinesis-video-streams.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Kinesis(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/kinesis.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Quicksight(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/quicksight.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Redshift(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/redshift.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) ElasticsearchService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/elasticsearch-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) EmrHdfsCluster(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/emr-hdfs-cluster.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Emr(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/emr.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) GlueCrawlers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/glue-crawlers.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) GlueDataCatalog(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/glue-data-catalog.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) RedshiftDenseComputeNode(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/redshift-dense-compute-node.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) RedshiftDenseStorageNode(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/redshift-dense-storage-node.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Analytics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/analytics.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Athena(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/athena.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) CloudsearchSearchDocuments(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/cloudsearch-search-documents.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Cloudsearch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/cloudsearch.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) EmrCluster(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/emr-cluster.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) LakeFormation(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/lake-formation.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) KinesisDataStreams(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/kinesis-data-streams.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) ManagedStreamingForKafka(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/managed-streaming-for-kafka.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) DataPipeline(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/data-pipeline.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) Glue(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/glue.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *analyticsContainer) KinesisDataAnalytics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/analytics/kinesis-data-analytics.png")}, c.opts, opts)
	return node.New(nopts...)
}
