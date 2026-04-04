---
title: "Between Two Worlds"
layout: layouts/films.njk
slug: between-two-worlds-2022
ogImage: content/bill/films/backdrops/between-two-worlds-2022.jpg
description: "Marianne Winckler relocates to the port city of Caen in order to pass herself off as a member of a large community of itinerant workers desperate to make ends meet. She gains employment as a cleaner on a ferry travelling between Ouistreham and Portsmouth, recording the drudgery of the work she and her colleagues are required to do."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-tragedy-of-macbeth-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">87 / 100</a>
  </div>
  <div class="next">
    <a href="../eo-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Tragedy of Macbeth
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      EO
    </span>
  </div>
</nav>

<article class="film slug-between-two-worlds-2022">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Ouistreham.
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
  <img src="../films/profiles/1137.jpg" alt="Juliette Binoche" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Juliette Binoche</span>
    <span class="cast-card-character">Marianne Winckler</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Hélène Lambert</span>
    <span class="cast-card-character">Chrystèle</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Louise Pociecka</span>
    <span class="cast-card-character">Louise</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Steve Papagiannis</span>
    <span class="cast-card-character">Steve</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Jérémy Lechevallier</span>
    <span class="cast-card-character">Eric</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Kévin Maspimby</span>
    <span class="cast-card-character">Kévin</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Faïçal Zoua</span>
    <span class="cast-card-character">Faïçal</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Arnaud Duval</span>
    <span class="cast-card-character">M. Mathieu</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Didier Pupin</span>
    <span class="cast-card-character">Cédric</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Léa Carne</span>
    <span class="cast-card-character">Marilou</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Nathalie Lecornu</span>
    <span class="cast-card-character">Nathalie</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Joël Graindorge</span>
    <span class="cast-card-character">Joël</span>
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
