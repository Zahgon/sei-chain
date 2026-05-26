package crypto

const (
	blockTypePrivKey = "TENDERMINT PRIVATE KEY"
	blockTypeKeyInfo = "TENDERMINT KEY INFO"
	blockTypePubKey  = "TENDERMINT PUBLIC KEY"

	defaultAlgo = "secp256k1"

	headerVersion = "version"
	headerType    = "type"
)

func EncodeArmor(blockType string, headers map[string]string, data []byte) string {
	_ = "STUB: not implemented"
	return ""
}

func DecodeArmor(armorStr string) (blockType string, headers map[string]string, data []byte, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil, nil
}

const nonceLen = 24
const secretLen = 32

// secret must be 32 bytes long. Use something like Sha256(Bcrypt(passphrase))
// The ciphertext is (secretbox.Overhead + 24) bytes longer than the plaintext.
func EncryptSymmetric(plaintext []byte, secret []byte) (ciphertext []byte) {
	_ = "STUB: not implemented"
	return nil
}

// secret must be 32 bytes long. Use something like Sha256(Bcrypt(passphrase))
// The ciphertext is (secretbox.Overhead + 24) bytes longer than the plaintext.
func DecryptSymmetric(ciphertext []byte, secret []byte) (plaintext []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BcryptSecurityParameter is security parameter var, and it can be changed within the lcd test.
// Making the bcrypt security parameter a var shouldn't be a security issue:
// One can't verify an invalid key by maliciously changing the bcrypt
// parameter during a runtime vulnerability. The main security
// threat this then exposes would be something that changes this during
// runtime before the user creates their key. This vulnerability must
// succeed to update this to that same value before every subsequent call
// to the keys command in future startups / or the attacker must get access
// to the filesystem. However, with a similar threat model (changing
// variables in runtime), one can cause the user to sign a different tx
// than what they see, which is a significantly cheaper attack then breaking
// a bcrypt hash. (Recall that the nonce still exists to break rainbow tables)
// For further notes on security parameter choice, see README.md
var BcryptSecurityParameter = 12

//-----------------------------------------------------------------
// encrypt/decrypt with armor

// Encrypt and armor the private key.
func EncryptArmorPrivKey(privKeyBytes []byte, passphrase string, algo string) string {
	_ = "STUB: not implemented"
	return ""
}

// encrypt the given privKey with the passphrase using a randomly
// generated salt and the xsalsa20 cipher. returns the salt and the
// encrypted priv key.
func encryptPrivKey(privKeyBytes []byte, passphrase string) (saltBytes []byte, encBytes []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get 32 bytes

// UnarmorDecryptPrivKey returns the privkey byte slice, a string of the algo type, and an error
func UnarmorDecryptPrivKey(armorStr string, passphrase string) (privKey []byte, algo string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func decryptPrivKey(saltBytes []byte, encBytes []byte, passphrase string) (privKey []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get 32 bytes

// return legacy.PrivKeyFromBytes(privKeyBytes)
