---
title: "The Fabelmans"
layout: layouts/films.njk
slug: the-fabelmans-2022
ogImage: content/bill/films/backdrops/the-fabelmans-2022.jpg
description: "Growing up in post-World War II era Arizona, young Sammy Fabelman aspires to become a filmmaker as he reaches adolescence, but soon discovers a shattering family secret and explores how the power of films can help him see the truth."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-banshees-of-inisherin-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">92 / 100</a>
  </div>
  <div class="next">
    <a href="../anatomy-of-a-fall-2023">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Banshees of Inisherin
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Anatomy of a Fall
    </span>
  </div>
</nav>

<article class="film slug-the-fabelmans-2022">
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
  <img src="../films/profiles/1812.jpg" alt="Michelle Williams" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Michelle Williams</span>
    <span class="cast-card-character">Mitzi Fabelman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17142.jpg" alt="Paul Dano" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Paul Dano</span>
    <span class="cast-card-character">Burt Fabelman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/19274.jpg" alt="Seth Rogen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Seth Rogen</span>
    <span class="cast-card-character">Benny Loewy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1476330.jpg" alt="Gabriel LaBelle" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gabriel LaBelle</span>
    <span class="cast-card-character">Sammy Fabelman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3731015.jpg" alt="Mateo Zoryan Francis-DeFord" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Mateo Zoryan Francis-DeFord</span>
    <span class="cast-card-character">Younger Sammy Fabelman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3690404.jpg" alt="Keeley Karsten" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Keeley Karsten</span>
    <span class="cast-card-character">Natalie Fabelman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3690411.jpg" alt="Alina Brace" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Alina Brace</span>
    <span class="cast-card-character">Young Natalie Fabelman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1651362.jpg" alt="Julia Butters" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Julia Butters</span>
    <span class="cast-card-character">Reggie Fabelman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3690410.jpg" alt="Birdie Borria" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Birdie Borria</span>
    <span class="cast-card-character">Younger Reggie Fabelman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6167.jpg" alt="Judd Hirsch" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Judd Hirsch</span>
    <span class="cast-card-character">Uncle Boris</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3233032.jpg" alt="Sophia Kopera" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sophia Kopera</span>
    <span class="cast-card-character">Lisa Fabelman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/124377.jpg" alt="Jeannie Berlin" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jeannie Berlin</span>
    <span class="cast-card-character">Hadassah Fabelman</span>
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
    <li><a href="../fight-club-1999">Fight Club</a> and <a href="../magnolia-1999">Magnolia</a> because of Ezra Buzzington</li>
<li><a href="../the-straight-story-1999">The Straight Story</a> and <a href="../lucky-2017">Lucky</a> because of David Lynch</li>
<li><a href="../uncut-gems-2019">Uncut Gems</a> because of Judd Hirsch</li>
<li><a href="../licorice-pizza-2021">Licorice Pizza</a> because of Isabelle Kusman and Paige Locke</li>
  </ul>
</section>

</article>
