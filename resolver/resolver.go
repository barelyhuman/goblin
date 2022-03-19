package resolver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/google/go-github/v28/github"
	goSem "github.com/tj/go-semver"
	"golang.org/x/oauth2"

	"github.com/Masterminds/semver"
)

const DEFAULT_PROXY_URL = "https://proxy.golang.org"

const hashRegex = "^[0-9a-f]{7,40}$"

var gh = &GitHub{}

type Resolver struct {
	Pkg             string
	Value           string
	Hash            bool
	ConstraintCheck *semver.Constraints
	ghClient        *github.Client
}

type GitHub struct {
	// Client is the GitHub client.
	Client *github.Client
}

type VersionInfo struct {
	Version string    // version string
	Time    time.Time // commit time
}

func init() {
	ctx := context.Background()

	ghClient := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: os.Getenv("GITHUB_TOKEN"),
		},
	)

	gh.Client = github.NewClient(oauth2.NewClient(ctx, ghClient))
}

// Resolve the version for the given package by
// checking with the proxy for either the specified version
// or getting the latest version on the proxy
func (v *Resolver) ResolveVersion() (string, error) {
	if len(v.Value) == 0 {
		version, err := v.ResolveLatestVersion()
		if err != nil {
			if v.isGithubPKG() {
				return v.GithubFallbackResolveVersion()
			}
			return "", err
		}

		if len(version.Version) == 0 {
			return v.GithubFallbackResolveVersion()
		}

		return version.Version, err
	}

	if v.Hash {
		return v.Value, nil
	}

	versionString, err := v.ResolveClosestVersion()
	if err != nil {
		return "", err
	}

	if len(versionString) == 0 && v.isGithubPKG() {
		return v.GithubFallbackResolveVersion()
	}

	return versionString, nil
}

func (v *Resolver) isGithubPKG() bool {
	parts := strings.Split(v.Pkg, "/")
	return parts[0] == "github.com"
}

func (v *Resolver) GithubFallbackResolveVersion() (string, error) {
	parts := strings.Split(v.Pkg, "/")

	version := "master"
	if len(v.Value) == 0 {
		version = v.Value
	}

	resolvedV, err := gh.resolve(parts[1], parts[2], version)
	if err != nil {
		return "", err
	}

	return resolvedV, nil
}

// resolve the latest version from the proxy
func (v *Resolver) ResolveLatestVersion() (VersionInfo, error) {
	var versionInfo VersionInfo

	resp, err := http.Get(getVersionLatestProxyURL(v.Pkg))
	if err != nil {
		return versionInfo, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return versionInfo, err
	}

	if err := json.Unmarshal(respBytes, &versionInfo); err != nil {
		return versionInfo, err
	}

	return versionInfo, nil
}

// Parse the given string to be either a semver version string
// or a commit hash
func (v *Resolver) ParseVersion(version string) error {
	v.Value = version

	// just send back if no version is provided
	if len(version) == 0 {
		return nil
	}

	// return the string back if it's a valid hash string
	if !isSemver(version) && !isValidSemverConstraint(version) {
		matched, err := regexp.MatchString(hashRegex, version)
		if matched {
			v.Hash = true
			return nil
		}
		// if not a hash or a semver, just return an error
		if err != nil {
			return err
		}
	}

	if isSemver(version) {
		check, err := semver.NewConstraint("= " + version)
		if err != nil {
			return err
		}

		v.ConstraintCheck = check
	}

	if isValidSemverConstraint(version) {
		check, err := semver.NewConstraint(version)
		if err != nil {
			return err
		}

		v.ConstraintCheck = check
	}

	return nil
}

// Resolve the closes version to the given semver from the proxy
func (v *Resolver) ResolveClosestVersion() (string, error) {
	var versionTags []string

	resp, err := http.Get(getVersionListProxyURL(v.Pkg))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	versionTags = strings.Split(string(data), "\n")
	matchedVersion := ""

	var sortedVersionTags []*semver.Version

	for _, versionTag := range versionTags {
		if len(versionTag) == 0 {
			continue
		}

		ver, err := semver.NewVersion(versionTag)
		if err != nil {
			return "", err
		}
		sortedVersionTags = append(sortedVersionTags, ver)
	}
	sort.Sort(semver.Collection(sortedVersionTags))

	for _, versionTag := range sortedVersionTags {
		if !v.ConstraintCheck.Check(
			versionTag,
		) {
			continue
		}
		matchedVersion = versionTag.String()
		break
	}

	if len(matchedVersion) == 0 {
		return "", nil
	}

	return matchedVersion, nil
}

// check if the given string is valid semver string and if yest
// create a constraint checker out of it
func isSemver(version string) bool {
	_, err := semver.NewVersion(version)
	return err == nil
}

func isValidSemverConstraint(version string) bool {
	versionRegex := `v?([0-9|x|X|\*]+)(\.[0-9|x|X|\*]+)?(\.[0-9|x|X|\*]+)?` +
		`(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` +
		`(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?`
	constraintOperations := `=||!=|>|<|>=|=>|<=|=<|~|~>|\^`
	validConstraintRegex := regexp.MustCompile(fmt.Sprintf(
		`^(\s*(%s)\s*(%s)\s*\,?)+$`,
		constraintOperations,
		versionRegex))

	return validConstraintRegex.MatchString(version)
}

// normalize the proxy url to
// - not have traling slashes
func normalizeUrl(url string) string {
	if strings.HasSuffix(url, "/") {
		ind := strings.LastIndex(url, "/")
		if ind == -1 {
			return url
		}
		return strings.Join([]string{url[:ind], "", url[ind+1:]}, "")
	}
	return url
}

// get the proxy url for the latest version
func getVersionLatestProxyURL(pkg string) string {
	urlPrefix := normalizeUrl(DEFAULT_PROXY_URL)
	return urlPrefix + "/" + pkg + "/@latest"
}

// get the proxy url for the entire version list
func getVersionListProxyURL(pkg string) string {
	urlPrefix := normalizeUrl(DEFAULT_PROXY_URL)
	return urlPrefix + "/" + pkg + "/@v/list"
}

func (g *GitHub) versions(owner, repo string) (versions []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	page := 1

	for {
		options := &github.ListOptions{
			Page:    page,
			PerPage: 100,
		}

		tags, _, err := g.Client.Repositories.ListTags(ctx, owner, repo, options)
		if err != nil {
			return nil, fmt.Errorf("listing tags: %w", err)
		}

		if len(tags) == 0 {
			break
		}

		for _, t := range tags {
			versions = append(versions, t.GetName())
		}

		page++
	}

	if len(versions) == 0 {
		return nil, errors.New("no versions defined")
	}

	return
}

// Resolve implementation.
func (g *GitHub) resolve(owner, repo, version string) (string, error) {
	// fetch tags
	tags, err := g.versions(owner, repo)
	if err != nil {
		return "", err
	}

	// convert to semver, ignoring malformed
	var versions []goSem.Version
	for _, t := range tags {
		if v, err := goSem.Parse(t); err == nil {
			versions = append(versions, v)
		}
	}

	// no versions, it has tags but they're not semver
	if len(versions) == 0 {
		return "", errors.New("no versions matched")
	}

	// master special-case
	if version == "master" {
		return versions[0].String(), nil
	}

	// match requested semver range
	vr, err := goSem.ParseRange(version)
	if err != nil {
		return "", fmt.Errorf("parsing version range: %w", err)
	}

	for _, v := range versions {
		if vr.Match(v) {
			return v.String(), nil
		}
	}

	return "", errors.New("no versions matched")
}
