---
title: "Rear Window"
layout: layouts/films.njk
slug: rear-window-1954
ogImage: content/bill/films/backdrops/rear-window-1954.jpg
description: "A wheelchair-bound photographer spies on his neighbors from his apartment window and becomes convinced one of them has committed murder."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../la-strada-1954"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">3 / 100</a>
  </div>
  <div class="next">
    <a href="../breathless-1960">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      La Strada
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Breathless
    </span>
  </div>
</nav>

<article class="film slug-rear-window-1954">
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
  <img src="../films/profiles/854.jpg" alt="James Stewart" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">James Stewart</span>
    <span class="cast-card-character">L.B. 'Jeff' Jefferies</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4070.jpg" alt="Grace Kelly" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Grace Kelly</span>
    <span class="cast-card-character">Lisa Fremont</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7683.jpg" alt="Wendell Corey" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Wendell Corey</span>
    <span class="cast-card-character">Det. Lt. Thomas J. Doyle</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7684.jpg" alt="Thelma Ritter" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Thelma Ritter</span>
    <span class="cast-card-character">Stella</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7685.jpg" alt="Raymond Burr" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Raymond Burr</span>
    <span class="cast-card-character">Lars Thorwald</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7686.jpg" alt="Judith Evelyn" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Judith Evelyn</span>
    <span class="cast-card-character">Miss Lonelyhearts</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/49906.jpg" alt="Ross Bagdasarian" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ross Bagdasarian</span>
    <span class="cast-card-character">Songwriter</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/161738.jpg" alt="Georgine Darcy" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Georgine Darcy</span>
    <span class="cast-card-character">Miss Torso</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/121038.jpg" alt="Sara Berner" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sara Berner</span>
    <span class="cast-card-character">Woman on Fire Escape</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/93622.jpg" alt="Frank Cady" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Frank Cady</span>
    <span class="cast-card-character">Man on Fire Escape</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/85846.jpg" alt="Jesslyn Fax" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jesslyn Fax</span>
    <span class="cast-card-character">Miss Hearing Aid</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Rand Harper</span>
    <span class="cast-card-character">Newlywed</span>
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
    <li><a href="../its-a-wonderful-life-1946">It's a Wonderful Life</a> because of James Stewart</li>
<li><a href="../parasite-2019">Parasite</a> because of Alfred Hitchcock</li>
  </ul>
</section>

</article>
