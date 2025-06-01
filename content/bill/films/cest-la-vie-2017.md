---
title: "C'est la vie!"
layout: layouts/home.njk
slug: cest-la-vie-2017
ogImage: content/bill/films/backdrops/cest-la-vie-2017.jpg
description: "Max is a battle-weary veteran of the wedding-planning racket. His latest — and last — gig is a hell of a fête, involving stuffy period costumes for the caterers, a vain, hyper- sensitive singer who thinks he's a Gallic James Brown, and a morose, micromanaging groom determined to make Max's night as miserable as possible. But what makes the affair too bitter to endure is that Max's colleague and ostensible girlfriend, Joisette, seems to have written him off, coolly going about her professional duties while openly flirting with a much younger server. It's going to be a very long night… especially once the groom's aerial serenade gets underway."
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-handmaiden-2016">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../lucky-2017">Next</a>
</nav>

<p>67 / 100</p>

<article class="film slug-cest-la-vie-2017">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
  </div>

  <h1>{{ film.title }} ({{ film | filmYear }})</h1>

  <p>Also known as <strong>Le Sens de la fête</strong></p>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>

  {% if films.reviews[slug] %}
    <blockquote> 
      {{ films.reviews[slug] | safe }} <em>— Bill</em>
    </blockquote> 
  {% endif %}

  <h2>
    Cast
  </h2>
  <ul>
    {%- for cast in film.credits.cast -%}
      <li>
        {{ cast.name }} as <em>{{ cast.character }}</em>
      </li>
    {%- endfor -%}
  </ul>

  <h2>
    Crew
  </h2>
  <ul>
    {%- for crew in film.credits.crew -%}
      <li>
        {{ crew.name }} &mdash; <em>{{ crew.job }}</em>
      </li>
    {%- endfor -%}
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
