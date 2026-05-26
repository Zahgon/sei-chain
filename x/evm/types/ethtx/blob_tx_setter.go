package ethtx

import sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

func (tx *BlobTx) SetTo(v string) { _ = "STUB: not implemented"; return }

func (tx *BlobTx) SetAmount(v sdk.Int) { _ = "STUB: not implemented"; return }

func (tx *BlobTx) SetGasFeeCap(v sdk.Int) { _ = "STUB: not implemented"; return }

func (tx *BlobTx) SetGasTipCap(v sdk.Int) { _ = "STUB: not implemented"; return }

func (tx *BlobTx) SetAccesses(v AccessList) { _ = "STUB: not implemented"; return }

func (tx *BlobTx) SetBlobFeeCap(v sdk.Int) { _ = "STUB: not implemented"; return }

func (tx *BlobTx) SetBlobHashes(v [][]byte) { _ = "STUB: not implemented"; return }

func (tx *BlobTx) SetBlobSidecar(v *BlobTxSidecar) { _ = "STUB: not implemented"; return }
