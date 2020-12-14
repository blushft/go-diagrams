package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ociDatabaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var OciDatabase = &ociDatabaseContainer{path: "assets/oci/database"}

func (c *ociDatabaseContainer) Dcat(opts ...attr.Attribute) *node.Node {
	return node.New("dcat", attr.AssetImage("assets/oci/database/dcat.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) Autonomous(opts ...attr.Attribute) *node.Node {
	return node.New("autonomous", attr.AssetImage("assets/oci/database/autonomous.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) BigdataServiceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("bigdata-service-white", attr.AssetImage("assets/oci/database/bigdata-service-white.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) BigdataService(opts ...attr.Attribute) *node.Node {
	return node.New("bigdata-service", attr.AssetImage("assets/oci/database/bigdata-service.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) Dis(opts ...attr.Attribute) *node.Node {
	return node.New("dis", attr.AssetImage("assets/oci/database/dis.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) DmsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("dms-white", attr.AssetImage("assets/oci/database/dms-white.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) Dms(opts ...attr.Attribute) *node.Node {
	return node.New("dms", attr.AssetImage("assets/oci/database/dms.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) ScienceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("science-white", attr.AssetImage("assets/oci/database/science-white.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) AutonomousWhite(opts ...attr.Attribute) *node.Node {
	return node.New("autonomous-white", attr.AssetImage("assets/oci/database/autonomous-white.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) DatabaseServiceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("database-service-white", attr.AssetImage("assets/oci/database/database-service-white.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) DatabaseService(opts ...attr.Attribute) *node.Node {
	return node.New("database-service", attr.AssetImage("assets/oci/database/database-service.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) DataflowApache(opts ...attr.Attribute) *node.Node {
	return node.New("dataflow-apache", attr.AssetImage("assets/oci/database/dataflow-apache.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) DisWhite(opts ...attr.Attribute) *node.Node {
	return node.New("dis-white", attr.AssetImage("assets/oci/database/dis-white.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) StreamWhite(opts ...attr.Attribute) *node.Node {
	return node.New("stream-white", attr.AssetImage("assets/oci/database/stream-white.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) DataflowApacheWhite(opts ...attr.Attribute) *node.Node {
	return node.New("dataflow-apache-white", attr.AssetImage("assets/oci/database/dataflow-apache-white.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) DcatWhite(opts ...attr.Attribute) *node.Node {
	return node.New("dcat-white", attr.AssetImage("assets/oci/database/dcat-white.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) Science(opts ...attr.Attribute) *node.Node {
	return node.New("science", attr.AssetImage("assets/oci/database/science.png"), attr.Shape(attr.None))
}

func (c *ociDatabaseContainer) Stream(opts ...attr.Attribute) *node.Node {
	return node.New("stream", attr.AssetImage("assets/oci/database/stream.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "oci.database"
  Register(namespace, "Dcat", OciDatabase.Dcat)
  Register(namespace, "Autonomous", OciDatabase.Autonomous)
  Register(namespace, "BigdataServiceWhite", OciDatabase.BigdataServiceWhite)
  Register(namespace, "BigdataService", OciDatabase.BigdataService)
  Register(namespace, "Dis", OciDatabase.Dis)
  Register(namespace, "DmsWhite", OciDatabase.DmsWhite)
  Register(namespace, "Dms", OciDatabase.Dms)
  Register(namespace, "ScienceWhite", OciDatabase.ScienceWhite)
  Register(namespace, "AutonomousWhite", OciDatabase.AutonomousWhite)
  Register(namespace, "DatabaseServiceWhite", OciDatabase.DatabaseServiceWhite)
  Register(namespace, "DatabaseService", OciDatabase.DatabaseService)
  Register(namespace, "DataflowApache", OciDatabase.DataflowApache)
  Register(namespace, "DisWhite", OciDatabase.DisWhite)
  Register(namespace, "StreamWhite", OciDatabase.StreamWhite)
  Register(namespace, "DataflowApacheWhite", OciDatabase.DataflowApacheWhite)
  Register(namespace, "DcatWhite", OciDatabase.DcatWhite)
  Register(namespace, "Science", OciDatabase.Science)
  Register(namespace, "Stream", OciDatabase.Stream)
}
