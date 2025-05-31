---
title: "La Dolce Vita"
layout: layouts/home.njk
slug: la-dolce-vita-1960
ogImage: content/bill/films/backdrops/la-dolce-vita-1960.jpg
description: "Episodic journey of journalist Marcello who struggles to find his place in the world, torn between the allure of Rome's elite social scene and the stifling domesticity offered by his girlfriend, all the while searching for a way to become a serious writer."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../im-all-right-jack-1959">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../purple-noon-1960">Next</a>
</nav>

<p>6 / 100</p>

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
