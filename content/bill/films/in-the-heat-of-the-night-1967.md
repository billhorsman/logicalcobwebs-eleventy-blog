---
title: "In the Heat of the Night"
layout: layouts/home.njk
slug: in-the-heat-of-the-night-1967
ogImage: content/bill/films/backdrops/in-the-heat-of-the-night-1967.jpg
description: "African-American Philadelphia police detective Virgil Tibbs is arrested on suspicion of murder by Bill Gillespie, the racist police chief of tiny Sparta, Mississippi. After Tibbs proves not only his own innocence but that of another man, he joins forces with Gillespie to track down the real killer. Their investigation takes them through every social level of the town, with Tibbs making enemies as well as unlikely friends as he hunts for the truth."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../barefoot-in-the-park-1967">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../2001-a-space-odyssey-1968">Next</a>
</nav>

<p>10 / 100</p>

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
