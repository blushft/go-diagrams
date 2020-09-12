package oci

import "github.com/blushft/go-diagrams/node"

type storageContainer struct {
	path string
	opts []node.Option
}

var Storage = &storageContainer{
	opts: node.OptionSet{node.Provider("oci"), node.Shape("none")},
	path: "assets/oci/storage",
}

func (c *storageContainer) ElasticPerformance(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/elastic-performance.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) ObjectStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/object-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) StorageGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/storage-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) BackupRestoreWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/backup-restore-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) BlockStorageWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/block-storage-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) BucketsWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/buckets-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) FileStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/file-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) ObjectStorageWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/object-storage-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) StorageGatewayWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/storage-gateway-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) BlockStorageClone(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/block-storage-clone.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) BlockStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/block-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) ElasticPerformanceWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/elastic-performance-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) FileStorageWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/file-storage-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) BackupRestore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/backup-restore.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Buckets(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/buckets.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) DataTransfer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/data-transfer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) BlockStorageCloneWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/block-storage-clone-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) DataTransferWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/storage/data-transfer-white.png")}, c.opts, opts)
	return node.New(nopts...)
}
