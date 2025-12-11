---
title: "Mr. Turner"
layout: layouts/films.njk
slug: mr-turner-2014
ogImage: content/bill/films/backdrops/mr-turner-2014.jpg
description: "Eccentric British painter J.M.W. Turner  lives his last 25 years with gusto and secretly becomes involved with a seaside landlady, while his faithful housekeeper bears an unrequited love for him."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../dallas-buyers-club-2013"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">57 / 100</a>
  </div>
  <div class="next">
    <a href="../the-grand-budapest-hotel-2014">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Dallas Buyers Club
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Grand Budapest Hotel
    </span>
  </div>
</nav>

<article class="film slug-mr-turner-2014">
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
    <li><a href="../brazil-1985">Brazil</a> because of Roger Ashton-Griffiths</li>
<li><a href="../trainspotting-1996">Trainspotting</a> because of Stuart McQuarrie</li>
<li><a href="../the-bourne-identity-2002">The Bourne Identity</a> because of Vincent Franklin</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> because of Karl Johnson and Peter Wight</li>
<li><a href="../in-bruges-2008">In Bruges</a> because of Elizabeth Berrington</li>
<li><a href="../the-party-2017">The Party</a> because of Timothy Spall</li>
<li><a href="../little-women-2019">Little Women</a> because of James Norton</li>
  </ul>
</section>

</article>
