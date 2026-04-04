---
title: "La Strada"
layout: layouts/films.njk
slug: la-strada-1954
ogImage: content/bill/films/backdrops/la-strada-1954.jpg
description: "When Gelsomina, a naïve young woman, is purchased from her impoverished mother by brutish circus strongman Zampanò to be his wife and partner, she loyally endures her husband's coldness and abuse as they travel the Italian countryside performing together. Soon Zampanò must deal with his jealousy and conflicted feelings about Gelsomina when she finds a kindred spirit in Il Matto, the carefree circus fool, and contemplates leaving Zampanò."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../its-a-wonderful-life-1946"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">2 / 100</a>
  </div>
  <div class="next">
    <a href="../rear-window-1954">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      It's a Wonderful Life
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Rear Window
    </span>
  </div>
</nav>

<article class="film slug-la-strada-1954">
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
  <img src="../films/profiles/5402.jpg" alt="Giulietta Masina" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Giulietta Masina</span>
    <span class="cast-card-character">Gelsomina</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5401.jpg" alt="Anthony Quinn" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Anthony Quinn</span>
    <span class="cast-card-character">Zampanò</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5403.jpg" alt="Richard Basehart" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Richard Basehart</span>
    <span class="cast-card-character">Il 'Matto'</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5404.jpg" alt="Aldo Silvani" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Aldo Silvani</span>
    <span class="cast-card-character">Il Signor Giraffa</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5405.jpg" alt="Marcella Rovere" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Marcella Rovere</span>
    <span class="cast-card-character">La Vedova</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Livia Venturini</span>
    <span class="cast-card-character">La Suorina</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/240913.jpg" alt="Pietro Ceccarelli" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Pietro Ceccarelli</span>
    <span class="cast-card-character">Innkeeper (uncredited)</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Giovanna Galli</span>
    <span class="cast-card-character">Prostitute at the Inn (uncredited)</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Gustavo Giorgi</span>
    <span class="cast-card-character">(uncredited)</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Yami Kamadeva</span>
    <span class="cast-card-character">Prostitute (uncredited)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5409.jpg" alt="Mario Passante" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Mario Passante</span>
    <span class="cast-card-character">Waiter (uncredited)</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Anna Primula</span>
    <span class="cast-card-character">Gelsomina's Mother (uncredited)</span>
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
    <li><a href="../la-dolce-vita-1960">La Dolce Vita</a> because of Federico Fellini</li>
  </ul>
</section>

</article>
