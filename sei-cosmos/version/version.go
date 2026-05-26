// Package version is a convenience utility that provides SDK
// consumers with a ready-to-use version command that
// produces apps versioning information based on flags
// passed at compile time.
//
// # Configure the version command
//
// The version command can be just added to your cobra root command.
// At build time, the variables Name, Version, Commit, and BuildTags
// can be passed as build flags as shown in the following example:
//
//	go build -X github.com/cosmos/cosmos-sdk/version.Name=gaia \
//	 -X github.com/cosmos/cosmos-sdk/version.AppName=gaiad \
//	 -X github.com/cosmos/cosmos-sdk/version.Version=1.0 \
//	 -X github.com/cosmos/cosmos-sdk/version.Commit=f0f7b7dab7e36c20b757cebce0e8f4fc5b95de60 \
//	 -X "github.com/sei-protocol/sei-chain/sei-cosmos/version.BuildTags=linux darwin amd64"
package version

import (
	"runtime/debug"
)

var (
	// application's name
	Name = ""
	// application binary name
	AppName = "<appd>"
	// application's version string
	Version = ""
	// commit
	Commit = ""
	// build tags
	BuildTags = ""
)

func getSDKVersion() string { _ = "STUB: not implemented"; return "" }

// Info defines the application version information.
type Info struct {
	Name             string     `json:"name" yaml:"name"`
	AppName          string     `json:"server_name" yaml:"server_name"`
	Version          string     `json:"version" yaml:"version"`
	GitCommit        string     `json:"commit" yaml:"commit"`
	BuildTags        string     `json:"build_tags" yaml:"build_tags"`
	GoVersion        string     `json:"go" yaml:"go"`
	BuildDeps        []buildDep `json:"build_deps" yaml:"build_deps"`
	CosmosSdkVersion string     `json:"cosmos_sdk_version" yaml:"cosmos_sdk_version"`
}

func NewInfo() Info { _ = "STUB: not implemented"; return *new(Info) }

func (vi Info) String() string { _ = "STUB: not implemented"; return "" }

func depsFromBuildInfo() (deps []buildDep) { _ = "STUB: not implemented"; return nil }

type buildDep struct {
	*debug.Module
}

func (d buildDep) String() string { _ = "STUB: not implemented"; return "" }

func (d buildDep) MarshalJSON() ([]byte, error)      { _ = "STUB: not implemented"; return nil, nil }
func (d buildDep) MarshalYAML() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
