---
title: "All Is Lost"
layout: layouts/home.njk
slug: all-is-lost-2013
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../le-havre-2011">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../dallas-buyers-club-2013">Next</a>
</nav>

<p>61 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/all-is-lost-2013.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/all-is-lost-2013.jpg" alt="">
  </div>

  <h1>All Is Lost ({{ film | filmYear }})</h1>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>


  <h2>
    Cast
  </h2>
  <ul>
            <li><strong>Robert Redford</strong> as <em>Our Man</em></li>
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
