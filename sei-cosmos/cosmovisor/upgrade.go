package cosmovisor

// DoUpgrade will be called after the log message has been parsed and the process has terminated.
// We can now make any changes to the underlying directory without interference and leave it
// in a state, so we can make a proper restart
func DoUpgrade(cfg *Config, info *UpgradeInfo) error {
	_ = "STUB: not implemented"
	// Simplest case is to switch the link
	return nil
}

// we have the binary - do it

// if auto-download is disabled, we fail

// if the dir is there already, don't download either

// If not there, then we try to download it... maybe

// and then set the binary again

// DownloadBinary will grab the binary and place it in the proper directory
func DownloadBinary(cfg *Config, info *UpgradeInfo) error { _ = "STUB: not implemented"; return nil }

// download into the bin dir (works for one file)

// if this fails, let's see if it is a zipped directory

// copy binary to binPath from dirPath if zipped directory don't contain bin directory to wrap the binary

// if it is successful, let's ensure the binary is executable

// MarkExecutable will try to set the executable bits if not already set
// Fails if file doesn't exist or we cannot set those bits
func MarkExecutable(path string) error { _ = "STUB: not implemented"; return nil }

// end early if world exec already set

// now try to set all exec bits

// UpgradeConfig is expected format for the info field to allow auto-download
type UpgradeConfig struct {
	Binaries map[string]string `json:"binaries"`
}

// GetDownloadURL will check if there is an arch-dependent binary specified in Info
func GetDownloadURL(info *UpgradeInfo) (string, error) { _ = "STUB: not implemented"; return "", nil }

// if this is a url, then we download that and try to get a new doc with the real info

// if download worked properly, then we use this new file as the binary map to parse

// check if it is the upgrade config

func OSArch() string { _ = "STUB: not implemented"; return "" }

// SetCurrentUpgrade sets the named upgrade to be the current link, returns error if this binary doesn't exist
func (cfg *Config) SetCurrentUpgrade(upgradeName string) error {
	_ = "STUB: not implemented"
	// ensure named upgrade exists
	return nil
}

// set a symbolic link

// remove link if it exists

// point to the new directory

// EnsureBinary ensures the file exists and is executable, or returns an error
func EnsureBinary(path string) error { _ = "STUB: not implemented"; return nil }

// this checks if the world-executable bit is set (we cannot check owner easily)
