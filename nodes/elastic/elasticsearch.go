package elastic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type elasticsearchContainer struct {
	path  string
	attrs []attr.Attribute
}

var Elasticsearch = &elasticsearchContainer{path: "assets/elastic/elasticsearch"}

func (c *elasticsearchContainer) Alerting(opts ...attr.Attribute) *node.Node {
	return node.New("alerting", attr.AssetImage("assets/elastic/elasticsearch/alerting.png"), attr.Shape(attr.None))
}

func (c *elasticsearchContainer) Beats(opts ...attr.Attribute) *node.Node {
	return node.New("beats", attr.AssetImage("assets/elastic/elasticsearch/beats.png"), attr.Shape(attr.None))
}

func (c *elasticsearchContainer) Sql(opts ...attr.Attribute) *node.Node {
	return node.New("sql", attr.AssetImage("assets/elastic/elasticsearch/sql.png"), attr.Shape(attr.None))
}

func (c *elasticsearchContainer) Elasticsearch(opts ...attr.Attribute) *node.Node {
	return node.New("elasticsearch", attr.AssetImage("assets/elastic/elasticsearch/elasticsearch.png"), attr.Shape(attr.None))
}

func (c *elasticsearchContainer) Kibana(opts ...attr.Attribute) *node.Node {
	return node.New("kibana", attr.AssetImage("assets/elastic/elasticsearch/kibana.png"), attr.Shape(attr.None))
}

func (c *elasticsearchContainer) Logstash(opts ...attr.Attribute) *node.Node {
	return node.New("logstash", attr.AssetImage("assets/elastic/elasticsearch/logstash.png"), attr.Shape(attr.None))
}

func (c *elasticsearchContainer) MachineLearning(opts ...attr.Attribute) *node.Node {
	return node.New("machine-learning", attr.AssetImage("assets/elastic/elasticsearch/machine-learning.png"), attr.Shape(attr.None))
}

func (c *elasticsearchContainer) Maps(opts ...attr.Attribute) *node.Node {
	return node.New("maps", attr.AssetImage("assets/elastic/elasticsearch/maps.png"), attr.Shape(attr.None))
}

func (c *elasticsearchContainer) Monitoring(opts ...attr.Attribute) *node.Node {
	return node.New("monitoring", attr.AssetImage("assets/elastic/elasticsearch/monitoring.png"), attr.Shape(attr.None))
}

func (c *elasticsearchContainer) SecuritySettings(opts ...attr.Attribute) *node.Node {
	return node.New("security-settings", attr.AssetImage("assets/elastic/elasticsearch/security-settings.png"), attr.Shape(attr.None))
}
