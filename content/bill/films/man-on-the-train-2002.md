---
title: "Man on the Train"
layout: layouts/home.njk
slug: man-on-the-train-2002
ogImage: content/bill/films/backdrops/man-on-the-train-2002.jpg
description: "A man, Milan steps off a train, into a small French village. As he waits for the day when he will rob the town bank, he runs into an old retired poetry teacher named M. Manesquier. The two men strike up a strange friendship and explore the road not taken, each wanting to live the other's life."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../24-hour-party-people-2002"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../the-bourne-identity-2002">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>47 / 100</p>

<article class="film slug-man-on-the-train-2002">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as L'Homme du train.
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
