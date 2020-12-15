package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsStorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSStorage = &awsStorageContainer{path: "assets/aws/storage"}

func (c *awsStorageContainer) S3Glacier(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/s3-glacier.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("s3-glacier", opts...)
}

func (c *awsStorageContainer) EfsInfrequentaccessPrimaryBg(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/efs-infrequentaccess-primary-bg.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("efs-infrequentaccess-primary-bg", opts...)
}

func (c *awsStorageContainer) EfsStandardPrimaryBg(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/efs-standard-primary-bg.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("efs-standard-primary-bg", opts...)
}

func (c *awsStorageContainer) ElasticFileSystemEfs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/elastic-file-system-efs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-file-system-efs", opts...)
}

func (c *awsStorageContainer) FsxForWindowsFileServer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/fsx-for-windows-file-server.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("fsx-for-windows-file-server", opts...)
}

func (c *awsStorageContainer) ElasticBlockStoreEbs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/elastic-block-store-ebs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-block-store-ebs", opts...)
}

func (c *awsStorageContainer) Storage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage", opts...)
}

func (c *awsStorageContainer) Backup(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/backup.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("backup", opts...)
}

func (c *awsStorageContainer) Fsx(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/fsx.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("fsx", opts...)
}

func (c *awsStorageContainer) StorageGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/storage-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage-gateway", opts...)
}

func (c *awsStorageContainer) Snowball(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/snowball.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("snowball", opts...)
}

func (c *awsStorageContainer) Snowmobile(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/snowmobile.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("snowmobile", opts...)
}

func (c *awsStorageContainer) CloudendureDisasterRecovery(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/cloudendure-disaster-recovery.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudendure-disaster-recovery", opts...)
}

func (c *awsStorageContainer) FsxForLustre(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/fsx-for-lustre.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("fsx-for-lustre", opts...)
}

func (c *awsStorageContainer) SimpleStorageServiceS3(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/simple-storage-service-s3.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("simple-storage-service-s3", opts...)
}

func (c *awsStorageContainer) SnowballEdge(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/storage/snowball-edge.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("snowball-edge", opts...)
}

func init() {
  const namespace= "aws.storage"
  Register(namespace, "S3Glacier", AWSStorage.S3Glacier)
  Register(namespace, "EfsInfrequentaccessPrimaryBg", AWSStorage.EfsInfrequentaccessPrimaryBg)
  Register(namespace, "EfsStandardPrimaryBg", AWSStorage.EfsStandardPrimaryBg)
  Register(namespace, "ElasticFileSystemEfs", AWSStorage.ElasticFileSystemEfs)
  Register(namespace, "FsxForWindowsFileServer", AWSStorage.FsxForWindowsFileServer)
  Register(namespace, "ElasticBlockStoreEbs", AWSStorage.ElasticBlockStoreEbs)
  Register(namespace, "AWSStorage", AWSStorage.Storage)
  Register(namespace, "Backup", AWSStorage.Backup)
  Register(namespace, "Fsx", AWSStorage.Fsx)
  Register(namespace, "StorageGateway", AWSStorage.StorageGateway)
  Register(namespace, "Snowball", AWSStorage.Snowball)
  Register(namespace, "Snowmobile", AWSStorage.Snowmobile)
  Register(namespace, "CloudendureDisasterRecovery", AWSStorage.CloudendureDisasterRecovery)
  Register(namespace, "FsxForLustre", AWSStorage.FsxForLustre)
  Register(namespace, "SimpleStorageServiceS3", AWSStorage.SimpleStorageServiceS3)
  Register(namespace, "SnowballEdge", AWSStorage.SnowballEdge)
}
