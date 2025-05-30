---
title: "Portrait of a Lady on Fire"
layout: layouts/home.njk
slug: portrait-of-a-lady-on-fire-2019
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../parasite-2019">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../first-cow-2020">Next</a>
</nav>

<p>74 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/portrait-of-a-lady-on-fire-2019.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/portrait-of-a-lady-on-fire-2019.jpg" alt="">
  </div>

  <h1>Portrait of a Lady on Fire ({{ film | filmYear }})</h1>

  <p>Also known as <strong>Portrait de la jeune fille en feu</strong></p>

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>


  <h2>
    Cast
  </h2>
  <ul>
            <li><strong>Noémie Merlant</strong> as <em>Marianne</em></li>
        <li><strong>Adèle Haenel</strong> as <em>Héloïse</em></li>
        <li><strong>Luàna Bajrami</strong> as <em>Sophie</em></li>
        <li><strong>Valeria Golino</strong> as <em>La Comtesse</em></li>
        <li><strong>Christel Baras</strong> as <em>La faiseuse d'ange</em></li>
        <li><strong>Armande Boulanger</strong> as <em>L'élève atelier</em></li>
        <li><strong>Guy Delamarche</strong> as <em>L'homme salon</em></li>
        <li><strong>Clément Bouyssou</strong> as <em>Le batelier</em></li>
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
