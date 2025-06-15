require 'uri'
require 'net/http'
require "json"

# See https://developer.themoviedb.org/docs/getting-started

# 1. Fetch all pages from TMDB
# See https://developer.themoviedb.org/reference/list-details

list_id = 8291691
list = []

def get_api_response(endpoint:, page: nil, language: "en-US", append_to_response: [])
  api_key = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3ODk1NmI5NjEzMGU5NmVhZTYwNDNjMDZiYjBjMTliNiIsIm5iZiI6MTcwODUwODMwMS4yNDg5OTk4LCJzdWIiOiI2NWQ1YzQ4ZGYyOWQ2NjAxN2RlOTRmMzgiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.TU6KG5MWHtgbsGRNDG786tlP22TY3y86-Q02sDUUkEs"

  url = URI("https://api.themoviedb.org/3/#{endpoint}?language=#{language}#{"&append_to_response=#{append_to_response.join(",")}" if append_to_response.any?}#{"&page=#{page}" if page}")

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Get.new(url)
  request["accept"] = "application/json"
  request["Authorization"] = "Bearer #{api_key}"

  http.request(request)
end

(1..10).each do |page|
  puts "Fetching page #{page}"
  response = get_api_response(endpoint: "list/#{list_id}", page:)
  items = JSON.parse(response.read_body)["items"]
  if items.any?
    list += items
  else
    puts "Page #{page} is empty"
    break
  end
end

# Sort the films by date
list.sort_by! { |film| [film["release_date"]] }

# Generate year and slug for each item
list.each do |film|
  film["year"] = film["release_date"].split("-").first
  film["slug"] = film["title"].downcase.gsub(/[^a-z0-9 ]+/, "").gsub(" ", "-") + "-" + film["year"].to_s
end

# Get the details for each film
list.each do |film|
  filename = "../_data/films/#{film["slug"]}.json"
  if !File.exist?(filename) || ARGV[0] == "--force"
    puts "Fetching details for #{film["title"]}"
    response = get_api_response(endpoint: "movie/#{film["id"]}", append_to_response: ["credits"])
    data = JSON.parse(response.read_body)
    data["slug"] = film["slug"]
    data["year"] = film["year"]
    data["language"] = data["spoken_languages"]&.detect { |language| language["iso_639_1"] == data["original_language"] }&.dig("english_name")
    IO.write(filename, JSON.pretty_generate(data))
  end
end

# 4. Download backdrop and poster images
list.each do |film|
  %w[poster backdrop].each do |type|
    if image_slug = film["#{type}_path"]
      path = "#{type}s/#{film["slug"]}.jpg"
      full_path = "../content/bill/films/#{path}"
      if !File.exist?(full_path)
        url = "https://image.tmdb.org/t/p/w500#{image_slug}"
        response = Net::HTTP.get(URI(url))
        puts "Writing #{response.length} bytes to #{path}"
        File.write(full_path, response)    
      end
      film[type] = path
    end
  end
end
