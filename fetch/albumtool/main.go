// albumtool maintains the album data behind logicalcobwebs.com/albums,
// the sibling of filmtool. Data comes from musicbrainz.org (no API key
// needed) and cover art from coverartarchive.org.
//
//   - _data/albums/<slug>.json     release-group details
//   - content/bill/albums/covers/*.jpg  front covers
//
// _data/top_albums.json (the top 100) and _data/top_5_albums.json are
// curated by hand; `sort` keeps the top 100 in release-date order.
package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprint(os.Stderr, `Usage: albumtool <command> [arguments]

Commands:
  search <title> [by <artist>]  search MusicBrainz, pick a match, then add it
  add <mbid-or-url>   fetch one release group by MusicBrainz id or URL
  sort                sort _data/top_albums.json by first release date
  cover <slug> <release>  replace an album's cover with a specific release's front image
  covers              retry missing covers (Cover Art Archive, then Discogs)
  links               backfill direct Spotify links from MusicBrainz
  tracks [<slug> <release>]  backfill missing tracklists, or pin one edition's
  montage             regenerate the OpenGraph cover grid for the index
  playlist [name]     build a Spotify playlist of the top 100 (needs SPOTIFY_CLIENT_ID)
`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	root, err := findRoot()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	cmd, args := os.Args[1], os.Args[2:]
	switch cmd {
	case "search":
		err = cmdSearch(root, args)
	case "add":
		err = cmdAdd(root, args)
	case "sort":
		err = cmdSort(root)
	case "cover":
		err = cmdCover(root, args)
	case "covers":
		err = cmdCovers(root)
	case "links":
		err = cmdLinks(root)
	case "tracks":
		err = cmdTracks(root, args)
	case "montage":
		err = cmdMontage(root)
	case "playlist":
		err = cmdPlaylist(root, args)
	case "help", "-h", "--help":
		usage()
		return
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n", cmd)
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
