---
title: Films
eleventyComputed:
  description: "Bill Horsman's top 100 films, from {{ films.list[0].year }} to {{ films.list[films.count - 1].year }} including Diva, Ghost Dog: The Way of the Samurai, Night on Earth, Woman at War and Portrait of a Lady on Fire."
ogImage: content/bill/films/montage.jpg
layout: layouts/home.njk
---

## Favourite Films

{% set firstFilm = films[top_films.list[0]] %}
{% set lastFilm = films[top_films.list[top_films.list.length - 1]] %}
Starting from {% filmLink firstFilm.slug %} in {{ firstFilm.year }} through to {% filmLink lastFilm.slug %} in {{ lastFilm.year }}, here are my {{ top_films.list.length }} favourite films. There are lots that are brilliant and don't make the cut but it's a nice round number. 

Based on these films, my favourite director is {{ films.top_directors[0].name }} ({{ films.top_directors[0].count }} films), followed by {{ films.top_directors[1].name }} ({{ films.top_directors[1].count }}) and {{ films.top_directors[2].name }} ({{ films.top_directors[2].count }}).

A few films are in here because I watched them a long time ago and they have a nostalgic appeal. Films like {% summariseFilms nostalgic_films %}.

One film is in here twice: {% filmLink "purple-noon-1960" %} (aka Plan Soleil) and {% filmLink "the-talented-mr-ripley-1999" %}. I would have added it again for Steven Zaillian's 2024 TV miniseries <em>Ripley</em> with Andrew Scott but that's not a film.

Actors that crop up a lot are <span class="sentence-list">{% for cast in films.top_cast %}<span>{{ cast.name }}</span>{% endfor %}</span>.

### Top {{ must_see_films.list.length }}

If you only watch {{ must_see_films.list.length }} films&hellip;

<ul class="film-list">
{% for filmSlug in must_see_films.list %}
  {% set review = films.reviews[filmSlug] %}
  <li>
    {% include "film.njk" %}
  </li>
{% endfor %}
</ul>

<section class="list">
  <header>
    <h3>All {{ top_films.list.length }}</h3>
    <button type="button" data-toggle-list="posters" aria-label="Toggle view">
      <i class="fa-solid fa-list"></i>
      <i class="fa-solid fa-grip posters"></i> 
    </button>
  </header>
  <ul class="film-list">
  {% for filmSlug in top_films.list %}
  <li>
    {% include "film.njk" %}
  </li>
  {% endfor %}
  </ul>
</section>

<footer>
  <a href="about">About this list</a>
</footer>
