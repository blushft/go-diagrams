package oci

import "github.com/blushft/go-diagrams/diagram"

type databaseContainer struct {
	path string
	opts []diagram.NodeOption
}

var Database = &databaseContainer{
	opts: diagram.OptionSet{diagram.Provider("oci"), diagram.NodeShape("none")},
	path: "assets/oci/database",
}

func (c *databaseContainer) DcatWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/dcat-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DisWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/dis-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Dms(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/dms.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) ScienceWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/science-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) StreamWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/stream-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) BigdataServiceWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/bigdata-service-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Science(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/science.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Stream(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/stream.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) AutonomousWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/autonomous-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Autonomous(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/autonomous.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DatabaseService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/database-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DataflowApacheWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/dataflow-apache-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Dis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/dis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DmsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/dms-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) BigdataService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/bigdata-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DatabaseServiceWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/database-service-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) DataflowApache(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/dataflow-apache.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *databaseContainer) Dcat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/database/dcat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
