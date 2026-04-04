---
title: "Withnail & I"
layout: layouts/films.njk
slug: withnail--i-1987
ogImage: content/bill/films/backdrops/withnail--i-1987.jpg
description: "Two out-of-work actors -- the anxious, luckless Marwood and his acerbic, alcoholic friend, Withnail -- spend their days drifting between their squalid flat, the unemployment office and the pub. When they take a holiday \"by mistake\" at the country house of Withnail's flamboyantly gay uncle, Monty, they encounter the unpleasant side of the English countryside: tedium, terrifying locals and torrential rain."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../brazil-1985"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">23 / 100</a>
  </div>
  <div class="next">
    <a href="../delicatessen-1991">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Brazil
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Delicatessen
    </span>
  </div>
</nav>

<article class="film slug-withnail--i-1987">
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
  <img src="../films/profiles/20766.jpg" alt="Richard E. Grant" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Richard E. Grant</span>
    <span class="cast-card-character">Withnail</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/47654.jpg" alt="Paul McGann" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Paul McGann</span>
    <span class="cast-card-character">...& I</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/10983.jpg" alt="Richard Griffiths" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Richard Griffiths</span>
    <span class="cast-card-character">Monty</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/53916.jpg" alt="Ralph Brown" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ralph Brown</span>
    <span class="cast-card-character">Danny</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1821.jpg" alt="Michael Elphick" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Michael Elphick</span>
    <span class="cast-card-character">Jake</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/144863.jpg" alt="Daragh O'Malley" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Daragh O'Malley</span>
    <span class="cast-card-character">Irishman</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Michael Wardle</span>
    <span class="cast-card-character">Isaac Parkin</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1368909.jpg" alt="Una Brandon-Jones" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Una Brandon-Jones</span>
    <span class="cast-card-character">Mrs. Parkin</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Noel Johnson</span>
    <span class="cast-card-character">General</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Irene Sutcliffe</span>
    <span class="cast-card-character">Waitress</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1221507.jpg" alt="Llewellyn Rees" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Llewellyn Rees</span>
    <span class="cast-card-character">Tea Shop Proprietor</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Oates</span>
    <span class="cast-card-character">Policeman 1</span>
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
