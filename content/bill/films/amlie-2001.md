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
    <a class="simple" href="../">42 / 100</a>
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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../diva-1981">Diva</a> (1981) by Dominique Pinon</li>
<li><a href="../delicatessen-1991">Delicatessen</a> (1991) by Dominique Pinon, Ticky Holgado, Rufus and Patrick Paroux</li>
<li><a href="../micmacs-2009">Micmacs</a> (2009) by Dominique Pinon, Yolande Moreau and André Dussollier</li>
  </ul>

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
</article>
