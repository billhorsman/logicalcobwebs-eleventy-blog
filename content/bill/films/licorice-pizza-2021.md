---
title: "Licorice Pizza"
layout: layouts/films.njk
slug: licorice-pizza-2021
ogImage: content/bill/films/backdrops/licorice-pizza-2021.jpg
description: "The story of Gary Valentine and Alana Kane growing up, running around and going through the treacherous navigation of first love in the San Fernando Valley, 1973."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../house-of-gucci-2021"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">80 / 100</a>
  </div>
  <div class="next">
    <a href="../nomadland-2021">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      House of Gucci
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Nomadland
    </span>
  </div>
</nav>

<article class="film slug-licorice-pizza-2021">
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
  <img src="../films/profiles/1741367.jpg" alt="Alana Haim" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Alana Haim</span>
    <span class="cast-card-character">Alana Kane</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2764542.jpg" alt="Cooper Hoffman" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Cooper Hoffman</span>
    <span class="cast-card-character">Gary Valentine</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2228.jpg" alt="Sean Penn" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sean Penn</span>
    <span class="cast-card-character">Jack Holden</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2887.jpg" alt="Tom Waits" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Tom Waits</span>
    <span class="cast-card-character">Rex Blau</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/51329.jpg" alt="Bradley Cooper" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Bradley Cooper</span>
    <span class="cast-card-character">Jon Peters</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/227564.jpg" alt="Benny Safdie" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Benny Safdie</span>
    <span class="cast-card-character">Joel Wachs</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/61263.jpg" alt="Skyler Gisondo" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Skyler Gisondo</span>
    <span class="cast-card-character">Lance Brannigan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/139309.jpg" alt="Mary Elizabeth Ellis" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Mary Elizabeth Ellis</span>
    <span class="cast-card-character">Momma Anita</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8265.jpg" alt="John Michael Higgins" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">John Michael Higgins</span>
    <span class="cast-card-character">Jerry Frick</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4003.jpg" alt="Christine Ebersole" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Christine Ebersole</span>
    <span class="cast-card-character">Lucille Doolittle</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/538.jpg" alt="Harriet Sansom Harris" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Harriet Sansom Harris</span>
    <span class="cast-card-character">Mary Grady</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/33528.jpg" alt="Joseph Cross" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Joseph Cross</span>
    <span class="cast-card-character">Matthew</span>
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
    <li><a href="../magnolia-1999">Magnolia</a> because of John C. Reilly, Mark Flanagan, Paul Thomas Anderson and Brian Kehew</li>
<li><a href="../uncut-gems-2019">Uncut Gems</a> because of Benny Safdie</li>
<li><a href="../the-fabelmans-2022">The Fabelmans</a> because of Isabelle Kusman and Paige Locke</li>
  </ul>
</section>

</article>
