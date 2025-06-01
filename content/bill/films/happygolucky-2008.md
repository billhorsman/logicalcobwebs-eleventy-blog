---
title: "Happy-Go-Lucky"
layout: layouts/home.njk
slug: happygolucky-2008
ogImage: content/bill/films/backdrops/happygolucky-2008.jpg
description: "A look at a few chapters in the life of Poppy, a cheery, colorful, North London schoolteacher whose optimism tends to exasperate those around her."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../in-bruges-2008">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../district-9-2009">Next</a>
</nav>

<p>54 / 100</p>

<article class="film slug-happygolucky-2008">
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
