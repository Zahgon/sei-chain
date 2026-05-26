package operations

import (
	"encoding/json"
	"os"
	"sync/atomic"
	"time"

	"github.com/spf13/cobra"
)

const defaultTraceProfileTimeout = 120 * time.Second

type traceProfileRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int64         `json:"id"`
}

type traceProfileRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int64           `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *traceRPCError  `json:"error,omitempty"`
}

type traceRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type blockByNumberResult struct {
	Number       string   `json:"number"`
	Hash         string   `json:"hash"`
	Transactions []string `json:"transactions"`
}

type traceProfileResponse struct {
	Trace   json.RawMessage      `json:"trace"`
	Profile traceProfileEnvelope `json:"profile"`
}

type traceProfileEnvelope struct {
	TotalNanos              int64              `json:"totalNanos"`
	HistoricalDBLookupNanos int64              `json:"historicalDbLookupNanos"`
	OtherNanos              int64              `json:"otherNanos"`
	Phases                  traceProfilePhases `json:"phases"`
	Store                   *traceStoreDump    `json:"store,omitempty"`
}

type traceProfilePhases struct {
	LookupTransactionNanos   int64 `json:"lookupTransactionNanos"`
	LoadBlockNanos           int64 `json:"loadBlockNanos"`
	ReplayHistoricalTxsNanos int64 `json:"replayHistoricalTxsNanos"`
	BuildBlockContextNanos   int64 `json:"buildBlockContextNanos"`
	PrepareTxNanos           int64 `json:"prepareTxNanos"`
	ExecutionNanos           int64 `json:"executionNanos"`
	TraceResultNanos         int64 `json:"traceResultNanos"`
}

type traceStoreDump struct {
	Modules map[string]traceStoreModule `json:"modules"`
}

type traceStoreModule struct {
	Stats map[string]traceOperationSummary `json:"stats"`
}

type traceOperationSummary struct {
	Count      int   `json:"count"`
	TotalNanos int64 `json:"totalNanos"`
}

type traceJob struct {
	BlockNumber int64
	BlockHash   string
	TxHash      string
}

type traceRecord struct {
	BlockNumber int64                 `json:"blockNumber"`
	BlockHash   string                `json:"blockHash"`
	TxHash      string                `json:"txHash"`
	Result      *traceProfileResponse `json:"result,omitempty"`
	Error       string                `json:"error,omitempty"`
}

type aggregateOp struct {
	Name       string `json:"name"`
	Count      int    `json:"count"`
	TotalNanos int64  `json:"totalNanos"`
}

type txSummary struct {
	TxHash          string `json:"txHash"`
	BlockNumber     int64  `json:"blockNumber"`
	TotalNanos      int64  `json:"totalNanos"`
	HistoricalNanos int64  `json:"historicalNanos"`
	ExecutionNanos  int64  `json:"executionNanos"`
}

type blockSummary struct {
	BlockNumber int64 `json:"blockNumber"`
	TxCount     int   `json:"txCount"`
	TotalNanos  int64 `json:"totalNanos"`
}

type traceSummary struct {
	Endpoint               string         `json:"endpoint"`
	StartBlock             int64          `json:"startBlock"`
	EndBlock               int64          `json:"endBlock"`
	BlockCount             int            `json:"blockCount"`
	TxCount                int            `json:"txCount"`
	SuccessCount           int            `json:"successCount"`
	ErrorCount             int            `json:"errorCount"`
	GeneratedAt            time.Time      `json:"generatedAt"`
	AverageTotalNanos      int64          `json:"averageTotalNanos"`
	AverageHistoricalNanos int64          `json:"averageHistoricalNanos"`
	AverageExecutionNanos  int64          `json:"averageExecutionNanos"`
	P50TotalNanos          int64          `json:"p50TotalNanos"`
	P95TotalNanos          int64          `json:"p95TotalNanos"`
	P50HistoricalNanos     int64          `json:"p50HistoricalNanos"`
	P95HistoricalNanos     int64          `json:"p95HistoricalNanos"`
	PhaseTotals            []aggregateOp  `json:"phaseTotals"`
	StoreTotals            []aggregateOp  `json:"storeTotals"`
	TopTransactions        []txSummary    `json:"topTransactions"`
	TopBlocks              []blockSummary `json:"topBlocks"`
}

func TraceProfileReportCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runTraceProfileReport(endpoint string, startBlock, endBlock int64, outputDir string, concurrency, maxTransactions int, traceConfig map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // outputDir is an operator-supplied CLI flag for this offline seidb tool

func collectTraceJobs(endpoint string, startBlock, endBlock int64, maxTransactions int) ([]traceJob, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func runTraceWorkers(endpoint string, jobs []traceJob, concurrency int, traceConfig map[string]interface{}) ([]traceRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeAndSummarize(results []traceRecord, rawFile *os.File, endpoint string, startBlock, endBlock int64, blockCount int) (traceSummary, error) {
	_ = "STUB: not implemented"
	return *new(traceSummary), nil
}

func topNTxs(items []txSummary, n int) []txSummary { _ = "STUB: not implemented"; return nil }

func topNBlocks(items map[int64]*blockSummary, n int) []blockSummary {
	_ = "STUB: not implemented"
	return nil
}

func addOp(m map[string]traceOperationSummary, name string, nanos int64) {
	_ = "STUB: not implemented"
	return
}

func addNamedOp(m map[string]traceOperationSummary, name string, stats traceOperationSummary) {
	_ = "STUB: not implemented"
	return
}

func sortedOps(m map[string]traceOperationSummary, limit int) []aggregateOp {
	_ = "STUB: not implemented"
	return nil
}

func percentile(values []int64, pct float64) int64 { _ = "STUB: not implemented"; return 0 }

var traceReqID atomic.Int64

func fetchBlockByNumber(endpoint string, blockNumber int64) (*blockByNumberResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fetchTraceProfile(endpoint string, job traceJob, traceConfig map[string]interface{}) traceRecord {
	_ = "STUB: not implemented"
	return *new(traceRecord)
}

func doRPC(endpoint, method string, params []interface{}, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
