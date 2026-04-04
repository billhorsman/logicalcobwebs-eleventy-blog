---
title: "Black Hawk Down"
layout: layouts/films.njk
slug: black-hawk-down-2001
ogImage: content/bill/films/backdrops/black-hawk-down-2001.jpg
description: "When U.S. Rangers and an elite Delta Force team attempt to kidnap two underlings of a Somali warlord, their Black Hawk helicopters are shot down, and the Americans suffer heavy casualties, facing intense fighting from the militia on the ground."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../amlie-2001"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">39 / 100</a>
  </div>
  <div class="next">
    <a href="../24-hour-party-people-2002">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Amélie
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      24 Hour Party People
    </span>
  </div>
</nav>

<article class="film slug-black-hawk-down-2001">
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
  <img src="../films/profiles/2299.jpg" alt="Josh Hartnett" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Josh Hartnett</span>
    <span class="cast-card-character">SSG Matthew Eversmann</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8783.jpg" alt="Eric Bana" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Eric Bana</span>
    <span class="cast-card-character">SFC Norm 'Hoot' Gibson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3061.jpg" alt="Ewan McGregor" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ewan McGregor</span>
    <span class="cast-card-character">SPC John Grimes</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3197.jpg" alt="Tom Sizemore" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Tom Sizemore</span>
    <span class="cast-card-character">LTC Danny McKnight</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/886.jpg" alt="William Fichtner" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">William Fichtner</span>
    <span class="cast-card-character">SFC Jeff Sanderson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9880.jpg" alt="Sam Shepard" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sam Shepard</span>
    <span class="cast-card-character">MG William F. Garrison</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/11355.jpg" alt="Jason Isaacs" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jason Isaacs</span>
    <span class="cast-card-character">CPT Mike Steele</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1125.jpg" alt="Ewen Bremner" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ewen Bremner</span>
    <span class="cast-card-character">SPC Shawn Nelson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/114.jpg" alt="Orlando Bloom" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Orlando Bloom</span>
    <span class="cast-card-character">PFC Todd Blackburn</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2524.jpg" alt="Tom Hardy" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Tom Hardy</span>
    <span class="cast-card-character">SPC Lance Twombly</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/12790.jpg" alt="Charlie Hofheimer" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Charlie Hofheimer</span>
    <span class="cast-card-character">CPL James 'Jamie' Smith</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/12791.jpg" alt="Hugh Dancy" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Hugh Dancy</span>
    <span class="cast-card-character">SFC Kurt 'Doc' Schmid</span>
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
    <li><a href="../blade-runner-1982">Blade Runner</a> and <a href="../house-of-gucci-2021">House of Gucci</a> because of Ridley Scott</li>
<li><a href="../delicatessen-1991">Delicatessen</a> because of Nikky Smedley</li>
<li><a href="../shallow-grave-1994">Shallow Grave</a> because of Ewan McGregor</li>
<li><a href="../trainspotting-1996">Trainspotting</a> because of Ewan McGregor and Ewen Bremner</li>
<li><a href="../in-bruges-2008">In Bruges</a> because of Zeljko Ivanek</li>
  </ul>
</section>

</article>
