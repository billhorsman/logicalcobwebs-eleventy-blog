require "json"

top_film_slugs = JSON.parse(File.read("../_data/top_films.json"))

top_films = top_film_slugs.map { |slug|
  JSON.parse(File.read("../_data/films/#{slug}.json")).tap do |film|
    film["slug"] = slug
  end
}

top_films.sort_by! { |film| [film["year"], film["title"]] }

IO.write("../_data/top_films.json", JSON.pretty_generate(top_films.map { |film| film["slug"] }))
