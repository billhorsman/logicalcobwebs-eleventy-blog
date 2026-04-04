---
title: "Brazil"
layout: layouts/films.njk
slug: brazil-1985
ogImage: content/bill/films/backdrops/brazil-1985.jpg
description: "Low-level bureaucrat Sam Lowry escapes the monotony of his day-to-day life through a recurring daydream of himself as a virtuous hero saving a beautiful damsel. Investigating a case that led to the wrongful arrest and eventual death of an innocent man instead of wanted terrorist Harry Tuttle, he meets the woman from his daydream, and in trying to help her gets caught in a web of mistaken identities, mindless bureaucracy and lies."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../paris-texas-1984"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">22 / 100</a>
  </div>
  <div class="next">
    <a href="../withnail--i-1987">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Paris, Texas
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Withnail & I
    </span>
  </div>
</nav>

<article class="film slug-brazil-1985">
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
  <img src="../films/profiles/378.jpg" alt="Jonathan Pryce" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jonathan Pryce</span>
    <span class="cast-card-character">Sam Lowry</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/380.jpg" alt="Robert De Niro" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Robert De Niro</span>
    <span class="cast-card-character">Harry Tuttle</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/381.jpg" alt="Katherine Helmond" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Katherine Helmond</span>
    <span class="cast-card-character">Mrs. Ida Lowry</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/65.jpg" alt="Ian Holm" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ian Holm</span>
    <span class="cast-card-character">Mr. Kurtzmann</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/382.jpg" alt="Bob Hoskins" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Bob Hoskins</span>
    <span class="cast-card-character">Spoor</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/383.jpg" alt="Michael Palin" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Michael Palin</span>
    <span class="cast-card-character">Jack Lint</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/385.jpg" alt="Ian Richardson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ian Richardson</span>
    <span class="cast-card-character">Mr. Warrenn</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/386.jpg" alt="Peter Vaughan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Vaughan</span>
    <span class="cast-card-character">Mr. Helpmann</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/387.jpg" alt="Kim Greist" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kim Greist</span>
    <span class="cast-card-character">Jill Layton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/388.jpg" alt="Jim Broadbent" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jim Broadbent</span>
    <span class="cast-card-character">Dr. Jaffe</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/97167.jpg" alt="Barbara Hicks" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Barbara Hicks</span>
    <span class="cast-card-character">Mrs. Alma Terrain</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/374.jpg" alt="Charles McKeown" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Charles McKeown</span>
    <span class="cast-card-character">Lime</span>
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
    <li><a href="../the-deer-hunter-1978">The Deer Hunter</a> because of Robert De Niro</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> because of Jim Broadbent</li>
<li><a href="../mr-turner-2014">Mr. Turner</a> because of Roger Ashton-Griffiths</li>
  </ul>
</section>

</article>
