---
title: "All Is Lost"
layout: layouts/films.njk
slug: all-is-lost-2013
ogImage: content/bill/films/backdrops/all-is-lost-2013.jpg
description: "During a solo voyage in the Indian Ocean, a veteran mariner awakes to find his vessel taking on water after a collision with a stray shipping container. With his radio and navigation equipment disabled, he sails unknowingly into a violent storm and barely escapes with his life. With any luck, the ocean currents may carry him into a shipping lane -- but, with supplies dwindling and the sharks circling, the sailor is forced to face his own mortality."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../kill-bill-the-whole-bloody-affair-2011"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">58 / 100</a>
  </div>
  <div class="next">
    <a href="../dallas-buyers-club-2013">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Kill Bill: The Whole Bloody Affair
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Dallas Buyers Club
    </span>
  </div>
</nav>

<article class="film slug-all-is-lost-2013">
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
    <li><a href="../butch-cassidy-and-the-sundance-kid-1969">Butch Cassidy and the Sundance Kid</a>, <a href="../the-sting-1973">The Sting</a> and <a href="../three-days-of-the-condor-1975">Three Days of the Condor</a> because of Robert Redford</li>
  </ul>
</section>

</article>
