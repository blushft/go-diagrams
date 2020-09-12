package aws

import "github.com/blushft/go-diagrams/node"

type blockchainContainer struct {
	path string
	opts []node.Option
}

var Blockchain = &blockchainContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/blockchain",
}

func (c *blockchainContainer) ManagedBlockchain(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/blockchain/managed-blockchain.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *blockchainContainer) QuantumLedgerDatabaseQldb(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/blockchain/quantum-ledger-database-qldb.png")}, c.opts, opts)
	return node.New(nopts...)
}
