package keyring

func infoKey(name string) string   { _ = "STUB: not implemented"; return "" }
func infoKeyBz(name string) []byte { _ = "STUB: not implemented"; return nil }

// KeybaseOption overrides options for the db.
type KeybaseOption func(*kbOptions)

type kbOptions struct {
}
