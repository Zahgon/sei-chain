// Program linkpatch rewrites absolute URLs pointing to targets in GitHub in
// Markdown link tags to target a different branch.
//
// This is used to update documentation links for backport branches.
// See https://github.com/tendermint/tendermint/issues/7675 for context.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	repoName     = flag.String("repo", "tendermint/tendermint", "Repository name to match")
	sourceBranch = flag.String("source", "master", "Source branch name (required)")
	targetBranch = flag.String("target", "", "Target branch name (required)")
	doRecur      = flag.Bool("recur", false, "Recur into subdirectories")

	skipPath  stringList
	skipMatch regexpFlag

	// Match markdown links pointing to absolute URLs.
	// This only works for "inline" links, not referenced links.
	// The submetch selects the URL.
	linkRE = regexp.MustCompile(`(?m)\[.*?\]\((https?://.*?)\)`)
)

func init() {
	flag.Var(&skipPath, "skip-path", "Skip these paths (comma-separated)")
	flag.Var(&skipMatch, "skip-match", "Skip URLs matching this regexp (RE2)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %[1]s [options] <file/dir>...

Rewrite absolute Markdown links targeting the specified GitHub repository
and source branch name to point to the target branch instead. Matching
files are updated in-place.

Each path names either a directory to list, or a single file path to
rewrite. By default, only the top level of a directory is scanned; use -recur
to recur into subdirectories.

Options:
`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	switch {
	case *repoName == "":
		log.Fatal("You must specify a non-empty -repo name (org/repo)")
	case *targetBranch == "":
		log.Fatal("You must specify a non-empty -target branch")
	case *sourceBranch == "":
		log.Fatal("You must specify a non-empty -source branch")
	case *sourceBranch == *targetBranch:
		log.Fatalf("Source and target branch are the same (%q)", *sourceBranch)
	case flag.NArg() == 0:
		log.Fatal("You must specify at least one file/directory to rewrite")
	}

	r, err := regexp.Compile(fmt.Sprintf(`^https?://github.com/%s/(?:blob|tree)/%s`,
		*repoName, *sourceBranch))
	if err != nil {
		log.Fatalf("Compiling regexp: %v", err)
	}
	for _, path := range flag.Args() {
		if err := processPath(r, path); err != nil {
			log.Fatalf("Processing %q failed: %v", path, err)
		}
	}
}

func processPath(r *regexp.Regexp, path string) error { _ = "STUB: not implemented"; return nil }

// nothing to do with links, device files, sockets, etc.

func processDir(r *regexp.Regexp, root string) error { _ = "STUB: not implemented"; return nil }

// explicitly skipped

// skipped because we aren't recurring

// nothing else to do for directories

// explicitly skipped

// nothing to do for non-Markdown files

func processFile(r *regexp.Regexp, path string) error { _ = "STUB: not implemented"; return nil }

// copy the existing data as-is

// Copy everything before the URL as-is, then write the replacement.
// everything up to the URL

// Write out the tail of the match, everything after the URL.

// the rest of the file

// stringList implements the flag.Value interface for a comma-separated list of strings.
type stringList []string

func (lst *stringList) Set(s string) error { _ = "STUB: not implemented"; return nil }

// Contains reports whether lst contains s.
func (lst stringList) Contains(s string) bool { _ = "STUB: not implemented"; return false }

func (lst stringList) String() string { _ = "STUB: not implemented"; return "" }

// regexpFlag implements the flag.Value interface for a regular expression.
type regexpFlag struct{ *regexp.Regexp }

func (r regexpFlag) MatchString(s string) bool { _ = "STUB: not implemented"; return false }

func (r *regexpFlag) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (r regexpFlag) String() string { _ = "STUB: not implemented"; return "" }
