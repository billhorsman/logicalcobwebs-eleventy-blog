---
title: "First Cow"
layout: layouts/films.njk
slug: first-cow-2020
ogImage: content/bill/films/backdrops/first-cow-2020.jpg
description: "In the 1820s, a taciturn loner and skilled cook travels west to Oregon Territory, where he meets a Chinese immigrant also seeking his fortune. Soon the two team up on a dangerous scheme to steal milk from the wealthy landowner’s prized Jersey cow – the first, and only, in the territory."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../eternal-beauty-2020"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">75 / 100</a>
  </div>
  <div class="next">
    <a href="../limbo-2020">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Eternal Beauty
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Limbo
    </span>
  </div>
</nav>

<article class="film slug-first-cow-2020">
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
