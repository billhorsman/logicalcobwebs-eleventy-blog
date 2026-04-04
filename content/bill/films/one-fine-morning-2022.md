---
title: "One Fine Morning"
layout: layouts/films.njk
slug: one-fine-morning-2022
ogImage: content/bill/films/backdrops/one-fine-morning-2022.jpg
description: "With a father suffering from neurodegenerative disease, a young woman lives with her eight-year-old daughter. While struggling to secure a decent nursing home, she runs into an unavailable friend with whom she embarks on an affair."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../empire-of-light-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">90 / 100</a>
  </div>
  <div class="next">
    <a href="../the-banshees-of-inisherin-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Empire of Light
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Banshees of Inisherin
    </span>
  </div>
</nav>

<article class="film slug-one-fine-morning-2022">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Un beau matin.
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
  <img src="../films/profiles/121529.jpg" alt="Léa Seydoux" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Léa Seydoux</span>
    <span class="cast-card-character">Sandra Kienzler</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/16923.jpg" alt="Pascal Greggory" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Pascal Greggory</span>
    <span class="cast-card-character">Georg Kienzler</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/26100.jpg" alt="Melvil Poupaud" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Melvil Poupaud</span>
    <span class="cast-card-character">Clément</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/24485.jpg" alt="Nicole Garcia" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Nicole Garcia</span>
    <span class="cast-card-character">Françoise</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3549778.jpg" alt="Camille Leban Martins" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Camille Leban Martins</span>
    <span class="cast-card-character">Linn</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1096268.jpg" alt="Sarah Le Picard" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sarah Le Picard</span>
    <span class="cast-card-character">Elodie Kienzler</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2814135.jpg" alt="Pierre Meunier" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Pierre Meunier</span>
    <span class="cast-card-character">Michel</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1770282.jpg" alt="Fejria Deliba" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Fejria Deliba</span>
    <span class="cast-card-character">Leila</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Jacqueline Hansen-Løve</span>
    <span class="cast-card-character">Jacqueline Kienzler</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/59045.jpg" alt="Catherine Vinatier" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Catherine Vinatier</span>
    <span class="cast-card-character">Soeur de Georg</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1652770.jpg" alt="Samuel Achache" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Samuel Achache</span>
    <span class="cast-card-character">Mari d'Elodie</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Esther Wajeman</span>
    <span class="cast-card-character">Enfant d'Elodie</span>
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
    <li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Léa Seydoux</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Léa Seydoux and Sharif Andoura</li>
  </ul>
</section>

</article>
