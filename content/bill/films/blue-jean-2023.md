---
title: "Blue Jean"
layout: layouts/home.njk
slug: blue-jean-2023
ogImage: content/bill/films/backdrops/blue-jean-2023.jpg
description: "Jean, a PE teacher, is forced to live a double life. When a new student arrives and threatens to expose her sexuality, Jean is pushed to extreme lengths to keep her job and her integrity."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../empire-of-light-2022">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../asteroid-city-2023">Next</a>
</nav>

<p>97 / 100</p>

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
      {{ films.reviews[slug] }} <em>— Bill</em>
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
