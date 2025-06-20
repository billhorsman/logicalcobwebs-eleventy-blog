---
title: "District 9"
layout: layouts/films.njk
slug: district-9-2009
ogImage: content/bill/films/backdrops/district-9-2009.jpg
description: "Thirty years ago, aliens arrive on Earth. Not to conquer or give aid, but to find refuge from their dying planet. Separated from humans in a South African area called District 9, the aliens are managed by Multi-National United, which is unconcerned with the aliens' welfare but will do anything to master their advanced technology. When a company field agent contracts a mysterious virus that begins to alter his DNA, there is only one place he can hide: District 9."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../in-bruges-2008"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">54 / 100</a>
  </div>
  <div class="next">
    <a href="../fantastic-mr-fox-2009">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      In Bruges
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Fantastic Mr. Fox
    </span>
  </div>
</nav>

<article class="film slug-district-9-2009">
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
