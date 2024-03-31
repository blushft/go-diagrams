package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/storage",
}

func (c *storageContainer) CloudStorageGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/storage/cloud-storage-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) FileStorageHdfs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/storage/file-storage-hdfs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) FileStorageNas(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/storage/file-storage-nas.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) HybridBackupRecovery(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/storage/hybrid-backup-recovery.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) HybridCloudDisasterRecovery(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/storage/hybrid-cloud-disaster-recovery.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) Imm(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/storage/imm.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ObjectStorageService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/storage/object-storage-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ObjectTableStore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/storage/object-table-store.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
