package harness

import (
	"sync"
)

// extractOnce ensures we only extract the archive once per test run
var extractOnce sync.Once

// LoadStateTest loads a state test from a JSON file
func LoadStateTest(filePath string) (map[string]*StateTestJSON, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadStateTestsFromDir loads all state tests from a directory
func LoadStateTestsFromDir(dirPath string) (map[string]*StateTestJSON, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use relative path + test name as key to avoid collisions

// GetStateTestsPath returns the path to GeneralStateTests, extracting from archive if needed
func GetStateTestsPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Always verify the directory exists after sync.Once completes,
// regardless of whether this was the first invocation

// extractArchive extracts a .tgz archive to the destination directory
func extractArchive(archivePath, destDir string) error { _ = "STUB: not implemented"; return nil }
