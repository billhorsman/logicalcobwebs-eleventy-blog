---
title: "The Grand Budapest Hotel"
layout: layouts/films.njk
slug: the-grand-budapest-hotel-2014
ogImage: content/bill/films/backdrops/the-grand-budapest-hotel-2014.jpg
description: "The Grand Budapest Hotel tells of a legendary concierge at a famous European hotel between the wars and his friendship with a young employee who becomes his trusted protégé. The story involves the theft and recovery of a priceless Renaissance painting, the battle for an enormous family fortune and the slow and then sudden upheavals that transformed Europe during the first half of the 20th century."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../mr-turner-2014"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">62 / 100</a>
  </div>
  <div class="next">
    <a href="../maudie-2016">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Mr. Turner
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Maudie
    </span>
  </div>
</nav>

<article class="film slug-the-grand-budapest-hotel-2014">
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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../fight-club-1999">Fight Club</a> (1999) by Edward Norton</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> (2021) by Edward Norton, Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Owen Wilson, Adrien Brody, Mathieu Amalric, Saoirse Ronan, Léa Seydoux, Tilda Swinton, Tony Revolori, Larry Pine, Bob Balaban and Fisher Stevens</li>
<li><a href="../asteroid-city-2023">Asteroid City</a> (2023) by Edward Norton, Jason Schwartzman, Willem Dafoe, Adrien Brody, Jeff Goldblum, Tilda Swinton, Tony Revolori, Bob Balaban and Fisher Stevens</li>
<li><a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> (1999) by Jude Law</li>
<li><a href="../in-bruges-2008">In Bruges</a> (2008) by Ralph Fiennes</li>
<li><a href="../fantastic-mr-fox-2009">Fantastic Mr. Fox</a> (2009) by Jason Schwartzman, Wallace Wolodarsky, Willem Dafoe, Bill Murray, Robin Hurlstone, Owen Wilson and Adrien Brody</li>
<li><a href="../the-lighthouse-2019">The Lighthouse</a> (2019) by Willem Dafoe</li>
<li><a href="../sink-or-swim-2018">Sink or Swim</a> (2018) by Mathieu Amalric</li>
<li><a href="../one-fine-morning-2022">One Fine Morning</a> (2022) by Léa Seydoux</li>
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
