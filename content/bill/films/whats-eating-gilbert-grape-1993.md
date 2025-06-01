---
title: "What's Eating Gilbert Grape"
layout: layouts/home.njk
slug: whats-eating-gilbert-grape-1993
ogImage: content/bill/films/backdrops/whats-eating-gilbert-grape-1993.jpg
description: "Gilbert Grape is a small-town young man with a lot of responsibility. Chief among his concerns are his mother, who is so overweight that she can't leave the house, and his mentally impaired younger brother, Arnie, who has a knack for finding trouble. Settled into a job at a grocery store and an ongoing affair with local woman Betty Carver, Gilbert finally has his life shaken up by the free-spirited Becky."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-fugitive-1993">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../four-weddings-and-a-funeral-1994">Next</a>
</nav>

<p>33 / 100</p>

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
