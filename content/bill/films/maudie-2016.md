---
title: "Maudie"
layout: layouts/films.njk
slug: maudie-2016
ogImage: content/bill/films/backdrops/maudie-2016.jpg
description: "Canadian folk artist Maud Lewis falls in love with a fishmonger while working for him as a live-in housekeeper."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-grand-budapest-hotel-2014"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">63 / 100</a>
  </div>
  <div class="next">
    <a href="../the-handmaiden-2016">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Grand Budapest Hotel
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Handmaiden
    </span>
  </div>
</nav>

<article class="film slug-maudie-2016">
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
  <li><a href="../happygolucky-2008">Happy-Go-Lucky</a> and <a href="../eternal-beauty-2020">Eternal Beauty</a> by Sally Hawkins</li>
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
