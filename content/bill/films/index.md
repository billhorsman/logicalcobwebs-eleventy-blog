---
title: Films
description: Films that Bill Horsman likes
excerpt: A work in progress while I attempt to remember all the films I like
canonical: https://logicalcobwebs.com/bill/films
ogImage: content/bill/bill.webp
layout: layouts/home.njk
---

## Top 100 Films

<div class="film-list">
{% for film in films.list %}
  <img src="{{ film.poster }}" alt="{{ film.title }}">
{% endfor %}
</div>
