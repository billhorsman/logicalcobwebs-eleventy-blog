---
title: Albums
permalink: /bill/albums/
ogImage: content/bill/albums/montage.jpg
eleventyComputed:
  description: "Bill Horsman's top {{ top_albums.length }} albums, from {{ albums[top_albums[0]].year }} to {{ albums[top_albums[top_albums.length - 1]].year }}."
---

<h1 class="visually-hidden">Albums</h1>

## Favourite Albums

It's very hard to pick a top {{ top_albums.length }} albums, so this is me throwing in some to start with. I'm expecting to add more as I go (and to remove those that no longer make the cut).

Only {{ top_artists.length }} artists appear more than once: {% namesAndCounts top_artists, 10, "albums" %}.

I might have to pick a top 100 songs as well. There are lots of songs that I love aren't in this list. For example, I like [One Bourbon, One Scotch and One Beer](https://open.spotify.com/track/2dp14VWbIxOVNmaWKkVB1r?si=2b29f5661ce94c44) by _John Lee Hooker_ and I could just lookup whatever album that's on (or pick one of the albums it's on) but that wouldn't be quite right. 

### Top {{ top_5_albums.length }}

If you only listen to {{ top_5_albums.length }} albums&hellip; well, that would be surprising. At the moment, these are my top {{ top_5_albums.length }}. They may change from time to time but Nina Simone will always be in there. _Little Girl Blue_ is secretly my #1.

<section class="album-grid top-five">
{% set gridSlugs = top_5_albums %}
{% include "album-grid.njk" %}
</section>

### Top {{ top_albums.length }}

In order of release date, rather than a ranking. 

<p class="list-search">
  <input type="search" data-list-search data-endpoint="/bill/albums/search.json" data-scope="#all-albums" placeholder="Search title, artist or track&hellip;" aria-label="Search the top {{ top_albums.length }} albums">
</p>

<section class="album-grid" id="all-albums">
{% set gridSlugs = top_albums %}
{% include "album-grid.njk" %}
</section>
