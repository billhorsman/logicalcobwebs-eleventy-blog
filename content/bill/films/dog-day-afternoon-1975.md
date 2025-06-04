---
title: "Dog Day Afternoon"
layout: layouts/films.njk
slug: dog-day-afternoon-1975
ogImage: content/bill/films/backdrops/dog-day-afternoon-1975.jpg
description: "Based on the true story of would-be Brooklyn bank robbers John Wojtowicz and Salvatore Naturile. Sonny and Sal attempt a bank heist which quickly turns sour and escalates into a hostage situation and stand-off with the police. As Sonny's motives for the robbery are slowly revealed and things become more complicated, the heist turns into a media circus."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-sting-1973"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">16 / 100</a>
  </div>
  <div class="next">
    <a href="../three-days-of-the-condor-1975">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<article class="film slug-dog-day-afternoon-1975">
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

  <details>
    <summary>
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

</article>
<footer>
  <a href="../about">About this list</a>
</footer>
