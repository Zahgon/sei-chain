package migration

import (
	ics23 "github.com/confio/ics23/go"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	db "github.com/tendermint/tm-db"
)

// Route binds a set of module/store names to the database accessors
// (reader, writer, and optionally iterator and proof builders) that
// should be used to access them. A ModuleRouter dispatches reads,
// writes, iteration and proof requests to the matching Route.
type Route struct {
	// The module names to route to this destination. Guaranteed to
	// contain no duplicates by NewRoute.
	modules []string
	// For reading values from the database.
	reader DBReader
	// For writing values to the database.
	writer DBWriter
	// For getting an iterator over a range of keys in a store. If nil, the route does not support iteration.
	iteratorBuilder DBIteratorBuilder
	// For building a proof of the value for a key in a store. If nil, the route does not support proofs.
	proofBuilder DBProofBuilder
}

// NewRoute creates a new Route.
//
// modules may be empty (the route will simply receive no traffic), but
// each name listed must be unique: a duplicate is rejected as a
// misconfiguration.
func NewRoute(
	// For reading values from the database.
	reader DBReader,
	// For writing values to the database.
	writer DBWriter,
	// For getting an iterator over a range of keys in a store. If nil, the route does not support iteration.
	iteratorBuilder DBIteratorBuilder,
	// For building a proof of the value for a key in a store. If nil, the route does not support proofs.
	proofBuilder DBProofBuilder,
	// The module names to route to this destination. Must not contain
	// duplicates.
	modules ...string,
) (*Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Defensive copy: callers passing a slice with `mods...` should not
// be able to mutate our internal state after construction.

var _ Router = (*ModuleRouter)(nil)

// ModuleRouter routes reads and writes between any number of databases
// based on the module/store name they target. Each module name must be
// registered with exactly one Route; reads or writes to an unregistered
// module return an error.
type ModuleRouter struct {
	// The routes managed by this router, in the order they were
	// registered. ApplyChangeSets dispatches to each route's writer
	// sequentially in registration order.
	routes []*Route

	// Lookup from module/store name to the route that owns it.
	moduleToRoute map[string]*Route
}

// NewModuleRouter creates a new ModuleRouter from one or more Routes.
//
// At least one Route must be provided. No module name may appear in more
// than one Route. Data targeting a module that is not registered with any
// Route returns an error.
//
// This is intentionally fragile to misconfiguration: it is important to
// very specifically specify which modules belong to which database.
func NewModuleRouter(routes ...*Route) (*ModuleRouter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Defensive copy: callers passing a slice with `routes...` should not
// be able to mutate our internal state after construction.

// ApplyChangeSets splits changesets across the registered routes based
// on the module name of each changeset and applies them sequentially in
// registration order. If any changeset targets a module that is not
// registered with any route, no writes are performed and an error is
// returned.
//
// Non-atomic across routes; atomicity must be ensured by the caller.
func (m *ModuleRouter) ApplyChangeSets(changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

// Read returns the value for key in store, dispatching to the route
// that owns store. Returns an error if store is not registered with any
// route.
func (m *ModuleRouter) Read(store string, key []byte) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *ModuleRouter) GetProof(store string, key []byte) (*ics23.CommitmentProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ModuleRouter) Iterator(store string, start []byte, end []byte, ascending bool) (db.Iterator, error) {
	_ = "STUB: not implemented"
	return *new(db.Iterator), nil
}
