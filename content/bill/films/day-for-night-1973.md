---
title: "Day for Night"
layout: layouts/films.njk
slug: day-for-night-1973
ogImage: content/bill/films/backdrops/day-for-night-1973.jpg
description: "A committed film director struggles to complete his movie while coping with a myriad of crises, personal and professional, among the cast and crew."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-sting-1973"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">12 / 100</a>
  </div>
  <div class="next">
    <a href="../dog-day-afternoon-1975">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Sting
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Dog Day Afternoon
    </span>
  </div>
</nav>

<article class="film slug-day-for-night-1973">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as La Nuit américaine.
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
  <img src="../films/profiles/14061.jpg" alt="Jacqueline Bisset" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jacqueline Bisset</span>
    <span class="cast-card-character">Julie Baker</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/24499.jpg" alt="Valentina Cortese" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Valentina Cortese</span>
    <span class="cast-card-character">Séverine</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3591.jpg" alt="Dani" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Dani</span>
    <span class="cast-card-character">Liliane, the Trainee Script Girl</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/18765.jpg" alt="Alexandra Stewart" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Alexandra Stewart</span>
    <span class="cast-card-character">Stacey</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/18766.jpg" alt="Jean-Pierre Aumont" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Pierre Aumont</span>
    <span class="cast-card-character">Alexandre</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/18767.jpg" alt="Jean Champion" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jean Champion</span>
    <span class="cast-card-character">Bertrand, the Producer</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1653.jpg" alt="Jean-Pierre Léaud" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Pierre Léaud</span>
    <span class="cast-card-character">Alphonse</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1650.jpg" alt="François Truffaut" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">François Truffaut</span>
    <span class="cast-card-character">Ferrand, the Director</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/18768.jpg" alt="Niké Arrighi" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Niké Arrighi</span>
    <span class="cast-card-character">Odile, the Makeup Artist</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/136761.jpg" alt="Nathalie Baye" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Nathalie Baye</span>
    <span class="cast-card-character">Joelle, the Script Girl</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/578333.jpg" alt="Maurice Seveno" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Maurice Seveno</span>
    <span class="cast-card-character">TV Reporter</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">David Markham</span>
    <span class="cast-card-character">Doctor Michael Nelson</span>
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
    <li><a href="../bullitt-1968">Bullitt</a> because of Jacqueline Bisset</li>
<li><a href="../le-havre-2011">Le Havre</a> because of Jean-Pierre Léaud</li>
<li><a href="../man-on-the-train-2002">Man on the Train</a> because of Jean-François Stévenin</li>
  </ul>
</section>

</article>
