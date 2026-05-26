// Copyright 2022 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package evmrpc

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const JwtExpiryTimeout = 60 * time.Second

type jwtHandler struct {
	keyFunc func(token *jwt.Token) (interface{}, error)
	next    http.Handler
}

// newJWTHandler creates a http.Handler with jwt authentication support.
func newJWTHandler(secret []byte, next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// ServeHTTP implements http.Handler
func (handler *jwtHandler) ServeHTTP(out http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// We explicitly set only HS256 allowed, and also disables the
// claim-check: the RegisteredClaims internally requires 'iat' to
// be no later than 'now', but we allow for a bit of drift.

// optional
