---
title: "Man on the Train"
layout: layouts/films.njk
slug: man-on-the-train-2002
ogImage: content/bill/films/backdrops/man-on-the-train-2002.jpg
description: "A man, Milan steps off a train, into a small French village. As he waits for the day when he will rob the town bank, he runs into an old retired poetry teacher named M. Manesquier. The two men strike up a strange friendship and explore the road not taken, each wanting to live the other's life."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../24-hour-party-people-2002"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">41 / 100</a>
  </div>
  <div class="next">
    <a href="../the-bourne-identity-2002">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      24 Hour Party People
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Bourne Identity
    </span>
  </div>
</nav>

<article class="film slug-man-on-the-train-2002">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as L'Homme du train.
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
  <img src="../films/profiles/24421.jpg" alt="Jean Rochefort" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean Rochefort</span>
    <span class="cast-card-character">Monsieur Manesquier</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/35084.jpg" alt="Johnny Hallyday" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Johnny Hallyday</span>
    <span class="cast-card-character">Milan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6304.jpg" alt="Jean-François Stévenin" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-François Stévenin</span>
    <span class="cast-card-character">Luigi</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/91271.jpg" alt="Pascal Parmentier" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Pascal Parmentier</span>
    <span class="cast-card-character">Sadko</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1543580.jpg" alt="Charlie Nelson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Charlie Nelson</span>
    <span class="cast-card-character">Max</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/91272.jpg" alt="Isabelle Petit-Jacques" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Isabelle Petit-Jacques</span>
    <span class="cast-card-character">Viviane, Manesquier's mistress</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/27980.jpg" alt="Édith Scob" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Édith Scob</span>
    <span class="cast-card-character">Manesquier's Sister</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/68722.jpg" alt="Maurice Chevit" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Maurice Chevit</span>
    <span class="cast-card-character">Hairdresser</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/35882.jpg" alt="Riton Liebman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Riton Liebman</span>
    <span class="cast-card-character">Burly Guy</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Olivier Fauron</span>
    <span class="cast-card-character">Schoolboy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1291135.jpg" alt="Véronique Kapoyan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Véronique Kapoyan</span>
    <span class="cast-card-character">Baker</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Armand Chagot</span>
    <span class="cast-card-character">Gardener of Manesquier</span>
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
    <li><a href="../day-for-night-1973">Day for Night</a> because of Jean-François Stévenin</li>
  </ul>
</section>

</article>
