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
    <a class="simple" href="../">32 / 100</a>
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
