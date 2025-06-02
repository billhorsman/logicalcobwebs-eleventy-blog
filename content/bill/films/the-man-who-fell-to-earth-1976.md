---
title: "The Man Who Fell to Earth"
layout: layouts/home.njk
slug: the-man-who-fell-to-earth-1976
ogImage: content/bill/films/backdrops/the-man-who-fell-to-earth-1976.jpg
description: "Thomas Jerome Newton is an alien who has come to Earth in search of water to save his home planet. Aided by lawyer Oliver Farnsworth, Thomas uses his knowledge of advanced technology to create profitable inventions. While developing a method to transport water, Thomas meets Mary-Lou, a quiet hotel clerk, and begins to fall in love with her. Just as he is ready to leave Earth, Thomas is intercepted by the U.S. government, and his entire plan is threatened."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../three-days-of-the-condor-1975"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../the-deer-hunter-1978">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>18 / 100</p>

<article class="film slug-the-man-who-fell-to-earth-1976">
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
