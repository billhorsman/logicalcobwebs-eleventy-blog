---
title: "The Big Lebowski"
layout: layouts/films.njk
slug: the-big-lebowski-1998
ogImage: content/bill/films/backdrops/the-big-lebowski-1998.jpg
description: "Jeffrey 'The Dude' Lebowski, a Los Angeles slacker who only wants to bowl and drink White Russians, is mistaken for another Jeffrey Lebowski, a wheelchair-bound millionaire, and finds himself dragged into a strange series of events involving nihilists, adult film producers, ferrets, errant toes, and large sums of money."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../good-will-hunting-1997"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">31 / 100</a>
  </div>
  <div class="next">
    <a href="../fight-club-1999">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Good Will Hunting
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Fight Club
    </span>
  </div>
</nav>

<article class="film slug-the-big-lebowski-1998">
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
  <img src="../films/profiles/1229.jpg" alt="Jeff Bridges" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jeff Bridges</span>
    <span class="cast-card-character">The Dude</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1230.jpg" alt="John Goodman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">John Goodman</span>
    <span class="cast-card-character">Walter Sobchak</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1231.jpg" alt="Julianne Moore" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Julianne Moore</span>
    <span class="cast-card-character">Maude Lebowski</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/884.jpg" alt="Steve Buscemi" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Steve Buscemi</span>
    <span class="cast-card-character">Donny</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1232.jpg" alt="David Huddleston" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">David Huddleston</span>
    <span class="cast-card-character">The Big Lebowski</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1233.jpg" alt="Philip Seymour Hoffman" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Philip Seymour Hoffman</span>
    <span class="cast-card-character">Brandt</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1234.jpg" alt="Tara Reid" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Tara Reid</span>
    <span class="cast-card-character">Bunny Lebowski</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1235.jpg" alt="Philip Moon" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Philip Moon</span>
    <span class="cast-card-character">Treehorn Thug</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1236.jpg" alt="Mark Pellegrino" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Mark Pellegrino</span>
    <span class="cast-card-character">Treehorn Thug</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/53.jpg" alt="Peter Stormare" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Peter Stormare</span>
    <span class="cast-card-character">Nihilist</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1237.jpg" alt="Flea" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Flea</span>
    <span class="cast-card-character">Nihilist</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1238.jpg" alt="Torsten Voges" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Torsten Voges</span>
    <span class="cast-card-character">Nihilist</span>
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
    <li><a href="../fargo-1996">Fargo</a> because of Steve Buscemi, Peter Stormare, Warren Keith and Joel Coen</li>
<li><a href="../nomadland-2021">Nomadland</a> because of Warren Keith</li>
<li><a href="../no-country-for-old-men-2007">No Country for Old Men</a> and <a href="../the-tragedy-of-macbeth-2021">The Tragedy of Macbeth</a> because of Joel Coen</li>
<li><a href="../magnolia-1999">Magnolia</a> because of Julianne Moore, Philip Seymour Hoffman and Aimee Mann</li>
<li><a href="../the-talented-mr-ripley-1999">The Talented Mr. Ripley</a> because of Philip Seymour Hoffman</li>
<li><a href="../eternal-beauty-2020">Eternal Beauty</a> because of David Thewlis</li>
  </ul>
</section>

</article>
