package azure

import "github.com/blushft/go-diagrams/node"

type storageContainer struct {
	path string
	opts []node.Option
}

var Storage = &storageContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/storage",
}

func (c *storageContainer) Azurefxtedgefiler(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/azurefxtedgefiler.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) DataBoxEdgeDataBoxGateway(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/data-box-edge---data-box-gateway.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) DataBox(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/data-box.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) DataLakeStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/data-lake-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) QueuesStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/queues-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) StorageAccountsClassic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/storage-accounts-classic.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) StorsimpleDataManagers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/storsimple-data-managers.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) GeneralStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/general-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) StorageSyncServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/storage-sync-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) TableStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/table-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) StorageExplorer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/storage-explorer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) ArchiveStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/archive-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) BlobStorage(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/blob-storage.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) NetappFiles(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/netapp-files.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) StorageAccounts(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/storage-accounts.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *storageContainer) StorsimpleDeviceManagers(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/storage/storsimple-device-managers.png")}, c.opts, opts)
	return node.New(nopts...)
}
