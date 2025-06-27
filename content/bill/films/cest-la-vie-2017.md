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
    <a class="simple" href="../">65 / 100</a>
  </div>
  <div class="next">
    <a href="../lucky-2017">Next <i class="fa-solid fa-chevron-right fa-xs"></i></a>
  </div>
  <div class="hint">
    <span class="prev-hint">
      <span class="sr-only">Previous film:</span>
      The Handmaiden
    </span>
    <span class="next-hint">
      <span class="sr-only">Next film:</span>
      Lucky
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

  <p class="related-films">Related to:</p>
  <ul class="related-films">
  <li><a href="../sink-or-swim-2018">Sink or Swim</a> by Alban Ivanov</li>
<li><a href="../the-french-dispatch-2021">The French Dispatch</a> by Benjamin Lavernhe</li>
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
