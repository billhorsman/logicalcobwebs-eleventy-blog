---
title: "Amélie"
layout: layouts/films.njk
slug: amlie-2001
ogImage: content/bill/films/backdrops/amlie-2001.jpg
description: "At a tiny Parisian café, the adorable yet painfully shy Amélie accidentally discovers a gift for helping others. Soon Amelie is spending her days as a matchmaker, guardian angel, and all-around do-gooder. But when she bumps into a handsome stranger, will she find the courage to become the star of her very own love story?"
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../billy-elliot-2000"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">38 / 100</a>
  </div>
  <div class="next">
    <a href="../black-hawk-down-2001">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Billy Elliot
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Black Hawk Down
    </span>
  </div>
</nav>

<article class="film slug-amlie-2001">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Le Fabuleux Destin d'Amélie Poulain.
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
  <img src="../films/profiles/2405.jpg" alt="Audrey Tautou" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Audrey Tautou</span>
    <span class="cast-card-character">Amélie Poulain</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2406.jpg" alt="Mathieu Kassovitz" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Mathieu Kassovitz</span>
    <span class="cast-card-character">Nino Quincampoix</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2407.jpg" alt="Rufus" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Rufus</span>
    <span class="cast-card-character">Raphaël Poulain</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2410.jpg" alt="Serge Merlin" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Serge Merlin</span>
    <span class="cast-card-character">Raymond Dufayel</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2408.jpg" alt="Jamel Debbouze" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jamel Debbouze</span>
    <span class="cast-card-character">Lucien</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2411.jpg" alt="Clotilde Mollet" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Clotilde Mollet</span>
    <span class="cast-card-character">Gina</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1654.jpg" alt="Claire Maurier" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Claire Maurier</span>
    <span class="cast-card-character">Suzanne</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2412.jpg" alt="Isabelle Nanty" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Isabelle Nanty</span>
    <span class="cast-card-character">Georgette</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2413.jpg" alt="Dominique Pinon" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Dominique Pinon</span>
    <span class="cast-card-character">Joseph</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2414.jpg" alt="Artus de Penguern" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Artus de Penguern</span>
    <span class="cast-card-character">Hipolito</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2415.jpg" alt="Yolande Moreau" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Yolande Moreau</span>
    <span class="cast-card-character">Madeleine Wallace</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2416.jpg" alt="Urbain Cancelier" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Urbain Cancelier</span>
    <span class="cast-card-character">Collignon</span>
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
    <li><a href="../diva-1981">Diva</a> because of Dominique Pinon</li>
<li><a href="../delicatessen-1991">Delicatessen</a> because of Dominique Pinon, Ticky Holgado, Rufus, Patrick Paroux and Jean-Pierre Jeunet</li>
<li><a href="../micmacs-2009">Micmacs</a> because of Dominique Pinon, Jean-Pierre Jeunet, Yolande Moreau and André Dussollier</li>
  </ul>
</section>

</article>
