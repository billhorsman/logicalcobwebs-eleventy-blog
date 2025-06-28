---
title: "Belfast"
layout: layouts/films.njk
slug: belfast-2021
ogImage: content/bill/films/backdrops/belfast-2021.jpg
description: "Buddy is a young boy on the cusp of adolescence, whose life is filled with familial love, childhood hijinks, and a blossoming romance. Yet, with his beloved hometown caught up in increasing turmoil, his family faces a momentous choice: hope the conflict will pass or leave everything they know behind for a new life."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-truffle-hunters-2020"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">79 / 100</a>
  </div>
  <div class="next">
    <a href="../coda-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Truffle Hunters
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      CODA
    </span>
  </div>
</nav>

<article class="film slug-belfast-2021">
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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../in-bruges-2008">In Bruges</a> by Ciarán Hinds</li>
  </ul>

  <section class="film-detail">
    <div>
      <details>
        <summary>
          <i class="fa-solid fa-masks-theater"></i>
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
          <i class="fa-solid fa-clapperboard"></i>
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
    </div>
  </section>
</article>
