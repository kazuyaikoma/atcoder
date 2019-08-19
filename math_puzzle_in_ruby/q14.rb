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
  nexts = remaining.select { |r| r[0] == prev[-1] }
  nexts.each { |n| search(remaining - [n], n, depth + 1) }
  @max_depth = [@max_depth, depth].max if nexts.empty?
end

@countries.each { |c| search(@countries - [c], c, 1) }

puts @max_depth
