---
title: "Local Hero"
layout: layouts/films.njk
slug: local-hero-1983
ogImage: content/bill/films/backdrops/local-hero-1983.jpg
description: "An American oil company sends a man to Scotland to buy up an entire village where they want to build a refinery. But things don't go as expected."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../blade-runner-1982"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">20 / 100</a>
  </div>
  <div class="next">
    <a href="../paris-texas-1984">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Blade Runner
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Paris, Texas
    </span>
  </div>
</nav>

<article class="film slug-local-hero-1983">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
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
  <img src="../films/profiles/13784.jpg" alt="Burt Lancaster" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Burt Lancaster</span>
    <span class="cast-card-character">Felix Happer</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/20899.jpg" alt="Peter Riegert" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Riegert</span>
    <span class="cast-card-character">Mac</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/47698.jpg" alt="Denis Lawson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Denis Lawson</span>
    <span class="cast-card-character">Urquhart</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/68686.jpg" alt="Fulton Mackay" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Fulton Mackay</span>
    <span class="cast-card-character">Ben</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/12982.jpg" alt="Peter Capaldi" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Capaldi</span>
    <span class="cast-card-character">Oldsen</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/116125.jpg" alt="Jennifer Black" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jennifer Black</span>
    <span class="cast-card-character">Stella</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/36820.jpg" alt="Jenny Seagrove" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jenny Seagrove</span>
    <span class="cast-card-character">Marina</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/116123.jpg" alt="Norman Chancer" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Norman Chancer</span>
    <span class="cast-card-character">Moritz</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/105964.jpg" alt="Rikki Fulton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rikki Fulton</span>
    <span class="cast-card-character">Geddes</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2451.jpg" alt="Alex Norton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Alex Norton</span>
    <span class="cast-card-character">Watt</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/116126.jpg" alt="Christopher Rozycki" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Christopher Rozycki</span>
    <span class="cast-card-character">Victor</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Gyearbuor Asante</span>
    <span class="cast-card-character">Rev Macpherson</span>
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

  
</article>
