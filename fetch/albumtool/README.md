# albumtool

The sibling of filmtool, for the (work in progress, unlinked) /albums
section. Album data comes from [MusicBrainz](https://musicbrainz.org/)
release groups and covers from the Cover Art Archive. No API key needed.

Build:

```sh
go -C fetch/albumtool build -o ../../bin/albumtool .
```

## Adding an album

```sh
./bin/albumtool search little girl blue
./bin/albumtool add https://musicbrainz.org/release-group/19ca48c1-af98-3000-bc9a-97ef5fa3e3be
```

Both write `_data/albums/<slug>.json` and download the front cover to
`content/albums/covers/<slug>.jpg`. A search with a single match skips
the picker.

## Curating the top 100 (and top 5)

Like films, the lists are curated by hand:

- `_data/top_albums.json` — the top 100, kept in release-date order by
  `./bin/albumtool sort`
- `_data/top_5_albums.json` — the top 5, in rank order (position 1 is
  the best album)

Album pages and the /albums index render from these lists via Eleventy
pagination — there is no generate step. The pages are noindexed and
excluded from collections (sitemap, feeds) while the section is WIP.
