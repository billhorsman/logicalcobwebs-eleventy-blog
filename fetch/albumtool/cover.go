package main

import (
	"fmt"
	"os"
)

// cmdCover replaces an album's cover with a specific release's front
// image, for when the automatic pick lands on the wrong edition:
//
//	albumtool cover snap-1983 https://musicbrainz.org/release/<mbid>
func cmdCover(root string, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: albumtool cover <slug> <release-mbid-or-url>")
	}
	slug, arg := args[0], args[1]
	if _, err := os.Stat(albumDataPath(root, slug)); err != nil {
		return fmt.Errorf("unknown album %q (no _data/albums/%s.json)", slug, slug)
	}
	mbid := mbidPattern.FindString(arg)
	if mbid == "" {
		return fmt.Errorf("%q is not a MusicBrainz release id or URL", arg)
	}

	body, err := fetchImage("https://coverartarchive.org/release/" + mbid + "/front-500")
	if err != nil {
		return err
	}
	if !isSquarish(body) {
		fmt.Println("Warning: this cover is not square either")
	}
	return writeCover(coverPath(root, slug), body)
}
