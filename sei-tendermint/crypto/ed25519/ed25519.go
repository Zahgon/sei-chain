package ed25519

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519/extra/cache"

	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type ErrBadSig struct {
	Idx int // Index of the first invalid signature.
}

func (e ErrBadSig) Error() string { _ = "STUB: not implemented"; return "" }

// cacheSize is the number of public keys that will be cached in
// an expanded format for repeated signature verification.
//
// TODO/perf: Either this should exclude single verification, or be
// tuned to `> validatorSize + maxTxnsPerBlock` to avoid cache
// thrashing.
const cacheSize = 4096

// curve25519-voi's Ed25519 implementation supports configurable
// verification behavior, and tendermint uses the ZIP-215 verification
// semantics.
var verifyOptions = ed25519.VerifyOptionsZIP_215
var cachingVerifier = cache.NewVerifier(cache.NewLRUCache(cacheSize))

// SecretKey represents a secret key in the Ed25519 signature scheme.
type SecretKey struct {
	// When using the key make sure to use runtime.KeepAlive, so that key is not zeroized midway.
	// This is a pointer to avoid copying the secret to stack when used.
	// This is a pointer to pointer, so that runtime.AddCleanup can actually work:
	// AddCleanup requires the referenced pointer to be unreachable, even from the cleanup function.
	// This is a closure returning pointer to pointer , so that secret is not extractable via golang reflection,
	// because reflection is not able to call closures stored in private fields of types in other modules.
	// this protects us from reflection-based traversing (like default stringer implementation)
	// We cannot reuse the closure interface as the outer pointer (i.e. reduce it to func() *[...]byte),
	// because closure is not a pointer in go.
	// The cost of dereferencing multiple times is assumed to be dwarfed by the crypto operations.
	// Secret is not comparable and that's intentional. Compare the public keys instead.
	key func() **[ed25519.PrivateKeySize]byte
}

// WARNING: this function should only be used when persisting the private key.
// WARNING: caller is responsible for zeroizing the returned slice.
func (k SecretKey) SecretBytes() []byte { _ = "STUB: not implemented"; return nil }

// SecretKeyFromSecretBytes constructs a secret key from a raw secret material.
// WARNING: this function zeroes the content of the input slice.
func SecretKeyFromSecretBytes(b []byte) (SecretKey, error) {
	_ = "STUB: not implemented"
	return *new(SecretKey), nil
}

// Zero the memory to avoid leaking the secret.

// Zero the input slice to avoid leaking the secret.

// TestSecretKey generates a testonly secret key.
func TestSecretKey(seed []byte) SecretKey { _ = "STUB: not implemented"; return *new(SecretKey) }

// GenerateSecretKey generates a new secret key using a cryptographically secure random number generator.
func GenerateSecretKey() SecretKey { _ = "STUB: not implemented"; return *new(SecretKey) }

// rand.Read is documented to never return an error.

// Generated key is always valid.

// Zeroize the seed after generation.

// Public returns the public key corresponding to the secret key.
func (k SecretKey) Public() PublicKey { _ = "STUB: not implemented"; return *new(PublicKey) }

// PublicKey represents a public key in the Ed25519 signature scheme.
type PublicKey struct {
	utils.ReadOnly
	key [ed25519.PublicKeySize]byte
}

// Signature represents a signature in the Ed25519 signature scheme.
type Signature struct {
	utils.ReadOnly
	sig [ed25519.SignatureSize]byte
}

// Sign signs a message using the secret key.
func (k SecretKey) Sign(message []byte) Signature {
	_ = "STUB: not implemented"
	return *new(Signature)
}

// Domain separation tag.
type Tag struct{ tag string }

func NewTag(tag string) (Tag, error) { _ = "STUB: not implemented"; return *new(Tag), nil }

// SignWithTag signs a message with a domain separation tag.
// It is safe to assume that signatures for messages with different tags do not collide.
// It is also safe to assume that Sign() signatures do not collide with SignWithTag() signatures.
// It is secure to use the same secret key for signing with both Sign() and SignWithTag() [https://datatracker.ietf.org/doc/html/rfc8032#section-8.6].
func (k SecretKey) SignWithTag(tag Tag, msg []byte) Signature {
	_ = "STUB: not implemented"
	return *new(Signature)
}

// Returns no error if opts.Context is of correct size.

// Compare defines a total order on public keys.
func (k PublicKey) Compare(other PublicKey) int { _ = "STUB: not implemented"; return 0 }

// Verify verifies a signature.
func (k PublicKey) Verify(msg []byte, sig Signature) error { _ = "STUB: not implemented"; return nil }

// Verify verifies a signature, given domain separation tag.
func (k PublicKey) VerifyWithTag(tag Tag, msg []byte, sig Signature) error {
	_ = "STUB: not implemented"
	return nil
}

// BatchVerifier implements batch verification for ed25519.
type BatchVerifier struct {
	inner *ed25519.BatchVerifier
}

func NewBatchVerifier() *BatchVerifier { _ = "STUB: not implemented"; return nil }

func (b *BatchVerifier) Add(key PublicKey, msg []byte, sig Signature) {
	_ = "STUB: not implemented"
	return
}

func (b *BatchVerifier) AddWithTag(key PublicKey, tag Tag, msg []byte, sig Signature) {
	_ = "STUB: not implemented"
	return
}

// Verify verifies the batched signatures using OS entropy.
// If any signature is invalid, returns ErrBadSig with an index
// of the first invalid signature.
func (b *BatchVerifier) Verify() error { _ = "STUB: not implemented"; return nil }
