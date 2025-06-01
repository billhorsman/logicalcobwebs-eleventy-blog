---
title: "Whisky Galore!"
layout: layouts/home.njk
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
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../la-strada-1954">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>2 / 100</p>

<article class="film slug-whisky-galore-1949">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  <blockquote> 
    {{ film.overview }} <em>—&nbsp;<a href="https://www.themoviedb.org/movie/{{ film.id }}">TMDB</a></em>
  </blockquote> 

  {%- if films.reviews[slug] -%}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href="/bill">Bill</a></em>
    </blockquote> 
  {%- endif -%}

  <h2>
    Cast
  </h2>
  <ul>
    {%- for cast in film.credits.cast -%}
      <li>
        {{ cast.name }} as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>

  <h2>
    Crew
  </h2>
  <ul>
    {%- for crew in film.credits.crew -%}
      <li>
        {{ crew.name }} &mdash; <em>{{ crew.job }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
