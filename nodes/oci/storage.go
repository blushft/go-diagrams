package oci

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("oci"), diagram.NodeShape("none")},
	path: "assets/oci/storage",
}

func (c *storageContainer) BlockStorageCloneWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/block-storage-clone-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) BlockStorageWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/block-storage-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) BucketsWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/buckets-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ObjectStorageWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/object-storage-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ObjectStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/object-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) StorageGatewayWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/storage-gateway-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) BackupRestore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/backup-restore.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) DataTransferWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/data-transfer-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) FileStorageWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/file-storage-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) FileStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/file-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) StorageGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/storage-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) BackupRestoreWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/backup-restore-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) BlockStorageClone(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/block-storage-clone.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) DataTransfer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/data-transfer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ElasticPerformanceWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/elastic-performance-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ElasticPerformance(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/elastic-performance.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) BlockStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/block-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Buckets(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/storage/buckets.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
