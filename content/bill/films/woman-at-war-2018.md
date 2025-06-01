---
title: "Woman at War"
layout: layouts/home.njk
slug: woman-at-war-2018
ogImage: content/bill/films/backdrops/woman-at-war-2018.jpg
description: "Halla declares a one-woman-war on the local aluminium industry. She is prepared to risk everything to protect the pristine Icelandic Highlands she loves… Until an orphan unexpectedly enters her life."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../sink-or-swim-2018">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../parasite-2019">Next</a>
</nav>

<p>72 / 100</p>

<article class="film slug-woman-at-war-2018">
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
