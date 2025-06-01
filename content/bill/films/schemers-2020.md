---
title: "Schemers"
layout: layouts/home.njk
slug: schemers-2020
ogImage: content/bill/films/backdrops/schemers-2020.jpg
description: "Set in late-1970s Dundee, Schemers is based on writer-producer David McLean’s early years in the music business. After a run in with a local gangster, a fledgling promoter and his two friends raise their ambitions to booking major bands in order to escape their debt."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../first-cow-2020">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../eternal-beauty-2020">Next</a>
</nav>

<p>76 / 100</p>

<article class="film slug-schemers-2020">
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
