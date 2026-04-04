---
title: "In Bruges"
layout: layouts/films.njk
slug: in-bruges-2008
ogImage: content/bill/films/backdrops/in-bruges-2008.jpg
description: "Ray and Ken, two hit men, are in Bruges, Belgium, waiting for their next mission. While they are there they have time to think and discuss their previous assignment. When the mission is revealed to Ken, it is not what he expected."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../no-country-for-old-men-2007"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">47 / 100</a>
  </div>
  <div class="next">
    <a href="../fantastic-mr-fox-2009">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      No Country for Old Men
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Fantastic Mr. Fox
    </span>
  </div>
</nav>

<article class="film slug-in-bruges-2008">
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
    <span class="cast-card-character">Ray</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2039.jpg" alt="Brendan Gleeson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Brendan Gleeson</span>
    <span class="cast-card-character">Ken</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/5469.jpg" alt="Ralph Fiennes" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ralph Fiennes</span>
    <span class="cast-card-character">Harry</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/11291.jpg" alt="Clémence Poésy" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Clémence Poésy</span>
    <span class="cast-card-character">Chloë</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/45749.jpg" alt="Thekla Reuten" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Thekla Reuten</span>
    <span class="cast-card-character">Marie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/54476.jpg" alt="Jordan Prentice" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jordan Prentice</span>
    <span class="cast-card-character">Jimmy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/42998.jpg" alt="Elizabeth Berrington" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Elizabeth Berrington</span>
    <span class="cast-card-character">Natalie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/51325.jpg" alt="Jérémie Renier" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jérémie Renier</span>
    <span class="cast-card-character">Eirik</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/211413.jpg" alt="Mark Donovan" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Mark Donovan</span>
    <span class="cast-card-character">Overweight Man</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/145299.jpg" alt="Éric Godon" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Éric Godon</span>
    <span class="cast-card-character">Yuri</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/72857.jpg" alt="Anna Madeley" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Anna Madeley</span>
    <span class="cast-card-character">Denise</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/588742.jpg" alt="Theo Stevenson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Theo Stevenson</span>
    <span class="cast-card-character">Boy in Church</span>
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
    <li><a href="../black-hawk-down-2001">Black Hawk Down</a> because of Zeljko Ivanek</li>
<li><a href="../phone-booth-2003">Phone Booth</a> because of Colin Farrell</li>
<li><a href="../the-banshees-of-inisherin-2022">The Banshees of Inisherin</a> because of Colin Farrell, Brendan Gleeson and Martin McDonagh</li>
<li><a href="../the-tragedy-of-macbeth-2021">The Tragedy of Macbeth</a> because of Brendan Gleeson</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Ralph Fiennes</li>
<li><a href="../mr-turner-2014">Mr. Turner</a> because of Elizabeth Berrington</li>
<li><a href="../belfast-2021">Belfast</a> because of Ciarán Hinds</li>
  </ul>
</section>

</article>
