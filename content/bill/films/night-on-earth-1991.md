---
title: "Night on Earth"
layout: layouts/films.njk
slug: night-on-earth-1991
ogImage: content/bill/films/backdrops/night-on-earth-1991.jpg
description: "An anthology of 5 different cab drivers in 5 American and European cities and their remarkable fares on the same eventful night."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../delicatessen-1991"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">25 / 100</a>
  </div>
  <div class="next">
    <a href="../lon-the-professional-1994">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Delicatessen
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Léon: The Professional
    </span>
  </div>
</nav>

<article class="film slug-night-on-earth-1991">
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
  <img src="../films/profiles/1920.jpg" alt="Winona Ryder" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Winona Ryder</span>
    <span class="cast-card-character">Corky</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4800.jpg" alt="Gena Rowlands" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gena Rowlands</span>
    <span class="cast-card-character">Victoria Snelling</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4808.jpg" alt="Giancarlo Esposito" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Giancarlo Esposito</span>
    <span class="cast-card-character">YoYo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/12647.jpg" alt="Armin Mueller-Stahl" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Armin Mueller-Stahl</span>
    <span class="cast-card-character">Helmut Grokenberger</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4810.jpg" alt="Rosie Perez" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rosie Perez</span>
    <span class="cast-card-character">Angela</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4812.jpg" alt="Isaach de Bankolé" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Isaach de Bankolé</span>
    <span class="cast-card-character">Cab Driver Paris</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4813.jpg" alt="Béatrice Dalle" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Béatrice Dalle</span>
    <span class="cast-card-character">Blind Woman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4818.jpg" alt="Roberto Benigni" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Roberto Benigni</span>
    <span class="cast-card-character">Cab Driver Rome</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4819.jpg" alt="Paolo Bonacelli" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Paolo Bonacelli</span>
    <span class="cast-card-character">Priest</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4826.jpg" alt="Matti Pellonpää" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Matti Pellonpää</span>
    <span class="cast-card-character">Mika</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/122402.jpg" alt="Kari Väänänen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kari Väänänen</span>
    <span class="cast-card-character">Man #1 Helsinki</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4828.jpg" alt="Sakari Kuosmanen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sakari Kuosmanen</span>
    <span class="cast-card-character">Man #2 Helsinki</span>
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
    <li><a href="../ghost-dog-the-way-of-the-samurai-1999">Ghost Dog: The Way of the Samurai</a> because of Isaach de Bankolé and Jim Jarmusch</li>
  </ul>
</section>

</article>
