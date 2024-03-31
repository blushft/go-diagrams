package elastic

import "github.com/blushft/go-diagrams/diagram"

type elasticsearchContainer struct {
	path string
	opts []diagram.NodeOption
}

var Elasticsearch = &elasticsearchContainer{
	opts: diagram.OptionSet{diagram.Provider("elastic"), diagram.NodeShape("none")},
	path: "assets/elastic/elasticsearch",
}

func (c *elasticsearchContainer) SecuritySettings(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/security-settings.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *elasticsearchContainer) Sql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/sql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *elasticsearchContainer) Alerting(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/alerting.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *elasticsearchContainer) Beats(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/beats.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *elasticsearchContainer) Elasticsearch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/elasticsearch.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *elasticsearchContainer) Logstash(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/logstash.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *elasticsearchContainer) Maps(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/maps.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *elasticsearchContainer) Monitoring(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/monitoring.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *elasticsearchContainer) Kibana(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/kibana.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *elasticsearchContainer) MachineLearning(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/elastic/elasticsearch/machine-learning.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
