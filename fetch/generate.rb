require "json"

top_film_slugs = JSON.parse(File.read("../_data/top_films.json"))

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
    ogImage: content/bill/films/backdrops/#{film["slug"]}.jpg
    description: "#{film["overview"].gsub('"', '\"')}"
    ---

    {% set film = films[slug] %}

    <nav class="films">
      <div class=\"prev\">
        #{prev_film ? "<a href=\"../#{prev_film["slug"]}\"><i class=\"fa-solid fa-chevron-left fa-xs\"></i> Previous</a>" : "<span><i class=\"fa-solid fa-chevron-left fa-xs\"></i> Previous</span>"}
      </div>
      <div>
        <a href="../">Film list</a>
      </div>
      <div class=\"next\">
        #{next_film ? "<a href=\"../#{next_film["slug"]}\">Next <i class=\"fa-solid fa-chevron-right fa-xs\"></i></a>" : "<span>Next <i class=\"fa-solid fa-chevron-right fa-xs\"></i></span>"}
      </div>
    </nav>

    <p>#{index + 1} / #{top_films.length}</p>
    
    <article class="film slug-#{film["slug"]}">
      <div class="backdrop-and-poster">
        <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
        <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
      </div>

      <h1>{{ film.title }} ({{ film | filmYear }})</h1>

      <p>
        {%- if film.language -%}Language: {{ film.language }}.{% endif %}
        #{ "Also known as #{film["original_title"]}." if film["original_title"].downcase != film["title"].downcase }
      </p>

      <p class="director">
        Directed by <strong>{{ film | directors }}</strong>
      </p>

      {%- if films.reviews[slug] -%}
        <blockquote> 
          {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href="/bill">Bill</a></em>
        </blockquote> 
      {%- endif -%}

      <details>
        <summary>
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

    </article>
    <footer>
      <a href="../about">About this list</a>
    </footer>
  MD
end
