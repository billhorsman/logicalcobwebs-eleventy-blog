---
title: "Sweetheart"
layout: layouts/home.njk
slug: sweetheart-2021
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../dune-2021">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-french-dispatch-2021">Next</a>
</nav>

<p>84 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/sweetheart-2021.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/sweetheart-2021.jpg" alt="">
  </div>

  <h1>Sweetheart ({{ film | filmYear }})</h1>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>


  <h2>
    Cast
  </h2>
  <ul>
            <li><strong>Nell Barlow</strong> as <em>A.J.</em></li>
        <li><strong>Ella-Rae Smith</strong> as <em>Isla</em></li>
        <li><strong>Jo Hartley</strong> as <em>Tina</em></li>
        <li><strong>Sophia Di Martino</strong> as <em>Lucy</em></li>
        <li><strong>Samuel Anderson</strong> as <em>Steve</em></li>
        <li><strong>Tabitha Byron</strong> as <em>Dayna</em></li>
        <li><strong>Steffan Cennydd</strong> as <em>Nathan</em></li>
        <li><strong>Spike Fearn</strong> as <em>Elvis</em></li>
        <li><strong>William Andrews</strong> as <em>Phil the Magician</em></li>
        <li><strong>Anna Antoniades</strong> as <em>Gemma G</em></li>
        <li><strong>Celeste De Veazey</strong> as <em>Bendy Wendy</em></li>
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
