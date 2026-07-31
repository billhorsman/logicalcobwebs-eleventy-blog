# albumtool

The sibling of filmtool, for the (work in progress, unlinked) /albums
section. Album data comes from [MusicBrainz](https://musicbrainz.org/)
release groups and covers from the Cover Art Archive, with Discogs as a
cover fallback. MusicBrainz needs no key; for Discogs, put a personal
access token (https://www.discogs.com/settings/developers) in `.env`
as `DISCOGS_TOKEN=...`. Run `./bin/albumtool covers` to retry albums
with missing covers.

Build:

```sh
go -C fetch/albumtool build -o ../../bin/albumtool .
```

## Adding an album

```sh
./bin/albumtool search little girl blue
./bin/albumtool search wish you were here by pink floyd
./bin/albumtool add https://musicbrainz.org/release-group/19ca48c1-af98-3000-bc9a-97ef5fa3e3be
```

Both write `_data/albums/<slug>.json` and download the front cover to
`content/bill/albums/covers/<slug>.jpg`. A search with a single match skips
the picker. While the top 100 has fewer than 100 entries, added albums
join it automatically (kept in release-date order); once it's full,
swaps are made by hand.

## Better artwork

If a cover is wrong, ugly, or missing:

- **Pin a specific edition's art.** Find the release on musicbrainz.org
  whose cover you want (the release page's Cover Art tab shows it) and:

  ```sh
  ./bin/albumtool cover <slug> https://musicbrainz.org/release/<mbid>
  ```

- **Re-run the automatic hunt.** Delete
  `content/bill/albums/covers/<slug>.jpg`, then:

  ```sh
  ./bin/albumtool covers
  ```

  This retries every album with a missing cover: the Cover Art
  Archive first (preferring square images), then Discogs (needs
  `DISCOGS_TOKEN` in `.env`).

- **Manual override.** Drop any square JPEG at
  `content/bill/albums/covers/<slug>.jpg` — the tool never overwrites
  an existing cover.

The index crops covers to square in CSS, so a slightly-off scan only
loses a sliver at the edges.

## Notes

Drop a markdown file at `_data/album_notes/<slug>.md` and it renders
above the tracklist on that album's page. No file, no note.

## Tracklists

Album pages list the tracks of the earliest official release. To use a
different edition's tracklist:

```sh
./bin/albumtool tracks <slug> https://musicbrainz.org/release/<mbid>
```

`./bin/albumtool tracks` (no arguments) backfills albums without one.

## OpenGraph montage

The albums index shares as a grid of 18 covers sampled across the top
100. Regenerate it after the list changes much:

```sh
./bin/albumtool montage
```

## Curating the top 100 (and top 5)

Like films, the lists are curated by hand:

- `_data/top_albums.json` — the top 100, kept in release-date order by
  `./bin/albumtool sort`
- `_data/top_5_albums.json` — the top 5, unranked and equal; `sort` keeps
  it in release-date order too

Album pages and the /bill/albums index render from these lists via
Eleventy pagination — there is no generate step.
