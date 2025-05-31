---
title: "The Motorcycle Diaries"
layout: layouts/home.njk
slug: the-motorcycle-diaries-2004
ogImage: content/bill/films/backdrops/the-motorcycle-diaries-2004.jpg
description: "Based on the journals of Che Guevara, leader of the Cuban Revolution. In his memoirs, Guevara recounts adventures he and best friend Alberto Granado had while crossing South America by motorcycle in the early 1950s."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../phone-booth-2003">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../hot-fuzz-2007">Next</a>
</nav>

<p>51 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>Diarios de motocicleta</strong></p>

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
