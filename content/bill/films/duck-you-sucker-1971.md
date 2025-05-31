---
title: "Duck, You Sucker"
layout: layouts/home.njk
slug: duck-you-sucker-1971
ogImage: content/bill/films/backdrops/duck-you-sucker-1971.jpg
description: "At the beginning of the 1913 Mexican Revolution, greedy bandit Juan Miranda and idealist John H. Mallory, an Irish Republican Army explosives expert on the lam from the British, fall in with a band of revolutionaries plotting to strike a national bank. When it turns out that the government has been using the bank as a hiding place for illegally detained political prisoners -- who are freed by the blast -- Miranda becomes a revolutionary hero against his will."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../butch-cassidy-and-the-sundance-kid-1969">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-sting-1973">Next</a>
</nav>

<p>15 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>Giù la testa</strong></p>

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
