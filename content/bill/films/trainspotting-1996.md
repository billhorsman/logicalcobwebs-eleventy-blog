---
title: "Trainspotting"
layout: layouts/films.njk
slug: trainspotting-1996
ogImage: content/bill/films/backdrops/trainspotting-1996.jpg
description: "Hilarious but harrowing, the film charts the disintegration of the friendship between Renton, Spud, Sick Boy, Tommy and Begbie as they proceed seemingly towards a psychotic, drug-fuelled self-destruction."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../fargo-1996"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">29 / 100</a>
  </div>
  <div class="next">
    <a href="../good-will-hunting-1997">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Fargo
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Good Will Hunting
    </span>
  </div>
</nav>

<article class="film slug-trainspotting-1996">
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
  <img src="../films/profiles/3061.jpg" alt="Ewan McGregor" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ewan McGregor</span>
    <span class="cast-card-character">Renton</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1125.jpg" alt="Ewen Bremner" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Ewen Bremner</span>
    <span class="cast-card-character">Spud</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9012.jpg" alt="Jonny Lee Miller" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jonny Lee Miller</span>
    <span class="cast-card-character">Sick Boy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9013.jpg" alt="Kevin McKidd" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Kevin McKidd</span>
    <span class="cast-card-character">Tommy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/18023.jpg" alt="Robert Carlyle" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Carlyle</span>
    <span class="cast-card-character">Begbie</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9015.jpg" alt="Kelly Macdonald" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Kelly Macdonald</span>
    <span class="cast-card-character">Diane</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3064.jpg" alt="Peter Mullan" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Mullan</span>
    <span class="cast-card-character">Swanney</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2467.jpg" alt="James Cosmo" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">James Cosmo</span>
    <span class="cast-card-character">Renton's Father</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9016.jpg" alt="Eileen Nicholas" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Eileen Nicholas</span>
    <span class="cast-card-character">Renton's Mother</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1837.jpg" alt="Susan Vidler" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Susan Vidler</span>
    <span class="cast-card-character">Allison</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/9017.jpg" alt="Pauline Lynch" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Pauline Lynch</span>
    <span class="cast-card-character">Lizzy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1834.jpg" alt="Shirley Henderson" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Shirley Henderson</span>
    <span class="cast-card-character">Gail</span>
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
    <li><a href="../shallow-grave-1994">Shallow Grave</a> because of Ewan McGregor, Keith Allen, Peter Mullan, Victor Eadie, Billy Riddoch and Danny Boyle</li>
<li><a href="../black-hawk-down-2001">Black Hawk Down</a> because of Ewan McGregor and Ewen Bremner</li>
<li><a href="../24-hour-party-people-2002">24 Hour Party People</a> because of Keith Allen and Shirley Henderson</li>
<li><a href="../no-country-for-old-men-2007">No Country for Old Men</a> because of Kelly Macdonald</li>
<li><a href="../mr-turner-2014">Mr. Turner</a> because of Stuart McQuarrie</li>
  </ul>
</section>

</article>
