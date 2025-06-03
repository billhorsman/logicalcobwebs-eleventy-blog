---
title: "Amélie"
layout: layouts/home.njk
slug: amlie-2001
ogImage: content/bill/films/backdrops/amlie-2001.jpg
description: "At a tiny Parisian café, the adorable yet painfully shy Amélie accidentally discovers a gift for helping others. Soon Amelie is spending her days as a matchmaker, guardian angel, and all-around do-gooder. But when she bumps into a handsome stranger, will she find the courage to become the star of her very own love story?"
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../billy-elliot-2000"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../black-hawk-down-2001">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>44 / 100</p>

<article class="film slug-amlie-2001">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Le Fabuleux Destin d'Amélie Poulain.
  </p>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {%- if films.reviews[slug] -%}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href="/bill">Bill</a></em>
    </blockquote> 
  {%- endif -%}

  <details>
    <summary>
      Cast
    </summary>
  <ul>
    {%- for cast in film.credits.cast -%}
      <li>
        {{ cast.name }} as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>
  </details>

  <details>
    <summary>
      Crew
    </summary>
    <ul>
      {%- for crew in film.credits.crew -%}
        <li>
          {{ crew.name }} &mdash; <em>{{ crew.job }}</em>
        </li>
      {%- endfor -%}
    </ul>
  </details>
  
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
