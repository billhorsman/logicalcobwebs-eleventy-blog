---
title: "Ghost Dog: The Way of the Samurai"
layout: layouts/films.njk
slug: ghost-dog-the-way-of-the-samurai-1999
ogImage: content/bill/films/backdrops/ghost-dog-the-way-of-the-samurai-1999.jpg
description: "An African-American Mafia hit man who models himself after the samurai of ancient Japan finds himself targeted for death by the mob."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../fight-club-1999"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">35 / 100</a>
  </div>
  <div class="next">
    <a href="../magnolia-1999">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Fight Club
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Magnolia
    </span>
  </div>
</nav>

<article class="film slug-ghost-dog-the-way-of-the-samurai-1999">
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
    <li><a href="../night-on-earth-1991">Night on Earth</a> because of Isaach de Bankolé and Jim Jarmusch</li>
<li><a href="../phone-booth-2003">Phone Booth</a> because of Forest Whitaker</li>
  </ul>
</section>

</article>
