package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Discogs is the cover-art fallback for albums the Cover Art Archive
// lacks. It needs a free personal access token from
// https://www.discogs.com/settings/developers in DISCOGS_TOKEN (either
// the environment or .env at the repo root).

func discogsToken(root string) string {
	if token := os.Getenv("DISCOGS_TOKEN"); token != "" {
		return token
	}
	content, err := os.ReadFile(filepath.Join(root, ".env"))
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(content), "\n") {
		if value, found := strings.CutPrefix(strings.TrimSpace(line), "DISCOGS_TOKEN="); found {
			return strings.Trim(value, `"'`)
		}
	}
	return ""
}

// discogsCover searches Discogs for the album and returns the first
// square-ish cover image, preferring master entries (the canonical
// album) over individual releases.
func discogsCover(root, title, artist string) ([]byte, error) {
	token := discogsToken(root)
	if token == "" {
		return nil, fmt.Errorf("DISCOGS_TOKEN is not set (in the environment or .env); create one at https://www.discogs.com/settings/developers")
	}

	for _, searchType := range []string{"master", "release"} {
		params := url.Values{
			"release_title": {title},
			"artist":        {artist},
			"type":          {searchType},
			"per_page":      {"10"},
			"token":         {token},
		}
		req, err := http.NewRequest("GET", "https://api.discogs.com/database/search?"+params.Encode(), nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", userAgent)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("Discogs search: %s: %s", resp.Status, body)
		}

		var data struct {
			Results []struct {
				CoverImage string `json:"cover_image"`
			} `json:"results"`
		}
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, err
		}
		for _, result := range data.Results {
			// spacer.gif is Discogs' image-less placeholder
			if result.CoverImage == "" || strings.HasSuffix(result.CoverImage, "spacer.gif") {
				continue
			}
			image, err := fetchImage(result.CoverImage)
			if err != nil {
				continue
			}
			if isSquarish(image) {
				return image, nil
			}
		}
	}
	return nil, fmt.Errorf("no square cover on Discogs for %q by %q", title, artist)
}
