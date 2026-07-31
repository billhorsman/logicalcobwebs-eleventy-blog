# logicalcobwebs.com

The Logical Cobwebs blog — built with [Eleventy](https://www.11ty.dev/),
deployed on Netlify. Started from
[eleventy-base-blog](https://github.com/11ty/eleventy-base-blog).

## Development

```sh
npm install
npm start        # dev server with live reload
npm run build    # production build to _site/
```

Requires Node 22+ (see `.nvmrc`) and Go (for the film tooling below).

## Film tooling

`bin/filmtool` maintains the film data (`_data/films/`), images
(`content/bill/films/`), and the generated top-100 pages. Build it with:

```sh
go -C fetch/filmtool build -o ../../bin/filmtool .
```

It needs a TMDB read access token (https://www.themoviedb.org/settings/api),
read from $TMDB_API_TOKEN or from a line in `.env` at the repo root
(gitignored):

```
TMDB_API_TOKEN=eyJ...
```

See [fetch/filmtool/README.md](fetch/filmtool/README.md) for full details.

## Adding a film review blog post

1. Fetch the film's data and images — search by title and pick from the
   matches:

   ```sh
   ./bin/filmtool search the last viking
   ```

   or paste a TMDB id or URL if you already have it:

   ```sh
   ./bin/filmtool add https://www.themoviedb.org/movie/1295400-den-sidste-viking
   ```

2. A scaffold post is created at `content/blog/<year>/<film-name>.md`
   (unless a review already exists). Write the review and set `stars:`
   in the frontmatter.

3. Check it locally with `npm start`. The cast grid and film details
   render from the fetched data — no generation step needed.

A reviewed film does not have to be in the top 100.

## Rebuilding the top 100

Film data comes from a [TMDB list](https://www.themoviedb.org/list/8291691)
holding every film of interest (not just the top 100); the top 100 itself
is the hand-curated `_data/top_films.json`. To change it:

1. Add the film to the TMDB list (if it isn't already on it).
2. Add its slug to `_data/top_films.json` (anywhere — it gets sorted),
   and remove one to keep it at 100.
3. Run:

   ```sh
   ./bin/filmtool top-100
   ```

This syncs film data and images from the TMDB list, sorts
`_data/top_films.json` by release date, and regenerates the film pages
with the new numbering. Review the diff and commit. (The steps can also
be run individually: `sync`, `sort`, `generate`.)

To show a review quote on a film's top-100 page, add a JSON file to
`_data/films/reviews/<slug>.json`.
