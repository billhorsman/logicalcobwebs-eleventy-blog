---
title: "Lady Bird"
layout: layouts/films.njk
slug: lady-bird-2017
ogImage: content/bill/films/backdrops/lady-bird-2017.jpg
description: "Lady Bird McPherson, a strong willed, deeply opinionated, artistic 17 year old comes of age in Sacramento. Her relationship with her mother and her upbringing are questioned and tested as she plans to head off to college."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../cest-la-vie-2017"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">62 / 100</a>
  </div>
  <div class="next">
    <a href="../lucky-2017">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      C'est la vie!
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Lucky
    </span>
  </div>
</nav>

<article class="film slug-lady-bird-2017">
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

  <section class="related-films">
  <h2>Related films</h2>
  <ul>
    <li><a href="../fight-club-1999">Fight Club</a> because of Bob Stephenson</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Saoirse Ronan and Lucas Hedges</li>
<li><a href="../little-women-2019">Little Women</a> because of Saoirse Ronan, Tracy Letts, Timothée Chalamet and Greta Gerwig</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Saoirse Ronan, Timothée Chalamet and Lois Smith</li>
<li><a href="../dune-2021">Dune</a> because of Timothée Chalamet and Stephen McKinley Henderson</li>
  </ul>
</section>

</article>
