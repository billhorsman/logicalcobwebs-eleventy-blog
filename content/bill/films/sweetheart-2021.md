---
title: "Sweetheart"
layout: layouts/home.njk
slug: sweetheart-2021
ogImage: content/bill/films/backdrops/sweetheart-2021.jpg
description: "A socially awkward, environmentally-conscious teenager named AJ is dragged to a coastal holiday park by her painfully 'normal' family, where she becomes unexpectedly captivated by a chlorine smelling, sun-loving lifeguard named Isla."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../dune-2021">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-french-dispatch-2021">Next</a>
</nav>

<p>84 / 100</p>

<article class="film slug-sweetheart-2021">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {% if films.reviews[slug] %}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>— Bill</em>
    </blockquote> 
  {% endif %}

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
