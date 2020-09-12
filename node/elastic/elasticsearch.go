package elastic

import "github.com/blushft/go-diagrams/node"

type elasticsearchContainer struct {
	path string
	opts []node.Option
}

var Elasticsearch = &elasticsearchContainer{
	opts: node.OptionSet{node.Provider("elastic"), node.Shape("none")},
	path: "assets/elastic/elasticsearch",
}

func (c *elasticsearchContainer) Elasticsearch(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/elasticsearch.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *elasticsearchContainer) Logstash(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/logstash.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *elasticsearchContainer) SecuritySettings(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/security-settings.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *elasticsearchContainer) Alerting(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/alerting.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *elasticsearchContainer) Kibana(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/kibana.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *elasticsearchContainer) MachineLearning(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/machine-learning.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *elasticsearchContainer) Maps(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/maps.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *elasticsearchContainer) Monitoring(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/monitoring.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *elasticsearchContainer) Sql(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/sql.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *elasticsearchContainer) Beats(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/elastic/elasticsearch/beats.png")}, c.opts, opts)
	return node.New(nopts...)
}
