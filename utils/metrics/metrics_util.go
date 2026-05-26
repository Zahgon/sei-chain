package metrics

import (
	"math/big"
	"time"

	metrics "github.com/armon/go-metrics"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

func SetupOtelMetricsProvider() error { _ = "STUB: not implemented"; return nil }

func SafeTelemetryIncrCounter(val float32, keys ...string) { _ = "STUB: not implemented"; return }

func SafeTelemetryIncrCounterWithLabels(keys []string, val float32, labels []metrics.Label) {
	_ = "STUB: not implemented"
	return
}

func SafeMetricsIncrCounterWithLabels(keys []string, val float32, labels []metrics.Label) {
	_ = "STUB: not implemented"
	return
}

// Gauge metric with seid version and git commit as labels
// Metric Name:
//
//	seid_version_and_commit
func GaugeSeidVersionAndCommit(version string, commit string) { _ = "STUB: not implemented"; return }

// sei_tx_process_type_count
func IncrTxProcessTypeCounter(processType string) { _ = "STUB: not implemented"; return }

// sei_giga_fallback_to_v2_count
func IncrGigaFallbackToV2Counter() { _ = "STUB: not implemented"; return }

// Measures the time taken to process a block by the process type
// Metric Names:
//
//	sei_process_block_miliseconds
//	sei_process_block_miliseconds_count
//	sei_process_block_miliseconds_sum
func BlockProcessLatency(start time.Time, processType string) { _ = "STUB: not implemented"; return }

// Measures the time taken to execute a sudo msg
// Metric Names:
//
//	sei_deliver_tx_duration_miliseconds
//	sei_deliver_tx_duration_miliseconds_count
//	sei_deliver_tx_duration_miliseconds_sum
func MeasureDeliverTxDuration(start time.Time) { _ = "STUB: not implemented"; return }

// Measures the time taken to execute a batch tx
// Metric Names:
//
//	sei_deliver_batch_tx_duration_miliseconds
//	sei_deliver_batch_tx_duration_miliseconds_count
//	sei_deliver_batch_tx_duration_miliseconds_sum
func MeasureDeliverBatchTxDuration(start time.Time) { _ = "STUB: not implemented"; return }

// sei_oracle_vote_penalty_count
func SetOracleVotePenaltyCount(count uint64, valAddr string, penaltyType string) {
	_ = "STUB: not implemented"
	return
}

// sei_epoch_new
func SetEpochNew(epochNum uint64) { _ = "STUB: not implemented"; return }

// sei_evm_zero_storage_pruned_keys
func IncrEvmZeroStoragePrunedKeys(count uint64) { _ = "STUB: not implemented"; return }

// sei_evm_zero_storage_processed_keys
func IncrEvmZeroStorageProcessedKeys(count uint64) { _ = "STUB: not implemented"; return }

// sei_evm_zero_storage_pruned_bytes
func IncrEvmZeroStoragePrunedBytes(bytes uint64) { _ = "STUB: not implemented"; return }

// Measures number of times a denom's price is updated
// Metric Name:
//
//	sei_oracle_price_update_count
func IncrPriceUpdateDenom(denom string) { _ = "STUB: not implemented"; return }

// Measures the number of times the total block gas wanted in the proposal exceeds the max
// Metric Name:
//
//	sei_failed_total_gas_wanted_check
func IncrFailedTotalGasWantedCheck(proposer string) { _ = "STUB: not implemented"; return }

// Measures number of times a denom's price is updated
// Metric Name:
//
//	sei_oracle_price_update_count
func SetCoinsMinted(amount uint64, denom string) { _ = "STUB: not implemented"; return }

// Measures the number of times the total block gas wanted in the proposal exceeds the max
// Metric Name:
//
//	sei_tx_gas_counter
func IncrGasCounter(gasType string, value int64) { _ = "STUB: not implemented"; return }

// Maximum safe integer representable in float32 (16,777,215)

// Log negative values but don't panic - this shouldn't happen in normal operation

// Cap the value to prevent overflow while logging the incident

// Measures the number of times optimistic processing runs
// Metric Name:
//
//	sei_optimistic_processing_counter
func IncrementOptimisticProcessingCounter(enabled bool) { _ = "STUB: not implemented"; return }

// TODO(PLT-326): remove once dashboards are migrated to evmrpc_* OTEL metrics.
// Measures number of new websocket connects
// Metric Name:
//
//	sei_websocket_connect
func IncWebsocketConnects() { _ = "STUB: not implemented"; return }

// TODO(PLT-326): remove once dashboards are migrated to evmrpc_* OTEL metrics.
// Measures RPC endpoint request throughput
// Metric Name:
//
//	sei_rpc_request_counter
func IncrementRpcRequestCounter(endpoint string, connectionType string, success bool) {
	_ = "STUB: not implemented"
	return
}

// TODO(PLT-326): remove once dashboards are migrated to evmrpc_* OTEL metrics.
// Measures the RPC request latency in milliseconds
// Metric Name:
//
//	sei_rpc_request_latency_ms
func MeasureRpcRequestLatency(endpoint string, connectionType string, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func IncrementErrorMetrics(scenario string, err error) { _ = "STUB: not implemented"; return }

// add other error types to handle as metrics

func IncrementAssociationError(scenario string, err types.AssociationMissingErr) {
	_ = "STUB: not implemented"
	return
}

func IncrementNonceMismatch(tooHigh bool) { _ = "STUB: not implemented"; return }

func AddHistogramMetric(key []string, value float32) { _ = "STUB: not implemented"; return }

// Gauge for gas price paid for transactions
// Metric Name:
//
// sei_evm_effective_gas_price
func HistogramEvmEffectiveGasPrice(gasPrice *big.Int) { _ = "STUB: not implemented"; return }

// Gauge for block base fee
// Metric Name:
//
// sei_evm_block_base_fee
func GaugeEvmBlockBaseFee(baseFee *big.Int, blockHeight int64) { _ = "STUB: not implemented"; return }
