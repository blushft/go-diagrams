package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alibabaCloudStorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabacloudStorage =&alibabaCloudStorageContainer{path: "assets/alibabacloud/storage"}

func (c *alibabaCloudStorageContainer) ObjectStorageService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/storage/object-storage-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("object-storage-service", opts...)
}

func (c *alibabaCloudStorageContainer) ObjectTableStore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/storage/object-table-store.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("object-table-store", opts...)
}

func (c *alibabaCloudStorageContainer) CloudStorageGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/storage/cloud-storage-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-storage-gateway", opts...)
}

func (c *alibabaCloudStorageContainer) FileStorageHdfs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/storage/file-storage-hdfs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("file-storage-hdfs", opts...)
}

func (c *alibabaCloudStorageContainer) FileStorageNas(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/storage/file-storage-nas.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("file-storage-nas", opts...)
}

func (c *alibabaCloudStorageContainer) HybridBackupRecovery(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/storage/hybrid-backup-recovery.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hybrid-backup-recovery", opts...)
}

func (c *alibabaCloudStorageContainer) HybridCloudDisasterRecovery(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/storage/hybrid-cloud-disaster-recovery.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hybrid-cloud-disaster-recovery", opts...)
}

func (c *alibabaCloudStorageContainer) Imm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/storage/imm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("imm", opts...)
}

func init() {
  const namespace = "alibabacloud.storage"
  Register(namespace, "ObjectStorageService", AlibabacloudStorage.ObjectStorageService)
  Register(namespace, "ObjectTableStore", AlibabacloudStorage.ObjectTableStore)
  Register(namespace, "CloudStorageGateway", AlibabacloudStorage.CloudStorageGateway)
  Register(namespace, "FileStorageHdfs", AlibabacloudStorage.FileStorageHdfs)
  Register(namespace, "FileStorageNas", AlibabacloudStorage.FileStorageNas)
  Register(namespace, "HybridBackupRecovery", AlibabacloudStorage.HybridBackupRecovery)
  Register(namespace, "HybridCloudDisasterRecovery", AlibabacloudStorage.HybridCloudDisasterRecovery)
  Register(namespace, "Imm", AlibabacloudStorage.Imm)
}
