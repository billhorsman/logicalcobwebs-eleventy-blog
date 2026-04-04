---
title: "Anatomy of a Fall"
layout: layouts/films.njk
slug: anatomy-of-a-fall-2023
ogImage: content/bill/films/backdrops/anatomy-of-a-fall-2023.jpg
description: "A woman is suspected of her husband's murder, and their blind son faces a moral dilemma as the sole witness."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-fabelmans-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">93 / 100</a>
  </div>
  <div class="next">
    <a href="../all-of-us-strangers-2023">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Fabelmans
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      All of Us Strangers
    </span>
  </div>
</nav>

<article class="film slug-anatomy-of-a-fall-2023">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Anatomie d'une chute.
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
  <img src="../films/profiles/7152.jpg" alt="Sandra Hüller" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sandra Hüller</span>
    <span class="cast-card-character">Sandra Voyter</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/145120.jpg" alt="Swann Arlaud" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Swann Arlaud</span>
    <span class="cast-card-character">Maître Vincent Renzi</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3302971.jpg" alt="Milo Machado-Graner" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Milo Machado-Graner</span>
    <span class="cast-card-character">Daniel</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1795152.jpg" alt="Antoine Reinartz" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Antoine Reinartz</span>
    <span class="cast-card-character">Advocate General</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1312061.jpg" alt="Samuel Theis" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Samuel Theis</span>
    <span class="cast-card-character">Samuel Maleski</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1029285.jpg" alt="Jehnny Beth" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jehnny Beth</span>
    <span class="cast-card-character">Marge Berger</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1613123.jpg" alt="Saadia Bentaïeb" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Saadia Bentaïeb</span>
    <span class="cast-card-character">Maître Nour Boudaoud</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1033745.jpg" alt="Camille Rutherford" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Camille Rutherford</span>
    <span class="cast-card-character">Zoé Solidor</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2670733.jpg" alt="Anne Rotger" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Anne Rotger</span>
    <span class="cast-card-character">President of the Court</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/233491.jpg" alt="Sophie Fillières" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Sophie Fillières</span>
    <span class="cast-card-character">Monica</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Julien Comte</span>
    <span class="cast-card-character">Forensic Pathologist</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3040614.jpg" alt="Pierre-François Garel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Pierre-François Garel</span>
    <span class="cast-card-character">Judge Janvier</span>
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
    <li><a href="../nouvelle-vague-2025">Nouvelle Vague</a> because of Pierre-François Garel</li>
  </ul>
</section>

</article>
