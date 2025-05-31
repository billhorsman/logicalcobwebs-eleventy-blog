---
title: "The Party"
layout: layouts/home.njk
slug: the-party-2017
ogImage: content/bill/films/backdrops/the-party-2017.jpg
description: "Various individuals think they’re coming together for a party in a private home, but a series of revelations results in a huge crisis that throws their belief systems – and their values – into total disarray."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../maudie-2016">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../lucky-2017">Next</a>
</nav>

<p>68 / 100</p>

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
