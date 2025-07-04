---
title: "Fight Club"
layout: layouts/films.njk
slug: fight-club-1999
ogImage: content/bill/films/backdrops/fight-club-1999.jpg
description: "A ticking-time-bomb insomniac and a slippery soap salesman channel primal male aggression into a shocking new form of therapy. Their concept catches on, with underground \"fight clubs\" forming in every town, until an eccentric gets in the way and ignites an out-of-control spiral toward oblivion."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-big-lebowski-1998"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">34 / 100</a>
  </div>
  <div class="next">
    <a href="../ghost-dog-the-way-of-the-samurai-1999">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Big Lebowski
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Ghost Dog: The Way of the Samurai
    </span>
  </div>
</nav>

<article class="film slug-fight-club-1999">
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

  <section class="related-films">
  <h2>Related films</h2>
  <ul>
    <li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a>, <a href="../the-french-dispatch-2021">The French Dispatch</a> and <a href="../asteroid-city-2023">Asteroid City</a> because of Edward Norton</li>
<li><a href="../phone-booth-2003">Phone Booth</a>, <a href="../dallas-buyers-club-2013">Dallas Buyers Club</a> and <a href="../house-of-gucci-2021">House of Gucci</a> because of Jared Leto</li>
<li><a href="../magnolia-1999">Magnolia</a> because of Ezra Buzzington, Michael Shamus Wiles, Greg Bronson and Phil Hawn</li>
<li><a href="../the-fabelmans-2022">The Fabelmans</a> because of Ezra Buzzington</li>
<li><a href="../lady-bird-2017">Lady Bird</a> because of Bob Stephenson</li>
  </ul>
</section>

</article>
