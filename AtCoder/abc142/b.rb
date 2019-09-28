n, k = gets.split.map(&:to_i)
hs = gets.split.map(&:to_i)

cnt = 0

hs.each do |h|
  cnt += 1 if h >= k
end

puts cnt
