---
title: "Le Havre"
layout: layouts/home.njk
slug: le-havre-2011
ogImage: content/bill/films/backdrops/le-havre-2011.jpg
description: "In the French harbor city of Le Havre, fate throws young African refugee Idrissa into the path of Marcel Marx, a well-spoken bohemian who works as a shoe-shiner. With innate optimism and the tireless support of his community, Marcel stands up to officials pursuing the boy for deportation."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../tomboy-2011">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../all-is-lost-2013">Next</a>
</nav>

<p>60 / 100</p>

<article class="film slug-le-havre-2011">
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
