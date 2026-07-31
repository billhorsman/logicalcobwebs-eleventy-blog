package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"sort"
)

var mbidPattern = regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)

// parseMBID accepts a bare MusicBrainz id or a release-group URL
// ("https://musicbrainz.org/release-group/19ca48c1-...").
func parseMBID(arg string) (string, error) {
	if m := mbidPattern.FindString(arg); m != "" {
		return m, nil
	}
	return "", fmt.Errorf("%q is not a MusicBrainz release-group id or URL", arg)
}

func cmdAdd(root string, args []string) error {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	force := fs.Bool("force", false, "refetch album data even if it already exists")
	fs.Parse(args)
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: albumtool add [-force] <mbid-or-url>")
	}
	mbid, err := parseMBID(fs.Arg(0))
	if err != nil {
		return err
	}
	return addAlbum(root, mbid, *force)
}

// addAlbum fetches one release group's data and cover art. The curated
// top-100 list (_data/top_albums.json) is never touched.
func addAlbum(root, mbid string, force bool) error {
	data, err := mbGet("release-group/"+mbid, url.Values{"inc": {"artist-credits+genres"}})
	if err != nil {
		return err
	}

	title := getString(data, "title")
	year := yearOf(getString(data, "first-release-date"))
	slug := slugify(title, year)
	credits, _ := data["artist-credit"].([]any)

	data["slug"] = slug
	data["year"] = year
	data["artist"] = artistName(credits)
	data["top-genres"] = topGenres(data, 5)

	path := albumDataPath(root, slug)
	if _, err := os.Stat(path); os.IsNotExist(err) || force {
		out, err := marshalPretty(data)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(albumDataPath(root, ""), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(path, out, 0o644); err != nil {
			return err
		}
		fmt.Printf("Wrote _data/albums/%s.json\n", slug)
	} else {
		fmt.Printf("_data/albums/%s.json already exists (use -force to refetch)\n", slug)
	}

	if err := downloadCover(mbid, coverPath(root, slug)); err != nil {
		return err
	}

	fmt.Printf("\n%s — %s (%s)\nAdd %q to _data/top_albums.json (and _data/top_5_albums.json?) to show it.\n",
		title, data["artist"], year, slug)
	return nil
}

// topGenres returns the release group's genre names, most-voted first.
func topGenres(data map[string]any, limit int) []string {
	genres, _ := data["genres"].([]any)
	type genre struct {
		name  string
		count int64
	}
	list := make([]genre, 0, len(genres))
	for _, g := range genres {
		m, ok := g.(map[string]any)
		if !ok {
			continue
		}
		count, _ := m["count"].(interface{ Int64() (int64, error) })
		var n int64
		if count != nil {
			n, _ = count.Int64()
		}
		list = append(list, genre{getString(m, "name"), n})
	}
	sort.SliceStable(list, func(i, j int) bool { return list[i].count > list[j].count })
	names := make([]string, 0, limit)
	for _, g := range list[:min(limit, len(list))] {
		names = append(names, g.name)
	}
	return names
}
