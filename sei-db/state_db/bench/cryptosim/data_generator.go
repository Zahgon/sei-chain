package cryptosim

import (
	"github.com/sei-protocol/sei-chain/sei-db/common/keys"
	crand "github.com/sei-protocol/sei-chain/sei-db/common/rand"
)

const (
	// Used to store the next account ID in the database.
	accountIdCounterKey = "accountIdCounterKey"
	// Used to store the next ERC20 contract ID in the database.
	erc20IdCounterKey = "erc20IdCounterKey"
	// Used to store the next block number in the database.
	blockNumberCounterKey = "blockNumberCounterKey"

	// Use the code hash as a proxy. There is currently no mechanism to force FlatKV to update the account balance
	// field, and code hash keys will cause the account DB to get updated, which is the important part for this
	// simulation.
	accountKeyPrefix = keys.EVMKeyCodeHash
)

// Generates random data for the benchmark. This is not a thread safe utility.
type DataGenerator struct {
	config *CryptoSimConfig

	// The next account ID to be used when creating a new account.
	nextAccountID int64

	// The next ERC20 contract ID to be used when creating a new ERC20 contract.
	nextErc20ContractID int64

	// The next block number at startup time. Not updated after initialization;
	// the block builder tracks the ongoing value.
	initialNextBlockNumber uint64

	// The random number generator.
	rand *crand.CannedRandom

	// The address of the fee account (i.e. the account that collects gas fees). This is a special account
	// and has account ID 0. Since we reuse this account very often, it is cached for performance.
	feeCollectionAddress []byte

	// The database for the benchmark.
	database *Database

	// The highest account ID that has been read in the current block.
	// Since there are multiple threads of execution, it's possible that one executor may create a new account,
	// and another may attempt to read/write it before the account is actually created. To avoid this, we will only
	// choose read/write targets from acccounts that were created before the current block. This field tracks the
	// highest account ID that was created before the current block.
	highestSafeAccountIDInBlock int64

	// The current number of cold accounts. These are accounts that are not used frequently, but are not
	// entirely dormant.
	numberOfColdAccounts int64

	// The metrics for the benchmark.
	metrics *CryptosimMetrics
}

// Creates a new data generator.
func NewDataGenerator(
	config *CryptoSimConfig,
	database *Database,
	rand *crand.CannedRandom,
	metrics *CryptosimMetrics,
) (*DataGenerator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // G115 - persisted counter value, overflow acceptable

//nolint:gosec // G115 - persisted counter value, overflow acceptable

//nolint:gosec

// Get the next account ID to be used when creating a new account. This is also the total number of accounts
// currently in the database.
func (d *DataGenerator) NextAccountID() int64 { _ = "STUB: not implemented"; return 0 }

// NumberOfColdAccounts returns the current count of cold accounts.
func (d *DataGenerator) NumberOfColdAccounts() int64 { _ = "STUB: not implemented"; return 0 }

// ReportAccountCounts updates the metrics with the current account counts (total, hot, cold).
func (d *DataGenerator) ReportAccountCounts() { _ = "STUB: not implemented"; return }

// Get the next ERC20 contract ID to be used when creating a new ERC20 contract. This is also the total number of
// ERC20 contracts currently in the database.
func (d *DataGenerator) NextErc20ContractID() int64 { _ = "STUB: not implemented"; return 0 }

// Get the next block number as it was at startup time.
func (d *DataGenerator) InitialNextBlockNumber() uint64 { _ = "STUB: not implemented"; return 0 }

// Creates a new account and optionally writes it to the database. Returns the address of the new
// account and whether it is a cold account (vs dormant).
func (d *DataGenerator) CreateNewAccount(
	// The number of bytes to allocate for the account data.
	accountSize int,
	// If true, the account will be immediately written to the database.
	write bool,
) (id int64, address []byte, isCold bool, err error) {
	_ = "STUB: not implemented"
	return 0, nil, false, nil
}

//nolint:gosec // G115 - balance is benchmark simulation value, overflow acceptable

// The remaining bytes are random data for padding.

// Creates a new ERC20 contract and optionally writes it to the database. Returns the address of the new ERC20 contract.
func (d *DataGenerator) CreateNewErc20Contract(
	// The number of bytes to allocate for the ERC20 contract data.
	erc20ContractSize int,
	// If true, the ERC20 contract will be immediately written to the database.
	write bool,
) (id int64, address []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Select a random account for a transaction. If an existing account is selected then its ID is guaranteed to be
// less or equal to maxAccountID. If a new account is created, it may have an ID greater than maxAccountID.
func (d *DataGenerator) RandomAccount() (id int64, address []byte, isNew bool, err error) {
	_ = "STUB: not implemented"
	return 0, nil, false, nil
}

// create a new account

// select an existing account at random

// Selects a random account slot for a transaction.
// Uses EVMKeyStorage with addr||slot (AddressLen+SlotLen bytes) for proper storage slot format.
func (d *DataGenerator) randomAccountSlot(accountID int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Selects a random ERC20 contract for a transaction.
func (d *DataGenerator) randomErc20Contract() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Otherwise, select a cold ERC20 contract at random.

// Close the data generator and release any resources.
func (d *DataGenerator) Close() {
	_ = "STUB: not implemented"
	// Specifically release rand, since it's likely to hold a lot of memory.
	return
}

// Get the address of the fee collection account.
func (d *DataGenerator) FeeCollectionAddress() []byte { _ = "STUB: not implemented"; return nil }

// Call this to signal that we have reached the end of a block. This is a signal that it is now safe to use
// recently created accounts as read/write targets.
func (d *DataGenerator) ReportEndOfBlock() { _ = "STUB: not implemented"; return }

// Get the random number generator. Note that the random number generator is not thread safe, and
// so the caller is responsible for ensuring that it is not used concurrently with other calls to the data generator.
func (d *DataGenerator) Rand() *crand.CannedRandom { _ = "STUB: not implemented"; return nil }
