---
title: "The Banshees of Inisherin"
layout: layouts/films.njk
slug: the-banshees-of-inisherin-2022
ogImage: content/bill/films/backdrops/the-banshees-of-inisherin-2022.jpg
description: "Two lifelong friends find themselves at an impasse when one abruptly ends their relationship, with alarming consequences for both of them."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../one-fine-morning-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">95 / 100</a>
  </div>
  <div class="next">
    <a href="../the-fabelmans-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<article class="film slug-the-banshees-of-inisherin-2022">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    
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
