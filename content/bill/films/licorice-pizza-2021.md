---
title: "Licorice Pizza"
layout: layouts/films.njk
slug: licorice-pizza-2021
ogImage: content/bill/films/backdrops/licorice-pizza-2021.jpg
description: "The story of Gary Valentine and Alana Kane growing up, running around and going through the treacherous navigation of first love in the San Fernando Valley, 1973."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../house-of-gucci-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">82 / 100</a>
  </div>
  <div class="next">
    <a href="../nomadland-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      House of Gucci
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Nomadland
    </span>
  </div>
</nav>

<article class="film slug-licorice-pizza-2021">
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
    <li><a href="../magnolia-1999">Magnolia</a> because of John C. Reilly, Mark Flanagan, Paul Thomas Anderson and Brian Kehew</li>
<li><a href="../uncut-gems-2019">Uncut Gems</a> because of Benny Safdie</li>
<li><a href="../the-fabelmans-2022">The Fabelmans</a> because of Isabelle Kusman and Paige Locke</li>
  </ul>
</section>

</article>
