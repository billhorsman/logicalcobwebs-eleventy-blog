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
    <a href="../happygolucky-2008"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">51 / 100</a>
  </div>
  <div class="next">
    <a href="../district-9-2009">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Happy-Go-Lucky
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      District 9
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
