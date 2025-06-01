---
title: "Delicatessen"
layout: layouts/home.njk
slug: delicatessen-1991
ogImage: content/bill/films/backdrops/delicatessen-1991.jpg
description: "In a post-apocalyptic world, the residents of an apartment above the butcher shop receive an occasional delicacy of meat, something that is in low supply. A young man new in town falls in love with the butcher's daughter, which causes conflicts in her family, who need the young man for other business-related purposes."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../withnail--i-1987">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../night-on-earth-1991">Next</a>
</nav>

<p>30 / 100</p>

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
