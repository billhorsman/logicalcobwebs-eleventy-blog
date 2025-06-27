---
title: Films
eleventyComputed:
  description: "Bill Horsman's top 100 films, from {{ films[top_films[0]].year }} to {{ films[top_films[top_films.length - 1]].year }} including Diva, Ghost Dog: The Way of the Samurai, Night on Earth, Woman at War and Portrait of a Lady on Fire."
ogImage: content/bill/films/montage.jpg
layout: layouts/films.njk
---

## Favourite Films

{% set firstFilm = films[top_films[0]] %}
{% set lastFilm = films[top_films[top_films.length - 1]] %}
Starting from {% filmLink firstFilm.slug %} in {{ firstFilm.year }} through to {% filmLink lastFilm.slug %} in {{ lastFilm.year }}, here are my {{ top_films.length }} favourite films. There are lots that are brilliant and don't make the cut but it's a nice round number. 

Based on these films, my favourite directors are {% namesAndCounts top_directors, 5 %}.

Actors that crop up a lot are {% namesAndCounts top_cast, 6 %}.

A few films are in here because I watched them a long time ago and they have a nostalgic appeal. Films like {% summariseFilms nostalgic_films %}.

Some older films I only discovered relatively recently: {% summariseFilms ["la-strada-1954", "la-dolce-vita-1960", "dog-day-afternoon-1975"] %}.

One film is in here twice: {% filmLink "purple-noon-1960" %} (aka Plan Soleil) and {% filmLink "the-talented-mr-ripley-1999" %}. I would have added it again for Steven Zaillian's 2024 TV miniseries <em>Ripley</em> with Andrew Scott but that's not a film.

Only one person is connected with more than one film in my top 5: Isaach de Bankolé, who plays a Parisian cab driver in {% filmLink "night-on-earth-1991" %} and Raymond in {% filmLink "ghost-dog-the-way-of-the-samurai-1999" %}.

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
    <h3>All {{ top_films.length }}</h3>
    <button type="button" data-toggle-list="posters" aria-label="Toggle view">
      <span class="caption">Toggle display</span>
      <i class="fa-solid fa-list fa-fw"></i>
      <i class="fa-solid fa-grip posters fa-fw"></i> 
    </button>
  </header>
  <p>In order of release&hellip;</p>
  <ul class="film-list">
  {% for filmSlug in top_films %}
  <li>
    {% include "film.njk" %}
  </li>
  {% endfor %}
  </ul>
</section>
