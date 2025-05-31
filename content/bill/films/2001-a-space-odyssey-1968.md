---
title: "2001: A Space Odyssey"
layout: layouts/home.njk
slug: 2001-a-space-odyssey-1968
ogImage: content/bill/films/backdrops/2001-a-space-odyssey-1968.jpg
description: "Humanity finds a mysterious object buried beneath the lunar surface and sets off to find its origins with the help of HAL 9000, the world's most advanced super computer."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../in-the-heat-of-the-night-1967">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../bullitt-1968">Next</a>
</nav>

<p>11 / 100</p>

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
