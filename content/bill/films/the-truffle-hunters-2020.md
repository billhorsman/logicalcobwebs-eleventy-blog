---
title: "The Truffle Hunters"
layout: layouts/home.njk
slug: the-truffle-hunters-2020
ogImage: content/bill/films/backdrops/the-truffle-hunters-2020.jpg
description: "In the secret forests of Northern Italy, a dwindling group of joyful old men and their faithful dogs search for the world’s most expensive ingredient, the white Alba truffle. Their stories form a real-life fairy tale that celebrates human passion in a fragile land that seems forgotten in time."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../limbo-2020">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../nomadland-2021">Next</a>
</nav>

<p>79 / 100</p>

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
        <strong>{{ cast.name }}</strong> as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
