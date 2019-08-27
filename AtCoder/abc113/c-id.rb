City = Struct.new(:prefecture, :year, :serial, :idx)

n, m = gets.split.map(&:to_i)
cities = []
(1..m).each do |i|
  prefecture, year = gets.split.map(&:to_i)
  cities << City.new(prefecture, year, '0', i)
end

prefectures = cities.group_by(&:prefecture)
prefectures.keys.each do |pi|
  cities = prefectures[pi]&.sort_by(&:year)
  next if cities.nil?

  cities.map.with_index do |city, idx|
    first = format('%06d', pi)
    last = format('%06d', idx + 1)
    city.serial = first + last
  end
end

puts prefectures.values.flatten.sort_by(&:idx).map(&:serial)
