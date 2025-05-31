---
title: "Maudie"
layout: layouts/home.njk
slug: maudie-2016
ogImage: content/bill/films/backdrops/maudie-2016.jpg
description: "Canadian folk artist Maud Lewis falls in love with a fishmonger while working for him as a live-in housekeeper."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-handmaiden-2016">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-party-2017">Next</a>
</nav>

<p>67 / 100</p>

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
