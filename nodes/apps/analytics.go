package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type analyticsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Analytics = &analyticsContainer{path: "assets/apps/analytics"}

func (c *analyticsContainer) Dbt(opts ...attr.Attribute) *node.Node {
	return node.New("dbt", attr.AssetImage("assets/apps/analytics/dbt.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Hive(opts ...attr.Attribute) *node.Node {
	return node.New("hive", attr.AssetImage("assets/apps/analytics/hive.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Norikra(opts ...attr.Attribute) *node.Node {
	return node.New("norikra", attr.AssetImage("assets/apps/analytics/norikra.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Singer(opts ...attr.Attribute) *node.Node {
	return node.New("singer", attr.AssetImage("assets/apps/analytics/singer.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Spark(opts ...attr.Attribute) *node.Node {
	return node.New("spark", attr.AssetImage("assets/apps/analytics/spark.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Storm(opts ...attr.Attribute) *node.Node {
	return node.New("storm", attr.AssetImage("assets/apps/analytics/storm.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Databricks(opts ...attr.Attribute) *node.Node {
	return node.New("databricks", attr.AssetImage("assets/apps/analytics/databricks.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Flink(opts ...attr.Attribute) *node.Node {
	return node.New("flink", attr.AssetImage("assets/apps/analytics/flink.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Hadoop(opts ...attr.Attribute) *node.Node {
	return node.New("hadoop", attr.AssetImage("assets/apps/analytics/hadoop.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Metabase(opts ...attr.Attribute) *node.Node {
	return node.New("metabase", attr.AssetImage("assets/apps/analytics/metabase.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Tableau(opts ...attr.Attribute) *node.Node {
	return node.New("tableau", attr.AssetImage("assets/apps/analytics/tableau.png"), attr.Shape(attr.None))
}

func (c *analyticsContainer) Beam(opts ...attr.Attribute) *node.Node {
	return node.New("beam", attr.AssetImage("assets/apps/analytics/beam.png"), attr.Shape(attr.None))
}
