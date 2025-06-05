---
title: "Blue Jean"
layout: layouts/films.njk
slug: blue-jean-2023
ogImage: content/bill/films/backdrops/blue-jean-2023.jpg
description: "Jean, a PE teacher, is forced to live a double life. When a new student arrives and threatens to expose her sexuality, Jean is pushed to extreme lengths to keep her job and her integrity."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../asteroid-city-2023"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">99 / 100</a>
  </div>
  <div class="next">
    <a href="../killers-of-the-flower-moon-2023">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Asteroid City
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Killers of the Flower Moon
    </span>
  </div>
</nav>

<article class="film slug-blue-jean-2023">
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
