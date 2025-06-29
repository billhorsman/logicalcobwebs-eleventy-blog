---
title: "Little Women"
layout: layouts/films.njk
slug: little-women-2019
ogImage: content/bill/films/backdrops/little-women-2019.jpg
description: "Four sisters come of age in America in the aftermath of the Civil War."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../woman-at-war-2018"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">69 / 100</a>
  </div>
  <div class="next">
    <a href="../parasite-2019">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Woman at War
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Parasite
    </span>
  </div>
</nav>

<article class="film slug-little-women-2019">
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
    <li><a href="../the-deer-hunter-1978">The Deer Hunter</a> and <a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> because of Meryl Streep</li>
<li><a href="../the-bourne-identity-2002">The Bourne Identity</a> because of Chris Cooper</li>
<li><a href="../interstellar-2014">Interstellar</a> and <a href="../dune-2021">Dune</a> because of Timothée Chalamet</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Timothée Chalamet and Saoirse Ronan</li>
<li><a href="../mr-turner-2014">Mr. Turner</a> because of James Norton</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Saoirse Ronan</li>
<li><a href="../coda-2021">CODA</a> because of Lonnie Farmer</li>
  </ul>
</section>

</article>
