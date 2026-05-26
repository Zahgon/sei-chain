package cosmovisor

const (
	rootName    = "cosmovisor"
	genesisDir  = "genesis"
	upgradesDir = "upgrades"
	currentLink = "current"
)

// Config is the information passed in to control the daemon
type Config struct {
	Home                  string
	Name                  string
	AllowDownloadBinaries bool
	RestartAfterUpgrade   bool
	LogBufferSize         int
}

// Root returns the root directory where all info lives
func (cfg *Config) Root() string { _ = "STUB: not implemented"; return "" }

// GenesisBin is the path to the genesis binary - must be in place to start manager
func (cfg *Config) GenesisBin() string { _ = "STUB: not implemented"; return "" }

// UpgradeBin is the path to the binary for the named upgrade
func (cfg *Config) UpgradeBin(upgradeName string) string { _ = "STUB: not implemented"; return "" }

// UpgradeDir is the directory named upgrade
func (cfg *Config) UpgradeDir(upgradeName string) string { _ = "STUB: not implemented"; return "" }

// Symlink to genesis
func (cfg *Config) SymLinkToGenesis() (string, error) { _ = "STUB: not implemented"; return "", nil }

// and return the genesis binary

// CurrentBin is the path to the currently selected binary (genesis if no link is set)
// This will resolve the symlink to the underlying directory to make it easier to debug
func (cfg *Config) CurrentBin() (string, error) { _ = "STUB: not implemented"; return "", nil }

// if nothing here, fallback to genesis

//Create symlink to the genesis

// if it is there, ensure it is a symlink

//Create symlink to the genesis

// resolve it

//Create symlink to the genesis

// and return the binary

// GetConfigFromEnv will read the environmental variables into a config
// and then validate it is reasonable
func GetConfigFromEnv() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// validate returns an error if this config is invalid.
// it enforces Home/cosmovisor is a valid directory and exists,
// and that Name is set
func (cfg *Config) validate() error { _ = "STUB: not implemented"; return nil }

// ensure the root directory exists
