package memiavl

type FileLock interface {
	Unlock() error
	Destroy() error
}

func LockFile(fname string) (FileLock, error) {
	_ = "STUB: not implemented"
	return *new(FileLock), nil
}
