---
title: Albums
permalink: /bill/albums/
eleventyComputed:
  description: "Bill Horsman's top {{ top_albums.length }} albums, from {{ albums[top_albums[0]].year }} to {{ albums[top_albums[top_albums.length - 1]].year }}."
---

## Favourite Albums

{% set firstAlbum = albums[top_albums[0]] %}
{% set lastAlbum = albums[top_albums[top_albums.length - 1]] %}
Starting from *{{ firstAlbum.title }}* in {{ firstAlbum.year }} through to *{{ lastAlbum.title }}* in {{ lastAlbum.year }}, here are my {{ top_albums.length }} favourite albums.

Almost impossible to pick, so this is me throwing in some to start with. I'm expecting to add more as I go (and to remove those that no longer make the cut).

### Top 5

If you only listen to {{ top_5_albums.length }} albums&hellip;

<section class="album-grid top-five">
{% set gridSlugs = top_5_albums %}
{% set showRank = true %}
{% include "album-grid.njk" %}
</section>

### Top {{ top_albums.length }}

<section class="album-grid">
{% set gridSlugs = top_albums %}
{% set showRank = false %}
{% include "album-grid.njk" %}
</section>
