package types

import (
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

// BlockMeta contains meta information.
type BlockMeta struct {
	BlockID   BlockID `json:"block_id"`
	BlockSize int     `json:"block_size,string"`
	Header    Header  `json:"header"`
	NumTxs    int     `json:"num_txs,string"`
}

// NewBlockMeta returns a new BlockMeta.
func NewBlockMeta(block *Block, blockParts *PartSet) *BlockMeta {
	_ = "STUB: not implemented"
	return nil
}

func (bm *BlockMeta) ToProto() *tmproto.BlockMeta { _ = "STUB: not implemented"; return nil }

func BlockMetaFromProto(pb *tmproto.BlockMeta) (*BlockMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateBasic performs basic validation.
func (bm *BlockMeta) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
