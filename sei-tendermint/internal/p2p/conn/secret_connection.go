package conn

import (
	"context"
	"crypto/cipher"
	"errors"
	"net/netip"

	"golang.org/x/crypto/chacha20poly1305"
)

var errAEAD = errors.New("decoding failed")
var errDH = errors.New("DH secret failure")

// 4 + 1024 == 1028 total frame size
const (
	dataSizeLen   = 4
	dataSizeMax   = 1024
	frameSize     = dataSizeLen + dataSizeMax
	aeadOverhead  = chacha20poly1305.Overhead
	aeadNonceSize = chacha20poly1305.NonceSize

	labelEphemeralLowerPublicKey = "EPHEMERAL_LOWER_PUBLIC_KEY"
	labelEphemeralUpperPublicKey = "EPHEMERAL_UPPER_PUBLIC_KEY"
	labelDHSecret                = "DH_SECRET"
	labelSecretConnectionMac     = "SECRET_CONNECTION_MAC"
)

type asyncMutex[T any] struct {
	mu chan struct{}
	v  T
}

func newAsyncMutex[T any](v T) asyncMutex[T] { _ = "STUB: not implemented"; return nil }

func (m *asyncMutex[T]) Lock(ctx context.Context, yield func(T) error) error {
	_ = "STUB: not implemented"
	return nil
}

var secretConnKeyAndChallengeGen = []byte("TENDERMINT_SECRET_CONNECTION_KEY_AND_CHALLENGE_GEN")

type sendState struct {
	cipher cipher.AEAD
	frame  []byte
	data   []byte
	nonce  uint64
}

type recvState struct {
	cipher cipher.AEAD
	frame  []byte
	data   []byte
	nonce  uint64
}

func newSendState(cipher cipher.AEAD) *sendState { _ = "STUB: not implemented"; return nil }

func newRecvState(cipher cipher.AEAD) *recvState { _ = "STUB: not implemented"; return nil }

var _ Conn = (*SecretConnection)(nil)

type Challenge [32]byte

// SecretConnection implements Conn.
// It is an implementation of the STS protocol.
// See https://github.com/tendermint/tendermint/blob/0.1/docs/sts-final.pdf for
// details on the protocol.
//
// Consumers of the SecretConnection are responsible for authenticating
// the remote peer's pubkey against known information, like a nodeID.
// Otherwise they are vulnerable to MITM.
// (TODO(ismail): see also https://github.com/tendermint/tendermint/issues/3010)
type SecretConnection struct {
	conn      Conn
	challenge Challenge
	recvState asyncMutex[*recvState]
	sendState asyncMutex[*sendState]
}

func (sc *SecretConnection) Challenge() Challenge {
	_ = "STUB: not implemented"
	return *new(Challenge)
}

func newSecretConnection(conn Conn, loc ephSecret, rem ephPublic) (*SecretConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generate the secret used for receiving, sending, challenge via HKDF-SHA2
// on the transcript state (which itself also uses HKDF-SHA2 to derive a key
// from the dhSecret).

// MakeSecretConnection performs handshake and returns an encrypted SecretConnection.
// To authenticate the secret connection, you need to sign the Challenge() and exchange the signatures
// with the peer. See docs/sts-final.pdf for more information.
func MakeSecretConnection(ctx context.Context, conn Conn) (*SecretConnection, error) {
	_ = "STUB: not implemented"
	// Write local ephemeral pubkey and receive one too.
	// NOTE: every 32-byte string is accepted as a Curve25519 public key (see
	// DJB's Curve25519 paper: http://cr.yp.to/ecdh/curve25519-20060209.pdf)
	return nil, nil
}

// Generate ephemeral key for perfect forward secrecy.

// Writes encrypted frames of `totalFrameSize + aeadSizeOverhead`.
func (sc *SecretConnection) Write(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SecretConnection) flush(ctx context.Context, sendState *sendState) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // data length bounded by frame size

// We use a predeclared stack-allocated buffer, to prevent Seal from doing heap allocation.
// I'm not sure whether this optimization is needed though.

// Zeroize the frame to avoid resending data from the previous frame.
// Security-wise it doesn't make any difference, it is here just to avoid people raising concerns.

func (sc *SecretConnection) Read(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Implements Conn
func (sc *SecretConnection) Flush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (sc *SecretConnection) LocalAddr() netip.AddrPort {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort)
}
func (sc *SecretConnection) RemoteAddr() netip.AddrPort {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort)
}
func (sc *SecretConnection) Close() { _ = "STUB: not implemented"; return }

type ephPublic [32]byte

type ephSecret struct {
	secret [32]byte
	public ephPublic
}

func genEphKey() ephSecret {
	_ = "STUB: not implemented"
	// TODO: Probably not a problem but ask Tony: different from the rust implementation (uses x25519-dalek),
	// we do not "clamp" the private key scalar:
	// see: https://github.com/dalek-cryptography/x25519-dalek/blob/34676d336049df2bba763cc076a75e47ae1f170f/src/x25519.rs#L56-L74
	return *new(ephSecret)
}

type dhSecret [32]byte
type aeadSecret [chacha20poly1305.KeySize]byte

type aeadSecrets struct {
	send aeadSecret
	recv aeadSecret
}

func (s aeadSecret) Cipher() cipher.AEAD {
	_ = "STUB: not implemented"
	// Never returns an error on input of correct size.
	return *new(cipher.AEAD)
}

func (s dhSecret) AeadSecrets(locIsLeast bool) aeadSecrets {
	_ = "STUB: not implemented"
	return *new(aeadSecrets)
}

// hkdf reader never returns an error.

// computeDHSecret computes a Diffie-Hellman shared secret key
// from our own local private key and the other's public key.
func (s ephSecret) DhSecret(remPubKey ephPublic) (dhSecret, error) {
	_ = "STUB: not implemented"
	return *new(dhSecret), nil
}
