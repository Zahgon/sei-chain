// Program condiff performs a keyspace diff on two TOML documents.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/creachadair/tomledit"
	"github.com/creachadair/tomledit/transform"
)

var (
	doDesnake = flag.Bool("desnake", false, "Convert snake_case to kebab-case before comparing")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %[1]s [options] f1 f2

Diff the keyspaces of the TOML documents in files f1 and f2.
The output prints one line per key that differs:

   -S name    -- section exists in f1 but not f2
   +S name    -- section exists in f2 but not f1
   -M name    -- mapping exists in f1 but not f2
   +M name    -- mapping exists in f2 but not f1

Comments, order, and values are ignored for comparison purposes.

Options:
`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		log.Fatalf("Usage: %[1]s <lhs> <rhs>", filepath.Base(os.Args[0]))
	}
	lhs := mustParse(flag.Arg(0))
	rhs := mustParse(flag.Arg(1))
	if *doDesnake {
		log.Printf("Converting all names from snake_case to kebab-case")
		fix := transform.SnakeToKebab()
		_ = fix(context.Background(), lhs)
		_ = fix(context.Background(), rhs)
	}
	diffDocs(os.Stdout, lhs, rhs)
}

func mustParse(path string) *tomledit.Document { _ = "STUB: not implemented"; return nil }

func allKeys(s *tomledit.Section) []string { _ = "STUB: not implemented"; return nil }

const (
	delSection = "-S"
	delMapping = "-M"
	addSection = "+S"
	addMapping = "+M"

	delMapSep = "\n" + delMapping + " "
	addMapSep = "\n" + addMapping + " "
)

func diffDocs(w io.Writer, lhs, rhs *tomledit.Document) { _ = "STUB: not implemented"; return }

func diffSections(w io.Writer, lhs, rhs *tomledit.Section) { _ = "STUB: not implemented"; return }

func diffKeys(w io.Writer, lhs, rhs []string) { _ = "STUB: not implemented"; return }
