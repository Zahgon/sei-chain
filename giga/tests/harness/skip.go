package harness

// SkipList defines tests and categories to skip during state test execution
type SkipList struct {
	Tests      map[string]string `json:"skipped_tests"`      // testName -> reason
	Categories []string          `json:"skipped_categories"` // entire categories to skip
}

// LoadSkipList loads the skip list from the data directory
func LoadSkipList() (*SkipList, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadSkipListFromPath loads a skip list from a specific path
func LoadSkipListFromPath(path string) (*SkipList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return empty skip list if file doesn't exist

// Initialize maps if nil

// ShouldSkip checks if a test should be skipped
// Returns whether to skip and the reason
func (sl *SkipList) ShouldSkip(category, testName string) (skip bool, reason string) {
	_ = "STUB: not implemented"
	// Check if entire category is skipped
	return false, ""
}

// Check specific test by full name (category/testName)

// Also check with just the test name (for backward compatibility)

// Check for partial matches - allows skipping by short name like "category/testBase"
// which will match "category/testBase.json/..." style full test names

// Pattern format: "category/shortName" should match if testName contains shortName

// Match if the test name starts with the short name (before any .json or /)

// SkippedCount returns the number of tests that would be skipped
func (sl *SkipList) SkippedCount() int { _ = "STUB: not implemented"; return 0 }

// SkippedCategoriesCount returns the number of skipped categories
func (sl *SkipList) SkippedCategoriesCount() int { _ = "STUB: not implemented"; return 0 }

// IsCategorySkipped checks if an entire category is skipped
func (sl *SkipList) IsCategorySkipped(category string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetSkipReason returns the skip reason for a test, or empty string if not skipped
func (sl *SkipList) GetSkipReason(category, testName string) string {
	_ = "STUB: not implemented"
	return ""
}
