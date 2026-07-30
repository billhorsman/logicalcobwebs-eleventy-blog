package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func topFilmsPath(root string) string {
	return filepath.Join(root, "_data", "top_films.json")
}

func readTopFilmSlugs(root string) ([]string, error) {
	raw, err := os.ReadFile(topFilmsPath(root))
	if err != nil {
		return nil, err
	}
	var slugs []string
	if err := json.Unmarshal(raw, &slugs); err != nil {
		return nil, err
	}
	return slugs, nil
}

// cmdSort is the port of sort.rb: order top_films.json by year then title.
func cmdSort(root string) error {
	slugs, err := readTopFilmSlugs(root)
	if err != nil {
		return err
	}
	films := make([]*Film, 0, len(slugs))
	for _, slug := range slugs {
		film, err := loadFilm(root, slug)
		if err != nil {
			return err
		}
		films = append(films, film)
	}

	sort.SliceStable(films, func(i, j int) bool {
		if films[i].Year != films[j].Year {
			return films[i].Year < films[j].Year
		}
		return films[i].Title < films[j].Title
	})

	sorted := make([]string, len(films))
	for i, film := range films {
		sorted[i] = film.Slug
	}

	// top_films.json is tab-indented with a trailing newline (unlike the
	// generated data files), so keep that style to avoid diff noise.
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	if err := enc.Encode(sorted); err != nil {
		return err
	}
	if err := os.WriteFile(topFilmsPath(root), buf.Bytes(), 0o644); err != nil {
		return err
	}
	fmt.Printf("Sorted %d films\n", len(sorted))
	return nil
}
