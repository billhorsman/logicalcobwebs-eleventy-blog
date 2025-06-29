---
title: "Uncut Gems"
layout: layouts/films.njk
slug: uncut-gems-2019
ogImage: content/bill/films/backdrops/uncut-gems-2019.jpg
description: "A charismatic New York City jeweler always on the lookout for the next big score makes a series of high-stakes bets that could lead to the windfall of a lifetime. Howard must perform a precarious high-wire act, balancing business, family, and encroaching adversaries on all sides in his relentless pursuit of the ultimate win."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-lighthouse-2019"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">73 / 100</a>
  </div>
  <div class="next">
    <a href="../eternal-beauty-2020">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Lighthouse
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Eternal Beauty
    </span>
  </div>
</nav>

<article class="film slug-uncut-gems-2019">
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
    <li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> and <a href="../the-french-dispatch-2021">The French Dispatch</a> because of Tilda Swinton</li>
<li><a href="../asteroid-city-2023">Asteroid City</a> because of Tilda Swinton and Jake Ryan</li>
<li><a href="../the-fabelmans-2022">The Fabelmans</a> because of Judd Hirsch</li>
<li><a href="../licorice-pizza-2021">Licorice Pizza</a> because of Benny Safdie</li>
  </ul>
</section>

</article>
