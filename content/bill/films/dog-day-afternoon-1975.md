---
title: "Dog Day Afternoon"
layout: layouts/home.njk
slug: dog-day-afternoon-1975
ogImage: content/bill/films/backdrops/dog-day-afternoon-1975.jpg
description: "Based on the true story of would-be Brooklyn bank robbers John Wojtowicz and Salvatore Naturile. Sonny and Sal attempt a bank heist which quickly turns sour and escalates into a hostage situation and stand-off with the police. As Sonny's motives for the robbery are slowly revealed and things become more complicated, the heist turns into a media circus."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-sting-1973">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../three-days-of-the-condor-1975">Next</a>
</nav>

<p>17 / 100</p>

<article class="film">
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
