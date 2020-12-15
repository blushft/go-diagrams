package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type elasticsearchContainer struct {
	path  string
	attrs []attr.Attribute
}

var Elasticsearch = &elasticsearchContainer{path: "assets/elastic/elasticsearch"}

func (c *elasticsearchContainer) Monitoring(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/monitoring.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("monitoring", opts...)
}

func (c *elasticsearchContainer) Sql(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/sql.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("sql", opts...)
}

func (c *elasticsearchContainer) Alerting(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/alerting.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("alerting", opts...)
}

func (c *elasticsearchContainer) MachineLearning(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/machine-learning.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("machine-learning", opts...)
}

func (c *elasticsearchContainer) Maps(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/maps.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("maps", opts...)
}

func (c *elasticsearchContainer) Logstash(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/logstash.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("logstash", opts...)
}

func (c *elasticsearchContainer) SecuritySettings(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/security-settings.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security-settings", opts...)
}

func (c *elasticsearchContainer) Beats(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/beats.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("beats", opts...)
}

func (c *elasticsearchContainer) Elasticsearch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/elasticsearch.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elasticsearch", opts...)
}

func (c *elasticsearchContainer) Kibana(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/elastic/elasticsearch/kibana.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kibana", opts...)
}
