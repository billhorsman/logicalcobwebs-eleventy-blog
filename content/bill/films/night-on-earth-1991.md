---
title: "Night on Earth"
layout: layouts/home.njk
slug: night-on-earth-1991
ogImage: content/bill/films/backdrops/night-on-earth-1991.jpg
description: "An anthology of 5 different cab drivers in 5 American and European cities and their remarkable fares on the same eventful night."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../delicatessen-1991">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-fugitive-1993">Next</a>
</nav>

<p>31 / 100</p>

<article class="film slug-night-on-earth-1991">
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
