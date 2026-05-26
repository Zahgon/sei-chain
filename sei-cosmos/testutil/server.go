package testutil

type TestAppOpts struct{}

func (t TestAppOpts) Get(s string) interface{} { _ = "STUB: not implemented"; return nil }
