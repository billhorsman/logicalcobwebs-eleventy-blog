---
title: "Three Days of the Condor"
layout: layouts/films.njk
slug: three-days-of-the-condor-1975
ogImage: content/bill/films/backdrops/three-days-of-the-condor-1975.jpg
description: "When bookish CIA researcher Joe Turner finds all his co-workers dead, he, together with a woman he has kidnapped, must work together to outwit those responsible until he determines who he can really trust."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../dog-day-afternoon-1975"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">14 / 100</a>
  </div>
  <div class="next">
    <a href="../the-man-who-fell-to-earth-1976">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Dog Day Afternoon
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Man Who Fell to Earth
    </span>
  </div>
</nav>

<article class="film slug-three-days-of-the-condor-1975">
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
  <img src="../films/profiles/4135.jpg" alt="Robert Redford" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Redford</span>
    <span class="cast-card-character">Joseph Turner</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6450.jpg" alt="Faye Dunaway" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Faye Dunaway</span>
    <span class="cast-card-character">Kathy Hale</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/19153.jpg" alt="Cliff Robertson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Cliff Robertson</span>
    <span class="cast-card-character">J. Higgins</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2201.jpg" alt="Max von Sydow" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Max von Sydow</span>
    <span class="cast-card-character">G. Joubert</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/11783.jpg" alt="John Houseman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">John Houseman</span>
    <span class="cast-card-character">Mr. Wabash</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13948.jpg" alt="Addison Powell" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Addison Powell</span>
    <span class="cast-card-character">Leonard Atwood</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/81734.jpg" alt="Walter McGinn" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Walter McGinn</span>
    <span class="cast-card-character">Sam Barber</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/44550.jpg" alt="Tina Chen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tina Chen</span>
    <span class="cast-card-character">Janice Chon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6758.jpg" alt="John Randolph Jones" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">John Randolph Jones</span>
    <span class="cast-card-character">Beefy Man</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/72659.jpg" alt="Michael Kane" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Michael Kane</span>
    <span class="cast-card-character">S.W. Wicks</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Don McHenry</span>
    <span class="cast-card-character">Dr. Ferdinand Lappe</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Jess Osuna</span>
    <span class="cast-card-character">The Major</span>
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
    <li><a href="../the-sting-1973">The Sting</a> and <a href="../all-is-lost-2013">All Is Lost</a> because of Robert Redford</li>
<li><a href="../apocalypse-now-1979">Apocalypse Now</a> because of James Keane</li>
  </ul>
</section>

</article>
