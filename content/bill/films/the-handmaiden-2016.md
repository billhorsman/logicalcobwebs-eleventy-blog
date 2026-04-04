---
title: "The Handmaiden"
layout: layouts/films.njk
slug: the-handmaiden-2016
ogImage: content/bill/films/backdrops/the-handmaiden-2016.jpg
description: "1930s Korea, in the period of Japanese occupation, a new girl, Sookee, is hired as a handmaiden to a Japanese heiress, Hideko, who lives a secluded life on a large countryside estate with her domineering Uncle Kouzuki. But the maid has a secret. She is a pickpocket recruited by a swindler posing as a Japanese Count to help him seduce the Lady to elope with him, rob her of her fortune, and lock her up in a madhouse. The plan seems to proceed according to plan until Sookee and Hideko discover some unexpected emotions."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../maudie-2016"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">59 / 100</a>
  </div>
  <div class="next">
    <a href="../cest-la-vie-2017">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      Maudie
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      C'est la vie!
    </span>
  </div>
</nav>

<article class="film slug-the-handmaiden-2016">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="" eleventy:ignore>
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="" eleventy:ignore>
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as 아가씨.
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
  <img src="../films/profiles/123664.jpg" alt="Kim Min-hee" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kim Min-hee</span>
    <span class="cast-card-character">Lady Hideko</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1537768.jpg" alt="Kim Tae-ri" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kim Tae-ri</span>
    <span class="cast-card-character">Sook-hee</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/75913.jpg" alt="Ha Jung-woo" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Ha Jung-woo</span>
    <span class="cast-card-character">Count Fujiwara</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/138336.jpg" alt="Cho Jin-woong" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Cho Jin-woong</span>
    <span class="cast-card-character">Uncle Kouzuki</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/93252.jpg" alt="Kim Hae-sook" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kim Hae-sook</span>
    <span class="cast-card-character">Ms. Sasaki</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/83122.jpg" alt="Moon So-ri" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Moon So-ri</span>
    <span class="cast-card-character">Hideko's Aunt</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/557743.jpg" alt="Lee Yong-nyeo" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lee Yong-nyeo</span>
    <span class="cast-card-character">Bok-soon</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/3553480.jpg" alt="Kwak Eun-jin" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Kwak Eun-jin</span>
    <span class="cast-card-character">Kkeut-dan</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1418580.jpg" alt="Lee Dong-hwi" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Lee Dong-hwi</span>
    <span class="cast-card-character">Goo-gai</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1993723.jpg" alt="Jo Eun-hyung" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Jo Eun-hyung</span>
    <span class="cast-card-character">Young Hideko</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1379307.jpg" alt="Rina Takagi" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Rina Takagi</span>
    <span class="cast-card-character">Hideko's Mother</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/2124000.jpg" alt="Han Ha-na" loading="lazy" eleventy:ignore>
  <div class="cast-card-info">
    <span class="cast-card-name">Han Ha-na</span>
    <span class="cast-card-character">Junko</span>
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
    <li><a href="../parasite-2019">Parasite</a> because of Lee Ji-hye and Ahn Seong-bong</li>
  </ul>
</section>

</article>
