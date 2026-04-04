---
title: "Nouvelle Vague"
layout: layouts/films.njk
slug: nouvelle-vague-2025
ogImage: content/bill/films/backdrops/nouvelle-vague-2025.jpg
description: "After writing for Cahiers du cinéma, a young Jean-Luc Godard decides making films is the best film criticism. He convinces producer Georges de Beauregard to fund a low-budget feature, and creates a treatment with fellow New Wave filmmaker François Truffaut about a gangster couple. The result? Breathless, one of the first features of the Nouvelle Vague era of French cinema."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-ballad-of-wallis-island-2025"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">99 / 100</a>
  </div>
  <div class="next">
    <a href="../the-secret-agent-2025">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Ballad of Wallis Island
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Secret Agent
    </span>
  </div>
</nav>

<article class="film slug-nouvelle-vague-2025">
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
  <img src="../films/profiles/3624024.jpg" alt="Guillaume Marbeck" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Guillaume Marbeck</span>
    <span class="cast-card-character">Jean-Luc Godard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1059597.jpg" alt="Zoey Deutch" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Zoey Deutch</span>
    <span class="cast-card-character">Jean Seberg</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4620704.jpg" alt="Aubry Dullin" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Aubry Dullin</span>
    <span class="cast-card-character">Jean-Paul Belmondo</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4782604.jpg" alt="Adrien Rouyard" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Adrien Rouyard</span>
    <span class="cast-card-character">François Truffaut</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2623513.jpg" alt="Antoine Besson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Antoine Besson</span>
    <span class="cast-card-character">Claude Chabrol</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4620691.jpg" alt="Jodie Ruth-Forest" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jodie Ruth-Forest</span>
    <span class="cast-card-character">Suzanne Schiffman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2250217.jpg" alt="Bruno Dreyfürst" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Bruno Dreyfürst</span>
    <span class="cast-card-character">Georges de Beauregard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4195610.jpg" alt="Benjamin Cléry" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Benjamin Cléry</span>
    <span class="cast-card-character">Pierre Rissient</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2408218.jpg" alt="Matthieu Penchinat" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Matthieu Penchinat</span>
    <span class="cast-card-character">Raoul Coutard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3402505.jpg" alt="Pauline Belle" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Pauline Belle</span>
    <span class="cast-card-character">Suzon Faye</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2158224.jpg" alt="Frank Cicurel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Frank Cicurel</span>
    <span class="cast-card-character">Raymond Cauchetier</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5227510.jpg" alt="Blaise Pettebone" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Blaise Pettebone</span>
    <span class="cast-card-character">Marc Pierret</span>
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
    <li><a href="../anatomy-of-a-fall-2023">Anatomy of a Fall</a> because of Pierre-François Garel</li>
  </ul>
</section>

</article>
