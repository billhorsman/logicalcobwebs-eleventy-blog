require "json"

top_film_slugs = JSON.parse(File.read("../_data/top_films.json"))["list"]

top_films = top_film_slugs.map { |slug|
  JSON.parse(File.read("../_data/films/#{slug}.json")).tap do |film|
    film["slug"] = slug
  end
}

top_cast = top_films.map { |film| 
  film.dig("credits", "cast").slice(0, 3).map { |cast| cast["name"] }
}.flatten.tally.sort_by { |cast_name, count| 
  [-count, cast_name.downcase] 
}.map { |cast_name, count| 
  { name: cast_name, count: } 
}
(1..5).each do |i|
  if top_cast.length < 10
    break
  else
    top_cast = top_cast.select { |cast| cast[:count] > i }
  end
end
IO.write("../_data/top_cast.json", JSON.pretty_generate(top_cast))

top_directors = top_films.map { |film| 
  film.dig("credits", "crew").select { |crew| crew["job"] == "Director" }.map { |crew| crew["name"] }
}.flatten.tally.select { |director, count| 
  count > 1 
}.sort_by { |director, count| 
  [-count, director.downcase] 
}.map { |director, count| 
  { name: director, count: count } 
}
IO.write("../_data/top_directors.json", JSON.pretty_generate(top_directors))


top_films.each_with_index do |film, index|
  next_film = top_films[index + 1]
  prev_film = index > 0 ? top_films[index - 1] : nil

  puts "Wring #{film["slug"]}.md"
  path = "../content/bill/films/#{film["slug"]}.md"
  File.write(path, <<~MD)
    ---
    title: "#{film["title"]}"
    layout: layouts/home.njk
    slug: #{film["slug"]}
    ---

    {% set film = films[slug] %}

    <nav class="films">
      #{prev_film ? "<a class=\"prev\" href=\"../#{prev_film["slug"]}\">Previous</a>" : "<span class=\"prev\">Previous</span>"}
      <a href="../">Film list</a>
      #{next_film ? "<a class=\"next\" href=\"../#{next_film["slug"]}\">Next</a>" : "<span class=\"next\">Next</span>"}
    </nav>

    <p>#{index + 1} / #{top_films.length}</p>
    
    <article class="film">
      <div class="backdrop-and-poster">
        <img class="poster" src="../films/posters/#{film["slug"]}.jpg" alt="">
        <img class="backdrop" src="../films/backdrops/#{film["slug"]}.jpg" alt="">
      </div>

      <h1>#{film["title"]} ({{ film | filmYear }})</h1>

      <p class="director">
        Directed by <strong>{{ film | directors }}</strong>
      </p>


      <h2>
        Cast
      </h2>
      <ul>
        #{film.dig("credits", "cast").map { |cast| "        <li><strong>#{cast["name"]}</strong> as <em>#{cast["character"]}</em></li>" }.join("\n")}
      </ul>
    </article>
    <footer>
      <a href="../about">About this list</a>
    </footer>
  MD
end
