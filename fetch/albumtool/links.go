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

// Each album links directly to Spotify. MusicBrainz releases carry
// streaming URL relationships we can mine for the album id.

var spotifyAlbum = regexp.MustCompile(`open\.spotify\.com/album/([A-Za-z0-9]+)`)

// spotifyURL finds a Spotify album URL among the release group's
// releases. Empty when the group has no Spotify relationship.
func spotifyURL(mbid string) (string, error) {
	data, err := mbGet("release", url.Values{
		"release-group": {mbid},
		"inc":           {"url-rels"},
		"limit":         {"100"},
	})
	if err != nil {
		return "", err
	}
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
			if match := spotifyAlbum.FindStringSubmatch(getString(urlMap, "resource")); match != nil {
				return "https://open.spotify.com/album/" + match[1], nil
			}
		}
	}
	return "", nil
}

// cmdLinks backfills the spotify field for every album that lacks one,
// migrating old album.link values where the Spotify id is already known.
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
		data, err := readAlbumData(root, slug)
		if err != nil {
			return err
		}

		oldLink := getString(data, "album-link")
		if getString(data, "spotify") != "" && oldLink == "" {
			continue
		}
		delete(data, "album-link")

		spotify := getString(data, "spotify")
		if spotify == "" {
			// album.link/s/<id> URLs already carry the Spotify id
			if id, found := strings.CutPrefix(oldLink, "https://album.link/s/"); found {
				spotify = "https://open.spotify.com/album/" + id
			}
		}
		if spotify == "" {
			spotify, err = spotifyURL(getString(data, "id"))
			// MusicBrainz asks for at most one request per second
			time.Sleep(1100 * time.Millisecond)
			if err != nil {
				fmt.Printf("Skipping %s: %v\n", slug, err)
				missing++
				continue
			}
		}

		if spotify == "" {
			fmt.Printf("No Spotify link for %s\n", slug)
			missing++
		} else {
			data["spotify"] = spotify
			fmt.Printf("%s -> %s\n", slug, spotify)
		}
		if err := writeAlbumData(root, slug, data); err != nil {
			return err
		}
	}
	if missing > 0 {
		fmt.Printf("%d album(s) without Spotify links\n", missing)
	}
	return nil
}
