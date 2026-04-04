---
title: "Purple Noon"
layout: layouts/films.njk
slug: purple-noon-1960
ogImage: content/bill/films/backdrops/purple-noon-1960.jpg
description: "Tom Ripley is a talented mimic, moocher, forger and all-around criminal improviser; but there's more to Tom Ripley than even he can guess."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../la-dolce-vita-1960"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">6 / 100</a>
  </div>
  <div class="next">
    <a href="../in-the-heat-of-the-night-1967">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      La Dolce Vita
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      In the Heat of the Night
    </span>
  </div>
</nav>

<article class="film slug-purple-noon-1960">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Plein soleil.
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
  <img src="../films/profiles/15135.jpg" alt="Alain Delon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Alain Delon</span>
    <span class="cast-card-character">Tom Ripley</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39644.jpg" alt="Marie Laforêt" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Marie Laforêt</span>
    <span class="cast-card-character">Marge Duval</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/15395.jpg" alt="Maurice Ronet" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Maurice Ronet</span>
    <span class="cast-card-character">Philippe Greenleaf</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/37728.jpg" alt="Erno Crisa" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Erno Crisa</span>
    <span class="cast-card-character">Inspector Riccordi</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13733.jpg" alt="Frank Latimore" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Frank Latimore</span>
    <span class="cast-card-character">O'Brien</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/35958.jpg" alt="Billy Kearns" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Billy Kearns</span>
    <span class="cast-card-character">Freddy Miles</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/40165.jpg" alt="Ave Ninchi" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ave Ninchi</span>
    <span class="cast-card-character">Signora Gianna</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/64926.jpg" alt="Viviane Chantel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Viviane Chantel</span>
    <span class="cast-card-character">Belgian Lady</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/99530.jpg" alt="Nerio Bernardi" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Nerio Bernardi</span>
    <span class="cast-card-character">Agency Director</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Barbel Fanger</span>
    <span class="cast-card-character">Mr. Greenleaf</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Lily Romanelli</span>
    <span class="cast-card-character">Housekeeper</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Nicolas Petrov</span>
    <span class="cast-card-character">Boris</span>
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
