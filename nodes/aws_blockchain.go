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
	return node.New("managed-blockchain", attr.AssetImage("assets/aws/blockchain/managed-blockchain.png"), attr.Shape(attr.None))
}

func (c *awsBlockchainContainer) QuantumLedgerDatabaseQldb(opts ...attr.Attribute) *node.Node {
	return node.New("quantum-ledger-database-qldb", attr.AssetImage("assets/aws/blockchain/quantum-ledger-database-qldb.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.blockchain"
  Register(namespace, "ManagedBlockchain", AWSBlockchain.ManagedBlockchain)
  Register(namespace, "QuantumLedgerDatabaseQldb", AWSBlockchain.QuantumLedgerDatabaseQldb)
}
