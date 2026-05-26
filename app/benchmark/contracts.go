package benchmark

// getERC20DeployData returns the deployment bytecode for ERC20 with constructor args.
func getERC20DeployData() []byte { _ = "STUB: not implemented"; return nil }

// Constructor args: name = "LoadToken", symbol = "LT"
// Use Constructor.Inputs.Pack directly to ensure correct encoding

// getERC721DeployData returns the deployment bytecode for ERC721 with constructor args.
func getERC721DeployData() []byte { _ = "STUB: not implemented"; return nil }

// Constructor args: name = "LoadNFT", symbol = "LNFT"

// getERC20ConflictDeployData returns the deployment bytecode for ERC20Conflict.
func getERC20ConflictDeployData() []byte { _ = "STUB: not implemented"; return nil }

// Constructor args: name = "ConflictToken", symbol = "CT"

// getERC20NoopDeployData returns the deployment bytecode for ERC20Noop.
func getERC20NoopDeployData() []byte { _ = "STUB: not implemented"; return nil }

// Constructor args: name = "NoopToken", symbol = "NT"

// getDisperseDeployData returns the deployment bytecode for Disperse.
func getDisperseDeployData() []byte {
	_ = "STUB: not implemented"
	// Disperse has no constructor args
	return nil
}
