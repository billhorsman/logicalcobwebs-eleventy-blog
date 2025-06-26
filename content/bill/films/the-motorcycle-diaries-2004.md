---
title: "The Motorcycle Diaries"
layout: layouts/films.njk
slug: the-motorcycle-diaries-2004
ogImage: content/bill/films/backdrops/the-motorcycle-diaries-2004.jpg
description: "Based on the journals of Che Guevara, leader of the Cuban Revolution. In his memoirs, Guevara recounts adventures he and best friend Alberto Granado had while crossing South America by motorcycle in the early 1950s."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../phone-booth-2003"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">48 / 100</a>
  </div>
  <div class="next">
    <a href="../hot-fuzz-2007">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Phone Booth
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Hot Fuzz
    </span>
  </div>
</nav>

<article class="film slug-the-motorcycle-diaries-2004">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Diarios de motocicleta.
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
