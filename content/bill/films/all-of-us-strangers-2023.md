---
title: "All of Us Strangers"
layout: layouts/films.njk
slug: all-of-us-strangers-2023
ogImage: content/bill/films/backdrops/all-of-us-strangers-2023.jpg
description: "One night in his near-empty tower block in contemporary London, Adam has a chance encounter with a mysterious neighbor Harry, which punctures the rhythm of his everyday life."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../anatomy-of-a-fall-2023"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">94 / 100</a>
  </div>
  <div class="next">
    <a href="../perfect-days-2023">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Anatomy of a Fall
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Perfect Days
    </span>
  </div>
</nav>

<article class="film slug-all-of-us-strangers-2023">
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
  <img src="../films/profiles/125660.jpg" alt="Andrew Scott" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Andrew Scott</span>
    <span class="cast-card-character">Adam</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2326151.jpg" alt="Paul Mescal" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Paul Mescal</span>
    <span class="cast-card-character">Harry</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Carter John Grout</span>
    <span class="cast-card-character">Young Adam</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/478.jpg" alt="Jamie Bell" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jamie Bell</span>
    <span class="cast-card-character">Dad</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/120932.jpg" alt="Claire Foy" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Claire Foy</span>
    <span class="cast-card-character">Mum</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3893524.jpg" alt="Ami Tredrea" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ami Tredrea</span>
    <span class="cast-card-character">Waitress</span>
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
    <li><a href="../billy-elliot-2000">Billy Elliot</a> because of Jamie Bell</li>
<li><a href="../1917-2019">1917</a> because of Andrew Scott</li>
  </ul>
</section>

</article>
