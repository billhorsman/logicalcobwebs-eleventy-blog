---
title: "The Grand Budapest Hotel"
layout: layouts/films.njk
slug: the-grand-budapest-hotel-2014
ogImage: content/bill/films/backdrops/the-grand-budapest-hotel-2014.jpg
description: "The Grand Budapest Hotel tells of a legendary concierge at a famous European hotel between the wars and his friendship with a young employee who becomes his trusted protégé. The story involves the theft and recovery of a priceless Renaissance painting, the battle for an enormous family fortune and the slow and then sudden upheavals that transformed Europe during the first half of the 20th century."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../mr-turner-2014"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">56 / 100</a>
  </div>
  <div class="next">
    <a href="../taxi-2015">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Mr. Turner
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Taxi
    </span>
  </div>
</nav>

<article class="film slug-the-grand-budapest-hotel-2014">
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
  <img src="../films/profiles/5469.jpg" alt="Ralph Fiennes" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ralph Fiennes</span>
    <span class="cast-card-character">M. Gustave</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1164.jpg" alt="F. Murray Abraham" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">F. Murray Abraham</span>
    <span class="cast-card-character">Mr. Moustafa</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8789.jpg" alt="Mathieu Amalric" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Mathieu Amalric</span>
    <span class="cast-card-character">Serge X.</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3490.jpg" alt="Adrien Brody" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Adrien Brody</span>
    <span class="cast-card-character">Dmitri</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5293.jpg" alt="Willem Dafoe" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Willem Dafoe</span>
    <span class="cast-card-character">Jopling</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4785.jpg" alt="Jeff Goldblum" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jeff Goldblum</span>
    <span class="cast-card-character">Deputy Kovacs</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1037.jpg" alt="Harvey Keitel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Harvey Keitel</span>
    <span class="cast-card-character">Ludwig</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9642.jpg" alt="Jude Law" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jude Law</span>
    <span class="cast-card-character">Young Author</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1532.jpg" alt="Bill Murray" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Bill Murray</span>
    <span class="cast-card-character">M. Ivan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/819.jpg" alt="Edward Norton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Edward Norton</span>
    <span class="cast-card-character">Henckels</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/36592.jpg" alt="Saoirse Ronan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Saoirse Ronan</span>
    <span class="cast-card-character">Agatha</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17881.jpg" alt="Jason Schwartzman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jason Schwartzman</span>
    <span class="cast-card-character">M. Jean</span>
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
    <li><a href="../fight-club-1999">Fight Club</a> because of Edward Norton</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Edward Norton, Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Wes Anderson, Owen Wilson, Adrien Brody, Mathieu Amalric, Saoirse Ronan, Léa Seydoux, Tilda Swinton, Tony Revolori, Larry Pine, Bob Balaban and Fisher Stevens</li>
<li><a href="../asteroid-city-2023">Asteroid City</a> because of Edward Norton, Jason Schwartzman, Willem Dafoe, Wes Anderson, Adrien Brody, Jeff Goldblum, Tilda Swinton, Tony Revolori, Bob Balaban and Fisher Stevens</li>
<li><a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> because of Jude Law</li>
<li><a href="../in-bruges-2008">In Bruges</a> because of Ralph Fiennes</li>
<li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> because of Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Robin Hurlstone, Wes Anderson, Owen Wilson and Adrien Brody</li>
<li><a href="../the-lighthouse-2019">The Lighthouse</a> because of Willem Dafoe</li>
<li><a href="../sink-or-swim-2018">Sink or Swim</a> because of Mathieu Amalric</li>
<li><a href="../lady-bird-2017">Lady Bird</a> because of Saoirse Ronan and Lucas Hedges</li>
<li><a href="../little-women-2019">Little Women</a> because of Saoirse Ronan</li>
<li><a href="../one-fine-morning-2022">One Fine Morning</a> because of Léa Seydoux</li>
<li><a href="../uncut-gems-2019">Uncut Gems</a> because of Tilda Swinton</li>
  </ul>
</section>

</article>
