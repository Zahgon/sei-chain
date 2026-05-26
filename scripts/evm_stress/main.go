package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	evmRPC     = "http://127.0.0.1:8545"
	chainID    = 713714 // default EVM chain ID for "sei-chain"
	targetTPS  = 500
	numWorkers = 250

	// Total unique sender accounts pre-funded in genesis. Every tx sent by
	// the stress test has a distinct sender with nonce=0. At targetTPS, the
	// pool lasts totalAccounts/targetTPS seconds.
	totalAccounts = 50_000
)

var (
	bigChainID  = big.NewInt(chainID)
	signer      = types.NewLondonSigner(bigChainID)
	maxFee      = big.NewInt(1_000_000_000_000) // 1000 gwei
	priorityFee = big.NewInt(1_000_000_000)     // 1 gwei
	txValue     = big.NewInt(1_000_000_000_001) // 10^12+1 wei: touches both usei balance and wei remainder
)

// nextKey returns a unique deterministic private key for the given index.
func nextKey(idx uint64) *ecdsa.PrivateKey { _ = "STUB: not implemented"; return nil }

// use upper 8 bytes for the index so seed is never all-zero

func keyAddr(key *ecdsa.PrivateKey) common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func evmToSei(addr common.Address) string { _ = "STUB: not implemented"; return "" }

func signTx(tx *types.Transaction, key *ecdsa.PrivateKey) *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

func transfer(nonce uint64, to common.Address, key *ecdsa.PrivateKey) *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

func waitForBalance(ctx context.Context, client *ethclient.Client, addr common.Address) {
	_ = "STUB: not implemented"
	return
}

func main() {
	dumpSeiAddrs := flag.Bool("dump-sei-addrs", false, "print sender sei bech32 addresses for genesis funding and exit")
	flag.Parse()

	// Key 0 = recipient; keys 1..totalAccounts = one-time genesis-funded senders.
	recipient := keyAddr(nextKey(0))

	if *dumpSeiAddrs {
		for i := uint64(1); i <= totalAccounts; i++ {
			fmt.Println(evmToSei(keyAddr(nextKey(i))))
		}
		return
	}

	ctx := context.Background()
	client, err := ethclient.Dial(evmRPC)
	if err != nil {
		panic(fmt.Sprintf("dial %s: %v", evmRPC, err))
	}
	defer client.Close()

	fmt.Printf("recipient: %s\n", recipient.Hex())

	// Wait for genesis accounts to have balance — confirms the node is live.
	waitForBalance(ctx, client, keyAddr(nextKey(1)))

	// Pre-fill the work queue. Each key is used for exactly one tx (nonce=0).
	funded := make(chan *ecdsa.PrivateKey, totalAccounts)
	for i := uint64(1); i <= totalAccounts; i++ {
		funded <- nextKey(i)
	}
	close(funded)

	// Shared rate limiter across all workers: one tick per tx slot.
	ticker := time.NewTicker(time.Second / time.Duration(targetTPS))
	defer ticker.Stop()

	fmt.Printf("starting %d workers, %d unique senders, target %d TPS\n",
		numWorkers, totalAccounts, targetTPS)

	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for key := range funded {
				<-ticker.C
				tx := transfer(0, recipient, key)
				_ = client.SendTransaction(ctx, tx)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("all %d accounts exhausted\n", totalAccounts)
}
