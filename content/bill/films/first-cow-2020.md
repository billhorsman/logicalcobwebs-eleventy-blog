---
title: "First Cow"
layout: layouts/home.njk
slug: first-cow-2020
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../portrait-of-a-lady-on-fire-2019">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../schemers-2020">Next</a>
</nav>

<p>75 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/first-cow-2020.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/first-cow-2020.jpg" alt="">
  </div>

  <h1>First Cow ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>


  <h2>
    Cast
  </h2>
  <ul>
            <li><strong>John Magaro</strong> as <em>Cookie</em></li>
        <li><strong>Orion Lee</strong> as <em>King-Lu</em></li>
        <li><strong>Toby Jones</strong> as <em>Chief Factor</em></li>
        <li><strong>Ewen Bremner</strong> as <em>Lloyd</em></li>
        <li><strong>Scott Shepherd</strong> as <em>Captain</em></li>
        <li><strong>Gary Farmer</strong> as <em>Totillicum</em></li>
        <li><strong>Lily Gladstone</strong> as <em>Chief Factor’s Wife</em></li>
        <li><strong>Alia Shawkat</strong> as <em>Woman with Dog</em></li>
        <li><strong>Dylan Smith</strong> as <em>Trapper Jack</em></li>
        <li><strong>Ryan Findley</strong> as <em>Trapper Dame</em></li>
        <li><strong>Manuel Rodriguez</strong> as <em>Trapper Bill</em></li>
        <li><strong>Patrick D. Green</strong> as <em>Russian Trapper</em></li>
        <li><strong>Jared Kasowski</strong> as <em>Thomas</em></li>
        <li><strong>René Auberjonois</strong> as <em>Man with Raven</em></li>
        <li><strong>Jean-Luc Boucherot</strong> as <em>Sailor in Saloon</em></li>
        <li><strong>Jeb Berrier</strong> as <em>Cribbage Player</em></li>
        <li><strong>John Keating</strong> as <em>Heckler in Saloon</em></li>
        <li><strong>Don MacEllis</strong> as <em>Brilliant William</em></li>
        <li><strong>Todd A. Robinson</strong> as <em>Fort Trapper</em></li>
        <li><strong>Kevin-Michael Moore</strong> as <em>Fort Trapper</em></li>
        <li><strong>Eric Martin Reid</strong> as <em>Fort Trapper</em></li>
        <li><strong>Ted Rooney</strong> as <em>Fort Trapper</em></li>
        <li><strong>Phelan Davis</strong> as <em>Fort Trapper</em></li>
        <li><strong>Mike Wood</strong> as <em>Fort Trapper</em></li>
        <li><strong>Sabrina Mary Morrison</strong> as <em>Totillicum's Wife</em></li>
        <li><strong>Mitchell Saddleback</strong> as <em>Chief Factor's Servant</em></li>
        <li><strong>Mary Ann Perreira</strong> as <em>Hawaiian Woman</em></li>
        <li><strong>T. Dan Hopkins</strong> as <em>Hawaiian Man</em></li>
        <li><strong>James Lee Jones</strong> as <em>Man with Canoe</em></li>
        <li><strong>James Ridley</strong> as <em>Soldier</em></li>
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
