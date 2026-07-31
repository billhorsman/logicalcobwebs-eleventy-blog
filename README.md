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

2. The command prints a blog post scaffold. Paste it into
   `content/blog/<year>/<film-name>/index.md`, write the review, and set
   the star rating in `{% outOfFive 5 %}`.

3. Check it locally with `npm start`. The cast grid and film details
   render from the fetched data — no generation step needed.

A reviewed film does not have to be in the top 100.

## Rebuilding the top 100

The top 100 is driven by a [TMDB list](https://www.themoviedb.org/list/8291691).
After changing the list there:

```sh
./bin/filmtool sync        # fetch data + images for new films on the list
./bin/filmtool sort        # order _data/top_films.json by release date
./bin/filmtool generate    # regenerate content/bill/films/*.md pages
```

Then review the diff and commit. `sort` renumbers the films, so run
`generate` after it to keep the page numbering consistent.

To show a review quote on a film's top-100 page, add a JSON file to
`_data/films/reviews/<slug>.json`.
