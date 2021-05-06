package azure

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Storage = &storageContainer{path: "assets/azure/storage"}

func (c *storageContainer) ArchiveStorage(opts ...attr.Attribute) *node.Node {
	return node.New("archive-storage", attr.AssetImage("assets/azure/storage/archive-storage.png"), attr.Shape(attr.None))
}

func (c *storageContainer) DataBoxEdgeDataBoxGateway(opts ...attr.Attribute) *node.Node {
	return node.New("data-box-edge---data-box-gateway", attr.AssetImage("assets/azure/storage/data-box-edge---data-box-gateway.png"), attr.Shape(attr.None))
}

func (c *storageContainer) DataLakeStorage(opts ...attr.Attribute) *node.Node {
	return node.New("data-lake-storage", attr.AssetImage("assets/azure/storage/data-lake-storage.png"), attr.Shape(attr.None))
}

func (c *storageContainer) NetappFiles(opts ...attr.Attribute) *node.Node {
	return node.New("netapp-files", attr.AssetImage("assets/azure/storage/netapp-files.png"), attr.Shape(attr.None))
}

func (c *storageContainer) StorsimpleDataManagers(opts ...attr.Attribute) *node.Node {
	return node.New("storsimple-data-managers", attr.AssetImage("assets/azure/storage/storsimple-data-managers.png"), attr.Shape(attr.None))
}

func (c *storageContainer) StorageExplorer(opts ...attr.Attribute) *node.Node {
	return node.New("storage-explorer", attr.AssetImage("assets/azure/storage/storage-explorer.png"), attr.Shape(attr.None))
}

func (c *storageContainer) StorsimpleDeviceManagers(opts ...attr.Attribute) *node.Node {
	return node.New("storsimple-device-managers", attr.AssetImage("assets/azure/storage/storsimple-device-managers.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Azurefxtedgefiler(opts ...attr.Attribute) *node.Node {
	return node.New("azurefxtedgefiler", attr.AssetImage("assets/azure/storage/azurefxtedgefiler.png"), attr.Shape(attr.None))
}

func (c *storageContainer) DataBox(opts ...attr.Attribute) *node.Node {
	return node.New("data-box", attr.AssetImage("assets/azure/storage/data-box.png"), attr.Shape(attr.None))
}

func (c *storageContainer) GeneralStorage(opts ...attr.Attribute) *node.Node {
	return node.New("general-storage", attr.AssetImage("assets/azure/storage/general-storage.png"), attr.Shape(attr.None))
}

func (c *storageContainer) QueuesStorage(opts ...attr.Attribute) *node.Node {
	return node.New("queues-storage", attr.AssetImage("assets/azure/storage/queues-storage.png"), attr.Shape(attr.None))
}

func (c *storageContainer) StorageAccountsClassic(opts ...attr.Attribute) *node.Node {
	return node.New("storage-accounts-classic", attr.AssetImage("assets/azure/storage/storage-accounts-classic.png"), attr.Shape(attr.None))
}

func (c *storageContainer) BlobStorage(opts ...attr.Attribute) *node.Node {
	return node.New("blob-storage", attr.AssetImage("assets/azure/storage/blob-storage.png"), attr.Shape(attr.None))
}

func (c *storageContainer) StorageAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("storage-accounts", attr.AssetImage("assets/azure/storage/storage-accounts.png"), attr.Shape(attr.None))
}

func (c *storageContainer) StorageSyncServices(opts ...attr.Attribute) *node.Node {
	return node.New("storage-sync-services", attr.AssetImage("assets/azure/storage/storage-sync-services.png"), attr.Shape(attr.None))
}

func (c *storageContainer) TableStorage(opts ...attr.Attribute) *node.Node {
	return node.New("table-storage", attr.AssetImage("assets/azure/storage/table-storage.png"), attr.Shape(attr.None))
}
