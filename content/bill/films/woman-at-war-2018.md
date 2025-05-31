---
title: "Woman at War"
layout: layouts/home.njk
slug: woman-at-war-2018
ogImage: content/bill/films/backdrops/woman-at-war-2018.jpg
description: "Halla declares a one-woman-war on the local aluminium industry. She is prepared to risk everything to protect the pristine Icelandic Highlands she loves… Until an orphan unexpectedly enters her life."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../cest-la-vie-2017">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../sink-or-swim-2018">Next</a>
</nav>

<p>71 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>Kona fer í stríð</strong></p>

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
