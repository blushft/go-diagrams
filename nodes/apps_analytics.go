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
	opts = append(opts, attr.AssetImage("assets/apps/analytics/dbt.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dbt", opts...)
}

func (c *appsAnalyicsContainer) Hive(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/hive.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hive", opts...)
}

func (c *appsAnalyicsContainer) Norikra(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/norikra.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("norikra", opts...)
}

func (c *appsAnalyicsContainer) Singer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/singer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("singer", opts...)
}

func (c *appsAnalyicsContainer) Spark(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/spark.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("spark", opts...)
}

func (c *appsAnalyicsContainer) Storm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/storm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storm", opts...)
}

func (c *appsAnalyicsContainer) Databricks(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/databricks.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("databricks", opts...)
}

func (c *appsAnalyicsContainer) Flink(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/flink.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("flink", opts...)
}

func (c *appsAnalyicsContainer) Hadoop(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/hadoop.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hadoop", opts...)
}

func (c *appsAnalyicsContainer) Metabase(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/metabase.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("metabase", opts...)
}

func (c *appsAnalyicsContainer) Tableau(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/tableau.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tableau", opts...)
}

func (c *appsAnalyicsContainer) Beam(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/analytics/beam.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("beam", opts...)
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