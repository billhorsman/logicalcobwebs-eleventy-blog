---
title: "Breathless"
layout: layouts/films.njk
slug: breathless-1960
ogImage: content/bill/films/backdrops/breathless-1960.jpg
description: "A small-time thief steals a car and impulsively murders a motorcycle policeman. Wanted by the authorities, he attempts to persuade a girl to run away to Italy with him."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../rear-window-1954"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">4 / 100</a>
  </div>
  <div class="next">
    <a href="../la-dolce-vita-1960">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Rear Window
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      La Dolce Vita
    </span>
  </div>
</nav>

<article class="film slug-breathless-1960">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as À bout de souffle.
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
  <img src="../films/profiles/3829.jpg" alt="Jean-Paul Belmondo" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Paul Belmondo</span>
    <span class="cast-card-character">Michel Poiccard / László Kovács</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3830.jpg" alt="Jean Seberg" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean Seberg</span>
    <span class="cast-card-character">Patricia Franchini</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3573.jpg" alt="Daniel Boulanger" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Daniel Boulanger</span>
    <span class="cast-card-character">Police Inspector Vital</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3832.jpg" alt="Henri-Jacques Huet" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Henri-Jacques Huet</span>
    <span class="cast-card-character">Antonio Berrutti</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3836.jpg" alt="Roger Hanin" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Roger Hanin</span>
    <span class="cast-card-character">Carl Zubart</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3833.jpg" alt="Van Doude" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Van Doude</span>
    <span class="cast-card-character">American Journalist, Patricia's Friend</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3834.jpg" alt="Claude Mansard" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Claude Mansard</span>
    <span class="cast-card-character">Claudius Mansard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/331806.jpg" alt="Liliane Dreyfus" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Liliane Dreyfus</span>
    <span class="cast-card-character">Liliane / Minouche</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Michel Fabre</span>
    <span class="cast-card-character">Police Inspector #2</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3831.jpg" alt="Jean-Pierre Melville" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Pierre Melville</span>
    <span class="cast-card-character">Parvulesco the Writer</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3776.jpg" alt="Jean-Luc Godard" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Luc Godard</span>
    <span class="cast-card-character">The Snitch</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3835.jpg" alt="Richard Balducci" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Richard Balducci</span>
    <span class="cast-card-character">Tolmatchoff</span>
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
