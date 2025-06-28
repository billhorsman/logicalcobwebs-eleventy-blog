---
title: "Mr. Turner"
layout: layouts/films.njk
slug: mr-turner-2014
ogImage: content/bill/films/backdrops/mr-turner-2014.jpg
description: "Eccentric British painter J.M.W. Turner  lives his last 25 years with gusto and secretly becomes involved with a seaside landlady, while his faithful housekeeper bears an unrequited love for him."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../interstellar-2014"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">61 / 100</a>
  </div>
  <div class="next">
    <a href="../the-grand-budapest-hotel-2014">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Interstellar
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Grand Budapest Hotel
    </span>
  </div>
</nav>

<article class="film slug-mr-turner-2014">
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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../brazil-1985">Brazil</a> by Roger Ashton-Griffiths</li>
<li><a href="../the-bourne-identity-2002">The Bourne Identity</a> by Vincent Franklin</li>
<li><a href="../hot-fuzz-2007">Hot Fuzz</a> by Karl Johnson and Peter Wight</li>
<li><a href="../happygolucky-2008">Happy-Go-Lucky</a> by Sylvestra Le Touzel, Kate O'Flynn, Oliver Maltman, Karina Fernandez, Sinead Matthews and Mike Leigh</li>
<li><a href="../in-bruges-2008">In Bruges</a> by Elizabeth Berrington</li>
<li><a href="../the-party-2017">The Party</a> by Timothy Spall</li>
  </ul>

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
