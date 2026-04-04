---
title: "Léon: The Professional"
layout: layouts/films.njk
slug: lon-the-professional-1994
ogImage: content/bill/films/backdrops/lon-the-professional-1994.jpg
description: "Léon, the top hit man in New York, has earned a rep as an effective \"cleaner\". But when his next-door neighbors are wiped out by a loose-cannon DEA agent, he becomes the unwilling custodian of 12-year-old Mathilda. Before long, Mathilda's thoughts turn to revenge, and she considers following in Léon's footsteps."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../night-on-earth-1991"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">26 / 100</a>
  </div>
  <div class="next">
    <a href="../shallow-grave-1994">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Night on Earth
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Shallow Grave
    </span>
  </div>
</nav>

<article class="film slug-lon-the-professional-1994">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Léon.
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
  <img src="../films/profiles/1003.jpg" alt="Jean Reno" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean Reno</span>
    <span class="cast-card-character">Léon Montana</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/524.jpg" alt="Natalie Portman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Natalie Portman</span>
    <span class="cast-card-character">Mathilda Lando</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/64.jpg" alt="Gary Oldman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Gary Oldman</span>
    <span class="cast-card-character">Norman Stansfield</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1004.jpg" alt="Danny Aiello" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Danny Aiello</span>
    <span class="cast-card-character">Tony</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1005.jpg" alt="Peter Appel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Appel</span>
    <span class="cast-card-character">Malky</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1010.jpg" alt="Michael Badalucco" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Michael Badalucco</span>
    <span class="cast-card-character">Mathilda's Father</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13420.jpg" alt="Ellen Greene" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ellen Greene</span>
    <span class="cast-card-character">Mathilda's Mother</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/122230.jpg" alt="Elizabeth Regen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Elizabeth Regen</span>
    <span class="cast-card-character">Mathilda's Sister</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/977543.jpg" alt="Carl J. Matusovich" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Carl J. Matusovich</span>
    <span class="cast-card-character">Mathilda's Brother</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1171233.jpg" alt="Eric Challier" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Eric Challier</span>
    <span class="cast-card-character">Bodyguard Chief</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1006.jpg" alt="Willi One Blood" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Willi One Blood</span>
    <span class="cast-card-character">1st Stansfield Man</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1007.jpg" alt="Don Creech" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Don Creech</span>
    <span class="cast-card-character">2nd Stansfield man</span>
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
    <li><a href="../sink-or-swim-2018">Sink or Swim</a> because of Jean-Hugues Anglade</li>
  </ul>
</section>

</article>
