---
title: "Killers of the Flower Moon"
layout: layouts/films.njk
slug: killers-of-the-flower-moon-2023
ogImage: content/bill/films/backdrops/killers-of-the-flower-moon-2023.jpg
description: "When oil is discovered in 1920s Oklahoma under Osage Nation land, the Osage people are murdered one by one—until the FBI steps in to unravel the mystery."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../blue-jean-2023"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">99 / 100</a>
  </div>
  <div class="next">
    <a href="../the-ballad-of-wallis-island-2025">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Blue Jean
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Ballad of Wallis Island
    </span>
  </div>
</nav>

<article class="film slug-killers-of-the-flower-moon-2023">
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
    <li><a href="../the-deer-hunter-1978">The Deer Hunter</a> and <a href="../brazil-1985">Brazil</a> because of Robert De Niro</li>
<li><a href="../whats-eating-gilbert-grape-1993">What's Eating Gilbert Grape</a> because of Leonardo DiCaprio</li>
<li><a href="../magnolia-1999">Magnolia</a> because of Pat Healy</li>
<li><a href="../no-country-for-old-men-2007">No Country for Old Men</a> because of Barry Corbin and Gene Jones</li>
<li><a href="../dallas-buyers-club-2013">Dallas Buyers Club</a> because of Carl Palmer</li>
<li><a href="../interstellar-2014">Interstellar</a> because of John Lithgow</li>
<li><a href="../first-cow-2020">First Cow</a> because of Scott Shepherd and Lily Gladstone</li>
<li><a href="../the-power-of-the-dog-2021">The Power of the Dog</a> because of Jesse Plemons</li>
  </ul>
</section>

</article>
