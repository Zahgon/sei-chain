package strings

// SplitAndTrimEmpty slices s into all subslices separated by sep and returns a
// slice of the string s with all leading and trailing Unicode code points
// contained in cutset removed. If sep is empty, SplitAndTrim splits after each
// UTF-8 sequence. First part is equivalent to strings.SplitN with a count of
// -1.  also filter out empty strings, only return non-empty strings.
func SplitAndTrimEmpty(s, sep, cutset string) []string { _ = "STUB: not implemented"; return nil }

// ASCIITrim removes spaces from an a ASCII string, erroring if the
// sequence is not an ASCII string.
func ASCIITrim(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// skip space

// StringSliceEqual checks if string slices a and b are equal
func StringSliceEqual(a, b []string) bool { _ = "STUB: not implemented"; return false }
