---
title: "Three Days of the Condor"
layout: layouts/home.njk
slug: three-days-of-the-condor-1975
ogImage: content/bill/films/backdrops/three-days-of-the-condor-1975.jpg
description: "When bookish CIA researcher Joe Turner finds all his co-workers dead, he, together with a woman he has kidnapped, must work together to outwit those responsible until he determines who he can really trust."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../dog-day-afternoon-1975">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-man-who-fell-to-earth-1976">Next</a>
</nav>

<p>18 / 100</p>

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
