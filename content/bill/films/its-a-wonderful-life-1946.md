---
title: "It's a Wonderful Life"
layout: layouts/films.njk
slug: its-a-wonderful-life-1946
ogImage: content/bill/films/backdrops/its-a-wonderful-life-1946.jpg
description: "A holiday favourite for generations...  George Bailey has spent his entire life giving to the people of Bedford Falls.  All that prevents rich skinflint Mr. Potter from taking over the entire town is George's modest building and loan company.  But on Christmas Eve the business's $8,000 is lost and George's troubles begin."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <span><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</span>
  </div>
  <div>
    <a class="simple" href="../">1 / 100</a>
  </div>
  <div class="next">
    <a href="../la-strada-1954">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Start of list
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      La Strada
    </span>
  </div>
</nav>

<article class="film slug-its-a-wonderful-life-1946">
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
    <span class="cast-card-character">George Bailey</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17752.jpg" alt="Donna Reed" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Donna Reed</span>
    <span class="cast-card-character">Mary Hatch</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17753.jpg" alt="Lionel Barrymore" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Lionel Barrymore</span>
    <span class="cast-card-character">Mr. Potter</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3383.jpg" alt="Thomas Mitchell" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Thomas Mitchell</span>
    <span class="cast-card-character">Uncle Billy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7666.jpg" alt="Henry Travers" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Henry Travers</span>
    <span class="cast-card-character">Clarence</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17755.jpg" alt="Beulah Bondi" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Beulah Bondi</span>
    <span class="cast-card-character">Mrs. Bailey</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17759.jpg" alt="Frank Faylen" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Frank Faylen</span>
    <span class="cast-card-character">Ernie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4303.jpg" alt="Ward Bond" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ward Bond</span>
    <span class="cast-card-character">Bert</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/77081.jpg" alt="Gloria Grahame" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Gloria Grahame</span>
    <span class="cast-card-character">Violet</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/33278.jpg" alt="H.B. Warner" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">H.B. Warner</span>
    <span class="cast-card-character">Mr. Gower</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/78902.jpg" alt="Frank Albertson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Frank Albertson</span>
    <span class="cast-card-character">Sam Wainwright</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Todd Karns</span>
    <span class="cast-card-character">Harry Bailey</span>
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
    <li><a href="../rear-window-1954">Rear Window</a> because of James Stewart</li>
  </ul>
</section>

</article>
