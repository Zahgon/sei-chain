package util

// RecursiveMove transfers files/directory trees from the source to the destination.
//
// If preserveOriginal is false, then the files at the source will be deleted when this method returns.
// If preserveOriginal is true, then this function will leave behind a copy of the original files at the source.
//
// This function does not support symlinks. It will return an error if it encounters any symlinks in the source path.
func RecursiveMove(
	source string,
	destination string,
	preserveOriginal bool,
	fsync bool,
) error {
	_ = "STUB: not implemented"
	// Sanitize paths
	return nil
}

// Verify source exists

// Verify destination parent directory is writable

// If source is a file, handle it directly

// Source is a directory, handle recursively

// moveFile handles moving a single file
func moveFile(source string, destination string, preserveOriginal bool, fsync bool) error {
	_ = "STUB: not implemented"
	// Ensure parent directory exists
	return nil
}

// If not preserving original, try to move the file first (regardless of deep mode)

// Try simple rename first (works if on same filesystem)

// Rename failed (likely different filesystem), fall back to copy+delete

// Copy the file

// Sync if requested

// sync parent directory

// Remove source if not preserving original

// recursiveMoveDirectory handles moving a directory and its contents
func recursiveMoveDirectory(
	source string,
	destination string,
	preserveOriginal bool,
	fsync bool,
) error {
	_ = "STUB: not implemented"

	// Create destination directory if it doesn't exist
	return nil
}

// Walk through source directory

// Skip the root directory itself

// Calculate relative path and destination path

// Create directory at destination

// Move the file

// Sync destination directory if requested

// Remove source directory if not preserving original
