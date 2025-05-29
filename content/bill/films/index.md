---
title: Films
description: Films that Bill Horsman likes
excerpt: A work in progress while I attempt to remember all the films I like
canonical: https://logicalcobwebs.com/bill/films
ogImage: content/bill/bill.webp
layout: layouts/home.njk
---

## Top {{ films.count }} Films

{% set firstFilm = films.list[0] %}
{% set lastFilm = films.list[films.count - 1] %}
Starting from <em>{{ firstFilm.title }}</em> in {{ firstFilm.year }} through to <em>{{ lastFilm.title }}</em> in {{ lastFilm.year }}.

<div class="film-list">
{% for film in films.list %}
<div><a href="{{ film.slug }}"><img src="{{ film.poster }}" alt="{{ film.title }}"></a></div>
{% endfor %}
</div>
