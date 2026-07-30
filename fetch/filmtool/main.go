// filmtool maintains the film data behind logicalcobwebs.com.
//
// It talks to api.themoviedb.org (TMDB) and keeps three things up to date:
//
//   - _data/films/<slug>.json        film details + credits
//   - content/bill/films/{posters,backdrops,profiles}/*.jpg
//   - content/bill/films/<slug>.md   generated top-100 pages
//
// A TMDB read access token must be set in the TMDB_API_TOKEN environment
// variable (https://www.themoviedb.org/settings/api).
package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprint(os.Stderr, `Usage: filmtool <command> [arguments]

Commands:
  search <title>      search TMDB by title, pick a match, then add it
  add <id-or-url>     fetch one film by TMDB id or URL (for blog-only reviews)
  sync [-force]       refresh the top-100 films from the TMDB list
  generate            regenerate top-100 markdown pages and top cast/director data
  sort                sort _data/top_films.json by year then title

add and search fetch film data and images without touching the top 100,
and print a blog post scaffold for a review. Use sync + sort + generate
when the top-100 TMDB list itself changes.
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
	case "sync":
		err = cmdSync(root, args)
	case "generate":
		err = cmdGenerate(root)
	case "sort":
		err = cmdSort(root)
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
