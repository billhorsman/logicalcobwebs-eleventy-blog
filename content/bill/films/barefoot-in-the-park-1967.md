---
title: "Barefoot in the Park"
layout: layouts/films.njk
slug: barefoot-in-the-park-1967
ogImage: content/bill/films/backdrops/barefoot-in-the-park-1967.jpg
description: "In this film based on a Neil Simon play, newlyweds Corie, a free spirit, and Paul Bratter, an uptight lawyer, share a sixth-floor apartment in Greenwich Village. Soon after their marriage, Corie tries to find a companion for mother, Ethel, who is now alone, and sets up Ethel with neighbor Victor. Inappropriate behavior on a double date causes conflict, and the young couple considers divorce."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../purple-noon-1960"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">9 / 100</a>
  </div>
  <div class="next">
    <a href="../in-the-heat-of-the-night-1967">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Purple Noon
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      In the Heat of the Night
    </span>
  </div>
</nav>

<article class="film slug-barefoot-in-the-park-1967">
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
