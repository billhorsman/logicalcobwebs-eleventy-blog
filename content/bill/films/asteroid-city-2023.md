---
title: "Asteroid City"
layout: layouts/films.njk
slug: asteroid-city-2023
ogImage: content/bill/films/backdrops/asteroid-city-2023.jpg
description: "In an American desert town circa 1955, the itinerary of a Junior Stargazer/Space Cadet convention is spectacularly disrupted by world-changing events."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../perfect-days-2023"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">96 / 100</a>
  </div>
  <div class="next">
    <a href="../blue-jean-2023">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Perfect Days
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Blue Jean
    </span>
  </div>
</nav>

<article class="film slug-asteroid-city-2023">
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
  <img src="../films/profiles/17881.jpg" alt="Jason Schwartzman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jason Schwartzman</span>
    <span class="cast-card-character">Augie Steenbeck</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1245.jpg" alt="Scarlett Johansson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Scarlett Johansson</span>
    <span class="cast-card-character">Midge Campbell</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/31.jpg" alt="Tom Hanks" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tom Hanks</span>
    <span class="cast-card-character">Stanley Zak</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2954.jpg" alt="Jeffrey Wright" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jeffrey Wright</span>
    <span class="cast-card-character">General Gibson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3063.jpg" alt="Tilda Swinton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tilda Swinton</span>
    <span class="cast-card-character">Dr. Hickenlooper</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17419.jpg" alt="Bryan Cranston" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Bryan Cranston</span>
    <span class="cast-card-character">The Host</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/819.jpg" alt="Edward Norton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Edward Norton</span>
    <span class="cast-card-character">Conrad Earp</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3490.jpg" alt="Adrien Brody" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Adrien Brody</span>
    <span class="cast-card-character">Schubert Green</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/23626.jpg" alt="Liev Schreiber" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Liev Schreiber</span>
    <span class="cast-card-character">J.J. Kellogg</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/15250.jpg" alt="Hope Davis" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Hope Davis</span>
    <span class="cast-card-character">Sandy Borden</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4025.jpg" alt="Steve Park" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Steve Park</span>
    <span class="cast-card-character">Roger Cho</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/36669.jpg" alt="Rupert Friend" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rupert Friend</span>
    <span class="cast-card-character">Montana</span>
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
    <li><a href="../fargo-1996">Fargo</a> because of Steve Park</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Steve Park, Edward Norton, Jason Schwartzman, Willem Dafoe, Wes Anderson, Jarvis Cocker, Adrien Brody, Tilda Swinton, Tony Revolori, Bob Balaban, Fisher Stevens, Jeffrey Wright, Mohamed Belhadjine, Nicolas Avinée, Rupert Friend, Tom Hudson, Stéphane Bak, Liev Schreiber, Damien Bonnard, Rodolphe Pauly and Eliel Ford</li>
<li><a href="../fight-club-1999">Fight Club</a> because of Edward Norton</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Edward Norton, Jason Schwartzman, Willem Dafoe, Wes Anderson, Adrien Brody, Jeff Goldblum, Tilda Swinton, Tony Revolori, Bob Balaban and Fisher Stevens</li>
<li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> because of Jason Schwartzman, Willem Dafoe, Wes Anderson, Jarvis Cocker and Adrien Brody</li>
<li><a href="../the-lighthouse-2019">The Lighthouse</a> because of Willem Dafoe</li>
<li><a href="../uncut-gems-2019">Uncut Gems</a> because of Tilda Swinton and Jake Ryan</li>
  </ul>
</section>

</article>
