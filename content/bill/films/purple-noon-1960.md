---
title: "Purple Noon"
layout: layouts/home.njk
slug: purple-noon-1960
ogImage: content/bill/films/backdrops/purple-noon-1960.jpg
description: "Tom Ripley is a talented mimic, moocher, forger and all-around criminal improviser; but there's more to Tom Ripley than even he can guess."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../la-dolce-vita-1960">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../breathless-1960">Next</a>
</nav>

<p>7 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>Plein soleil</strong></p>

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
