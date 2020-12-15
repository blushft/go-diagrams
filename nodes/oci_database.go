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
	opts = append(opts, attr.AssetImage("assets/oci/database/dcat.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dcat", opts...)
}

func (c *ociDatabaseContainer) Autonomous(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/autonomous.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("autonomous", opts...)
}

func (c *ociDatabaseContainer) BigdataServiceWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/bigdata-service-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bigdata-service-white", opts...)
}

func (c *ociDatabaseContainer) BigdataService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/bigdata-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bigdata-service", opts...)
}

func (c *ociDatabaseContainer) Dis(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/dis.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dis", opts...)
}

func (c *ociDatabaseContainer) DmsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/dms-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dms-white", opts...)
}

func (c *ociDatabaseContainer) Dms(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/dms.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dms", opts...)
}

func (c *ociDatabaseContainer) ScienceWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/science-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("science-white", opts...)
}

func (c *ociDatabaseContainer) AutonomousWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/autonomous-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("autonomous-white", opts...)
}

func (c *ociDatabaseContainer) DatabaseServiceWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/database-service-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database-service-white", opts...)
}

func (c *ociDatabaseContainer) DatabaseService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/database-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("database-service", opts...)
}

func (c *ociDatabaseContainer) DataflowApache(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/dataflow-apache.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dataflow-apache", opts...)
}

func (c *ociDatabaseContainer) DisWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/dis-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dis-white", opts...)
}

func (c *ociDatabaseContainer) StreamWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/stream-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("stream-white", opts...)
}

func (c *ociDatabaseContainer) DataflowApacheWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/dataflow-apache-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dataflow-apache-white", opts...)
}

func (c *ociDatabaseContainer) DcatWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/dcat-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dcat-white", opts...)
}

func (c *ociDatabaseContainer) Science(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/science.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("science", opts...)
}

func (c *ociDatabaseContainer) Stream(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/database/stream.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("stream", opts...)
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
