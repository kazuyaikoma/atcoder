@countries = %w[
  Brazil Croatia Mexico
  Cameroon Spain Netherlands
  Chile Australia Colombia
  Greece Coted'lvoire Japan
  Uruguay CostaRica England
  Italy Switzerland Ecuador
  France Honduras Argentina
  BosniaAndHerzegovina Iran Nigeria
  Germany Portugal Ghana
  USA Belgium Algeria
  Russia KoreaRepublic
].each(&:downcase!)

@max_depth = 0

def search(remaining, prev, depth)
  next_c = remaining.select { |c| prev[-1] == c[0] }
  if next_c.size.positive?
    next_c.each do |c|
      search(remaining - [c], c, depth + 1)
    end
  else
    @max_depth = [@max_depth, depth].max
  end
end

@countries.each do |c|
  search(@countries - [c], c, 1)
end
puts @max_depth
