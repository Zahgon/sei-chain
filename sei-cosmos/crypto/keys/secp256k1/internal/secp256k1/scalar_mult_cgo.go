// Copyright 2015 Jeffrey Wilcke, Felix Lange, Gustav Simonsson. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

//go:build !gofuzz && cgo
// +build !gofuzz,cgo

package secp256k1

import (
	"math/big"
)

/*

#include "libsecp256k1/include/secp256k1.h"

extern int secp256k1_ext_scalar_mul(const secp256k1_context* ctx, const unsigned char *point, const unsigned char *scalar);

*/
import "C"

func (BitCurve *BitCurve) ScalarMult(Bx, By *big.Int, scalar []byte) (*big.Int, *big.Int) {
	_ = "STUB: not implemented"
	// Ensure scalar is exactly 32 bytes. We pad always, even if
	// scalar is 32 bytes long, to avoid a timing side channel.
	return nil, nil
}

// NOTE: potential timing issue

// Do the multiplication in C, updating point.

// Unpack the result and clear temporaries.
