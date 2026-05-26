package state

func (s *DBImpl) AddRefund(gas uint64) { _ = "STUB: not implemented"; return }

// Copied from go-ethereum as-is
// SubRefund removes gas from the refund counter.
// This method will panic if the refund counter goes below zero
func (s *DBImpl) SubRefund(gas uint64) { _ = "STUB: not implemented"; return }

func (s *DBImpl) GetRefund() uint64 { _ = "STUB: not implemented"; return 0 }
