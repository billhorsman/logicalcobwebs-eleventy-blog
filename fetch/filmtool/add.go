package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

var movieURLPattern = regexp.MustCompile(`themoviedb\.org/movie/(\d+)`)

// parseMovieID accepts a bare TMDB id ("1295400") or a TMDB movie URL
// ("https://www.themoviedb.org/movie/1295400-den-sidste-viking").
func parseMovieID(arg string) (int, error) {
	if id, err := strconv.Atoi(arg); err == nil {
		return id, nil
	}
	if m := movieURLPattern.FindStringSubmatch(arg); m != nil {
		return strconv.Atoi(m[1])
	}
	return 0, fmt.Errorf("%q is not a TMDB movie id or URL", arg)
}

func cmdAdd(root string, args []string) error {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	force := fs.Bool("force", false, "refetch film data even if it already exists")
	fs.Parse(args)
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: filmtool add [-force] <tmdb-id-or-url>")
	}
	id, err := parseMovieID(fs.Arg(0))
	if err != nil {
		return err
	}
	return addFilm(root, id, *force)
}

// addFilm fetches one film's data and images. It never touches the top-100
// list, so it is safe for films that are only referenced from blog reviews.
func addFilm(root string, id int, force bool) error {
	data, err := apiGetJSON("movie/"+strconv.Itoa(id), url.Values{"append_to_response": {"credits"}})
	if err != nil {
		return err
	}

	title := getString(data, "title")
	year := yearOf(getString(data, "release_date"))
	slug := slugify(title, year)

	data["slug"] = slug
	data["year"] = year
	data["language"] = englishLanguageName(data)

	path := filmDataPath(root, slug)
	if _, err := os.Stat(path); os.IsNotExist(err) || force {
		out, err := marshalPretty(data)
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, out, 0o644); err != nil {
			return err
		}
		fmt.Printf("Wrote _data/films/%s.json\n", slug)
	} else {
		fmt.Printf("_data/films/%s.json already exists (use -force to refetch)\n", slug)
	}

	if err := downloadFilmImages(root, slug, getString(data, "poster_path"), getString(data, "backdrop_path")); err != nil {
		return err
	}

	film, err := loadFilm(root, slug)
	if err != nil {
		return err
	}
	if err := downloadProfileImages(root, film); err != nil {
		return err
	}

	printBlogScaffold(title, slug, year)
	return nil
}

// englishLanguageName resolves the film's original language to its English
// name, e.g. "da" → "Danish".
func englishLanguageName(data map[string]any) any {
	original := getString(data, "original_language")
	languages, _ := data["spoken_languages"].([]any)
	for _, l := range languages {
		if m, ok := l.(map[string]any); ok && getString(m, "iso_639_1") == original {
			return getString(m, "english_name")
		}
	}
	return nil
}

func downloadFilmImages(root, slug, posterPath, backdropPath string) error {
	for kind, tmdbPath := range map[string]string{"posters": posterPath, "backdrops": backdropPath} {
		if tmdbPath == "" {
			continue
		}
		if err := downloadImage(tmdbPath, imagePath(root, kind, slug+".jpg")); err != nil {
			return err
		}
	}
	return nil
}

// downloadProfileImages saves profile photos for the top 12 cast members
// and the director(s) — the people shown on cast cards.
func downloadProfileImages(root string, film *Film) error {
	people := append(film.TopCast(), film.Directors()...)
	for _, p := range people {
		if p.ProfilePath == "" {
			continue
		}
		dest := imagePath(root, "profiles", p.ID.String()+".jpg")
		if err := downloadImage(p.ProfilePath, dest); err != nil {
			return err
		}
	}
	return nil
}

func printBlogScaffold(title, slug, year string) {
	today := time.Now().Format("2006-01-02")
	fmt.Printf(`
Done. Blog post scaffold for %s (e.g. content/blog/%s/%s.md):

---
layout: layouts/film-review.njk
filmSlug: %s
date: %s
tags: Film review
stars: 5
---

Your review here.

Title, description, and share image derive from the film data; add a
dca: block (date, cinema, seat, rating) for a ticket stub.
`, title, time.Now().Format("2006"), trimYearSuffix(slug, year), slug, today)
}

func trimYearSuffix(slug, year string) string {
	if suffix := "-" + year; len(slug) > len(suffix) {
		if slug[len(slug)-len(suffix):] == suffix {
			return slug[:len(slug)-len(suffix)]
		}
	}
	return slug
}
