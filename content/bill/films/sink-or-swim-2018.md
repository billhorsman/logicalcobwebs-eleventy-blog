---
title: "Sink or Swim"
layout: layouts/home.njk
slug: sink-or-swim-2018
ogImage: content/bill/films/backdrops/sink-or-swim-2018.jpg
description: "40-year-old Bertrand has been suffering from depression for the last two years and is barely able to keep his head above water. Despite the medication he gulps down all day, every day, and his wife's encouragement, he is unable to find any meaning in his life. Curiously, he will end up finding this sense of purpose at the swimming pool, by joining an all-male synchronised swimming team."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../roma-2018"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a href="../">Film list</a>
  </div>
  <div class="next">
    <a href="../woman-at-war-2018">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
</nav>

<p>71 / 100</p>

<article class="film slug-sink-or-swim-2018">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>Le Grand Bain</strong></p>

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
