---
title: Films
description: Films that Bill Horsman likes
excerpt: A work in progress while I attempt to remember all the films I like
canonical: https://logicalcobwebs.com/bill/films
ogImage: content/bill/bill.webp
layout: layouts/home.njk
---

## Favourite Films

{% set firstFilm = films.list[0] %}
{% set lastFilm = films.list[films.count - 1] %}
Starting from <a href="{{ firstFilm.slug }}">{{ firstFilm.title }}</a> in {{ firstFilm.year }} through to <a href="{{ lastFilm.slug }}">{{ lastFilm.title }}</a> in {{ lastFilm.year }}, here are my {{ films.list.length }} favourite films. There are lots that are brilliant and don't make the cut but it's a nice round number. 

Based on these films, my favourite director is {{ films.top_directors[0].name }} ({{ films.top_directors[0].count }} films), followed by {{ films.top_directors[1].name }} ({{ films.top_directors[1].count }}) and {{ films.top_directors[2].name }} ({{ films.top_directors[2].count }}).

A few films are in here because I watched them a long time ago and they have a nostalgic appeal. Films like <span class="sentence-list">{% for film in films.nostaligic %}<span><a href="{{ film.slug }}">{{ film.title }}</a></span>{% endfor %}</span>.

One film is in here twice: <a href="purple-noon">Purple Noon</a> (aka Plan Soleil) and <a href="the-talented-mr-ripley">The Talented Mr. Ripley</a>. I would have added it again for Steven Zaillian's 2024 TV miniseries <em>Ripley</em> with Andrew Scott but that's not a film.

Actors that crop up a lot are <span class="sentence-list">{% for cast in films.top_cast %}<span>{{ cast.name }}</span>{% endfor %}</span>.

### Top {{ films.must_see.length }}
If you only watch {{ films.must_see.length }} films&hellip;

<ul class="film-list">
{% for film in films.must_see %}
  <li>
    {% include "film.njk" %}
  </li>
{% endfor %}
</ul>

<section class="list">
  <header>
    <h3>All {{ films.list.length }}</h3>
    <button type="button" data-toggle-list="posters" aria-label="Toggle view">
      <i class="fa-solid fa-list"></i>
      <i class="fa-solid fa-grip posters"></i> 
    </button>
  </header>
  <ul class="film-list">
  {% for film in films.list %}
  <li>
    {% include "film.njk" %}
  </li>
  {% endfor %}
  </ul>
</section>

<footer>
  <a href="about">About this list</a>
</footer>
