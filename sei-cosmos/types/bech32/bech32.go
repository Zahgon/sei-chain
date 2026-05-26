package bech32

// ConvertAndEncode converts from a base64 encoded byte string to base32 encoded byte string and then to bech32.
func ConvertAndEncode(hrp string, data []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// DecodeAndConvert decodes a bech32 encoded string and converts to base64 encoded bytes.
func DecodeAndConvert(bech string) (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
