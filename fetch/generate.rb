require "digest"
require "json"

class Person
  def initialize(data)
    @data = data
  end

  def name
    @data["name"]
  end

  def director?
    @data["job"] == "Director"
  end
end

class Film
  attr_reader :slug

  def initialize(slug)
    @data = JSON.parse(File.read("../_data/films/#{slug}.json"))
    @slug = slug
  end

  def title
    @data["title"]
  end

  def original_title
    return if @data["original_title"].downcase == @data["title"].downcase
    @data["original_title"]
  end

  def year
    @data["year"]
  end

  def overview
    @data["overview"]
  end

  def cast
    @data.dig("credits", "cast").map { |credit| Person.new(credit) }
  end

  def crew
    @data.dig("credits", "crew").map { |credit| Person.new(credit) }
  end

  def directors
    crew.select(&:director?)
  end
end

def to_sentence(array)
  out = ""
  array.each_with_index do |x, index|
    out << x.to_s
    if index < array.length - 2
      out << ", "
    elsif index == array.length - 2
      out << " and "
    end
  end
  out
end

top_films = JSON.parse(File.read("../_data/top_films.json")).map { Film.new(_1) }

top_cast = top_films.map { |film| 
  film.cast.slice(0, 3).map(&:name)
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
  film.directors.map(&:name)
}.flatten.tally.select { |director, count| 
  count > 1 
}.sort_by { |director, count| 
  [-count, director.downcase] 
}.map { |director, count| 
  { name: director, count: count } 
}
IO.write("../_data/top_directors.json", JSON.pretty_generate(top_directors))


credit_lookup = {}
top_films.each do |film|
  film.cast.each do |credit|
    name = credit.name
    credit_lookup[name] ||= []
    credit_lookup[name] << film unless credit_lookup[name].include?(film)
  end
end

credit_lookup.select! { |cast_name, films| films.length > 1 }

related_films = {}
credit_lookup.each do |name, films|
  films.each do |film|
    other_films = films - [film]
    related_films[film] ||= {}
    other_films.each do |other_film|
      related_films[film][other_film] ||= []
      related_films[film][other_film] << name unless related_films[film][other_film].include?(name)
    end
  end
end

top_films.each_with_index do |film, index|
  next_film = top_films[index + 1]
  prev_film = index > 0 ? top_films[index - 1] : nil

  related = related_films[film]&.map { |film, names|
    "<a href=\"../#{film.slug}\">#{film.title}</a> (#{film.year}) by #{to_sentence(names)}"  
  } || []
  
  puts "Writing #{film.slug}.md"
  path = "../content/bill/films/#{film.slug}.md"
  File.write(path, <<~MD)
    ---
    title: "#{film.title}"
    layout: layouts/films.njk
    slug: #{film.slug}
    ogImage: content/bill/films/backdrops/#{film.slug}.jpg
    description: "#{film.overview.gsub('"', '\"')}"
    ---

    {% set film = films[slug] %}

    <nav class="films">
      <div class=\"prev\">
        #{prev_film ? "<a href=\"../#{prev_film.slug}\"><i class=\"fa-solid fa-chevron-left fa-xs\"></i> Previous</a>" : "<span><i class=\"fa-solid fa-chevron-left fa-xs\"></i> Previous</span>"}
      </div>
      <div>
        <a class="simple" href="../">#{index + 1} / #{top_films.length}</a>
      </div>
      <div class=\"next\">
        #{next_film ? "<a href=\"../#{next_film.slug}\">Next <i class=\"fa-solid fa-chevron-right fa-xs\"></i></a>" : "<span>Next <i class=\"fa-solid fa-chevron-right fa-xs\"></i></span>"}
      </div>
      <div class="hint">
        <span class="prev-hint">
          <span class="sr-only">Previous film:</span>
          #{prev_film&.title || "Start of list"}
        </span>
        <span class="next-hint">
          <span class="sr-only">Next film:</span>
          #{next_film&.title || "End of list"}
        </span>
      </div>
    </nav>

    <article class="film slug-#{film.slug}">
      <div class="backdrop-and-poster">
        <img class="poster" src="../films/posters/{{ slug }}.jpg" alt="">
        <img class="backdrop" src="../films/backdrops/{{ slug }}.jpg" alt="">
      </div>

      <h1>{{ film.title }} ({{ film | filmYear }})</h1>

      <p>
        {%- if film.language -%}Language: {{ film.language }}.{% endif %}
        #{ "Also known as #{film.original_title}." if film.original_title }
      </p>

      <p class="director">
        Directed by <strong>{{ film | directors }}</strong>
      </p>

      {%- if films.reviews[slug] -%}
        <blockquote> 
          {{ films.reviews[slug] | safe }} <em>—&nbsp;<a href="/bill">Bill</a></em>
        </blockquote> 
      {%- endif -%}

      #{"<p class=\"related-films\">Related to:</p>" if related.any?}
      #{"<ul class=\"related-films\">" if related.any?}
      #{related.map { |x| "<li>#{x}</li>" }.join("\n")}
      #{"</ul>" if related.any?}

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
  MD
end
