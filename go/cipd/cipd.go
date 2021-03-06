package cipd

/*
	Utilities for working with CIPD.
*/

//go:generate go run gen_versions.go

import (
	"context"
	"fmt"

	"go.chromium.org/luci/cipd/client/cipd"
	"go.chromium.org/luci/cipd/client/cipd/common"
	"go.skia.org/infra/go/auth"
)

const (
	// CIPD server to use for obtaining packages.
	SERVICE_URL = "https://chrome-infra-packages.appspot.com"
)

var (
	// CIPD package for the Go installation.
	PkgGo = &Package{
		Dest:    "go",
		Name:    "skia/bots/go",
		Version: fmt.Sprintf("version:%s", PKG_VERSIONS_FROM_ASSETS["go"]),
	}
)

// Package describes a CIPD package.
type Package struct {
	// Relative path within the root dir to install the package.
	Dest string

	// Name of the package.
	Name string

	// Version of the package. See the CIPD docs for valid version strings:
	// https://godoc.org/go.chromium.org/luci/cipd/client/cipd/common#ValidateInstanceVersion
	Version string
}

// Run "cipd ensure" to get the correct packages in the given location.
func Ensure(local bool, rootdir string, packages ...*Package) error {
	httpClient, err := auth.NewDefaultClient(local, auth.SCOPE_USERINFO_EMAIL, auth.SCOPE_PLUS_ME)
	if err != nil {
		return fmt.Errorf("Failed to create authenticated HTTP client: %s", err)
	}
	c, err := cipd.NewClient(cipd.ClientOptions{
		ServiceURL:          SERVICE_URL,
		Root:                rootdir,
		AuthenticatedClient: httpClient,
	})
	if err != nil {
		return fmt.Errorf("Failed to create CIPD client: %s", err)
	}

	ctx := context.Background()
	pkgs := common.PinSliceBySubdir{}
	for _, pkg := range packages {
		pin, err := c.ResolveVersion(ctx, pkg.Name, pkg.Version)
		if err != nil {
			return fmt.Errorf("Failed to resolve package version %q @ %q: %s", pkg.Name, pkg.Version, err)
		}
		pkgs[pkg.Dest] = common.PinSlice{pin}
	}
	if _, err := c.EnsurePackages(context.Background(), pkgs, false); err != nil {
		return fmt.Errorf("Failed to ensure packages: %s", err)
	}
	return nil
}
