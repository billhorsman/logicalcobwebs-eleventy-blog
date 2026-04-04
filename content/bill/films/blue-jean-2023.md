---
title: "Blue Jean"
layout: layouts/films.njk
slug: blue-jean-2023
ogImage: content/bill/films/backdrops/blue-jean-2023.jpg
description: "Jean, a PE teacher, is forced to live a double life. When a new student arrives and threatens to expose her sexuality, Jean is pushed to extreme lengths to keep her job and her integrity."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../asteroid-city-2023"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">97 / 100</a>
  </div>
  <div class="next">
    <a href="../the-ballad-of-wallis-island-2025">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Asteroid City
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Ballad of Wallis Island
    </span>
  </div>
</nav>

<article class="film slug-blue-jean-2023">
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
  <img src="../films/profiles/2648396.jpg" alt="Rosy McEwen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rosy McEwen</span>
    <span class="cast-card-character">Jean Newman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/570323.jpg" alt="Kerrie Hayes" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kerrie Hayes</span>
    <span class="cast-card-character">Vivian Highton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3536817.jpg" alt="Lucy Halliday" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lucy Halliday</span>
    <span class="cast-card-character">Lois Jackson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2664861.jpg" alt="Lydia Page" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lydia Page</span>
    <span class="cast-card-character">Siobhan Murphy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3163376.jpg" alt="Becky Lindsay" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Becky Lindsay</span>
    <span class="cast-card-character">Jill</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3924040.jpg" alt="Maya Torres" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Maya Torres</span>
    <span class="cast-card-character">Mindy Singh</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3924038.jpg" alt="Ellen Gowland" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ellen Gowland</span>
    <span class="cast-card-character">Carol Ridley</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2335665.jpg" alt="Amy Booth-Steel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Amy Booth-Steel</span>
    <span class="cast-card-character">Debbie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3063466.jpg" alt="Stacy Abalogun" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Stacy Abalogun</span>
    <span class="cast-card-character">Ace</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3640536.jpg" alt="Izzy Neish" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Izzy Neish</span>
    <span class="cast-card-character">Abi</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/562183.jpg" alt="Kate Soulsby" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kate Soulsby</span>
    <span class="cast-card-character">Joni</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2969577.jpg" alt="Lainey Shaw" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lainey Shaw</span>
    <span class="cast-card-character">Paula</span>
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
