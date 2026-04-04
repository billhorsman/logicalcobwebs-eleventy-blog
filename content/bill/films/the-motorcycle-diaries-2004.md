---
title: "The Motorcycle Diaries"
layout: layouts/films.njk
slug: the-motorcycle-diaries-2004
ogImage: content/bill/films/backdrops/the-motorcycle-diaries-2004.jpg
description: "Based on the journals of Che Guevara, leader of the Cuban Revolution. In his memoirs, Guevara recounts adventures he and best friend Alberto Granado had while crossing South America by motorcycle in the early 1950s."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../phone-booth-2003"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">44 / 100</a>
  </div>
  <div class="next">
    <a href="../hot-fuzz-2007">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Phone Booth
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Hot Fuzz
    </span>
  </div>
</nav>

<article class="film slug-the-motorcycle-diaries-2004">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Diarios de motocicleta.
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
  <img src="../films/profiles/258.jpg" alt="Gael García Bernal" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gael García Bernal</span>
    <span class="cast-card-character">Ernesto "Che" Guevara de la Serna</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/18478.jpg" alt="Rodrigo de la Serna" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rodrigo de la Serna</span>
    <span class="cast-card-character">Alberto Granado</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/18499.jpg" alt="Mercedes Morán" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Mercedes Morán</span>
    <span class="cast-card-character">Celia de la Serna</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6859.jpg" alt="Mía Maestro" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Mía Maestro</span>
    <span class="cast-card-character">Chichina Ferreyra</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/105738.jpg" alt="Jean Pierre Noher" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean Pierre Noher</span>
    <span class="cast-card-character">Ernesto Guevara Lynch</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Lucas Oro</span>
    <span class="cast-card-character">Roberto Guevara</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/64772.jpg" alt="Marina Glezer" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Marina Glezer</span>
    <span class="cast-card-character">Celita Guevara</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/929483.jpg" alt="Sofia Bertolotto" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sofia Bertolotto</span>
    <span class="cast-card-character">Ana María Guevara</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Franco Solazzi</span>
    <span class="cast-card-character">Juan Martín Guevara</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1313069.jpg" alt="Ricardo Díaz Mourelle" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ricardo Díaz Mourelle</span>
    <span class="cast-card-character">Uncle Jorge</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/147327.jpg" alt="Gustavo Bueno" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gustavo Bueno</span>
    <span class="cast-card-character">Dr. Hugo Pesce</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/125493.jpg" alt="Antonella Costa" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Antonella Costa</span>
    <span class="cast-card-character">Silvia</span>
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
