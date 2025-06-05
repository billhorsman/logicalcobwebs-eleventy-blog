---
title: "The Truffle Hunters"
layout: layouts/films.njk
slug: the-truffle-hunters-2020
ogImage: content/bill/films/backdrops/the-truffle-hunters-2020.jpg
description: "In the secret forests of Northern Italy, a dwindling group of joyful old men and their faithful dogs search for the world’s most expensive ingredient, the white Alba truffle. Their stories form a real-life fairy tale that celebrates human passion in a fragile land that seems forgotten in time."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../schemers-2020"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">79 / 100</a>
  </div>
  <div class="next">
    <a href="../belfast-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Schemers
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Belfast
    </span>
  </div>
</nav>

<article class="film slug-the-truffle-hunters-2020">
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
