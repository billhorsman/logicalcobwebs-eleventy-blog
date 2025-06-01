---
title: "What's Eating Gilbert Grape"
layout: layouts/home.njk
slug: whats-eating-gilbert-grape-1993
ogImage: content/bill/films/backdrops/whats-eating-gilbert-grape-1993.jpg
description: "Gilbert Grape is a small-town young man with a lot of responsibility. Chief among his concerns are his mother, who is so overweight that she can't leave the house, and his mentally impaired younger brother, Arnie, who has a knack for finding trouble. Settled into a job at a grocery store and an ongoing affair with local woman Betty Carver, Gilbert finally has his life shaken up by the free-spirited Becky."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-fugitive-1993"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../four-weddings-and-a-funeral-1994">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>32 / 100</p>

<article class="film slug-whats-eating-gilbert-grape-1993">
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
