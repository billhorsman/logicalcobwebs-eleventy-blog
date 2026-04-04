---
title: "Delicatessen"
layout: layouts/films.njk
slug: delicatessen-1991
ogImage: content/bill/films/backdrops/delicatessen-1991.jpg
description: "In a post-apocalyptic world, the residents of an apartment above the butcher shop receive an occasional delicacy of meat, something that is in low supply. A young man new in town falls in love with the butcher's daughter, which causes conflicts in her family, who need the young man for other business-related purposes."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../withnail--i-1987"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">24 / 100</a>
  </div>
  <div class="next">
    <a href="../night-on-earth-1991">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Withnail & I
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Night on Earth
    </span>
  </div>
</nav>

<article class="film slug-delicatessen-1991">
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
  <img src="../films/profiles/2413.jpg" alt="Dominique Pinon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Dominique Pinon</span>
    <span class="cast-card-character">Louison</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13686.jpg" alt="Marie-Laure Dougnac" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Marie-Laure Dougnac</span>
    <span class="cast-card-character">Julie Clapet</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13687.jpg" alt="Jean-Claude Dreyfus" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Claude Dreyfus</span>
    <span class="cast-card-character">Clapet</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13688.jpg" alt="Karin Viard" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Karin Viard</span>
    <span class="cast-card-character">Mademoiselle Plusse</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13689.jpg" alt="Ticky Holgado" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ticky Holgado</span>
    <span class="cast-card-character">Marcel Tapioca</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Pascal Benezech</span>
    <span class="cast-card-character">Tried to Escape</span>
  </div>
</div>
    <div class="cast-card">
  <div class="cast-card-no-image"><i class="fa-solid fa-user"></i></div>
  <div class="cast-card-info">
    <span class="cast-card-name">Edith Ker</span>
    <span class="cast-card-character">Grandmother</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2407.jpg" alt="Rufus" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rufus</span>
    <span class="cast-card-character">Robert Kube</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13692.jpg" alt="Jacques Mathou" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jacques Mathou</span>
    <span class="cast-card-character">Roger Kube</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13693.jpg" alt="Chick Ortega" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Chick Ortega</span>
    <span class="cast-card-character">Postman</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13694.jpg" alt="Jean-François Perrier" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-François Perrier</span>
    <span class="cast-card-character">Georges Interligator</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/13695.jpg" alt="Silvie Laguna" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Silvie Laguna</span>
    <span class="cast-card-character">Aurore Interligator</span>
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
<li><a href="../amlie-2001">Amélie</a> because of Dominique Pinon, Ticky Holgado, Rufus, Patrick Paroux and Jean-Pierre Jeunet</li>
<li><a href="../micmacs-2009">Micmacs</a> because of Dominique Pinon and Jean-Pierre Jeunet</li>
<li><a href="../black-hawk-down-2001">Black Hawk Down</a> because of Nikky Smedley</li>
  </ul>
</section>

</article>
