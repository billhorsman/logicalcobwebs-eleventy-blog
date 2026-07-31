package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// album.link (Odesli) pages link out to every streaming service, but
// are addressed by platform id. MusicBrainz releases carry streaming
// URL relationships we can mine for one.

var (
	spotifyAlbum = regexp.MustCompile(`open\.spotify\.com/album/([A-Za-z0-9]+)`)
	appleAlbum   = regexp.MustCompile(`music\.apple\.com/[a-z]{2}/album/(?:[^/]+/)?(\d+)`)
	deezerAlbum  = regexp.MustCompile(`deezer\.com/album/(\d+)`)
)

// albumLink finds a streaming id among the release group's releases and
// returns an album.link URL, preferring Spotify, then Apple, then
// Deezer. Empty when the group has no streaming relationships.
func albumLink(mbid string) (string, error) {
	data, err := mbGet("release", url.Values{
		"release-group": {mbid},
		"inc":           {"url-rels"},
		"limit":         {"100"},
	})
	if err != nil {
		return "", err
	}

	var spotify, apple, deezer string
	releases, _ := data["releases"].([]any)
	for _, r := range releases {
		m, ok := r.(map[string]any)
		if !ok {
			continue
		}
		relations, _ := m["relations"].([]any)
		for _, rel := range relations {
			relMap, ok := rel.(map[string]any)
			if !ok {
				continue
			}
			urlMap, _ := relMap["url"].(map[string]any)
			resource := getString(urlMap, "resource")
			if match := spotifyAlbum.FindStringSubmatch(resource); match != nil && spotify == "" {
				spotify = match[1]
			} else if match := appleAlbum.FindStringSubmatch(resource); match != nil && apple == "" {
				apple = match[1]
			} else if match := deezerAlbum.FindStringSubmatch(resource); match != nil && deezer == "" {
				deezer = match[1]
			}
		}
	}

	switch {
	case spotify != "":
		return "https://album.link/s/" + spotify, nil
	case apple != "":
		return "https://album.link/i/" + apple, nil
	case deezer != "":
		return "https://album.link/d/" + deezer, nil
	}
	return "", nil
}

// cmdLinks backfills album-link for every album that lacks one.
func cmdLinks(root string) error {
	entries, err := os.ReadDir(filepath.Join(root, "_data", "albums"))
	if err != nil {
		return err
	}

	missing := 0
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		slug := strings.TrimSuffix(entry.Name(), ".json")

		f, err := os.Open(albumDataPath(root, slug))
		if err != nil {
			return err
		}
		data, err := decodeJSON(f)
		f.Close()
		if err != nil {
			return err
		}
		if getString(data, "album-link") != "" {
			continue
		}

		link, err := albumLink(getString(data, "id"))
		// MusicBrainz asks for at most one request per second
		time.Sleep(1100 * time.Millisecond)
		if err != nil {
			// Transient (MusicBrainz 503s when busy) — skip; a re-run
			// picks up whatever is still missing
			fmt.Printf("Skipping %s: %v\n", slug, err)
			missing++
			continue
		}
		if link == "" {
			fmt.Printf("No streaming link for %s\n", slug)
			missing++
			continue
		}

		data["album-link"] = link
		out, err := marshalPretty(data)
		if err != nil {
			return err
		}
		if err := os.WriteFile(albumDataPath(root, slug), out, 0o644); err != nil {
			return err
		}
		fmt.Printf("%s -> %s\n", slug, link)
	}
	if missing > 0 {
		fmt.Printf("%d album(s) without streaming links\n", missing)
	}
	return nil
}
