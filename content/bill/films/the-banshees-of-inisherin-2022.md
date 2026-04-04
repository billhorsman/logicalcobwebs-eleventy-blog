---
title: "The Banshees of Inisherin"
layout: layouts/films.njk
slug: the-banshees-of-inisherin-2022
ogImage: content/bill/films/backdrops/the-banshees-of-inisherin-2022.jpg
description: "Two lifelong friends find themselves at an impasse when one abruptly ends their relationship, with alarming consequences for both of them."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../one-fine-morning-2022"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">91 / 100</a>
  </div>
  <div class="next">
    <a href="../the-fabelmans-2022">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      One Fine Morning
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Fabelmans
    </span>
  </div>
</nav>

<article class="film slug-the-banshees-of-inisherin-2022">
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
  <img src="../films/profiles/72466.jpg" alt="Colin Farrell" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Colin Farrell</span>
    <span class="cast-card-character">Pádraic Súilleabháin</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2039.jpg" alt="Brendan Gleeson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Brendan Gleeson</span>
    <span class="cast-card-character">Colm Doherty</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/62105.jpg" alt="Kerry Condon" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Kerry Condon</span>
    <span class="cast-card-character">Siobhán Súilleabháin</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1290466.jpg" alt="Barry Keoghan" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Barry Keoghan</span>
    <span class="cast-card-character">Dominic Kearney</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/93209.jpg" alt="Gary Lydon" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Gary Lydon</span>
    <span class="cast-card-character">Peadar Kearney</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/83278.jpg" alt="Pat Shortt" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Pat Shortt</span>
    <span class="cast-card-character">Jonjo Devine</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1597386.jpg" alt="Sheila Flitton" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sheila Flitton</span>
    <span class="cast-card-character">Mrs. McCormick</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1907071.jpg" alt="Bríd Ní Neachtain" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Bríd Ní Neachtain</span>
    <span class="cast-card-character">Mrs. O'Riordan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/37169.jpg" alt="Jon Kenny" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jon Kenny</span>
    <span class="cast-card-character">Gerry</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/210061.jpg" alt="Aaron Monaghan" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Aaron Monaghan</span>
    <span class="cast-card-character">Declan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1122014.jpg" alt="David Pearse" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">David Pearse</span>
    <span class="cast-card-character">Priest</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">John Carty</span>
    <span class="cast-card-character">Older Musician 1</span>
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
    <li><a href="../phone-booth-2003">Phone Booth</a> because of Colin Farrell</li>
<li><a href="../in-bruges-2008">In Bruges</a> because of Colin Farrell, Brendan Gleeson and Martin McDonagh</li>
<li><a href="../the-tragedy-of-macbeth-2021">The Tragedy of Macbeth</a> because of Brendan Gleeson</li>
  </ul>
</section>

</article>
