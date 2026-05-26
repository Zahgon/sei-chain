package light

import (
	"context"
	"sync"
	"time"

	tmmath "github.com/sei-protocol/sei-chain/sei-tendermint/libs/math"
	"github.com/sei-protocol/sei-chain/sei-tendermint/light/provider"
	"github.com/sei-protocol/sei-chain/sei-tendermint/light/store"
	"github.com/sei-protocol/seilog"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var logger = seilog.NewLogger("tendermint", "light")

type mode byte

const (
	sequential mode = iota + 1
	skipping

	defaultPruningSize = 1000

	// For verifySkipping, we need an algorithm to find what height to check
	// next to see if it has sufficient validator set overlap. The most
	// intuitive method is to take the halfway point i.e. if you trusted block
	// 1 and were not able to verify block 128 then your next try would be 64.
	//
	// However, because this implementation caches all the prior results, instead of always taking halfpoints
	// it is more efficient to re-check cached blocks. Take this simple example. Say
	// you failed to verify 64 but were able to verify block 32. Following a strict half-way policy,
	// you would start over again and try verify to block 128. If this failed
	// then the halfway point between 32 and 128 is 80. But you already have
	// block 64. Instead of requesting and waiting for another block it is far
	// better to try again with block 64. This is of course not directly in the
	// middle. In fact, no matter how the algrorithm plays out, the blocks in
	// cache are always going to be a little less than the halfway point (
	// maximum 1/8 less). To account for this we add a heuristic, bumping the
	// next height to 9/16 instead of 1/2
	verifySkippingNumerator   = 9
	verifySkippingDenominator = 16

	// 10s should cover most of the clients.
	// References:
	// - http://vancouver-webpages.com/time/web.html
	// - https://blog.codinghorror.com/keeping-time-on-the-pc/
	defaultMaxClockDrift = 10 * time.Second

	// 10s is sufficient for most networks.
	defaultMaxBlockLag = 10 * time.Second
)

// Option sets a parameter for the light client.
type Option func(*Client)

// SequentialVerification option configures the light client to sequentially
// check the blocks (every block, in ascending height order). Note this is
// much slower than SkippingVerification, albeit more secure.
func SequentialVerification() Option { _ = "STUB: not implemented"; return *new(Option) }

// SkippingVerification option configures the light client to skip blocks as
// long as {trustLevel} of the old validator set signed the new header. The
// verifySkipping algorithm from the specification is used for finding the minimal
// "trust path".
//
// trustLevel - fraction of the old validator set (in terms of voting power),
// which must sign the new header in order for us to trust it. NOTE this only
// applies to non-adjacent headers. For adjacent headers, sequential
// verification is used.
func SkippingVerification(trustLevel tmmath.Fraction) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// PruningSize option sets the maximum amount of light blocks that the light
// client stores. When Prune() is run, all light blocks that are earlier than
// the h amount of light blocks will be removed from the store.
// Default: 1000. A pruning size of 0 will not prune the light client at all.
func PruningSize(h uint16) Option { _ = "STUB: not implemented"; return *new(Option) }

// MaxClockDrift defines how much new header's time can drift into
// the future relative to the light clients local time. Default: 10s.
func MaxClockDrift(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// MaxBlockLag represents the maximum time difference between the realtime
// that a block is received and the timestamp of that block.
// One can approximate it to the maximum block production time
//
// As an example, say the light client received block B at a time
// 12:05 (this is the real time) and the time on the block
// was 12:00. Then the lag here is 5 minutes.
// Default: 10s
func MaxBlockLag(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// Client represents a light client, connected to a single chain, which gets
// light blocks from a primary provider, verifies them either sequentially or by
// skipping some and stores them in a trusted store (usually, a local FS).
//
// Default verification: SkippingVerification(DefaultTrustLevel)
type Client struct {
	chainID          string
	trustingPeriod   time.Duration // see TrustOptions.Period
	verificationMode mode
	trustLevel       tmmath.Fraction
	maxClockDrift    time.Duration
	maxBlockLag      time.Duration
	blacklistTTL     time.Duration

	// Mutex for locking during changes of the light clients providers
	providerMutex sync.Mutex
	// Primary provider of new headers.
	primary provider.Provider
	// Providers used to "witness" new headers.
	witnesses []provider.Provider

	// Map of witnesses, who have been removed
	// and not allowed to be added back as a provider,
	// to the timadd they were added to the blacklist
	blacklist map[string]time.Time

	// Where trusted light blocks are stored.
	trustedStore store.Store
	// Highest trusted light block from the store (height=H).
	latestTrustedBlock *types.LightBlock

	// See PruningSize option
	pruningSize uint16
}

// NewClient returns a new light client. It returns an error if it fails to
// obtain the light block from the primary, or they are invalid (e.g. trust
// hash does not match with the one from the headers).
//
// Witnesses are providers, which will be used for cross-checking the primary
// provider. At least one witness should be given when skipping verification is
// used (default). A verified header is compared with the headers at same height
// obtained from the specified witnesses. A witness can become a primary iff the
// current primary is unavailable.
//
// See all Option(s) for the additional configuration.
func NewClient(
	ctx context.Context,
	chainID string,
	trustOptions TrustOptions,
	primary provider.Provider,
	witnesses []provider.Provider,
	trustedStore store.Store,
	blacklistTTL time.Duration,
	options ...Option,
) (*Client, error) {
	_ = "STUB: not implemented"

	// Check whether the trusted store already has a trusted block. If so, then create
	// a new client from the trusted store instead of the trust options.
	return nil, nil
}

// Validate the number of witnesses.

// Validate trust options

// Validate trust level.

// Use the trusted hash and height to fetch the first weakly-trusted block
// from the primary provider. Assert that all the witnesses have the same block

// NewClientFromTrustedStore initializes an existing client from the trusted store.
// It does not check that the providers have the same trusted block.
func NewClientFromTrustedStore(
	chainID string,
	trustingPeriod time.Duration,
	primary provider.Provider,
	witnesses []provider.Provider,
	trustedStore store.Store,
	blacklistTTL time.Duration,
	options ...Option) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate trust level.

// Check that the trusted store has at least one block and

// isBlacklisted checks whether provider is black listed
// NOTE: requires a providerMutex lock
func (c *Client) isBlacklisted(p provider.Provider) bool { _ = "STUB: not implemented"; return false }

// If the provider is found, check the TTL

// Remove from blacklist if TTL expired

// restoreTrustedLightBlock loads the latest trusted light block from the store
func (c *Client) restoreTrustedLightBlock() error { _ = "STUB: not implemented"; return nil }

// initializeWithTrustOptions fetches the weakly-trusted light block from
// primary provider, matches it to the trusted hash, and sets it as the
// lastTrustedBlock. It then asserts that all witnesses have the same light block.
func (c *Client) initializeWithTrustOptions(ctx context.Context, options TrustOptions) error {
	_ = "STUB: not implemented"
	// 1) Fetch and verify the light block. Note that we do not verify the time of the first block
	return nil
}

// 2) Assert that the hashes match

// 3) Ensure that +2/3 of validators signed correctly. This also sanity checks that the
// chain ID is the same.

// 4) Cross-verify with witnesses to ensure everybody has the same state.

// 5) Persist both of them and continue.

// TrustedLightBlock returns a trusted light block at the given height (0 - the latest).
//
// It returns an error if:
//   - there are some issues with the trusted store, although that should not
//     happen normally;
//   - negative height is passed;
//   - header has not been verified yet and is therefore not in the store
//
// Safe for concurrent use by multiple goroutines.
func (c *Client) TrustedLightBlock(height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) compareWithLatestHeight(height int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Update attempts to advance the state by downloading the latest light
// block and verifying it. It returns a new light block on a successful
// update. Otherwise, it returns nil (plus an error, if any).
func (c *Client) Update(ctx context.Context, now time.Time) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no light blocks yet => wait

// If there is a new light block then verify it

// else return the latestTrustedBlock

// VerifyLightBlockAtHeight fetches the light block at the given height
// and verifies it. It returns the block immediately if it exists in
// the trustedStore (no verification is needed).
//
// height must be > 0.
//
// It returns provider.ErrlightBlockNotFound if light block is not found by
// primary.
//
// It will replace the primary provider if an error from a request to the provider occurs
func (c *Client) VerifyLightBlockAtHeight(ctx context.Context, height int64, now time.Time) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the light block is already verified.

// Return already trusted light block

// Request the light block from primary

// VerifyHeader verifies a new header against the trusted state. It returns
// immediately if newHeader exists in trustedStore (no verification is
// needed). Else it performs one of the two types of verification:
//
// SequentialVerification: verifies that 2/3 of the trusted validator set has
// signed the new header. If the headers are not adjacent, **all** intermediate
// headers will be requested. Intermediate headers are not saved to database.
//
// SkippingVerification(trustLevel): verifies that {trustLevel} of the trusted
// validator set has signed the new header. If it's not the case and the
// headers are not adjacent, verifySkipping is performed and necessary (not all)
// intermediate headers will be requested. See the specification for details.
// Intermediate headers are not saved to database.
// https://github.com/tendermint/tendermint/blob/master/spec/light-client/README.md
//
// If the header, which is older than the currently trusted header, is
// requested and the light client does not have it, VerifyHeader will perform:
//
//	a) verifySkipping verification if nearest trusted header is found & not expired
//	b) backwards verification in all other cases
//
// It returns ErrOldHeaderExpired if the latest trusted header expired.
//
// If the primary provides an invalid header (ErrInvalidHeader), it is rejected
// and replaced by another provider until all are exhausted.
//
// If, at any moment, a LightBlock is not found by the primary provider as part of
// verification then the provider will be replaced by another and the process will
// restart.
func (c *Client) VerifyHeader(ctx context.Context, newHeader *types.Header, now time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if newHeader already verified.

// Make sure it's the same header.

// Request the header and the vals.

func (c *Client) verifyLightBlock(ctx context.Context, newLightBlock *types.LightBlock, now time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// Verifying forwards

// Verifying backwards

// Verifying between first and last trusted light block. In this situation
// we find the closest block prior to the target height then perform
// verification forwards.

// Once verified, save and return

// see VerifyHeader
func (c *Client) verifySequential(
	ctx context.Context,
	trustedBlock *types.LightBlock,
	newLightBlock *types.LightBlock,
	now time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// 1) Fetch interim light block if needed.
// last light block

// intermediate light blocks

// 2) Verify them

// If the target header is invalid, return immediately.

// If some intermediate header is invalid, remove the primary and try again.

// attempt to verify header again

// 3) Update verifiedBlock

// 4) Add verifiedBlock to trace

// Compare header with the witnesses to ensure it's not a fork.
// More witnesses we have, more chance to notice one.
//
// CORRECTNESS ASSUMPTION: there's at least 1 correct full node
// (primary or one of the witnesses).

// see VerifyHeader
//
// verifySkipping finds the middle light block between a trusted and new light block,
// reiterating the action until it verifies a light block. A cache of light blocks
// requested from source is kept such that when a verification is made, and the
// light client tries again to verify the new light block in the middle, the light
// client does not need to ask for all the same light blocks again.
//
// If this function errors, it should always wrap it in a `ErrVerifcationFailed`
// struct so that the calling function can determine where it failed and handle
// it accordingly.
func (c *Client) verifySkipping(
	ctx context.Context,
	source provider.Provider,
	trustedBlock *types.LightBlock,
	newLightBlock *types.LightBlock,
	now time.Time) ([]*types.LightBlock, error) {
	_ = "STUB: not implemented"

	// The block cache is ordered in height from highest to lowest. We start
	// with the newLightBlock and for any height requested in between we add
	// it.
	return nil, nil
}

// Verify the untrusted header. This function is equivalent to
// ValidAndVerified in the spec

// If we have verified the last header then depth will be 0 and we
// can return a success along with the trace of intermediate headers

// If not, update the lower bound to the previous upper bound

// Remove the light block at the lower bound in the header cache - it will no longer be needed

// Reset the cache depth so that we start from the upper bound again

// add verifiedBlock to the trace

// the light block current passed validation, but the validator
// set is too different to verify it. We keep the block because it
// may become valuable later on.
//
// If we have reached the end of the cache we need to request a
// completely new block else we recycle a previously requested one.
// In both cases we are taking a block with a closer height to the
// previously verified one in the hope that it has a better chance
// of having a similar validator set

// schedule what the next height we need to fetch is

// for any verification error we abort the operation and return the error

// schedule works out the next height to attempt sequential verification
func (c *Client) schedule(lastVerifiedHeight, lastFailedHeight int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// verifySkippingAgainstPrimary does verifySkipping plus it compares new header with
// witnesses and replaces primary if it sends the light client an invalid header
func (c *Client) verifySkippingAgainstPrimary(
	ctx context.Context,
	trustedBlock *types.LightBlock,
	newLightBlock *types.LightBlock,
	now time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// Success! Now compare the header with the witnesses to ensure it's not a fork.
// More witnesses we have, more chance to notice one.
//
// CORRECTNESS ASSUMPTION: there's at least 1 correct full node
// (primary or one of the witnesses).

// all errors from verify skipping should be `ErrVerificationFailed`
// if it's not we just return the error directly

// Verification returned an invalid header

// If it was the target header, return immediately.

// If some intermediate header is invalid, remove the primary and try
// again.

// An intermediate header expired. We can no longer validate it as there is
// no longer the ability to punish invalid blocks as evidence of misbehavior

// This happens if there was a problem in finding the next block or a
// context was canceled.

// if we've reached here we're attempting to retry verification with a
// different provider

// attempt to verify the header again from the trusted block

// LastTrustedHeight returns a last trusted height. -1 and nil are returned if
// there are no trusted headers.
//
// Safe for concurrent use by multiple goroutines.
func (c *Client) LastTrustedHeight() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// FirstTrustedHeight returns a first trusted height. -1 and nil are returned if
// there are no trusted headers.
//
// Safe for concurrent use by multiple goroutines.
func (c *Client) FirstTrustedHeight() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// ChainID returns the chain ID the light client was configured with.
//
// Safe for concurrent use by multiple goroutines.
func (c *Client) ChainID() string {
	_ = "STUB: not implemented"

	// Primary returns the primary provider.
	//
	// NOTE: provider may be not safe for concurrent access.
	return ""
}

func (c *Client) Primary() provider.Provider {
	_ = "STUB: not implemented"
	return *new(provider.Provider)
}

// Witnesses returns the witness providers.
//
// NOTE: providers may be not safe for concurrent access.
func (c *Client) Witnesses() []provider.Provider { _ = "STUB: not implemented"; return nil }

// BlacklistedWitnessIDS returns the blacklisted witness IDs.
//
// NOTE: providers may be not safe for concurrent access.
func (c *Client) BlacklistedWitnessIDs() []string { _ = "STUB: not implemented"; return nil }

// AddProvider adds a providers to the light clients set
//
// NOTE: The light client does not check for uniqueness
func (c *Client) AddProvider(p provider.Provider) { _ = "STUB: not implemented"; return }

// If the provider is blacklisted, don't add it

// Cleanup removes all the data (headers and validator sets) stored. Note: the
// client must be stopped at this point.
func (c *Client) Cleanup() error { _ = "STUB: not implemented"; return nil }

func (c *Client) updateTrustedLightBlock(l *types.LightBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// backwards verification (see VerifyHeaderBackwards func in the spec) verifies
// headers before a trusted header. If a sent header is invalid the primary is
// replaced with another provider and the operation is repeated.
func (c *Client) backwards(
	ctx context.Context,
	trustedHeader *types.Header,
	newHeader *types.Header) error {
	_ = "STUB: not implemented"
	return nil
}

// verification has failed

// the client tries to see if it can get a witness to continue with the request

// before continuing we must check that they have the same target header to validate

// return the original error

// try again with the new primary

// lightBlockFromPrimary retrieves the lightBlock from the primary provider
// at the specified height. This method also handles provider behavior as follows:
//
//  1. If the provider does not respond or does not have the block, it tries again
//     with a different provider
//  2. If all providers return the same error, the light client forwards the error to
//     where the initial request came from
//  3. If the provider provides an invalid light block, is deemed unreliable or returns
//     any other error, the primary is permanently dropped and is replaced by a witness.
func (c *Client) lightBlockFromPrimary(ctx context.Context, height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Everything went smoothly. We reset the lightBlockRequests and return the light block

// catch canceled contexts or deadlines

// we find a new witness to replace the primary

// The light client has most likely received either provider.ErrUnreliableProvider or provider.ErrBadLightBlock
// These errors mean that the light client should drop the primary and try with another provider instead

func (c *Client) getLightBlock(ctx context.Context, p provider.Provider, height int64) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// addWitnessToBlacklist adds a witness to the blacklist
// NOTE: requires a providerMutex lock
func (c *Client) addWitnessesToBlacklist(providers []provider.Provider) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) findIndexForWitness(ID types.NodeID) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// RemoveProviderByID removes a witness from the light client.
func (c *Client) RemoveProviderByID(ID types.NodeID) error { _ = "STUB: not implemented"; return nil }

// NOTE: requires a providerMutex lock
func (c *Client) removeWitnesses(indexes []int) error { _ = "STUB: not implemented"; return nil }

// we need to make sure that we remove witnesses by index in the reverse
// order so as to not affect the indexes themselves

type witnessResponse struct {
	lb           *types.LightBlock
	witnessIndex int
	err          error
}

// findNewPrimary concurrently sends a light block request, promoting the first witness to return
// a valid light block as the new primary. The remove option indicates whether the primary should be
// entire removed or just appended to the back of the witnesses list. This method also handles witness
// errors. If no witness is available, it returns the last error of the witness.
func (c *Client) findNewPrimary(ctx context.Context, height int64, remove bool) (*types.LightBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// send out a light block request to all witnesses

// process all the responses as they come in

// success! We have found a new primary

// cancel all remaining requests to other witnesses

// wait for all goroutines to finish

// if we are not intending on removing the primary then append the old primary to the end of the witness slice

// promote respondent as the new primary

// add promoted witness to the list of witnesses to be removed

// remove witnesses marked as bad (the client must do this before we alter the witness slice and change the indexes
// of witnesses). Removal is done in descending order

// return the light block that new primary responded with

// catch canceled contexts or deadlines

// process benign errors by logging them only

// process malevolent errors like ErrUnreliableProvider and ErrBadLightBlock by removing the witness

// compareFirstLightBlockWithWitnesses concurrently compares light block l with all witnesses. If any
// witness reports a different header than h, the function returns an error.
func (c *Client) compareFirstLightBlockWithWitnesses(ctx context.Context, l *types.LightBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// handle errors from the header comparisons as they come in

// If witness sent us an invalid header, then remove it

// check for canceled contexts or deadlines

// the witness either didn't respond or didn't have the block. We ignore it.

// remove all witnesses that misbehaved

// providerShouldBeRemoved analyzes the nature of the error and whether the provider
// should be removed from the light clients set
func (c *Client) providerShouldBeRemoved(err error) bool { _ = "STUB: not implemented"; return false }

func (c *Client) Status(ctx context.Context) *types.LightClientInfo {
	_ = "STUB: not implemented"
	return nil
}

// If primary is in witness list we do not want to count it twice in the number of peers

// The caller of /status can deduce this from the two variables above
// Having a boolean flag improves readbility
