package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	mbBase       = "https://musicbrainz.org/ws/2/"
	coverArtBase = "https://coverartarchive.org/release-group/"
	// MusicBrainz requires an identifying User-Agent
	userAgent = "logicalcobwebs-albumtool/1.0 (https://logicalcobwebs.com)"
)

func findRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "eleventy.config.js")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("not inside the blog repository (eleventy.config.js not found)")
		}
		dir = parent
	}
}

func albumDataPath(root, slug string) string {
	return filepath.Join(root, "_data", "albums", slug+".json")
}

func coverPath(root, slug string) string {
	return filepath.Join(root, "content", "albums", "covers", slug+".jpg")
}

func mbGet(endpoint string, params url.Values) (map[string]any, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("fmt", "json")

	req, err := http.NewRequest("GET", mbBase+endpoint+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %s: %s", endpoint, resp.Status, body)
	}

	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	var data map[string]any
	if err := dec.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// downloadCover fetches the 500px front cover for a release group. The
// release-group endpoint needs a designated front image (and can be
// flaky), so on failure it falls back to trying the covers of the
// group's individual releases.
func downloadCover(mbid, dest string) error {
	if _, err := os.Stat(dest); err == nil {
		return nil
	}
	if fetchImage(coverArtBase+mbid+"/front-500", dest) == nil {
		return nil
	}

	releasesData, err := mbGet("release", url.Values{"release-group": {mbid}, "limit": {"25"}})
	if err != nil {
		return err
	}
	releases, _ := releasesData["releases"].([]any)
	for _, r := range releases {
		m, ok := r.(map[string]any)
		if !ok {
			continue
		}
		releaseID := getString(m, "id")
		if fetchImage("https://coverartarchive.org/release/"+releaseID+"/front-500", dest) == nil {
			fmt.Printf("Cover art from release %s\n", releaseID)
			return nil
		}
	}
	fmt.Printf("No cover art found for %s\n", mbid)
	return nil
}

func fetchImage(imageURL, dest string) error {
	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: %s", imageURL, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	fmt.Printf("Writing %d bytes to %s\n", len(body), dest)
	return os.WriteFile(dest, body, 0o644)
}

func marshalPretty(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}

var nonSlugChars = regexp.MustCompile(`[^a-z0-9 ]+`)

// slugify matches filmtool's slug rules: lowercase ascii, spaces to
// dashes, then the release year.
func slugify(title, year string) string {
	s := strings.ToLower(title)
	s = nonSlugChars.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "-")
	return s + "-" + year
}

func yearOf(date string) string {
	year, _, _ := strings.Cut(date, "-")
	return year
}

func getString(m map[string]any, key string) string {
	s, _ := m[key].(string)
	return s
}

// artistOverrides corrects credited names MusicBrainz preserves
// as-printed-at-release, e.g. an artist's dead name.
var artistOverrides = map[string]string{
	"Kate Tempest": "Kae Tempest",
}

// artistName joins a MusicBrainz artist-credit list into a display name,
// e.g. "Simon & Garfunkel" or "David Bowie & Queen".
func artistName(credits []any) string {
	var out strings.Builder
	for _, c := range credits {
		m, ok := c.(map[string]any)
		if !ok {
			continue
		}
		name := getString(m, "name")
		if override, ok := artistOverrides[name]; ok {
			name = override
		}
		out.WriteString(name)
		out.WriteString(getString(m, "joinphrase"))
	}
	return out.String()
}
