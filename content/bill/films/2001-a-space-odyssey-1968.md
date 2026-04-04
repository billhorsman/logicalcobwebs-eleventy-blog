---
title: "2001: A Space Odyssey"
layout: layouts/films.njk
slug: 2001-a-space-odyssey-1968
ogImage: content/bill/films/backdrops/2001-a-space-odyssey-1968.jpg
description: "Humanity finds a mysterious object buried beneath the lunar surface and sets off to find its origins with the help of HAL 9000, the world's most advanced super computer."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../in-the-heat-of-the-night-1967"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">8 / 100</a>
  </div>
  <div class="next">
    <a href="../bullitt-1968">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      In the Heat of the Night
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Bullitt
    </span>
  </div>
</nav>

<article class="film slug-2001-a-space-odyssey-1968">
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
  <img src="../films/profiles/245.jpg" alt="Keir Dullea" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Keir Dullea</span>
    <span class="cast-card-character">Dr. David Bowman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/246.jpg" alt="Gary Lockwood" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Gary Lockwood</span>
    <span class="cast-card-character">Dr. Frank Poole</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/247.jpg" alt="William Sylvester" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">William Sylvester</span>
    <span class="cast-card-character">Dr. Heywood Floyd</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/253.jpg" alt="Douglas Rain" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Douglas Rain</span>
    <span class="cast-card-character">HAL 9000 (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/248.jpg" alt="Daniel Richter" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Daniel Richter</span>
    <span class="cast-card-character">Moonwatcher</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/249.jpg" alt="Leonard Rossiter" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Leonard Rossiter</span>
    <span class="cast-card-character">Dr. Andrei Smyslov</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/250.jpg" alt="Margaret Tyzack" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Margaret Tyzack</span>
    <span class="cast-card-character">Elena</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/251.jpg" alt="Robert Beatty" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Beatty</span>
    <span class="cast-card-character">Dr. Ralph Halvorsen</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/252.jpg" alt="Sean Sullivan" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sean Sullivan</span>
    <span class="cast-card-character">Dr. Roy Michaels</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Frank W. Miller</span>
    <span class="cast-card-character">Mission Controller (voice)</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Bill Weston</span>
    <span class="cast-card-character">Astronaut</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/108277.jpg" alt="Ed Bishop" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ed Bishop</span>
    <span class="cast-card-character">Aries-1B Lunar Shuttle Captain</span>
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
