---
title: "C'est la vie!"
layout: layouts/films.njk
slug: cest-la-vie-2017
ogImage: content/bill/films/backdrops/cest-la-vie-2017.jpg
description: "Max is a battle-weary veteran of the wedding-planning racket. His latest — and last — gig is a hell of a fête, involving stuffy period costumes for the caterers, a vain, hyper- sensitive singer who thinks he's a Gallic James Brown, and a morose, micromanaging groom determined to make Max's night as miserable as possible. But what makes the affair too bitter to endure is that Max's colleague and ostensible girlfriend, Joisette, seems to have written him off, coolly going about her professional duties while openly flirting with a much younger server. It's going to be a very long night… especially once the groom's aerial serenade gets underway."
---

{% set film = films[slug] %}

<nav class="films">
  <div class="prev">
    <a href="../the-handmaiden-2016"><i class="fa-solid fa-chevron-left fa-xs"></i> Previous</a>
  </div>
  <div>
    <a class="simple" href="../">60 / 100</a>
  </div>
  <div class="next">
    <a href="../lady-bird-2017">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Handmaiden
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Lady Bird
    </span>
  </div>
</nav>

<article class="film slug-cest-la-vie-2017">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>
    {%- if film.language -%}Language: {{ film.language }}.{% endif %}
    Also known as Le Sens de la fête.
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
  <img src="../films/profiles/28281.jpg" alt="Jean-Pierre Bacri" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Pierre Bacri</span>
    <span class="cast-card-character">Max</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/54291.jpg" alt="Gilles Lellouche" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Gilles Lellouche</span>
    <span class="cast-card-character">James</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/16922.jpg" alt="Jean-Paul Rouve" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Jean-Paul Rouve</span>
    <span class="cast-card-character">Guy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/992623.jpg" alt="Vincent Macaigne" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Vincent Macaigne</span>
    <span class="cast-card-character">Julien</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/544681.jpg" alt="Alban Ivanov" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Alban Ivanov</span>
    <span class="cast-card-character">Samy</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/149989.jpg" alt="Eye Haïdara" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Eye Haïdara</span>
    <span class="cast-card-character">Adèle</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/142689.jpg" alt="Suzanne Clément" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Suzanne Clément</span>
    <span class="cast-card-character">Josiane</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1141.jpg" alt="Hélène Vincent" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Hélène Vincent</span>
    <span class="cast-card-character">Pierre's Mother</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1159954.jpg" alt="Benjamin Lavernhe" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Benjamin Lavernhe</span>
    <span class="cast-card-character">Pierre</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/72946.jpg" alt="Judith Chemla" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Judith Chemla</span>
    <span class="cast-card-character">Héléna</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/544682.jpg" alt="William Lebghil" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">William Lebghil</span>
    <span class="cast-card-character">Seb</span>
  </div>
</div>
    <div class="cast-card">
  <img src="../films/profiles/1132641.jpg" alt="Kévin Azaïs" loading="lazy">
  <div class="cast-card-info">
    <span class="cast-card-name">Kévin Azaïs</span>
    <span class="cast-card-character">Patrice</span>
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
    <li><a href="../sink-or-swim-2018">Sink or Swim</a> because of Gilles Lellouche and Alban Ivanov</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> because of Benjamin Lavernhe</li>
  </ul>
</section>

</article>
