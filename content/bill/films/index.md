---
title: Films
description: Films that Bill Horsman likes
excerpt: A work in progress while I attempt to remember all the films I like
canonical: https://logicalcobwebs.com/bill/films
ogImage: content/bill/bill.webp
layout: layouts/home.njk
---

## Top 100 Films

In order of release date:

<div class="film-list">
{% for film in films.list %}
<div><a href="{{ film.slug }}"><img src="{{ film.poster }}" alt="{{ film.title }}"></a></div>
{% endfor %}
</div>
