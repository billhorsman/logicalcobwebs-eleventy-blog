---
title: "La Strada"
layout: layouts/home.njk
slug: la-strada-1954
ogImage: content/bill/films/backdrops/la-strada-1954.jpg
description: "When Gelsomina, a naïve young woman, is purchased from her impoverished mother by brutish circus strongman Zampanò to be his wife and partner, she loyally endures her husband's coldness and abuse as they travel the Italian countryside performing together. Soon Zampanò must deal with his jealousy and conflicted feelings about Gelsomina when she finds a kindred spirit in Il Matto, the carefree circus fool, and contemplates leaving Zampanò."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../whisky-galore-1949"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../im-all-right-jack-1959">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>3 / 100</p>

<article class="film slug-la-strada-1954">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {%- if films.reviews[slug] -%}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href="/bill">Bill</a></em>
    </blockquote> 
  {%- endif -%}

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
