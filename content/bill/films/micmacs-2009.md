---
title: "Micmacs"
layout: layouts/films.njk
slug: micmacs-2009
ogImage: content/bill/films/backdrops/micmacs-2009.jpg
description: "While standing in the doorway of the video shop where he works, Bazil is inadvertently shot in the head. Now homeless and jobless, he is taken in by a troupe of misfits who live in a giant mound of trash. There Bazil begins his quest for revenge against the people who produced the gun that shot him."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../fantastic-mr-fox-2009"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">51 / 100</a>
  </div>
  <div class="next">
    <a href="../le-havre-2011">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Fantastic Mr. Fox
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Le Havre
    </span>
  </div>
</nav>

<article class="film slug-micmacs-2009">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Micmacs à tire-larigot.
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
    <li><a href="../diva-1981">Diva</a> because of Dominique Pinon</li>
<li><a href="../delicatessen-1991">Delicatessen</a> because of Dominique Pinon and Jean-Pierre Jeunet</li>
<li><a href="../amlie-2001">Amélie</a> because of Dominique Pinon, Jean-Pierre Jeunet, Yolande Moreau and André Dussollier</li>
  </ul>
</section>

</article>
