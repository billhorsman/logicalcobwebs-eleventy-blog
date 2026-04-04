---
title: "The Sting"
layout: layouts/films.njk
slug: the-sting-1973
ogImage: content/bill/films/backdrops/the-sting-1973.jpg
description: "A novice con man teams up with an acknowledged master to avenge the murder of a mutual friend by pulling off the ultimate big con and swindling a fortune from a big-time mobster."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../once-upon-a-time-in-the-west-1968"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">11 / 100</a>
  </div>
  <div class="next">
    <a href="../day-for-night-1973">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Once Upon a Time in the West
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Day for Night
    </span>
  </div>
</nav>

<article class="film slug-the-sting-1973">
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
  <img src="../films/profiles/3636.jpg" alt="Paul Newman" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Paul Newman</span>
    <span class="cast-card-character">Henry Gondorff</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4135.jpg" alt="Robert Redford" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Redford</span>
    <span class="cast-card-character">Johnny Hooker</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8606.jpg" alt="Robert Shaw" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Shaw</span>
    <span class="cast-card-character">Doyle Lonnegan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1466.jpg" alt="Charles Durning" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Charles Durning</span>
    <span class="cast-card-character">Lt. Wm. Snyder</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/4093.jpg" alt="Ray Walston" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ray Walston</span>
    <span class="cast-card-character">J.J. Singleton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/39015.jpg" alt="Eileen Brennan" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Eileen Brennan</span>
    <span class="cast-card-character">Billie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/14833.jpg" alt="Harold Gould" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Harold Gould</span>
    <span class="cast-card-character">Kid Twist</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/171056.jpg" alt="John Heffernan" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">John Heffernan</span>
    <span class="cast-card-character">Eddie Niles</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/37712.jpg" alt="Dana Elcar" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Dana Elcar</span>
    <span class="cast-card-character">F.B.I. Agent Polk</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1273.jpg" alt="Jack Kehoe" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jack Kehoe</span>
    <span class="cast-card-character">Erie Kid</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1220923.jpg" alt="Dimitra Arliss" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Dimitra Arliss</span>
    <span class="cast-card-character">Loretta</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/98927.jpg" alt="Robert Earl Jones" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Earl Jones</span>
    <span class="cast-card-character">Luther Coleman</span>
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
    <li><a href="../in-the-heat-of-the-night-1967">In the Heat of the Night</a> because of Larry D. Mann</li>
<li><a href="../three-days-of-the-condor-1975">Three Days of the Condor</a> and <a href="../all-is-lost-2013">All Is Lost</a> because of Robert Redford</li>
<li><a href="../dog-day-afternoon-1975">Dog Day Afternoon</a> because of Charles Durning</li>
  </ul>
</section>

</article>
