---
title: Albums
permalink: /bill/albums/
ogImage: content/bill/albums/montage.jpg
eleventyComputed:
  description: "Bill Horsman's top {{ top_albums.length }} albums, from {{ albums[top_albums[0]].year }} to {{ albums[top_albums[top_albums.length - 1]].year }}."
---

<h1 class="visually-hidden">Albums</h1>

## Favourite Albums

{% set firstAlbum = albums[top_albums[0]] %}
{% set lastAlbum = albums[top_albums[top_albums.length - 1]] %}
Starting from *{{ firstAlbum.title }}* in {{ firstAlbum.year }} through to *{{ lastAlbum.title }}* in {{ lastAlbum.year }}, here are my {{ top_albums.length }} favourite albums.

Almost impossible to pick, so this is me throwing in some to start with. I'm expecting to add more as I go (and to remove those that no longer make the cut).

### Top {{ top_5_albums.length }}

If you only listen to {{ top_5_albums.length }} albums&hellip; well, that would be surprising. At the moment, these are my top {{ top_5_albums.length }}. They may change from time to time but Nina Simone will always be in there. _Little Girl Blue_ is secretly my #1.

<section class="album-grid top-five">
{% set gridSlugs = top_5_albums %}
{% include "album-grid.njk" %}
</section>

### Top {{ top_albums.length }}

In order of release date, rather than a ranking. 

<section class="album-grid">
{% set gridSlugs = top_albums %}
{% include "album-grid.njk" %}
</section>
