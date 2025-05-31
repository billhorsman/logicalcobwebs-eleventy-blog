---
title: "All Is Lost"
layout: layouts/home.njk
slug: all-is-lost-2013
ogImage: content/bill/films/backdrops/all-is-lost-2013.jpg
description: "During a solo voyage in the Indian Ocean, a veteran mariner awakes to find his vessel taking on water after a collision with a stray shipping container. With his radio and navigation equipment disabled, he sails unknowingly into a violent storm and barely escapes with his life. With any luck, the ocean currents may carry him into a shipping lane -- but, with supplies dwindling and the sharks circling, the sailor is forced to face his own mortality."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../le-havre-2011">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../dallas-buyers-club-2013">Next</a>
</nav>

<p>61 / 100</p>

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
