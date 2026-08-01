# filmtool

Go rewrite of the Ruby scripts (`fetch.rb`, `generate.rb`, `sort.rb`) that
maintain the film data behind logicalcobwebs.com.

## Setup

Create a TMDB read access token at https://www.themoviedb.org/settings/api
and put it in `.env` at the repo root (gitignored), or export it:

```
TMDB_API_TOKEN=eyJ...
```

Build the binary (it's gitignored, so rebuild after pulling changes):

```sh
go -C fetch/filmtool build -o ../../bin/filmtool .
```

Run from the repo root:

```sh
./bin/filmtool <command>
```

## Adding a film for a blog review (not in the top 100)

Either search by title and pick from the matches:

```sh
./bin/filmtool search the last viking
```

or paste a TMDB id or URL directly:

```sh
./bin/filmtool add https://www.themoviedb.org/movie/1295400-den-sidste-viking
```

Both fetch the film's details and credits into `_data/films/<slug>.json`,
download the poster, backdrop, and cast profile photos, and create a
scaffold post at `content/blog/<year>/<name>.md` (unless the film already
has a review). A search with a single match skips the picker. The top
100 list is untouched.

## Maintaining the top 100

The TMDB list holds every film of interest; the top 100 itself is the
hand-curated `_data/top_films.json` (edit it by hand, `sort` keeps it
in release-date order).

```sh
./bin/filmtool top-100    # sync from the TMDB list and sort (add -force to refetch)
```

## Differences from the Ruby scripts

- The API token comes from `$TMDB_API_TOKEN` instead of being hardcoded.
- HTTP errors are reported instead of silently parsed as JSON.
- `fetch/people/*.json` are no longer fetched — the site only ever used the
  profile images, not the person details.
- Newly written `_data/films/*.json` files have alphabetically sorted keys
  (Go's JSON encoder); existing files keep TMDB's field order until
  refetched with `-force`. Both read the same.
