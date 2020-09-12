package aws

import "github.com/blushft/go-diagrams/node"

type storageContainer struct {
	path string
	opts []node.Option
}

var Storage = &storageContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/storage",
}

func (c *storageContainer) ElasticBlockStoreEbs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/elastic-block-store-ebs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) FsxForWindowsFileServer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/fsx-for-windows-file-server.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) S3Glacier(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/s3-glacier.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) SimpleStorageServiceS3(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/simple-storage-service-s3.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Snowball(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/snowball.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) StorageGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/storage-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Backup(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/backup.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) EfsInfrequentaccessPrimaryBg(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/efs-infrequentaccess-primary-bg.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) EfsStandardPrimaryBg(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/efs-standard-primary-bg.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) ElasticFileSystemEfs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/elastic-file-system-efs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) FsxForLustre(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/fsx-for-lustre.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Fsx(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/fsx.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Snowmobile(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/snowmobile.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) SnowballEdge(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/snowball-edge.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Storage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) CloudendureDisasterRecovery(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/storage/cloudendure-disaster-recovery.png")}, c.opts, opts)
	return node.New(nopts...)
}
