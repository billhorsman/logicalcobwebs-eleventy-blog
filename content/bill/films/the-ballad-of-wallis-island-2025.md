---
title: "The Ballad of Wallis Island"
layout: layouts/films.njk
slug: the-ballad-of-wallis-island-2025
ogImage: content/bill/films/backdrops/the-ballad-of-wallis-island-2025.jpg
description: "Eccentric lottery winner Charles lives alone on a remote island but dreams of hiring his favourite musician, Herb McGwyer, to play an exclusive, private gig. Unbeknownst to Herb, Charles has also hired Herb’s ex-bandmate and ex-girlfriend, Nell, with her new husband in town, to perform the old favourites. As tempers flare and old tensions resurface, the stormy weather traps them all on the island and Charles desperately looks for a way to salvage his dream gig."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../blue-jean-2023"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">98 / 100</a>
  </div>
  <div class="next">
    <a href="../nouvelle-vague-2025">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Blue Jean
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Nouvelle Vague
    </span>
  </div>
</nav>

<article class="film slug-the-ballad-of-wallis-island-2025">
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
  <img src="../films/profiles/134229.jpg" alt="Tom Basden" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tom Basden</span>
    <span class="cast-card-character">Herb McGwyer</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1016119.jpg" alt="Tim Key" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tim Key</span>
    <span class="cast-card-character">Charles Heath</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/36662.jpg" alt="Carey Mulligan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Carey Mulligan</span>
    <span class="cast-card-character">Nell Mortimer</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1648728.jpg" alt="Sian Clifford" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sian Clifford</span>
    <span class="cast-card-character">Amanda</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/550396.jpg" alt="Akemnji Ndifornyen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Akemnji Ndifornyen</span>
    <span class="cast-card-character">Michael</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Steve Marsh</span>
    <span class="cast-card-character">Peter</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Luka Downie</span>
    <span class="cast-card-character">Marcus</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Kerrie Thomason</span>
    <span class="cast-card-character">Marie</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Arron Long</span>
    <span class="cast-card-character">Tour guide (uncredited)</span>
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
