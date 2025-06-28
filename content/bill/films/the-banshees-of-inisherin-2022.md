---
title: "The Banshees of Inisherin"
layout: layouts/films.njk
slug: the-banshees-of-inisherin-2022
ogImage: content/bill/films/backdrops/the-banshees-of-inisherin-2022.jpg
description: "Two lifelong friends find themselves at an impasse when one abruptly ends their relationship, with alarming consequences for both of them."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../one-fine-morning-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">94 / 100</a>
  </div>
  <div class="next">
    <a href="../the-fabelmans-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      One Fine Morning
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Fabelmans
    </span>
  </div>
</nav>

<article class="film slug-the-banshees-of-inisherin-2022">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    
  </p>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {%- if films.reviews[slug] -%}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href="/bill">Bill</a></em>
    </blockquote> 
  {%- endif -%}

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../phone-booth-2003">Phone Booth</a> by Colin Farrell</li>
<li><a href="../in-bruges-2008">In Bruges</a> by Colin Farrell, Brendan Gleeson and Martin McDonagh</li>
<li><a href="../the-tragedy-of-macbeth-2021">The Tragedy of Macbeth</a> by Brendan Gleeson</li>
  </ul>

  <section class="film-detail">
    <div>
      <details>
        <summary>
          <i class="fa-solid fa-masks-theater"></i>
          Cast
        </summary>
        <ul>
          {%- for cast in film.credits.cast -%}
            <li>
              {{ cast.name }} as <em>{{ cast.character }}</em>
            </li>
          {%- endfor -%}
        </ul>
      </details>
      <details>
        <summary>
          <i class="fa-solid fa-clapperboard"></i>
          Crew
        </summary>
        <ul>
          {%- for crew in film.credits.crew -%}
            <li>
              {{ crew.name }} &mdash; <em>{{ crew.job }}</em>
            </li>
          {%- endfor -%}
        </ul>
      </details>
    </div>
  </section>
</article>
