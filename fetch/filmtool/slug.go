package main

import (
	"regexp"
	"strings"
)

var nonSlugChars = regexp.MustCompile(`[^a-z0-9 ]+`)

// slugify builds the film slug the same way the Ruby scripts did:
// lowercase, strip everything but ascii letters/digits/spaces (accented
// characters are dropped, e.g. "Amélie" → "amlie"), spaces to dashes,
// then the release year. Existing slugs depend on this stying stable.
func slugify(title, year string) string {
	s := strings.ToLower(title)
	s = nonSlugChars.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "-")
	return s + "-" + year
}

// yearOf extracts the year from a TMDB release_date (YYYY-MM-DD).
func yearOf(releaseDate string) string {
	year, _, _ := strings.Cut(releaseDate, "-")
	return year
}
