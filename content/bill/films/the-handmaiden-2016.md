---
title: "The Handmaiden"
layout: layouts/home.njk
slug: the-handmaiden-2016
ogImage: content/bill/films/backdrops/the-handmaiden-2016.jpg
description: "1930s Korea, in the period of Japanese occupation, a new girl, Sookee, is hired as a handmaiden to a Japanese heiress, Hideko, who lives a secluded life on a large countryside estate with her domineering Uncle Kouzuki. But the maid has a secret. She is a pickpocket recruited by a swindler posing as a Japanese Count to help him seduce the Lady to elope with him, rob her of her fortune, and lock her up in a madhouse. The plan seems to proceed according to plan until Sookee and Hideko discover some unexpected emotions."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../interstellar-2014">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../maudie-2016">Next</a>
</nav>

<p>66 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>아가씨</strong></p>

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
