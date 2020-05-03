n, k = gets.split.map(&:to_i)
arr = []
k.times do
  _ = gets.to_i
  got = gets.split.map(&:to_i)
  arr += got
  arr.uniq!
end

cnt = 0
(1..n).each do |a|
  cnt += 1 unless arr.include?(a)
end

puts cnt
