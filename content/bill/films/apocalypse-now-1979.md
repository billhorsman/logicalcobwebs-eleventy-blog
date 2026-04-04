---
title: "Apocalypse Now"
layout: layouts/films.njk
slug: apocalypse-now-1979
ogImage: content/bill/films/backdrops/apocalypse-now-1979.jpg
description: "At the height of the Vietnam war, Captain Benjamin Willard is sent on a dangerous mission that, officially, \"does not exist, nor will it ever exist.\" His goal is to locate - and eliminate - a mysterious Green Beret Colonel named Walter Kurtz, who has been leading his personal army on illegal guerrilla missions into enemy territory."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-deer-hunter-1978"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">17 / 100</a>
  </div>
  <div class="next">
    <a href="../diva-1981">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Deer Hunter
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Diva
    </span>
  </div>
</nav>

<article class="film slug-apocalypse-now-1979">
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
  <img src="../films/profiles/8349.jpg" alt="Martin Sheen" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Martin Sheen</span>
    <span class="cast-card-character">Captain Benjamin Willard</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3084.jpg" alt="Marlon Brando" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Marlon Brando</span>
    <span class="cast-card-character">Colonel Walter Kurtz</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8354.jpg" alt="Albert Hall" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Albert Hall</span>
    <span class="cast-card-character">Chief Phillips</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8351.jpg" alt="Frederic Forrest" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Frederic Forrest</span>
    <span class="cast-card-character">Jay 'Chef' Hicks</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2975.jpg" alt="Laurence Fishburne" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Laurence Fishburne</span>
    <span class="cast-card-character">Tyrone 'Clean' Miller</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8350.jpg" alt="Sam Bottoms" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Sam Bottoms</span>
    <span class="cast-card-character">Lance B. Johnson</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3087.jpg" alt="Robert Duvall" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Robert Duvall</span>
    <span class="cast-card-character">Lieutenant Colonel Bill Kilgore</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2778.jpg" alt="Dennis Hopper" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Dennis Hopper</span>
    <span class="cast-card-character">Photojournalist</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3173.jpg" alt="G. D. Spradlin" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">G. D. Spradlin</span>
    <span class="cast-card-character">General Corman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3.jpg" alt="Harrison Ford" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Harrison Ford</span>
    <span class="cast-card-character">Colonel Lucas</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/8346.jpg" alt="Jerry Ziesmer" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jerry Ziesmer</span>
    <span class="cast-card-character">Jerry, Civilian</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/349.jpg" alt="Scott Glenn" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Scott Glenn</span>
    <span class="cast-card-character">Lieutenant Richard M. Colby</span>
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
    <li><a href="../bullitt-1968">Bullitt</a> because of Robert Duvall</li>
<li><a href="../three-days-of-the-condor-1975">Three Days of the Condor</a> because of James Keane</li>
<li><a href="../blade-runner-1982">Blade Runner</a> because of Harrison Ford</li>
<li><a href="../paris-texas-1984">Paris, Texas</a> because of Aurore Clément</li>
  </ul>
</section>

</article>
