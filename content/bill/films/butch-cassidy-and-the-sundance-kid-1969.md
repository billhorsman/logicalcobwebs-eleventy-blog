---
title: "Butch Cassidy and the Sundance Kid"
layout: layouts/home.njk
slug: butch-cassidy-and-the-sundance-kid-1969
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../once-upon-a-time-in-the-west-1968">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../duck-you-sucker-1971">Next</a>
</nav>

<p>14 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ film.slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ film.slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

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
