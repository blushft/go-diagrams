package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsAnalyicsContainer struct {
	path  string
	attrs []attr.Attribute
}

var AppsAnalytics = &appsAnalyicsContainer{path: "assets/apps/analytics"}

func (c *appsAnalyicsContainer) Dbt(opts ...attr.Attribute) *node.Node {
	return node.New("dbt", attr.AssetImage("assets/apps/analytics/dbt.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Hive(opts ...attr.Attribute) *node.Node {
	return node.New("hive", attr.AssetImage("assets/apps/analytics/hive.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Norikra(opts ...attr.Attribute) *node.Node {
	return node.New("norikra", attr.AssetImage("assets/apps/analytics/norikra.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Singer(opts ...attr.Attribute) *node.Node {
	return node.New("singer", attr.AssetImage("assets/apps/analytics/singer.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Spark(opts ...attr.Attribute) *node.Node {
	return node.New("spark", attr.AssetImage("assets/apps/analytics/spark.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Storm(opts ...attr.Attribute) *node.Node {
	return node.New("storm", attr.AssetImage("assets/apps/analytics/storm.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Databricks(opts ...attr.Attribute) *node.Node {
	return node.New("databricks", attr.AssetImage("assets/apps/analytics/databricks.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Flink(opts ...attr.Attribute) *node.Node {
	return node.New("flink", attr.AssetImage("assets/apps/analytics/flink.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Hadoop(opts ...attr.Attribute) *node.Node {
	return node.New("hadoop", attr.AssetImage("assets/apps/analytics/hadoop.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Metabase(opts ...attr.Attribute) *node.Node {
	return node.New("metabase", attr.AssetImage("assets/apps/analytics/metabase.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Tableau(opts ...attr.Attribute) *node.Node {
	return node.New("tableau", attr.AssetImage("assets/apps/analytics/tableau.png"), attr.Shape(attr.None))
}

func (c *appsAnalyicsContainer) Beam(opts ...attr.Attribute) *node.Node {
	return node.New("beam", attr.AssetImage("assets/apps/analytics/beam.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.analytics"
	Register(namespace, "Dbt", AppsAnalytics.Dbt)
	Register(namespace, "Hive", AppsAnalytics.Hive)
	Register(namespace, "Norikra", AppsAnalytics.Norikra)
	Register(namespace, "Singer", AppsAnalytics.Singer)
	Register(namespace, "Spark", AppsAnalytics.Spark)
	Register(namespace, "Storm", AppsAnalytics.Storm)
	Register(namespace, "Databricks", AppsAnalytics.Databricks)
	Register(namespace, "Flink", AppsAnalytics.Flink)
	Register(namespace, "Hadoop", AppsAnalytics.Hadoop)
	Register(namespace, "Metabase", AppsAnalytics.Metabase)
	Register(namespace, "Tableau", AppsAnalytics.Tableau)
	Register(namespace, "Beam", AppsAnalytics.Beam)
}