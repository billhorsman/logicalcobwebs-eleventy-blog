---
title: "Empire of Light"
layout: layouts/films.njk
slug: empire-of-light-2022
ogImage: content/bill/films/backdrops/empire-of-light-2022.jpg
description: "The duty manager of a seaside cinema, who is struggling with her mental health, forms a relationship with a new employee on the south coast of England in the 1980s."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../eo-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">92 / 100</a>
  </div>
  <div class="next">
    <a href="../one-fine-morning-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      EO
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      One Fine Morning
    </span>
  </div>
</nav>

<article class="film slug-empire-of-light-2022">
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
  <li><a href="../24-hour-party-people-2002">24 Hour Party People</a> by Ron Cook</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> by Ron Cook and Olivia Colman</li>
<li><a href="../first-cow-2020">First Cow</a> by Toby Jones</li>
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
