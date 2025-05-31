---
title: "Being There"
layout: layouts/home.njk
slug: being-there-1979
ogImage: content/bill/films/backdrops/being-there-1979.jpg
description: "A simple-minded gardener named Chance has spent all his life in the Washington D.C. house of an old man. When the man dies, Chance is put out on the street with no knowledge of the world except what he has learned from television."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../apocalypse-now-1979">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../diva-1981">Next</a>
</nav>

<p>22 / 100</p>

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
