package alibabacloud

import "github.com/blushft/go-diagrams/node"

type storageContainer struct {
	path string
	opts []node.Option
}

var Storage = &storageContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/storage",
}

func (c *storageContainer) ObjectTableStore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/storage/object-table-store.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) CloudStorageGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/storage/cloud-storage-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) FileStorageHdfs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/storage/file-storage-hdfs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) FileStorageNas(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/storage/file-storage-nas.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) HybridBackupRecovery(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/storage/hybrid-backup-recovery.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) HybridCloudDisasterRecovery(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/storage/hybrid-cloud-disaster-recovery.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) Imm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/storage/imm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) ObjectStorageService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/storage/object-storage-service.png")}, c.opts, opts)
	return node.New(nopts...)
}
