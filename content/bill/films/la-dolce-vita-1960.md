---
title: "La Dolce Vita"
layout: layouts/films.njk
slug: la-dolce-vita-1960
ogImage: content/bill/films/backdrops/la-dolce-vita-1960.jpg
description: "Episodic journey of journalist Marcello who struggles to find his place in the world, torn between the allure of Rome's elite social scene and the stifling domesticity offered by his girlfriend, all the while searching for a way to become a serious writer."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../breathless-1960"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">5 / 100</a>
  </div>
  <div class="next">
    <a href="../purple-noon-1960">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Breathless
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Purple Noon
    </span>
  </div>
</nav>

<article class="film slug-la-dolce-vita-1960">
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
  <img src="../films/profiles/5676.jpg" alt="Marcello Mastroianni" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Marcello Mastroianni</span>
    <span class="cast-card-character">Marcello Rubini</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5961.jpg" alt="Anita Ekberg" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Anita Ekberg</span>
    <span class="cast-card-character">Sylvia</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5682.jpg" alt="Anouk Aimée" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Anouk Aimée</span>
    <span class="cast-card-character">Maddalena</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5962.jpg" alt="Yvonne Furneaux" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Yvonne Furneaux</span>
    <span class="cast-card-character">Emma</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5963.jpg" alt="Magali Noël" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Magali Noël</span>
    <span class="cast-card-character">Fanny</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5964.jpg" alt="Alain Cuny" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Alain Cuny</span>
    <span class="cast-card-character">Steiner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5965.jpg" alt="Annibale Ninchi" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Annibale Ninchi</span>
    <span class="cast-card-character">Marcello's father</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5966.jpg" alt="Walter Santesso" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Walter Santesso</span>
    <span class="cast-card-character">Paparazzo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5967.jpg" alt="Valeria Ciangottini" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Valeria Ciangottini</span>
    <span class="cast-card-character">Paola</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5968.jpg" alt="Riccardo Garrone" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Riccardo Garrone</span>
    <span class="cast-card-character">Riccardo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5969.jpg" alt="Ida Galli" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ida Galli</span>
    <span class="cast-card-character">Debutante of the year</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1736893.jpg" alt="Audrey McDonald" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Audrey McDonald</span>
    <span class="cast-card-character">Jane</span>
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
    <li><a href="../la-strada-1954">La Strada</a> because of Federico Fellini</li>
  </ul>
</section>

</article>
