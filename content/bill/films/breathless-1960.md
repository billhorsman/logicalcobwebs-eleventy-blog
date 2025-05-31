---
title: "Breathless"
layout: layouts/home.njk
slug: breathless-1960
ogImage: content/bill/films/backdrops/breathless-1960.jpg
description: "A small-time thief steals a car and impulsively murders a motorcycle policeman. Wanted by the authorities, he attempts to persuade a girl to run away to Italy with him."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../purple-noon-1960">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../barefoot-in-the-park-1967">Next</a>
</nav>

<p>8 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>À bout de souffle</strong></p>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {% if films.reviews[slug] %}
    <blockquote> 
      {{ films.reviews[slug] }} <em>— Bill</em>
    </blockquote> 
  {% endif %}

  <h2>
    Cast
  </h2>
  <ul>
    {%- for cast in film.credits.cast -%}
      <li>
        <strong>{{ cast.name }}</strong> as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
