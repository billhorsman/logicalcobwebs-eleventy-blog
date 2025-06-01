---
title: "Between Two Worlds"
layout: layouts/home.njk
slug: between-two-worlds-2022
ogImage: content/bill/films/backdrops/between-two-worlds-2022.jpg
description: "Marianne Winckler relocates to the port city of Caen in order to pass herself off as a member of a large community of itinerant workers desperate to make ends meet. She gains employment as a cleaner on a ferry travelling between Ouistreham and Portsmouth, recording the drudgery of the work she and her colleagues are required to do."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-tragedy-of-macbeth-2021">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../eo-2022">Next</a>
</nav>

<p>91 / 100</p>

<article class="film slug-between-two-worlds-2022">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>Ouistreham</strong></p>

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
