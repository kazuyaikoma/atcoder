n = gets.to_i
odd = 0

(1..n).each do |i|
  odd += 1 if i.odd?
end

puts odd.to_f / n.to_f
