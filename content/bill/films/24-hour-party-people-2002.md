---
title: "24 Hour Party People"
layout: layouts/films.njk
slug: 24-hour-party-people-2002
ogImage: content/bill/films/backdrops/24-hour-party-people-2002.jpg
description: "Manchester, 1976. Tony Wilson is an ambitious but frustrated local TV news reporter looking for a way to make his mark. After witnessing a life-changing concert by a band known as the Sex Pistols, he persuades his station to televise one of their performances, and soon Manchester's punk groups are clamoring for him to manage them. Riding the wave of a musical revolution, Wilson and his friends create the legendary Factory Records label and The Hacienda club."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../black-hawk-down-2001"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">44 / 100</a>
  </div>
  <div class="next">
    <a href="../man-on-the-train-2002">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Black Hawk Down
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Man on the Train
    </span>
  </div>
</nav>

<article class="film slug-24-hour-party-people-2002">
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
  <li><a href="../shallow-grave-1994">Shallow Grave</a> by Christopher Eccleston and Keith Allen</li>
<li><a href="../trainspotting-1996">Trainspotting</a> by Keith Allen and Shirley Henderson</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> by Steve Coogan, Paddy Considine, Simon Pegg and Ron Cook</li>
<li><a href="../empire-of-light-2022">Empire of Light</a> by Ron Cook</li>
<li><a href="../dune-2021">Dune</a> by Neil Bell</li>
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
