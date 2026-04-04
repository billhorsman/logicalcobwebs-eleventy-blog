---
title: "Bullitt"
layout: layouts/films.njk
slug: bullitt-1968
ogImage: content/bill/films/backdrops/bullitt-1968.jpg
description: "Senator Walter Chalmers is aiming to take down mob boss Pete Ross with the help of testimony from the criminal's hothead brother Johnny, who is in protective custody in San Francisco under the watch of police lieutenant Frank Bullitt. When a pair of mob hitmen enter the scene, Bullitt follows their trail through a maze of complications and double-crosses. This thriller includes one of the most famous car chases ever filmed."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../2001-a-space-odyssey-1968"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">9 / 100</a>
  </div>
  <div class="next">
    <a href="../once-upon-a-time-in-the-west-1968">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      2001: A Space Odyssey
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Once Upon a Time in the West
    </span>
  </div>
</nav>

<article class="film slug-bullitt-1968">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
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
  <img src="../films/profiles/13565.jpg" alt="Steve McQueen" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Steve McQueen</span>
    <span class="cast-card-character">Lt. Frank Bullitt</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14060.jpg" alt="Robert Vaughn" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Vaughn</span>
    <span class="cast-card-character">Walter Chalmers</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14061.jpg" alt="Jacqueline Bisset" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jacqueline Bisset</span>
    <span class="cast-card-character">Cathy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14062.jpg" alt="Don Gordon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Don Gordon</span>
    <span class="cast-card-character">Lt. Delgetti</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3087.jpg" alt="Robert Duvall" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Duvall</span>
    <span class="cast-card-character">Cabbie Weissberg</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14063.jpg" alt="Simon Oakland" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Simon Oakland</span>
    <span class="cast-card-character">Captain Sam Bennett</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14064.jpg" alt="Norman Fell" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Norman Fell</span>
    <span class="cast-card-character">Captain Baker</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14065.jpg" alt="Georg Stanford Brown" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Georg Stanford Brown</span>
    <span class="cast-card-character">Dr. Willard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14066.jpg" alt="Justin Tarr" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Justin Tarr</span>
    <span class="cast-card-character">Eddy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14067.jpg" alt="Carl Reindel" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Carl Reindel</span>
    <span class="cast-card-character">Detective Stanton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14068.jpg" alt="Felice Orlandi" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Felice Orlandi</span>
    <span class="cast-card-character">Albert E. Renick</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14069.jpg" alt="Vic Tayback" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Vic Tayback</span>
    <span class="cast-card-character">Pete Ross</span>
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
    <li><a href="../day-for-night-1973">Day for Night</a> because of Jacqueline Bisset</li>
<li><a href="../apocalypse-now-1979">Apocalypse Now</a> because of Robert Duvall</li>
  </ul>
</section>

</article>
