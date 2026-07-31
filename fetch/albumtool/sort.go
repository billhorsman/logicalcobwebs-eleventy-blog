package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func topAlbumsPath(root string) string {
	return filepath.Join(root, "_data", "top_albums.json")
}

// cmdSort orders _data/top_albums.json by first release date (title
// breaks ties), matching how the film top 100 is sorted.
func cmdSort(root string) error {
	raw, err := os.ReadFile(topAlbumsPath(root))
	if err != nil {
		return err
	}
	var slugs []string
	if err := json.Unmarshal(raw, &slugs); err != nil {
		return err
	}

	type album struct {
		slug, date, title string
	}
	albums := make([]album, 0, len(slugs))
	for _, slug := range slugs {
		data, err := os.ReadFile(albumDataPath(root, slug))
		if err != nil {
			return err
		}
		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}
		albums = append(albums, album{slug, getString(m, "first-release-date"), getString(m, "title")})
	}

	sort.SliceStable(albums, func(i, j int) bool {
		if albums[i].date != albums[j].date {
			return albums[i].date < albums[j].date
		}
		return albums[i].title < albums[j].title
	})

	sorted := make([]string, len(albums))
	for i, a := range albums {
		sorted[i] = a.slug
	}

	// Same style as top_films.json: tab-indented with a trailing newline
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	if err := enc.Encode(sorted); err != nil {
		return err
	}
	if err := os.WriteFile(topAlbumsPath(root), buf.Bytes(), 0o644); err != nil {
		return err
	}
	fmt.Printf("Sorted %d albums\n", len(sorted))
	return nil
}
