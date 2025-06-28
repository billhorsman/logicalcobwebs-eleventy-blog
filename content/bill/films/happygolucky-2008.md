---
title: "Happy-Go-Lucky"
layout: layouts/films.njk
slug: happygolucky-2008
ogImage: content/bill/films/backdrops/happygolucky-2008.jpg
description: "A look at a few chapters in the life of Poppy, a cheery, colorful, North London schoolteacher whose optimism tends to exasperate those around her."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../no-country-for-old-men-2007"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">51 / 100</a>
  </div>
  <div class="next">
    <a href="../in-bruges-2008">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      No Country for Old Men
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      In Bruges
    </span>
  </div>
</nav>

<article class="film slug-happygolucky-2008">
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
  <li><a href="../maudie-2016">Maudie</a> and <a href="../eternal-beauty-2020">Eternal Beauty</a> by Sally Hawkins</li>
<li><a href="../mr-turner-2014">Mr. Turner</a> by Sylvestra Le Touzel, Kate O'Flynn, Oliver Maltman, Karina Fernandez, Sinead Matthews and Mike Leigh</li>
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
