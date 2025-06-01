---
title: "The Deer Hunter"
layout: layouts/home.njk
slug: the-deer-hunter-1978
ogImage: content/bill/films/backdrops/the-deer-hunter-1978.jpg
description: "A group of working-class friends decide to enlist in the Army during the Vietnam War and finds it to be hellish chaos -- not the noble venture they imagined. Before they left, Steven married his pregnant girlfriend -- and Michael and Nick were in love with the same woman. But all three are different men upon their return."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-man-who-fell-to-earth-1976"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../apocalypse-now-1979">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>19 / 100</p>

<article class="film slug-the-deer-hunter-1978">
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
