package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

const evmoneVersion = "0.12.0"

type Platform struct {
	Archive string
	Hash    string // SHA256 hash of the archive
	LibHash string // SHA256 hash of the extracted library file
	OS      string
	Arch    string
	LibPath string
	Ext     string
}

var platforms = []Platform{
	{
		Archive: "evmone-0.12.0-linux-x86_64.tar.gz",
		Hash:    "1c7b5eba0c8c3b3b2a7a05101e2d01a13a2f84b323989a29be66285dba4136ce",
		LibHash: "0fec5d79f4c9a466bb680e8b0b9c770aea38f3dd6d2e4af23535c893d0d18d40",
		OS:      "linux",
		Arch:    "amd64",
		LibPath: "lib/libevmone.so.0.12.0",
		Ext:     "so",
	},
	{
		Archive: "evmone-0.12.0-darwin-arm64.tar.gz",
		Hash:    "e164e0d2b985cc1cca07b501538b2e804bf872d1d8d531f9241d518a886234a6",
		LibHash: "cb1c555b3849a0a6a9402bc907a9a7bfe14e1c08c483c488ec6ba5f19e986847",
		OS:      "darwin",
		Arch:    "arm64",
		LibPath: "lib/libevmone.0.12.0.dylib",
		Ext:     "dylib",
	},
}

// This program downloads evmone shared libraries for supported platforms.
//
// To upgrade evmone:
//  1. Visit https://github.com/ethereum/evmone/releases
//  2. Update evmoneVersion constant below
//  3. Update the Archive filenames and SHA256 hashes in the platforms slice
//  4. Update the LibHash values for extracted library files
//  5. Run: go run download_evmone.go <output-dir>
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <output-dir>\n", os.Args[0])
	}
	outDir := filepath.Clean(os.Args[1])

	// Create context that cancels on SIGINT or SIGTERM
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	for _, p := range platforms {
		if err := downloadAndExtract(ctx, p, outDir); err != nil {
			if ctx.Err() != nil {
				log.Fatal("Interrupted")
			}
			log.Fatalf("Failed %s-%s: %v\n", p.OS, p.Arch, err)
		}
	}
	log.Println("All platforms downloaded successfully!")
}

// libFileName returns the output filename for a platform's library
func (p Platform) libFileName() string { _ = "STUB: not implemented"; return "" }

// checkExistingLib checks if the library file already exists with the correct hash.
// Returns true if the file exists and has the correct hash, false otherwise.
func checkExistingLib(p Platform, outDir string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func downloadAndExtract(ctx context.Context, p Platform, outDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if file already exists with correct hash

// Check for cancellation after download

// Verify SHA-256 hash of the downloaded archive

//nolint:gosec

// 100MiB maximum copy
