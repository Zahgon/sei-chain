package types

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/hashable"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// LaneID represents a lane identifier (currently it is the same as NodeID,
// since the producer uniquely identifies the lane).
type LaneID = PublicKey

// NodeID represents a unique identifier for a node in the network.
type NodeID string

// BlockNumber is the number of a block in a lane.
type BlockNumber uint64

// GlobalBlockNumber is the number of a block in the global chain.
type GlobalBlockNumber uint64

// BlockHeaderHash is the hash of a BlockHeader.
type BlockHeaderHash hashable.Hash[*pb.BlockHeader]

// Bytes converts the BlockHeaderHash to a byte slice.
func (h BlockHeaderHash) Bytes() []byte  { _ = "STUB: not implemented"; return nil }
func (h BlockHeaderHash) String() string { _ = "STUB: not implemented"; return "" }

func ParseBlockHeaderHash(bytes []byte) (BlockHeaderHash, error) {
	_ = "STUB: not implemented"
	return *new(BlockHeaderHash), nil
}

// BlockHeader .
type BlockHeader struct {
	utils.ReadOnly
	lane        LaneID
	blockNumber BlockNumber
	parentHash  BlockHeaderHash
	payloadHash PayloadHash
}

// Lane .
func (h *BlockHeader) Lane() LaneID {
	_ = "STUB: not implemented"

	// BlockNumber .
	return *new(LaneID)
}

func (h *BlockHeader) BlockNumber() BlockNumber {
	_ = "STUB: not implemented"
	return *

	// ParentHash .
	new(BlockNumber)
}

func (h *BlockHeader) ParentHash() BlockHeaderHash {
	_ = "STUB: not implemented"
	return *

	// PayloadHash .
	new(BlockHeaderHash)
}

func (h *BlockHeader) PayloadHash() PayloadHash {
	_ = "STUB: not implemented"
	return *

	// Next return the block number of the next header.
	new(PayloadHash)
}

func (h *BlockHeader) Next() BlockNumber {
	_ = "STUB: not implemented"
	return *

	// Verify verifies the BlockHeader against the committee.
	new(BlockNumber)
}

func (h *BlockHeader) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

const standardTxBytes uint64 = 1024

// Maximum number of transactions in a block.
const MaxTxsPerBlock uint64 = 2000

// Maximum total size of all the transactions.
// It can be split arbitrarily across transactions (1 large, 2000 small ones, etc.)
// up to MaxTxsPerBlock limit.
const MaxTxsBytesPerBlock = MaxTxsPerBlock * standardTxBytes

// Upper bound on the block proto encoding.
var MaxBlockProtoSize = func() uint64 {
	// Payload.Txs represents the variable part of the Block size.
	// Proto size is maximized if we distribute data evenly across transactions.
	tx := make([]byte, standardTxBytes)
	txs := make([][]byte, MaxTxsPerBlock)
	for i := range txs {
		txs[i] = tx
	}
	// Crude estimate of all other fields.
	const otherFields = 100 * 1024
	return otherFields + uint64(protoutils.Size(&pb.Block{Payload: &pb.Payload{Txs: txs}}))
}()

// Block .
type Block struct {
	utils.ReadOnly
	header  *BlockHeader
	payload *Payload
}

// GlobalBlock is a finalized block with global block number.
type GlobalBlock struct {
	Header       *BlockHeader
	Timestamp    time.Time
	GlobalNumber GlobalBlockNumber
	Payload      *Payload
	// Highest known finalized state.
	FinalAppState utils.Option[*AppProposal]
}

// NewBlock creates a new Block.
func NewBlock(
	lane LaneID,
	blockNumber BlockNumber,
	parentHash BlockHeaderHash,
	payload *Payload,
) *Block {
	_ = "STUB: not implemented"
	return nil
}

// Header .
func (b *Block) Header() *BlockHeader {
	_ = "STUB: not implemented"

	// Payload .
	return nil
}

func (b *Block) Payload() *Payload {
	_ = "STUB: not implemented"

	// Verify validates the Block.
	return nil
}

func (b *Block) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// Hash of the BlockHeader.
func (h *BlockHeader) Hash() BlockHeaderHash {
	_ = "STUB: not implemented"
	return *new(BlockHeaderHash)
}

// PayloadHash is the hash of a Payload.
type PayloadHash hashable.Hash[*pb.Payload]

// PayloadBuilder builds a Payload.
type PayloadBuilder struct {
	CreatedAt time.Time
	TotalGas  uint64
	EdgeCount int64
	Coinbase  []byte
	Basefee   int64
	Txs       [][]byte
}

// Payload .
type Payload struct {
	utils.ReadOnly
	p PayloadBuilder
}

// Build builds the Payload.
func (b PayloadBuilder) Build() (*Payload, error) { _ = "STUB: not implemented"; return nil, nil }

// ToBuilder converts the Payload to a PayloadBuilder.
func (p *Payload) ToBuilder() PayloadBuilder {
	_ = "STUB: not implemented"

	// CreatedAt .
	return *new(PayloadBuilder)
}

func (p *Payload) CreatedAt() time.Time {
	_ = "STUB: not implemented"
	return *

	// TotalGas .
	new(time.Time)
}

func (p *Payload) TotalGas() uint64 {
	_ = "STUB: not implemented"

	// EdgeCount .
	return 0
}

func (p *Payload) EdgeCount() int64 { _ = "STUB: not implemented"; return 0 }

// Coinbase .
func (p *Payload) Coinbase() []byte {
	_ = "STUB: not implemented"

	// Basefee .
	return nil
}

func (p *Payload) Basefee() int64 {
	_ = "STUB: not implemented"

	// Txs .
	return 0
}

func (p *Payload) Txs() [][]byte {
	_ = "STUB: not implemented"

	// Hash of the Payload.
	return nil
}

func (p *Payload) Hash() PayloadHash { _ = "STUB: not implemented"; return *new(PayloadHash) }

// BlockHeaderConv is a protobuf converter for BlockHeader.
var BlockHeaderConv = protoutils.Conv[*BlockHeader, *pb.BlockHeader]{
	Encode: func(h *BlockHeader) *pb.BlockHeader {
		return &pb.BlockHeader{
			Lane:        PublicKeyConv.Encode(h.lane),
			BlockNumber: utils.Alloc(uint64(h.blockNumber)),
			ParentHash:  h.parentHash[:],
			PayloadHash: h.payloadHash[:],
		}
	},
	Decode: func(h *pb.BlockHeader) (*BlockHeader, error) {
		payloadHash, err := hashable.ParseHash[*pb.Payload](h.PayloadHash)
		if err != nil {
			return nil, fmt.Errorf("PayloadHash: %w", err)
		}
		parentHash, err := hashable.ParseHash[*pb.BlockHeader](h.ParentHash)
		if err != nil {
			return nil, fmt.Errorf("ParentHash: %w", err)
		}
		lane, err := PublicKeyConv.DecodeReq(h.Lane)
		if err != nil {
			return nil, fmt.Errorf("lane: %w", err)
		}
		if h.BlockNumber == nil {
			return nil, fmt.Errorf("BlockNumber: missing")
		}
		return &BlockHeader{
			lane:        lane,
			blockNumber: BlockNumber(*h.BlockNumber),
			parentHash:  BlockHeaderHash(parentHash),
			payloadHash: PayloadHash(payloadHash),
		}, nil
	},
}

// PayloadConv is a protobuf converter for Payload.
var PayloadConv = protoutils.Conv[*Payload, *pb.Payload]{
	Encode: func(p *Payload) *pb.Payload {
		return &pb.Payload{
			CreatedAt: TimeConv.Encode(p.p.CreatedAt),
			TotalGas:  utils.Alloc(p.p.TotalGas),
			EdgeCount: utils.Alloc(p.p.EdgeCount),
			Coinbase:  p.p.Coinbase,
			Basefee:   utils.Alloc(p.p.Basefee),
			Txs:       p.p.Txs,
		}
	},
	Decode: func(p *pb.Payload) (*Payload, error) {
		createdAt, err := TimeConv.DecodeReq(p.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("created_at: %w", err)
		}
		if p.TotalGas == nil {
			return nil, fmt.Errorf("TotalGas: missing")
		}
		if p.EdgeCount == nil {
			return nil, fmt.Errorf("EdgeCount: missing")
		}
		if p.Basefee == nil {
			return nil, fmt.Errorf("Basefee: missing")
		}
		return PayloadBuilder{
			CreatedAt: createdAt,
			TotalGas:  *p.TotalGas,
			EdgeCount: *p.EdgeCount,
			Coinbase:  p.Coinbase,
			Basefee:   *p.Basefee,
			Txs:       p.Txs,
		}.Build()
	},
}

// BlockConv is a protobuf converter for Block.
var BlockConv = protoutils.Conv[*Block, *pb.Block]{
	Encode: func(b *Block) *pb.Block {
		return &pb.Block{
			Header:  BlockHeaderConv.Encode(b.header),
			Payload: PayloadConv.Encode(b.payload),
		}
	},
	Decode: func(b *pb.Block) (*Block, error) {
		header, err := BlockHeaderConv.DecodeReq(b.Header)
		if err != nil {
			return nil, err
		}
		payload, err := PayloadConv.DecodeReq(b.Payload)
		if err != nil {
			return nil, err
		}
		return &Block{header: header, payload: payload}, nil
	},
}

// CalculateBlockHash calculates the hash of a block.
func (b *GlobalBlock) CalculateBlockHash() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

//nolint:gosec // block timestamps are always positive post-epoch values
//nolint:gosec // block numbers are within int64 range for all practical chain heights
