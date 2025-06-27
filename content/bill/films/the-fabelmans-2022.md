---
title: "The Fabelmans"
layout: layouts/films.njk
slug: the-fabelmans-2022
ogImage: content/bill/films/backdrops/the-fabelmans-2022.jpg
description: "Growing up in post-World War II era Arizona, young Sammy Fabelman aspires to become a filmmaker as he reaches adolescence, but soon discovers a shattering family secret and explores how the power of films can help him see the truth."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-banshees-of-inisherin-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">95 / 100</a>
  </div>
  <div class="next">
    <a href="../all-of-us-strangers-2023">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Banshees of Inisherin
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      All of Us Strangers
    </span>
  </div>
</nav>

<article class="film slug-the-fabelmans-2022">
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
  <li><a href="../fight-club-1999">Fight Club</a> (1999) by Ezra Buzzington</li>
<li><a href="../magnolia-1999">Magnolia</a> (1999) by Ezra Buzzington</li>
<li><a href="../lucky-2017">Lucky</a> (2017) by David Lynch</li>
<li><a href="../licorice-pizza-2021">Licorice Pizza</a> (2021) by Isabelle Kusman and Paige Locke</li>
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
