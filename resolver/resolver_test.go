package resolver

import (
	"testing"
)

func TestGetVersionListProxyURL(t *testing.T) {
	url := getVersionListProxyURL("github.com/barelyhuman/commitlog")
	if url != "https://proxy.golang.org/github.com/barelyhuman/commitlog/@v/list" {
		t.Fatalf("Invalid proxy version list url")
	}
}

func TestGetVersionLatestProxyURL(t *testing.T) {
	url := getVersionLatestProxyURL("github.com/barelyhuman/commitlog")
	if url != "https://proxy.golang.org/github.com/barelyhuman/commitlog/@latest" {
		t.Fatalf("Invalid proxy version latest url, result:%s", url)
	}
}

func TestNormalizeURL(t *testing.T) {
	url := normalizeUrl("https://proxy.golang.org/")
	if normalizeUrl("https://proxy.golang.org/") != "https://proxy.golang.org" {
		t.Fatalf("Failed to normalize proxy url,result:%s", url)
	}
	url = normalizeUrl("https://proxy.golang.org")
	if normalizeUrl("https://proxy.golang.org") != "https://proxy.golang.org" {
		t.Fatalf("Failed to normalize proxy url,result:%s", url)
	}
}

func TestParseVersion(t *testing.T) {
	r := Resolver{
		Pkg: "barelyhuman/commitlog",
	}
	err := r.ParseVersion("0.0.6")
	if err != nil {
		t.Fatalf("Failed to parse version, err:%v", err)
	}
	if len(r.Value) == 0 {
		t.Fatalf("Failed to assign value")
	}

	if r.ConstraintCheck == nil {
		t.Fatalf("Failed to create a constraint checker")
	}
}

func TestParseVersionWithHash(t *testing.T) {
	r := Resolver{
		Pkg: "barelyhuman/commitlog",
	}
	commitHash := "bba8d7a63d622e4f12dbea9722b647cd985be8ad"
	err := r.ParseVersion("bba8d7a63d622e4f12dbea9722b647cd985be8ad")
	if err != nil {
		t.Fatalf("Failed to parse version, err:%v", err)
	}

	if r.Value != commitHash {
		t.Fatalf("Failed to assign value, value:%v, hash:%v", r.Value, commitHash)
	}

	if r.ConstraintCheck != nil {
		t.Fatalf("Created a constraint check for invalid semver")
	}
}

func TestResolveClosestVersion(t *testing.T) {
	versionToResolve := "v0.0.6"
	inputVersion := "0.0.6"
	r := Resolver{
		Pkg: "github.com/barelyhuman/commitlog",
	}

	// parse with a version
	err := r.ParseVersion(inputVersion)
	if err != nil {
		t.Fatalf("Failed to parse version,err:%v", err)
	}
	version, err := r.ResolveClosestVersion()
	if err != nil {
		t.Fatalf("Failed to get closest version,err:%v", err)
	}
	if version != versionToResolve {
		t.Fatalf("Resolved invalid version,resolved version:%v,should resolve to:%v,input:%v", version, versionToResolve, inputVersion)
	}

}

func TestResolveLatestVersion(t *testing.T) {
	versionToResolve := "v0.0.10"
	inputVersion := ""
	r := Resolver{
		Pkg: "github.com/barelyhuman/commitlog",
	}
	err := r.ParseVersion(inputVersion)
	if err != nil {
		t.Fatalf("Failed to parse version,err:%v", err)
	}
	versionInfo, err := r.ResolveLatestVersion()
	if err != nil {
		t.Fatalf("Failed to get closes version,err:%v", err)
	}
	if versionInfo.Version != versionToResolve {
		t.Fatalf("Resolved invalid version,resolved version:%v,should resolve to:%v,input:%v", versionInfo.Version, versionToResolve, inputVersion)
	}
}

func TestResolveVersionWithVersion(t *testing.T) {
	versionToResolve := "v0.0.7-dev.5"
	inputVersion := "0.0.7-dev.5"
	r := Resolver{
		Pkg: "github.com/barelyhuman/commitlog",
	}
	err := r.ParseVersion(inputVersion)
	if err != nil {
		t.Fatalf("Failed to parse version, err:%v", err)
	}

	version, err := r.ResolveVersion()
	if err != nil {
		t.Fatalf("Failed to resolve, err:%v", err)
	}
	if version != versionToResolve {
		t.Fatalf("Failed to resolve, resolved:%v,expected resolve:%v", version, versionToResolve)
	}
}

func TestResolveVersionWithoutVersion(t *testing.T) {
	versionToResolve := "v0.0.10"
	inputVersion := ""
	r := Resolver{
		Pkg: "github.com/barelyhuman/commitlog",
	}
	err := r.ParseVersion(inputVersion)
	if err != nil {
		t.Fatalf("Failed to parse version, err:%v", err)
	}

	version, err := r.ResolveVersion()
	if err != nil {
		t.Fatalf("Failed to resolve, err:%v", err)
	}
	if version != versionToResolve {
		t.Fatalf("Failed to resolve, resolved:%v,expected resolve:%v", version, versionToResolve)
	}
}

func TestResolveVersionWithHash(t *testing.T) {
	versionToResolve := "7e0664aba1db8e44d11f9d457bd5bb583a8000ba"
	inputVersion := "7e0664aba1db8e44d11f9d457bd5bb583a8000ba"
	r := Resolver{
		Pkg: "github.com/barelyhuman/commitlog",
	}
	err := r.ParseVersion(inputVersion)
	if err != nil {
		t.Fatalf("Failed to parse version, err:%v", err)
	}

	version, err := r.ResolveVersion()
	if err != nil {
		t.Fatalf("Failed to resolve, err:%v", err)
	}
	if version != versionToResolve {
		t.Fatalf("Failed to resolve, resolved:%v,expected resolve:%v", version, versionToResolve)
	}
}

func TestFailGettingLatest(t *testing.T) {
	pkgName := "barelyhuman/commitlog"
	version := "0.0.6"

	r := Resolver{
		Pkg: pkgName,
	}
	r.ParseVersion(version)

	versionInfo, err := r.ResolveLatestVersion()
	if err == nil {
		t.Fatalf("Resolved for invalid package, pkg:%s , version:%s", pkgName, versionInfo.Version)
	}
}

func TestFailGettingClosest(t *testing.T) {
	pkgName := "barelyhuman/commitlog"
	version := "0.0.6"

	r := Resolver{
		Pkg: pkgName,
	}
	r.ParseVersion(version)

	versionInfo, err := r.ResolveClosestVersion()
	if err == nil {
		t.Fatalf("Resolved for invalid package, pkg:%s , version:%s", pkgName, versionInfo)
	}
}
