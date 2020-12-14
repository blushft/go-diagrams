package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azurestorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureStorage = &azurestorageContainer{path: "assets/azure/storage"}

func (c *azurestorageContainer) Azurefxtedgefiler(opts ...attr.Attribute) *node.Node {
	return node.New("azurefxtedgefiler", attr.AssetImage("assets/azure/storage/azurefxtedgefiler.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) NetappFiles(opts ...attr.Attribute) *node.Node {
	return node.New("netapp-files", attr.AssetImage("assets/azure/storage/netapp-files.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) QueuesStorage(opts ...attr.Attribute) *node.Node {
	return node.New("queues-storage", attr.AssetImage("assets/azure/storage/queues-storage.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) StorageExplorer(opts ...attr.Attribute) *node.Node {
	return node.New("storage-explorer", attr.AssetImage("assets/azure/storage/storage-explorer.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) DataBoxEdgeDataBoxGateway(opts ...attr.Attribute) *node.Node {
	return node.New("data-box-edge---data-box-gateway", attr.AssetImage("assets/azure/storage/data-box-edge---data-box-gateway.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) StorsimpleDataManagers(opts ...attr.Attribute) *node.Node {
	return node.New("storsimple-data-managers", attr.AssetImage("assets/azure/storage/storsimple-data-managers.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) StorsimpleDeviceManagers(opts ...attr.Attribute) *node.Node {
	return node.New("storsimple-device-managers", attr.AssetImage("assets/azure/storage/storsimple-device-managers.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) ArchiveStorage(opts ...attr.Attribute) *node.Node {
	return node.New("archive-storage", attr.AssetImage("assets/azure/storage/archive-storage.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) BlobStorage(opts ...attr.Attribute) *node.Node {
	return node.New("blob-storage", attr.AssetImage("assets/azure/storage/blob-storage.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) DataBox(opts ...attr.Attribute) *node.Node {
	return node.New("data-box", attr.AssetImage("assets/azure/storage/data-box.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) DataLakeStorage(opts ...attr.Attribute) *node.Node {
	return node.New("data-lake-storage", attr.AssetImage("assets/azure/storage/data-lake-storage.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) StorageAccounts(opts ...attr.Attribute) *node.Node {
	return node.New("storage-accounts", attr.AssetImage("assets/azure/storage/storage-accounts.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) GeneralStorage(opts ...attr.Attribute) *node.Node {
	return node.New("general-storage", attr.AssetImage("assets/azure/storage/general-storage.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) StorageAccountsClassic(opts ...attr.Attribute) *node.Node {
	return node.New("storage-accounts-classic", attr.AssetImage("assets/azure/storage/storage-accounts-classic.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) StorageSyncServices(opts ...attr.Attribute) *node.Node {
	return node.New("storage-sync-services", attr.AssetImage("assets/azure/storage/storage-sync-services.png"), attr.Shape(attr.None))
}

func (c *azurestorageContainer) TableStorage(opts ...attr.Attribute) *node.Node {
	return node.New("table-storage", attr.AssetImage("assets/azure/storage/table-storage.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.storage"
  Register(namespace, "Azurefxtedgefiler", azureStorage.Azurefxtedgefiler)
  Register(namespace, "NetappFiles", azureStorage.NetappFiles)
  Register(namespace, "QueuesStorage", azureStorage.QueuesStorage)
  Register(namespace, "StorageExplorer", azureStorage.StorageExplorer)
  Register(namespace, "DataBoxEdgeDataBoxGateway", azureStorage.DataBoxEdgeDataBoxGateway)
  Register(namespace, "StorsimpleDataManagers", azureStorage.StorsimpleDataManagers)
  Register(namespace, "StorsimpleDeviceManagers", azureStorage.StorsimpleDeviceManagers)
  Register(namespace, "ArchiveStorage", azureStorage.ArchiveStorage)
  Register(namespace, "BlobStorage", azureStorage.BlobStorage)
  Register(namespace, "DataBox", azureStorage.DataBox)
  Register(namespace, "DataLakeStorage", azureStorage.DataLakeStorage)
  Register(namespace, "StorageAccounts", azureStorage.StorageAccounts)
  Register(namespace, "GeneralStorage", azureStorage.GeneralStorage)
  Register(namespace, "StorageAccountsClassic", azureStorage.StorageAccountsClassic)
  Register(namespace, "StorageSyncServices", azureStorage.StorageSyncServices)
  Register(namespace, "TableStorage", azureStorage.TableStorage)
}
