package utils

import (
	mrand "math/rand"
	"sync"
)

const (
	strChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" // 62 characters
)

type Rand struct {
	sync.Mutex
	rand *mrand.Rand
}

var grand *Rand

func init() {
	grand = NewRand()
	grand.init()
}

func NewRand() *Rand { _ = "STUB: not implemented"; return nil }

func (r *Rand) init() {
	bz := cRandBytes(8)
	var seed uint64
	for i := 0; i < 8; i++ {
		seed |= uint64(bz[i])
		seed <<= 8
	}
	r.reset(int64(seed)) //#nosec G115 -- intentional conversion; full uint64 entropy is desired, sign bit is irrelevant for seeding
}

func (r *Rand) reset(seed int64) { _ = "STUB: not implemented"; return }

//nolint:gosec // G404: seeded from crypto/rand, used for non-security purposes

func (r *Rand) Int() int { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Int63() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Rand) Str(length int) string { _ = "STUB: not implemented"; return "" }

// rightmost 6 bits
// only 62 characters in strChars

func Int() int { _ = "STUB: not implemented"; return 0 }

func cRandBytes(numBytes int) []byte { _ = "STUB: not implemented"; return nil }
