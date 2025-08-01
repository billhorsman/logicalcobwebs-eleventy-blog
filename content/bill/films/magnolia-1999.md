---
title: "Magnolia"
layout: layouts/films.njk
slug: magnolia-1999
ogImage: content/bill/films/backdrops/magnolia-1999.jpg
description: "On one random day in the San Fernando Valley, a dying father, a young wife, a male caretaker, a famous lost son, a police officer in love, a boy genius, an ex-boy genius, a game show host and an estranged daughter will each become part of a dazzling multiplicity of plots, but one story."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../ghost-dog-the-way-of-the-samurai-1999"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">36 / 100</a>
  </div>
  <div class="next">
    <a href="../the-straight-story-1999">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Ghost Dog: The Way of the Samurai
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Straight Story
    </span>
  </div>
</nav>

<article class="film slug-magnolia-1999">
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
    <li><a href="../once-upon-a-time-in-the-west-1968">Once Upon a Time in the West</a> because of Jason Robards</li>
<li><a href="../fargo-1996">Fargo</a> because of William H. Macy</li>
<li><a href="../the-big-lebowski-1998">The Big Lebowski</a> because of Julianne Moore, Philip Seymour Hoffman and Aimee Mann</li>
<li><a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> because of Philip Seymour Hoffman and Philip Baker Hall</li>
<li><a href="../fight-club-1999">Fight Club</a> because of Ezra Buzzington, Michael Shamus Wiles, Greg Bronson and Phil Hawn</li>
<li><a href="../the-fabelmans-2022">The Fabelmans</a> because of Ezra Buzzington</li>
<li><a href="../licorice-pizza-2021">Licorice Pizza</a> because of John C. Reilly, Mark Flanagan, Paul Thomas Anderson and Brian Kehew</li>
<li><a href="../kill-bill-the-whole-bloody-affair-2011">Kill Bill: The Whole Bloody Affair</a> because of Michael Bowen</li>
<li><a href="../killers-of-the-flower-moon-2023">Killers of the Flower Moon</a> because of Pat Healy</li>
  </ul>
</section>

</article>
