---
title: "Sink or Swim"
layout: layouts/films.njk
slug: sink-or-swim-2018
ogImage: content/bill/films/backdrops/sink-or-swim-2018.jpg
description: "40-year-old Bertrand has been suffering from depression for the last two years and is barely able to keep his head above water. Despite the medication he gulps down all day, every day, and his wife's encouragement, he is unable to find any meaning in his life. Curiously, he will end up finding this sense of purpose at the swimming pool, by joining an all-male synchronised swimming team."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../roma-2018"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">65 / 100</a>
  </div>
  <div class="next">
    <a href="../woman-at-war-2018">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Roma
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Woman at War
    </span>
  </div>
</nav>

<article class="film slug-sink-or-swim-2018">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Le Grand Bain.
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
  <img src="../films/profiles/8789.jpg" alt="Mathieu Amalric" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Mathieu Amalric</span>
    <span class="cast-card-character">Bertrand</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/19866.jpg" alt="Guillaume Canet" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Guillaume Canet</span>
    <span class="cast-card-character">Laurent</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/47820.jpg" alt="Benoît Poelvoorde" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Benoît Poelvoorde</span>
    <span class="cast-card-character">Marcus</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/7037.jpg" alt="Jean-Hugues Anglade" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Hugues Anglade</span>
    <span class="cast-card-character">Simon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/118178.jpg" alt="Virginie Efira" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Virginie Efira</span>
    <span class="cast-card-character">Delphine</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/23383.jpg" alt="Leïla Bekhti" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Leïla Bekhti</span>
    <span class="cast-card-character">Amanda</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/76820.jpg" alt="Marina Foïs" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Marina Foïs</span>
    <span class="cast-card-character">Claire</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/119501.jpg" alt="Philippe Katerine" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Philippe Katerine</span>
    <span class="cast-card-character">Thierry</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/930288.jpg" alt="Félix Moati" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Félix Moati</span>
    <span class="cast-card-character">John</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/544681.jpg" alt="Alban Ivanov" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Alban Ivanov</span>
    <span class="cast-card-character">Basile</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3144543.jpg" alt="Balasingham Thamilchelvan" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Balasingham Thamilchelvan</span>
    <span class="cast-card-character">Avanish</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/17499.jpg" alt="Jonathan Zaccaï" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jonathan Zaccaï</span>
    <span class="cast-card-character">Thibault</span>
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
    <li><a href="../lon-the-professional-1994">Léon: The Professional</a> because of Jean-Hugues Anglade</li>
<li><a href="../the-grand-budapest-hotel-2014">The Grand Budapest Hotel</a> because of Mathieu Amalric</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Mathieu Amalric and Félix Moati</li>
<li><a href="../cest-la-vie-2017">C'est la vie!</a> because of Gilles Lellouche and Alban Ivanov</li>
  </ul>
</section>

</article>
