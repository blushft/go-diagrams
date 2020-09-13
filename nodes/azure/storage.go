package azure

import "github.com/blushft/go-diagrams/diagram"

type storageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Storage = &storageContainer{
	opts: diagram.OptionSet{diagram.Provider("azure"), diagram.NodeShape("none")},
	path: "assets/azure/storage",
}

func (c *storageContainer) Azurefxtedgefiler(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/azurefxtedgefiler.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) DataBoxEdgeDataBoxGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/data-box-edge---data-box-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) NetappFiles(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/netapp-files.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) TableStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/table-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) StorsimpleDeviceManagers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/storsimple-device-managers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) ArchiveStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/archive-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) BlobStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/blob-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) DataBox(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/data-box.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) StorageAccountsClassic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/storage-accounts-classic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) StorageAccounts(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/storage-accounts.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) StorageExplorer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/storage-explorer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) GeneralStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/general-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) QueuesStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/queues-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) DataLakeStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/data-lake-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) StorageSyncServices(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/storage-sync-services.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *storageContainer) StorsimpleDataManagers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/azure/storage/storsimple-data-managers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
