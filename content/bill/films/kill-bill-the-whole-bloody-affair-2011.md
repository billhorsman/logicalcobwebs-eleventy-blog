---
title: "Kill Bill: The Whole Bloody Affair"
layout: layouts/films.njk
slug: kill-bill-the-whole-bloody-affair-2011
ogImage: content/bill/films/backdrops/kill-bill-the-whole-bloody-affair-2011.jpg
description: "An assassin is shot and almost killed by her ruthless employer, Bill, and other members of their assassination circle – but she lives to plot her vengeance."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../tomboy-2011"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">55 / 100</a>
  </div>
  <div class="next">
    <a href="../all-is-lost-2013">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Tomboy
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      All Is Lost
    </span>
  </div>
</nav>

<article class="film slug-kill-bill-the-whole-bloody-affair-2011">
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
    <li><a href="../blade-runner-1982">Blade Runner</a> because of Daryl Hannah</li>
<li><a href="../magnolia-1999">Magnolia</a> because of Michael Bowen</li>
<li><a href="../phone-booth-2003">Phone Booth</a> because of Shu Lan Tuan</li>
  </ul>
</section>

</article>
