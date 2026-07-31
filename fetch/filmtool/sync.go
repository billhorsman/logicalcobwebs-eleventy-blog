package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
)

// The TMDB list holding the top 100 films: https://www.themoviedb.org/list/8291691
const topFilmsListID = 8291691

// cmdTop100 rebuilds the top 100 in one go: sync, sort, then generate.
func cmdTop100(root string, args []string) error {
	if err := cmdSync(root, args); err != nil {
		return err
	}
	if err := cmdSort(root); err != nil {
		return err
	}
	return cmdGenerate(root)
}

// cmdSync is the port of fetch.rb: it pulls every film on the TMDB top-100
// list and makes sure its data file and images exist locally.
func cmdSync(root string, args []string) error {
	fs := flag.NewFlagSet("sync", flag.ExitOnError)
	force := fs.Bool("force", false, "refetch film data even if it already exists")
	listID := fs.Int("list", topFilmsListID, "TMDB list id")
	fs.Parse(args)

	var items []map[string]any
	for page := 1; page <= 10; page++ {
		fmt.Printf("Fetching page %d\n", page)
		data, err := apiGetJSON("list/"+strconv.Itoa(*listID), url.Values{"page": {strconv.Itoa(page)}})
		if err != nil {
			return err
		}
		pageItems, _ := data["items"].([]any)
		if len(pageItems) == 0 {
			break
		}
		for _, item := range pageItems {
			if m, ok := item.(map[string]any); ok {
				items = append(items, m)
			}
		}
	}

	sort.SliceStable(items, func(i, j int) bool {
		return getString(items[i], "release_date") < getString(items[j], "release_date")
	})

	var slugs []string
	for _, item := range items {
		title := getString(item, "title")
		year := yearOf(getString(item, "release_date"))
		slug := slugify(title, year)
		slugs = append(slugs, slug)

		path := filmDataPath(root, slug)
		if _, err := os.Stat(path); os.IsNotExist(err) || *force {
			fmt.Printf("Fetching details for %s\n", title)
			id, ok := item["id"].(interface{ String() string })
			if !ok {
				return fmt.Errorf("film %q has no id", title)
			}
			movieID, err := strconv.Atoi(id.String())
			if err != nil {
				return err
			}
			if err := fetchFilmData(root, movieID, slug, year); err != nil {
				return err
			}
		}

		if err := downloadFilmImages(root, slug, getString(item, "poster_path"), getString(item, "backdrop_path")); err != nil {
			return err
		}
	}

	// Note: the TMDB list holds every film of interest, not just the top
	// 100 — _data/top_films.json is curated by hand and never written here.

	// Profile photos for everyone shown on cast cards, across all films.
	for _, slug := range slugs {
		film, err := loadFilm(root, slug)
		if err != nil {
			return err
		}
		if err := downloadProfileImages(root, film); err != nil {
			return err
		}
	}

	fmt.Printf("Synced %d films\n", len(slugs))
	return nil
}

// fetchFilmData fetches details+credits for one film and writes its data
// file, using the slug/year derived from the list entry.
func fetchFilmData(root string, id int, slug, year string) error {
	data, err := apiGetJSON("movie/"+strconv.Itoa(id), url.Values{"append_to_response": {"credits"}})
	if err != nil {
		return err
	}
	data["slug"] = slug
	data["year"] = year
	data["language"] = englishLanguageName(data)
	out, err := marshalPretty(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filmDataPath(root, slug), out, 0o644)
}
