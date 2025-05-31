---
title: "Eternal Beauty"
layout: layouts/home.njk
slug: eternal-beauty-2020
ogImage: content/bill/films/backdrops/eternal-beauty-2020.jpg
description: "When Jane is rejected by life, she spirals into a chaotic, schizophrenic world where love and normality collide with humorous consequences."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../schemers-2020">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../limbo-2020">Next</a>
</nav>

<p>77 / 100</p>

<article class="film">
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
        <strong>{{ cast.name }}</strong> as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
