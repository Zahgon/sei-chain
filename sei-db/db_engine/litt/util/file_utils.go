package util

// SwapFileExtension is the file extension used for temporary swap files created during atomic writes.
const SwapFileExtension = ".swap"

// IsSymlink checks if the given path is a symlink.
func IsSymlink(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Path does not exist, so it can't be a symlink

// ErrIfSymlink checks if the given path is a symlink and returns an error if it is.
func ErrIfSymlink(path string) error { _ = "STUB: not implemented"; return nil }

// IsDirectory checks if the given path is a directory. Returns false if the path is not a directory or does not exist.
func IsDirectory(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Path does not exist, so it can't be a directory

// SanitizePath returns a sanitized version of the given path, doing things like expanding
// "~" to the user's home directory, converting to absolute path, normalizing slashes, etc.
func SanitizePath(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// DeleteOrphanedSwapFiles deletes any swap files in the given directory, i.e. files that end with ".swap".
func DeleteOrphanedSwapFiles(directory string) error { _ = "STUB: not implemented"; return nil }

// AtomicWrite writes data to a file atomically. The parent directory must exist and be writable.
// If the destination file already exists, it will be overwritten.
//
// This method creates a temporary swap file in the same directory as the destination, but with SwapFileExtension
// appended to the filename. If there is a crash during this method's execution, it may leave this swap file behind.
func AtomicWrite(destination string, data []byte, fsync bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Write the data into the swap file.
//nolint:gosec // caller-supplied destination path

// Ensure the data in the swap file is fully written to disk.

// Rename the swap file to the destination file.

// AtomicRename renames a file from oldPath to newPath atomically.
func AtomicRename(oldPath string, newPath string, fsync bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure that the rename is committed to disk.
//nolint:gosec // derived from caller-supplied path

// ErrIfNotWritableFile verifies that a path is either a regular file with read+write permissions,
// or that it is legal to create a new regular file with read+write permissions in the parent directory.
//
// A file is considered to have the correct permissions/type if:
// - it exists and is a standard file with read+write permissions
// - if it does not exist but its parent directory has read+write permissions.
//
// The arguments for the function are the result of os.Stat(path). There is no need to do error checking on the
// result of os.Stat in the calling context (this method does it for you).
func ErrIfNotWritableFile(path string) (exists bool, size int64, err error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

// The file does not exist. Check the parent.

// File exists. Check if it is a regular file and that it is readable+writeable.

// ErrIfNotWritableDirectory checks if a directory exists and is writable, or if it doesn't exist but it would
// be legal to create it.
func ErrIfNotWritableDirectory(dirPath string) error { _ = "STUB: not implemented"; return nil }

// Directory doesn't exist, check parent permissions

// Path exists, verify it's a directory with write permissions

// Returns an error if the given path exists, otherwise returns nil.
func ErrIfExists(path string) error { _ = "STUB: not implemented"; return nil }

// Returns an error if the given path does not exist, otherwise returns nil.
func ErrIfNotExists(path string) error { _ = "STUB: not implemented"; return nil }

// Exists checks if a file or directory exists at the given path. More aesthetically pleasant than os.Stat.
func Exists(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// SyncFile syncs a file/directory
func SyncPath(path string) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // caller-supplied path

// SyncParentPath syncs the parent directory of the given path.
func SyncParentPath(path string) error { _ = "STUB: not implemented"; return nil }

// CopyRegularFile copies a regular file from src to dst. If a file already exists at dst, it will be removed
// before copying.
func CopyRegularFile(src string, dst string, fsync bool) error {
	_ = "STUB: not implemented"
	// Ensure parent directory exists
	return nil
}

// Open source file
//nolint:gosec // caller-supplied source path

// If there is already a file at the destination, remove it.
// This ensures we don't have issues with file permissions or existing symlinks

// Create destination file
//nolint:gosec // caller-supplied destination path

// Copy content

// Sync if requested

// EnsureParentDirectoryExists ensures the parent directory of the given path exists and is writable.
// Creates parent directories if they don't exist.
func EnsureParentDirectoryExists(path string, fsync bool) error {
	_ = "STUB: not implemented"
	return nil
}

// EnsureDirectoryExists ensures a directory exists with the given permissions.
// If the directory already exists, it verifies it has write permissions.
// If fsync is true, all newly created directories are synced to disk.
func EnsureDirectoryExists(dirPath string, fsync bool) error {
	_ = "STUB: not implemented"
	// Convert to absolute path to ensure clean processing
	return nil
}

// Find the first ancestor that exists

// Check if current path exists

// Path exists, verify it's a directory

// Found existing ancestor

// Path doesn't exist, add to list of paths to create

// Move to parent directory

// Reached filesystem root. filepath.Dir("/") returns "/", so we stop here.

// Create directories from top-level to bottom-level and possibly sync each one

// Create the directory

// Sync the newly created directory

// Also sync the parent directory to ensure the directory entry is persisted

// DeepDelete deletes a regular file. If the file is a symlink, the symlink and the file pointed to by the symlink
// are both deleted. This method can delete an empty directory, but will return an error if asked to delete a
// non-empty directory. For the sake of simplicity, this method does not traverse chain of symlinks. If the symlink
// points to another symlink, it will only delete original symlink and the symlink that the original symlink points to.
func DeepDelete(path string) error { _ = "STUB: not implemented"; return nil }

// remove the file where the symlink points
