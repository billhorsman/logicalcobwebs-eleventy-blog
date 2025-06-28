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
    <a href="../all-of-us-strangers-2023"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">97 / 100</a>
  </div>
  <div class="next">
    <a href="../blue-jean-2023">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      All of Us Strangers
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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../fargo-1996">Fargo</a> by Steve Park</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> by Steve Park, Edward Norton, Jason Schwartzman, Willem Dafoe, Wes Anderson, Jarvis Cocker, Adrien Brody, Tilda Swinton, Tony Revolori, Bob Balaban, Fisher Stevens, Jeffrey Wright, Mohamed Belhadjine, Nicolas Avinée, Rupert Friend, Tom Hudson, Stéphane Bak, Liev Schreiber, Damien Bonnard, Rodolphe Pauly and Eliel Ford</li>
<li><a href="../fight-club-1999">Fight Club</a> by Edward Norton</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> by Edward Norton, Jason Schwartzman, Willem Dafoe, Wes Anderson, Adrien Brody, Jeff Goldblum, Tilda Swinton, Tony Revolori, Bob Balaban and Fisher Stevens</li>
<li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> by Jason Schwartzman, Willem Dafoe, Wes Anderson, Jarvis Cocker and Adrien Brody</li>
<li><a href="../the-lighthouse-2019">The Lighthouse</a> by Willem Dafoe</li>
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
