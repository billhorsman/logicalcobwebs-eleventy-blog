---
title: "Whisky Galore!"
layout: layouts/films.njk
slug: whisky-galore-1949
ogImage: content/bill/films/backdrops/whisky-galore-1949.jpg
description: "Based on a true story. The name of the real ship, that sunk Feb 5 1941 - during WWII - was S/S Politician. Having left Liverpool two days earlier, heading for Jamaica, it sank outside Eriskay, The Outer Hebrides, Scotland, in bad weather, containing 250,000 bottles of whisky. The locals gathered as many bottles as they could, before the proper authorities arrived, and even today, bottles are found in the sand or in the sea every other year."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../its-a-wonderful-life-1946"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">2 / 100</a>
  </div>
  <div class="next">
    <a href="../la-strada-1954">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      It's a Wonderful Life
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      La Strada
    </span>
  </div>
</nav>

<article class="film slug-whisky-galore-1949">
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
