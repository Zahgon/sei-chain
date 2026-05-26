package types

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

const (
	// module name
	ModuleName = "evm"

	RouterKey = ModuleName

	// StoreKey is string representation of the store key for auth
	StoreKey = "evm"

	TransientStoreKey = "evm_transient"

	ReceiptStoreKey = "receipt"

	// QuerierRoute is the querier route for auth
	QuerierRoute = ModuleName
)

var (
	EVMAddressToSeiAddressKeyPrefix            = []byte{0x01}
	SeiAddressToEVMAddressKeyPrefix            = []byte{0x02}
	StateKeyPrefix                             = []byte{0x03}
	TransientStateKeyPrefix                    = []byte{0x04} // deprecated
	AccountTransientStateKeyPrefix             = []byte{0x05} // deprecated
	TransientModuleStateKeyPrefix              = []byte{0x06} // deprecated
	CodeKeyPrefix                              = []byte{0x07}
	CodeHashKeyPrefix                          = []byte{0x08}
	CodeSizeKeyPrefix                          = []byte{0x09}
	NonceKeyPrefix                             = []byte{0x0a}
	ReceiptKeyPrefix                           = []byte{0x0b}
	WhitelistedCodeHashesForBankSendPrefix     = []byte{0x0c}
	BlockBloomPrefix                           = []byte{0x0d}
	TxHashesPrefix                             = []byte{0x0e} // deprecated
	WhitelistedCodeHashesForDelegateCallPrefix = []byte{0x0f}

	// TxHashPrefix  = []byte{0x10}
	// TxBloomPrefix = []byte{0x11}

	ReplaySeenAddrPrefix = []byte{0x12}
	ReplayedHeight       = []byte{0x13}
	ReplayInitialHeight  = []byte{0x14}

	PointerRegistryPrefix        = []byte{0x15}
	PointerCWCodePrefix          = []byte{0x16}
	PointerReverseRegistryPrefix = []byte{0x17}

	AnteSurplusPrefix  = []byte{0x18} // transient
	DeferredInfoPrefix = []byte{0x19} // transient

	LegacyBlockBloomCutoffHeightKey = []byte{0x1a}
	BaseFeePerGasPrefix             = []byte{0x1b}
	NextBaseFeePerGasPrefix         = []byte{0x1c}
	EvmOnlyBlockBloomPrefix         = []byte{0x1d}
	ZeroStorageCleanupCheckpointKey = []byte{0x1e}
)

var (
	PointerERC20NativePrefix   = []byte{0x0}
	PointerERC20CW20Prefix     = []byte{0x1}
	PointerERC721CW721Prefix   = []byte{0x2}
	PointerCW20ERC20Prefix     = []byte{0x3}
	PointerCW721ERC721Prefix   = []byte{0x4}
	PointerERC1155CW1155Prefix = []byte{0x5}
	PointerCW1155ERC1155Prefix = []byte{0x6}
)

func EVMAddressToSeiAddressKey(evmAddress common.Address) []byte {
	_ = "STUB: not implemented"
	return nil
}

func SeiAddressToEVMAddressKey(seiAddress sdk.AccAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

func StateKey(evmAddress common.Address) []byte { _ = "STUB: not implemented"; return nil }

func ReceiptKey(txHash common.Hash) []byte { _ = "STUB: not implemented"; return nil }

type TransientReceiptKey []byte

func NewTransientReceiptKey(txIndex uint64, txHash common.Hash) TransientReceiptKey {
	_ = "STUB: not implemented"
	return *new(TransientReceiptKey)
}

func (trk TransientReceiptKey) TransactionHash() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}

func BlockBloomKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec

func TxHashesKey(height int64) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec

func PointerERC20NativeKey(token string) []byte { _ = "STUB: not implemented"; return nil }

func PointerERC20CW20Key(cw20Address string) []byte { _ = "STUB: not implemented"; return nil }

func PointerERC721CW721Key(cw721Address string) []byte { _ = "STUB: not implemented"; return nil }

func PointerERC1155CW1155Key(cw1155Address string) []byte { _ = "STUB: not implemented"; return nil }

func PointerCW20ERC20Key(erc20Addr common.Address) []byte { _ = "STUB: not implemented"; return nil }

func PointerCW721ERC721Key(erc721Addr common.Address) []byte { _ = "STUB: not implemented"; return nil }

func PointerCW1155ERC1155Key(erc1155Addr common.Address) []byte {
	_ = "STUB: not implemented"
	return nil
}

func PointerReverseRegistryKey(addr common.Address) []byte { _ = "STUB: not implemented"; return nil }
