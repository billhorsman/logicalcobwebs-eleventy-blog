---
title: "Once Upon a Time in the West"
layout: layouts/films.njk
slug: once-upon-a-time-in-the-west-1968
ogImage: content/bill/films/backdrops/once-upon-a-time-in-the-west-1968.jpg
description: "As the railroad builders advance unstoppably through the Arizona desert on their way to the sea, Jill arrives in the small town of Flagstone with the intention of starting a new life."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../bullitt-1968"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">10 / 100</a>
  </div>
  <div class="next">
    <a href="../the-sting-1973">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Bullitt
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      The Sting
    </span>
  </div>
</nav>

<article class="film slug-once-upon-a-time-in-the-west-1968">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as C'era una volta il West.
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
  <img src="../films/profiles/4959.jpg" alt="Claudia Cardinale" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Claudia Cardinale</span>
    <span class="cast-card-character">Jill</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4958.jpg" alt="Henry Fonda" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Henry Fonda</span>
    <span class="cast-card-character">Frank</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4765.jpg" alt="Jason Robards" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jason Robards</span>
    <span class="cast-card-character">'Cheyenne'</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4960.jpg" alt="Charles Bronson" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Charles Bronson</span>
    <span class="cast-card-character">'Harmonica'</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4961.jpg" alt="Gabriele Ferzetti" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gabriele Ferzetti</span>
    <span class="cast-card-character">Morton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4962.jpg" alt="Paolo Stoppa" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Paolo Stoppa</span>
    <span class="cast-card-character">Sam</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4963.jpg" alt="Woody Strode" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Woody Strode</span>
    <span class="cast-card-character">Frank's Gunman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4965.jpg" alt="Jack Elam" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jack Elam</span>
    <span class="cast-card-character">Frank's Gunman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4966.jpg" alt="Keenan Wynn" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Keenan Wynn</span>
    <span class="cast-card-character">Sheriff</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4968.jpg" alt="Frank Wolff" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Frank Wolff</span>
    <span class="cast-card-character">Brett McBain</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4969.jpg" alt="Lionel Stander" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lionel Stander</span>
    <span class="cast-card-character">Innkeeper</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/72888.jpg" alt="Frank Braña" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Frank Braña</span>
    <span class="cast-card-character">Frank's Gunman (uncredited)</span>
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
    <li><a href="../magnolia-1999">Magnolia</a> because of Jason Robards</li>
  </ul>
</section>

</article>
