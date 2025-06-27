---
title: "House of Gucci"
layout: layouts/films.njk
slug: house-of-gucci-2021
ogImage: content/bill/films/backdrops/house-of-gucci-2021.jpg
description: "When Patrizia Reggiani, an outsider from humble beginnings, marries into the Gucci family, her unbridled ambition begins to unravel the family legacy and triggers a reckless spiral of betrayal, decadence, revenge, and ultimately… murder."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../dune-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">82 / 100</a>
  </div>
  <div class="next">
    <a href="../licorice-pizza-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Dune
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Licorice Pizza
    </span>
  </div>
</nav>

<article class="film slug-house-of-gucci-2021">
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
  <li><a href="../dog-day-afternoon-1975">Dog Day Afternoon</a> (1975) by Al Pacino</li>
<li><a href="../fight-club-1999">Fight Club</a> (1999) by Jared Leto</li>
<li><a href="../phone-booth-2003">Phone Booth</a> (2003) by Jared Leto</li>
<li><a href="../dallas-buyers-club-2013">Dallas Buyers Club</a> (2013) by Jared Leto</li>
<li><a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> (1999) by Pietro Ragusa</li>
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
