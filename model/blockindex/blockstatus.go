package blockindex

const (
	// BlockValidUnknown : Unused.
	BlockValidUnknown uint32 = 0

	// BlockValidHeader : parsed , version ok , hash satisfies claimed PoW, 1 <= vtx count <=max
	// timestamp not in future
	BlockValidHeader uint32 = 1

	// BlockValidTree : All parent headers found, difficulty matches, timestamp>= median
	// previous , checkpoint , Implies all parents are also at least TREE
	BlockValidTree uint32 = 2

	// BlockValidTransactions : Only first tx is coinBase, 2 <= coinBase input script length <= 100,
	// transactions valid, no duplicate txIds , sigOps , size , merkle root .
	// Implies all parents are at least TREE but not necessarily TRANSACTIONS.
	// When all parent blocks also have TRANSACTIONS , CBlockIndex ::nChainTx wll be set
	BlockValidTransactions uint32 = 3

	// BlockValidChain : outputs do not overspend inputs , no double spends , coinBase output ok
	// no immature coinBase spends , BIP30.
	// Implies all parents are also at least CHAIN.
	BlockValidChain uint32 = 4

	// BlockValidScripts : Scripts & Signatures ok. Implies all parents are also at least SCRIPTS.
	BlockValidScripts uint32 = 5

	// BlockValidMask : All validity bits
	BlockValidMask = BlockValidHeader |
		BlockValidTree |
		BlockValidTransactions |
		BlockValidChain |
		BlockValidScripts

	// BlockHaveData : full block available in blk*.dat
	BlockHaveData uint32 = 8

	BlockHaveUndo uint32 = 16
	BlockHaveMask        = BlockHaveData | BlockHaveUndo

	// BlockFailedValid : stage after last reached validness failed
	BlockFailedValid uint32 = 32
	// BlockFailedChild : descends from failed block
	BlockFailedChild uint32 = 64
	BlockFailedMask         = BlockFailedValid | BlockFailedChild
)
