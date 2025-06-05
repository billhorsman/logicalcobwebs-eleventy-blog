---
title: "Schemers"
layout: layouts/films.njk
slug: schemers-2020
ogImage: content/bill/films/backdrops/schemers-2020.jpg
description: "Set in late-1970s Dundee, Schemers is based on writer-producer David McLean’s early years in the music business. After a run in with a local gangster, a fledgling promoter and his two friends raise their ambitions to booking major bands in order to escape their debt."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../limbo-2020"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">78 / 100</a>
  </div>
  <div class="next">
    <a href="../the-truffle-hunters-2020">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Limbo
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Truffle Hunters
    </span>
  </div>
</nav>

<article class="film slug-schemers-2020">
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
