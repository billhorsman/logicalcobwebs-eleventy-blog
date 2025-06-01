---
title: "Four Weddings and a Funeral"
layout: layouts/home.njk
slug: four-weddings-and-a-funeral-1994
ogImage: content/bill/films/backdrops/four-weddings-and-a-funeral-1994.jpg
description: "Over the course of five social occasions, a committed bachelor must consider the notion that he may have discovered love."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../whats-eating-gilbert-grape-1993"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../lon-the-professional-1994">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>33 / 100</p>

<article class="film slug-four-weddings-and-a-funeral-1994">
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
