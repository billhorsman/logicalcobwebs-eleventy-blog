---
title: "Being There"
layout: layouts/films.njk
slug: being-there-1979
ogImage: content/bill/films/backdrops/being-there-1979.jpg
description: "A simple-minded gardener named Chance has spent all his life in the Washington D.C. house of an old man. When the man dies, Chance is put out on the street with no knowledge of the world except what he has learned from television."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../apocalypse-now-1979"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">21 / 100</a>
  </div>
  <div class="next">
    <a href="../diva-1981">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Apocalypse Now
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Diva
    </span>
  </div>
</nav>

<article class="film slug-being-there-1979">
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
