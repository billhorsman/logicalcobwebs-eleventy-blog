---
title: "Tomboy"
layout: layouts/home.njk
slug: tomboy-2011
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../micmacs-2009">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../le-havre-2011">Next</a>
</nav>

<p>59 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/tomboy-2011.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/tomboy-2011.jpg" alt="">
  </div>

  <h1>Tomboy ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>


  <h2>
    Cast
  </h2>
  <ul>
            <li><strong>Zoé Héran</strong> as <em>Laure / Mickaël</em></li>
        <li><strong>Malonn Lévana</strong> as <em>Jeanne</em></li>
        <li><strong>Jeanne Disson</strong> as <em>Lisa</em></li>
        <li><strong>Sophie Cattani</strong> as <em>La mère</em></li>
        <li><strong>Mathieu Demy</strong> as <em>Le père</em></li>
        <li><strong>Rayan Boubekri</strong> as <em>Rayan</em></li>
        <li><strong>Yohan Vero</strong> as <em>Vince</em></li>
        <li><strong>Noah Vero</strong> as <em>Noah</em></li>
        <li><strong>Cheyenne Lainé</strong> as <em>Cheyenne</em></li>
        <li><strong>Christel Baras</strong> as <em>La mère de Lisa</em></li>
        <li><strong>Valérie Roucher</strong> as <em>La mère de Rayan</em></li>
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
