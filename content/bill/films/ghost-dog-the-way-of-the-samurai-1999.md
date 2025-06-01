---
title: "Ghost Dog: The Way of the Samurai"
layout: layouts/home.njk
slug: ghost-dog-the-way-of-the-samurai-1999
ogImage: content/bill/films/backdrops/ghost-dog-the-way-of-the-samurai-1999.jpg
description: "An African-American Mafia hit man who models himself after the samurai of ancient Japan finds himself targeted for death by the mob."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../fight-club-1999">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../magnolia-1999">Next</a>
</nav>

<p>39 / 100</p>

<article class="film slug-ghost-dog-the-way-of-the-samurai-1999">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {% if films.reviews[slug] %}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>— Bill</em>
    </blockquote> 
  {% endif %}

  <h2>
    Cast
  </h2>
  <ul>
    {%- for cast in film.credits.cast -%}
      <li>
        {{ cast.name }} as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>

  <h2>
    Crew
  </h2>
  <ul>
    {%- for crew in film.credits.crew -%}
      <li>
        {{ crew.name }} &mdash; <em>{{ crew.job }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
