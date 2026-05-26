package wal

type ErrBadOffset struct{ error }

type logView struct {
	headPath string
	firstIdx int
	nextIdx  int
}

func (v *logView) tailPath(idx int) string { _ = "STUB: not implemented"; return "" }

func (v *logView) PathByOffset(fileOffset int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func loadLogView(headPath string) (*logView, error) { _ = "STUB: not implemented"; return nil, nil }

// Find the idx range.

func (v *logView) TailSize() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (v *logView) Rotate(cfg *Config) error {
	_ = "STUB: not implemented"
	// Move head to tail.
	// WARNING: Rename is atomic, but not crash safe.
	// We don't need to sync directory - it is ok if the rename gets
	// reverted due to crash.
	return nil
}

// truncate to acceptable size.

// There is no head, so just fetch tail size.

// Sync directory after each deletion to ensure that files are removed in order.
