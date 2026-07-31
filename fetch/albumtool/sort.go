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

func top5AlbumsPath(root string) string {
	return filepath.Join(root, "_data", "top_5_albums.json")
}

func readSlugList(path string) ([]string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var slugs []string
	if err := json.Unmarshal(raw, &slugs); err != nil {
		return nil, err
	}
	return slugs, nil
}

// writeSlugList matches top_films.json's style: tab-indented with a
// trailing newline.
func writeSlugList(path string, slugs []string) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	if err := enc.Encode(slugs); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0o644)
}

func cmdSort(root string) error {
	for _, path := range []string{topAlbumsPath(root), top5AlbumsPath(root)} {
		n, err := sortAlbumList(root, path)
		if err != nil {
			return err
		}
		fmt.Printf("Sorted %d albums in %s\n", n, filepath.Base(path))
	}
	return nil
}

// sortAlbumList orders a slug-list file by first release date (title
// breaks ties), matching how the film top 100 is sorted.
func sortAlbumList(root, path string) (int, error) {
	slugs, err := readSlugList(path)
	if err != nil {
		return 0, err
	}

	type album struct {
		slug, date, title string
	}
	albums := make([]album, 0, len(slugs))
	for _, slug := range slugs {
		data, err := os.ReadFile(albumDataPath(root, slug))
		if err != nil {
			return 0, err
		}
		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			return 0, err
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
	return len(sorted), writeSlugList(path, sorted)
}
