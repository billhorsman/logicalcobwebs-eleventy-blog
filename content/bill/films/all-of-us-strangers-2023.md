---
title: "All of Us Strangers"
layout: layouts/home.njk
slug: all-of-us-strangers-2023
ogImage: content/bill/films/backdrops/all-of-us-strangers-2023.jpg
description: "One night in his near-empty tower block in contemporary London, Adam has a chance encounter with a mysterious neighbor Harry, which punctures the rhythm of his everyday life."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-fabelmans-2022">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../asteroid-city-2023">Next</a>
</nav>

<p>97 / 100</p>

<article class="film slug-all-of-us-strangers-2023">
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
