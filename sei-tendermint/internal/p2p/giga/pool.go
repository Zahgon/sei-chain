package giga

import (
	"context"

	"github.com/google/btree"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

type poolEntry[V any] struct {
	idx     uint64
	deleted utils.AtomicSend[bool]
	val     V
}

func (e *poolEntry[V]) Run(ctx context.Context, task func(context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

type poolInner[K comparable, V any] struct {
	nextIdx uint64
	byIdx   *btree.BTreeG[*poolEntry[V]]
	byKey   map[K]*poolEntry[V]
}

type Pool[K comparable, V any] struct {
	inner utils.Watch[*poolInner[K, V]]
}

func NewPool[K comparable, V any]() *Pool[K, V] { _ = "STUB: not implemented"; return nil }

func (p *Pool[K, V]) InsertAndRun(ctx context.Context, key K, val V, task func(context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Pool[K, V]) RunForEach(ctx context.Context, task func(context.Context, V) error) error {
	_ = "STUB: not implemented"
	return nil
}
