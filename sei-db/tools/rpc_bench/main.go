package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LatencyStats struct {
	Method    string
	Total     int
	Errors    int
	Duration  time.Duration
	Latencies []time.Duration
}

func (s *LatencyStats) Report() { _ = "STUB: not implemented"; return }

func configureOutput(outputFile string) (func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var httpClient = &http.Client{
	Timeout: 120 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        200,
		MaxIdleConnsPerHost: 200,
		IdleConnTimeout:     90 * time.Second,
	},
}

var reqID atomic.Int64

func rpcCall(endpoint, method string, params []interface{}) (*RPCResponse, time.Duration, error) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration), nil
}

// benchMethod defines a single RPC method to benchmark.
type benchMethod struct {
	name   string
	params func() []interface{}
	weight int
	heavy  bool // heavy methods get dedicated concurrent phases
}

func (m *benchMethod) call(endpoint string) (string, time.Duration, error) {
	_ = "STUB: not implemented"
	return "", *new(time.Duration), nil
}

type storageSlot struct {
	Address string
	Slot    string
}

type BlockInfo struct {
	Number       int64
	Hash         string
	GasUsed      uint64
	Transactions []string
	Addresses    []string
}

type PerBlockTraceSample struct {
	Block   int64
	Txs     int
	GasUsed uint64
	Latency time.Duration
}

func discoverStorageSlots(endpoint string, txHashes []string, maxTxs int) []storageSlot {
	_ = "STUB: not implemented"
	return nil
}

func getLatestBlockNumber(endpoint string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func getBlockInfo(endpoint string, blockNum int64) (*BlockInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func runConcurrent(concurrency, total int, workFn func(i int) (string, time.Duration, error)) map[string]*LatencyStats {
	_ = "STUB: not implemented"
	return nil
}

func printStats(title string, stats map[string]*LatencyStats) { _ = "STUB: not implemented"; return }

func buildBlockNumbers(latestBlock int64, blockCount int, startBlock, endBlock int64) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeLabel(img *image.RGBA, x, y int, text string, col color.Color) {
	_ = "STUB: not implemented"
	return
}

func setPixel(img *image.RGBA, x, y int, col color.Color) { _ = "STUB: not implemented"; return }

func drawLine(img *image.RGBA, x1, y1, x2, y2 int, col color.Color) {
	_ = "STUB: not implemented"
	return
}

func fillCircle(img *image.RGBA, cx, cy, r int, col color.Color) { _ = "STUB: not implemented"; return }

func scaleValue(value, minValue, maxValue, start, span float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func formatTick(value float64) string { _ = "STUB: not implemented"; return "" }

func writePlotPNG(path, title, xLabel, yLabel string, points [][2]float64, connectPoints bool) error {
	_ = "STUB: not implemented"
	return nil
}

// The output filename is fixed by the caller and joined onto a cleaned plot directory.

func writePerBlockTracePlots(plotDir string, samples []PerBlockTraceSample) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	var (
		endpoint      string
		concurrency   int
		blockCount    int
		startBlock    int64
		endBlock      int64
		requestsPer   int
		methodsFlag   string
		traceDiscover int
		plotDir       string
		outputFile    string
	)
	flag.StringVar(&endpoint, "endpoint", "", "RPC endpoint URL (required)")
	flag.IntVar(&concurrency, "concurrency", 16, "number of concurrent workers")
	flag.IntVar(&blockCount, "blocks", 20, "number of recent blocks to sample")
	flag.Int64Var(&startBlock, "start-block", 0, "explicit starting block number to benchmark (inclusive)")
	flag.Int64Var(&endBlock, "end-block", 0, "explicit ending block number to benchmark (inclusive); defaults to start-block when omitted")
	flag.IntVar(&requestsPer, "requests", 100, "requests per method per phase")
	flag.StringVar(&methodsFlag, "methods", "", "comma-separated methods to run (default: all)")
	flag.IntVar(&traceDiscover, "trace-discover", 5, "txs to trace for storage slot discovery (0 to disable)")
	flag.StringVar(&plotDir, "plot-dir", "", "directory to write per-block trace PNG charts (empty disables plots)")
	flag.StringVar(&outputFile, "output-file", "", "file to write benchmark output to in addition to stdout")
	flag.Parse()

	if endpoint == "" {
		fmt.Fprintf(os.Stderr, "Usage: go run main.go -endpoint <rpc-url> [-concurrency 16] [-blocks 20] [-start-block 100 -end-block 200] [-requests 100] [-methods debug_traceBlockByNumber,eth_getLogs] [-output-file bench.txt]\n")
		os.Exit(1)
	}
	closeOutput, err := configureOutput(outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to configure output file: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := closeOutput(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close output file: %v\n", err)
		}
	}()

	// =========================================================================
	// Discover recent blocks, transactions, and addresses
	// =========================================================================
	fmt.Printf("RPC Read Benchmark\n")
	fmt.Printf("  endpoint:    %s\n", endpoint)
	fmt.Printf("  concurrency: %d\n", concurrency)
	if startBlock > 0 || endBlock > 0 {
		effectiveEndBlock := endBlock
		if effectiveEndBlock == 0 {
			effectiveEndBlock = startBlock
		}
		fmt.Printf("  range:       %d-%d\n", startBlock, effectiveEndBlock)
	} else {
		fmt.Printf("  blocks:      %d recent blocks\n", blockCount)
	}
	fmt.Printf("  requests:    %d per method per phase\n", requestsPer)
	if outputFile != "" {
		fmt.Printf("  output file: %s\n", filepath.Clean(outputFile))
	}

	fmt.Printf("\n--- Discovering blocks ---\n")
	latestBlock, err := getLatestBlockNumber(endpoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get latest block: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Latest block: %d\n", latestBlock)

	blockNums, err := buildBlockNumbers(latestBlock, blockCount, startBlock, endBlock)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid block selection: %v\n", err)
		os.Exit(1)
	}

	var blocks []*BlockInfo
	var allTxHashes []string
	var allAddresses []string
	addrSeen := make(map[string]bool)

	for _, blockNum := range blockNums {
		info, err := getBlockInfo(endpoint, blockNum)
		if err != nil {
			fmt.Printf("  block %d: error %v\n", blockNum, err)
			continue
		}
		blocks = append(blocks, info)
		allTxHashes = append(allTxHashes, info.Transactions...)
		for _, addr := range info.Addresses {
			if !addrSeen[addr] {
				addrSeen[addr] = true
				allAddresses = append(allAddresses, addr)
			}
		}
		avgGasPerTx := 0.0
		if len(info.Transactions) > 0 {
			avgGasPerTx = float64(info.GasUsed) / float64(len(info.Transactions))
		}
		fmt.Printf("  block %d: %d txs, gas=%d, avg_gas/tx=%.1f, %d addresses\n",
			blockNum, len(info.Transactions), info.GasUsed, avgGasPerTx, len(info.Addresses))
	}

	if len(blocks) == 0 {
		fmt.Fprintf(os.Stderr, "No blocks discovered\n")
		os.Exit(1)
	}
	if len(allAddresses) == 0 {
		fmt.Fprintf(os.Stderr, "No addresses found in selected blocks\n")
		os.Exit(1)
	}
	fmt.Printf("Discovered %d blocks, %d transactions, %d unique addresses\n",
		len(blocks), len(allTxHashes), len(allAddresses))

	var allStorageSlots []storageSlot

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	var rngMu sync.Mutex
	randomIntn := func(n int) int {
		rngMu.Lock()
		defer rngMu.Unlock()
		return rng.Intn(n)
	}
	referenceBlock := latestBlock
	if startBlock > 0 || endBlock > 0 {
		referenceBlock = blocks[len(blocks)-1].Number
	}
	referenceHex := fmt.Sprintf("0x%x", referenceBlock)
	randBlock := func() *BlockInfo { return blocks[randomIntn(len(blocks))] }
	randAddr := func() string { return allAddresses[randomIntn(len(allAddresses))] }
	randTxHash := func() string {
		if len(allTxHashes) == 0 {
			return ""
		}
		return allTxHashes[randomIntn(len(allTxHashes))]
	}
	randStorageParams := func() []interface{} {
		if len(allStorageSlots) > 0 {
			s := allStorageSlots[randomIntn(len(allStorageSlots))]
			return []interface{}{s.Address, s.Slot, referenceHex}
		}
		return []interface{}{randAddr(), fmt.Sprintf("0x%064x", randomIntn(10)), referenceHex}
	}
	randLogsParams := func() []interface{} {
		first := randBlock().Number
		second := randBlock().Number
		fromBlock := min(first, second)
		toBlock := max(first, second)
		return []interface{}{map[string]interface{}{
			"fromBlock": fmt.Sprintf("0x%x", fromBlock),
			"toBlock":   fmt.Sprintf("0x%x", toBlock),
		}}
	}

	// =========================================================================
	// Method registry — add new methods here (one line each)
	// =========================================================================
	allMethods := []benchMethod{
		{"debug_traceBlockByNumber", func() []interface{} { return []interface{}{fmt.Sprintf("0x%x", randBlock().Number)} }, 10, true},
		{"debug_traceTransaction", func() []interface{} { return []interface{}{randTxHash()} }, 10, true},
		{"eth_getLogs", func() []interface{} { return randLogsParams() }, 20, true},
		{"eth_getBalance", func() []interface{} { return []interface{}{randAddr(), referenceHex} }, 25, false},
		{"eth_getTransactionCount", func() []interface{} { return []interface{}{randAddr(), referenceHex} }, 15, false},
		{"eth_getCode", func() []interface{} { return []interface{}{randAddr(), referenceHex} }, 15, false},
		{"eth_getStorageAt", func() []interface{} { return randStorageParams() }, 25, false},
	}

	// Skip debug_traceTransaction if no txs discovered
	if len(allTxHashes) == 0 {
		filtered := allMethods[:0]
		for _, m := range allMethods {
			if m.name != "debug_traceTransaction" {
				filtered = append(filtered, m)
			}
		}
		allMethods = filtered
	}

	// Filter by -methods flag if provided
	if methodsFlag != "" {
		allowed := make(map[string]bool)
		for _, m := range strings.Split(methodsFlag, ",") {
			allowed[strings.TrimSpace(m)] = true
		}
		filtered := allMethods[:0]
		for _, m := range allMethods {
			if allowed[m.name] {
				filtered = append(filtered, m)
			}
		}
		allMethods = filtered
	}

	if len(allMethods) == 0 {
		fmt.Fprintf(os.Stderr, "No methods selected\n")
		os.Exit(1)
	}
	hasMethod := func(name string) bool {
		for _, m := range allMethods {
			if m.name == name {
				return true
			}
		}
		return false
	}

	if hasMethod("eth_getStorageAt") && traceDiscover > 0 && len(allTxHashes) > 0 {
		fmt.Printf("\n--- Discovering storage slots (tracing %d txs) ---\n", min(traceDiscover, len(allTxHashes)))
		allStorageSlots = discoverStorageSlots(endpoint, allTxHashes, traceDiscover)
		fmt.Printf("Discovered %d unique storage slots\n", len(allStorageSlots))
	}

	fmt.Printf("  reference:   block %d\n", referenceBlock)
	fmt.Printf("  methods:     ")
	for i, m := range allMethods {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", m.name)
	}
	fmt.Printf("\n")

	// =========================================================================
	// Phase 1: Per-block trace — one trace per discovered block, prints each result
	// =========================================================================
	if hasMethod("debug_traceBlockByNumber") {
		fmt.Printf("\n--- Per-block trace (1 req per block, %d blocks) ---\n", len(blocks))
		fmt.Printf("  %-12s  %-6s  %-12s  %-12s  %s\n", "BLOCK", "TXS", "GAS_USED", "AVG_GAS/TX", "LATENCY")
		fmt.Printf("  %-12s  %-6s  %-12s  %-12s  %s\n", "-----", "---", "--------", "----------", "-------")
		perBlockStats := &LatencyStats{Method: "debug_traceBlockByNumber"}
		perBlockSamples := make([]PerBlockTraceSample, 0, len(blocks))
		for _, b := range blocks {
			hexNum := fmt.Sprintf("0x%x", b.Number)
			resp, lat, err := rpcCall(endpoint, "debug_traceBlockByNumber", []interface{}{hexNum})
			if err == nil && resp != nil && resp.Error != nil {
				err = fmt.Errorf("rpc: %s", resp.Error.Message)
			}
			perBlockStats.Total++
			perBlockStats.Latencies = append(perBlockStats.Latencies, lat)
			errStr := ""
			if err != nil {
				perBlockStats.Errors++
				errStr = fmt.Sprintf("  ERR: %v", err)
			}
			avgGasPerTx := 0.0
			if len(b.Transactions) > 0 {
				avgGasPerTx = float64(b.GasUsed) / float64(len(b.Transactions))
			}
			perBlockSamples = append(perBlockSamples, PerBlockTraceSample{
				Block:   b.Number,
				Txs:     len(b.Transactions),
				GasUsed: b.GasUsed,
				Latency: lat,
			})
			fmt.Printf("  %-12d  %-6d  %-12d  %-12.1f  %s%s\n",
				b.Number, len(b.Transactions), b.GasUsed, avgGasPerTx, lat.Round(time.Millisecond), errStr)
		}
		totalTime := time.Duration(0)
		for _, lat := range perBlockStats.Latencies {
			totalTime += lat
		}
		perBlockStats.Duration = totalTime
		printStats("Per-block trace summary", map[string]*LatencyStats{"debug_traceBlockByNumber": perBlockStats})
		if plotDir != "" {
			paths, err := writePerBlockTracePlots(plotDir, perBlockSamples)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to write plots: %v\n", err)
			} else {
				fmt.Printf("\nWrote per-block trace plots:\n")
				for _, path := range paths {
					fmt.Printf("  %s\n", path)
				}
			}
		}
	}

	// =========================================================================
	// Phase 2: Heavy methods — concurrent blast
	// =========================================================================
	for i := range allMethods {
		m := &allMethods[i]
		if !m.heavy {
			continue
		}
		title := fmt.Sprintf("%s (concurrent x%d)", m.name, concurrency)
		fmt.Printf("\n--- %s ---\n", title)
		s := runConcurrent(concurrency, requestsPer, func(_ int) (string, time.Duration, error) {
			return m.call(endpoint)
		})
		printStats(title, s)
	}

	// =========================================================================
	// Phase 3: Light methods — concurrent per-method
	// =========================================================================
	lightStats := make(map[string]*LatencyStats)
	hasLight := false
	for i := range allMethods {
		m := &allMethods[i]
		if m.heavy {
			continue
		}
		hasLight = true
		s := runConcurrent(concurrency, requestsPer, func(_ int) (string, time.Duration, error) {
			return m.call(endpoint)
		})
		for k, v := range s {
			lightStats[k] = v
		}
	}
	if hasLight {
		printStats(fmt.Sprintf("State reads (concurrent x%d, %d reqs each)", concurrency, requestsPer), lightStats)
	}

	// =========================================================================
	// Phase 4: Mixed workload — all methods, weighted random
	// =========================================================================
	totalWeight := 0
	for _, m := range allMethods {
		totalWeight += m.weight
	}

	totalMixed := requestsPer * 3
	fmt.Printf("\n--- Mixed workload (concurrent x%d, %d total reqs) ---\n", concurrency, totalMixed)
	stats := runConcurrent(concurrency, totalMixed, func(_ int) (string, time.Duration, error) {
		r := randomIntn(totalWeight)
		cumulative := 0
		for i := range allMethods {
			cumulative += allMethods[i].weight
			if r < cumulative {
				return allMethods[i].call(endpoint)
			}
		}
		return allMethods[len(allMethods)-1].call(endpoint)
	})
	printStats("Mixed workload", stats)

	fmt.Printf("\nBenchmark complete.\n")
}
