---
title: "Eternal Beauty"
layout: layouts/films.njk
slug: eternal-beauty-2020
ogImage: content/bill/films/backdrops/eternal-beauty-2020.jpg
description: "When Jane is rejected by life, she spirals into a chaotic, schizophrenic world where love and normality collide with humorous consequences."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../1917-2019"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
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
      1917
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Limbo
    </span>
  </div>
</nav>

<article class="film slug-eternal-beauty-2020">
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
    <li><a href="../the-big-lebowski-1998">The Big Lebowski</a> because of David Thewlis</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> because of Alice Lowe</li>
<li><a href="../happygolucky-2008">Happy-Go-Lucky</a> and <a href="../maudie-2016">Maudie</a> because of Sally Hawkins</li>
  </ul>
</section>

</article>
