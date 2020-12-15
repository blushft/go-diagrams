package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsBlockchainContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSBlockchain = &awsBlockchainContainer{path: "assets/aws/blockchain"}

func (c *awsBlockchainContainer) ManagedBlockchain(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/blockchain/managed-blockchain.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("managed-blockchain", opts...)
}

func (c *awsBlockchainContainer) QuantumLedgerDatabaseQldb(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/blockchain/quantum-ledger-database-qldb.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("quantum-ledger-database-qldb", opts...)
}

func init() {
  const namespace = "aws.blockchain"
  Register(namespace, "ManagedBlockchain", AWSBlockchain.ManagedBlockchain)
  Register(namespace, "QuantumLedgerDatabaseQldb", AWSBlockchain.QuantumLedgerDatabaseQldb)
}
