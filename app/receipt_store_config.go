package app

import (
	seidbconfig "github.com/sei-protocol/sei-chain/sei-db/config"
)

const (
	receiptStoreBackendKey              = "receipt-store.rs-backend"
	receiptStoreDBDirectoryKey          = "receipt-store.db-directory"
	receiptStoreAsyncWriteBufferKey     = "receipt-store.async-write-buffer"
	receiptStorePruneIntervalSecondsKey = "receipt-store.prune-interval-seconds"
)

func readReceiptStoreConfig(homePath string, appOpts seidbconfig.AppOptions) (seidbconfig.ReceiptStoreConfig, error) {
	_ = "STUB: not implemented"
	return *new(seidbconfig.ReceiptStoreConfig), nil
}
