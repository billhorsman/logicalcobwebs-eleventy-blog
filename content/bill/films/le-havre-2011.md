---
title: "Le Havre"
layout: layouts/films.njk
slug: le-havre-2011
ogImage: content/bill/films/backdrops/le-havre-2011.jpg
description: "In the French harbor city of Le Havre, fate throws young African refugee Idrissa into the path of Marcel Marx, a well-spoken bohemian who works as a shoe-shiner. With innate optimism and the tireless support of his community, Marcel stands up to officials pursuing the boy for deportation."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../micmacs-2009"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">50 / 100</a>
  </div>
  <div class="next">
    <a href="../tomboy-2011">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Micmacs
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Tomboy
    </span>
  </div>
</nav>

<article class="film slug-le-havre-2011">
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
  <img src="../films/profiles/20853.jpg" alt="André Wilms" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">André Wilms</span>
    <span class="cast-card-character">Marcel Marx</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5999.jpg" alt="Kati Outinen" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Kati Outinen</span>
    <span class="cast-card-character">Arletty</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/28463.jpg" alt="Jean-Pierre Darroussin" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Pierre Darroussin</span>
    <span class="cast-card-character">Monet</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/569369.jpg" alt="Blondin Miguel" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Blondin Miguel</span>
    <span class="cast-card-character">Idrissa</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/53507.jpg" alt="Elina Salo" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Elina Salo</span>
    <span class="cast-card-character">Claire</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/49756.jpg" alt="Evelyne Didi" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Evelyne Didi</span>
    <span class="cast-card-character">Yvette</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Quoc Dung Nguyen</span>
    <span class="cast-card-character">Chang</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">François Monnié</span>
    <span class="cast-card-character">Greengrocer</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Roberto Piazza</span>
    <span class="cast-card-character">Little Bob</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/548816.jpg" alt="Pierre Étaix" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Pierre Étaix</span>
    <span class="cast-card-character">Dr. Becker</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1653.jpg" alt="Jean-Pierre Léaud" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Pierre Léaud</span>
    <span class="cast-card-character">The Whistleblower</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Vincent Lebodo</span>
    <span class="cast-card-character">Francis</span>
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
    <li><a href="../day-for-night-1973">Day for Night</a> because of Jean-Pierre Léaud</li>
  </ul>
</section>

</article>
