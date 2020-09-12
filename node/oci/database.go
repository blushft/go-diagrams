package oci

import "github.com/blushft/go-diagrams/node"

type databaseContainer struct {
	path string
	opts []node.Option
}

var Database = &databaseContainer{
	opts: node.OptionSet{node.Provider("oci"), node.Shape("none")},
	path: "assets/oci/database",
}

func (c *databaseContainer) DcatWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/dcat-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Dms(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/dms.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) AutonomousWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/autonomous-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DataflowApacheWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/dataflow-apache-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DmsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/dms-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Science(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/science.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) StreamWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/stream-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Stream(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/stream.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Autonomous(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/autonomous.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) BigdataServiceWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/bigdata-service-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) BigdataService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/bigdata-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DatabaseServiceWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/database-service-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Dcat(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/dcat.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DisWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/dis-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DatabaseService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/database-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) DataflowApache(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/dataflow-apache.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) Dis(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/dis.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *databaseContainer) ScienceWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/database/science-white.png")}, c.opts, opts)
	return node.New(nopts...)
}
