package client

func unmarshalResponseBytes(responseBytes []byte, expectedID string, result interface{}) error {
	_ = "STUB: not implemented"
	// Read response.  If rpc/core/types is imported, the result will unmarshal
	// into the correct type.
	return nil
}

// Unmarshal the RawMessage into the result.

func unmarshalResponseBytesArray(responseBytes []byte, expectedIDs []string, results []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Intersect IDs from responses with expectedIDs.

func validateResponseIDs(ids, expectedIDs []string) error { _ = "STUB: not implemented"; return nil }
