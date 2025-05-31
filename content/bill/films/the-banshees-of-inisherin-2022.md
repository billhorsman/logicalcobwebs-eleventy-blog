---
title: "The Banshees of Inisherin"
layout: layouts/home.njk
slug: the-banshees-of-inisherin-2022
ogImage: content/bill/films/backdrops/the-banshees-of-inisherin-2022.jpg
description: "Two lifelong friends find themselves at an impasse when one abruptly ends their relationship, with alarming consequences for both of them."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../one-fine-morning-2022">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-fabelmans-2022">Next</a>
</nav>

<p>94 / 100</p>

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
