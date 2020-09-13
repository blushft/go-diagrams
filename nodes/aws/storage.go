package aws

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/storage",
}

func (c *storageContainer) CloudendureDisasterRecovery(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/cloudendure-disaster-recovery.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) FsxForLustre(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/fsx-for-lustre.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Fsx(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/fsx.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Backup(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/backup.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) SimpleStorageServiceS3(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/simple-storage-service-s3.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Snowmobile(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/snowmobile.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) EfsInfrequentaccessPrimaryBg(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/efs-infrequentaccess-primary-bg.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) FsxForWindowsFileServer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/fsx-for-windows-file-server.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) SnowballEdge(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/snowball-edge.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) StorageGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/storage-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Storage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) EfsStandardPrimaryBg(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/efs-standard-primary-bg.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ElasticBlockStoreEbs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/elastic-block-store-ebs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ElasticFileSystemEfs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/elastic-file-system-efs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) S3Glacier(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/s3-glacier.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Snowball(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/storage/snowball.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
