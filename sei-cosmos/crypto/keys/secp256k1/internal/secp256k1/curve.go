// Copyright 2010 The Go Authors. All rights reserved.
// Copyright 2011 ThePiachu. All rights reserved.
// Copyright 2015 Jeffrey Wilcke, Felix Lange, Gustav Simonsson. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// * Redistributions of source code must retain the above copyright
//   notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
//   copyright notice, this list of conditions and the following disclaimer
//   in the documentation and/or other materials provided with the
//   distribution.
// * Neither the name of Google Inc. nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
// * The name of ThePiachu may not be used to endorse or promote products
//   derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package secp256k1

import (
	"crypto/elliptic"
	"math/big"
)

const (
	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)

// readBits encodes the absolute value of bigint as big-endian bytes. Callers
// must ensure that buf has enough space. If buf is too short the result will
// be incomplete.
func readBits(bigint *big.Int, buf []byte) { _ = "STUB: not implemented"; return }

// This code is from https://github.com/ThePiachu/GoBit and implements
// several Koblitz elliptic curves over prime fields.
//
// The curve methods, internally, on Jacobian coordinates. For a given
// (x, y) position on the curve, the Jacobian coordinates are (x1, y1,
// z1) where x = x1/z1² and y = y1/z1³. The greatest speedups come
// when the whole calculation can be performed within the transform
// (as in ScalarMult and ScalarBaseMult). But even for Add and Double,
// it's faster to apply and reverse the transform than to operate in
// affine coordinates.

// A BitCurve represents a Koblitz Curve with a=0.
// See http://www.hyperelliptic.org/EFD/g1p/auto-shortw.html
type BitCurve struct {
	P       *big.Int // the order of the underlying field
	N       *big.Int // the order of the base point
	B       *big.Int // the constant of the BitCurve equation
	Gx, Gy  *big.Int // (x,y) of the base point
	BitSize int      // the size of the underlying field
}

func (BitCurve *BitCurve) Params() *elliptic.CurveParams { _ = "STUB: not implemented"; return nil }

// IsOnCurve returns true if the given (x,y) lies on the BitCurve.
func (BitCurve *BitCurve) IsOnCurve(x, y *big.Int) bool {
	_ = "STUB: not implemented"
	// y² = x³ + b
	return false
}

//y²
//y²%P

//x²
//x³

//x³+B
//(x³+B)%P

// TODO: double check if the function is okay
// affineFromJacobian reverses the Jacobian transform. See the comment at the
// top of the file.
func (BitCurve *BitCurve) affineFromJacobian(x, y, z *big.Int) (xOut, yOut *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add returns the sum of (x1,y1) and (x2,y2)
func (BitCurve *BitCurve) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int) {
	_ = "STUB: not implemented"
	// If one point is at infinity, return the other point.
	// Adding the point at infinity to any point will preserve the other point.
	return nil, nil
}

// addJacobian takes two points in Jacobian coordinates, (x1, y1, z1) and
// (x2, y2, z2) and returns their sum, also in Jacobian form.
func (BitCurve *BitCurve) addJacobian(x1, y1, z1, x2, y2, z2 *big.Int) (*big.Int, *big.Int, *big.Int) {
	_ = "STUB: not implemented"
	// See http://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-0.html#addition-add-2007-bl
	return nil, nil, nil
}

// Double returns 2*(x,y)
func (BitCurve *BitCurve) Double(x1, y1 *big.Int) (*big.Int, *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// doubleJacobian takes a point in Jacobian coordinates, (x, y, z), and
// returns its double, also in Jacobian form.
func (BitCurve *BitCurve) doubleJacobian(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
	_ = "STUB: not implemented"
	// See http://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-0.html#doubling-dbl-2009-l
	return nil, nil, nil
}

//X1²
//Y1²
//B²

//X1+B
//(X1+B)²
//(X1+B)²-A
//(X1+B)²-A-C
//2*((X1+B)²-A-C)

//3*A
//E²

//2*D
//F-2*D

//D-X3
//E*(D-X3)
//E*(D-X3)-8*C

//Y1*Z1
//3*Y1*Z1

// ScalarBaseMult returns k*G, where G is the base point of the group and k is
// an integer in big-endian form.
func (BitCurve *BitCurve) ScalarBaseMult(k []byte) (*big.Int, *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marshal converts a point into the form specified in section 4.3.6 of ANSI
// X9.62.
func (BitCurve *BitCurve) Marshal(x, y *big.Int) []byte { _ = "STUB: not implemented"; return nil }

// uncompressed point flag

// Unmarshal converts a point, serialised by Marshal, into an x, y pair. On
// error, x = nil.
func (BitCurve *BitCurve) Unmarshal(data []byte) (x, y *big.Int) {
	_ = "STUB: not implemented"
	return nil, nil
}

// uncompressed form

var theCurve = new(BitCurve)

func init() {
	// See SEC 2 section 2.7.1
	// curve parameters taken from:
	// http://www.secg.org/sec2-v2.pdf
	theCurve.P, _ = new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 0)
	theCurve.N, _ = new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 0)
	theCurve.B, _ = new(big.Int).SetString("0x0000000000000000000000000000000000000000000000000000000000000007", 0)
	theCurve.Gx, _ = new(big.Int).SetString("0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 0)
	theCurve.Gy, _ = new(big.Int).SetString("0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 0)
	theCurve.BitSize = 256
}

// S256 returns a BitCurve which implements secp256k1.
func S256() *BitCurve { _ = "STUB: not implemented"; return nil }
