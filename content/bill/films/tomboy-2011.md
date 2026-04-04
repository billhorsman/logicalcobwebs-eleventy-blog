---
title: "Tomboy"
layout: layouts/films.njk
slug: tomboy-2011
ogImage: content/bill/films/backdrops/tomboy-2011.jpg
description: "A French family moves to a new neighborhood with during the summer holidays. The story follows a 10-year-old gender non-conforming child, Laure, who experiments with their gender presentation, adopting the name Mikäel."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../le-havre-2011"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">51 / 100</a>
  </div>
  <div class="next">
    <a href="../kill-bill-the-whole-bloody-affair-2011">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Le Havre
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Kill Bill: The Whole Bloody Affair
    </span>
  </div>
</nav>

<article class="film slug-tomboy-2011">
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
  <img src="../films/profiles/240732.jpg" alt="Zoé Héran" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Zoé Héran</span>
    <span class="cast-card-character">Laure / Mickaël</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/240733.jpg" alt="Malonn Lévana" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Malonn Lévana</span>
    <span class="cast-card-character">Jeanne</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/240734.jpg" alt="Jeanne Disson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jeanne Disson</span>
    <span class="cast-card-character">Lisa</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/124072.jpg" alt="Sophie Cattani" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sophie Cattani</span>
    <span class="cast-card-character">La mère</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/64751.jpg" alt="Mathieu Demy" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Mathieu Demy</span>
    <span class="cast-card-character">Le père</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Rayan Boubekri</span>
    <span class="cast-card-character">Rayan</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Yohan Vero</span>
    <span class="cast-card-character">Vince</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1119267.jpg" alt="Noah Vero" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Noah Vero</span>
    <span class="cast-card-character">Noah</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Cheyenne Lainé</span>
    <span class="cast-card-character">Cheyenne</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/96295.jpg" alt="Christel Baras" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Christel Baras</span>
    <span class="cast-card-character">La mère de Lisa</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Valérie Roucher</span>
    <span class="cast-card-character">La mère de Rayan</span>
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
    <li><a href="../portrait-of-a-lady-on-fire-2019">Portrait of a Lady on Fire</a> because of Christel Baras and Céline Sciamma</li>
<li><a href="../petite-maman-2021">Petite Maman</a> because of Céline Sciamma</li>
  </ul>
</section>

</article>
