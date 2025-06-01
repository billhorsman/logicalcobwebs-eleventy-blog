---
title: "The Power of the Dog"
layout: layouts/home.njk
slug: the-power-of-the-dog-2021
ogImage: content/bill/films/backdrops/the-power-of-the-dog-2021.jpg
description: "A domineering but charismatic rancher wages a war of intimidation on his brother's new wife and her teen son, until long-hidden secrets come to light."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-french-dispatch-2021">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../belfast-2021">Next</a>
</nav>

<p>86 / 100</p>

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
