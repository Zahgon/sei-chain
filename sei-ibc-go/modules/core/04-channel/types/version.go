package types

const ChannelVersionDelimiter = ":"

// SplitChannelVersion splits the channel version string
// into the outermost middleware version and the underlying app version.
// It will use the default delimiter `:` for middleware versions.
// In case there's no delimeter, this function returns an empty string for the middleware version (first return argument),
// and the full input as the second underlying app version.
func SplitChannelVersion(version string) (middlewareVersion, appVersion string) {
	_ = "STUB: not implemented"
	// only split out the first middleware version
	return "", ""
}

// MergeChannelVersions merges the provided versions together with the channel version delimiter
// the versions should be passed in from the highest-level middleware to the base application
func MergeChannelVersions(versions ...string) string { _ = "STUB: not implemented"; return "" }
