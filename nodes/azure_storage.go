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
	opts = append(opts, attr.AssetImage("assets/azure/storage/azurefxtedgefiler.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("azurefxtedgefiler", opts...)
}

func (c *azurestorageContainer) NetappFiles(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/netapp-files.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("netapp-files", opts...)
}

func (c *azurestorageContainer) QueuesStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/queues-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("queues-storage", opts...)
}

func (c *azurestorageContainer) StorageExplorer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/storage-explorer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage-explorer", opts...)
}

func (c *azurestorageContainer) DataBoxEdgeDataBoxGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/data-box-edge---data-box-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-box-edge---data-box-gateway", opts...)
}

func (c *azurestorageContainer) StorsimpleDataManagers(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/storsimple-data-managers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storsimple-data-managers", opts...)
}

func (c *azurestorageContainer) StorsimpleDeviceManagers(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/storsimple-device-managers.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storsimple-device-managers", opts...)
}

func (c *azurestorageContainer) ArchiveStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/archive-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("archive-storage", opts...)
}

func (c *azurestorageContainer) BlobStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/blob-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("blob-storage", opts...)
}

func (c *azurestorageContainer) DataBox(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/data-box.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-box", opts...)
}

func (c *azurestorageContainer) DataLakeStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/data-lake-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-lake-storage", opts...)
}

func (c *azurestorageContainer) StorageAccounts(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/storage-accounts.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage-accounts", opts...)
}

func (c *azurestorageContainer) GeneralStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/general-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("general-storage", opts...)
}

func (c *azurestorageContainer) StorageAccountsClassic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/storage-accounts-classic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage-accounts-classic", opts...)
}

func (c *azurestorageContainer) StorageSyncServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/storage-sync-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage-sync-services", opts...)
}

func (c *azurestorageContainer) TableStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/storage/table-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("table-storage", opts...)
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
