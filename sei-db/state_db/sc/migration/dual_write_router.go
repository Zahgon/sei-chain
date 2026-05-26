package migration

import (
	ics23 "github.com/confio/ics23/go"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	db "github.com/tendermint/tm-db"
)

var _ Router = (*TestOnlyDualWriteRouter)(nil)

// A router that dual-writes traffic, sending each batch of changesets to both backends. Read
// requests, requests for proofs, and requests for iteration are not dual-written, and are instead
// served exclusively by the primary backend.
//
// CRITICAL: this is a test-only router and should never be deployed to production machines.
type TestOnlyDualWriteRouter struct {
	primary   *Route
	secondary DBWriter
}

// Create a new test-only dual-write router.
//
// CRITICAL: this is a test-only router and should never be deployed to production machines.
func NewTestOnlyDualWriteRouter(
	// Read, proof, and iteration traffic is served by this route, and writes are also sent here.
	// Module names associated with this route are ignored; this route forwards all regardless of the module names.
	primary *Route,
	// Write traffic is dual-written and also sent here. Reads, proofs, and iteration are not sent here.
	secondary DBWriter,
) (*TestOnlyDualWriteRouter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TestOnlyDualWriteRouter) ApplyChangeSets(changesets []*proto.NamedChangeSet) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestOnlyDualWriteRouter) GetProof(store string, key []byte) (*ics23.CommitmentProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TestOnlyDualWriteRouter) Iterator(
	store string,
	start []byte,
	end []byte,
	ascending bool,
) (db.Iterator, error) {
	_ = "STUB: not implemented"
	return *new(db.Iterator), nil
}

func (t *TestOnlyDualWriteRouter) Read(store string, key []byte) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// BuildRoute returns a Route that dispatches the given module names to
// this DualWriteRouter. Reads, writes, iteration and proof requests
// for those modules will all flow through this dual-write router.
//
// Module names must be unique; NewRoute's validation rules apply. The
// returned Route may be passed to NewModuleRouter alongside other
// Routes to compose multi-database setups.
func (t *TestOnlyDualWriteRouter) BuildRoute(moduleNames ...string) (*Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
