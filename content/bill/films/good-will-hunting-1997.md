---
title: "Good Will Hunting"
layout: layouts/films.njk
slug: good-will-hunting-1997
ogImage: content/bill/films/backdrops/good-will-hunting-1997.jpg
description: "Headstrong yet aimless, Will Hunting has a genius-level IQ but chooses to work as a janitor at MIT. When he secretly solves highly difficult graduate-level math problems, his talents are discovered by Professor Gerald Lambeau, who decides to help the misguided youth reach his potential. When Will is arrested for attacking a police officer, Professor Lambeau makes a deal to get leniency for him if he gets court-ordered therapy. Eventually, therapist Dr. Sean Maguire helps Will confront the demons that are holding him back."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../trainspotting-1996"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">30 / 100</a>
  </div>
  <div class="next">
    <a href="../the-big-lebowski-1998">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Trainspotting
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Big Lebowski
    </span>
  </div>
</nav>

<article class="film slug-good-will-hunting-1997">
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
  <img src="../films/profiles/1892.jpg" alt="Matt Damon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Matt Damon</span>
    <span class="cast-card-character">Will Hunting</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2157.jpg" alt="Robin Williams" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Robin Williams</span>
    <span class="cast-card-character">Sean Maguire</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/880.jpg" alt="Ben Affleck" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ben Affleck</span>
    <span class="cast-card-character">Chuckie Sullivan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1640.jpg" alt="Stellan Skarsgård" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Stellan Skarsgård</span>
    <span class="cast-card-character">Gerald Lambeau</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6613.jpg" alt="Minnie Driver" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Minnie Driver</span>
    <span class="cast-card-character">Skylar</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1893.jpg" alt="Casey Affleck" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Casey Affleck</span>
    <span class="cast-card-character">Morgan O'Mally</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6614.jpg" alt="Cole Hauser" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Cole Hauser</span>
    <span class="cast-card-character">Billy McBride</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6623.jpg" alt="Vik Sahay" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Vik Sahay</span>
    <span class="cast-card-character">MIT Student</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/6615.jpg" alt="John Mighton" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">John Mighton</span>
    <span class="cast-card-character">Tom</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Rachel Majorowski</span>
    <span class="cast-card-character">Krystyn</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Colleen McCauley</span>
    <span class="cast-card-character">Cathy</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Matt Mercier</span>
    <span class="cast-card-character">Barbershop Quartet #1</span>
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
    <li><a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> and <a href="../the-bourne-identity-2002">The Bourne Identity</a> because of Matt Damon</li>
<li><a href="../dune-2021">Dune</a> because of Stellan Skarsgård</li>
  </ul>
</section>

</article>
