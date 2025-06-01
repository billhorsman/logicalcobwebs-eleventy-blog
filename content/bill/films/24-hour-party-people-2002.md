---
title: "24 Hour Party People"
layout: layouts/home.njk
slug: 24-hour-party-people-2002
ogImage: content/bill/films/backdrops/24-hour-party-people-2002.jpg
description: "Manchester, 1976. Tony Wilson is an ambitious but frustrated local TV news reporter looking for a way to make his mark. After witnessing a life-changing concert by a band known as the Sex Pistols, he persuades his station to televise one of their performances, and soon Manchester's punk groups are clamoring for him to manage them. Riding the wave of a musical revolution, Wilson and his friends create the legendary Factory Records label and The Hacienda club."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../black-hawk-down-2001">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-bourne-identity-2002">Next</a>
</nav>

<p>47 / 100</p>

<article class="film slug-24-hour-party-people-2002">
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
