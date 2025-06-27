---
title: "The Tragedy of Macbeth"
layout: layouts/films.njk
slug: the-tragedy-of-macbeth-2021
ogImage: content/bill/films/backdrops/the-tragedy-of-macbeth-2021.jpg
description: "Macbeth, the Thane of Glamis, receives a prophecy from a trio of witches that one day he will become King of Scotland. Consumed by ambition and spurred to action by his wife, Macbeth murders his king and takes the throne for himself."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-power-of-the-dog-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">89 / 100</a>
  </div>
  <div class="next">
    <a href="../between-two-worlds-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Power of the Dog
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Between Two Worlds
    </span>
  </div>
</nav>

<article class="film slug-the-tragedy-of-macbeth-2021">
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
  <li><a href="../fargo-1996">Fargo</a> (1996) by Frances McDormand</li>
<li><a href="../nomadland-2021">Nomadland</a> (2021) by Frances McDormand</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> (2021) by Frances McDormand</li>
<li><a href="../no-country-for-old-men-2007">No Country for Old Men</a> (2007) by Stephen Root</li>
<li><a href="../in-bruges-2008">In Bruges</a> (2008) by Brendan Gleeson</li>
<li><a href="../the-banshees-of-inisherin-2022">The Banshees of Inisherin</a> (2022) by Brendan Gleeson</li>
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
