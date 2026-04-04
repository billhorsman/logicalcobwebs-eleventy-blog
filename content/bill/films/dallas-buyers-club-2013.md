---
title: "Dallas Buyers Club"
layout: layouts/films.njk
slug: dallas-buyers-club-2013
ogImage: content/bill/films/backdrops/dallas-buyers-club-2013.jpg
description: "Loosely based on the true-life tale of Ron Woodroof, a drug-taking, women-loving, homophobic man who in 1986 was diagnosed with HIV/AIDS and given thirty days to live."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../all-is-lost-2013"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">54 / 100</a>
  </div>
  <div class="next">
    <a href="../mr-turner-2014">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      All Is Lost
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Mr. Turner
    </span>
  </div>
</nav>

<article class="film slug-dallas-buyers-club-2013">
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
  <img src="../films/profiles/10297.jpg" alt="Matthew McConaughey" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Matthew McConaughey</span>
    <span class="cast-card-character">Ron Woodroof</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9278.jpg" alt="Jennifer Garner" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jennifer Garner</span>
    <span class="cast-card-character">Eve</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7499.jpg" alt="Jared Leto" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jared Leto</span>
    <span class="cast-card-character">Rayon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/81681.jpg" alt="Denis O'Hare" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Denis O'Hare</span>
    <span class="cast-card-character">Dr. Sevard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/18324.jpg" alt="Steve Zahn" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Steve Zahn</span>
    <span class="cast-card-character">Tucker</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/21710.jpg" alt="Michael O'Neill" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Michael O'Neill</span>
    <span class="cast-card-character">Richard Barkley</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/424.jpg" alt="Dallas Roberts" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Dallas Roberts</span>
    <span class="cast-card-character">David Wayne</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2171.jpg" alt="Griffin Dunne" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Griffin Dunne</span>
    <span class="cast-card-character">Dr. Vass</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/114000.jpg" alt="Kevin Rankin" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kevin Rankin</span>
    <span class="cast-card-character">T.J.</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/230995.jpg" alt="Donna DuPlantier" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Donna DuPlantier</span>
    <span class="cast-card-character">Nurse Frazin</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/29933.jpg" alt="Deneen Tyler" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Deneen Tyler</span>
    <span class="cast-card-character">Denise</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/129868.jpg" alt="J.D. Evermore" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">J.D. Evermore</span>
    <span class="cast-card-character">Clint</span>
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
    <li><a href="../fight-club-1999">Fight Club</a>, <a href="../phone-booth-2003">Phone Booth</a> and <a href="../house-of-gucci-2021">House of Gucci</a> because of Jared Leto</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Griffin Dunne</li>
  </ul>
</section>

</article>
