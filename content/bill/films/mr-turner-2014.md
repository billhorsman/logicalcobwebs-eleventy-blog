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
    <a class="simple" href="../">55 / 100</a>
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

  <section class="cast-grid">
  <div class="cast-grid-cards">
    <div class="cast-card">
  <img src="../films/profiles/9191.jpg" alt="Timothy Spall" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Timothy Spall</span>
    <span class="cast-card-character">JMW Turner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1221056.jpg" alt="Dorothy Atkinson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Dorothy Atkinson</span>
    <span class="cast-card-character">Hannah Danby</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/72307.jpg" alt="Marion Bailey" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Marion Bailey</span>
    <span class="cast-card-character">Sophia Booth</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/72308.jpg" alt="Paul Jesson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Paul Jesson</span>
    <span class="cast-card-character">William Turner Snr</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/72305.jpg" alt="Lesley Manville" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Lesley Manville</span>
    <span class="cast-card-character">Mary Somerville</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/30328.jpg" alt="Martin Savage" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Martin Savage</span>
    <span class="cast-card-character">Benjamin Robert Haydon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/53367.jpg" alt="Ruth Sheen" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ruth Sheen</span>
    <span class="cast-card-character">Sarah Danby</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/133032.jpg" alt="David Horovitch" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">David Horovitch</span>
    <span class="cast-card-character">Dr Price</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39186.jpg" alt="Karl Johnson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Karl Johnson</span>
    <span class="cast-card-character">Mr. Booth</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17476.jpg" alt="Peter Wight" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Wight</span>
    <span class="cast-card-character">Joseph Gillot</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1394375.jpg" alt="Joshua McGuire" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Joshua McGuire</span>
    <span class="cast-card-character">John Ruskin</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2258.jpg" alt="Stuart McQuarrie" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Stuart McQuarrie</span>
    <span class="cast-card-character">Ruskin's Father</span>
  </div>
</div>
  </div>
</section>

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
