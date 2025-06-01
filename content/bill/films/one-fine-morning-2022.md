---
title: "One Fine Morning"
layout: layouts/home.njk
slug: one-fine-morning-2022
ogImage: content/bill/films/backdrops/one-fine-morning-2022.jpg
description: "With a father suffering from neurodegenerative disease, a young woman lives with her eight-year-old daughter. While struggling to secure a decent nursing home, she runs into an unavailable friend with whom she embarks on an affair."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../empire-of-light-2022">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-banshees-of-inisherin-2022">Next</a>
</nav>

<p>94 / 100</p>

<article class="film slug-one-fine-morning-2022">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>Un beau matin</strong></p>

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
