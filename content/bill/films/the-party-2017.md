---
title: "The Party"
layout: layouts/films.njk
slug: the-party-2017
ogImage: content/bill/films/backdrops/the-party-2017.jpg
description: "Various individuals think they’re coming together for a party in a private home, but a series of revelations results in a huge crisis that throws their belief systems – and their values – into total disarray."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../lucky-2017"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">63 / 100</a>
  </div>
  <div class="next">
    <a href="../roma-2018">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Lucky
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Roma
    </span>
  </div>
</nav>

<article class="film slug-the-party-2017">
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
  <img src="../films/profiles/1276.jpg" alt="Patricia Clarkson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Patricia Clarkson</span>
    <span class="cast-card-character">April</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1956.jpg" alt="Cherry Jones" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Cherry Jones</span>
    <span class="cast-card-character">Martha</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5470.jpg" alt="Kristin Scott Thomas" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Kristin Scott Thomas</span>
    <span class="cast-card-character">Janet</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2310.jpg" alt="Bruno Ganz" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Bruno Ganz</span>
    <span class="cast-card-character">Godfried</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9191.jpg" alt="Timothy Spall" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Timothy Spall</span>
    <span class="cast-card-character">Bill</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1246.jpg" alt="Emily Mortimer" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Emily Mortimer</span>
    <span class="cast-card-character">Jinny</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2037.jpg" alt="Cillian Murphy" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Cillian Murphy</span>
    <span class="cast-card-character">Tom</span>
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
    <li><a href="../mr-turner-2014">Mr. Turner</a> because of Timothy Spall</li>
  </ul>
</section>

</article>
