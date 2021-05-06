package oci

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type databaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var Database = &databaseContainer{path: "assets/oci/database"}

func (c *databaseContainer) BigdataService(opts ...attr.Attribute) *node.Node {
	return node.New("bigdata-service", attr.AssetImage("assets/oci/database/bigdata-service.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DcatWhite(opts ...attr.Attribute) *node.Node {
	return node.New("dcat-white", attr.AssetImage("assets/oci/database/dcat-white.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Dis(opts ...attr.Attribute) *node.Node {
	return node.New("dis", attr.AssetImage("assets/oci/database/dis.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DmsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("dms-white", attr.AssetImage("assets/oci/database/dms-white.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) ScienceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("science-white", attr.AssetImage("assets/oci/database/science-white.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DatabaseService(opts ...attr.Attribute) *node.Node {
	return node.New("database-service", attr.AssetImage("assets/oci/database/database-service.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Science(opts ...attr.Attribute) *node.Node {
	return node.New("science", attr.AssetImage("assets/oci/database/science.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) StreamWhite(opts ...attr.Attribute) *node.Node {
	return node.New("stream-white", attr.AssetImage("assets/oci/database/stream-white.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) BigdataServiceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("bigdata-service-white", attr.AssetImage("assets/oci/database/bigdata-service-white.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DataflowApacheWhite(opts ...attr.Attribute) *node.Node {
	return node.New("dataflow-apache-white", attr.AssetImage("assets/oci/database/dataflow-apache-white.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Dcat(opts ...attr.Attribute) *node.Node {
	return node.New("dcat", attr.AssetImage("assets/oci/database/dcat.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DisWhite(opts ...attr.Attribute) *node.Node {
	return node.New("dis-white", attr.AssetImage("assets/oci/database/dis-white.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) AutonomousWhite(opts ...attr.Attribute) *node.Node {
	return node.New("autonomous-white", attr.AssetImage("assets/oci/database/autonomous-white.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Autonomous(opts ...attr.Attribute) *node.Node {
	return node.New("autonomous", attr.AssetImage("assets/oci/database/autonomous.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DatabaseServiceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("database-service-white", attr.AssetImage("assets/oci/database/database-service-white.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) DataflowApache(opts ...attr.Attribute) *node.Node {
	return node.New("dataflow-apache", attr.AssetImage("assets/oci/database/dataflow-apache.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Dms(opts ...attr.Attribute) *node.Node {
	return node.New("dms", attr.AssetImage("assets/oci/database/dms.png"), attr.Shape(attr.None))
}

func (c *databaseContainer) Stream(opts ...attr.Attribute) *node.Node {
	return node.New("stream", attr.AssetImage("assets/oci/database/stream.png"), attr.Shape(attr.None))
}
