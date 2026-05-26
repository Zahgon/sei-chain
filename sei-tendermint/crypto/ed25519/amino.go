package ed25519

// Another way to leak a secret. Needed for LegacyAminoEncoding.
func (k SecretKey) MarshalAmino() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *SecretKey) UnmarshalAmino(secretBytes []byte) error { _ = "STUB: not implemented"; return nil }

func (k PublicKey) MarshalAmino() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (k *PublicKey) UnmarshalAmino(bytes []byte) error { _ = "STUB: not implemented"; return nil }
