---
title: "Tomboy"
layout: layouts/home.njk
slug: tomboy-2011
ogImage: content/bill/films/backdrops/tomboy-2011.jpg
description: "A French family moves to a new neighborhood with during the summer holidays. The story follows a 10-year-old gender non-conforming child, Laure, who experiments with their gender presentation, adopting the name Mikäel."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../le-havre-2011">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../all-is-lost-2013">Next</a>
</nav>

<p>59 / 100</p>

<article class="film slug-tomboy-2011">
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
