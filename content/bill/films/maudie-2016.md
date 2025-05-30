---
title: "Maudie"
layout: layouts/home.njk
slug: maudie-2016
---

{% set film = films[slug] %}

<nav class="films">
  <a class="prev" href="../the-handmaiden-2016">Previous</a>
  <a href="../">Film list</a>
  <a class="next" href="../the-party-2017">Next</a>
</nav>

<p>67 / 100</p>

<article class="film">
  <div class="backdrop-and-poster">
    <img class="poster" src="../films/posters/maudie-2016.jpg" alt="">
    <img class="backdrop" src="../films/backdrops/maudie-2016.jpg" alt="">
  </div>

  <h1>Maudie ({{ film | filmYear }})</h1>

  

  <p class="director">
    Directed by <strong>{{ film | directors }}</strong>
  </p>


  <h2>
    Cast
  </h2>
  <ul>
            <li><strong>Sally Hawkins</strong> as <em>Maud Lewis</em></li>
        <li><strong>Ethan Hawke</strong> as <em>Everett Lewis</em></li>
        <li><strong>Gabrielle Rose</strong> as <em>Aunt Ida</em></li>
        <li><strong>Billy MacLellan</strong> as <em>Frank</em></li>
        <li><strong>Zachary Bennett</strong> as <em>Charles Dowley</em></li>
        <li><strong>Kari Matchett</strong> as <em>Sandra</em></li>
        <li><strong>David Feehan</strong> as <em>Paul</em></li>
        <li><strong>Lawrence Barry</strong> as <em>Mr. Davis (Shopkeeper)</em></li>
        <li><strong>Marthe Bernard</strong> as <em>Kay</em></li>
        <li><strong>Greg Malone</strong> as <em>Mr. Hill</em></li>
        <li><strong>Nik Sexton</strong> as <em>Steven (CBC Reporter)</em></li>
        <li><strong>Brian Marler</strong> as <em>Doctor</em></li>
        <li><strong>Denise Sinnott</strong> as <em>Hospital Nurse</em></li>
        <li><strong>Mike Daly</strong> as <em>Man at Bar</em></li>
        <li><strong>Judy Hancock</strong> as <em>Ida's Nurse</em></li>
        <li><strong>Lisa Machin</strong> as <em>Nurse (uncredited)</em></li>
        <li><strong>Erin Mick</strong> as <em></em></li>
  </ul>
</article>
<footer>
  <a href="../about">About this list</a>
</footer>
