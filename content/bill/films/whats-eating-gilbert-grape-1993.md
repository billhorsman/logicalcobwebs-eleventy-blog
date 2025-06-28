---
title: "What's Eating Gilbert Grape"
layout: layouts/films.njk
slug: whats-eating-gilbert-grape-1993
ogImage: content/bill/films/backdrops/whats-eating-gilbert-grape-1993.jpg
description: "Gilbert Grape is a small-town young man with a lot of responsibility. Chief among his concerns are his mother, who is so overweight that she can't leave the house, and his mentally impaired younger brother, Arnie, who has a knack for finding trouble. Settled into a job at a grocery store and an ongoing affair with local woman Betty Carver, Gilbert finally has his life shaken up by the free-spirited Becky."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../night-on-earth-1991"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">29 / 100</a>
  </div>
  <div class="next">
    <a href="../lon-the-professional-1994">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Night on Earth
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Léon: The Professional
    </span>
  </div>
</nav>

<article class="film slug-whats-eating-gilbert-grape-1993">
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
  <li><a href="../killers-of-the-flower-moon-2023">Killers of the Flower Moon</a> by Leonardo DiCaprio</li>
<li><a href="../magnolia-1999">Magnolia</a> and <a href="../licorice-pizza-2021">Licorice Pizza</a> by John C. Reilly</li>
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
