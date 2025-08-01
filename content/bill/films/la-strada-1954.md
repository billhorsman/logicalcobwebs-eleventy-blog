---
title: "La Strada"
layout: layouts/films.njk
slug: la-strada-1954
ogImage: content/bill/films/backdrops/la-strada-1954.jpg
description: "When Gelsomina, a naïve young woman, is purchased from her impoverished mother by brutish circus strongman Zampanò to be his wife and partner, she loyally endures her husband's coldness and abuse as they travel the Italian countryside performing together. Soon Zampanò must deal with his jealousy and conflicted feelings about Gelsomina when she finds a kindred spirit in Il Matto, the carefree circus fool, and contemplates leaving Zampanò."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../its-a-wonderful-life-1946"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">2 / 100</a>
  </div>
  <div class="next">
    <a href="../rear-window-1954">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      It's a Wonderful Life
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Rear Window
    </span>
  </div>
</nav>

<article class="film slug-la-strada-1954">
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
    <li><a href="../being-there-1979">Being There</a> because of Richard Basehart</li>
<li><a href="../la-dolce-vita-1960">La Dolce Vita</a> because of Federico Fellini</li>
  </ul>
</section>

</article>
