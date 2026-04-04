---
title: "In the Heat of the Night"
layout: layouts/films.njk
slug: in-the-heat-of-the-night-1967
ogImage: content/bill/films/backdrops/in-the-heat-of-the-night-1967.jpg
description: "African-American Philadelphia police detective Virgil Tibbs is arrested on suspicion of murder by Bill Gillespie, the racist police chief of tiny Sparta, Mississippi. After Tibbs proves not only his own innocence but that of another man, he joins forces with Gillespie to track down the real killer. Their investigation takes them through every social level of the town, with Tibbs making enemies as well as unlikely friends as he hunts for the truth."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../purple-noon-1960"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">7 / 100</a>
  </div>
  <div class="next">
    <a href="../2001-a-space-odyssey-1968">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Purple Noon
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      2001: A Space Odyssey
    </span>
  </div>
</nav>

<article class="film slug-in-the-heat-of-the-night-1967">
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
  <img src="../films/profiles/16897.jpg" alt="Sidney Poitier" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sidney Poitier</span>
    <span class="cast-card-character">Virgil Tibbs</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/522.jpg" alt="Rod Steiger" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Rod Steiger</span>
    <span class="cast-card-character">Police Chief Bill Gillespie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8255.jpg" alt="Warren Oates" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Warren Oates</span>
    <span class="cast-card-character">Deputy Sam Wood</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/89527.jpg" alt="Peter Whitney" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Whitney</span>
    <span class="cast-card-character">Deputy Courtney</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/30123.jpg" alt="Lee Grant" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Lee Grant</span>
    <span class="cast-card-character">Mrs. Leslie Colbert</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3715.jpg" alt="Anthony James" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Anthony James</span>
    <span class="cast-card-character">Ralph</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/15992.jpg" alt="William Schallert" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">William Schallert</span>
    <span class="cast-card-character">Mayor Schubert</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6914.jpg" alt="Scott Wilson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Scott Wilson</span>
    <span class="cast-card-character">Harvey Oberst</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3640.jpg" alt="Larry Gates" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Larry Gates</span>
    <span class="cast-card-character">Eric Endicott</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/105049.jpg" alt="James Patterson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">James Patterson</span>
    <span class="cast-card-character">Mr. Purdy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/196406.jpg" alt="Quentin Dean" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Quentin Dean</span>
    <span class="cast-card-character">Delores Purdy</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Kermit Murdock</span>
    <span class="cast-card-character">Henderson</span>
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
    <li><a href="../the-sting-1973">The Sting</a> because of Larry D. Mann</li>
  </ul>
</section>

</article>
