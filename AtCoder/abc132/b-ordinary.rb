n = gets.to_i
elms = gets.split.map(&:to_i)
cnt = 0

(1...n - 1).each do |i|
  arr = [elms[i - 1], elms[i], elms[i + 1]].sort
  cnt += 1 if arr.index(elms[i]) == 1
end

puts cnt
