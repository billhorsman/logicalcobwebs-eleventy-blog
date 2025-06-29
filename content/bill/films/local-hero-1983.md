---
title: "Local Hero"
layout: layouts/films.njk
slug: local-hero-1983
ogImage: content/bill/films/backdrops/local-hero-1983.jpg
description: "An American oil company sends a man to Scotland to buy up an entire village where they want to build a refinery. But things don't go as expected."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../blade-runner-1982"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">22 / 100</a>
  </div>
  <div class="next">
    <a href="../paris-texas-1984">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Blade Runner
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Paris, Texas
    </span>
  </div>
</nav>

<article class="film slug-local-hero-1983">
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
