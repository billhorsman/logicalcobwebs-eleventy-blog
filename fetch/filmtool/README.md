# filmtool

Go rewrite of the Ruby scripts (`fetch.rb`, `generate.rb`, `sort.rb`) that
maintain the film data behind logicalcobwebs.com.

## Setup

Create a TMDB read access token at https://www.themoviedb.org/settings/api
and export it:

```sh
export TMDB_API_TOKEN="eyJ..."
```

Run from anywhere in the repo:

```sh
go run ./fetch/filmtool <command>
```

## Adding a film for a blog review (not in the top 100)

Either search by title and pick from the matches:

```sh
go run ./fetch/filmtool search the last viking
```

or paste a TMDB id or URL directly:

```sh
go run ./fetch/filmtool add https://www.themoviedb.org/movie/1295400-den-sidste-viking
```

Both fetch the film's details and credits into `_data/films/<slug>.json`,
download the poster, backdrop, and cast profile photos, and print a blog
post scaffold to paste into `content/blog/<year>/<name>/index.md`. The top
100 list is untouched.

## Maintaining the top 100

```sh
go run ./fetch/filmtool sync       # pull films from the TMDB list (add -force to refetch)
go run ./fetch/filmtool sort       # order _data/top_films.json by year then title
go run ./fetch/filmtool generate   # regenerate content/bill/films/*.md and top cast/director data
```

`generate` reproduces the Ruby output byte-for-byte, so running it after a
sync only changes what actually changed.

## Differences from the Ruby scripts

- The API token comes from `$TMDB_API_TOKEN` instead of being hardcoded.
- HTTP errors are reported instead of silently parsed as JSON.
- `fetch/people/*.json` are no longer fetched — the site only ever used the
  profile images, not the person details.
- Newly written `_data/films/*.json` files have alphabetically sorted keys
  (Go's JSON encoder); existing files keep TMDB's field order until
  refetched with `-force`. Both read the same.
