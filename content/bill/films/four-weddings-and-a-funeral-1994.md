---
title: "Four Weddings and a Funeral"
layout: layouts/films.njk
slug: four-weddings-and-a-funeral-1994
ogImage: content/bill/films/backdrops/four-weddings-and-a-funeral-1994.jpg
description: "Over the course of five social occasions, a committed bachelor must consider the notion that he may have discovered love."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../whats-eating-gilbert-grape-1993"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">33 / 100</a>
  </div>
  <div class="next">
    <a href="../lon-the-professional-1994">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      What's Eating Gilbert Grape
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Léon: The Professional
    </span>
  </div>
</nav>

<article class="film slug-four-weddings-and-a-funeral-1994">
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
