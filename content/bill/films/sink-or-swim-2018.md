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
    <a class="simple" href="../">67 / 100</a>
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
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
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
